// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package iam

import (
	"context"
	"iter"
	"net/url"
	"strconv"

	basaltic "github.com/basaltic-sh/sdk-go"
)

// ListPoliciesParams are the optional filters and pagination controls for
// [Client.ListPolicies]. A nil *ListPoliciesParams sends none of them.
type ListPoliciesParams struct {
	// CRN exact returned CRN, combined with name using AND before pagination.
	// Malformed or empty CRNs return 400; valid mismatched or foreign CRNs
	// return an empty page. Resources without a CRN never match.
	// Organization-scoped CRNs use the authenticated organization.
	CRN string

	// Limit maximum number of items to return. A value above the maximum is
	// clamped to it rather than rejected, so a page shorter than the one
	// you asked for is normal — page until `meta.has_more` is false, not
	// until a page looks short.
	Limit int

	// Marker opaque pagination cursor. Echo back the `meta.marker` value from the
	// previous page to fetch the next one; do not construct or parse it.
	// The token's internal form varies by endpoint (a resource ID, a
	// timestamp, …) and is not guaranteed stable across releases.
	Marker string

	// Name exact resource name, combined with crn using AND before pagination.
	// Empty values are filters. Resources without a name never match.
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListPoliciesParams) query() url.Values {
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
func (p *ListPoliciesParams) withMarker(marker string) *ListPoliciesParams {
	var out ListPoliciesParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListPolicyRolesParams are the optional filters and pagination controls for
// [Client.ListPolicyRoles]. A nil *ListPolicyRolesParams sends none of them.
type ListPolicyRolesParams struct {
	// CRN exact returned CRN, combined with name using AND before pagination.
	// Malformed or empty CRNs return 400; valid mismatched or foreign CRNs
	// return an empty page. Resources without a CRN never match.
	// Organization-scoped CRNs use the authenticated organization.
	CRN string

	// Limit maximum number of items to return. A value above the maximum is
	// clamped to it rather than rejected, so a page shorter than the one
	// you asked for is normal — page until `meta.has_more` is false, not
	// until a page looks short.
	Limit int

	// Marker opaque pagination cursor. Echo back the `meta.marker` value from the
	// previous page to fetch the next one; do not construct or parse it.
	// The token's internal form varies by endpoint (a resource ID, a
	// timestamp, …) and is not guaranteed stable across releases.
	Marker string

	// Name exact resource name, combined with crn using AND before pagination.
	// Empty values are filters. Resources without a name never match.
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListPolicyRolesParams) query() url.Values {
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
func (p *ListPolicyRolesParams) withMarker(marker string) *ListPolicyRolesParams {
	var out ListPolicyRolesParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListPolicyServiceAccountsParams are the optional filters and pagination controls for
// [Client.ListPolicyServiceAccounts]. A nil *ListPolicyServiceAccountsParams sends none of them.
type ListPolicyServiceAccountsParams struct {
	// CRN exact returned CRN, combined with name using AND before pagination.
	// Malformed or empty CRNs return 400; valid mismatched or foreign CRNs
	// return an empty page. Resources without a CRN never match.
	// Organization-scoped CRNs use the authenticated organization.
	CRN string

	// Limit maximum number of items to return. A value above the maximum is
	// clamped to it rather than rejected, so a page shorter than the one
	// you asked for is normal — page until `meta.has_more` is false, not
	// until a page looks short.
	Limit int

	// Marker opaque pagination cursor. Echo back the `meta.marker` value from the
	// previous page to fetch the next one; do not construct or parse it.
	// The token's internal form varies by endpoint (a resource ID, a
	// timestamp, …) and is not guaranteed stable across releases.
	Marker string

	// Name exact resource name, combined with crn using AND before pagination.
	// Empty values are filters. Resources without a name never match.
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListPolicyServiceAccountsParams) query() url.Values {
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
func (p *ListPolicyServiceAccountsParams) withMarker(marker string) *ListPolicyServiceAccountsParams {
	var out ListPolicyServiceAccountsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListRegionsParams are the optional filters and pagination controls for
// [Client.ListRegions]. A nil *ListRegionsParams sends none of them.
type ListRegionsParams struct {
	// CRN exact returned CRN, combined with name using AND before pagination.
	// Malformed or empty CRNs return 400; valid mismatched or foreign CRNs
	// return an empty page. Resources without a CRN never match.
	// Organization-scoped CRNs use the authenticated organization.
	CRN string

	// Name exact resource name, combined with crn using AND before pagination.
	// Empty values are filters. Resources without a name never match.
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListRegionsParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.CRN != "" {
		q.Set("crn", p.CRN)
	}
	if p.Name != "" {
		q.Set("name", p.Name)
	}
	return q
}

// ListRoleInlinePoliciesParams are the optional filters and pagination controls for
// [Client.ListRoleInlinePolicies]. A nil *ListRoleInlinePoliciesParams sends none of them.
type ListRoleInlinePoliciesParams struct {
	// CRN exact returned CRN, combined with name using AND before pagination.
	// Malformed or empty CRNs return 400; valid mismatched or foreign CRNs
	// return an empty page. Resources without a CRN never match.
	// Organization-scoped CRNs use the authenticated organization.
	CRN string

	// Name exact resource name, combined with crn using AND before pagination.
	// Empty values are filters. Resources without a name never match.
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListRoleInlinePoliciesParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.CRN != "" {
		q.Set("crn", p.CRN)
	}
	if p.Name != "" {
		q.Set("name", p.Name)
	}
	return q
}

// ListRolePoliciesParams are the optional filters and pagination controls for
// [Client.ListRolePolicies]. A nil *ListRolePoliciesParams sends none of them.
type ListRolePoliciesParams struct {
	// CRN exact returned CRN, combined with name using AND before pagination.
	// Malformed or empty CRNs return 400; valid mismatched or foreign CRNs
	// return an empty page. Resources without a CRN never match.
	// Organization-scoped CRNs use the authenticated organization.
	CRN string

	// Name exact resource name, combined with crn using AND before pagination.
	// Empty values are filters. Resources without a name never match.
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListRolePoliciesParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.CRN != "" {
		q.Set("crn", p.CRN)
	}
	if p.Name != "" {
		q.Set("name", p.Name)
	}
	return q
}

// ListRolesParams are the optional filters and pagination controls for
// [Client.ListRoles]. A nil *ListRolesParams sends none of them.
type ListRolesParams struct {
	// CRN exact returned CRN, combined with name using AND before pagination.
	// Malformed or empty CRNs return 400; valid mismatched or foreign CRNs
	// return an empty page. Resources without a CRN never match.
	// Organization-scoped CRNs use the authenticated organization.
	CRN string

	// Limit maximum number of items to return. A value above the maximum is
	// clamped to it rather than rejected, so a page shorter than the one
	// you asked for is normal — page until `meta.has_more` is false, not
	// until a page looks short.
	Limit int

	// Marker opaque pagination cursor. Echo back the `meta.marker` value from the
	// previous page to fetch the next one; do not construct or parse it.
	// The token's internal form varies by endpoint (a resource ID, a
	// timestamp, …) and is not guaranteed stable across releases.
	Marker string

	// Name exact resource name, combined with crn using AND before pagination.
	// Empty values are filters. Resources without a name never match.
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListRolesParams) query() url.Values {
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
func (p *ListRolesParams) withMarker(marker string) *ListRolesParams {
	var out ListRolesParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListSTSSessionsParams are the optional filters and pagination controls for
// [Client.ListSTSSessions]. A nil *ListSTSSessionsParams sends none of them.
type ListSTSSessionsParams struct {
	// ActiveOnly only show active (non-expired, non-revoked) sessions
	ActiveOnly *bool

	// CRN exact returned CRN, combined with name using AND before pagination.
	// Malformed or empty CRNs return 400; valid mismatched or foreign CRNs
	// return an empty page. Resources without a CRN never match.
	// Organization-scoped CRNs use the authenticated organization.
	CRN string

	// Limit maximum number of items to return. A value above the maximum is
	// clamped to it rather than rejected, so a page shorter than the one
	// you asked for is normal — page until `meta.has_more` is false, not
	// until a page looks short.
	Limit int

	// Marker opaque pagination cursor. Echo back the `meta.marker` value from the
	// previous page to fetch the next one; do not construct or parse it.
	// The token's internal form varies by endpoint (a resource ID, a
	// timestamp, …) and is not guaranteed stable across releases.
	Marker string

	// Name exact resource name, combined with crn using AND before pagination.
	// Empty values are filters. Resources without a name never match.
	Name string

	// Principal reference; principal_type is required. Users and
	// assumed-role sessions accept UUID or CRN; roles and service accounts
	// also accept names in the owning account.
	Principal string

	// PrincipalType filter by principal type. `assumed_role` selects the sessions role
	// chaining produces, where an existing assumed-role session assumed
	// another role.
	//
	//
	// One of: "user", "service_account", "role", "assumed_role".
	PrincipalType string

	// Role UUID, immutable name, or account-qualified role CRN.
	Role string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListSTSSessionsParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.ActiveOnly != nil {
		q.Set("active_only", strconv.FormatBool(*p.ActiveOnly))
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
	if p.Principal != "" {
		q.Set("principal", p.Principal)
	}
	if p.PrincipalType != "" {
		q.Set("principal_type", p.PrincipalType)
	}
	if p.Role != "" {
		q.Set("role", p.Role)
	}
	return q
}

// withMarker copies p with the pagination cursor replaced, leaving the
// caller's value untouched across pages.
func (p *ListSTSSessionsParams) withMarker(marker string) *ListSTSSessionsParams {
	var out ListSTSSessionsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListServiceAccountCredentialsParams are the optional filters and pagination controls for
// [Client.ListServiceAccountCredentials]. A nil *ListServiceAccountCredentialsParams sends none of them.
type ListServiceAccountCredentialsParams struct {
	// CRN exact returned CRN, combined with name using AND before pagination.
	// Malformed or empty CRNs return 400; valid mismatched or foreign CRNs
	// return an empty page. Resources without a CRN never match.
	// Organization-scoped CRNs use the authenticated organization.
	CRN string

	// Name exact resource name, combined with crn using AND before pagination.
	// Empty values are filters. Resources without a name never match.
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListServiceAccountCredentialsParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.CRN != "" {
		q.Set("crn", p.CRN)
	}
	if p.Name != "" {
		q.Set("name", p.Name)
	}
	return q
}

// ListServiceAccountInlinePoliciesParams are the optional filters and pagination controls for
// [Client.ListServiceAccountInlinePolicies]. A nil *ListServiceAccountInlinePoliciesParams sends none of them.
type ListServiceAccountInlinePoliciesParams struct {
	// CRN exact returned CRN, combined with name using AND before pagination.
	// Malformed or empty CRNs return 400; valid mismatched or foreign CRNs
	// return an empty page. Resources without a CRN never match.
	// Organization-scoped CRNs use the authenticated organization.
	CRN string

	// Name exact resource name, combined with crn using AND before pagination.
	// Empty values are filters. Resources without a name never match.
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListServiceAccountInlinePoliciesParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.CRN != "" {
		q.Set("crn", p.CRN)
	}
	if p.Name != "" {
		q.Set("name", p.Name)
	}
	return q
}

// ListServiceAccountPoliciesParams are the optional filters and pagination controls for
// [Client.ListServiceAccountPolicies]. A nil *ListServiceAccountPoliciesParams sends none of them.
type ListServiceAccountPoliciesParams struct {
	// CRN exact returned CRN, combined with name using AND before pagination.
	// Malformed or empty CRNs return 400; valid mismatched or foreign CRNs
	// return an empty page. Resources without a CRN never match.
	// Organization-scoped CRNs use the authenticated organization.
	CRN string

	// Name exact resource name, combined with crn using AND before pagination.
	// Empty values are filters. Resources without a name never match.
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListServiceAccountPoliciesParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.CRN != "" {
		q.Set("crn", p.CRN)
	}
	if p.Name != "" {
		q.Set("name", p.Name)
	}
	return q
}

// ListServiceAccountsParams are the optional filters and pagination controls for
// [Client.ListServiceAccounts]. A nil *ListServiceAccountsParams sends none of them.
type ListServiceAccountsParams struct {
	// CRN exact returned CRN, combined with name using AND before pagination.
	// Malformed or empty CRNs return 400; valid mismatched or foreign CRNs
	// return an empty page. Resources without a CRN never match.
	// Organization-scoped CRNs use the authenticated organization.
	CRN string

	// Limit maximum number of items to return. A value above the maximum is
	// clamped to it rather than rejected, so a page shorter than the one
	// you asked for is normal — page until `meta.has_more` is false, not
	// until a page looks short.
	Limit int

	// Marker opaque pagination cursor. Echo back the `meta.marker` value from the
	// previous page to fetch the next one; do not construct or parse it.
	// The token's internal form varies by endpoint (a resource ID, a
	// timestamp, …) and is not guaranteed stable across releases.
	Marker string

	// Name exact resource name, combined with crn using AND before pagination.
	// Empty values are filters. Resources without a name never match.
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListServiceAccountsParams) query() url.Values {
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
func (p *ListServiceAccountsParams) withMarker(marker string) *ListServiceAccountsParams {
	var out ListServiceAccountsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// AssumeRole assumes role.
//
// Assume a role in this account or another account in the same
// organization. The caller needs iam:AssumeRole permission (or an
// effective Workspace account-role assignment), and the target role must
// independently trust the caller.
//
// The session account comes from the target role, not the
// selected-account header. The session uses the target role permissions
// and delegated organization policies, subject to applicable boundaries,
// explicit denies, and an optional narrowing session policy. Source
// principal permissions are not unioned into the session.
func (c *Client) AssumeRole(ctx context.Context, body *AssumeRoleRequest, opts ...basaltic.RequestOption) (*AssumeRoleResponse, error) {
	op := &basaltic.Operation{
		ID:     "assumeRole",
		Method: "POST",
		Path:   "/v1/assume-role",
		Body:   body,
	}
	var out AssumeRoleResponse
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// AssumeRoleWithWebIdentity assumes role with web identity.
//
// Exchange an identity token issued by a federation provider this
// platform trusts for temporary credentials. The result is the same
// assumed-role session `POST /v1/assume-role` mints, and is used the
// same way.
//
// This request carries **no signature**, and it is the only
// credential-vending call that does not. A federated caller holds no
// Basaltic credential yet — that is what the exchange is for — so
// the token in the body *is* the credential being presented. A signature
// sent anyway is ignored, and nothing is taken from the request context:
// `role` and `account` are read from the body like every other field.
//
// That does not leave the endpoint open. Two independent gates have to
// pass, and they fail differently.
//
// **The token has to verify.** This happens before any role is read, so
// a forged token never reaches a trust policy. The signature must chain
// to a key the provider publishes, the audience must be the one this
// platform accepts, and `exp` must be in the future. A wrong signer, a
// token minted for some other consumer, and an expired token all answer
// `401` with the same message — the response does not say which check
// failed.
//
// **The role has to agree.** Verifying the token establishes who is
// calling; it grants nothing. The role named in `role` is assumable only
// if its own trust policy admits this identity. Its `principals` must
// name the federation provider, written
// `crn:iam:::oidc-provider/<provider>` — the one case where a trust
// policy principal is not the caller's own CRN, because a federated
// identity has no CRN and what is trusted is the source that vouched for
// it. Every entry in `conditions` must then hold against the token's
// claims: `basalt:webidentity:Subject` carries the token's `sub` and
// `basalt:webidentity:Audience` its `aud`, so a role can bind one
// identity instead of accepting everything that provider will ever
// issue. A condition on a claim the token does not carry fails closed.
//
// A role whose trust policy names no provider therefore cannot be
// assumed this way at all, however good the token is. That is the line
// between the two failures: `401` means the token is not trustworthy,
// `403` means it is and the role still will not have it.
//
// The credentials come back scoped to `account`, carrying the role's own
// permissions. There is no `policy` field here — unlike `POST
// /v1/assume-role`, a federated session cannot be scoped down at
// exchange time, so the role's attached policies are the whole grant.
// Size the role accordingly.
//
// Because it takes no credentials, requests are rate-limited per client
// IP.
//
// Which providers are trusted is part of the platform's own
// configuration. There is no API for registering an identity provider of
// your own yet, so this operation is live but has no external provider
// whose tokens it would accept; the roles that use it today are
// platform-managed.
//
// Sends no bearer token: the credentials in the request are the
// authentication.
func (c *Client) AssumeRoleWithWebIdentity(ctx context.Context, body *AssumeRoleWithWebIdentityRequest, opts ...basaltic.RequestOption) (*AssumeRoleResponse, error) {
	op := &basaltic.Operation{
		ID:              "assumeRoleWithWebIdentity",
		Method:          "POST",
		Path:            "/v1/assume-role-with-web-identity",
		Body:            body,
		Unauthenticated: true,
	}
	var out AssumeRoleResponse
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// AttachRolePolicy attaches policy to role.
//
// Attach a policy to a role.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) AttachRolePolicy(ctx context.Context, roleID string, body *RolePolicyAttachRequest, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "attachRolePolicy",
		Method:   "POST",
		Path:     "/v1/roles/{role_id}/policies",
		PathArgs: []string{roleID},
		Body:     body,
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// AttachServiceAccountPolicy attaches policy to service account.
//
// Attach a policy to a service account.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) AttachServiceAccountPolicy(ctx context.Context, serviceAccountID string, body *PolicyAttachRequest, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "attachServiceAccountPolicy",
		Method:   "POST",
		Path:     "/v1/service-accounts/{service_account_id}/policies",
		PathArgs: []string{serviceAccountID},
		Body:     body,
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// AuthorizeOAuthClient — Approve a CLI login and issue an authorization code.
//
// Approve a client to act as you, and receive the authorization code
// that hands it a token.
//
// **This is the console's endpoint, not a client's.** It is called by
// the Basaltic console's consent page on behalf of a signed-in user; the
// CLI never calls it. A CLI prints a URL for that page, and the page
// calls this. Anything driving it directly would need the user's console
// session, at which point it already has everything the code would
// grant.
//
// It is the half of the authorization-code flow that establishes WHO is
// approving. The user must already be signed in — including any second
// factor — and must be a member of the organization named in
// `organization`. The organization is explicit rather than inferred: a
// person in several has no single obvious answer, and choosing one for
// them would scope the resulting token to something they did not pick.
//
// Unlike the token endpoint, this answers in the usual API envelope. It
// is not part of the surface a third-party OAuth client talks to, so it
// follows the caller — and the caller is our own front end.
func (c *Client) AuthorizeOAuthClient(ctx context.Context, body *OAuthAuthorizeRequest, opts ...basaltic.RequestOption) (*OAuthAuthorizeResponse, error) {
	op := &basaltic.Operation{
		ID:     "authorizeOAuthClient",
		Method: "POST",
		Path:   "/v1/oauth/authorize",
		Body:   body,
	}
	var out OAuthAuthorizeResponse
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// CreatePersonalSSHKey adds personal SSH key.
//
// Requires the signed-in human user. Personal keys apply to that user
// across organization memberships; the organization selects only the
// Linux identity. A key never carries a role or grants access to a VM by
// itself. Linux names and numeric IDs are allocated by the platform and
// do not change on key rotation or display-name changes.
func (c *Client) CreatePersonalSSHKey(ctx context.Context, body *SSHKeyCreateRequest, opts ...basaltic.RequestOption) (*SSHKey, error) {
	op := &basaltic.Operation{
		ID:     "createPersonalSSHKey",
		Method: "POST",
		Path:   "/v1/auth/ssh-keys",
		Body:   body,
	}
	var out struct {
		SSHKey *SSHKey `json:"ssh_key"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.SSHKey, nil
}

// CreatePolicy creates policy.
//
// Create a new policy in the selected account.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreatePolicy(ctx context.Context, body *PolicyCreateRequest, opts ...basaltic.RequestOption) (*Policy, error) {
	op := &basaltic.Operation{
		ID:     "createPolicy",
		Method: "POST",
		Path:   "/v1/policies",
		Body:   body,
	}
	var out struct {
		Policy *Policy `json:"policy"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Policy, nil
}

// CreateRole creates role.
//
// Create a new role in the selected account.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateRole(ctx context.Context, body *RoleCreateRequest, opts ...basaltic.RequestOption) (*Role, error) {
	op := &basaltic.Operation{
		ID:     "createRole",
		Method: "POST",
		Path:   "/v1/roles",
		Body:   body,
	}
	var out struct {
		Role *Role `json:"role"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Role, nil
}

// CreateServiceAccount creates service account.
//
// Create a new service account bound to the account selected via
// X-Account-Id. Permanent keys minted for the SA authenticate to S3 as
// that account.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateServiceAccount(ctx context.Context, body *ServiceAccountCreateRequest, opts ...basaltic.RequestOption) (*ServiceAccount, error) {
	op := &basaltic.Operation{
		ID:     "createServiceAccount",
		Method: "POST",
		Path:   "/v1/service-accounts",
		Body:   body,
	}
	var out struct {
		ServiceAccount *ServiceAccount `json:"service_account"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.ServiceAccount, nil
}

// CreateServiceAccountCredential creates credential.
//
// Create a new credential for a service account. The secret access key
// is only returned once at creation time.
func (c *Client) CreateServiceAccountCredential(ctx context.Context, serviceAccountID string, body *CredentialCreateRequest, opts ...basaltic.RequestOption) (*CredentialCreateResponse, error) {
	op := &basaltic.Operation{
		ID:       "createServiceAccountCredential",
		Method:   "POST",
		Path:     "/v1/service-accounts/{service_account_id}/credentials",
		PathArgs: []string{serviceAccountID},
		Body:     body,
	}
	var out CredentialCreateResponse
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// CreateServiceAccountSSHKey adds service-account SSH key.
//
// Requires the selected account and ordinary IAM authorization against
// the service account, including tag conditions and explicit denies. A
// key never carries a role or grants access to a VM by itself. Linux
// names and numeric IDs are allocated by the platform and do not change
// on key rotation or display-name changes.
func (c *Client) CreateServiceAccountSSHKey(ctx context.Context, serviceAccountID string, body *SSHKeyCreateRequest, opts ...basaltic.RequestOption) (*SSHKey, error) {
	op := &basaltic.Operation{
		ID:       "createServiceAccountSSHKey",
		Method:   "POST",
		Path:     "/v1/service-accounts/{service_account_id}/ssh-keys",
		PathArgs: []string{serviceAccountID},
		Body:     body,
	}
	var out struct {
		SSHKey *SSHKey `json:"ssh_key"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.SSHKey, nil
}

// DeletePersonalSSHKey revokes personal SSH key.
//
// Requires the signed-in human user. Personal keys apply to that user
// across organization memberships; the organization selects only the
// Linux identity. A key never carries a role or grants access to a VM by
// itself. Linux names and numeric IDs are allocated by the platform and
// do not change on key rotation or display-name changes.
func (c *Client) DeletePersonalSSHKey(ctx context.Context, sshKeyID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deletePersonalSSHKey",
		Method:   "DELETE",
		Path:     "/v1/auth/ssh-keys/{ssh_key_id}",
		PathArgs: []string{sshKeyID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeletePolicy deletes policy.
//
// Delete a policy.
func (c *Client) DeletePolicy(ctx context.Context, policyID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deletePolicy",
		Method:   "DELETE",
		Path:     "/v1/policies/{policy_id}",
		PathArgs: []string{policyID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteRole deletes role.
//
// Delete a role.
func (c *Client) DeleteRole(ctx context.Context, roleID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteRole",
		Method:   "DELETE",
		Path:     "/v1/roles/{role_id}",
		PathArgs: []string{roleID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteRoleInlinePolicy deletes a role's inline policy by name.
func (c *Client) DeleteRoleInlinePolicy(ctx context.Context, roleID string, policyName string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteRoleInlinePolicy",
		Method:   "DELETE",
		Path:     "/v1/roles/{role_id}/inline-policies/{policy_name}",
		PathArgs: []string{roleID, policyName},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteServiceAccount deletes service account.
//
// Delete a service account and all its credentials.
func (c *Client) DeleteServiceAccount(ctx context.Context, serviceAccountID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteServiceAccount",
		Method:   "DELETE",
		Path:     "/v1/service-accounts/{service_account_id}",
		PathArgs: []string{serviceAccountID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteServiceAccountCredential deletes credential.
//
// Delete a credential.
func (c *Client) DeleteServiceAccountCredential(ctx context.Context, serviceAccountID string, credentialID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteServiceAccountCredential",
		Method:   "DELETE",
		Path:     "/v1/service-accounts/{service_account_id}/credentials/{credential_id}",
		PathArgs: []string{serviceAccountID, credentialID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteServiceAccountInlinePolicy deletes a service account's inline policy by name.
func (c *Client) DeleteServiceAccountInlinePolicy(ctx context.Context, serviceAccountID string, policyName string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteServiceAccountInlinePolicy",
		Method:   "DELETE",
		Path:     "/v1/service-accounts/{service_account_id}/inline-policies/{policy_name}",
		PathArgs: []string{serviceAccountID, policyName},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteServiceAccountSSHKey revokes service-account SSH key.
//
// Requires the selected account and ordinary IAM authorization against
// the service account, including tag conditions and explicit denies. A
// key never carries a role or grants access to a VM by itself. Linux
// names and numeric IDs are allocated by the platform and do not change
// on key rotation or display-name changes.
func (c *Client) DeleteServiceAccountSSHKey(ctx context.Context, serviceAccountID string, sshKeyID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteServiceAccountSSHKey",
		Method:   "DELETE",
		Path:     "/v1/service-accounts/{service_account_id}/ssh-keys/{ssh_key_id}",
		PathArgs: []string{serviceAccountID, sshKeyID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DetachRolePolicy detaches policy from role.
//
// Detach a policy from a role.
func (c *Client) DetachRolePolicy(ctx context.Context, roleID string, policyID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "detachRolePolicy",
		Method:   "DELETE",
		Path:     "/v1/roles/{role_id}/policies/{policy_id}",
		PathArgs: []string{roleID, policyID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DetachServiceAccountPolicy detaches policy from service account.
//
// Detach a policy from a service account.
func (c *Client) DetachServiceAccountPolicy(ctx context.Context, serviceAccountID string, policyID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "detachServiceAccountPolicy",
		Method:   "DELETE",
		Path:     "/v1/service-accounts/{service_account_id}/policies/{policy_id}",
		PathArgs: []string{serviceAccountID, policyID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// GetOAuthToken exchanges an access key for a bearer token.
//
// Exchange a service account's access key pair for a short-lived bearer
// token, then send that token as `Authorization: Bearer <token>` on
// every other call.
//
// This is the ordinary way to authenticate. The access key pair stays
// the one long-lived credential a service account has; what changes is
// that you present a token derived from it rather than signing each
// request.
//
// ```
//
//	curl -s -u "$KEY_ID:$SECRET" -d grant_type=client_credentials \
//	  https://iam.basaltic.sh/v1/oauth/token
//
// ```
//
// The same key pair is *also* the AWS SigV4 credential for the
// S3-compatible object endpoint, which speaks nothing else. Use the
// token for this API and the key pair for S3; there is no need to
// choose.
//
// **Errors here use the OAuth 2.0 shape, not this API's usual envelope**
// — `{"error": "...", "error_description": "..."}` — because every
// OAuth client library parses that and nothing else. Two answers matter
// and their remedies are opposite. `invalid_client` means the key was
// rejected: check or rotate it. `invalid_grant` means the key is fine
// and the organization is suspended or still onboarding, where rotating
// a working key would waste your time.
//
// An unknown access key and a wrong secret both answer `invalid_client`
// with the same message, so the endpoint cannot be used to discover
// which keys exist.
//
// Sends no bearer token: the credentials in the request are the
// authentication.
func (c *Client) GetOAuthToken(ctx context.Context, body *OAuthTokenRequest, opts ...basaltic.RequestOption) (*OAuthTokenResponse, error) {
	op := &basaltic.Operation{
		ID:              "getOAuthToken",
		Method:          "POST",
		Path:            "/v1/oauth/token",
		Body:            body,
		Unauthenticated: true,
	}
	var out OAuthTokenResponse
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetPersonalLinuxIdentity gets personal Linux identity.
//
// Requires the signed-in human user. Personal keys apply to that user
// across organization memberships; the organization selects only the
// Linux identity. A key never carries a role or grants access to a VM by
// itself. Linux names and numeric IDs are allocated by the platform and
// do not change on key rotation or display-name changes.
func (c *Client) GetPersonalLinuxIdentity(ctx context.Context, opts ...basaltic.RequestOption) (*LinuxIdentity, error) {
	op := &basaltic.Operation{
		ID:     "getPersonalLinuxIdentity",
		Method: "GET",
		Path:   "/v1/auth/linux-identity",
	}
	var out struct {
		LinuxIdentity *LinuxIdentity `json:"linux_identity"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.LinuxIdentity, nil
}

// GetPolicy gets policy.
//
// Get details of a specific policy.
func (c *Client) GetPolicy(ctx context.Context, policyID string, opts ...basaltic.RequestOption) (*Policy, error) {
	op := &basaltic.Operation{
		ID:       "getPolicy",
		Method:   "GET",
		Path:     "/v1/policies/{policy_id}",
		PathArgs: []string{policyID},
	}
	var out struct {
		Policy *Policy `json:"policy"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Policy, nil
}

// GetRole gets role.
//
// Get details of a specific role.
func (c *Client) GetRole(ctx context.Context, roleID string, opts ...basaltic.RequestOption) (*Role, error) {
	op := &basaltic.Operation{
		ID:       "getRole",
		Method:   "GET",
		Path:     "/v1/roles/{role_id}",
		PathArgs: []string{roleID},
	}
	var out struct {
		Role *Role `json:"role"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Role, nil
}

// GetRoleInlinePolicy gets a role's inline policy by name.
func (c *Client) GetRoleInlinePolicy(ctx context.Context, roleID string, policyName string, opts ...basaltic.RequestOption) (*InlinePolicy, error) {
	op := &basaltic.Operation{
		ID:       "getRoleInlinePolicy",
		Method:   "GET",
		Path:     "/v1/roles/{role_id}/inline-policies/{policy_name}",
		PathArgs: []string{roleID, policyName},
	}
	var out struct {
		InlinePolicy *InlinePolicy `json:"inline_policy"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.InlinePolicy, nil
}

// GetRolePermissionBoundary gets a role's permission boundary.
func (c *Client) GetRolePermissionBoundary(ctx context.Context, roleID string, opts ...basaltic.RequestOption) (*PermissionBoundary, error) {
	op := &basaltic.Operation{
		ID:       "getRolePermissionBoundary",
		Method:   "GET",
		Path:     "/v1/roles/{role_id}/permission-boundary",
		PathArgs: []string{roleID},
	}
	var out struct {
		PermissionBoundary *PermissionBoundary `json:"permission_boundary"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.PermissionBoundary, nil
}

// GetSTSSession gets STS session.
//
// Get details of a specific STS session. Requires `iam:GetSTSSession`
// permission.
func (c *Client) GetSTSSession(ctx context.Context, sessionID string, opts ...basaltic.RequestOption) (*STSSession, error) {
	op := &basaltic.Operation{
		ID:       "getSTSSession",
		Method:   "GET",
		Path:     "/v1/sts-sessions/{session_id}",
		PathArgs: []string{sessionID},
	}
	var out struct {
		STSSession *STSSession `json:"sts_session"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.STSSession, nil
}

// GetServiceAccount gets service account.
//
// Get details of a specific service account.
func (c *Client) GetServiceAccount(ctx context.Context, serviceAccountID string, opts ...basaltic.RequestOption) (*ServiceAccount, error) {
	op := &basaltic.Operation{
		ID:       "getServiceAccount",
		Method:   "GET",
		Path:     "/v1/service-accounts/{service_account_id}",
		PathArgs: []string{serviceAccountID},
	}
	var out struct {
		ServiceAccount *ServiceAccount `json:"service_account"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.ServiceAccount, nil
}

// GetServiceAccountInlinePolicy gets a service account's inline policy by name.
func (c *Client) GetServiceAccountInlinePolicy(ctx context.Context, serviceAccountID string, policyName string, opts ...basaltic.RequestOption) (*InlinePolicy, error) {
	op := &basaltic.Operation{
		ID:       "getServiceAccountInlinePolicy",
		Method:   "GET",
		Path:     "/v1/service-accounts/{service_account_id}/inline-policies/{policy_name}",
		PathArgs: []string{serviceAccountID, policyName},
	}
	var out struct {
		InlinePolicy *InlinePolicy `json:"inline_policy"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.InlinePolicy, nil
}

// GetServiceAccountLinuxIdentity gets serviceaccount Linux identity.
//
// Requires the selected account and ordinary IAM authorization against
// the service account, including tag conditions and explicit denies. A
// key never carries a role or grants access to a VM by itself. Linux
// names and numeric IDs are allocated by the platform and do not change
// on key rotation or display-name changes.
func (c *Client) GetServiceAccountLinuxIdentity(ctx context.Context, serviceAccountID string, opts ...basaltic.RequestOption) (*LinuxIdentity, error) {
	op := &basaltic.Operation{
		ID:       "getServiceAccountLinuxIdentity",
		Method:   "GET",
		Path:     "/v1/service-accounts/{service_account_id}/linux-identity",
		PathArgs: []string{serviceAccountID},
	}
	var out struct {
		LinuxIdentity *LinuxIdentity `json:"linux_identity"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.LinuxIdentity, nil
}

// GetServiceAccountPermissionBoundary gets a service account's permission boundary.
func (c *Client) GetServiceAccountPermissionBoundary(ctx context.Context, serviceAccountID string, opts ...basaltic.RequestOption) (*PermissionBoundary, error) {
	op := &basaltic.Operation{
		ID:       "getServiceAccountPermissionBoundary",
		Method:   "GET",
		Path:     "/v1/service-accounts/{service_account_id}/permission-boundary",
		PathArgs: []string{serviceAccountID},
	}
	var out struct {
		PermissionBoundary *PermissionBoundary `json:"permission_boundary"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.PermissionBoundary, nil
}

// ListPersonalSSHKeys lists personal SSH keys.
//
// Requires the signed-in human user. Personal keys apply to that user
// across organization memberships; the organization selects only the
// Linux identity. A key never carries a role or grants access to a VM by
// itself. Linux names and numeric IDs are allocated by the platform and
// do not change on key rotation or display-name changes.
func (c *Client) ListPersonalSSHKeys(ctx context.Context, opts ...basaltic.RequestOption) ([]*SSHKey, error) {
	op := &basaltic.Operation{
		ID:     "listPersonalSSHKeys",
		Method: "GET",
		Path:   "/v1/auth/ssh-keys",
	}
	var out struct {
		SSHKeys []*SSHKey `json:"ssh_keys"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.SSHKeys, nil
}

// ListPolicies lists policies.
//
// List all policies in the selected account.
//
// Returns one page. Use ListPoliciesAll to walk every page.
func (c *Client) ListPolicies(ctx context.Context, params *ListPoliciesParams, opts ...basaltic.RequestOption) (*basaltic.Page[Policy], error) {
	op := &basaltic.Operation{
		ID:     "listPolicies",
		Method: "GET",
		Path:   "/v1/policies",
	}
	op.Query = params.query()
	var out struct {
		Items []Policy `json:"policies"`
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
	page := &basaltic.Page[Policy]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListPoliciesAll walks every page of ListPolicies, yielding one item at
// a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListPoliciesAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListPoliciesAll(ctx context.Context, params *ListPoliciesParams, opts ...basaltic.RequestOption) iter.Seq2[Policy, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Policy], error) {
		return c.ListPolicies(ctx, params.withMarker(marker), opts...)
	})
}

// ListPolicyRoles lists roles with policy.
//
// List all roles that have this policy attached.
//
// Returns one page. Use ListPolicyRolesAll to walk every page.
func (c *Client) ListPolicyRoles(ctx context.Context, policyID string, params *ListPolicyRolesParams, opts ...basaltic.RequestOption) (*basaltic.Page[Role], error) {
	op := &basaltic.Operation{
		ID:       "listPolicyRoles",
		Method:   "GET",
		Path:     "/v1/policies/{policy_id}/roles",
		PathArgs: []string{policyID},
	}
	op.Query = params.query()
	var out struct {
		Items []Role `json:"roles"`
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
	page := &basaltic.Page[Role]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListPolicyRolesAll walks every page of ListPolicyRoles, yielding one
// item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListPolicyRolesAll(ctx, policyID, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListPolicyRolesAll(ctx context.Context, policyID string, params *ListPolicyRolesParams, opts ...basaltic.RequestOption) iter.Seq2[Role, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Role], error) {
		return c.ListPolicyRoles(ctx, policyID, params.withMarker(marker), opts...)
	})
}

// ListPolicyServiceAccounts lists service accounts with policy.
//
// List all service accounts that have this policy attached.
//
// Returns one page. Use ListPolicyServiceAccountsAll to walk every page.
func (c *Client) ListPolicyServiceAccounts(ctx context.Context, policyID string, params *ListPolicyServiceAccountsParams, opts ...basaltic.RequestOption) (*basaltic.Page[ServiceAccount], error) {
	op := &basaltic.Operation{
		ID:       "listPolicyServiceAccounts",
		Method:   "GET",
		Path:     "/v1/policies/{policy_id}/service-accounts",
		PathArgs: []string{policyID},
	}
	op.Query = params.query()
	var out struct {
		Items []ServiceAccount `json:"service_accounts"`
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
	page := &basaltic.Page[ServiceAccount]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListPolicyServiceAccountsAll walks every page of
// ListPolicyServiceAccounts, yielding one item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListPolicyServiceAccountsAll(ctx, policyID, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListPolicyServiceAccountsAll(ctx context.Context, policyID string, params *ListPolicyServiceAccountsParams, opts ...basaltic.RequestOption) iter.Seq2[ServiceAccount, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[ServiceAccount], error) {
		return c.ListPolicyServiceAccounts(ctx, policyID, params.withMarker(marker), opts...)
	})
}

// ListRegions lists regions (legacy IAM).
//
// Legacy compatibility endpoint. New clients should use
// catalog.basaltic.sh/v1/regions. This endpoint retains IAM-namespaced
// region CRNs and its existing response shape. List all published
// regions. This endpoint is public and does not require authentication.
// Returns all regions with their availability status.
//
// Because it takes no credentials, requests are rate-limited per client
// IP.
//
// Sends no bearer token: the credentials in the request are the
// authentication.
func (c *Client) ListRegions(ctx context.Context, params *ListRegionsParams, opts ...basaltic.RequestOption) (*ListRegionsResult, error) {
	op := &basaltic.Operation{
		ID:              "listRegions",
		Method:          "GET",
		Path:            "/v1/regions",
		Unauthenticated: true,
	}
	op.Query = params.query()
	var out ListRegionsResult
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListRoleInlinePolicies lists a role's inline policies.
func (c *Client) ListRoleInlinePolicies(ctx context.Context, roleID string, params *ListRoleInlinePoliciesParams, opts ...basaltic.RequestOption) (*basaltic.Page[InlinePolicy], error) {
	op := &basaltic.Operation{
		ID:       "listRoleInlinePolicies",
		Method:   "GET",
		Path:     "/v1/roles/{role_id}/inline-policies",
		PathArgs: []string{roleID},
	}
	op.Query = params.query()
	var out struct {
		Items []InlinePolicy `json:"inline_policies"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[InlinePolicy]{Items: out.Items}
	return page, nil
}

// ListRolePolicies lists role policies.
//
// List all policies attached to a role.
func (c *Client) ListRolePolicies(ctx context.Context, roleID string, params *ListRolePoliciesParams, opts ...basaltic.RequestOption) (*basaltic.Page[Policy], error) {
	op := &basaltic.Operation{
		ID:       "listRolePolicies",
		Method:   "GET",
		Path:     "/v1/roles/{role_id}/policies",
		PathArgs: []string{roleID},
	}
	op.Query = params.query()
	var out struct {
		Items []Policy `json:"policies"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[Policy]{Items: out.Items}
	return page, nil
}

// ListRoles lists roles.
//
// List all roles in the selected account.
//
// Returns one page. Use ListRolesAll to walk every page.
func (c *Client) ListRoles(ctx context.Context, params *ListRolesParams, opts ...basaltic.RequestOption) (*basaltic.Page[Role], error) {
	op := &basaltic.Operation{
		ID:     "listRoles",
		Method: "GET",
		Path:   "/v1/roles",
	}
	op.Query = params.query()
	var out struct {
		Items []Role `json:"roles"`
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
	page := &basaltic.Page[Role]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListRolesAll walks every page of ListRoles, yielding one item at a
// time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListRolesAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListRolesAll(ctx context.Context, params *ListRolesParams, opts ...basaltic.RequestOption) iter.Seq2[Role, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Role], error) {
		return c.ListRoles(ctx, params.withMarker(marker), opts...)
	})
}

// ListSTSSessions lists STS sessions.
//
// List active and recent STS (role assumption) sessions for the selected
// account. Requires `iam:ListSTSSessions` permission.
//
// Returns one page. Use ListSTSSessionsAll to walk every page.
func (c *Client) ListSTSSessions(ctx context.Context, params *ListSTSSessionsParams, opts ...basaltic.RequestOption) (*basaltic.Page[STSSession], error) {
	op := &basaltic.Operation{
		ID:     "listSTSSessions",
		Method: "GET",
		Path:   "/v1/sts-sessions",
	}
	op.Query = params.query()
	var out struct {
		Items []STSSession `json:"sts_sessions"`
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
	page := &basaltic.Page[STSSession]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListSTSSessionsAll walks every page of ListSTSSessions, yielding one
// item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListSTSSessionsAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListSTSSessionsAll(ctx context.Context, params *ListSTSSessionsParams, opts ...basaltic.RequestOption) iter.Seq2[STSSession, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[STSSession], error) {
		return c.ListSTSSessions(ctx, params.withMarker(marker), opts...)
	})
}

// ListServiceAccountCredentials lists credentials.
//
// List all credentials for a service account.
func (c *Client) ListServiceAccountCredentials(ctx context.Context, serviceAccountID string, params *ListServiceAccountCredentialsParams, opts ...basaltic.RequestOption) (*basaltic.Page[Credential], error) {
	op := &basaltic.Operation{
		ID:       "listServiceAccountCredentials",
		Method:   "GET",
		Path:     "/v1/service-accounts/{service_account_id}/credentials",
		PathArgs: []string{serviceAccountID},
	}
	op.Query = params.query()
	var out struct {
		Items []Credential `json:"credentials"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[Credential]{Items: out.Items}
	return page, nil
}

// ListServiceAccountInlinePolicies lists a service account's inline policies.
func (c *Client) ListServiceAccountInlinePolicies(ctx context.Context, serviceAccountID string, params *ListServiceAccountInlinePoliciesParams, opts ...basaltic.RequestOption) (*basaltic.Page[InlinePolicy], error) {
	op := &basaltic.Operation{
		ID:       "listServiceAccountInlinePolicies",
		Method:   "GET",
		Path:     "/v1/service-accounts/{service_account_id}/inline-policies",
		PathArgs: []string{serviceAccountID},
	}
	op.Query = params.query()
	var out struct {
		Items []InlinePolicy `json:"inline_policies"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[InlinePolicy]{Items: out.Items}
	return page, nil
}

// ListServiceAccountPolicies lists service account policies.
//
// List all policies attached to a service account.
func (c *Client) ListServiceAccountPolicies(ctx context.Context, serviceAccountID string, params *ListServiceAccountPoliciesParams, opts ...basaltic.RequestOption) (*basaltic.Page[Policy], error) {
	op := &basaltic.Operation{
		ID:       "listServiceAccountPolicies",
		Method:   "GET",
		Path:     "/v1/service-accounts/{service_account_id}/policies",
		PathArgs: []string{serviceAccountID},
	}
	op.Query = params.query()
	var out struct {
		Items []Policy `json:"policies"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[Policy]{Items: out.Items}
	return page, nil
}

// ListServiceAccountSSHKeys lists service-account SSH keys.
//
// Requires the selected account and ordinary IAM authorization against
// the service account, including tag conditions and explicit denies. A
// key never carries a role or grants access to a VM by itself. Linux
// names and numeric IDs are allocated by the platform and do not change
// on key rotation or display-name changes.
func (c *Client) ListServiceAccountSSHKeys(ctx context.Context, serviceAccountID string, opts ...basaltic.RequestOption) ([]*SSHKey, error) {
	op := &basaltic.Operation{
		ID:       "listServiceAccountSSHKeys",
		Method:   "GET",
		Path:     "/v1/service-accounts/{service_account_id}/ssh-keys",
		PathArgs: []string{serviceAccountID},
	}
	var out struct {
		SSHKeys []*SSHKey `json:"ssh_keys"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.SSHKeys, nil
}

// ListServiceAccounts lists service accounts.
//
// List service accounts owned by the account selected via X-Account-Id
// within the selected account.
//
// Returns one page. Use ListServiceAccountsAll to walk every page.
func (c *Client) ListServiceAccounts(ctx context.Context, params *ListServiceAccountsParams, opts ...basaltic.RequestOption) (*basaltic.Page[ServiceAccount], error) {
	op := &basaltic.Operation{
		ID:     "listServiceAccounts",
		Method: "GET",
		Path:   "/v1/service-accounts",
	}
	op.Query = params.query()
	var out struct {
		Items []ServiceAccount `json:"service_accounts"`
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
	page := &basaltic.Page[ServiceAccount]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListServiceAccountsAll walks every page of ListServiceAccounts,
// yielding one item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListServiceAccountsAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListServiceAccountsAll(ctx context.Context, params *ListServiceAccountsParams, opts ...basaltic.RequestOption) iter.Seq2[ServiceAccount, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[ServiceAccount], error) {
		return c.ListServiceAccounts(ctx, params.withMarker(marker), opts...)
	})
}

// PutRoleInlinePolicy creates or replace a role's inline policy.
func (c *Client) PutRoleInlinePolicy(ctx context.Context, roleID string, policyName string, body *PutInlinePolicyRequest, opts ...basaltic.RequestOption) (*InlinePolicy, error) {
	op := &basaltic.Operation{
		ID:       "putRoleInlinePolicy",
		Method:   "PUT",
		Path:     "/v1/roles/{role_id}/inline-policies/{policy_name}",
		PathArgs: []string{roleID, policyName},
		Body:     body,
	}
	var out struct {
		InlinePolicy *InlinePolicy `json:"inline_policy"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.InlinePolicy, nil
}

// PutServiceAccountInlinePolicy creates or replace a service account's inline policy.
func (c *Client) PutServiceAccountInlinePolicy(ctx context.Context, serviceAccountID string, policyName string, body *PutInlinePolicyRequest, opts ...basaltic.RequestOption) (*InlinePolicy, error) {
	op := &basaltic.Operation{
		ID:       "putServiceAccountInlinePolicy",
		Method:   "PUT",
		Path:     "/v1/service-accounts/{service_account_id}/inline-policies/{policy_name}",
		PathArgs: []string{serviceAccountID, policyName},
		Body:     body,
	}
	var out struct {
		InlinePolicy *InlinePolicy `json:"inline_policy"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.InlinePolicy, nil
}

// RemoveRolePermissionBoundary removes a role's permission boundary.
func (c *Client) RemoveRolePermissionBoundary(ctx context.Context, roleID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "removeRolePermissionBoundary",
		Method:   "DELETE",
		Path:     "/v1/roles/{role_id}/permission-boundary",
		PathArgs: []string{roleID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// RemoveServiceAccountPermissionBoundary removes a service account's permission boundary.
func (c *Client) RemoveServiceAccountPermissionBoundary(ctx context.Context, serviceAccountID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "removeServiceAccountPermissionBoundary",
		Method:   "DELETE",
		Path:     "/v1/service-accounts/{service_account_id}/permission-boundary",
		PathArgs: []string{serviceAccountID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// RevokeOAuthToken revokes a bearer token.
//
// Revoke an access token, ending the session behind it. Every credential
// that session issued stops working on the next request, rather than at
// the token's expiry.
//
// Answers `200` whether or not anything was revoked — an unknown,
// expired or already-revoked token is not distinguished from a live one.
// That is required by RFC 7009 and it is the point: an endpoint that
// reported the difference would tell anyone holding a token whether it
// is still good.
//
// A token belonging to another organization is silently ignored for the
// same reason.
func (c *Client) RevokeOAuthToken(ctx context.Context, body *OAuthRevokeRequest, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:     "revokeOAuthToken",
		Method: "POST",
		Path:   "/v1/oauth/revoke",
		Body:   body,
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// RevokeSTSSession revokes STS session.
//
// Revoke an active STS session, invalidating its credentials.
//
// Needs **two** actions, not one. `iam:RevokeSession` decides whether
// the revocation happens; the response then re-reads the session, which
// checks `iam:GetSTSSession`. A caller holding only the first revokes
// the session successfully and still receives a `403` — the
// credentials are already dead at that point.
func (c *Client) RevokeSTSSession(ctx context.Context, sessionID string, body *RevokeSTSSessionRequest, opts ...basaltic.RequestOption) (*STSSession, error) {
	op := &basaltic.Operation{
		ID:       "revokeSTSSession",
		Method:   "DELETE",
		Path:     "/v1/sts-sessions/{session_id}",
		PathArgs: []string{sessionID},
		Body:     body,
	}
	var out struct {
		STSSession *STSSession `json:"sts_session"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.STSSession, nil
}

// SetRolePermissionBoundary sets a role's permission boundary.
//
// Set the permission boundary — the boundary caps the principal's
// effective permissions to the intersection with its identity policies.
func (c *Client) SetRolePermissionBoundary(ctx context.Context, roleID string, body *SetBoundaryRequest, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "setRolePermissionBoundary",
		Method:   "PUT",
		Path:     "/v1/roles/{role_id}/permission-boundary",
		PathArgs: []string{roleID},
		Body:     body,
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// SetServiceAccountPermissionBoundary sets a service account's permission boundary.
//
// Set the permission boundary — the boundary caps the principal's
// effective permissions to the intersection with its identity policies.
func (c *Client) SetServiceAccountPermissionBoundary(ctx context.Context, serviceAccountID string, body *SetBoundaryRequest, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "setServiceAccountPermissionBoundary",
		Method:   "PUT",
		Path:     "/v1/service-accounts/{service_account_id}/permission-boundary",
		PathArgs: []string{serviceAccountID},
		Body:     body,
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// UpdatePolicy updates policy.
//
// Update an existing policy.
func (c *Client) UpdatePolicy(ctx context.Context, policyID string, body *PolicyUpdateRequest, opts ...basaltic.RequestOption) (*Policy, error) {
	op := &basaltic.Operation{
		ID:       "updatePolicy",
		Method:   "PATCH",
		Path:     "/v1/policies/{policy_id}",
		PathArgs: []string{policyID},
		Body:     body,
	}
	var out struct {
		Policy *Policy `json:"policy"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Policy, nil
}

// UpdateRole updates role.
//
// Update an existing role.
func (c *Client) UpdateRole(ctx context.Context, roleID string, body *RoleUpdateRequest, opts ...basaltic.RequestOption) (*Role, error) {
	op := &basaltic.Operation{
		ID:       "updateRole",
		Method:   "PATCH",
		Path:     "/v1/roles/{role_id}",
		PathArgs: []string{roleID},
		Body:     body,
	}
	var out struct {
		Role *Role `json:"role"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Role, nil
}

// UpdateServiceAccount updates service account.
//
// Update a service account.
func (c *Client) UpdateServiceAccount(ctx context.Context, serviceAccountID string, body *ServiceAccountUpdateRequest, opts ...basaltic.RequestOption) (*ServiceAccount, error) {
	op := &basaltic.Operation{
		ID:       "updateServiceAccount",
		Method:   "PATCH",
		Path:     "/v1/service-accounts/{service_account_id}",
		PathArgs: []string{serviceAccountID},
		Body:     body,
	}
	var out struct {
		ServiceAccount *ServiceAccount `json:"service_account"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.ServiceAccount, nil
}

// GetPolicyByReference fetches one policy by an id, a CRN or a name.
//
// The reference is classified by its syntax alone, exactly as the
// platform does (see [basaltic.ParseReference]): an id is fetched with
// [Client.GetPolicy]; a CRN or a name goes to [Client.ListPolicies] as
// an exact filter, together with any filters already set on scope, which
// may be nil. A miss is a not-found error for the kind the string was
// read as — no other kind is tried — and more than one match is a
// [basaltic.AmbiguousReferenceError].
func (c *Client) GetPolicyByReference(ctx context.Context, ref string, scope *ListPoliciesParams, opts ...basaltic.RequestOption) (*Policy, error) {
	return basaltic.ResolveByReference(ctx, ref, "policy", "listPolicies", true,
		func(ctx context.Context, refID string) (*Policy, error) {
			return c.GetPolicy(ctx, refID, opts...)
		},
		func(ctx context.Context, refName, refCRN string) (*basaltic.Page[Policy], error) {
			var p ListPoliciesParams
			if scope != nil {
				p = *scope
			}
			p.Name = refName
			p.CRN = refCRN
			p.Limit = 2
			return c.ListPolicies(ctx, &p, opts...)
		})
}

// GetRoleByReference fetches one role by an id, a CRN or a name.
//
// The reference is classified by its syntax alone, exactly as the
// platform does (see [basaltic.ParseReference]): an id is fetched with
// [Client.GetRole]; a CRN or a name goes to [Client.ListRoles] as an
// exact filter, together with any filters already set on scope, which
// may be nil. A miss is a not-found error for the kind the string was
// read as — no other kind is tried — and more than one match is a
// [basaltic.AmbiguousReferenceError].
func (c *Client) GetRoleByReference(ctx context.Context, ref string, scope *ListRolesParams, opts ...basaltic.RequestOption) (*Role, error) {
	return basaltic.ResolveByReference(ctx, ref, "role", "listRoles", true,
		func(ctx context.Context, refID string) (*Role, error) {
			return c.GetRole(ctx, refID, opts...)
		},
		func(ctx context.Context, refName, refCRN string) (*basaltic.Page[Role], error) {
			var p ListRolesParams
			if scope != nil {
				p = *scope
			}
			p.Name = refName
			p.CRN = refCRN
			p.Limit = 2
			return c.ListRoles(ctx, &p, opts...)
		})
}

// GetSTSSessionByReference fetches one sts session by an id, a CRN or a
// name.
//
// The reference is classified by its syntax alone, exactly as the
// platform does (see [basaltic.ParseReference]): an id is fetched with
// [Client.GetSTSSession]; a CRN or a name goes to
// [Client.ListSTSSessions] as an exact filter, together with any filters
// already set on scope, which may be nil. A miss is a not-found error
// for the kind the string was read as — no other kind is tried — and
// more than one match is a [basaltic.AmbiguousReferenceError].
//
// A name is unique only within its parent; fix it on scope (Principal,
// Role) or the lookup can match more than one.
func (c *Client) GetSTSSessionByReference(ctx context.Context, ref string, scope *ListSTSSessionsParams, opts ...basaltic.RequestOption) (*STSSession, error) {
	return basaltic.ResolveByReference(ctx, ref, "sts-session", "listSTSSessions", true,
		func(ctx context.Context, refID string) (*STSSession, error) {
			return c.GetSTSSession(ctx, refID, opts...)
		},
		func(ctx context.Context, refName, refCRN string) (*basaltic.Page[STSSession], error) {
			var p ListSTSSessionsParams
			if scope != nil {
				p = *scope
			}
			p.Name = refName
			p.CRN = refCRN
			p.Limit = 2
			return c.ListSTSSessions(ctx, &p, opts...)
		})
}

// GetServiceAccountByReference fetches one service account by an id, a
// CRN or a name.
//
// The reference is classified by its syntax alone, exactly as the
// platform does (see [basaltic.ParseReference]): an id is fetched with
// [Client.GetServiceAccount]; a CRN or a name goes to
// [Client.ListServiceAccounts] as an exact filter, together with any
// filters already set on scope, which may be nil. A miss is a not-found
// error for the kind the string was read as — no other kind is tried
// — and more than one match is a [basaltic.AmbiguousReferenceError].
func (c *Client) GetServiceAccountByReference(ctx context.Context, ref string, scope *ListServiceAccountsParams, opts ...basaltic.RequestOption) (*ServiceAccount, error) {
	return basaltic.ResolveByReference(ctx, ref, "service-account", "listServiceAccounts", true,
		func(ctx context.Context, refID string) (*ServiceAccount, error) {
			return c.GetServiceAccount(ctx, refID, opts...)
		},
		func(ctx context.Context, refName, refCRN string) (*basaltic.Page[ServiceAccount], error) {
			var p ListServiceAccountsParams
			if scope != nil {
				p = *scope
			}
			p.Name = refName
			p.CRN = refCRN
			p.Limit = 2
			return c.ListServiceAccounts(ctx, &p, opts...)
		})
}
