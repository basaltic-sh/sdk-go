// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

// Package email is the Email API.
//
// Regional API for sending and receiving email, identities and DNS
// records, templates, configuration sets, IP pools, recipient rules,
// message content and attachments.
//
// Identities must be verified separately in every region. Create an
// identity, publish its dns_records, then request verification. Only
// verified identities can send or enable receiving.
//
// Request references accept a UUID, an account-unique name or a CRN of
// the expected resource type. Names are immutable. Malformed references
// or the wrong resource type/region return 400; missing or
// foreign-account references return 404. Resource paths use UUIDs. List
// filters are exact, account-scoped and applied before pagination;
// missing or foreign references return empty results.
//
// Sent and received message content expires after 30 days. Attachment
// objects follow their bucket lifecycle independently. Use Telemetry for
// aggregate metrics and the diagnostic endpoints for individual events
// and DMARC report records.
//
// Build a client from a shared [basaltic.Config]:
//
//	c := email.New(cfg)
//
// Clients are safe for concurrent use.
package email

import (
	basaltic "github.com/basaltic-sh/sdk-go"
)

// ServiceID is the short name the SDK addresses this service by. Use it
// with [basaltic.WithServiceEndpoint] to point this one client elsewhere.
const ServiceID = "email"

// endpointTemplate is the server URL this service's specification
// declares. Any {region} in it is substituted per request.
const endpointTemplate = "https://email.{region}.basaltic.sh"

func init() { basaltic.RegisterServiceEndpoint(ServiceID, endpointTemplate) }

// Client calls the Email API.
//
// Build one with [New]. It is safe for concurrent use.
type Client struct {
	rt *basaltic.Client
}

// New builds a email client from a shared configuration.
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
