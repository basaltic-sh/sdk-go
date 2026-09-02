package basaltic

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// newTestConfig wires a Config at a test server with a fixed token, so a test
// that is not about authentication does not have to stand up a token endpoint.
func newTestConfig(t *testing.T, srv *httptest.Server, opts ...Option) *Config {
	t.Helper()
	base := []Option{
		WithAccessToken("test-token"),
		WithServiceEndpoint("compute", srv.URL),
		WithRegion("sa-saopaulo-1"),
	}
	cfg, err := NewConfig(context.Background(), append(base, opts...)...)
	if err != nil {
		t.Fatalf("NewConfig: %v", err)
	}
	return cfg
}

func TestExpandPath(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		args    []string
		want    string
		wantErr string
	}{
		{name: "no placeholders", pattern: "/v1/instances", want: "/v1/instances"},
		{name: "one", pattern: "/v1/instances/{id}", args: []string{"abc"}, want: "/v1/instances/abc"},
		{
			name:    "two in order",
			pattern: "/v1/buckets/{bucket}/objects/{key}",
			args:    []string{"photos", "cat.png"},
			want:    "/v1/buckets/photos/objects/cat.png",
		},
		{
			// A slash in an identifier must not be able to reach another
			// route: /v1/instances/../flavors would be a different endpoint.
			name:    "escapes path separators",
			pattern: "/v1/instances/{id}",
			args:    []string{"a/../b"},
			want:    "/v1/instances/a%2F..%2Fb",
		},
		{
			// "" would silently turn a get into a list.
			name:    "rejects an empty argument",
			pattern: "/v1/instances/{instance_id}",
			args:    []string{""},
			wantErr: "instance_id must not be empty",
		},
		{
			name:    "rejects too few arguments",
			pattern: "/v1/instances/{id}",
			wantErr: "needs a value",
		},
		{
			name:    "rejects too many arguments",
			pattern: "/v1/instances/{id}",
			args:    []string{"a", "b"},
			wantErr: "takes 1 values, got 2",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := expandPath(tc.pattern, tc.args)
			if tc.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
					t.Fatalf("expandPath() error = %v, want it to contain %q", err, tc.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("expandPath() unexpected error: %v", err)
			}
			if got != tc.want {
				t.Errorf("expandPath() = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestDoSendsBearerTokenAndAccount(t *testing.T) {
	var gotAuth, gotAccount, gotUA, gotAccept string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
		gotAccount = r.Header.Get("X-Account-Id")
		gotUA = r.Header.Get("User-Agent")
		gotAccept = r.Header.Get("Accept")
		fmt.Fprint(w, `{"instance":{"id":"i-1"}}`)
	}))
	defer srv.Close()

	cfg := newTestConfig(t, srv, WithAccountID("123456789012"))
	c := NewClient(cfg, "compute")

	var out struct {
		Instance struct {
			ID string `json:"id"`
		} `json:"instance"`
	}
	op := &Operation{ID: "getInstance", Method: "GET", Path: "/v1/instances/{id}", PathArgs: []string{"i-1"}}
	if err := c.Do(context.Background(), op, &out); err != nil {
		t.Fatalf("Do: %v", err)
	}

	if want := "Bearer test-token"; gotAuth != want {
		t.Errorf("Authorization = %q, want %q", gotAuth, want)
	}
	if want := "123456789012"; gotAccount != want {
		t.Errorf("X-Account-Id = %q, want %q", gotAccount, want)
	}
	if !strings.HasPrefix(gotUA, "basaltic-go/") {
		t.Errorf("User-Agent = %q, want it to start with basaltic-go/", gotUA)
	}
	if want := "application/json"; gotAccept != want {
		t.Errorf("Accept = %q, want %q", gotAccept, want)
	}
	if out.Instance.ID != "i-1" {
		t.Errorf("decoded id = %q, want %q", out.Instance.ID, "i-1")
	}
}

func TestWithRequestAccountIDOverridesAndClears(t *testing.T) {
	var got string
	var seen bool
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got = r.Header.Get("X-Account-Id")
		_, seen = r.Header["X-Account-Id"]
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c := NewClient(newTestConfig(t, srv, WithAccountID("111111111111")), "compute")
	op := &Operation{ID: "op", Method: "GET", Path: "/v1/x"}

	if err := c.Do(context.Background(), op, nil, WithRequestAccountID("222222222222")); err != nil {
		t.Fatalf("Do: %v", err)
	}
	if want := "222222222222"; got != want {
		t.Errorf("X-Account-Id = %q, want the per-request override %q", got, want)
	}

	// An explicit empty account means "use the credential's own account",
	// which has to send no header at all rather than an empty one.
	if err := c.Do(context.Background(), op, nil, WithRequestAccountID("")); err != nil {
		t.Fatalf("Do: %v", err)
	}
	if seen {
		t.Errorf("X-Account-Id was sent for an explicitly empty account")
	}
}

func TestErrorEnvelopeIsParsed(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `{"error":{"code":"NOT_FOUND","message":"no such instance","request_id":"req-42"}}`)
	}))
	defer srv.Close()

	c := NewClient(newTestConfig(t, srv), "compute")
	err := c.Do(context.Background(), &Operation{ID: "getInstance", Method: "GET", Path: "/v1/x"}, nil)
	if err == nil {
		t.Fatal("Do() succeeded, want a not-found error")
	}
	apiErr, ok := AsError(err)
	if !ok {
		t.Fatalf("AsError() = false for %T", err)
	}
	if apiErr.Code != "NOT_FOUND" || apiErr.RequestID != "req-42" || apiErr.StatusCode != 404 {
		t.Errorf("got %+v, want code NOT_FOUND, request req-42, status 404", apiErr)
	}
	if !IsNotFound(err) {
		t.Error("IsNotFound() = false")
	}
	if IsAccessDenied(err) || IsConflict(err) || IsRateLimited(err) {
		t.Error("a not-found error matched another class")
	}
	if !strings.Contains(err.Error(), "req-42") {
		t.Errorf("Error() = %q, want it to quote the request id", err.Error())
	}
}

func TestErrorWithoutEnvelopeFallsBackToStatus(t *testing.T) {
	// A proxy in front of the service answers without the platform envelope.
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "<html>forbidden</html>")
	}))
	defer srv.Close()

	c := NewClient(newTestConfig(t, srv), "compute")
	err := c.Do(context.Background(), &Operation{ID: "op", Method: "GET", Path: "/v1/x"}, nil)
	if !IsAccessDenied(err) {
		t.Errorf("IsAccessDenied() = false for a bare 403: %v", err)
	}
	apiErr, _ := AsError(err)
	if string(apiErr.Body) != "<html>forbidden</html>" {
		t.Errorf("Body = %q, want the raw response kept", apiErr.Body)
	}
}

func TestRetriesServerErrorsThenSucceeds(t *testing.T) {
	var calls int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if atomic.AddInt32(&calls, 1) < 3 {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		fmt.Fprint(w, `{"ok":true}`)
	}))
	defer srv.Close()

	cfg := newTestConfig(t, srv, WithRetry(RetryConfig{MaxAttempts: 4, BaseDelay: time.Millisecond, MaxDelay: 5 * time.Millisecond}))
	c := NewClient(cfg, "compute")
	if err := c.Do(context.Background(), &Operation{ID: "op", Method: "GET", Path: "/v1/x"}, nil); err != nil {
		t.Fatalf("Do: %v", err)
	}
	if got := atomic.LoadInt32(&calls); got != 3 {
		t.Errorf("server saw %d calls, want 3", got)
	}
}

func TestPostIsNotRetriedWithoutAnIdempotencyKey(t *testing.T) {
	// A create that timed out may already have succeeded. Repeating it blindly
	// is how duplicate resources happen, so it must be sent exactly once.
	var calls int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&calls, 1)
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer srv.Close()

	cfg := newTestConfig(t, srv, WithRetry(RetryConfig{MaxAttempts: 4, BaseDelay: time.Millisecond, MaxDelay: time.Millisecond}))
	c := NewClient(cfg, "compute")
	op := &Operation{ID: "createInstance", Method: "POST", Path: "/v1/x", Body: map[string]string{"name": "web"}}

	if err := c.Do(context.Background(), op, nil); err == nil {
		t.Fatal("Do() succeeded, want the 503 to surface")
	}
	if got := atomic.LoadInt32(&calls); got != 1 {
		t.Errorf("server saw %d calls, want exactly 1 — a POST without an idempotency key must not be repeated", got)
	}
}

func TestPostIsRetriedWithAnIdempotencyKey(t *testing.T) {
	var calls int32
	var keys []string
	var mu sync.Mutex
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		keys = append(keys, r.Header.Get("Idempotency-Key"))
		mu.Unlock()
		if atomic.AddInt32(&calls, 1) < 2 {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		fmt.Fprint(w, `{}`)
	}))
	defer srv.Close()

	cfg := newTestConfig(t, srv, WithRetry(RetryConfig{MaxAttempts: 3, BaseDelay: time.Millisecond, MaxDelay: time.Millisecond}))
	c := NewClient(cfg, "compute")
	op := &Operation{ID: "createInstance", Method: "POST", Path: "/v1/x", Body: map[string]string{"name": "web"}}

	if err := c.Do(context.Background(), op, nil, WithIdempotencyKey("key-1")); err != nil {
		t.Fatalf("Do: %v", err)
	}
	if got := atomic.LoadInt32(&calls); got != 2 {
		t.Errorf("server saw %d calls, want 2", got)
	}
	// The retry must present the same key, or the platform cannot recognise
	// it as a replay.
	mu.Lock()
	defer mu.Unlock()
	for i, k := range keys {
		if k != "key-1" {
			t.Errorf("attempt %d sent Idempotency-Key %q, want key-1", i+1, k)
		}
	}
}

func TestRetryHonoursRetryAfter(t *testing.T) {
	var calls int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if atomic.AddInt32(&calls, 1) == 1 {
			w.Header().Set("Retry-After", "1")
			w.Header().Set("X-RateLimit-Limit", "100")
			w.Header().Set("X-RateLimit-Remaining", "0")
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		fmt.Fprint(w, `{}`)
	}))
	defer srv.Close()

	// MaxDelay caps what Retry-After can ask for, so the test does not sleep
	// a real second while still exercising the header path.
	cfg := newTestConfig(t, srv, WithRetry(RetryConfig{MaxAttempts: 2, BaseDelay: time.Hour, MaxDelay: 10 * time.Millisecond}))
	c := NewClient(cfg, "compute")

	start := time.Now()
	if err := c.Do(context.Background(), &Operation{ID: "op", Method: "GET", Path: "/v1/x"}, nil); err != nil {
		t.Fatalf("Do: %v", err)
	}
	if elapsed := time.Since(start); elapsed > time.Second {
		t.Errorf("waited %v, want the delay capped by MaxDelay", elapsed)
	}
	if got := atomic.LoadInt32(&calls); got != 2 {
		t.Errorf("server saw %d calls, want 2", got)
	}
}

func TestRateLimitHeadersAreReadable(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Retry-After", "30")
		w.Header().Set("X-RateLimit-Limit", "100")
		w.Header().Set("X-RateLimit-Remaining", "0")
		w.WriteHeader(http.StatusTooManyRequests)
		fmt.Fprint(w, `{"error":{"code":"TOO_MANY_REQUESTS","message":"slow down","request_id":"r1"}}`)
	}))
	defer srv.Close()

	c := NewClient(newTestConfig(t, srv, WithoutRetry()), "compute")
	err := c.Do(context.Background(), &Operation{ID: "op", Method: "GET", Path: "/v1/x"}, nil)
	if !IsRateLimited(err) {
		t.Fatalf("IsRateLimited() = false: %v", err)
	}
	apiErr, _ := AsError(err)
	if d, ok := apiErr.RetryAfter(); !ok || d != 30*time.Second {
		t.Errorf("RetryAfter() = (%v, %v), want (30s, true)", d, ok)
	}
	limit, remaining, ok := apiErr.RateLimit()
	if !ok || limit != 100 || remaining != 0 {
		t.Errorf("RateLimit() = (%d, %d, %v), want (100, 0, true)", limit, remaining, ok)
	}
}

func TestNoContentLeavesOutputUntouched(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c := NewClient(newTestConfig(t, srv), "compute")
	out := struct {
		Name string `json:"name"`
	}{Name: "unchanged"}
	if err := c.Do(context.Background(), &Operation{ID: "op", Method: "DELETE", Path: "/v1/x"}, &out); err != nil {
		t.Fatalf("Do: %v", err)
	}
	if out.Name != "unchanged" {
		t.Errorf("out.Name = %q, want an empty body to decode into nothing", out.Name)
	}
}

func TestDoStreamReturnsTheBodyUnread(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/octet-stream")
		fmt.Fprint(w, "the object bytes")
	}))
	defer srv.Close()

	c := NewClient(newTestConfig(t, srv), "compute")
	body, header, err := c.DoStream(context.Background(), &Operation{ID: "getObject", Method: "GET", Path: "/v1/x"})
	if err != nil {
		t.Fatalf("DoStream: %v", err)
	}
	defer body.Close()
	data, err := io.ReadAll(body)
	if err != nil {
		t.Fatalf("reading stream: %v", err)
	}
	if string(data) != "the object bytes" {
		t.Errorf("stream = %q, want the object bytes", data)
	}
	if got := header.Get("Content-Type"); got != "application/octet-stream" {
		t.Errorf("Content-Type = %q", got)
	}
}

func TestQueryAndOperationHeadersAreSent(t *testing.T) {
	var gotQuery, gotCustom string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotQuery = r.URL.RawQuery
		gotCustom = r.Header.Get("X-Trace")
		fmt.Fprint(w, `{}`)
	}))
	defer srv.Close()

	c := NewClient(newTestConfig(t, srv), "compute")
	op := &Operation{ID: "listInstances", Method: "GET", Path: "/v1/instances"}
	op.Query = map[string][]string{"limit": {"50"}}
	err := c.Do(context.Background(), op, nil,
		WithQueryParam("name", "web-*"),
		WithRequestHeader("X-Trace", "abc"))
	if err != nil {
		t.Fatalf("Do: %v", err)
	}
	if !strings.Contains(gotQuery, "limit=50") || !strings.Contains(gotQuery, "name=web-%2A") {
		t.Errorf("query = %q, want both the operation and the request-option parameters", gotQuery)
	}
	if gotCustom != "abc" {
		t.Errorf("X-Trace = %q, want abc", gotCustom)
	}
}

func TestUnauthenticatedOperationSendsNoToken(t *testing.T) {
	var sawAuth bool
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, sawAuth = r.Header["Authorization"]
		fmt.Fprint(w, `{}`)
	}))
	defer srv.Close()

	c := NewClient(newTestConfig(t, srv), "compute")
	op := &Operation{ID: "getOAuthToken", Method: "POST", Path: "/v1/oauth/token", Unauthenticated: true}
	if err := c.Do(context.Background(), op, nil); err != nil {
		t.Fatalf("Do: %v", err)
	}
	if sawAuth {
		t.Error("an unauthenticated operation sent an Authorization header")
	}
}

func TestJSONBodyIsSentOnce(t *testing.T) {
	var gotBody, gotType string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		gotBody, gotType = string(b), r.Header.Get("Content-Type")
		fmt.Fprint(w, `{}`)
	}))
	defer srv.Close()

	c := NewClient(newTestConfig(t, srv), "compute")
	op := &Operation{
		ID: "createInstance", Method: "POST", Path: "/v1/instances",
		Body: map[string]any{"name": "web-01"},
	}
	if err := c.Do(context.Background(), op, nil); err != nil {
		t.Fatalf("Do: %v", err)
	}
	var decoded map[string]any
	if err := json.Unmarshal([]byte(gotBody), &decoded); err != nil {
		t.Fatalf("body was not JSON: %v", err)
	}
	if decoded["name"] != "web-01" {
		t.Errorf("body = %s, want name web-01", gotBody)
	}
	if gotType != "application/json" {
		t.Errorf("Content-Type = %q, want application/json", gotType)
	}
}

// TestErrorClassificationUsesRealPlatformCodes checks the classifiers against
// codes the platform actually emits. The platform names errors per resource —
// INSTANCE_NOT_FOUND rather than NOT_FOUND — so a classifier that matched a
// fixed list of codes would miss almost all of them.
func TestErrorClassificationUsesRealPlatformCodes(t *testing.T) {
	tests := []struct {
		code   string
		status int
		want   string
	}{
		{"INSTANCE_NOT_FOUND", 404, "not-found"},
		{"VOLUME_NOT_FOUND", 404, "not-found"},
		{"DATABASE_USER_NOT_FOUND", 404, "not-found"},
		{"NOT_FOUND", 404, "not-found"},
		{"IMAGE_GONE", 410, "not-found"},

		{"ACCESS_DENIED", 403, "access-denied"},
		{"AUTHORIZATION_FAILED", 403, "access-denied"},

		// Shares its status with access-denied, and needs the opposite
		// response: raising the quota, not changing a policy.
		{"QUOTA_EXCEEDED", 403, "quota"},
		{"INSTANCE_QUOTA_EXCEEDED", 403, "quota"},

		{"CERTIFICATE_NAME_EXISTS", 409, "conflict"},
		{"SUBNET_IN_USE", 409, "conflict"},
		{"EMAIL_DEDICATED_IP_TAKEN", 409, "conflict"},

		{"INVALID_INPUT", 400, "invalid"},
		{"CERTIFICATE_INVALID_STATUS", 400, "invalid"},
		{"VALIDATION_ERROR", 422, "invalid"},

		{"AUTH_INVALID_TOKEN", 401, "unauthorized"},
		{"AUTH_SESSION_NOT_FOUND", 401, "unauthorized"},

		{"TOO_MANY_REQUESTS", 429, "rate-limited"},
	}

	classify := func(err error) string {
		switch {
		case IsQuotaExceeded(err):
			return "quota"
		case IsNotFound(err):
			return "not-found"
		case IsAccessDenied(err):
			return "access-denied"
		case IsUnauthorized(err):
			return "unauthorized"
		case IsConflict(err):
			return "conflict"
		case IsRateLimited(err):
			return "rate-limited"
		case IsInvalidInput(err):
			return "invalid"
		}
		return "unclassified"
	}

	for _, tc := range tests {
		t.Run(tc.code, func(t *testing.T) {
			err := &Error{Code: tc.code, StatusCode: tc.status, Message: "m", RequestID: "r"}
			if got := classify(err); got != tc.want {
				t.Errorf("%s (http %d) classified as %q, want %q", tc.code, tc.status, got, tc.want)
			}
		})
	}
}

// A quota refusal must not read as an authorization failure: the two share a
// status and their remedies are opposite.
func TestQuotaIsNotAccessDenied(t *testing.T) {
	quota := &Error{Code: "QUOTA_EXCEEDED", StatusCode: 403}
	if IsAccessDenied(quota) {
		t.Error("IsAccessDenied() = true for a quota refusal")
	}
	if !IsQuotaExceeded(quota) {
		t.Error("IsQuotaExceeded() = false for QUOTA_EXCEEDED")
	}

	denied := &Error{Code: "ACCESS_DENIED", StatusCode: 403}
	if IsQuotaExceeded(denied) {
		t.Error("IsQuotaExceeded() = true for an authorization failure")
	}
	if !IsAccessDenied(denied) {
		t.Error("IsAccessDenied() = false for ACCESS_DENIED")
	}
}

func TestIsTransient(t *testing.T) {
	for _, status := range []int{429, 500, 502, 503, 504} {
		if !IsTransient(&Error{StatusCode: status}) {
			t.Errorf("IsTransient() = false for http %d", status)
		}
	}
	for _, status := range []int{400, 403, 404, 409, 422} {
		if IsTransient(&Error{StatusCode: status}) {
			t.Errorf("IsTransient() = true for http %d", status)
		}
	}
}
