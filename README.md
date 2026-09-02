# Basaltic Go SDK

The official Go client for the [Basaltic](https://basaltic.sh) cloud platform.

```bash
go get github.com/basaltic-sh/sdk-go
```

Requires Go 1.23 or later. It has no dependencies outside the standard library.

## Quick start

```go
package main

import (
	"context"
	"fmt"
	"log"

	basaltic "github.com/basaltic-sh/sdk-go"
	"github.com/basaltic-sh/sdk-go/compute"
)

func main() {
	ctx := context.Background()

	cfg, err := basaltic.NewConfig(ctx,
		basaltic.WithClientCredentials(keyID, secret),
		basaltic.WithRegion("sa-saopaulo-1"),
	)
	if err != nil {
		log.Fatal(err)
	}

	c := compute.New(cfg)

	inst, err := c.CreateInstance(ctx, &compute.InstanceCreateRequest{
		Name:     "web-01",
		FlavorID: flavorID,
		// Optional fields in a request body are pointers, so that sending a
		// zero value stays distinguishable from leaving a field alone.
		// basaltic.String, .Int, .Bool and .Ptr are there for the literals.
		ImageID: basaltic.String("debian-13"),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(inst.ID, inst.VMState)
}
```

With no options, `NewConfig` reads the environment:

| Variable | Meaning |
| --- | --- |
| `BASALTIC_ACCESS_KEY_ID` | Service account access key id |
| `BASALTIC_SECRET_ACCESS_KEY` | Service account secret |
| `BASALTIC_ACCESS_TOKEN` | A bearer token to use directly, instead of the key pair |
| `BASALTIC_REGION` | Region for regional services |
| `BASALTIC_ACCOUNT_ID` | Account to act on |

Options beat the environment, so a program can take its region from the
environment and still override it at one call site.

## Services

One package per service. Import only what you use.

| Package | Scope | Package | Scope |
| --- | --- | --- | --- |
| `compute` | regional | `iam` | global |
| `network` | regional | `dns` | global |
| `storage` | regional | `billing` | global |
| `database` | regional | `audit` | global |
| `loadbalancer` | regional | `quota` | global |
| `kms` | regional | | |
| `secrets` | regional | | |
| `certificate` | regional | | |
| `telemetry` | regional | | |

Regional services need a region; global ones ignore it.

**Object storage is not here.** The S3-compatible endpoint speaks AWS SigV4
and nothing else, so use an AWS SDK against it with the same access key pair.
The `storage` package in this SDK covers block volumes, snapshots and the
native bucket API, not the S3 protocol.

## Authentication

The platform authenticates with OAuth 2.0 bearer tokens. `WithClientCredentials`
exchanges a service account's access key pair for a short-lived token at the
IAM token endpoint and refreshes it before expiry — a long-running process
performs about one exchange an hour, not one per request. Concurrent calls
arriving while an exchange is in flight wait for that exchange rather than
starting their own.

The access key pair stays the one long-lived credential a service account has.
What the SDK changes is how it is presented.

Share one `Config` across every service client. They then share a token cache,
so authenticating a program that talks to six services still costs one
exchange:

```go
cfg, err := basaltic.NewConfig(ctx, basaltic.WithClientCredentials(keyID, secret))

instances := compute.New(cfg)
networks  := network.New(cfg)
volumes   := storage.New(cfg)
```

If something upstream already holds a token — an assumed-role session, or an
in-VM agent reading one from instance metadata — pass it with
`WithAccessToken`, or implement `TokenSource` for anything the SDK cannot mint
itself.

## Pagination

List operations return one page:

```go
page, err := c.ListInstances(ctx, &compute.ListInstancesParams{
	Limit:   50,
	VMState: compute.VMStateRunning,
})
for _, inst := range page.Items {
	fmt.Println(inst.Name)
}
```

Or walk every page with an iterator:

```go
for inst, err := range c.ListInstancesAll(ctx, nil) {
	if err != nil {
		return err
	}
	fmt.Println(inst.Name)
}
```

The iterator stops at the first error and yields it alongside a zero value, so
a loop that checks `err` on every step cannot mistake a truncated walk for a
complete one. Breaking out of the loop stops it; no further requests are made.

Page until `HasMore` is false, not until a page looks short — the platform
clamps an over-large limit rather than refusing it, so a short page is normal
and says nothing about whether more exist.

Query parameters are plain values and are simply omitted when zero, which is
what an empty filter means. Response and model types are plain values too —
only optional fields in request bodies are pointers.

## Errors

Every operation returns an `*basaltic.Error` when the platform answers with a
failure. Test for classes of failure with the helpers — **do not match on the
error code yourself**, however precise that looks. The platform names errors
per resource: `INSTANCE_NOT_FOUND`, `VOLUME_NOT_FOUND`,
`DATABASE_USER_NOT_FOUND`. There are 77 distinct codes ending in `_NOT_FOUND`
and 39 ending in `_EXISTS`, and adding a resource kind adds another. The HTTP
status is the stable signal, and it is what these helpers use.

```go
inst, err := c.GetInstance(ctx, id)
switch {
case basaltic.IsNotFound(err):
	// gone, or not visible to this credential — the platform answers the
	// same way for both
case basaltic.IsQuotaExceeded(err):
	// retrying will not clear this; raising the quota will
case basaltic.IsAccessDenied(err):
	// authenticated, but policy said no: a policy change, not a new key
case err != nil:
	return err
}
```

A quota refusal and an authorization refusal are both `403`, and their
remedies are opposite — raise the quota, versus change the policy. The two
helpers are mutually exclusive, so the order of the cases above does not
matter; what matters is asking `IsQuotaExceeded` at all rather than reading
every `403` as "permission denied".

`IsUnauthorized`, `IsConflict`, `IsInvalidInput`, `IsRateLimited` and
`IsTransient` round out the set. For the details, including the request id to quote when reporting a
problem:

```go
if apiErr, ok := basaltic.AsError(err); ok {
	log.Printf("%s failed: %s (request %s)", apiErr.OperationID, apiErr.Code, apiErr.RequestID)
}
```

## Retries and idempotency

Transport failures and the statuses that mean "try again" — 429, 500, 502,
503, 504 — are retried with exponential backoff and full jitter, honouring
`Retry-After` when the platform sends one. Four attempts by default; change it
with `WithRetry` or turn it off with `WithoutRetry`.

`GET`, `HEAD`, `PUT` and `DELETE` are safe to repeat and are retried.
**A `POST` is not**, because a create that timed out may already have
succeeded. Give it an idempotency key to make it both replay-safe and
retryable:

```go
key := basaltic.NewIdempotencyKey()

inst, err := c.CreateInstance(ctx, req, basaltic.WithIdempotencyKey(key))
```

Generate the key once and reuse it across retries of the same logical create —
a key generated per attempt defeats the point. The platform honours a key for
24 hours and replays the original outcome rather than creating a second
resource.

## Accounts

`WithAccountID` selects which account requests act on, sent as `X-Account-Id`.
Omit it to act on the credential's own account.

Reaching another account you are authorized for **requires** it: without the
header those resources answer `404`, not `403`, because an account you are not
acting in is one whose resources you cannot see. Override it for a single call
with `WithRequestAccountID`.

## Per-call options

Every operation takes a trailing `...basaltic.RequestOption`, which is what
keeps the signatures stable: a new knob arrives as an option rather than as a
changed method signature.

```go
var hdr http.Header

inst, err := c.CreateInstance(ctx, req,
	basaltic.WithIdempotencyKey(key),
	basaltic.WithRequestTimeout(10*time.Second),
	basaltic.WithResponseHeader(&hdr),
)
```

## Generated code

The service packages are generated from the platform's OpenAPI
specifications, which live in their own repository and are the single source
of truth. Nothing is vendored here — the generator reads them from a local
path:

```bash
make generate SPEC=/path/to/openapi
```

The hand-written half is the runtime in the root package: configuration,
authentication, the request pipeline, retries, errors and pagination. The
generated half is the 393 operations and their types, which are mechanical.

Generated files are committed, because `go get` has to work without anyone
running the generator. Do not edit them; change the specification, or the
generator, and regenerate. `make check-generated` fails when the committed
output no longer matches the specs.

## Releasing

Versions are git tags; there is nothing to bump in the repository. The bump
itself is derived from the exported API rather than from commit messages —
`scripts/apidiff.sh` compares against the last tag and says what changed and
what it implies, and CI runs it on every pull request.

The SDK is v0.x while the surface settles, so it carries no compatibility
promise yet. See [RELEASING.md](RELEASING.md), which also explains why a major
version is a deliberate edit here: from v2 onward Go changes the module path,
and a v2 tag without that change is a release nothing can import.

## Tests

```bash
make test          # unit tests, plus the generator's own
make integration   # read-only, against a live region; needs credentials
make apidiff       # what this change does to the exported API
```

There are no generated per-operation tests, deliberately. A test asserting
that `GetInstance` sends `GET /v1/instances/{instance_id}` would read its
expectation from the same specification node the code was emitted from, so an
emitter bug would produce a matching bug in the test.

What the suite does instead is walk all 393 generated methods by reflection
and assert properties the specification cannot express: that no placeholder
survives into a URL, that positional arguments reach it in the order they were
passed, that a credential is attached unless the operation is one of the four
that deliberately has none, and that every paginator terminates. Those are
checked against the SDK's own contract rather than against the specification,
so they catch real emitter faults.

## License

Apache 2.0. See [LICENSE](LICENSE).
