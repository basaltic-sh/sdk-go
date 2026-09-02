//go:build integration

// Integration tests run against a live region. They are read-only: every call
// is a GET, nothing is created, changed or deleted.
//
//	BASALTIC_ACCESS_KEY_ID=... BASALTIC_SECRET_ACCESS_KEY=... \
//	BASALTIC_REGION=sa-saopaulo-1 BASALTIC_ACCOUNT_ID=... \
//	go test -tags integration -v .
//
// They exist because the unit tests cannot see the one thing that matters
// most: whether the SDK's idea of the API matches the API. The first live run
// of this suite found that error codes are named per resource
// (INSTANCE_NOT_FOUND, not NOT_FOUND), which no test written against the
// specification would have caught.
package basaltic_test

import (
	"context"
	"os"
	"testing"
	"time"

	basaltic "github.com/basaltic-sh/sdk-go"
	"github.com/basaltic-sh/sdk-go/compute"
	"github.com/basaltic-sh/sdk-go/iam"
	"github.com/basaltic-sh/sdk-go/network"
)

func liveConfig(t *testing.T) *basaltic.Config {
	t.Helper()
	keyID := os.Getenv(basaltic.EnvAccessKeyID)
	secret := os.Getenv(basaltic.EnvSecretAccessKey)
	if keyID == "" || secret == "" {
		t.Skipf("set %s and %s to run integration tests", basaltic.EnvAccessKeyID, basaltic.EnvSecretAccessKey)
	}
	cfg, err := basaltic.NewConfig(context.Background(),
		basaltic.WithUserAgent("sdk-go-integration/1"),
	)
	if err != nil {
		t.Fatalf("NewConfig: %v", err)
	}
	if cfg.Region == "" {
		t.Skipf("set %s to run integration tests", basaltic.EnvRegion)
	}
	return cfg
}

func liveContext(t *testing.T) context.Context {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	t.Cleanup(cancel)
	return ctx
}

// The client-credentials exchange against the real IAM token endpoint.
func TestLiveTokenExchange(t *testing.T) {
	cfg := liveConfig(t)
	ctx := liveContext(t)

	token, err := cfg.TokenSource.Token(ctx)
	if err != nil {
		t.Fatalf("token exchange: %v", err)
	}
	if token == "" {
		t.Fatal("token exchange returned an empty token")
	}

	// A second call must be served from cache rather than exchanging again.
	again, err := cfg.TokenSource.Token(ctx)
	if err != nil {
		t.Fatalf("second token call: %v", err)
	}
	if again != token {
		t.Error("the second call exchanged a new token instead of using the cached one")
	}
}

// Global and regional endpoints resolve and answer.
func TestLiveEndpointRouting(t *testing.T) {
	cfg := liveConfig(t)
	ctx := liveContext(t)

	// iam is global: no {region} in its endpoint.
	regions, err := iam.New(cfg).ListRegions(ctx)
	if err != nil {
		t.Fatalf("iam.ListRegions (global): %v", err)
	}
	if len(regions.Regions) == 0 {
		t.Error("iam.ListRegions returned no regions")
	}

	// compute and network are regional, on separate hosts.
	if _, err := compute.New(cfg).ListFlavors(ctx, nil); err != nil {
		t.Fatalf("compute.ListFlavors (regional): %v", err)
	}
	if _, err := network.New(cfg).ListVPCs(ctx, nil); err != nil {
		t.Fatalf("network.ListVPCs (regional): %v", err)
	}
}

// A real multi-page walk over opaque cursors.
func TestLivePagination(t *testing.T) {
	cfg := liveConfig(t)
	ctx := liveContext(t)
	c := iam.New(cfg)

	// A limit of one forces a page per item, so the marker round-trip is
	// exercised rather than assumed.
	first, err := c.ListPolicies(ctx, &iam.ListPoliciesParams{Limit: 1})
	if err != nil {
		t.Fatalf("iam.ListPolicies: %v", err)
	}
	if !first.HasMore {
		t.Skip("fewer than two policies in this account; nothing to page over")
	}

	seen := map[string]bool{}
	count := 0
	for pol, err := range c.ListPoliciesAll(ctx, &iam.ListPoliciesParams{Limit: 1}) {
		if err != nil {
			t.Fatalf("iam.ListPoliciesAll: %v", err)
		}
		if seen[pol.ID] {
			t.Fatalf("policy %s yielded twice — the cursor is not advancing", pol.ID)
		}
		seen[pol.ID] = true
		count++
		if count > 1000 {
			t.Fatal("the paginator did not terminate")
		}
	}
	if count < 2 {
		t.Errorf("walked %d policies but the first page reported more", count)
	}
}

// A real failure, classified. This is the test that found the SDK was
// matching error codes by name: the platform answers INSTANCE_NOT_FOUND.
func TestLiveNotFoundIsClassified(t *testing.T) {
	cfg := liveConfig(t)
	ctx := liveContext(t)

	_, err := compute.New(cfg).GetInstance(ctx, "00000000-0000-4000-8000-000000000000")
	if err == nil {
		t.Fatal("getting a nonexistent instance succeeded")
	}
	if !basaltic.IsNotFound(err) {
		t.Fatalf("IsNotFound() = false for a missing instance: %v", err)
	}
	apiErr, ok := basaltic.AsError(err)
	if !ok {
		t.Fatalf("error is %T, want *basaltic.Error", err)
	}
	if apiErr.RequestID == "" {
		t.Error("the platform's error carried no request id")
	}
	// The point of classifying on status: the code is per-resource and the
	// generic NOT_FOUND is not what comes back.
	t.Logf("classified %s (http %d, request %s)", apiErr.Code, apiErr.StatusCode, apiErr.RequestID)
}

// Argument validation happens before anything reaches the network.
func TestLiveEmptyPathArgumentIsRefusedLocally(t *testing.T) {
	cfg := liveConfig(t)
	ctx := liveContext(t)

	// "" would turn a get into a list against a different endpoint.
	_, err := compute.New(cfg).GetInstance(ctx, "")
	if err == nil {
		t.Fatal("an empty instance id was accepted")
	}
	if _, isAPI := basaltic.AsError(err); isAPI {
		t.Error("the empty id reached the platform; it should be refused locally")
	}
}
