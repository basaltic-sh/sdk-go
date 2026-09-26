// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

// Package workspace is the Workspace API.
//
// Organization management: organizations, accounts, human users,
// users-only groups, and organization policies. Account IAM identities
// may receive explicitly delegated organization policies through this
// API. Personal authentication remains at the IAM endpoint.
//
// Organization resources are global and are resolved in the
// authenticated organization. Canonical CRNs are
// crn:workspace:::organization/<organization-uuid>/<type>/<name-or-uuid>.
// Organization policies are separate from account policies; shared
// system policies use crn:workspace:::policy/<name>.
//
// Build a client from a shared [basaltic.Config]:
//
//	c := workspace.New(cfg)
//
// Clients are safe for concurrent use.
package workspace

import (
	basaltic "github.com/basaltic-sh/sdk-go"
)

// ServiceID is the short name the SDK addresses this service by. Use it
// with [basaltic.WithServiceEndpoint] to point this one client elsewhere.
const ServiceID = "workspace"

// endpointTemplate is the server URL this service's specification
// declares. Any {region} in it is substituted per request.
const endpointTemplate = "https://workspace.basaltic.sh"

func init() { basaltic.RegisterServiceEndpoint(ServiceID, endpointTemplate) }

// Client calls the Workspace API.
//
// Build one with [New]. It is safe for concurrent use.
type Client struct {
	rt *basaltic.Client
}

// New builds a workspace client from a shared configuration.
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
