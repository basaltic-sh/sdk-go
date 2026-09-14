// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package audit

import (
	"time"
)

type AuditLog struct {
	// Action the action performed (e.g., "iam.policy.create", "instance.start")
	Action string `json:"action,omitempty"`

	// ActorCRN immutable event-time actor identity. Null for historical entries
	// without a snapshot. Users use crn:iam:::user/<uuid>; service
	// accounts use crn:iam:::service-account/<name>; assumed roles use
	// crn:iam:::role/<name>. System actors use
	// crn:iam::platform:system/<service>, with service names certificate,
	// registry, secrets, queue, notifications and email.
	ActorCRN string `json:"actor_crn,omitempty"`

	// ActorEmail email of the actor (for users only)
	ActorEmail string `json:"actor_email,omitempty"`

	// ActorName name of the actor at the time of the event
	ActorName string `json:"actor_name,omitempty"`

	// CRN canonical event identity, scoped to the authenticated organization.
	CRN string `json:"crn,omitempty"`

	// Details additional action-specific details. For assumed-role actors,
	// actor_session_crn records crn:iam:::sts-session/<id> to correlate
	// the event with its AssumeRole call.
	Details map[string]any `json:"details,omitempty"`

	// ErrorCode error code for failed actions
	ErrorCode string `json:"error_code,omitempty"`

	// ErrorMessage error message for failed actions
	ErrorMessage string `json:"error_message,omitempty"`
	ID           string `json:"id,omitempty"`

	// IPAddress IP address of the request origin
	IPAddress string `json:"ip_address,omitempty"`

	// RequestID Request ID for correlation
	RequestID string `json:"request_id,omitempty"`

	// ResourceCRN immutable event-time resource identity. Null for historical entries
	// without a snapshot and enumerated events whose target cannot be
	// identified from event-time data, such as an unknown-email password
	// reset or a lookup that never resolved a row. A known target UUID is
	// retained in details.resource_id; it is never substituted for the
	// immutable name in a CRN.
	ResourceCRN string `json:"resource_crn,omitempty"`

	// ResourceName name of the resource at the time of the event
	ResourceName string `json:"resource_name,omitempty"`

	// Status outcome of the action
	//
	// One of: "success", "failure", "denied".
	Status string `json:"status,omitempty"`

	// Timestamp when the event occurred
	Timestamp time.Time `json:"timestamp,omitempty"`

	// UserAgent user agent string from the request
	UserAgent string `json:"user_agent,omitempty"`
}
