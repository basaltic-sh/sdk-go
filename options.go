package basaltic

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// idempotencyHeader is the platform's replay-safety key.
const idempotencyHeader = "Idempotency-Key"

// accountHeader selects the account a request acts on.
const accountHeader = "X-Account-Id"

// RequestOption adjusts one call.
//
// Every generated operation takes a trailing ...RequestOption, which is what
// keeps the signatures additive: a new per-call knob arrives as an option
// rather than as a changed method signature.
type RequestOption func(*requestSettings) error

// requestSettings is the per-call state options accumulate into.
type requestSettings struct {
	header         http.Header
	query          url.Values
	accountID      string
	accountIDSet   bool
	timeout        time.Duration
	noRetry        bool
	responseHeader *http.Header
	stream         bool
}

func (rs *requestSettings) apply(opts []RequestOption) error {
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		if err := opt(rs); err != nil {
			return err
		}
	}
	return nil
}

// WithIdempotencyKey makes a create replay-safe.
//
// Retrying a request with the same key returns the original outcome verbatim
// instead of creating a second resource. Reusing a key with a different body
// is rejected; a key whose request is still in flight answers 409. Records
// are honoured for 24 hours.
//
// It also makes the call retryable: without a key the SDK will not repeat a
// POST, because a create that timed out may already have succeeded.
//
// Use [NewIdempotencyKey] if you have no natural key of your own.
func WithIdempotencyKey(key string) RequestOption {
	return func(rs *requestSettings) error {
		if key == "" {
			return fmt.Errorf("basaltic: WithIdempotencyKey needs a key")
		}
		rs.header.Set(idempotencyHeader, key)
		return nil
	}
}

// NewIdempotencyKey returns a random key suitable for [WithIdempotencyKey].
//
// Generate it once and reuse it across retries of the same logical create —
// a key generated per attempt defeats the purpose.
func NewIdempotencyKey() string {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		// Falling back to a timestamp keeps the call working. It is weaker
		// than random, but a key that only has to be unique to one caller's
		// retries does not need to be unguessable.
		return fmt.Sprintf("ik-%d", time.Now().UnixNano())
	}
	// RFC 4122 version 4 layout, so the key reads as a UUID like every other
	// identifier the platform hands back.
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	h := hex.EncodeToString(b[:])
	return h[0:8] + "-" + h[8:12] + "-" + h[12:16] + "-" + h[16:20] + "-" + h[20:32]
}

// WithRequestHeader sets a header on this call only.
func WithRequestHeader(name, value string) RequestOption {
	return func(rs *requestSettings) error {
		rs.header.Set(name, value)
		return nil
	}
}

// WithRequestAccountID overrides the config's account for this call,
// selecting which account the request acts on.
//
// Pass an empty string to send no account header at all, making the platform
// use the credential's own account.
func WithRequestAccountID(accountID string) RequestOption {
	return func(rs *requestSettings) error {
		rs.accountID, rs.accountIDSet = accountID, true
		return nil
	}
}

// WithQueryParam adds a query parameter to this call, for a filter the
// generated signature does not yet carry.
func WithQueryParam(name, value string) RequestOption {
	return func(rs *requestSettings) error {
		rs.query.Add(name, value)
		return nil
	}
}

// WithRequestTimeout bounds this call, including its retries. It layers on
// top of the context: whichever deadline is sooner applies.
func WithRequestTimeout(d time.Duration) RequestOption {
	return func(rs *requestSettings) error {
		rs.timeout = d
		return nil
	}
}

// WithoutRequestRetry sends this call exactly once, whatever the config's
// retry policy says.
func WithoutRequestRetry() RequestOption {
	return func(rs *requestSettings) error {
		rs.noRetry = true
		return nil
	}
}

// WithResponseHeader copies the response headers into h once the call
// completes. Use it to read rate-limit state, or the request id of a
// successful call.
func WithResponseHeader(h *http.Header) RequestOption {
	return func(rs *requestSettings) error {
		rs.responseHeader = h
		return nil
	}
}
