// Package basaltic is the Go SDK for the Basaltic cloud platform.
//
// The root package holds everything shared across services: configuration,
// authentication, the HTTP pipeline, errors and pagination. One package per
// service holds that service's operations and types, and those are generated
// from the platform's OpenAPI specifications.
//
// # Getting started
//
// A [Config] carries credentials, the target region and endpoint overrides.
// Build one once and share it across every service client; clients are safe
// for concurrent use and a shared Config means one token exchange serves all
// of them.
//
//	cfg, err := basaltic.NewConfig(ctx,
//		basaltic.WithClientCredentials(keyID, secret),
//		basaltic.WithRegion("sa-saopaulo-1"),
//	)
//	if err != nil {
//		return err
//	}
//
//	c := compute.New(cfg)
//	inst, err := c.GetInstance(ctx, instanceID)
//
// With no options, [NewConfig] reads the environment — see [NewConfig] for
// the variables it consults.
//
// # Authentication
//
// The platform authenticates with OAuth 2.0 bearer tokens. The usual path is
// [WithClientCredentials]: the SDK exchanges a service account's access key
// pair for a short-lived token at the IAM token endpoint and refreshes it
// before expiry, so a long-running process never has to think about it. The
// key pair stays the one long-lived credential; only its presentation
// changes.
//
// The S3-compatible object storage endpoint is the exception. It speaks AWS
// SigV4 and nothing else, so use the same key pair with an AWS SDK there.
// This SDK does not wrap it.
//
// # Errors
//
// Every operation returns an [*Error] when the platform answers with a
// failure status. It carries the platform's error code, message and request
// id — quote the request id when reporting a problem.
//
// Test for classes of failure with [IsNotFound], [IsAccessDenied] and their
// siblings. Do not match on the error code yourself, however precise that
// looks: the platform names errors per resource, so there are 77 distinct
// codes ending in _NOT_FOUND and 39 ending in _EXISTS, and adding a resource
// kind adds another. The HTTP status is the stable signal, and it is what
// these helpers use — a new resource kind adds a code, never a status.
//
// One status is not enough on its own. A quota refusal and an authorization
// refusal are both 403, and their remedies are opposite: raise the quota
// versus change the policy. [IsQuotaExceeded] is therefore carved out of
// [IsAccessDenied] by code, and is the one place the SDK reads one. The two
// are mutually exclusive, so the order of the cases does not matter — what
// matters is asking about quota at all, rather than reading every 403 as
// "permission denied".
//
// # Pagination
//
// List operations return one [Page] at a time. Where the platform paginates,
// each list operation also has an "All" form returning an iterator that walks
// every page for you:
//
//	for inst, err := range c.ListInstancesAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		fmt.Println(inst.Name)
//	}
package basaltic
