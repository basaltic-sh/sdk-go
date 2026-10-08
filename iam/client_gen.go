// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

// Package iam is the IAM API.
//
// Account identity and access management: service accounts, roles,
// account policies, and account-scoped temporary sessions.
// Authentication and personal sign-in remain in IAM. Organizations,
// accounts, users, groups, and organization policies are managed by the
// Workspace API.
//
// Roles and custom policies belong to the selected account. Their global
// CRNs use an empty region and the owning account handle. Shared system
// policies use crn:iam:::policy/<name>. Relationship inputs are
// classified once as CRN, UUID, or name, without syntax fallback.
//
// AssumeRole resolves the target role's owning account. The caller needs
// source permission and the target role must trust the caller; the
// resulting session uses only the target role's permissions, subject to
// boundaries and session restrictions.
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
