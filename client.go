package basaltic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"runtime"
	"runtime/debug"
	"strings"
	"time"
)

// Version is this SDK's version, reported in the User-Agent.
//
// Read from the build information of the program that imports this package,
// which is where the answer actually lives: Go records the resolved version of
// every dependency. A constant here would have to be edited in the same commit
// as the tag and would be wrong the moment someone forgot — which is how it
// was found reporting 0.1.0 two releases later.
//
// Falls back to "dev" for a build with no module information: `go run`, a
// binary built with -buildvcs=false, or this module built as the main one.
var Version = readVersion()

// modulePath is this module, used to find its own entry in the build info.
const modulePath = "github.com/basaltic-sh/sdk-go"

func readVersion() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "dev"
	}
	for _, dep := range info.Deps {
		if dep.Path != modulePath {
			continue
		}
		// A replaced module reports the replacement's version, which is what
		// is actually linked in.
		if dep.Replace != nil && dep.Replace.Version != "" {
			return strings.TrimPrefix(dep.Replace.Version, "v")
		}
		if dep.Version != "" && dep.Version != "(devel)" {
			return strings.TrimPrefix(dep.Version, "v")
		}
		return "dev"
	}
	return "dev"
}

// Client is the transport one service is reached through. Generated service
// packages wrap it; construct one directly only to call an endpoint the SDK
// does not yet generate.
//
// A Client is safe for concurrent use.
type Client struct {
	cfg     *Config
	service string
	// stream performs streaming calls. It shares the configured client's
	// transport — and therefore its connection pool and TLS settings — but
	// carries no total-request deadline, because a multi-gigabyte body cannot
	// fit a fixed one.
	stream *http.Client
}

// NewClient builds the transport for one service. service is the short name
// the endpoint resolver keys on: "compute", "iam", "network".
func NewClient(cfg *Config, service string) *Client {
	c := &Client{cfg: cfg, service: service}
	if cfg != nil && cfg.HTTPClient != nil {
		c.stream = &http.Client{
			Transport:     cfg.HTTPClient.Transport,
			CheckRedirect: cfg.HTTPClient.CheckRedirect,
			Jar:           cfg.HTTPClient.Jar,
		}
	} else {
		c.stream = &http.Client{}
	}
	return c
}

// Config returns the configuration this client was built from.
func (c *Client) Config() *Config { return c.cfg }

// Service returns the short service name this client addresses.
func (c *Client) Service() string { return c.service }

// Operation describes one API call. Generated code fills it in; it is
// exported so a caller can reach an endpoint the SDK does not yet generate.
type Operation struct {
	// ID is the operationId from the specification, used in error messages.
	ID string
	// Method is the HTTP method.
	Method string
	// Path is the request path with {placeholder} segments, exactly as the
	// specification writes it.
	Path string
	// PathArgs supplies values for the placeholders in Path, in the order
	// they appear. They are escaped before substitution.
	PathArgs []string
	// Query holds query parameters.
	Query url.Values
	// Header holds operation-specific headers.
	Header http.Header

	// Body is marshalled as JSON when non-nil.
	Body any
	// RawBody is sent verbatim with ContentType, for the operations whose
	// body is not JSON. It takes precedence over Body.
	RawBody []byte
	// Stream is sent verbatim without being buffered, for bodies too large to
	// hold in memory. It takes precedence over RawBody and Body, and disables
	// retries — a stream cannot be replayed.
	Stream io.Reader
	// ContentType labels RawBody or Stream. JSON bodies set it themselves.
	ContentType string
	// Accept is the media type requested. Empty asks for JSON.
	Accept string

	// Unauthenticated sends no Authorization header, for the few operations
	// whose credentials are the request — the token endpoint itself.
	Unauthenticated bool
}

// Do performs the operation and decodes a JSON response into out.
//
// A nil out discards the body, which is what an operation with no response
// content wants. A failure status yields an [*Error].
func (c *Client) Do(ctx context.Context, op *Operation, out any, opts ...RequestOption) error {
	body, _, err := c.do(ctx, op, opts)
	if err != nil {
		return err
	}
	defer body.Close()

	if out == nil {
		// Drain so the connection can be reused, but bound it: a body we are
		// discarding should never be able to hold the caller up.
		_, _ = io.Copy(io.Discard, io.LimitReader(body, 1<<20))
		return nil
	}
	data, err := io.ReadAll(body)
	if err != nil {
		return fmt.Errorf("basaltic: %s: reading response: %w", op.ID, err)
	}
	if len(bytes.TrimSpace(data)) == 0 {
		// 204, or a 202 whose body the platform left empty. Not an error:
		// out simply stays at its zero value.
		return nil
	}
	if err := json.Unmarshal(data, out); err != nil {
		return fmt.Errorf("basaltic: %s: decoding response: %w", op.ID, err)
	}
	return nil
}

// DoStream performs the operation and hands back the response body unread,
// for the operations that return bytes rather than JSON — an object's
// contents, a console screenshot, an invoice PDF.
//
// The caller must close the returned body.
func (c *Client) DoStream(ctx context.Context, op *Operation, opts ...RequestOption) (io.ReadCloser, http.Header, error) {
	opts = append([]RequestOption{func(rs *requestSettings) error {
		rs.stream = true
		return nil
	}}, opts...)
	return c.do(ctx, op, opts)
}

// do runs the request, with retries, and returns the body of a successful
// response for the caller to consume.
func (c *Client) do(ctx context.Context, op *Operation, opts []RequestOption) (io.ReadCloser, http.Header, error) {
	if c.cfg == nil {
		return nil, nil, fmt.Errorf("basaltic: %s: client has no configuration", op.ID)
	}
	rs := &requestSettings{header: http.Header{}, query: url.Values{}}
	if err := rs.apply(c.cfg.requestOptions); err != nil {
		return nil, nil, err
	}
	if err := rs.apply(opts); err != nil {
		return nil, nil, err
	}

	if rs.timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, rs.timeout)
		defer cancel()
	}

	endpoint, err := c.cfg.EndpointResolver.ResolveEndpoint(c.service, c.cfg.Region)
	if err != nil {
		return nil, nil, err
	}
	target, err := buildURL(endpoint, op, rs)
	if err != nil {
		return nil, nil, fmt.Errorf("basaltic: %s: %w", op.ID, err)
	}

	payload, contentType, err := op.encodeBody()
	if err != nil {
		return nil, nil, fmt.Errorf("basaltic: %s: encoding request: %w", op.ID, err)
	}

	attempts := c.cfg.Retry.MaxAttempts
	if attempts < 1 {
		attempts = 1
	}
	// A stream cannot be replayed, and neither can a caller who asked for one
	// attempt. Both collapse to a single try.
	if rs.noRetry || op.Stream != nil {
		attempts = 1
	}

	// tokenRetried guards the one extra attempt granted to a token the
	// platform refused before its expiry — a revoked session. Without it,
	// every request would keep presenting the same dead cached token.
	tokenRetried := false

	var lastErr error
	for attempt := 1; ; attempt++ {
		req, err := c.newRequest(ctx, op, target, payload, contentType, rs)
		if err != nil {
			return nil, nil, err
		}

		httpClient := c.cfg.HTTPClient
		// The deadline exemption follows the BODY, not the caller's choice of
		// method. DoStream sets rs.stream for a streaming response, but an
		// operation that streams its REQUEST — uploading an object — reaches
		// here through Do, because it returns a decoded value. Sending a
		// multi-gigabyte body under a 30-second total-request timeout fails at
		// 30 seconds regardless of how well it is going.
		if rs.stream || op.Stream != nil {
			httpClient = c.stream
		}
		resp, err := httpClient.Do(req)

		if err == nil && resp.StatusCode < 300 {
			if rs.responseHeader != nil {
				*rs.responseHeader = resp.Header.Clone()
			}
			return resp.Body, resp.Header, nil
		}

		if err == nil && resp.StatusCode == http.StatusUnauthorized && !tokenRetried && !op.Unauthenticated {
			if inv, ok := c.cfg.TokenSource.(interface{ Invalidate() }); ok {
				tokenRetried = true
				drainAndClose(resp)
				inv.Invalidate()
				attempt--
				continue
			}
		}

		var apiErr error
		if err != nil {
			apiErr = fmt.Errorf("basaltic: %s: %w", op.ID, err)
		} else {
			data, _ := io.ReadAll(io.LimitReader(resp.Body, 4<<20))
			e := parseError(op.ID, resp, data)
			if rs.responseHeader != nil {
				*rs.responseHeader = resp.Header.Clone()
			}
			apiErr = e
		}
		lastErr = apiErr

		retry := attempt < attempts && c.cfg.Retry.shouldRetry(req, resp, err)
		if !retry {
			if resp != nil {
				resp.Body.Close()
			}
			return nil, nil, lastErr
		}
		delay := c.cfg.Retry.backoff(attempt, resp)
		if resp != nil {
			resp.Body.Close()
		}
		select {
		case <-ctx.Done():
			// Report why the caller stopped waiting, but keep the failure
			// that caused the retry — it is the more informative of the two.
			return nil, nil, fmt.Errorf("basaltic: %s: %w (last failure: %v)", op.ID, ctx.Err(), lastErr)
		case <-time.After(delay):
		}
	}
}

// newRequest builds one attempt. Called per attempt rather than once, because
// a retried request needs a fresh body reader and a token that may since have
// been refreshed.
func (c *Client) newRequest(ctx context.Context, op *Operation, target, payload string, contentType string, rs *requestSettings) (*http.Request, error) {
	var bodyReader io.Reader
	switch {
	case op.Stream != nil:
		bodyReader = op.Stream
	case payload != "":
		bodyReader = strings.NewReader(payload)
	}

	req, err := http.NewRequestWithContext(ctx, op.Method, target, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("basaltic: %s: %w", op.ID, err)
	}

	for k, vs := range c.cfg.baseHeaders {
		for _, v := range vs {
			req.Header.Add(k, v)
		}
	}
	for k, vs := range op.Header {
		for _, v := range vs {
			req.Header.Add(k, v)
		}
	}
	for k, vs := range rs.header {
		req.Header[http.CanonicalHeaderKey(k)] = append([]string(nil), vs...)
	}

	if contentType != "" && req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", contentType)
	}
	accept := op.Accept
	if accept == "" {
		accept = "application/json"
	}
	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", accept)
	}
	if req.Header.Get("User-Agent") == "" {
		req.Header.Set("User-Agent", c.userAgent())
	}

	accountID := c.cfg.AccountID
	if rs.accountIDSet {
		accountID = rs.accountID
	}
	if accountID != "" && req.Header.Get(accountHeader) == "" {
		req.Header.Set(accountHeader, accountID)
	}

	if !op.Unauthenticated {
		if c.cfg.TokenSource == nil {
			return nil, fmt.Errorf("basaltic: %s: no token source configured", op.ID)
		}
		token, err := c.cfg.TokenSource.Token(ctx)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Authorization", "Bearer "+token)
	}
	return req, nil
}

func (c *Client) userAgent() string {
	ua := fmt.Sprintf("basaltic-go/%s Go/%s (%s; %s)",
		Version, strings.TrimPrefix(runtime.Version(), "go"), runtime.GOOS, runtime.GOARCH)
	if c.cfg.UserAgentExtra != "" {
		ua += " " + c.cfg.UserAgentExtra
	}
	return ua
}

// encodeBody renders the request body once, so every retry sends identical
// bytes rather than re-marshalling a value the caller may have mutated.
func (op *Operation) encodeBody() (payload, contentType string, err error) {
	switch {
	case op.Stream != nil:
		return "", op.ContentType, nil
	case op.RawBody != nil:
		ct := op.ContentType
		if ct == "" {
			ct = "application/octet-stream"
		}
		return string(op.RawBody), ct, nil
	case op.Body != nil:
		data, err := json.Marshal(op.Body)
		if err != nil {
			return "", "", err
		}
		return string(data), "application/json", nil
	}
	return "", "", nil
}

// buildURL substitutes the path placeholders and attaches the query.
func buildURL(endpoint string, op *Operation, rs *requestSettings) (string, error) {
	path, err := expandPath(op.Path, op.PathArgs)
	if err != nil {
		return "", err
	}
	q := url.Values{}
	for k, vs := range op.Query {
		q[k] = append([]string(nil), vs...)
	}
	for k, vs := range rs.query {
		q[k] = append(q[k], vs...)
	}
	u := endpoint + path
	if len(q) > 0 {
		u += "?" + q.Encode()
	}
	return u, nil
}

// expandPath replaces each {placeholder} with the next path argument.
//
// Arguments are escaped as path segments, so an identifier containing a slash
// cannot reach into another route. An empty argument is refused rather than
// silently collapsing the path onto a different endpoint — "" for an instance
// id would turn a get into a list.
func expandPath(pattern string, args []string) (string, error) {
	var b strings.Builder
	rest := pattern
	i := 0
	for {
		open := strings.IndexByte(rest, '{')
		if open < 0 {
			b.WriteString(rest)
			break
		}
		end := strings.IndexByte(rest[open:], '}')
		if end < 0 {
			return "", fmt.Errorf("malformed path template %q", pattern)
		}
		end += open
		name := rest[open+1 : end]
		if i >= len(args) {
			return "", fmt.Errorf("path template %q needs a value for {%s}", pattern, name)
		}
		if args[i] == "" {
			return "", fmt.Errorf("%s must not be empty", name)
		}
		b.WriteString(rest[:open])
		b.WriteString(url.PathEscape(args[i]))
		i++
		rest = rest[end+1:]
	}
	if i != len(args) {
		return "", fmt.Errorf("path template %q takes %d values, got %d", pattern, i, len(args))
	}
	return b.String(), nil
}

func drainAndClose(resp *http.Response) {
	if resp == nil || resp.Body == nil {
		return
	}
	_, _ = io.Copy(io.Discard, io.LimitReader(resp.Body, 1<<20))
	resp.Body.Close()
}
