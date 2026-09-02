package basaltic_test

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	basaltic "github.com/basaltic-sh/sdk-go"
	"github.com/basaltic-sh/sdk-go/compute"
	"github.com/basaltic-sh/sdk-go/network"
	"github.com/basaltic-sh/sdk-go/storage"
)

// Authenticate as a service account and create an instance.
func Example() {
	ctx := context.Background()

	cfg, err := basaltic.NewConfig(ctx,
		basaltic.WithClientCredentials(os.Getenv("BASALTIC_ACCESS_KEY_ID"), os.Getenv("BASALTIC_SECRET_ACCESS_KEY")),
		basaltic.WithRegion("sa-saopaulo-1"),
	)
	if err != nil {
		log.Fatal(err)
	}

	c := compute.New(cfg)

	inst, err := c.CreateInstance(ctx, &compute.InstanceCreateRequest{
		Name:     "web-01",
		FlavorID: "c8b0a4f2-1d3e-4a5b-8c7d-9e0f1a2b3c4d",
		ImageID:  basaltic.String("debian-13"),
		Tags:     compute.Tags{"environment": "production"},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(inst.ID, inst.VMState)
}

// One Config serves every service client, so authenticating a program that
// talks to three services still costs one token exchange.
func Example_sharedConfig() {
	ctx := context.Background()

	cfg, err := basaltic.NewConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	instances := compute.New(cfg)
	networks := network.New(cfg)
	volumes := storage.New(cfg)

	_, _, _ = instances, networks, volumes
}

// Walk every page of a list operation.
func ExampleClient_do_pagination() {
	ctx := context.Background()
	cfg, _ := basaltic.NewConfig(ctx)
	c := compute.New(cfg)

	for inst, err := range c.ListInstancesAll(ctx, &compute.ListInstancesParams{
		VMState: compute.VMStateRunning,
	}) {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(inst.Name, inst.PrimaryIP)
	}
}

// Fetch a single page and page manually.
func ExamplePage() {
	ctx := context.Background()
	cfg, _ := basaltic.NewConfig(ctx)
	c := compute.New(cfg)

	params := &compute.ListInstancesParams{Limit: 50, Name: "web-*"}
	for {
		page, err := c.ListInstances(ctx, params)
		if err != nil {
			log.Fatal(err)
		}
		for _, inst := range page.Items {
			fmt.Println(inst.Name)
		}
		// Page until HasMore is false, not until a page looks short: the
		// platform clamps an over-large limit rather than refusing it.
		if !page.HasMore {
			break
		}
		params.Marker = page.Marker
	}
}

// Distinguish the failures that need different responses.
func ExampleAsError() {
	ctx := context.Background()
	cfg, _ := basaltic.NewConfig(ctx)
	c := compute.New(cfg)

	inst, err := c.GetInstance(ctx, "i-does-not-exist")
	switch {
	case basaltic.IsNotFound(err):
		// Gone, or not visible to this credential. The platform answers the
		// same way for both, deliberately.
		fmt.Println("no such instance")
	case basaltic.IsAccessDenied(err):
		// Authenticated, but policy said no. The fix is a policy change, not
		// a new credential.
		fmt.Println("not permitted")
	case basaltic.IsQuotaExceeded(err):
		// Retrying will not clear this; raising the quota will.
		fmt.Println("at a quota limit")
	case err != nil:
		if apiErr, ok := basaltic.AsError(err); ok {
			log.Fatalf("%s failed: %s (request %s)", apiErr.OperationID, apiErr.Code, apiErr.RequestID)
		}
		log.Fatal(err)
	default:
		fmt.Println(inst.ID)
	}
}

// Make a create replay-safe, which also makes it retryable.
func ExampleWithIdempotencyKey() {
	ctx := context.Background()
	cfg, _ := basaltic.NewConfig(ctx)
	c := compute.New(cfg)

	// Generate the key once and reuse it across retries of the same logical
	// create. A key generated per attempt defeats the point.
	key := basaltic.NewIdempotencyKey()

	inst, err := c.CreateInstance(ctx, &compute.InstanceCreateRequest{
		Name:     "web-01",
		FlavorID: "c8b0a4f2-1d3e-4a5b-8c7d-9e0f1a2b3c4d",
	}, basaltic.WithIdempotencyKey(key))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(inst.ID)
}

// Read rate-limit state off a throttled response.
func ExampleError_RateLimit() {
	ctx := context.Background()
	cfg, _ := basaltic.NewConfig(ctx)
	c := compute.New(cfg)

	_, err := c.ListInstances(ctx, nil)
	var apiErr *basaltic.Error
	if errors.As(err, &apiErr) && basaltic.IsRateLimited(err) {
		limit, remaining, ok := apiErr.RateLimit()
		if ok {
			fmt.Printf("throttled: %d of %d left\n", remaining, limit)
		}
		if wait, ok := apiErr.RetryAfter(); ok {
			// Retrying before Retry-After is refused and extends the window.
			time.Sleep(wait)
		}
	}
}

// Stream an object's bytes rather than buffering them.
func ExampleClient_doStream() {
	ctx := context.Background()
	cfg, _ := basaltic.NewConfig(ctx)
	c := storage.New(cfg)

	body, err := c.GetObject(ctx, "reports", "2026/q1.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer body.Close()

	if _, err := io.Copy(os.Stdout, body); err != nil {
		log.Fatal(err)
	}
}

// Point the SDK at a non-production deployment.
func ExampleWithDomain() {
	ctx := context.Background()

	// compute becomes https://compute.sa-saopaulo-1.cloud.example.dev
	cfg, err := basaltic.NewConfig(ctx,
		basaltic.WithDomain("cloud.example.dev"),
		basaltic.WithRegion("sa-saopaulo-1"),
	)
	if err != nil {
		log.Fatal(err)
	}
	_ = compute.New(cfg)
}

// Capture the response headers of a successful call.
func ExampleWithResponseHeader() {
	ctx := context.Background()
	cfg, _ := basaltic.NewConfig(ctx)
	c := compute.New(cfg)

	var hdr http.Header
	inst, err := c.GetInstance(ctx, "i-1", basaltic.WithResponseHeader(&hdr))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(inst.ID, hdr.Get("X-Request-Id"))
}

// Authenticate with a token something upstream already obtained.
func ExampleWithAccessToken() {
	ctx := context.Background()

	cfg, err := basaltic.NewConfig(ctx,
		basaltic.WithAccessToken(os.Getenv("BASALTIC_ACCESS_TOKEN")),
		basaltic.WithRegion("sa-saopaulo-1"),
	)
	if err != nil {
		log.Fatal(err)
	}
	// The SDK cannot refresh a static token. Use WithClientCredentials
	// wherever the access key pair is available.
	_ = compute.New(cfg)
}
