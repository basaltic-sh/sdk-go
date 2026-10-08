package basaltic

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// tokenServer stands up an OAuth token endpoint that hands out a fresh token
// on every exchange and counts how many it performed.
func tokenServer(t *testing.T, lifetime int) (*httptest.Server, *int32) {
	t.Helper()
	var exchanges int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := atomic.AddInt32(&exchanges, 1)
		user, pass, ok := r.BasicAuth()
		if !ok || user != "key-1" || pass != "secret-1" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, `{"error":"invalid_client","error_description":"client authentication failed"}`)
			return
		}
		if err := r.ParseForm(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if got := r.Form.Get("grant_type"); got != "client_credentials" {
			t.Errorf("grant_type = %q, want client_credentials", got)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"access_token":"token-%d","token_type":"Bearer","expires_in":%d}`, n, lifetime)
	}))
	t.Cleanup(srv.Close)
	return srv, &exchanges
}

func TestClientCredentialsExchangesOnceAndCaches(t *testing.T) {
	tokenSrv, exchanges := tokenServer(t, 3600)

	var seen []string
	var mu sync.Mutex
	api := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		seen = append(seen, r.Header.Get("Authorization"))
		mu.Unlock()
		fmt.Fprint(w, `{}`)
	}))
	defer api.Close()

	cfg, err := NewConfig(context.Background(),
		WithClientCredentials("key-1", "secret-1"),
		WithTokenEndpoint(tokenSrv.URL),
		WithServiceEndpoint("compute", api.URL),
		WithRegion("sa-saopaulo-1"),
	)
	if err != nil {
		t.Fatalf("NewConfig: %v", err)
	}
	c := NewClient(cfg, "compute")

	for i := 0; i < 3; i++ {
		if err := c.Do(context.Background(), &Operation{ID: "op", Method: "GET", Path: "/v1/x"}, nil); err != nil {
			t.Fatalf("Do %d: %v", i, err)
		}
	}
	if got := atomic.LoadInt32(exchanges); got != 1 {
		t.Errorf("performed %d token exchanges for 3 requests, want 1", got)
	}
	mu.Lock()
	defer mu.Unlock()
	for i, a := range seen {
		if a != "Bearer token-1" {
			t.Errorf("request %d sent %q, want the cached Bearer token-1", i, a)
		}
	}
}

func TestClientCredentialsRefreshesBeforeExpiry(t *testing.T) {
	// A token whose whole life is shorter than the refresh skew is always
	// due for replacement, which is how the skew is exercised without
	// waiting an hour.
	tokenSrv, exchanges := tokenServer(t, 60)
	api := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{}`)
	}))
	defer api.Close()

	cfg, err := NewConfig(context.Background(),
		WithClientCredentials("key-1", "secret-1"),
		WithTokenEndpoint(tokenSrv.URL),
		WithServiceEndpoint("compute", api.URL),
	)
	if err != nil {
		t.Fatalf("NewConfig: %v", err)
	}
	c := NewClient(cfg, "compute")
	for i := 0; i < 2; i++ {
		if err := c.Do(context.Background(), &Operation{ID: "op", Method: "GET", Path: "/v1/x"}, nil); err != nil {
			t.Fatalf("Do %d: %v", i, err)
		}
	}
	if got := atomic.LoadInt32(exchanges); got != 2 {
		t.Errorf("performed %d exchanges, want one per request for a token inside the refresh skew", got)
	}
}

func TestConcurrentFirstRequestsShareOneExchange(t *testing.T) {
	var exchanges int32
	tokenSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&exchanges, 1)
		// Hold the exchange open so every goroutine is certain to arrive
		// while it is in flight.
		time.Sleep(50 * time.Millisecond)
		fmt.Fprint(w, `{"access_token":"shared","token_type":"Bearer","expires_in":3600}`)
	}))
	defer tokenSrv.Close()

	src := &ClientCredentialsSource{
		AccessKeyID:     "key-1",
		SecretAccessKey: "secret-1",
		TokenURL:        tokenSrv.URL,
	}

	const n = 20
	var wg sync.WaitGroup
	tokens := make([]string, n)
	errs := make([]error, n)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			tokens[i], errs[i] = src.Token(context.Background())
		}(i)
	}
	wg.Wait()

	for i, err := range errs {
		if err != nil {
			t.Fatalf("goroutine %d: %v", i, err)
		}
		if tokens[i] != "shared" {
			t.Errorf("goroutine %d got token %q, want shared", i, tokens[i])
		}
	}
	if got := atomic.LoadInt32(&exchanges); got != 1 {
		t.Errorf("%d goroutines caused %d exchanges, want 1", n, got)
	}
}

func TestTokenExchangeErrorsAreDistinguishable(t *testing.T) {
	tests := []struct {
		name        string
		status      int
		body        string
		wantIn      string
		wantTempora bool
	}{
		{
			name:   "invalid_client says to rotate the key",
			status: http.StatusUnauthorized,
			body:   `{"error":"invalid_client","error_description":"client authentication failed"}`,
			wantIn: "rotate the credential",
		},
		{
			// The key is fine here; telling someone to rotate it would waste
			// their afternoon.
			name:   "invalid_grant says the key is fine",
			status: http.StatusForbidden,
			body:   `{"error":"invalid_grant","error_description":"Organization is suspended"}`,
			wantIn: "the access key itself is fine",
		},
		{
			name:        "temporarily_unavailable is retryable",
			status:      http.StatusServiceUnavailable,
			body:        `{"error":"temporarily_unavailable","error_description":"try again shortly"}`,
			wantIn:      "temporarily_unavailable",
			wantTempora: true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tc.status)
				fmt.Fprint(w, tc.body)
			}))
			defer srv.Close()

			src := &ClientCredentialsSource{AccessKeyID: "k", SecretAccessKey: "s", TokenURL: srv.URL}
			_, err := src.Token(context.Background())
			if err == nil {
				t.Fatal("Token() succeeded, want an error")
			}
			if !contains(err.Error(), tc.wantIn) {
				t.Errorf("Error() = %q, want it to contain %q", err.Error(), tc.wantIn)
			}
			var authErr *AuthError
			if !asAuthError(err, &authErr) {
				t.Fatalf("error is %T, want *AuthError", err)
			}
			if authErr.Temporary() != tc.wantTempora {
				t.Errorf("Temporary() = %v, want %v", authErr.Temporary(), tc.wantTempora)
			}
		})
	}
}

func TestRevokedTokenIsRefreshedOnce(t *testing.T) {
	// A session revoked before its token expires makes the platform answer
	// 401 on a token the SDK still believes is good. Without invalidating the
	// cache, every later request would present the same dead token.
	tokenSrv, exchanges := tokenServer(t, 3600)

	var apiCalls int32
	api := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := atomic.AddInt32(&apiCalls, 1)
		if n == 1 {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, `{"error":{"code":"UNAUTHORIZED","message":"session revoked","request_id":"r"}}`)
			return
		}
		if got := r.Header.Get("Authorization"); got != "Bearer token-2" {
			t.Errorf("retry sent %q, want a freshly exchanged Bearer token-2", got)
		}
		fmt.Fprint(w, `{}`)
	}))
	defer api.Close()

	cfg, err := NewConfig(context.Background(),
		WithClientCredentials("key-1", "secret-1"),
		WithTokenEndpoint(tokenSrv.URL),
		WithServiceEndpoint("compute", api.URL),
	)
	if err != nil {
		t.Fatalf("NewConfig: %v", err)
	}
	c := NewClient(cfg, "compute")
	if err := c.Do(context.Background(), &Operation{ID: "op", Method: "GET", Path: "/v1/x"}, nil); err != nil {
		t.Fatalf("Do: %v", err)
	}
	if got := atomic.LoadInt32(exchanges); got != 2 {
		t.Errorf("performed %d exchanges, want 2 — the refused token should be re-exchanged once", got)
	}
	if got := atomic.LoadInt32(&apiCalls); got != 2 {
		t.Errorf("api saw %d calls, want 2", got)
	}
}

func TestPersistentUnauthorizedIsNotRetriedForever(t *testing.T) {
	tokenSrv, _ := tokenServer(t, 3600)
	var apiCalls int32
	api := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&apiCalls, 1)
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, `{"error":{"code":"UNAUTHORIZED","message":"nope","request_id":"r"}}`)
	}))
	defer api.Close()

	cfg, err := NewConfig(context.Background(),
		WithClientCredentials("key-1", "secret-1"),
		WithTokenEndpoint(tokenSrv.URL),
		WithServiceEndpoint("compute", api.URL),
	)
	if err != nil {
		t.Fatalf("NewConfig: %v", err)
	}
	c := NewClient(cfg, "compute")
	err = c.Do(context.Background(), &Operation{ID: "op", Method: "GET", Path: "/v1/x"}, nil)
	if !IsUnauthorized(err) {
		t.Fatalf("err = %v, want an unauthorized error to surface", err)
	}
	if got := atomic.LoadInt32(&apiCalls); got != 2 {
		t.Errorf("api saw %d calls, want exactly 2 — one token refresh, then give up", got)
	}
}

func TestNewConfigRequiresCredentials(t *testing.T) {
	t.Setenv(EnvAccessKeyID, "")
	t.Setenv(EnvSecretAccessKey, "")
	t.Setenv(EnvAccessToken, "")
	_, err := NewConfig(context.Background())
	if err == nil {
		t.Fatal("NewConfig() succeeded with no credentials, want an error at construction rather than on every request")
	}
}

func TestNewConfigReadsTheEnvironment(t *testing.T) {
	t.Setenv(EnvAccessToken, "env-token")
	t.Setenv(EnvRegion, "sa-saopaulo-1")
	t.Setenv(EnvAccountID, "999999999999")

	cfg, err := NewConfig(context.Background())
	if err != nil {
		t.Fatalf("NewConfig: %v", err)
	}
	if cfg.Region != "sa-saopaulo-1" {
		t.Errorf("Region = %q, want sa-saopaulo-1", cfg.Region)
	}
	if cfg.AccountID != "999999999999" {
		t.Errorf("AccountID = %q, want 999999999999", cfg.AccountID)
	}
	tok, err := cfg.TokenSource.Token(context.Background())
	if err != nil || tok != "env-token" {
		t.Errorf("Token() = (%q, %v), want (env-token, nil)", tok, err)
	}
}

func TestOptionsBeatTheEnvironment(t *testing.T) {
	t.Setenv(EnvAccessToken, "env-token")
	t.Setenv(EnvRegion, "sa-saopaulo-1")

	cfg, err := NewConfig(context.Background(), WithRegion("us-ashburn-1"))
	if err != nil {
		t.Fatalf("NewConfig: %v", err)
	}
	if cfg.Region != "us-ashburn-1" {
		t.Errorf("Region = %q, want the option to win", cfg.Region)
	}
}

func contains(haystack, needle string) bool {
	return len(needle) == 0 || (len(haystack) >= len(needle) && indexOf(haystack, needle) >= 0)
}

func indexOf(haystack, needle string) int {
	for i := 0; i+len(needle) <= len(haystack); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

func asAuthError(err error, target **AuthError) bool {
	if e, ok := err.(*AuthError); ok {
		*target = e
		return true
	}
	return false
}
