// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package queue

import (
	"context"
	"iter"
	"net/url"
	"strconv"

	basaltic "github.com/basaltic-sh/sdk-go"
)

// DeleteMessageParams are the optional filters and pagination controls for
// [Client.DeleteMessage]. A nil *DeleteMessageParams sends none of them.
type DeleteMessageParams struct {
	// ReceiptHandle the handle returned by the receive that leased this message.
	ReceiptHandle string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *DeleteMessageParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.ReceiptHandle != "" {
		q.Set("receipt_handle", p.ReceiptHandle)
	}
	return q
}

// ListQueuesParams are the optional filters and pagination controls for
// [Client.ListQueues]. A nil *ListQueuesParams sends none of them.
type ListQueuesParams struct {
	// CRN exact queue CRN. Malformed, foreign-account or mismatched service,
	// type or region returns an empty page. Intersects with name.
	CRN string

	// Limit maximum items to return. A larger value is clamped to the maximum
	// rather than rejected, so page until `meta.has_more` is false.
	Limit int

	// Marker id of the last item on the previous page.
	Marker string

	// Name exact account-scoped queue name. Intersects with crn when both are
	// supplied.
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListQueuesParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.CRN != "" {
		q.Set("crn", p.CRN)
	}
	if p.Limit != 0 {
		q.Set("limit", strconv.Itoa(int(p.Limit)))
	}
	if p.Marker != "" {
		q.Set("marker", p.Marker)
	}
	if p.Name != "" {
		q.Set("name", p.Name)
	}
	return q
}

// withMarker copies p with the pagination cursor replaced, leaving the
// caller's value untouched across pages.
func (p *ListQueuesParams) withMarker(marker string) *ListQueuesParams {
	var out ListQueuesParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// PeekMessagesParams are the optional filters and pagination controls for
// [Client.PeekMessages]. A nil *PeekMessagesParams sends none of them.
type PeekMessagesParams struct {
	// CRN exact message CRN under this queue. Foreign or mismatched CRNs
	// select an empty page before pagination; malformed CRNs return 400.
	CRN string

	// Limit maximum items to return. A larger value is clamped to the maximum
	// rather than rejected, so page until `meta.has_more` is false.
	Limit int

	// Marker id of the last item on the previous page.
	Marker string

	// Visibility restrict the peek to one population. `available` can be received
	// now; `in_flight` is leased by a consumer; `delayed` has never been
	// delivered and its send-time delay has not elapsed.
	//
	//
	// One of: "available", "in_flight", "delayed".
	Visibility string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *PeekMessagesParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.CRN != "" {
		q.Set("crn", p.CRN)
	}
	if p.Limit != 0 {
		q.Set("limit", strconv.Itoa(int(p.Limit)))
	}
	if p.Marker != "" {
		q.Set("marker", p.Marker)
	}
	if p.Visibility != "" {
		q.Set("visibility", p.Visibility)
	}
	return q
}

// withMarker copies p with the pagination cursor replaced, leaving the
// caller's value untouched across pages.
func (p *PeekMessagesParams) withMarker(marker string) *PeekMessagesParams {
	var out PeekMessagesParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ChangeMessageVisibility — Extend or shorten a lease.
//
// A consumer that needs more time extends its lease rather than letting
// the message be redelivered while it is still working. A timeout of 0
// releases the message immediately, which is how to hand back work you
// cannot finish without waiting out the timeout.
func (c *Client) ChangeMessageVisibility(ctx context.Context, queueID string, body *ChangeVisibilityRequest, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "changeMessageVisibility",
		Method:   "POST",
		Path:     "/v1/queues/{queue_id}/messages/visibility",
		PathArgs: []string{queueID},
		Body:     body,
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// ChangeMessageVisibilityBatch — Extend or shorten up to ten leases.
func (c *Client) ChangeMessageVisibilityBatch(ctx context.Context, queueID string, body *ChangeVisibilityBatchRequest, opts ...basaltic.RequestOption) (*basaltic.Page[BatchResult], error) {
	op := &basaltic.Operation{
		ID:       "changeMessageVisibilityBatch",
		Method:   "POST",
		Path:     "/v1/queues/{queue_id}/messages/visibility-batch",
		PathArgs: []string{queueID},
		Body:     body,
	}
	var out struct {
		Items []BatchResult `json:"results"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[BatchResult]{Items: out.Items}
	return page, nil
}

// CreateQueue creates a queue.
//
// The kind is derived from the name: a name ending in `.fifo` creates a
// FIFO queue, anything else a standard one. There is no separate flag,
// so the two can never disagree.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateQueue(ctx context.Context, body *CreateQueueRequest, opts ...basaltic.RequestOption) (*Queue, error) {
	op := &basaltic.Operation{
		ID:     "createQueue",
		Method: "POST",
		Path:   "/v1/queues",
		Body:   body,
	}
	var out struct {
		Queue *Queue `json:"queue"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Queue, nil
}

// DeleteMessage deletes a leased message.
//
// The receipt handle is a query parameter rather than a path segment
// because it belongs to a lease, not to a resource: there is no stable
// URL a handle could name twice.
func (c *Client) DeleteMessage(ctx context.Context, queueID string, params *DeleteMessageParams, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteMessage",
		Method:   "DELETE",
		Path:     "/v1/queues/{queue_id}/messages",
		PathArgs: []string{queueID},
	}
	op.Query = params.query()
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteMessageBatch deletes up to ten leased messages.
func (c *Client) DeleteMessageBatch(ctx context.Context, queueID string, body *DeleteMessageBatchRequest, opts ...basaltic.RequestOption) (*basaltic.Page[BatchResult], error) {
	op := &basaltic.Operation{
		ID:       "deleteMessageBatch",
		Method:   "POST",
		Path:     "/v1/queues/{queue_id}/messages/delete-batch",
		PathArgs: []string{queueID},
		Body:     body,
	}
	var out struct {
		Items []BatchResult `json:"results"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[BatchResult]{Items: out.Items}
	return page, nil
}

// DeleteQueue deletes a queue.
//
// Removes the queue and every message on it. Refused while another queue
// names this one as its dead-letter target: deleting it would leave that
// queue's redrive policy pointing at nothing, and its poison messages
// would accumulate instead of being quarantined.
func (c *Client) DeleteQueue(ctx context.Context, queueID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteQueue",
		Method:   "DELETE",
		Path:     "/v1/queues/{queue_id}",
		PathArgs: []string{queueID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DescribeQueue describes a queue.
func (c *Client) DescribeQueue(ctx context.Context, queueID string, opts ...basaltic.RequestOption) (*Queue, error) {
	op := &basaltic.Operation{
		ID:       "describeQueue",
		Method:   "GET",
		Path:     "/v1/queues/{queue_id}",
		PathArgs: []string{queueID},
	}
	var out struct {
		Queue *Queue `json:"queue"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Queue, nil
}

// ListQueues lists queues.
//
// Returns one page. Use ListQueuesAll to walk every page.
func (c *Client) ListQueues(ctx context.Context, params *ListQueuesParams, opts ...basaltic.RequestOption) (*basaltic.Page[Queue], error) {
	op := &basaltic.Operation{
		ID:     "listQueues",
		Method: "GET",
		Path:   "/v1/queues",
	}
	op.Query = params.query()
	var out struct {
		Items []Queue `json:"queues"`
		Meta  *struct {
			Total   int    `json:"total"`
			Limit   int    `json:"limit"`
			Marker  string `json:"marker"`
			HasMore bool   `json:"has_more"`
		} `json:"meta"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[Queue]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListQueuesAll walks every page of ListQueues, yielding one item at a
// time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListQueuesAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListQueuesAll(ctx context.Context, params *ListQueuesParams, opts ...basaltic.RequestOption) iter.Seq2[Queue, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Queue], error) {
		return c.ListQueues(ctx, params.withMarker(marker), opts...)
	})
}

// PeekMessages peeks at a queue's messages.
//
// Reads without leasing: it does not increment `receive_count`, does not
// hide anything from a real consumer, and returns no receipt handles.
// Its own permission (`queue:PeekMessages`) so an operator can be given
// visibility into a stuck queue without the ability to drain it.
//
// Returns one page. Use PeekMessagesAll to walk every page.
func (c *Client) PeekMessages(ctx context.Context, queueID string, params *PeekMessagesParams, opts ...basaltic.RequestOption) (*basaltic.Page[Message], error) {
	op := &basaltic.Operation{
		ID:       "peekMessages",
		Method:   "GET",
		Path:     "/v1/queues/{queue_id}/messages",
		PathArgs: []string{queueID},
	}
	op.Query = params.query()
	var out struct {
		Items []Message `json:"messages"`
		Meta  *struct {
			Total   int    `json:"total"`
			Limit   int    `json:"limit"`
			Marker  string `json:"marker"`
			HasMore bool   `json:"has_more"`
		} `json:"meta"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[Message]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// PeekMessagesAll walks every page of PeekMessages, yielding one item at
// a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.PeekMessagesAll(ctx, queueID, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) PeekMessagesAll(ctx context.Context, queueID string, params *PeekMessagesParams, opts ...basaltic.RequestOption) iter.Seq2[Message, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Message], error) {
		return c.PeekMessages(ctx, queueID, params.withMarker(marker), opts...)
	})
}

// PurgeQueue deletes every message on a queue.
//
// Its own permission (`queue:PurgeQueue`), separate from
// `queue:UpdateQueue`: an operator who should be able to retune a
// visibility timeout should not thereby be able to throw away a backlog.
func (c *Client) PurgeQueue(ctx context.Context, queueID string, opts ...basaltic.RequestOption) (int, error) {
	op := &basaltic.Operation{
		ID:       "purgeQueue",
		Method:   "POST",
		Path:     "/v1/queues/{queue_id}/purge",
		PathArgs: []string{queueID},
	}
	var out struct {
		MessagesDeleted int `json:"messages_deleted"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return 0, err
	}
	return out.MessagesDeleted, nil
}

// ReceiveMessages — Lease messages.
//
// POST rather than GET because it mutates: a receive takes a lease,
// increments the delivery count, and hides the messages from every other
// consumer. Modelling that as a GET would invite a proxy or a retrying
// client to consume the queue by replaying what looks like a safe read.
//
// Returns an empty list, not an error, when nothing is available.
func (c *Client) ReceiveMessages(ctx context.Context, queueID string, body *ReceiveRequest, opts ...basaltic.RequestOption) (*basaltic.Page[ReceivedMessage], error) {
	op := &basaltic.Operation{
		ID:       "receiveMessages",
		Method:   "POST",
		Path:     "/v1/queues/{queue_id}/messages/receive",
		PathArgs: []string{queueID},
		Body:     body,
	}
	var out struct {
		Items []ReceivedMessage `json:"messages"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[ReceivedMessage]{Items: out.Items}
	return page, nil
}

// SendMessage sends a message.
func (c *Client) SendMessage(ctx context.Context, queueID string, body *SendMessageRequest, opts ...basaltic.RequestOption) (*SendResult, error) {
	op := &basaltic.Operation{
		ID:       "sendMessage",
		Method:   "POST",
		Path:     "/v1/queues/{queue_id}/messages",
		PathArgs: []string{queueID},
		Body:     body,
	}
	var out struct {
		Result *SendResult `json:"result"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Result, nil
}

// SendMessageBatch sends up to ten messages.
//
// Reports per-entry outcomes rather than failing wholesale: in a batch
// of ten, one oversized body should not cost the other nine. The
// response is 200 even when every entry succeeded — the status
// describes whether the request was processed, not whether every message
// landed.
func (c *Client) SendMessageBatch(ctx context.Context, queueID string, body *SendMessageBatchRequest, opts ...basaltic.RequestOption) (*basaltic.Page[BatchResult], error) {
	op := &basaltic.Operation{
		ID:       "sendMessageBatch",
		Method:   "POST",
		Path:     "/v1/queues/{queue_id}/messages/batch",
		PathArgs: []string{queueID},
		Body:     body,
	}
	var out struct {
		Items []BatchResult `json:"results"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[BatchResult]{Items: out.Items}
	return page, nil
}

// UpdateQueue updates queue attributes.
//
// The name, the kind, and `content_based_deduplication` are immutable.
// The first two are the CRN and the delivery contract; flipping the
// third would change how sends already accepted inside the live
// deduplication window are interpreted.
//
// Changing `kms_key` re-keys the queue. `kms_key: null` restores the
// platform key; omitting `kms_key` leaves it unchanged. Messages already
// on it keep the key they were sealed with, so they stay readable under
// the old key until they drain — which is the point of rotating away
// from a key you no longer trust.
func (c *Client) UpdateQueue(ctx context.Context, queueID string, body *UpdateQueueRequest, opts ...basaltic.RequestOption) (*Queue, error) {
	op := &basaltic.Operation{
		ID:       "updateQueue",
		Method:   "PATCH",
		Path:     "/v1/queues/{queue_id}",
		PathArgs: []string{queueID},
		Body:     body,
	}
	var out struct {
		Queue *Queue `json:"queue"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Queue, nil
}

// DescribeQueueByReference fetches one queue by an id, a CRN or a name.
//
// The reference is classified by its syntax alone, exactly as the
// platform does (see [basaltic.ParseReference]): an id is fetched with
// [Client.DescribeQueue]; a CRN or a name goes to [Client.ListQueues] as
// an exact filter, together with any filters already set on scope, which
// may be nil. A miss is a not-found error for the kind the string was
// read as — no other kind is tried — and more than one match is a
// [basaltic.AmbiguousReferenceError].
func (c *Client) DescribeQueueByReference(ctx context.Context, ref string, scope *ListQueuesParams, opts ...basaltic.RequestOption) (*Queue, error) {
	return basaltic.ResolveByReference(ctx, ref, "queue", "listQueues", true,
		func(ctx context.Context, refID string) (*Queue, error) {
			return c.DescribeQueue(ctx, refID, opts...)
		},
		func(ctx context.Context, refName, refCRN string) (*basaltic.Page[Queue], error) {
			var p ListQueuesParams
			if scope != nil {
				p = *scope
			}
			p.Name = refName
			p.CRN = refCRN
			p.Limit = 2
			return c.ListQueues(ctx, &p, opts...)
		})
}
