// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

// Package registry is the Registry API.
//
// KMS bindings use `kms_key`: a UUID, account-scoped name, or complete
// `crn:kms:<region>:<account>:key/<name>` in the authenticated account
// and current region. The key must be enabled, symmetric, and usable for
// encryption. Bindings retain the key's immutable identity; responses
// render a ResourceRef as `kms_key`. Invalid or foreign-scope CRNs are
// rejected with HTTP 400; missing keys return HTTP 404 on binding or
// cryptographic use. Metadata reads remain available for deleted
// bindings, with a nameless `kms_key` reference retaining the immutable
// key identity. Private container image registries.
//
// **This spec documents the control plane, not the pull path.** The
// management API uses `registry.<region>.basaltic.sh`; Docker and OCI
// clients use `cr.<region>.basaltic.cloud`. Everything below is the
// `/v1` REST surface — repositories, images, lifecycle policies,
// resource policies — authenticated with a platform OAuth 2.0 bearer
// token or a console session, like every other Basaltic API. Separately,
// `/v2` implements the [OCI Distribution
// Specification](https://github.com/opencontainers/distribution-spec)
// v1.1, which is what `docker`, `podman`, `buildah`, `crane`, `skopeo`,
// and every Kubernetes node's container runtime speak. That surface is
// not described here because it is not ours to describe: its paths,
// status codes, headers, and error codes are fixed by the spec, and a
// client that finds anything else fails.
//
// **Getting a docker credential.** `POST /v1/authorization-token`
// returns a short-lived password. Feed it to `docker login`; the daemon
// then exchanges it at `/v2/token` for per-repository bearer tokens on
// its own, which is where the actual authorization decision is made. The
// login credential carries no authority of its own — revoking a policy
// takes effect within one bearer token's lifetime (five minutes) rather
// than at the end of a twelve-hour login.
//
// **The account handle is part of the image reference.** A region serves
// one registry host, so an image lives at
// `cr.<region>.basaltic.cloud/<account_handle>/<repository>:<tag>`. That
// first path component is the only thing separating two tenants' `api`
// repositories, which is why every repository response includes a
// ready-made `uri` rather than leaving clients to assemble one.
//
// **Everything except a tag is immutable.** A blob is named by the
// sha256 of its bytes and a manifest by the sha256 of its own JSON, so
// pushing the same content twice is a no-op the client detects before it
// uploads a byte. A tag is the one mutable pointer — and a repository
// may forbid even that, with `tag_mutability: immutable`, so a deployed
// digest cannot be swapped under a running fleet.
//
// **Deleting by digest and deleting by tag are different.** Removing a
// digest removes the manifest and every tag pointing at it. Removing a
// tag leaves the manifest in place, untagged: another tag may still name
// it, and `docker rmi` of one tag must not destroy an image someone else
// is pulling by another. Untagged manifests are what the `untagged`
// lifecycle rule exists to sweep.
//
// **Layers are shared and counted once.** Blobs are deduplicated across
// every repository in an account, so a repository's `size_bytes` is the
// sum of the distinct blobs its manifests reference — not the sum of
// its images' sizes, which would count every shared base layer once per
// image. Deleting an image does not immediately free space: whether a
// layer is still needed is a question about every other repository in
// the account, and the answer comes from a background sweep.
//
// **Sharing a repository.** A repository is private to its account until
// you attach a resource policy. `PUT /v1/repositories/{id}/policy` is
// how another account — or, with principal `*`, the public — is
// granted pull access. Cross-account access requires BOTH the resource
// policy and the caller's own IAM policy to allow it, so one tenant
// cannot grant another tenant's employee more than that employee's
// organization permits.
//
// Build a client from a shared [basaltic.Config]:
//
//	c := registry.New(cfg)
//
// Clients are safe for concurrent use.
package registry

import (
	basaltic "github.com/basaltic-sh/sdk-go"
)

// ServiceID is the short name the SDK addresses this service by. Use it
// with [basaltic.WithServiceEndpoint] to point this one client elsewhere.
const ServiceID = "registry"

// endpointTemplate is the server URL this service's specification
// declares. Any {region} in it is substituted per request.
const endpointTemplate = "https://registry.{region}.basaltic.sh"

func init() { basaltic.RegisterServiceEndpoint(ServiceID, endpointTemplate) }

// Client calls the Registry API.
//
// Build one with [New]. It is safe for concurrent use.
type Client struct {
	rt *basaltic.Client
}

// New builds a registry client from a shared configuration.
//
// Share one [basaltic.Config] across every service client: they then
// share a token, so authenticating costs one exchange rather than one
// per service.
func New(cfg *basaltic.Config) *Client {
	return &Client{rt: basaltic.NewClient(cfg, ServiceID)}
}

// Transport returns the underlying transport, for reaching an endpoint
// this package does not generate. See [basaltic.Client.Do].
func (c *Client) Transport() *basaltic.Client { return c.rt }
