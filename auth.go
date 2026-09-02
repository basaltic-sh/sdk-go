package basaltic

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

// TokenSource supplies the bearer token sent on every request.
//
// Implement it to source tokens from somewhere the SDK does not know about —
// a secrets manager, a sidecar, an instance metadata service. Token is called
// on every request and must be safe for concurrent use, so an implementation
// that fetches remotely should cache.
type TokenSource interface {
	// Token returns a token that is valid now. Returning an expired token is
	// an error, not a nil-error result: the SDK cannot tell the difference
	// between "expired" and "refused" once the platform answers 401.
	Token(ctx context.Context) (string, error)
}

// StaticTokenSource returns a TokenSource that always presents the same token.
//
// Use it when something upstream already obtained a token — an assumed-role
// session, or a token minted by an in-VM agent. The SDK cannot refresh a
// static token, so a process outliving the token's expiry will start seeing
// authentication failures; use [ClientCredentials] where that matters.
func StaticTokenSource(token string) TokenSource {
	return staticTokenSource(token)
}

type staticTokenSource string

func (s staticTokenSource) Token(context.Context) (string, error) {
	if s == "" {
		return "", fmt.Errorf("basaltic: empty access token")
	}
	return string(s), nil
}

// tokenRefreshSkew re-exchanges this long before a token expires. A token
// that lapses mid-request yields a 401 that reads like an authorization
// failure, which is a much worse outcome than one extra exchange.
const tokenRefreshSkew = 5 * time.Minute

// defaultTokenLifetime is assumed when the token endpoint reports no expiry.
// Out of spec, but assuming a short life is the safe direction: the cost is
// an extra exchange, not a stale token.
const defaultTokenLifetime = 15 * time.Minute

// ClientCredentialsSource exchanges a service account's access key pair for
// bearer tokens, refreshing before expiry.
//
// This is the ordinary way to authenticate. The key pair remains the one
// long-lived credential the service account has; the SDK presents a token
// derived from it rather than the pair itself.
//
// Tokens are cached in memory for the life of the source, so a long-running
// process performs roughly one exchange per token lifetime (an hour by
// default) rather than one per request. Concurrent callers arriving while an
// exchange is in flight wait for that exchange instead of starting their own.
type ClientCredentialsSource struct {
	// AccessKeyID and SecretAccessKey are the service account's key pair.
	AccessKeyID     string
	SecretAccessKey string

	// TokenURL is the IAM token endpoint. Empty means the endpoint derived
	// from the Config that owns this source.
	TokenURL string

	// Duration optionally requests a token lifetime. The platform accepts
	// 15 minutes to 12 hours and clamps anything outside that range rather
	// than refusing it. Zero takes the platform default of one hour.
	Duration time.Duration

	// HTTPClient performs the exchange. Nil uses the Config's client.
	HTTPClient *http.Client

	mu      sync.Mutex
	token   string
	expires time.Time
	// inflight is the exchange currently running, if any, so that N
	// concurrent first requests cost one exchange rather than N.
	inflight *tokenExchange
}

// tokenExchange is one in-flight exchange awaited by every caller that
// arrived while it was running.
type tokenExchange struct {
	done  chan struct{}
	token string
	err   error
}

// Token returns a cached token, exchanging a fresh one when the cache holds
// nothing usable.
func (s *ClientCredentialsSource) Token(ctx context.Context) (string, error) {
	s.mu.Lock()
	if s.token != "" && time.Now().Before(s.expires.Add(-tokenRefreshSkew)) {
		tok := s.token
		s.mu.Unlock()
		return tok, nil
	}
	if ex := s.inflight; ex != nil {
		s.mu.Unlock()
		select {
		case <-ex.done:
			return ex.token, ex.err
		case <-ctx.Done():
			return "", ctx.Err()
		}
	}
	ex := &tokenExchange{done: make(chan struct{})}
	s.inflight = ex
	s.mu.Unlock()

	token, expires, err := s.exchange(ctx)

	s.mu.Lock()
	s.inflight = nil
	if err == nil {
		s.token, s.expires = token, expires
	}
	s.mu.Unlock()

	ex.token, ex.err = token, err
	close(ex.done)
	return token, err
}

// Invalidate drops the cached token, forcing the next call to exchange a new
// one. The SDK calls this itself when the platform rejects a token that had
// not yet expired — a revoked session, say — so that one 401 costs one retry
// rather than failing every request until the cached token lapses.
func (s *ClientCredentialsSource) Invalidate() {
	s.mu.Lock()
	s.token, s.expires = "", time.Time{}
	s.mu.Unlock()
}

// oauthTokenResponse is the RFC 6749 success body.
type oauthTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// oauthErrorResponse is the RFC 6749 error body. The token endpoint answers
// in this shape rather than the platform envelope, deliberately: every OAuth
// client library parses this and nothing else.
type oauthErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func (s *ClientCredentialsSource) exchange(ctx context.Context) (string, time.Time, error) {
	if s.AccessKeyID == "" || s.SecretAccessKey == "" {
		return "", time.Time{}, fmt.Errorf("basaltic: no access key configured")
	}
	if s.TokenURL == "" {
		return "", time.Time{}, fmt.Errorf("basaltic: no token endpoint configured")
	}

	form := url.Values{"grant_type": {"client_credentials"}}
	if s.Duration > 0 {
		form.Set("duration_seconds", strconv.Itoa(int(s.Duration.Seconds())))
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.TokenURL, strings.NewReader(form.Encode()))
	if err != nil {
		return "", time.Time{}, err
	}
	// HTTP Basic is what RFC 6749 prefers and what the platform checks first.
	req.SetBasicAuth(s.AccessKeyID, s.SecretAccessKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	hc := s.HTTPClient
	if hc == nil {
		hc = http.DefaultClient
	}
	resp, err := hc.Do(req)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("basaltic: token exchange: %w", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(io.LimitReader(resp.Body, 1<<20))

	if resp.StatusCode != http.StatusOK {
		return "", time.Time{}, tokenExchangeError(resp.StatusCode, body)
	}
	var out oauthTokenResponse
	if err := json.Unmarshal(body, &out); err != nil {
		return "", time.Time{}, fmt.Errorf("basaltic: token exchange: malformed response: %w", err)
	}
	if out.AccessToken == "" {
		return "", time.Time{}, fmt.Errorf("basaltic: token exchange: response carried no access_token")
	}
	lifetime := time.Duration(out.ExpiresIn) * time.Second
	if lifetime <= 0 {
		lifetime = defaultTokenLifetime
	}
	return out.AccessToken, time.Now().Add(lifetime), nil
}

// AuthError is a failure to obtain a token.
//
// The distinction it preserves is the one that decides what to do next.
// "invalid_client" means the key was refused, and the fix is to check or
// rotate it. "invalid_grant" means the key is fine and the organization is
// suspended or still onboarding, where rotating a working key wastes an
// afternoon.
type AuthError struct {
	// StatusCode is the HTTP status from the token endpoint.
	StatusCode int
	// Code is the RFC 6749 error code, such as "invalid_client".
	Code string
	// Description is the endpoint's explanation, where it gave one.
	Description string
}

func (e *AuthError) Error() string {
	switch e.Code {
	case "invalid_client":
		return "basaltic: authentication failed: the access key was rejected — check or rotate the credential"
	case "invalid_grant":
		d := e.Description
		if d == "" {
			d = "the organization is not currently active"
		}
		return fmt.Sprintf("basaltic: authentication refused: %s (the access key itself is fine)", d)
	}
	if e.Description != "" {
		return fmt.Sprintf("basaltic: token exchange failed: %s: %s", e.Code, e.Description)
	}
	if e.Code != "" {
		return fmt.Sprintf("basaltic: token exchange failed: %s", e.Code)
	}
	return fmt.Sprintf("basaltic: token exchange failed: http %d", e.StatusCode)
}

// Temporary reports whether retrying the exchange could succeed without any
// change to the credential.
func (e *AuthError) Temporary() bool {
	return e.Code == "temporarily_unavailable" || e.StatusCode >= 500
}

func tokenExchangeError(status int, body []byte) error {
	e := &AuthError{StatusCode: status}
	var oe oauthErrorResponse
	if err := json.Unmarshal(body, &oe); err == nil && oe.Error != "" {
		e.Code = oe.Error
		e.Description = oe.ErrorDescription
	} else if trimmed := strings.TrimSpace(string(body)); trimmed != "" {
		if len(trimmed) > 300 {
			trimmed = trimmed[:300]
		}
		e.Description = trimmed
	}
	return e
}
