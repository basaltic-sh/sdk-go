// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

// Package telemetry is the Telemetry API.
//
// Regional telemetry API — log groups + log ingest/search, ClickHouse
// metrics (remote_write + structured query with Prometheus-compatible
// JSON envelopes), span ingest, and trace search/detail + per-account
// trace settings.
//
// Log-group request references accept UUIDs, CRNs or exact
// account-scoped names, classified by syntax without lookup fallback.
// CRNs must match the authenticated account handle and serving region.
// HTTP and gRPC OTLP logs use basaltic.log_group, falling back to
// service.name, with the same reference contract. Trace/span
// identifiers, metric labels and region search dimensions are protocol
// values, independent of resource resolution.
//
// Build a client from a shared [basaltic.Config]:
//
//	c := telemetry.New(cfg)
//
// Clients are safe for concurrent use.
package telemetry

import (
	basaltic "github.com/basaltic-sh/sdk-go"
)

// ServiceID is the short name the SDK addresses this service by. Use it
// with [basaltic.WithServiceEndpoint] to point this one client elsewhere.
const ServiceID = "telemetry"

// endpointTemplate is the server URL this service's specification
// declares. Any {region} in it is substituted per request.
const endpointTemplate = "https://telemetry.{region}.basaltic.sh"

func init() { basaltic.RegisterServiceEndpoint(ServiceID, endpointTemplate) }

// Client calls the Telemetry API.
//
// Build one with [New]. It is safe for concurrent use.
type Client struct {
	rt *basaltic.Client
}

// New builds a telemetry client from a shared configuration.
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
