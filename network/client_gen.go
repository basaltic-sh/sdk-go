// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

// Package network is the Network API.
//
// The VPC surface — networks and subnets, interfaces, route tables,
// security groups, internet, NAT and egress-only gateways, and floating
// IPs.
//
// Relationship inputs use one reference field: UUID, CRN, or an exact
// name when the request supplies its required parent scope. Subnets and
// route tables are VPC-scoped; interfaces are subnet-scoped. Nested CRNs
// use vpc/<vpc>/subnet/<subnet>, vpc/<vpc>/route-table/<table>, and
// vpc/<vpc>/subnet/<subnet>/interface/<interface>. Names are immutable.
// Unknown request fields and list parameters are rejected. Supplied
// empty references are invalid and failed lookups never fall back to
// another reference kind. Every list accepts exact name and crn filters.
//
// Build a client from a shared [basaltic.Config]:
//
//	c := network.New(cfg)
//
// Clients are safe for concurrent use.
package network

import (
	basaltic "github.com/basaltic-sh/sdk-go"
)

// ServiceID is the short name the SDK addresses this service by. Use it
// with [basaltic.WithServiceEndpoint] to point this one client elsewhere.
const ServiceID = "network"

// endpointTemplate is the server URL this service's specification
// declares. Any {region} in it is substituted per request.
const endpointTemplate = "https://network.{region}.basaltic.sh"

func init() { basaltic.RegisterServiceEndpoint(ServiceID, endpointTemplate) }

// Client calls the Network API.
//
// Build one with [New]. It is safe for concurrent use.
type Client struct {
	rt *basaltic.Client
}

// New builds a network client from a shared configuration.
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
