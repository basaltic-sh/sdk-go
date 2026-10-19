// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

// Package queue is the Queue API.
//
// KMS bindings use `kms_key`: a UUID, account-scoped name, or complete
// `crn:kms:<region>:<account>:key/<name>` in the authenticated account
// and current region. The key must be enabled, symmetric, and usable for
// encryption. Bindings retain the key's immutable identity; responses
// render its CRN. Invalid or foreign-scope CRNs are rejected with HTTP
// 400; missing keys return HTTP 404 on binding or cryptographic use.
// Metadata reads remain available for deleted bindings, with a `kms_key`
// ResourceRef containing its stored ID and CRN, without a name. Durable
// message queues with at-least-once delivery.
//
// **A receive is a lease, not a read.** `ReceiveMessage` hides the
// messages it returns for the queue's visibility timeout and hands back
// a *receipt handle*. The message comes back — to you or to another
// consumer — unless you delete it with that handle before the lease
// lapses. Everything else follows from this: a consumer that crashes
// mid-work loses nothing, and a message no consumer can process is
// caught by `receive_count` crossing the redrive threshold.
//
// **Receipt handles are per-lease, not per-message.** A new one is
// minted on every receive. Presenting a handle from a lease that has
// already lapsed fails with `QUEUE_RECEIPT_HANDLE_EXPIRED` rather than
// deleting a message that has since been handed to someone else — so a
// slow consumer finds out it was slow instead of silently destroying
// another's work.
//
// **Two kinds of queue, fixed at creation.** A *standard* queue delivers
// at least once with best-effort ordering and races consumers freely. A
// *FIFO* queue — whose name must end in `.fifo` — delivers strictly
// in order within a `message_group_id` and refuses to hand out a group's
// next message while an earlier one is still in flight; sends inside a
// five-minute window that repeat a `message_deduplication_id` collapse
// to one message. Ordering across groups is undefined, which is what
// lets groups run concurrently. The kind cannot be changed: a queue
// holding messages written under one set of rules cannot start honouring
// the other.
//
// **Long polling.** `wait_time_seconds` (up to 20) blocks the receive
// until a message arrives rather than returning empty immediately. It
// costs one request instead of the many an empty polling loop would
// make, and it returns as soon as a producer sends.
//
// **Message bodies are encrypted at rest.** Each queue has a data key
// that seals its messages; that key is itself wrapped by a platform key
// or, if you set `kms_key`, by your own. The unwrapped data key is
// cached for `kms_data_key_reuse_period_seconds` — so for up to that
// long after you disable your key, sends and receives on the queue keep
// working. Lower it to tighten that window at the cost of more KMS
// round-trips.
//
// Build a client from a shared [basaltic.Config]:
//
//	c := queue.New(cfg)
//
// Clients are safe for concurrent use.
package queue

import (
	basaltic "github.com/basaltic-sh/sdk-go"
)

// ServiceID is the short name the SDK addresses this service by. Use it
// with [basaltic.WithServiceEndpoint] to point this one client elsewhere.
const ServiceID = "queue"

// endpointTemplate is the server URL this service's specification
// declares. Any {region} in it is substituted per request.
const endpointTemplate = "https://queue.{region}.basaltic.sh"

func init() { basaltic.RegisterServiceEndpoint(ServiceID, endpointTemplate) }

// Client calls the Queue API.
//
// Build one with [New]. It is safe for concurrent use.
type Client struct {
	rt *basaltic.Client
}

// New builds a queue client from a shared configuration.
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
