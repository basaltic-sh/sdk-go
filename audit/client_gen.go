// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

// Package audit is the Audit API.
//
// Read-only access to the account's audit trail. Every mutating call to
// a Basaltic API records who made it, what it touched and whether it
// succeeded — failed attempts included. One trail per account, across
// every region.
//
// Build a client from a shared [basaltic.Config]:
//
//	c := audit.New(cfg)
//
// Clients are safe for concurrent use.
package audit

import (
	basaltic "github.com/basaltic-sh/sdk-go"
)

// ServiceID is the short name the SDK addresses this service by. Use it
// with [basaltic.WithServiceEndpoint] to point this one client elsewhere.
const ServiceID = "audit"

// endpointTemplate is the server URL this service's specification
// declares. Any {region} in it is substituted per request.
const endpointTemplate = "https://audit.basaltic.sh"

func init() { basaltic.RegisterServiceEndpoint(ServiceID, endpointTemplate) }

// Client calls the Audit API.
//
// Build one with [New]. It is safe for concurrent use.
type Client struct {
	rt *basaltic.Client
}

// New builds a audit client from a shared configuration.
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
