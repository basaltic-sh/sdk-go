// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package audit

import (
	"time"
)

type AuditAuthenticationStep struct {
	// Email user email captured when the request was authenticated, when
	// available.
	Email string `json:"email,omitempty"`

	// GrantType how this session was issued, such as assume_role,
	// client_credentials, imds or web_identity.
	GrantType string `json:"grant_type,omitempty"`

	// IssuedAt when this session was created.
	IssuedAt time.Time `json:"issued_at,omitempty"`

	// Issuer verified issuer of the originating web identity token, when
	// applicable.
	Issuer string `json:"issuer,omitempty"`

	// Name display name captured when the request was authenticated, when
	// available.
	Name string `json:"name,omitempty"`

	// PrincipalCRN recorded principal identity; for web identity this identifies the
	// verified provider.
	PrincipalCRN string `json:"principal_crn,omitempty"`

	// PrincipalID stable principal UUID. For a federated workload this is derived from
	// its verified issuer and subject.
	PrincipalID string `json:"principal_id"`

	// PrincipalType identity kind, such as user, service_account, assumed_role, instance
	// or web_identity.
	PrincipalType string `json:"principal_type"`

	// SessionCRN Account-qualified STS session identity, when applicable.
	SessionCRN string `json:"session_crn,omitempty"`
	SessionID  string `json:"session_id,omitempty"`

	// SessionName caller-supplied session label, not proof of identity.
	SessionName string `json:"session_name,omitempty"`

	// Subject verified subject of the originating web identity token, independent
	// of the session label.
	Subject string `json:"subject,omitempty"`
}

type AuditLog struct {
	// Action the action performed (e.g., "iam.policy.create", "instance.start")
	Action string `json:"action,omitempty"`

	// ActorCRN immutable event-time actor identity. Null for historical entries
	// without a snapshot. Authenticated user events use
	// crn:workspace:::user/<username>, with the user UUID as a fallback
	// when a username is unavailable; service accounts use
	// crn:iam::<account-handle>:service-account/<name>; assumed roles use
	// crn:iam::<account-handle>:role/<name>, with the session identity in
	// details.actor_session_crn and verified ancestry in
	// details.auth_chain. Federated assumptions identify the verified
	// provider in actor_crn and the verified subject in actor_name and
	// details.auth_chain. Historical entries retain their original CRNs,
	// including legacy IAM organization identities. System actors use
	// crn:iam::platform:system/<service>, with service names certificate,
	// registry, secrets, queue, notifications and email.
	ActorCRN *string `json:"actor_crn,omitempty"`

	// ActorEmail email of the actor (for users only)
	ActorEmail *string `json:"actor_email,omitempty"`

	// ActorName name of the actor at the time of the event
	ActorName *string `json:"actor_name,omitempty"`

	// CRN canonical event identity, scoped to the authenticated organization.
	CRN string `json:"crn,omitempty"`

	// Details additional action-specific details. For assumed-role actors,
	// actor_session_crn records crn:iam::<account-handle>:sts-session/<id>
	// to correlate the event with its AssumeRole call. auth_chain is an
	// immutable snapshot of verified identities, ordered from the original
	// user, service account or workload through every assumed role to the
	// effective actor. It contains identifiers and session metadata, never
	// credentials. Historical events without recorded ancestry omit
	// auth_chain; reads do not reconstruct it from current sessions or
	// identities. Authentication grants include session_id, session_crn,
	// grant_type and expires_at. A federated origin includes its verified
	// issuer and subject independently of the caller-supplied session
	// name. Initiated background work retains the original audit
	// attribution and request correlation.
	Details *AuditLogDetails `json:"details,omitempty"`

	// ErrorCode error code for failed actions
	ErrorCode *string `json:"error_code,omitempty"`

	// ErrorMessage error message for failed actions
	ErrorMessage *string `json:"error_message,omitempty"`
	ID           string  `json:"id,omitempty"`

	// IPAddress IP address of the request origin
	IPAddress *string `json:"ip_address,omitempty"`

	// RequestID Request ID for correlation
	RequestID *string `json:"request_id,omitempty"`

	// ResourceCRN immutable event-time resource identity. Null for historical entries
	// without a snapshot and enumerated events whose target cannot be
	// identified from event-time data, such as an unknown-email password
	// reset or a lookup that never resolved a row. A known target UUID is
	// retained in details.resource_id; it is never substituted for the
	// immutable name in a CRN. AssumeRole and AssumeRoleWithWebIdentity
	// target the role being assumed. GetToken targets the authenticated
	// service account. These authentication events record the resulting
	// session in details.session_crn.
	ResourceCRN *string `json:"resource_crn,omitempty"`

	// ResourceName name of the resource at the time of the event
	ResourceName *string `json:"resource_name,omitempty"`

	// Status outcome of the action
	//
	// One of: "success", "failure", "denied".
	Status string `json:"status,omitempty"`

	// Timestamp when the event occurred
	Timestamp time.Time `json:"timestamp,omitempty"`

	// UserAgent user agent string from the request
	UserAgent *string `json:"user_agent,omitempty"`
}

// AuditLogDetails additional action-specific details. For assumed-role actors,
// actor_session_crn records crn:iam::<account-handle>:sts-session/<id>
// to correlate the event with its AssumeRole call. auth_chain is an
// immutable snapshot of verified identities, ordered from the original
// user, service account or workload through every assumed role to the
// effective actor. It contains identifiers and session metadata, never
// credentials. Historical events without recorded ancestry omit
// auth_chain; reads do not reconstruct it from current sessions or
// identities. Authentication grants include session_id, session_crn,
// grant_type and expires_at. A federated origin includes its verified
// issuer and subject independently of the caller-supplied session name.
// Initiated background work retains the original audit attribution and
// request correlation.
type AuditLogDetails struct {
	AuthChain []*AuditAuthenticationStep `json:"auth_chain,omitempty"`
}
