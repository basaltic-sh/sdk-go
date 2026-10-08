// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package workspace

import (
	"time"
)

type Account struct {
	// BootstrapRoleCRN returned on creation. Assume this role to administer the new
	// account; source AssumeRole permission is still required.
	BootstrapRoleCRN string `json:"bootstrap_role_crn,omitempty"`

	// BootstrapRoleID returned on creation. Immutable ID of the AccountAdministrator role
	// trusted only to the account creator.
	BootstrapRoleID string    `json:"bootstrap_role_id,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`

	// CRN Canonical Workspace resource identity.
	CRN         string `json:"crn,omitempty"`
	Description string `json:"description,omitempty"`

	// Handle globally-unique, immutable handle (URL-safe identifier). Sent as
	// X-Account-Id on every request that needs account context and
	// embedded in CRNs.
	Handle string `json:"handle,omitempty"`

	// ID Internal UUID. Used for joins; the handle is the public identifier.
	ID             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	OrganizationID string `json:"organization_id,omitempty"`

	// One of: "active", "suspended", "deleted".
	Status    string    `json:"status,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type AccountPrincipalReference struct {
	AccountHandle string `json:"account_handle,omitempty"`
	AccountID     string `json:"account_id,omitempty"`
	ID            string `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
}

type AccountRole struct {
	AccountHandle string `json:"account_handle,omitempty"`
	AccountID     string `json:"account_id,omitempty"`
	AccountName   string `json:"account_name,omitempty"`
	RoleCRN       string `json:"role_crn,omitempty"`
	RoleID        string `json:"role_id,omitempty"`
	RoleName      string `json:"role_name,omitempty"`
}

type AccountRoleAssignment struct {
	AccountID string    `json:"account_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`

	// CRN organization-qualified identity of this assignment in its target
	// account.
	CRN         string `json:"crn,omitempty"`
	ID          string `json:"id,omitempty"`
	PrincipalID string `json:"principal_id,omitempty"`

	// One of: "user", "group".
	PrincipalType string `json:"principal_type,omitempty"`
	RoleID        string `json:"role_id,omitempty"`
	RoleName      string `json:"role_name,omitempty"`
}

type AccountRoleAssignmentCreateRequest struct {
	// PrincipalID Immutable UUID of a user or users-only group in this organization.
	//
	// Required.
	PrincipalID string `json:"principal_id"`

	// One of: "user", "group".
	//
	// Required.
	PrincipalType string `json:"principal_type"`

	// RoleID Immutable UUID of a role owned by the target account.
	//
	// Required.
	RoleID string `json:"role_id"`
}

type CreateAccountRequest struct {
	Description *string `json:"description,omitempty"`

	// Required.
	Handle string `json:"handle"`

	// Required.
	Name string `json:"name"`
}

type Group struct {
	CreatedAt time.Time `json:"created_at,omitempty"`

	// CRN Cloud Resource Name
	CRN         string `json:"crn,omitempty"`
	Description string `json:"description,omitempty"`
	ID          string `json:"id,omitempty"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	Name      string    `json:"name,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type GroupCreateRequest struct {
	Description *string `json:"description,omitempty"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	//
	// Required.
	Name string `json:"name"`
}

// GroupReference Group UUID, immutable name in the authenticated organization, or a
// Workspace CRN in the authenticated organization. Groups contain users
// only.
type GroupReference = string

type GroupSummary struct {
	ID string `json:"id,omitempty"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	Name string `json:"name,omitempty"`
}

// GroupUpdateRequest the resource name is immutable.
type GroupUpdateRequest struct {
	Description *string `json:"description,omitempty"`
}

// GroupUser a user in a group
type GroupUser struct {
	AddedAt time.Time `json:"added_at,omitempty"`

	// CRN of the user
	CRN   string `json:"crn,omitempty"`
	Email string `json:"email,omitempty"`
	ID    string `json:"id,omitempty"`

	// Name display name of the user
	Name string `json:"name,omitempty"`
}

type InlinePolicy struct {
	CreatedAt time.Time `json:"created_at,omitempty"`

	// CRN canonical principal-scoped inline policy identity; named principals
	// use their immutable name, users use UUID.
	CRN      string          `json:"crn"`
	Document *PolicyDocument `json:"document,omitempty"`
	ID       string          `json:"id,omitempty"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	Name        string `json:"name,omitempty"`
	PrincipalID string `json:"principal_id,omitempty"`

	// One of: "user", "group".
	PrincipalType string    `json:"principal_type,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}

type Invitation struct {
	CreatedAt time.Time `json:"created_at,omitempty"`
	CRN       string    `json:"crn,omitempty"`

	// Email address of the invited user
	Email     string    `json:"email,omitempty"`
	ExpiresAt time.Time `json:"expires_at,omitempty"`

	// Groups the user will be added to upon accepting
	Groups    []*GroupSummary      `json:"groups,omitempty"`
	ID        string               `json:"id,omitempty"`
	InvitedBy *InvitationInvitedBy `json:"invited_by,omitempty"`

	// One of: "pending", "accepted", "expired", "cancelled".
	Status string `json:"status,omitempty"`
}

type InvitationInvitedBy struct {
	// AccountID owning account for a service-account or assumed-role inviter;
	// omitted for a human inviter.
	AccountID string `json:"account_id,omitempty"`

	// CRN Canonical Workspace user CRN or account IAM service-account/session
	// CRN captured when invited.
	CRN string `json:"crn,omitempty"`

	// Email human inviter email; omitted for machine identities.
	Email string `json:"email,omitempty"`
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`

	// Type actual actor type. Assumed-role invitations record the session UUID
	// in id.
	//
	// One of: "user", "service_account", "assumed_role".
	Type string `json:"type,omitempty"`
}

// LinuxIdentity stable platform-managed identity. Primary GID equals UID. Removing and
// re-adding a membership allocates a new identity; retired IDs are never
// reused.
type LinuxIdentity struct {
	GID int32 `json:"gid"`

	// HomeDirectory new identities use /home/<username>. Existing identities retain
	// their original home path and numeric file ownership.
	HomeDirectory string `json:"home_directory"`
	UID           int32  `json:"uid"`
	Username      string `json:"username"`
}

type Organization struct {
	CreatedAt time.Time `json:"created_at,omitempty"`

	// CRN Canonical Workspace resource identity.
	CRN         string `json:"crn,omitempty"`
	Description string `json:"description,omitempty"`
	ID          string `json:"id,omitempty"`

	// Language for organization billing and operational emails,
	// independent of each user's console preference.
	//
	// One of: "en", "pt-BR", "es".
	Language string `json:"language,omitempty"`
	Name     string `json:"name,omitempty"`

	// OwnerID ID of the organization owner
	OwnerID string `json:"owner_id,omitempty"`

	// Status lifecycle state. A newly created organization is `pending` until its
	// owner has verified a phone number and attached a payment method;
	// until then every resource API refuses it with
	// `ORGANIZATION_ONBOARDING_REQUIRED`. `suspended` is a billing or
	// administrative hold, and `terminated` is irreversible.
	//
	// One of: "pending", "active", "suspended", "terminated".
	Status string `json:"status,omitempty"`

	// SuspensionReason why the organization is suspended; absent unless it is. The two are
	// the same `status` but not the same situation — a `billing` hold is
	// one the customer can clear by settling their account, and the
	// platform still grants organization context for it so they can reach
	// billing to do so. A `manual` hold is an operator decision and grants
	// nothing.
	//
	// One of: "billing", "manual".
	SuspensionReason string `json:"suspension_reason,omitempty"`

	// TimeZone IANA timezone for formatting organization emails. Does not change
	// billing periods or resource schedules.
	TimeZone  string    `json:"time_zone,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type OrganizationPolicyAttachRequest struct {
	// PolicyID Immutable UUID of the organization policy to attach.
	//
	// Required.
	PolicyID string `json:"policy_id"`
}

type OrganizationUpdateRequest struct {
	// CaptchaToken google reCAPTCHA token for bot protection
	//
	// Required.
	CaptchaToken string  `json:"captcha_token"`
	Description  *string `json:"description,omitempty"`

	// Language for organization billing and operational emails,
	// independent of each user's console preference.
	//
	// One of: "en", "pt-BR", "es".
	Language *string `json:"language,omitempty"`
	Name     *string `json:"name,omitempty"`

	// TimeZone IANA timezone for formatting organization emails. Does not change
	// billing periods or resource schedules.
	TimeZone *string `json:"time_zone,omitempty"`
}

type OrganizationWithMembership struct {
	CreatedAt time.Time `json:"created_at,omitempty"`

	// CRN Canonical Workspace resource identity.
	CRN         string `json:"crn,omitempty"`
	Description string `json:"description,omitempty"`
	ID          string `json:"id,omitempty"`

	// Language for organization billing and operational emails,
	// independent of each user's console preference.
	//
	// One of: "en", "pt-BR", "es".
	Language string `json:"language,omitempty"`
	Name     string `json:"name,omitempty"`

	// OwnerID ID of the organization owner
	OwnerID string `json:"owner_id,omitempty"`

	// Status lifecycle state. A newly created organization is `pending` until its
	// owner has verified a phone number and attached a payment method;
	// until then every resource API refuses it with
	// `ORGANIZATION_ONBOARDING_REQUIRED`. `suspended` is a billing or
	// administrative hold, and `terminated` is irreversible.
	//
	// One of: "pending", "active", "suspended", "terminated".
	Status string `json:"status,omitempty"`

	// SuspensionReason why the organization is suspended; absent unless it is. The two are
	// the same `status` but not the same situation — a `billing` hold is
	// one the customer can clear by settling their account, and the
	// platform still grants organization context for it so they can reach
	// billing to do so. A `manual` hold is an operator decision and grants
	// nothing.
	//
	// One of: "billing", "manual".
	SuspensionReason string `json:"suspension_reason,omitempty"`

	// TimeZone IANA timezone for formatting organization emails. Does not change
	// billing periods or resource schedules.
	TimeZone  string    `json:"time_zone,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type PermissionBoundary struct {
	CreatedAt   time.Time `json:"created_at,omitempty"`
	PolicyID    string    `json:"policy_id,omitempty"`
	PolicyName  string    `json:"policy_name,omitempty"`
	PrincipalID string    `json:"principal_id,omitempty"`

	// One of: "user".
	PrincipalType string `json:"principal_type,omitempty"`
}

type Policy struct {
	// CreatedAt creation timestamp (not present for system policies)
	CreatedAt time.Time `json:"created_at,omitempty"`

	// CRN managed policy CRN; absent on inline policy projections in
	// effective-policy lists.
	CRN         string          `json:"crn,omitempty"`
	Description string          `json:"description,omitempty"`
	Document    *PolicyDocument `json:"document,omitempty"`
	ID          string          `json:"id,omitempty"`

	// IsSystem whether this is a system-managed policy (cannot be modified or
	// deleted)
	IsSystem bool `json:"is_system,omitempty"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	Name string `json:"name,omitempty"`
	Tags Tags   `json:"tags,omitempty"`

	// UpdatedAt last update timestamp (not present for system policies)
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type PolicyAttachRequest struct {
	// Required.
	Policy PolicyReference `json:"policy"`
}

// PolicyCondition a condition that must be satisfied for the statement to apply
type PolicyCondition struct {
	// Key the condition key to evaluate
	Key string `json:"key"`

	// Operator the comparison operator
	//
	// One of: "equals", "not_equals", "starts_with", "ends_with", "contains", "in", "not_in", "greater_than", "less_than", "greater_than_or_equals", "less_than_or_equals", "exists", "not_exists", "ip_address", "not_ip_address".
	Operator string `json:"operator"`

	// SetOperator evaluates `operator` against a multi-valued context key (a set, such
	// as `basalt:TagKeys` — the tag keys a request carries) rather than
	// a single value. Omit for an ordinary single-valued condition.
	//
	// - `for_all_values` — holds when every member of the request set
	//   satisfies `operator`. An absent or empty set holds vacuously, so a
	//   request carrying no tags is not fenced by a tag-key restriction.
	// - `for_any_value` — holds when at least one member does. An absent or
	//   empty set does not hold.
	//
	// One of: "for_all_values", "for_any_value".
	SetOperator string `json:"set_operator,omitempty"`

	// Values to compare against
	Values []string `json:"values"`
}

type PolicyCreateRequest struct {
	Description *string `json:"description,omitempty"`

	// Required.
	Document *PolicyDocument `json:"document"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	//
	// Required.
	Name string `json:"name"`
	Tags Tags   `json:"tags,omitempty"`
}

// PolicyDocument IAM-style policy document
type PolicyDocument struct {
	Statements []*PolicyStatement `json:"statements"`

	// One of: "2024-01-01".
	Version string `json:"version"`
}

// PolicyReference organization policy UUID, immutable name in the authenticated
// organization, or fully qualified Workspace CRN. Shared system policies
// use crn:workspace:::system-policy/<name>. Account IAM policies cannot
// be attached through Workspace.
type PolicyReference = string

// PolicyStatement a single statement. The action set is named either positively
// (`actions`) or by exclusion (`not_actions`), and the resource set
// likewise (`resources` / `not_resources`) — exactly one of each pair.
// A statement that sets both sides of a pair, or neither, is rejected
// with `INVALID_INPUT` when the document is saved.
type PolicyStatement struct {
	// Actions in service:action format
	Actions []string `json:"actions,omitempty"`

	// Conditions optional conditions for the statement
	Conditions []*PolicyCondition `json:"conditions,omitempty"`

	// One of: "allow", "deny".
	Effect string `json:"effect"`

	// NotActions the statement covers every action *except* these. Pairs naturally
	// with `effect: deny` to carve a hole out of a broad allow; with
	// `effect: allow` it grants everything the listed patterns don't name,
	// including actions added by future services.
	NotActions []string `json:"not_actions,omitempty"`

	// NotResources the statement covers every resource *except* these. Same trade-off
	// as `not_actions`: with `effect: allow` it reaches resources that do
	// not exist yet.
	NotResources []string `json:"not_resources,omitempty"`

	// Resources resource identifiers or patterns
	Resources []string `json:"resources,omitempty"`

	// Sid statement identifier
	Sid string `json:"sid,omitempty"`
}

// PolicyUpdateRequest the resource name is immutable.
type PolicyUpdateRequest struct {
	Description *string         `json:"description,omitempty"`
	Document    *PolicyDocument `json:"document,omitempty"`
	Tags        Tags            `json:"tags,omitempty"`
}

type PutInlinePolicyRequest struct {
	// Required.
	Document *PolicyDocument `json:"document"`
}

type SetBoundaryRequest struct {
	// Required.
	Policy PolicyReference `json:"policy"`
}

type Tags = map[string]string

type UpdateAccountRequest struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
}

// User a platform user linked to the organization.
type User struct {
	AddedAt time.Time `json:"added_at,omitempty"`

	// CRN Cloud Resource Name
	CRN           string         `json:"crn,omitempty"`
	Email         string         `json:"email,omitempty"`
	ID            string         `json:"id,omitempty"`
	LinuxIdentity *LinuxIdentity `json:"linux_identity,omitempty"`
	Name          string         `json:"name,omitempty"`
	Tags          Tags           `json:"tags,omitempty"`

	// Username globally unique permanent login handle; empty until the user
	// completes username selection.
	Username string `json:"username,omitempty"`
}

type UserAddRequest struct {
	// Email of the user to add
	//
	// Required.
	Email string `json:"email"`

	// Groups to assign when the invitation is accepted. Each reference is
	// validated in the caller organization before the invitation is
	// created.
	Groups []GroupReference `json:"groups,omitempty"`
	Tags   Tags             `json:"tags,omitempty"`
}

type UserAddResponse struct {
	Invitation *Invitation `json:"invitation"`

	// Status always `invited` — adding a user always goes through an invitation
	// the invitee has to accept, whether or not they already have a
	// platform account.
	//
	// One of: "invited".
	Status string `json:"status"`
}

type UserGroupAddRequest struct {
	// Required.
	Group GroupReference `json:"group"`
}
