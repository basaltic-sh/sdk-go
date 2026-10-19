// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package queue

import (
	"time"
)

type BatchResult struct {
	Duplicate bool `json:"duplicate,omitempty"`

	// Error machine-readable code when this entry failed.
	Error string `json:"error,omitempty"`

	// ID the entry id from the request.
	ID string `json:"id"`

	// Message human-readable reason accompanying error.
	Message   string `json:"message,omitempty"`
	MessageID string `json:"message_id,omitempty"`
}

type ChangeVisibilityBatchRequest struct {
	// Required.
	Entries []*ChangeVisibilityBatchRequestEntry `json:"entries"`
}

type ChangeVisibilityBatchRequestEntry struct {
	// Required.
	ID string `json:"id"`

	// Required.
	ReceiptHandle string `json:"receipt_handle"`

	// Required.
	VisibilityTimeoutSeconds int `json:"visibility_timeout_seconds"`
}

type ChangeVisibilityRequest struct {
	// Required.
	ReceiptHandle string `json:"receipt_handle"`

	// VisibilityTimeoutSeconds measured from now, not from the original lease. 0 releases the
	// message immediately.
	//
	// Required.
	VisibilityTimeoutSeconds int `json:"visibility_timeout_seconds"`
}

type CreateQueueRequest struct {
	// ContentBasedDeduplication FIFO queues only.
	ContentBasedDeduplication    *bool `json:"content_based_deduplication,omitempty"`
	DelaySeconds                 *int  `json:"delay_seconds,omitempty"`
	KMSDataKeyReusePeriodSeconds *int  `json:"kms_data_key_reuse_period_seconds,omitempty"`

	// KMSKey KMS key UUID, CRN or account-scoped name. Omit to use the platform
	// key.
	KMSKey                  *string `json:"kms_key,omitempty"`
	MaxMessageBytes         *int    `json:"max_message_bytes,omitempty"`
	MessageRetentionSeconds *int    `json:"message_retention_seconds,omitempty"`

	// Name letters, digits, dot, dash, or underscore, starting with a letter or
	// digit. A name ending in `.fifo` creates a FIFO queue. Resource names
	// must not start with the literal crn: prefix or be UUIDs (canonical,
	// compact, braced, or urn:uuid: forms, in either case).
	//
	// Required.
	Name                     string              `json:"name"`
	ReceiveWaitTimeSeconds   *int                `json:"receive_wait_time_seconds,omitempty"`
	RedrivePolicy            *RedrivePolicyInput `json:"redrive_policy,omitempty"`
	Tags                     Tags                `json:"tags,omitempty"`
	VisibilityTimeoutSeconds *int                `json:"visibility_timeout_seconds,omitempty"`
}

type DeleteMessageBatchRequest struct {
	// Required.
	Entries []*DeleteMessageBatchRequestEntry `json:"entries"`
}

type DeleteMessageBatchRequestEntry struct {
	// Required.
	ID string `json:"id"`

	// Required.
	ReceiptHandle string `json:"receipt_handle"`
}

// Message a message as seen by the peek endpoint. It carries no receipt handle,
// because peeking leases nothing.
type Message struct {
	Attributes MessageAttributes `json:"attributes,omitempty"`

	// Body absent when the body could not be decrypted — the rest of the row
	// is still reported.
	Body string `json:"body,omitempty"`

	// BodyBytes plaintext size, counted against the queue's max_message_bytes.
	BodyBytes  int       `json:"body_bytes,omitempty"`
	CRN        string    `json:"crn"`
	EnqueuedAt time.Time `json:"enqueued_at"`

	// ExpiresAt when retention discards this message.
	ExpiresAt       time.Time  `json:"expires_at"`
	FirstReceivedAt *time.Time `json:"first_received_at,omitempty"`
	ID              string     `json:"id"`
	MessageGroupID  string     `json:"message_group_id,omitempty"`
	QueueID         string     `json:"queue_id"`
	ReceiveCount    int        `json:"receive_count"`

	// One of: "available", "in_flight", "delayed".
	Visibility string    `json:"visibility"`
	VisibleAt  time.Time `json:"visible_at"`
}

// MessageAttributes typed attributes carried alongside the body. At most 20; values at
// most 1024 characters. `Binary` values are standard base64. The queue
// never coerces a value — it only checks the value is representable as
// the declared type.
type MessageAttributes = map[string]*MessageAttributesValue

type MessageAttributesValue struct {
	// One of: "String", "Number", "Binary".
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Queue struct {
	Attributes *QueueAttributes `json:"attributes"`

	// ContentBasedDeduplication FIFO only. Derives a deduplication id from a hash of the body when
	// the sender supplies none.
	ContentBasedDeduplication bool        `json:"content_based_deduplication,omitempty"`
	CreatedAt                 time.Time   `json:"created_at"`
	CRN                       string      `json:"crn"`
	Depth                     *QueueDepth `json:"depth,omitempty"`
	ID                        string      `json:"id"`

	// Kind derived from the name and immutable.
	//
	// One of: "standard", "fifo".
	Kind string `json:"kind"`

	// KMSKey customer-managed key wrapping the queue data key. Omitted for
	// platform encryption. If key metadata is unavailable, the reference
	// retains its stored ID and CRN without a name.
	KMSKey *ResourceRef `json:"kms_key,omitempty"`

	// Name unique within the calling account. Immutable. Resource names must
	// not start with the literal crn: prefix or be UUIDs (canonical,
	// compact, braced, or urn:uuid: forms, in either case).
	Name          string         `json:"name"`
	RedrivePolicy *RedrivePolicy `json:"redrive_policy,omitempty"`
	Tags          Tags           `json:"tags,omitempty"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

// QueueAttributes a queue's tunable behaviour. All durations are whole seconds.
type QueueAttributes struct {
	// DelaySeconds default delay applied to every send that does not override it. Not
	// supported on FIFO queues.
	DelaySeconds int `json:"delay_seconds,omitempty"`

	// KMSDataKeyReusePeriodSeconds how long an unwrapped data key is cached. Sends and receives keep
	// working for up to this long after the wrapping key is disabled;
	// lower it to tighten that window at the cost of more KMS round-trips.
	KMSDataKeyReusePeriodSeconds int `json:"kms_data_key_reuse_period_seconds,omitempty"`

	// MaxMessageBytes largest accepted message. Attributes count against this alongside
	// the body.
	MaxMessageBytes int `json:"max_message_bytes,omitempty"`

	// MessageRetentionSeconds how long an undeleted message survives before it is discarded.
	MessageRetentionSeconds int `json:"message_retention_seconds,omitempty"`

	// ReceiveWaitTimeSeconds default long-poll window for a receive that does not override it.
	ReceiveWaitTimeSeconds int `json:"receive_wait_time_seconds,omitempty"`

	// VisibilityTimeoutSeconds how long a received message stays invisible before redelivery. 0
	// redelivers immediately.
	VisibilityTimeoutSeconds int `json:"visibility_timeout_seconds,omitempty"`
}

// QueueDepth approximate message census. Approximate by construction: an exact
// count of a deep queue is a sequential scan, and the number is stale
// the moment it is returned. Counts are capped at 10000 —
// `counts_capped` reports that at least one is a floor rather than an
// exact figure.
type QueueDepth struct {
	ApproximateAvailable int  `json:"approximate_available,omitempty"`
	ApproximateDelayed   int  `json:"approximate_delayed,omitempty"`
	ApproximateInFlight  int  `json:"approximate_in_flight,omitempty"`
	CountsCapped         bool `json:"counts_capped,omitempty"`
}

type ReceiveRequest struct {
	// MaxMessages how many messages to lease. Fewer may be returned even when more are
	// available.
	MaxMessages *int `json:"max_messages,omitempty"`

	// VisibilityTimeoutSeconds overrides the queue's default for this lease only.
	VisibilityTimeoutSeconds *int `json:"visibility_timeout_seconds,omitempty"`

	// WaitTimeSeconds block up to this long for a message rather than returning empty.
	// Defaults to the queue's receive_wait_time_seconds.
	WaitTimeSeconds *int `json:"wait_time_seconds,omitempty"`
}

type ReceivedMessage struct {
	Attributes     MessageAttributes `json:"attributes,omitempty"`
	Body           string            `json:"body"`
	EnqueuedAt     time.Time         `json:"enqueued_at"`
	ID             string            `json:"id"`
	MessageGroupID string            `json:"message_group_id,omitempty"`

	// ReceiptHandle the credential for this lease. Required to delete the message or
	// extend its lease, and valid only until the lease lapses.
	ReceiptHandle string `json:"receipt_handle"`

	// ReceiveCount how many times this message has been leased, including now. Compare
	// against the redrive policy to spot a poison message.
	ReceiveCount int `json:"receive_count"`

	// VisibleAt when this lease lapses and the message is redelivered.
	VisibleAt time.Time `json:"visible_at"`
}

// RedrivePolicy quarantines a message that has been leased and abandoned too many
// times. Once `receive_count` reaches `max_receive_count` the message
// stops being delivered and is moved to the dead-letter queue, which
// must be the same kind as this one.
type RedrivePolicy struct {
	DeadLetterQueue *ResourceRef `json:"dead_letter_queue"`
	MaxReceiveCount int          `json:"max_receive_count"`
}

// RedrivePolicyInput target queue UUID, CRN or account-scoped name, resolved once and
// stored as an immutable UUID binding. Must be owned by this account, in
// this region, of the same kind, and different from the source queue.
// Invalid syntax, type or region returns 400; missing or foreign targets
// return 404. No lookup fallback is performed.
type RedrivePolicyInput struct {
	// Required.
	DeadLetterQueue string `json:"dead_letter_queue"`

	// Required.
	MaxReceiveCount int `json:"max_receive_count"`
}

type ResourceRef struct {
	CRN string `json:"crn"`
	ID  string `json:"id"`

	// Name display label at response time — the resource's name where it has
	// one, otherwise its human label (an invoice number, an account
	// display name). Absent when the target is gone. Not a lookup key.
	Name string `json:"name,omitempty"`
}

type SendMessageBatchRequest struct {
	// Required.
	Entries []*SendMessageBatchRequestEntry `json:"entries"`
}

type SendMessageBatchRequestEntry struct {
	Attributes MessageAttributes `json:"attributes,omitempty"`

	// Body the payload. Counted against the queue's max_message_bytes together
	// with the attributes.
	//
	// Required.
	Body string `json:"body"`

	// DelaySeconds overrides the queue's default for this message. Rejected on FIFO
	// queues, where it would let a later message become visible before an
	// earlier one in the same group.
	DelaySeconds *int `json:"delay_seconds,omitempty"`

	// ID your own identifier for this entry, echoed back in the result. Must
	// be unique within the request.
	//
	// Required.
	ID string `json:"id"`

	// MessageDeduplicationID FIFO only. A repeat within five minutes is suppressed and the
	// original message's id is returned. Omit on a queue with
	// `content_based_deduplication` to derive it from the body.
	MessageDeduplicationID *string `json:"message_deduplication_id,omitempty"`

	// MessageGroupID required on FIFO queues, rejected on standard ones. Messages sharing
	// a group are delivered in order, and a group's next message is
	// withheld while an earlier one is in flight.
	MessageGroupID *string `json:"message_group_id,omitempty"`
}

type SendMessageRequest struct {
	Attributes MessageAttributes `json:"attributes,omitempty"`

	// Body the payload. Counted against the queue's max_message_bytes together
	// with the attributes.
	//
	// Required.
	Body string `json:"body"`

	// DelaySeconds overrides the queue's default for this message. Rejected on FIFO
	// queues, where it would let a later message become visible before an
	// earlier one in the same group.
	DelaySeconds *int `json:"delay_seconds,omitempty"`

	// MessageDeduplicationID FIFO only. A repeat within five minutes is suppressed and the
	// original message's id is returned. Omit on a queue with
	// `content_based_deduplication` to derive it from the body.
	MessageDeduplicationID *string `json:"message_deduplication_id,omitempty"`

	// MessageGroupID required on FIFO queues, rejected on standard ones. Messages sharing
	// a group are delivered in order, and a group's next message is
	// withheld while an earlier one is in flight.
	MessageGroupID *string `json:"message_group_id,omitempty"`
}

type SendResult struct {
	// Duplicate a live deduplication entry suppressed this send; message_id names
	// the message the original produced.
	Duplicate bool   `json:"duplicate,omitempty"`
	MessageID string `json:"message_id"`
}

type Tags = map[string]string

type UpdateQueueRequest struct {
	DelaySeconds                 *int `json:"delay_seconds,omitempty"`
	KMSDataKeyReusePeriodSeconds *int `json:"kms_data_key_reuse_period_seconds,omitempty"`

	// KMSKey KMS key UUID, CRN or account-scoped name. Omit to keep the current
	// key; null restores the platform key. Existing messages keep the key
	// that sealed them.
	KMSKey                  *string `json:"kms_key,omitempty"`
	MaxMessageBytes         *int    `json:"max_message_bytes,omitempty"`
	MessageRetentionSeconds *int    `json:"message_retention_seconds,omitempty"`
	ReceiveWaitTimeSeconds  *int    `json:"receive_wait_time_seconds,omitempty"`

	// RedrivePolicy omit to keep the current policy; null removes it. An object replaces
	// the policy.
	RedrivePolicy            *UpdateQueueRequestRedrivePolicy `json:"redrive_policy,omitempty"`
	Tags                     Tags                             `json:"tags,omitempty"`
	VisibilityTimeoutSeconds *int                             `json:"visibility_timeout_seconds,omitempty"`
}

// UpdateQueueRequestRedrivePolicy omit to keep the current policy; null removes it. An object replaces
// the policy.
type UpdateQueueRequestRedrivePolicy struct {
	// DeadLetterQueue Queue UUID, CRN or account-scoped name.
	//
	// Required.
	DeadLetterQueue string `json:"dead_letter_queue"`

	// Required.
	MaxReceiveCount int `json:"max_receive_count"`
}
