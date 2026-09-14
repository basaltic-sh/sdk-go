// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

// Package iam is the IAM API.
//
// Identity for the platform: organizations, accounts, users, groups,
// service accounts, roles and policies, together with the sign-in flows
// and the temporary STS credentials every other Basaltic API
// authenticates against.
//
// Relationship inputs are classified once as CRN, UUID or name. A lookup
// never falls back to another syntax. Policy, role and group names are
// immutable. Their organization CRNs use empty region and account slots;
// system policies use crn:iam::platform:policy/<name>. Bare policy names
// select only organization policies. UUID path and response identity
// fields retain their existing meaning.
//
// Every resource list accepts exact name and crn filters, combined with
// AND before pagination. A mismatched or foreign CRN returns an empty
// page. Organization-scoped CRNs resolve only in the authenticated
// organization.
//
// Build a client from a shared [basaltic.Config]:
//
//	c := iam.New(cfg)
//
// Clients are safe for concurrent use.
package iam

import (
	basaltic "github.com/basaltic-sh/sdk-go"
)

// ServiceID is the short name the SDK addresses this service by. Use it
// with [basaltic.WithServiceEndpoint] to point this one client elsewhere.
const ServiceID = "iam"

// endpointTemplate is the server URL this service's specification
// declares. Any {region} in it is substituted per request.
const endpointTemplate = "https://iam.basaltic.sh"

func init() { basaltic.RegisterServiceEndpoint(ServiceID, endpointTemplate) }

// Client calls the IAM API.
//
// Build one with [New]. It is safe for concurrent use.
type Client struct {
	rt *basaltic.Client
}

// New builds a iam client from a shared configuration.
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
