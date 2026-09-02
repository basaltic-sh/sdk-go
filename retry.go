package basaltic

import (
	"context"
	"crypto/rand"
	"errors"
	"math/big"
	"net/http"
	"strconv"
	"time"
)

// RetryConfig governs how a failed request is retried.
//
// The default policy retries transport failures and the statuses that mean
// "try again" — 429, 500, 502, 503, 504 — with exponential backoff and full
// jitter, honouring Retry-After when the platform sends one.
//
// It will not retry a request that is not safe to repeat. GET, HEAD, PUT and
// DELETE are; POST is only retried when the call carries an idempotency key,
// because a create that timed out may well have succeeded. Pass
// [WithIdempotencyKey] on a create to make it retryable — the platform then
// returns the original outcome instead of creating a second resource.
type RetryConfig struct {
	// MaxAttempts is the total number of attempts, including the first. Zero
	// or one disables retries.
	MaxAttempts int
	// BaseDelay is the backoff for the first retry; each subsequent retry
	// doubles it, up to MaxDelay. Actual delays are jittered.
	BaseDelay time.Duration
	// MaxDelay caps a single backoff.
	MaxDelay time.Duration
	// Retryable overrides which outcomes are retried. resp is nil when the
	// request never got an answer, in which case err says why. Leave nil for
	// the default policy.
	Retryable func(req *http.Request, resp *http.Response, err error) bool
}

// DefaultRetryConfig is the policy applied when none is configured: four
// attempts, backing off from 200ms to at most 20s.
func DefaultRetryConfig() RetryConfig {
	return RetryConfig{
		MaxAttempts: 4,
		BaseDelay:   200 * time.Millisecond,
		MaxDelay:    20 * time.Second,
	}
}

// shouldRetry decides whether another attempt is worth making.
func (rc RetryConfig) shouldRetry(req *http.Request, resp *http.Response, err error) bool {
	if rc.Retryable != nil {
		return rc.Retryable(req, resp, err)
	}
	if !methodIsRepeatable(req) {
		return false
	}
	if err != nil {
		// A request whose context was cancelled or whose deadline passed is
		// not a transport hiccup — the caller asked for it to stop.
		return !errors.Is(err, context.Canceled) && !errors.Is(err, context.DeadlineExceeded)
	}
	switch resp.StatusCode {
	case http.StatusTooManyRequests,
		http.StatusInternalServerError,
		http.StatusBadGateway,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout:
		return true
	}
	return false
}

// methodIsRepeatable reports whether sending this request twice is safe.
//
// The idempotency-key exemption is the whole reason POST is not simply
// excluded: a create that timed out may already have succeeded, and retrying
// it blindly is how duplicate resources happen. With a key the platform
// replays the original outcome, so the retry is safe.
func methodIsRepeatable(req *http.Request) bool {
	switch req.Method {
	case http.MethodGet, http.MethodHead, http.MethodOptions,
		http.MethodPut, http.MethodDelete:
		return true
	case http.MethodPost, http.MethodPatch:
		return req.Header.Get(idempotencyHeader) != ""
	}
	return false
}

// backoff returns how long to wait before attempt n (1-based, so the wait
// before the first retry is backoff(1)).
//
// Full jitter: a uniform draw from [0, exponential]. It spreads a fleet that
// failed together instead of having it retry in lockstep, which is how a
// recovering service gets knocked back over.
func (rc RetryConfig) backoff(attempt int, resp *http.Response) time.Duration {
	// The platform's own answer wins. Retrying before Retry-After is refused
	// and extends the window, so guessing shorter is worse than waiting.
	if resp != nil {
		if d, ok := parseRetryAfter(resp.Header.Get("Retry-After")); ok {
			if d > rc.MaxDelay {
				return rc.MaxDelay
			}
			return d
		}
	}
	base := rc.BaseDelay
	if base <= 0 {
		base = 200 * time.Millisecond
	}
	maxDelay := rc.MaxDelay
	if maxDelay <= 0 {
		maxDelay = 20 * time.Second
	}
	d := base << min(attempt-1, 16)
	if d > maxDelay || d <= 0 {
		d = maxDelay
	}
	return time.Duration(randInt64(int64(d)) + 1)
}

// randInt64 returns a uniform value in [0, n). Uses crypto/rand so the SDK
// neither seeds nor perturbs the program's math/rand stream.
func randInt64(n int64) int64 {
	if n <= 1 {
		return 0
	}
	v, err := rand.Int(rand.Reader, big.NewInt(n))
	if err != nil {
		// A failing entropy source should not turn a retryable request into a
		// hard error. Fall back to the full delay: unjittered, still correct.
		return n - 1
	}
	return v.Int64()
}

// parseRetryAfter reads a Retry-After header in either RFC 7231 form: a
// delay in seconds, or an HTTP date.
func parseRetryAfter(v string) (time.Duration, bool) {
	if v == "" {
		return 0, false
	}
	if secs, err := strconv.Atoi(v); err == nil {
		if secs < 0 {
			return 0, false
		}
		return time.Duration(secs) * time.Second, true
	}
	if t, err := http.ParseTime(v); err == nil {
		if d := time.Until(t); d > 0 {
			return d, true
		}
		return 0, true
	}
	return 0, false
}

func atoiHeader(v string) (int, error) { return strconv.Atoi(v) }
