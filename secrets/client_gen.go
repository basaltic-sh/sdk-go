// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

// Package secrets is the Secrets API.
//
// Versioned application secrets with KMS-backed envelope encryption.
//
// Secrets are regional. The value of every version is encrypted at rest
// under a KMS key — only opaque ciphertext is persisted. By default a
// secret uses the platform-managed key; pass kms_key_id on CreateSecret
// to bind it to one of your own KMS keys instead (the key is fixed for
// the secret's life, and every version is encrypted under it). Each
// PutSecretValue allocates a new monotonically-increasing version number
// and flips is_current on the previous row; older versions remain
// readable by explicit version query.
//
// Soft delete: DeleteSecret moves the secret into a recovery window
// (default 7 days, configurable 1-30). RestoreSecret exits the window.
// The hard-purge sweeper removes the row + every version once now() >=
// scheduled_purge_at.
//
// Values move on the wire as base64 to survive arbitrary binary payloads
// (max 64 KiB to mirror AWS Secrets Manager).
//
// Build a client from a shared [basaltic.Config]:
//
//	c := secrets.New(cfg)
//
// Clients are safe for concurrent use.
package secrets

import (
	basaltic "github.com/basaltic-sh/sdk-go"
)

// ServiceID is the short name the SDK addresses this service by. Use it
// with [basaltic.WithServiceEndpoint] to point this one client elsewhere.
const ServiceID = "secrets"

// endpointTemplate is the server URL this service's specification
// declares. Any {region} in it is substituted per request.
const endpointTemplate = "https://secrets.{region}.basaltic.sh"

func init() { basaltic.RegisterServiceEndpoint(ServiceID, endpointTemplate) }

// Client calls the Secrets API.
//
// Build one with [New]. It is safe for concurrent use.
type Client struct {
	rt *basaltic.Client
}

// New builds a secrets client from a shared configuration.
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
