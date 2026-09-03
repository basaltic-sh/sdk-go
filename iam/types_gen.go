// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package iam

import (
	"time"
)

type Account struct {
	CreatedAt   time.Time `json:"created_at,omitempty"`
	Description string    `json:"description,omitempty"`

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

type AssumeRoleRequest struct {
	// DurationSeconds credential validity duration (15 min to 12 hours)
	DurationSeconds *int                   `json:"duration_seconds,omitempty"`
	Policy          *SessionPolicyDocument `json:"policy,omitempty"`

	// Required.
	RoleID string `json:"role_id"`
}

// AssumeRoleResponse a role session's credentials, in both forms it can be presented.
//
// `access_token` is a bearer token for this API — send it as
// `Authorization: Bearer <token>`. The other four fields are AWS SigV4
// credentials for the S3-compatible object storage endpoint, which
// speaks nothing else.
//
// Both come from the same session and share its expiry, so revoking the
// session stops both at once. Use whichever the endpoint you are calling
// needs; there is no need to choose one at request time.
type AssumeRoleResponse struct {
	// AccessKeyID SigV4 access key id, for the S3 endpoint.
	AccessKeyID string `json:"access_key_id,omitempty"`

	// AccessToken bearer token for the Basaltic API. Present on every role session.
	AccessToken string `json:"access_token,omitempty"`

	// Expiration when the session — and therefore both credential forms —
	// expires.
	Expiration time.Time `json:"expiration,omitempty"`

	// ExpiresIn seconds until `access_token` expires.
	ExpiresIn int `json:"expires_in,omitempty"`

	// SecretAccessKey SigV4 secret, for the S3 endpoint.
	SecretAccessKey string `json:"secret_access_key,omitempty"`

	// SessionToken SigV4 session token, for the S3 endpoint. Send as
	// `X-Amz-Security-Token` and include it in `SignedHeaders`.
	SessionToken string `json:"session_token,omitempty"`

	// TokenType always `Bearer` when `access_token` is present.
	TokenType string `json:"token_type,omitempty"`
}

// AssumeRoleWithWebIdentityRequest the exchange a federated caller sends. It carries no signature — the
// token is the credential — so every field is read from the body and
// nothing is inferred from the request context.
type AssumeRoleWithWebIdentityRequest struct {
	// AccountID the account the resulting credentials act in — the ownership scope
	// stamped on the session, the same scope a signed request selects with
	// `X-Account-Id`. Mind the difference in form: the header carries the
	// account handle, this field carries the account's id. Naming an
	// account does not widen the session; the role's own policies remain
	// the ceiling.
	//
	// Required.
	AccountID string `json:"account_id"`

	// DurationSeconds credential validity duration (15 min to 12 hours). A value above the
	// role's own `max_session_duration` is rejected rather than clamped.
	DurationSeconds *int `json:"duration_seconds,omitempty"`

	// RoleID the role to assume. Its trust policy has to admit this token — see
	// the operation description.
	//
	// Required.
	RoleID string `json:"role_id"`

	// SessionName a label recorded on the session and in the audit trail. Defaults to
	// the token's `sub` claim, so an unnamed session still records which
	// identity it came from.
	SessionName *string `json:"session_name,omitempty"`

	// WebIdentityToken the identity token to exchange, as a signed JWT. It is verified
	// before any role is read: the signature must chain to a key the
	// trusted provider publishes, the audience must be the one this
	// platform was configured to accept, and `exp` must be in the future.
	//
	// Required.
	WebIdentityToken string `json:"web_identity_token"`
}

type CreateAccountRequest struct {
	Description *string `json:"description,omitempty"`

	// Required.
	Handle string `json:"handle"`

	// Required.
	Name string `json:"name"`
}

type Credential struct {
	AccessKeyID string    `json:"access_key_id,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`

	// CRN Cloud Resource Name
	CRN        string    `json:"crn,omitempty"`
	ExpiresAt  time.Time `json:"expires_at,omitempty"`
	ID         string    `json:"id,omitempty"`
	LastUsedAt time.Time `json:"last_used_at,omitempty"`
	Name       string    `json:"name,omitempty"`
}

type CredentialCreateRequest struct {
	// ExpiresAt optional expiration date
	ExpiresAt *time.Time `json:"expires_at,omitempty"`

	// Required.
	Name string `json:"name"`
}

type CredentialCreateResponse struct {
	Credential *Credential `json:"credential,omitempty"`

	// SecretAccessKey only returned once at creation time
	SecretAccessKey string `json:"secret_access_key,omitempty"`
}

type Group struct {
	CreatedAt time.Time `json:"created_at,omitempty"`

	// CRN Cloud Resource Name
	CRN         string    `json:"crn,omitempty"`
	Description string    `json:"description,omitempty"`
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type GroupCreateRequest struct {
	Description *string `json:"description,omitempty"`

	// Required.
	Name string `json:"name"`
}

// GroupServiceAccount a service account in a group
type GroupServiceAccount struct {
	AddedAt time.Time `json:"added_at,omitempty"`

	// CRN of the service account
	CRN string `json:"crn,omitempty"`

	// Description of the service account, when it has one
	Description string `json:"description,omitempty"`
	ID          string `json:"id,omitempty"`

	// Name of the service account
	Name string `json:"name,omitempty"`
}

type GroupSummary struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type GroupUpdateRequest struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
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
	CreatedAt   time.Time       `json:"created_at,omitempty"`
	Document    *PolicyDocument `json:"document,omitempty"`
	ID          string          `json:"id,omitempty"`
	Name        string          `json:"name,omitempty"`
	PrincipalID string          `json:"principal_id,omitempty"`

	// One of: "user", "service_account", "role", "group".
	PrincipalType string    `json:"principal_type,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}

type Invitation struct {
	CreatedAt time.Time `json:"created_at,omitempty"`

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
	Email string `json:"email,omitempty"`
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
}

type ListRegionsResult struct {
	// Default the default region code
	Default string `json:"default"`

	// Regions list of available regions
	Regions []*Region `json:"regions"`
}

// OAuthAuthorizeRequest a signed-in user approving a client, from the console's consent page.
type OAuthAuthorizeRequest struct {
	// ClientID the registered client being approved.
	//
	// Required.
	ClientID string `json:"client_id"`

	// CodeChallenge Base64url SHA-256 of the client's PKCE verifier, without padding.
	//
	// Required.
	CodeChallenge string `json:"code_challenge"`

	// CodeChallengeMethod S256 only. `plain` is refused rather than merely discouraged:
	// whoever intercepts the code also saw the challenge, so a plain
	// challenge protects nothing.
	//
	// One of: "S256".
	//
	// Required.
	CodeChallengeMethod string `json:"code_challenge_method"`

	// OrganizationID which organization the resulting session is scoped to. The user must
	// be a member of it.
	//
	// Required.
	OrganizationID string `json:"organization_id"`

	// RedirectURI where to deliver the code. For the CLI this must be a loopback
	// address with any port — `http://127.0.0.1:<port>/...` or
	// `http://[::1]:<port>/...` (RFC 8252 section 7.3). `localhost` is
	// refused: it is a name, and whatever resolves it decides where the
	// code goes.
	//
	// Required.
	RedirectURI string `json:"redirect_uri"`

	// State opaque value echoed back on the redirect, unchanged. The client
	// generated it and compares it on return.
	State *string `json:"state,omitempty"`
}

type OAuthRevokeRequest struct {
	// Token the access token to revoke.
	//
	// Required.
	Token string `json:"token"`

	// TokenTypeHint accepted and ignored — the token identifies itself. Present
	// because RFC 7009 clients send it.
	TokenTypeHint *string `json:"token_type_hint,omitempty"`
}

// OAuthTokenRequest An OAuth 2.0 token request. Form-encoded is what RFC 6749 specifies
// and what client libraries send; JSON is accepted too.
//
// Client credentials may be sent as HTTP Basic (`Authorization: Basic
// base64(key_id:secret)`, which is what most libraries do by default) or
// as `client_id` and `client_secret` fields. Basic wins if both are
// present.
type OAuthTokenRequest struct {
	// ClientID the access key id. Omit when using HTTP Basic.
	ClientID *string `json:"client_id,omitempty"`

	// ClientSecret the secret access key. Omit when using HTTP Basic.
	ClientSecret *string `json:"client_secret,omitempty"`

	// Code the authorization code from the consent redirect. Single use, and
	// valid for five minutes. `authorization_code` grant only.
	Code *string `json:"code,omitempty"`

	// CodeVerifier The PKCE verifier whose SHA-256 was sent as `code_challenge` when
	// the flow started (RFC 7636). Required with `authorization_code`: it
	// is what proves this is the client that began the flow, since a CLI
	// holds no client secret.
	CodeVerifier *string `json:"code_verifier,omitempty"`

	// DurationSeconds requested token lifetime. A Basaltic extension, not an OAuth
	// parameter — omit it and you get the default. Values outside the
	// range are clamped into it rather than refused, so asking for a day
	// yields the longest token allowed.
	DurationSeconds *int `json:"duration_seconds,omitempty"`

	// GrantType `client_credentials` is the one to use for a service account: it
	// exchanges an access key pair for a token, and needs nothing else.
	//
	// `authorization_code` and `refresh_token` belong to the interactive
	// login a person runs (`basaltic login`), where the token names a USER
	// rather than a service account. They are driven by the CLI, not
	// written by hand. Check the authorization-server metadata document
	// before branching on them — they are advertised only where an
	// authorization endpoint is configured.
	//
	// One of: "client_credentials", "authorization_code", "refresh_token".
	//
	// Required.
	GrantType string `json:"grant_type"`

	// RedirectURI the same `redirect_uri` the code was issued for. Re-checked here, so
	// a code cannot be redeemed against a different destination.
	RedirectURI *string `json:"redirect_uri,omitempty"`

	// RefreshToken `refresh_token` grant only. Renews a user session without another
	// trip through the browser. Rotated on every use — store the new
	// one.
	RefreshToken *string `json:"refresh_token,omitempty"`
}

// OAuthTokenResponse RFC 6749 token response.
type OAuthTokenResponse struct {
	// AccessToken send as `Authorization: Bearer <token>`. Opaque to clients: do not
	// parse it, and do not key anything on the token string.
	AccessToken string `json:"access_token"`

	// ExpiresIn seconds until the token expires.
	ExpiresIn int `json:"expires_in"`

	// RefreshToken returned only by the user grants (`authorization_code` and
	// `refresh_token`). Present it to the `refresh_token` grant to renew
	// without another browser round trip; it is ROTATED on each use, so
	// replace the stored copy every time.
	//
	// A service account gets none. It already holds a long-lived access
	// key and can simply run `client_credentials` again, so a refresh
	// token would be a second credential to store for no gain.
	RefreshToken string `json:"refresh_token,omitempty"`

	// One of: "Bearer".
	TokenType string `json:"token_type"`
}

type Organization struct {
	CreatedAt   time.Time `json:"created_at,omitempty"`
	Description string    `json:"description,omitempty"`
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`

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
	SuspensionReason string    `json:"suspension_reason,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
}

type OrganizationUpdateRequest struct {
	// CaptchaToken google reCAPTCHA token for bot protection
	//
	// Required.
	CaptchaToken string  `json:"captcha_token"`
	Description  *string `json:"description,omitempty"`
	Name         *string `json:"name,omitempty"`
}

type OrganizationWithMembership struct {
	CreatedAt   time.Time `json:"created_at,omitempty"`
	Description string    `json:"description,omitempty"`
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`

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
	SuspensionReason string    `json:"suspension_reason,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
}

type PermissionBoundary struct {
	CreatedAt   time.Time `json:"created_at,omitempty"`
	PolicyID    string    `json:"policy_id,omitempty"`
	PolicyName  string    `json:"policy_name,omitempty"`
	PrincipalID string    `json:"principal_id,omitempty"`

	// One of: "user", "service_account", "role".
	PrincipalType string `json:"principal_type,omitempty"`
}

type Policy struct {
	// CreatedAt creation timestamp (not present for system policies)
	CreatedAt time.Time `json:"created_at,omitempty"`

	// CRN Cloud Resource Name
	CRN         string          `json:"crn,omitempty"`
	Description string          `json:"description,omitempty"`
	Document    *PolicyDocument `json:"document,omitempty"`
	ID          string          `json:"id,omitempty"`

	// IsSystem whether this is a system-managed policy (cannot be modified or
	// deleted)
	IsSystem bool   `json:"is_system,omitempty"`
	Name     string `json:"name,omitempty"`
	Tags     Tags   `json:"tags,omitempty"`

	// UpdatedAt last update timestamp (not present for system policies)
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type PolicyAttachRequest struct {
	// Required.
	PolicyID string `json:"policy_id"`
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

type PolicyUpdateRequest struct {
	Description *string         `json:"description,omitempty"`
	Document    *PolicyDocument `json:"document,omitempty"`
	Name        *string         `json:"name,omitempty"`
	Tags        Tags            `json:"tags,omitempty"`
}

type PutInlinePolicyRequest struct {
	// Required.
	Document *PolicyDocument `json:"document"`
}

type Region struct {
	// Available whether the region is currently available for use
	Available bool `json:"available"`

	// Code unique region code used in API calls and CRNs
	Code string `json:"code"`

	// ComingSoon whether the region is announced but not yet available
	ComingSoon bool `json:"coming_soon"`

	// CountryCode ISO 3166-1 alpha-2 country code (used to display flag in UI)
	CountryCode string `json:"country_code"`

	// Location geographic location of the region
	Location string `json:"location"`

	// Name human-readable region name
	Name string `json:"name"`
}

type RevokeSTSSessionRequest struct {
	// Reason for revoking the session
	Reason *string `json:"reason,omitempty"`
}

type Role struct {
	CreatedAt time.Time `json:"created_at,omitempty"`

	// CRN Cloud Resource Name
	CRN         string       `json:"crn,omitempty"`
	Description string       `json:"description,omitempty"`
	ID          string       `json:"id,omitempty"`
	Name        string       `json:"name,omitempty"`
	Tags        Tags         `json:"tags,omitempty"`
	TrustPolicy *TrustPolicy `json:"trust_policy,omitempty"`
	UpdatedAt   time.Time    `json:"updated_at,omitempty"`
}

type RoleCreateRequest struct {
	Description *string `json:"description,omitempty"`

	// Required.
	Name        string       `json:"name"`
	Tags        Tags         `json:"tags,omitempty"`
	TrustPolicy *TrustPolicy `json:"trust_policy,omitempty"`
}

type RolePolicyAttachRequest struct {
	// Required.
	PolicyID string `json:"policy_id"`
}

type RoleUpdateRequest struct {
	Description *string      `json:"description,omitempty"`
	Name        *string      `json:"name,omitempty"`
	Tags        Tags         `json:"tags,omitempty"`
	TrustPolicy *TrustPolicy `json:"trust_policy,omitempty"`
}

// STSSession a set of temporary credentials, as returned by the STS session
// listing. There is no "active" flag — a session is usable while
// `revoked` is false and `expires_at` is in the future.
type STSSession struct {
	CreatedAt  time.Time `json:"created_at,omitempty"`
	ExpiresAt  time.Time `json:"expires_at,omitempty"`
	ID         string    `json:"id,omitempty"`
	LastUsedAt time.Time `json:"last_used_at,omitempty"`

	// PrincipalID ID of the principal assuming the role (the source identity)
	PrincipalID string `json:"principal_id,omitempty"`

	// One of: "user", "service_account", "assumed_role".
	PrincipalType string `json:"principal_type,omitempty"`

	// Revoked whether the session has been revoked
	Revoked       bool      `json:"revoked,omitempty"`
	RevokedAt     time.Time `json:"revoked_at,omitempty"`
	RevokedReason string    `json:"revoked_reason,omitempty"`

	// RoleID the role being assumed. Empty on a session minted by `POST
	// /v1/session-token`, which is bound to the user directly and assumes
	// no role.
	RoleID string `json:"role_id,omitempty"`

	// SessionName optional session identifier
	SessionName string `json:"session_name,omitempty"`

	// SourceIP IP address where the session was created
	SourceIP string `json:"source_ip,omitempty"`

	// UserAgent User-Agent of the caller that created the session
	UserAgent string `json:"user_agent,omitempty"`
}

// ServiceAccount An API-only identity for programmatic access, bound to an account
type ServiceAccount struct {
	// AccountID owning account (from X-Account-Id at create time)
	AccountID string    `json:"account_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`

	// CRN Cloud Resource Name
	CRN         string    `json:"crn,omitempty"`
	Description string    `json:"description,omitempty"`
	Enabled     bool      `json:"enabled,omitempty"`
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Tags        Tags      `json:"tags,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type ServiceAccountCreateRequest struct {
	Description *string `json:"description,omitempty"`

	// Required.
	Name string `json:"name"`
	Tags Tags   `json:"tags,omitempty"`
}

type ServiceAccountGroupAddRequest struct {
	// Required.
	GroupID string `json:"group_id"`
}

type ServiceAccountUpdateRequest struct {
	Description *string `json:"description,omitempty"`
	Enabled     *bool   `json:"enabled,omitempty"`
	Tags        Tags    `json:"tags,omitempty"`
}

// SessionPolicyDocument an inline policy that scopes down the credentials being minted. It
// grants nothing on its own: every request made with the resulting
// credentials must be allowed by the caller's identity policies *and* by
// this document, so it can only narrow what the caller already has.
//
// The document is validated on the way in and stored with the session
// — an invalid one fails the call with `INVALID_INPUT` rather than
// being ignored. Statements take the same shape as in a managed policy
// but carry no `conditions`; a session policy fences on actions and
// resources only.
type SessionPolicyDocument struct {
	// Required.
	Statements []*SessionPolicyStatement `json:"statements"`

	// One of: "2024-01-01".
	//
	// Required.
	Version string `json:"version"`
}

// SessionPolicyStatement a session-policy statement. Same action/resource exclusivity as
// `PolicyStatement` — exactly one of `actions`/`not_actions` and one
// of `resources`/`not_resources`.
type SessionPolicyStatement struct {
	// Actions in service:action format
	Actions []string `json:"actions,omitempty"`

	// One of: "allow", "deny".
	//
	// Required.
	Effect string `json:"effect"`

	// NotActions the statement covers every action except these
	NotActions []string `json:"not_actions,omitempty"`

	// NotResources the statement covers every resource except these
	NotResources []string `json:"not_resources,omitempty"`

	// Resources resource identifiers or patterns
	Resources []string `json:"resources,omitempty"`

	// Sid statement identifier
	Sid *string `json:"sid,omitempty"`
}

type SetBoundaryRequest struct {
	// Required.
	PolicyID string `json:"policy_id"`
}

type Tags = map[string]string

// TrustPolicy defines who/what can assume this role using CRN patterns
type TrustPolicy struct {
	// Conditions optional conditions for role assumption
	Conditions []*PolicyCondition `json:"conditions,omitempty"`

	// Principals CRN patterns that identify who can assume this role. Supports
	// wildcards (*) for matching multiple resources.
	//
	// These match the caller's own CRN: an instance presents
	// `crn:compute:<region>:<account>:instance/<id>` to AssumeRole via
	// IMDS, a service account presents `crn:iam:::service-account/<id>`.
	//
	// A federated caller is the exception. It has no CRN of its own, so a
	// role that accepts one names the **provider** instead, as
	// `crn:iam:::oidc-provider/<provider>`, and pins the individual
	// identity with `conditions` on the token's claims. See `POST
	// /v1/assume-role-with-web-identity`.
	Principals []string `json:"principals,omitempty"`
}

type UpdateAccountRequest struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
}

// User a platform user linked to this organization for IAM
type User struct {
	AddedAt time.Time `json:"added_at,omitempty"`

	// CRN Cloud Resource Name
	CRN   string `json:"crn,omitempty"`
	Email string `json:"email,omitempty"`
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Tags  Tags   `json:"tags,omitempty"`
}

type UserAddRequest struct {
	// Email of the user to add
	//
	// Required.
	Email string `json:"email"`

	// GroupIDs IDs of groups to add the user to
	GroupIDs []string `json:"group_ids,omitempty"`
	Tags     Tags     `json:"tags,omitempty"`
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
	GroupID string `json:"group_id"`
}
