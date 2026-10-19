// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package notifications

import (
	"context"
	"iter"
	"net/url"
	"strconv"

	basaltic "github.com/basaltic-sh/sdk-go"
)

// ListDeliveriesParams are the optional filters and pagination controls for
// [Client.ListDeliveries]. A nil *ListDeliveriesParams sends none of them.
type ListDeliveriesParams struct {
	// CRN exact message/UUID/delivery/UUID identity. Foreign or mismatched
	// CRNs return an empty page; malformed CRNs are invalid.
	CRN string

	// Limit maximum items to return. A larger value is clamped to the maximum
	// rather than rejected, so page until `meta.has_more` is false.
	Limit int

	// Marker id of the last item on the previous page.
	Marker string

	// Message restrict to one published message. Resolved before the listing runs,
	// so an unknown id is a 404 rather than an empty page. Message accepts
	// UUID or CRN, never a bare name.
	Message string

	// Status one of: "pending", "delivered", "failed".
	Status string

	// Subscription restrict to one subscription by UUID or nested
	// topic/name/subscription/UUID CRN. Bare names are invalid.
	Subscription string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListDeliveriesParams) query() url.Values {
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
	if p.Message != "" {
		q.Set("message", p.Message)
	}
	if p.Status != "" {
		q.Set("status", p.Status)
	}
	if p.Subscription != "" {
		q.Set("subscription", p.Subscription)
	}
	return q
}

// withMarker copies p with the pagination cursor replaced, leaving the
// caller's value untouched across pages.
func (p *ListDeliveriesParams) withMarker(marker string) *ListDeliveriesParams {
	var out ListDeliveriesParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListMessagesParams are the optional filters and pagination controls for
// [Client.ListMessages]. A nil *ListMessagesParams sends none of them.
type ListMessagesParams struct {
	// CRN exact message CRN. Foreign identities return an empty page;
	// malformed CRNs are invalid.
	CRN string

	// FanoutState one of: "pending", "complete".
	FanoutState string

	// Limit maximum items to return. A larger value is clamped to the maximum
	// rather than rejected, so page until `meta.has_more` is false.
	Limit int

	// Marker id of the last item on the previous page.
	Marker string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListMessagesParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.CRN != "" {
		q.Set("crn", p.CRN)
	}
	if p.FanoutState != "" {
		q.Set("fanout_state", p.FanoutState)
	}
	if p.Limit != 0 {
		q.Set("limit", strconv.Itoa(int(p.Limit)))
	}
	if p.Marker != "" {
		q.Set("marker", p.Marker)
	}
	return q
}

// withMarker copies p with the pagination cursor replaced, leaving the
// caller's value untouched across pages.
func (p *ListMessagesParams) withMarker(marker string) *ListMessagesParams {
	var out ListMessagesParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListSubscriptionsParams are the optional filters and pagination controls for
// [Client.ListSubscriptions]. A nil *ListSubscriptionsParams sends none of them.
type ListSubscriptionsParams struct {
	// CRN exact nested topic/name/subscription/UUID CRN; intersects the parent
	// topic and status. Foreign or mismatched CRNs return an empty page,
	// except an authorized foreign topic parent may identify a
	// caller-owned subscription. No name filter exists.
	CRN string

	// Limit maximum items to return. A larger value is clamped to the maximum
	// rather than rejected, so page until `meta.has_more` is false.
	Limit int

	// Marker id of the last item on the previous page.
	Marker string

	// Status one of: "pending_confirmation", "confirmed".
	Status string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListSubscriptionsParams) query() url.Values {
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
	if p.Status != "" {
		q.Set("status", p.Status)
	}
	return q
}

// withMarker copies p with the pagination cursor replaced, leaving the
// caller's value untouched across pages.
func (p *ListSubscriptionsParams) withMarker(marker string) *ListSubscriptionsParams {
	var out ListSubscriptionsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListTopicsParams are the optional filters and pagination controls for
// [Client.ListTopics]. A nil *ListTopicsParams sends none of them.
type ListTopicsParams struct {
	// CRN exact notifications/topic CRN; foreign or mismatched CRNs return an
	// empty page.
	CRN string

	// Limit maximum items to return. A larger value is clamped to the maximum
	// rather than rejected, so page until `meta.has_more` is false.
	Limit int

	// Marker id of the last item on the previous page.
	Marker string

	// Name exact immutable topic name in the caller account; intersects crn.
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListTopicsParams) query() url.Values {
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
func (p *ListTopicsParams) withMarker(marker string) *ListTopicsParams {
	var out ListTopicsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// CreateTopic creates a topic.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateTopic(ctx context.Context, body *CreateTopicRequest, opts ...basaltic.RequestOption) (*Topic, error) {
	op := &basaltic.Operation{
		ID:     "createTopic",
		Method: "POST",
		Path:   "/v1/topics",
		Body:   body,
	}
	var out struct {
		Topic *Topic `json:"topic"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Topic, nil
}

// DeleteTopic deletes a topic.
//
// Removes the topic and everything under it — every subscription,
// published message, and delivery record. This is how delivery is
// stopped for good, so nothing is left orphaned behind it.
func (c *Client) DeleteTopic(ctx context.Context, topicID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteTopic",
		Method:   "DELETE",
		Path:     "/v1/topics/{topic_id}",
		PathArgs: []string{topicID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteTopicPolicy deletes the topic's resource policy.
//
// Removes the policy; the topic reverts to owner-account-only access.
// Existing foreign subscriptions stop receiving new messages only via
// unsubscribe — deleting the policy gates new publishes and
// subscribes, it does not tear down what it previously admitted.
func (c *Client) DeleteTopicPolicy(ctx context.Context, topicID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteTopicPolicy",
		Method:   "DELETE",
		Path:     "/v1/topics/{topic_id}/policy",
		PathArgs: []string{topicID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DescribeMessage describes a published notification.
//
// `fanout_state` is `pending` until the worker has expanded the message.
// Once it is `complete`, `delivery_count` distinguishes "delivered to N
// subscriptions" from "no subscription matched" (complete with count 0).
func (c *Client) DescribeMessage(ctx context.Context, topicID string, messageID string, opts ...basaltic.RequestOption) (*Message, error) {
	op := &basaltic.Operation{
		ID:       "describeMessage",
		Method:   "GET",
		Path:     "/v1/topics/{topic_id}/messages/{message_id}",
		PathArgs: []string{topicID, messageID},
	}
	var out struct {
		Message *Message `json:"message"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Message, nil
}

// DescribeSubscription describes a subscription.
func (c *Client) DescribeSubscription(ctx context.Context, topicID string, subscriptionID string, opts ...basaltic.RequestOption) (*Subscription, error) {
	op := &basaltic.Operation{
		ID:       "describeSubscription",
		Method:   "GET",
		Path:     "/v1/topics/{topic_id}/subscriptions/{subscription_id}",
		PathArgs: []string{topicID, subscriptionID},
	}
	var out struct {
		Subscription *Subscription `json:"subscription"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Subscription, nil
}

// DescribeTopic describes a topic.
func (c *Client) DescribeTopic(ctx context.Context, topicID string, opts ...basaltic.RequestOption) (*Topic, error) {
	op := &basaltic.Operation{
		ID:       "describeTopic",
		Method:   "GET",
		Path:     "/v1/topics/{topic_id}",
		PathArgs: []string{topicID},
	}
	var out struct {
		Topic *Topic `json:"topic"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Topic, nil
}

// GetSubscriptionSigningSecret reads the webhook signing secret.
//
// Behind its own IAM action, separate from DescribeSubscription: holding
// this secret is what lets a party forge a delivery, so a policy can
// grant subscription visibility without granting it.
func (c *Client) GetSubscriptionSigningSecret(ctx context.Context, topicID string, subscriptionID string, opts ...basaltic.RequestOption) (*SigningSecretResponse, error) {
	op := &basaltic.Operation{
		ID:       "getSubscriptionSigningSecret",
		Method:   "GET",
		Path:     "/v1/topics/{topic_id}/subscriptions/{subscription_id}/signing-secret",
		PathArgs: []string{topicID, subscriptionID},
	}
	var out SigningSecretResponse
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetTopicPolicy reads the topic's resource policy.
func (c *Client) GetTopicPolicy(ctx context.Context, topicID string, opts ...basaltic.RequestOption) (TopicPolicy, error) {
	op := &basaltic.Operation{
		ID:       "getTopicPolicy",
		Method:   "GET",
		Path:     "/v1/topics/{topic_id}/policy",
		PathArgs: []string{topicID},
	}
	var out struct {
		Document TopicPolicy `json:"document"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Document, nil
}

// ListDeliveries lists delivery records.
//
// The "why didn't my webhook fire" endpoint: each record carries the
// attempt count, the last HTTP status, and the last error.
//
// Returns one page. Use ListDeliveriesAll to walk every page.
func (c *Client) ListDeliveries(ctx context.Context, topicID string, params *ListDeliveriesParams, opts ...basaltic.RequestOption) (*basaltic.Page[Delivery], error) {
	op := &basaltic.Operation{
		ID:       "listDeliveries",
		Method:   "GET",
		Path:     "/v1/topics/{topic_id}/deliveries",
		PathArgs: []string{topicID},
	}
	op.Query = params.query()
	var out struct {
		Items []Delivery `json:"deliveries"`
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
	page := &basaltic.Page[Delivery]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListDeliveriesAll walks every page of ListDeliveries, yielding one
// item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListDeliveriesAll(ctx, topicID, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListDeliveriesAll(ctx context.Context, topicID string, params *ListDeliveriesParams, opts ...basaltic.RequestOption) iter.Seq2[Delivery, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Delivery], error) {
		return c.ListDeliveries(ctx, topicID, params.withMarker(marker), opts...)
	})
}

// ListMessages lists messages.
//
// Lists published message metadata for this topic, newest IDs first. Use
// the marker to continue through older messages. Message bodies are not
// included. The optional filters intersect with this topic.
//
// Returns one page. Use ListMessagesAll to walk every page.
func (c *Client) ListMessages(ctx context.Context, topicID string, params *ListMessagesParams, opts ...basaltic.RequestOption) (*basaltic.Page[Message], error) {
	op := &basaltic.Operation{
		ID:       "listMessages",
		Method:   "GET",
		Path:     "/v1/topics/{topic_id}/messages",
		PathArgs: []string{topicID},
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

// ListMessagesAll walks every page of ListMessages, yielding one item at
// a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListMessagesAll(ctx, topicID, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListMessagesAll(ctx context.Context, topicID string, params *ListMessagesParams, opts ...basaltic.RequestOption) iter.Seq2[Message, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Message], error) {
		return c.ListMessages(ctx, topicID, params.withMarker(marker), opts...)
	})
}

// ListSubscriptions lists subscriptions.
//
// Returns one page. Use ListSubscriptionsAll to walk every page.
func (c *Client) ListSubscriptions(ctx context.Context, topicID string, params *ListSubscriptionsParams, opts ...basaltic.RequestOption) (*basaltic.Page[Subscription], error) {
	op := &basaltic.Operation{
		ID:       "listSubscriptions",
		Method:   "GET",
		Path:     "/v1/topics/{topic_id}/subscriptions",
		PathArgs: []string{topicID},
	}
	op.Query = params.query()
	var out struct {
		Items []Subscription `json:"subscriptions"`
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
	page := &basaltic.Page[Subscription]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListSubscriptionsAll walks every page of ListSubscriptions, yielding
// one item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListSubscriptionsAll(ctx, topicID, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListSubscriptionsAll(ctx context.Context, topicID string, params *ListSubscriptionsParams, opts ...basaltic.RequestOption) iter.Seq2[Subscription, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Subscription], error) {
		return c.ListSubscriptions(ctx, topicID, params.withMarker(marker), opts...)
	})
}

// ListTopics lists topics.
//
// Returns one page. Use ListTopicsAll to walk every page.
func (c *Client) ListTopics(ctx context.Context, params *ListTopicsParams, opts ...basaltic.RequestOption) (*basaltic.Page[Topic], error) {
	op := &basaltic.Operation{
		ID:     "listTopics",
		Method: "GET",
		Path:   "/v1/topics",
	}
	op.Query = params.query()
	var out struct {
		Items []Topic `json:"topics"`
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
	page := &basaltic.Page[Topic]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListTopicsAll walks every page of ListTopics, yielding one item at a
// time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListTopicsAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListTopicsAll(ctx context.Context, params *ListTopicsParams, opts ...basaltic.RequestOption) iter.Seq2[Topic, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Topic], error) {
		return c.ListTopics(ctx, params.withMarker(marker), opts...)
	})
}

// Publish publishes a notification.
//
// Returns 202 once the notification is durable and queued. Fan-out into
// per-subscription deliveries has NOT run at that point —
// `delivery_count` on the message is 0 until it does.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) Publish(ctx context.Context, topicID string, body *PublishRequest, opts ...basaltic.RequestOption) (string, error) {
	op := &basaltic.Operation{
		ID:       "publish",
		Method:   "POST",
		Path:     "/v1/topics/{topic_id}/publish",
		PathArgs: []string{topicID},
		Body:     body,
	}
	var out struct {
		MessageID string `json:"message_id"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return "", err
	}
	return out.MessageID, nil
}

// PublishBatch publishes up to 10 notifications.
//
// Not atomic. Entries are validated and queued independently, and the
// response reports each outcome against the caller's entry id. A 202
// does NOT mean every entry was accepted — read `failed`.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) PublishBatch(ctx context.Context, topicID string, body *PublishBatchRequest, opts ...basaltic.RequestOption) (*PublishBatchResponse, error) {
	op := &basaltic.Operation{
		ID:       "publishBatch",
		Method:   "POST",
		Path:     "/v1/topics/{topic_id}/publish-batch",
		PathArgs: []string{topicID},
		Body:     body,
	}
	var out PublishBatchResponse
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// PutTopicPolicy sets the topic's resource policy.
//
// Stores an access policy on the topic — the cross-account door. A
// principal outside the owning account may publish to or subscribe on
// this topic only when BOTH its own identity policy allows the action on
// this topic's CRN AND this policy grants it (bilateral consent,
// matching bucket policies).
//
// Statements may grant only `notifications:Publish` and
// `notifications:Subscribe`. Principals: `*`, an org CRN
// (`crn:iam:::org/<org-id>`), an account CRN
// (`crn:iam:::account/<account-id>`), or an exact principal id.
// Statement resources default to this topic's CRN; `*`, the exact CRN,
// or this topic's CRN followed by `/*` are accepted. The subtree matches
// subscriptions only; list both exact and subtree resources to cover the
// topic and its descendants. Owner-only.
func (c *Client) PutTopicPolicy(ctx context.Context, topicID string, body *PutTopicPolicyRequest, opts ...basaltic.RequestOption) (TopicPolicy, error) {
	op := &basaltic.Operation{
		ID:       "putTopicPolicy",
		Method:   "PUT",
		Path:     "/v1/topics/{topic_id}/policy",
		PathArgs: []string{topicID},
		Body:     body,
	}
	var out struct {
		Document TopicPolicy `json:"document"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Document, nil
}

// ResendConfirmation — Re-send the confirmation handshake.
//
// Mints a NEW confirmation token, invalidating the previous one, and
// re-sends it. Use this when the original never arrived: a plain
// re-Subscribe is idempotent and sends nothing, and the original token
// cannot be replayed because only its hash is stored.
func (c *Client) ResendConfirmation(ctx context.Context, topicID string, subscriptionID string, opts ...basaltic.RequestOption) (*Subscription, error) {
	op := &basaltic.Operation{
		ID:       "resendConfirmation",
		Method:   "POST",
		Path:     "/v1/topics/{topic_id}/subscriptions/{subscription_id}/resend-confirmation",
		PathArgs: []string{topicID, subscriptionID},
	}
	var out struct {
		Subscription *Subscription `json:"subscription"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Subscription, nil
}

// RotateSubscriptionSigningSecret rotates the webhook signing secret.
//
// The swap is immediate and not dual-keyed. Deliveries already in flight
// are signed with whichever key was current when they were built, so a
// strict verifier will reject a small number of messages around the
// rotation.
func (c *Client) RotateSubscriptionSigningSecret(ctx context.Context, topicID string, subscriptionID string, opts ...basaltic.RequestOption) (*SigningSecretResponse, error) {
	op := &basaltic.Operation{
		ID:       "rotateSubscriptionSigningSecret",
		Method:   "POST",
		Path:     "/v1/topics/{topic_id}/subscriptions/{subscription_id}/signing-secret/rotate",
		PathArgs: []string{topicID, subscriptionID},
	}
	var out SigningSecretResponse
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Subscribe subscribes an endpoint to a topic.
//
// The parent topic is selected by `topic_id` in the path. Foreign topics
// require bilateral Subscribe authorization. Email and webhook
// subscriptions start in `pending_confirmation` and receive a
// confirmation token. Nothing is delivered until the endpoint owner
// confirms. Queue subscriptions are confirmed immediately after target
// authorization.
//
// Idempotent per `(topic, protocol, endpoint)`: subscribing an address
// that already has a subscription returns the existing one, WITHOUT a
// `signing_secret` — re-minting one would invalidate the signature the
// subscriber is already verifying against.
//
// `signing_secret` is returned exactly once, on the first create of an
// http/https subscription. Store it, or read it back later from the
// signing-secret endpoint.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) Subscribe(ctx context.Context, topicID string, body *SubscribeRequest, opts ...basaltic.RequestOption) (*SubscribeResponse, error) {
	op := &basaltic.Operation{
		ID:       "subscribe",
		Method:   "POST",
		Path:     "/v1/topics/{topic_id}/subscriptions",
		PathArgs: []string{topicID},
		Body:     body,
	}
	var out SubscribeResponse
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// Unsubscribe unsubscribes.
func (c *Client) Unsubscribe(ctx context.Context, topicID string, subscriptionID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "unsubscribe",
		Method:   "DELETE",
		Path:     "/v1/topics/{topic_id}/subscriptions/{subscription_id}",
		PathArgs: []string{topicID, subscriptionID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// UpdateSubscription updates mutable subscription settings.
//
// Protocol and endpoint are immutable — changing either would move
// deliveries to an address that never completed the confirmation
// handshake.
func (c *Client) UpdateSubscription(ctx context.Context, topicID string, subscriptionID string, body *UpdateSubscriptionRequest, opts ...basaltic.RequestOption) (*Subscription, error) {
	op := &basaltic.Operation{
		ID:       "updateSubscription",
		Method:   "PATCH",
		Path:     "/v1/topics/{topic_id}/subscriptions/{subscription_id}",
		PathArgs: []string{topicID, subscriptionID},
		Body:     body,
	}
	var out struct {
		Subscription *Subscription `json:"subscription"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Subscription, nil
}

// UpdateTopic updates mutable topic metadata.
//
// The topic name is immutable: it is the resource id in the topic's CRN,
// so renaming would silently re-target every IAM policy that references
// it.
func (c *Client) UpdateTopic(ctx context.Context, topicID string, body *UpdateTopicRequest, opts ...basaltic.RequestOption) (*Topic, error) {
	op := &basaltic.Operation{
		ID:       "updateTopic",
		Method:   "PATCH",
		Path:     "/v1/topics/{topic_id}",
		PathArgs: []string{topicID},
		Body:     body,
	}
	var out struct {
		Topic *Topic `json:"topic"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Topic, nil
}

// DescribeSubscriptionByReference fetches one subscription by an id or a
// CRN.
//
// The reference is classified by its syntax alone, exactly as the
// platform does (see [basaltic.ParseReference]): an id is fetched with
// [Client.DescribeSubscription]; a CRN or a name goes to
// [Client.ListSubscriptions] as an exact filter, together with any
// filters already set on scope, which may be nil. A miss is a not-found
// error for the kind the string was read as — no other kind is tried
// — and more than one match is a [basaltic.AmbiguousReferenceError].
//
// This resource has no name; a name reference is refused before any request.
func (c *Client) DescribeSubscriptionByReference(ctx context.Context, topicID string, ref string, scope *ListSubscriptionsParams, opts ...basaltic.RequestOption) (*Subscription, error) {
	return basaltic.ResolveByReference(ctx, ref, "subscription", "listSubscriptions", false,
		func(ctx context.Context, refID string) (*Subscription, error) {
			return c.DescribeSubscription(ctx, topicID, refID, opts...)
		},
		func(ctx context.Context, refName, refCRN string) (*basaltic.Page[Subscription], error) {
			var p ListSubscriptionsParams
			if scope != nil {
				p = *scope
			}
			_ = refName
			p.CRN = refCRN
			p.Limit = 2
			return c.ListSubscriptions(ctx, topicID, &p, opts...)
		})
}

// DescribeTopicByReference fetches one topic by an id, a CRN or a name.
//
// The reference is classified by its syntax alone, exactly as the
// platform does (see [basaltic.ParseReference]): an id is fetched with
// [Client.DescribeTopic]; a CRN or a name goes to [Client.ListTopics] as
// an exact filter, together with any filters already set on scope, which
// may be nil. A miss is a not-found error for the kind the string was
// read as — no other kind is tried — and more than one match is a
// [basaltic.AmbiguousReferenceError].
func (c *Client) DescribeTopicByReference(ctx context.Context, ref string, scope *ListTopicsParams, opts ...basaltic.RequestOption) (*Topic, error) {
	return basaltic.ResolveByReference(ctx, ref, "topic", "listTopics", true,
		func(ctx context.Context, refID string) (*Topic, error) {
			return c.DescribeTopic(ctx, refID, opts...)
		},
		func(ctx context.Context, refName, refCRN string) (*basaltic.Page[Topic], error) {
			var p ListTopicsParams
			if scope != nil {
				p = *scope
			}
			p.Name = refName
			p.CRN = refCRN
			p.Limit = 2
			return c.ListTopics(ctx, &p, opts...)
		})
}
