package basaltic_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"sync"
	"testing"

	basaltic "github.com/basaltic-sh/sdk-go"
	"github.com/basaltic-sh/sdk-go/audit"
	"github.com/basaltic-sh/sdk-go/billing"
	"github.com/basaltic-sh/sdk-go/certificate"
	"github.com/basaltic-sh/sdk-go/compute"
	"github.com/basaltic-sh/sdk-go/dns"
	"github.com/basaltic-sh/sdk-go/iam"
	"github.com/basaltic-sh/sdk-go/kms"
	"github.com/basaltic-sh/sdk-go/loadbalancer"
	"github.com/basaltic-sh/sdk-go/network"
	"github.com/basaltic-sh/sdk-go/quota"
	"github.com/basaltic-sh/sdk-go/secrets"
	"github.com/basaltic-sh/sdk-go/storage"
	"github.com/basaltic-sh/sdk-go/telemetry"
)

// These tests walk every generated method by reflection and assert properties
// that are NOT re-derived from the specification.
//
// That distinction is the whole point. A generated test asserting that
// GetInstance sends GET /v1/instances/{instance_id} would read its expectation
// from the same specification node the code was emitted from, so an emitter
// bug would produce a matching bug in the test. What follows instead are
// invariants the specification cannot express and the generator does not know:
// that no placeholder survives into a URL, that a credential is attached, that
// every paginator terminates. Those catch real emitter faults.

// unauthenticatedMethods are the operations that deliberately send no bearer
// token, because the credentials in the request are the authentication, or
// because the endpoint is a public catalogue.
//
// Kept as an explicit list so that an operation silently losing its
// credentials fails here. It is the trap the specification's own conventions
// warn about: deleting a security scheme where it is the only entry turns the
// operation unauthenticated without any other sign.
var unauthenticatedMethods = map[string]bool{
	"iam.AssumeRoleWithWebIdentity": true,
	"iam.GetOAuthToken":             true,
	"iam.ListRegions":               true,
	"billing.ListPrices":            true,
}

// recorder captures what the SDK actually put on the wire.
type recorder struct {
	mu       sync.Mutex
	requests []*http.Request
}

func (r *recorder) reset() {
	r.mu.Lock()
	r.requests = nil
	r.mu.Unlock()
}

func (r *recorder) all() []*http.Request {
	r.mu.Lock()
	defer r.mu.Unlock()
	return append([]*http.Request(nil), r.requests...)
}

// newSurfaceHarness points every service package at one recording server.
func newSurfaceHarness(t *testing.T) (*basaltic.Config, *recorder) {
	t.Helper()
	rec := &recorder{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.Copy(io.Discard, r.Body)
		rec.mu.Lock()
		rec.requests = append(rec.requests, r.Clone(context.Background()))
		rec.mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		// An empty object decodes into every result type the SDK returns, and
		// leaves every page empty, so each paginator terminates after one
		// fetch.
		fmt.Fprint(w, `{}`)
	}))
	t.Cleanup(srv.Close)

	cfg, err := basaltic.NewConfig(context.Background(),
		basaltic.WithAccessToken("surface-token"),
		basaltic.WithRegion("sa-saopaulo-1"),
		basaltic.WithEndpointResolver(basaltic.EndpointResolverFunc(
			func(service, region string) (string, error) { return srv.URL, nil })),
		basaltic.WithoutRetry(),
	)
	if err != nil {
		t.Fatalf("NewConfig: %v", err)
	}
	return cfg, rec
}

// allClients constructs one client per generated service package.
func allClients(cfg *basaltic.Config) map[string]any {
	return map[string]any{
		"audit":        audit.New(cfg),
		"billing":      billing.New(cfg),
		"certificate":  certificate.New(cfg),
		"compute":      compute.New(cfg),
		"dns":          dns.New(cfg),
		"iam":          iam.New(cfg),
		"kms":          kms.New(cfg),
		"loadbalancer": loadbalancer.New(cfg),
		"network":      network.New(cfg),
		"quota":        quota.New(cfg),
		"secrets":      secrets.New(cfg),
		"storage":      storage.New(cfg),
		"telemetry":    telemetry.New(cfg),
	}
}

var (
	ctxType    = reflect.TypeOf((*context.Context)(nil)).Elem()
	errorType  = reflect.TypeOf((*error)(nil)).Elem()
	readerType = reflect.TypeOf((*io.Reader)(nil)).Elem()
)

// TestEveryGeneratedMethodIssuesAWellFormedRequest calls all 393 operations
// and checks what reached the wire.
func TestEveryGeneratedMethodIssuesAWellFormedRequest(t *testing.T) {
	cfg, rec := newSurfaceHarness(t)

	called, iterators := 0, 0
	for svc, client := range allClients(cfg) {
		v := reflect.ValueOf(client)
		for i := 0; i < v.NumMethod(); i++ {
			method := v.Type().Method(i)
			fn := v.Method(i)
			if !takesContext(fn.Type()) {
				// Transport(), and anything else that is not an operation.
				continue
			}
			name := svc + "." + method.Name

			t.Run(name, func(t *testing.T) {
				rec.reset()
				args, strArgs, ok := buildArgs(fn.Type())
				if !ok {
					t.Skipf("cannot synthesise arguments for %s", fn.Type())
				}
				out := fn.Call(args)

				if seq, isIter := findIterator(out); isIter {
					iterators++
					drainIterator(t, seq)
				} else if err := returnedError(out); err != nil {
					t.Fatalf("%s returned an error against a 200 {} response: %v", name, err)
				}
				called++

				reqs := rec.all()
				if len(reqs) != 1 {
					t.Fatalf("%s issued %d requests, want exactly 1", name, len(reqs))
				}
				checkRequest(t, name, reqs[0])
				checkPathArgOrder(t, name, reqs[0].URL.Path, strArgs)
			})
		}
	}

	// A guard on the walk itself: a refactor that stopped discovering methods
	// would otherwise turn this whole file into a no-op that still passes.
	if called < 380 {
		t.Errorf("exercised only %d operations, want the full generated surface (393)", called)
	}
	if iterators < 40 {
		t.Errorf("exercised only %d paginators, want them all (51)", iterators)
	}
}

func checkRequest(t *testing.T, name string, req *http.Request) {
	t.Helper()
	path := req.URL.Path

	// An unsubstituted placeholder means a path argument never made it into
	// the URL. The request would 404 against the real API, having quietly
	// asked for a resource literally named "{instance_id}".
	if strings.ContainsAny(path, "{}") {
		t.Errorf("%s: path %q still carries a placeholder", name, path)
	}
	// An empty segment means an argument vanished, which silently addresses a
	// different endpoint: /v1/instances//volumes is not /v1/instances/x/volumes.
	if strings.Contains(path, "//") {
		t.Errorf("%s: path %q has an empty segment", name, path)
	}
	if !strings.HasPrefix(path, "/v1/") {
		t.Errorf("%s: path %q does not start with /v1/", name, path)
	}
	switch req.Method {
	case http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch,
		http.MethodDelete, http.MethodHead:
	default:
		t.Errorf("%s: unexpected method %q", name, req.Method)
	}

	auth := req.Header.Get("Authorization")
	if unauthenticatedMethods[name] {
		if auth != "" {
			t.Errorf("%s is meant to be unauthenticated but sent an Authorization header", name)
		}
		return
	}
	if auth != "Bearer surface-token" {
		t.Errorf("%s sent Authorization %q, want the configured bearer token — "+
			"if this operation is deliberately unauthenticated, add it to unauthenticatedMethods", name, auth)
	}
}

// takesContext reports whether a bound method looks like an operation.
func takesContext(t reflect.Type) bool {
	return t.NumIn() > 0 && t.In(0) == ctxType
}

// buildArgs synthesises a call.
//
// Each string argument gets a value that names its own position — "sdkarg0",
// "sdkarg1" — rather than one plausible identifier for all of them. That is
// what makes a swapped pair visible: identical arguments would produce an
// identical URL whichever order they went in, and two transposed resource ids
// address a different resource while looking perfectly well-formed.
//
// Bodies and parameter structs are left nil, which every operation accepts.
func buildArgs(t reflect.Type) ([]reflect.Value, []string, bool) {
	n := t.NumIn()
	if t.IsVariadic() {
		n-- // the ...RequestOption tail; Call fills it with nothing
	}
	args := make([]reflect.Value, 0, n)
	var strArgs []string
	for i := 0; i < n; i++ {
		in := t.In(i)
		switch {
		case i == 0 && in == ctxType:
			args = append(args, reflect.ValueOf(context.Background()))
		case in.Kind() == reflect.String:
			v := fmt.Sprintf("sdkarg%d", len(strArgs))
			strArgs = append(strArgs, v)
			args = append(args, reflect.ValueOf(v))
		case in == readerType:
			args = append(args, reflect.ValueOf(strings.NewReader("")))
		case in.Kind() == reflect.Pointer, in.Kind() == reflect.Map, in.Kind() == reflect.Slice:
			args = append(args, reflect.Zero(in))
		default:
			return nil, nil, false
		}
	}
	return args, strArgs, true
}

// checkPathArgOrder asserts that the positional arguments reached the URL in
// the order they were passed.
//
// This is the SDK's own contract — path parameters are positional, in the
// order the placeholders appear in the path — so it is checked against the
// call, not against the specification. A generator that emitted PathArgs in
// the specification's parameter order rather than the path's would transpose
// two ids here while still producing a URL that looks entirely valid.
//
// Arguments that do not appear in the path are skipped: an operation whose
// body is a string (a zone file, say) takes one that never reaches the URL.
func checkPathArgOrder(t *testing.T, name, path string, strArgs []string) {
	t.Helper()
	positions := make([]int, 0, len(strArgs))
	for _, a := range strArgs {
		if at := strings.Index(path, a); at >= 0 {
			positions = append(positions, at)
		}
	}
	for i := 1; i < len(positions); i++ {
		if positions[i] < positions[i-1] {
			t.Errorf("%s: path %q has its positional arguments out of order — "+
				"argument %d appears before argument %d, so two identifiers are transposed",
				name, path, i, i-1)
			return
		}
	}
}

func returnedError(out []reflect.Value) error {
	for _, v := range out {
		if v.Type() == errorType && !v.IsNil() {
			return v.Interface().(error)
		}
	}
	return nil
}

// findIterator spots an iter.Seq2 return: a func taking one func and
// returning nothing.
func findIterator(out []reflect.Value) (reflect.Value, bool) {
	for _, v := range out {
		t := v.Type()
		if t.Kind() == reflect.Func && t.NumIn() == 1 && t.NumOut() == 0 && t.In(0).Kind() == reflect.Func {
			return v, true
		}
	}
	return reflect.Value{}, false
}

// drainIterator ranges a generated "All" method to exhaustion, proving it
// terminates. A paginator whose cursor never advances would spin here rather
// than in production.
func drainIterator(t *testing.T, seq reflect.Value) {
	t.Helper()
	const limit = 50
	yieldType := seq.Type().In(0)
	count := 0
	var iterErr error

	yield := reflect.MakeFunc(yieldType, func(args []reflect.Value) []reflect.Value {
		count++
		if len(args) > 1 && args[1].Type() == errorType && !args[1].IsNil() {
			iterErr = args[1].Interface().(error)
		}
		return []reflect.Value{reflect.ValueOf(count < limit)}
	})
	seq.Call([]reflect.Value{yield})

	if iterErr != nil {
		t.Fatalf("paginator yielded an error against a 200 {} response: %v", iterErr)
	}
	if count >= limit {
		t.Fatalf("paginator did not terminate on a single empty page")
	}
}

// TestEveryServiceResolvesToItsOwnHost checks the endpoint each generated
// package registered, which the recording harness above deliberately bypasses.
func TestEveryServiceResolvesToItsOwnHost(t *testing.T) {
	cfg, err := basaltic.NewConfig(context.Background(),
		basaltic.WithAccessToken("t"),
		basaltic.WithRegion("sa-saopaulo-1"),
	)
	if err != nil {
		t.Fatalf("NewConfig: %v", err)
	}

	regional := map[string]bool{
		"certificate": true, "compute": true, "kms": true,
		"loadbalancer": true, "network": true, "secrets": true, "storage": true,
		"telemetry": true,
	}

	for svc := range allClients(cfg) {
		t.Run(svc, func(t *testing.T) {
			got, err := cfg.EndpointResolver.ResolveEndpoint(svc, cfg.Region)
			if err != nil {
				t.Fatalf("ResolveEndpoint(%q): %v", svc, err)
			}
			want := "https://" + svc + ".basaltic.sh"
			if regional[svc] {
				want = "https://" + svc + ".sa-saopaulo-1.basaltic.sh"
			}
			if got != want {
				t.Errorf("ResolveEndpoint(%q) = %q, want %q", svc, got, want)
			}
		})
	}
}

// TestRegionalServicesRefuseWithoutARegion proves the region requirement is
// carried by the specification's server template rather than by a hand-kept
// list that could drift from it.
func TestRegionalServicesRefuseWithoutARegion(t *testing.T) {
	cfg, err := basaltic.NewConfig(context.Background(), basaltic.WithAccessToken("t"))
	if err != nil {
		t.Fatalf("NewConfig: %v", err)
	}
	if _, err := cfg.EndpointResolver.ResolveEndpoint("compute", ""); err == nil {
		t.Error("compute resolved with no region, want a refusal")
	}
	if _, err := cfg.EndpointResolver.ResolveEndpoint("iam", ""); err != nil {
		t.Errorf("iam is global and should resolve without a region: %v", err)
	}
}
