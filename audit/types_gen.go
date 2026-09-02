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

	// ActorEmail email of the actor (for users only)
	ActorEmail string `json:"actor_email,omitempty"`

	// ActorID ID of the user or service account that performed the action
	ActorID string `json:"actor_id,omitempty"`

	// ActorName name of the actor at the time of the event
	ActorName string `json:"actor_name,omitempty"`

	// ActorType type of actor
	//
	// One of: "user", "service_account", "system".
	ActorType string `json:"actor_type,omitempty"`

	// Details additional action-specific details
	Details map[string]any `json:"details,omitempty"`

	// ErrorCode error code for failed actions
	ErrorCode string `json:"error_code,omitempty"`

	// ErrorMessage error message for failed actions
	ErrorMessage string `json:"error_message,omitempty"`
	ID           string `json:"id,omitempty"`

	// IPAddress IP address of the request origin
	IPAddress string `json:"ip_address,omitempty"`

	// OrganizationID organization context of the action
	OrganizationID string `json:"organization_id,omitempty"`

	// RequestID Request ID for correlation
	RequestID string `json:"request_id,omitempty"`

	// ResourceID ID of the resource affected
	ResourceID string `json:"resource_id,omitempty"`

	// ResourceName name of the resource at the time of the event
	ResourceName string `json:"resource_name,omitempty"`

	// ResourceType type of resource affected
	ResourceType string `json:"resource_type,omitempty"`

	// Status outcome of the action
	//
	// One of: "success", "failure", "denied".
	Status string `json:"status,omitempty"`

	// Timestamp when the event occurred
	Timestamp time.Time `json:"timestamp,omitempty"`

	// UserAgent user agent string from the request
	UserAgent string `json:"user_agent,omitempty"`
}
