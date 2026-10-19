// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package notifications

import (
	"time"
)

type CreateTopicRequest struct {
	DeliveryPolicy *DeliveryPolicy `json:"delivery_policy,omitempty"`
	DisplayName    *string         `json:"display_name,omitempty"`

	// Name 1-256 characters of letters, digits, dot, dash, or underscore.
	// Resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	//
	// Required.
	Name string `json:"name"`
	Tags Tags   `json:"tags,omitempty"`
}

type Delivery struct {
	Attempts    int        `json:"attempts"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`

	// CRN canonical message/UUID/delivery/UUID CRN.
	CRN string `json:"crn"`

	// Endpoint recorded at fan-out.
	Endpoint      string     `json:"endpoint"`
	ID            string     `json:"id"`
	LastAttemptAt *time.Time `json:"last_attempt_at,omitempty"`

	// LastError truncated to 1024 characters.
	LastError string `json:"last_error,omitempty"`

	// LastStatusCode HTTP status of the last attempt. Absent for a transport-level
	// failure or an email delivery.
	LastStatusCode *int         `json:"last_status_code,omitempty"`
	Message        *ResourceRef `json:"message"`

	// One of: "email", "https", "http", "queue".
	Protocol string `json:"protocol"`

	// One of: "pending", "delivered", "failed".
	Status       string       `json:"status"`
	Subscription *ResourceRef `json:"subscription"`
}

// DeliveryPolicy retry behaviour. Omitted fields fall back to the topic's policy, then
// to platform defaults (6 attempts, 20s to 20m exponential backoff with
// jitter, 15s per-request timeout). A subscription's policy is layered
// on its topic's field by field, so overriding only `max_attempts` keeps
// the topic's backoff curve.
type DeliveryPolicy struct {
	// MaxAttempts total attempts including the first. 1 disables retries.
	MaxAttempts *int `json:"max_attempts,omitempty"`

	// MaxDelaySeconds cap on the exponential curve.
	MaxDelaySeconds *int `json:"max_delay_seconds,omitempty"`

	// MinDelaySeconds backoff before the second attempt.
	MinDelaySeconds *int `json:"min_delay_seconds,omitempty"`

	// TimeoutSeconds per-request timeout. Ignored for email.
	TimeoutSeconds *int `json:"timeout_seconds,omitempty"`
}

// FilterPolicy matches publication attributes, not fields inside the message body.
// Every named attribute must match at least one alternative (AND across
// names, OR within each array). Names and string comparisons are
// case-sensitive. An omitted or empty policy matches every message.
//
// Alternatives are exact strings or numbers, or single-operator objects:
// `{"prefix": "order."}`, `{"suffix": ".failed"}`, `{"anything-but":
// ["debug", "trace"]}`, `{"exists": true}`, `{"exists": false}`, or
// `{"numeric": [">=", 100, "<", 1000]}`. Prefix and suffix values must
// be nonempty. `anything-but` accepts one string or number, or a list of
// 1–50 such values.
//
// Only `exists: false` matches an absent attribute. Exclusion also
// requires presence. Numeric operators are `=`, `<`, `<=`, `>`, and
// `>=`; every operator/number pair in one numeric matcher must succeed.
// Attribute values that cannot be parsed as numbers do not match. Bare
// numeric alternatives compare string representations; use `numeric` for
// numeric equality, such as matching both "100" and "100.0".
//
// At most 20 nonblank attribute names and 1–50 alternatives per name.
// Each operator object has exactly one operator. Wildcards, regular
// expressions, body filtering, and nested operators are not supported.
// PATCH replaces the entire policy; `{}` clears it and omission
// preserves it. See [Filter policies](/notifications/filter-policies)
// for worked examples.
type FilterPolicy = map[string][]any

type Message struct {
	Attributes MessageAttributes `json:"attributes,omitempty"`
	CRN        string            `json:"crn"`

	// DeliveryCount deliveries fan-out produced. Meaningful only once fanout_state is
	// complete.
	DeliveryCount int        `json:"delivery_count"`
	FannedOutAt   *time.Time `json:"fanned_out_at,omitempty"`

	// FanoutState `pending` until the worker has expanded the message into deliveries.
	//
	// One of: "pending", "complete".
	FanoutState string `json:"fanout_state"`
	ID          string `json:"id"`

	// One of: "", "json".
	MessageStructure string    `json:"message_structure,omitempty"`
	PublishedAt      time.Time `json:"published_at"`
	Subject          string    `json:"subject,omitempty"`
}

// MessageAttributes typed attributes carried with a notification and matched against
// subscription filter policies. At most 20 attributes; values at most
// 1024 bytes. Binary attributes are not accepted — they cannot be
// filtered on.
type MessageAttributes = map[string]*MessageAttributesValue

type MessageAttributesValue struct {
	// One of: "String", "Number".
	Type  string `json:"type"`
	Value string `json:"value"`
}

type PublishBatchRequest struct {
	// Required.
	Entries []*PublishBatchRequestEntry `json:"entries"`
}

type PublishBatchRequestEntry struct {
	Attributes MessageAttributes `json:"attributes,omitempty"`

	// ID caller-supplied correlation id, unique within the batch. Echoed back
	// in the result.
	//
	// Required.
	ID string `json:"id"`

	// Message the payload. Up to 256 KiB.
	//
	// Required.
	Message string `json:"message"`

	// MessageStructure omit to deliver `message` verbatim to every protocol. Set to `json`
	// to have `message` read as an object of protocol keys (`default`,
	// `email`, `https`, `http`), letting one publish carry a different
	// payload per transport. A `default` key is required — it is what a
	// protocol with no explicit entry gets.
	//
	// One of: "", "json".
	MessageStructure *string `json:"message_structure,omitempty"`

	// Subject used as the email subject. No line breaks.
	Subject *string `json:"subject,omitempty"`
}

type PublishBatchResponse struct {
	Failed     []*PublishBatchResponseFailedItem     `json:"failed,omitempty"`
	Successful []*PublishBatchResponseSuccessfulItem `json:"successful,omitempty"`
}

type PublishBatchResponseFailedItem struct {
	Code    string `json:"code"`
	ID      string `json:"id"`
	Message string `json:"message"`
}

type PublishBatchResponseSuccessfulItem struct {
	ID        string `json:"id"`
	MessageID string `json:"message_id"`
}

type PublishRequest struct {
	Attributes MessageAttributes `json:"attributes,omitempty"`

	// Message the payload. Up to 256 KiB.
	//
	// Required.
	Message string `json:"message"`

	// MessageStructure omit to deliver `message` verbatim to every protocol. Set to `json`
	// to have `message` read as an object of protocol keys (`default`,
	// `email`, `https`, `http`), letting one publish carry a different
	// payload per transport. A `default` key is required — it is what a
	// protocol with no explicit entry gets.
	//
	// One of: "", "json".
	MessageStructure *string `json:"message_structure,omitempty"`

	// Subject used as the email subject. No line breaks.
	Subject *string `json:"subject,omitempty"`
}

type PutTopicPolicyRequest struct {
	// Required.
	Document TopicPolicy `json:"document"`
}

type RedrivePolicy struct {
	// DeadLetterQueue the bound queue. A deleted target retains its stored ID and CRN
	// without a name.
	DeadLetterQueue *ResourceRef `json:"dead_letter_queue"`
}

// RedrivePolicyInput names a queue in the subscriber's own account that receives each
// terminally failed delivery — the full original notification plus a
// `failure` block (delivery id, endpoint, attempts, last error, status
// code) — as a `delivery_failure` message. Without one, terminal
// failures remain visible in delivery records and Telemetry metrics.
//
// The caller must hold `queue:SendMessage` on the target when the policy
// is set: writes happen later with no principal present, so consent is
// established up front, exactly as for a `protocol=queue` endpoint. For
// a `queue` subscription the target must differ from the endpoint queue.
type RedrivePolicyInput struct {
	// DeadLetterQueue Queue UUID, CRN or exact name on input, scoped to the subscriber
	// account and serving region. Empty on PATCH clears the target.
	//
	// Required.
	DeadLetterQueue string `json:"dead_letter_queue"`
}

type ResourceRef struct {
	CRN string `json:"crn"`
	ID  string `json:"id"`

	// Name display label at response time — the resource's name where it has
	// one, otherwise its human label (an invoice number, an account
	// display name). Absent when the target is gone. Not a lookup key.
	Name string `json:"name,omitempty"`
}

type SigningSecretResponse struct {
	SigningSecret string       `json:"signing_secret"`
	Subscription  *ResourceRef `json:"subscription"`
}

type SubscribeRequest struct {
	DeliveryPolicy *DeliveryPolicy `json:"delivery_policy,omitempty"`

	// Endpoint an email address; an http/https URL whose scheme matches the
	// protocol and resolves to a publicly-routable address; or, for
	// `queue`, a queue UUID, CRN or exact name in the subscriber account
	// and serving region. Queue targets are stored and returned as
	// canonical CRNs.
	//
	// A `queue` subscription skips the confirmation handshake and is
	// created already confirmed. The handshake exists to prove that
	// whoever controls the endpoint consents to receive, which is
	// necessary for an address or a URL you may not own. A queue is your
	// own resource, so consent is established instead by requiring
	// `queue:SendMessage` on it at subscribe time — an authorization
	// decision rather than a demonstration of reachability.
	//
	// Required.
	Endpoint     string       `json:"endpoint"`
	FilterPolicy FilterPolicy `json:"filter_policy,omitempty"`

	// One of: "email", "https", "http", "queue".
	//
	// Required.
	Protocol string `json:"protocol"`

	// RawMessageDelivery http/https only.
	RawMessageDelivery *bool               `json:"raw_message_delivery,omitempty"`
	RedrivePolicy      *RedrivePolicyInput `json:"redrive_policy,omitempty"`
	Tags               Tags                `json:"tags,omitempty"`
}

type SubscribeResponse struct {
	// SigningSecret The HMAC key for verifying webhook deliveries. Returned only on the
	// first create of an http/https subscription — absent for email, and
	// absent when an existing subscription is returned.
	SigningSecret string        `json:"signing_secret,omitempty"`
	Subscription  *Subscription `json:"subscription"`
}

type Subscription struct {
	// ConfirmationExpiresAt when the pending confirmation lapses. Absent once confirmed.
	ConfirmationExpiresAt *time.Time `json:"confirmation_expires_at,omitempty"`
	ConfirmedAt           *time.Time `json:"confirmed_at,omitempty"`
	CreatedAt             time.Time  `json:"created_at"`

	// CRN nested under the actual topic name and owner account, including
	// cross-account subscriptions.
	CRN            string          `json:"crn"`
	DeliveryPolicy *DeliveryPolicy `json:"delivery_policy,omitempty"`

	// Endpoint Immutable. An email address, an http/https URL, or a canonical queue
	// CRN when protocol is queue.
	Endpoint     string       `json:"endpoint"`
	FilterPolicy FilterPolicy `json:"filter_policy,omitempty"`
	ID           string       `json:"id"`

	// One of: "email", "https", "http", "queue".
	Protocol string `json:"protocol"`

	// RawMessageDelivery http/https only. Deliver the published payload verbatim instead of
	// the notification envelope; the metadata moves into headers.
	RawMessageDelivery bool           `json:"raw_message_delivery"`
	RedrivePolicy      *RedrivePolicy `json:"redrive_policy,omitempty"`

	// One of: "pending_confirmation", "confirmed".
	Status    string    `json:"status"`
	Tags      Tags      `json:"tags,omitempty"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Tags = map[string]string

type Topic struct {
	CreatedAt      time.Time       `json:"created_at"`
	CRN            string          `json:"crn"`
	DeliveryPolicy *DeliveryPolicy `json:"delivery_policy,omitempty"`

	// DisplayName human-facing label; used as the email subject when a notification
	// carries none.
	DisplayName string `json:"display_name,omitempty"`
	ID          string `json:"id"`

	// Name unique within the calling account. Immutable. Resource names must
	// not start with the literal crn: prefix or be UUIDs (canonical,
	// compact, braced, or urn:uuid: forms, in either case).
	Name      string    `json:"name"`
	Tags      Tags      `json:"tags,omitempty"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TopicPolicy pkg/policy.PolicyDocument serialization, the same engine and statement
// shape as bucket policies. Declared opaque here so the spec doesn't
// track schema changes inside the policy engine.
type TopicPolicy = map[string]any

type UpdateSubscriptionRequest struct {
	DeliveryPolicy     *DeliveryPolicy `json:"delivery_policy,omitempty"`
	FilterPolicy       FilterPolicy    `json:"filter_policy,omitempty"`
	RawMessageDelivery *bool           `json:"raw_message_delivery,omitempty"`

	// RedrivePolicy omitted leaves the dead-letter queue unchanged. Present with an
	// empty `dead_letter_queue` clears it. Setting a new target re-checks
	// `queue:SendMessage` on it.
	RedrivePolicy *RedrivePolicyInput `json:"redrive_policy,omitempty"`
	Tags          Tags                `json:"tags,omitempty"`
}

type UpdateTopicRequest struct {
	DeliveryPolicy *DeliveryPolicy `json:"delivery_policy,omitempty"`
	DisplayName    *string         `json:"display_name,omitempty"`
	Tags           Tags            `json:"tags,omitempty"`
}
