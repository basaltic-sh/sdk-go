// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

// Package notifications is the Notifications API.
//
// Publish/subscribe notification topics with email and webhook delivery.
//
// A topic is a fan-out root. Subscriptions bind it to endpoints;
// publishing to the topic delivers to every confirmed subscription whose
// filter policy matches the message.
//
// **Publishing is asynchronous.** Publish returns 202 as soon as the
// notification is durable and queued: no delivery has happened yet.
// Fan-out into per-subscription deliveries runs on a worker, so publish
// latency does not grow with subscriber count. Poll `GET
// /v1/topics/{topic_id}/messages/{message_id}` for the fan-out state and
// `GET /v1/topics/{topic_id}/deliveries?message=...` for per-endpoint
// outcomes.
//
// **Subscriptions must be confirmed.** Subscribe creates the
// subscription in `pending_confirmation` and sends a token to the
// endpoint — a link for email, a `subscription_confirmation` POST for
// webhooks. Nothing is delivered until the endpoint owner presents that
// token back. The confirmation window is 72 hours; unconfirmed
// subscriptions are reclaimed after it. If the handshake never arrived,
// `POST /v1/topics/{topic_id}/subscriptions/{id}/resend-confirmation`
// mints a fresh one — the original token is not recoverable, only its
// hash is stored.
//
// **Webhook endpoints must be publicly routable.** Endpoints resolving
// to loopback, RFC1918, link-local, or unique-local addresses are
// refused at subscribe time and again at delivery time.
//
// **Webhook deliveries are signed.** Each POST carries
// `X-Basaltic-Signature: v1,t=<unix-seconds>,s=<hex>`, an HMAC-SHA256
// over `"<t>.<raw-request-body>"` keyed by the subscription's signing
// secret. The secret is returned once by Subscribe and is readable
// afterwards via the signing-secret endpoint. Verify by recomputing the
// MAC over the raw body you received — not a re-serialised copy —
// and reject timestamps more than five minutes from now.
//
// Build a client from a shared [basaltic.Config]:
//
//	c := notifications.New(cfg)
//
// Clients are safe for concurrent use.
package notifications

import (
	basaltic "github.com/basaltic-sh/sdk-go"
)

// ServiceID is the short name the SDK addresses this service by. Use it
// with [basaltic.WithServiceEndpoint] to point this one client elsewhere.
const ServiceID = "notifications"

// endpointTemplate is the server URL this service's specification
// declares. Any {region} in it is substituted per request.
const endpointTemplate = "https://notifications.{region}.basaltic.sh"

func init() { basaltic.RegisterServiceEndpoint(ServiceID, endpointTemplate) }

// Client calls the Notifications API.
//
// Build one with [New]. It is safe for concurrent use.
type Client struct {
	rt *basaltic.Client
}

// New builds a notifications client from a shared configuration.
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
