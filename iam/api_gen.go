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

// ListAccountsParams are the optional filters and pagination controls for
// [Client.ListAccounts]. A nil *ListAccountsParams sends none of them.
type ListAccountsParams struct {
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
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListAccountsParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
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
func (p *ListAccountsParams) withMarker(marker string) *ListAccountsParams {
	var out ListAccountsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListGroupServiceAccountsParams are the optional filters and pagination controls for
// [Client.ListGroupServiceAccounts]. A nil *ListGroupServiceAccountsParams sends none of them.
type ListGroupServiceAccountsParams struct {
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
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListGroupServiceAccountsParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
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
func (p *ListGroupServiceAccountsParams) withMarker(marker string) *ListGroupServiceAccountsParams {
	var out ListGroupServiceAccountsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListGroupUsersParams are the optional filters and pagination controls for
// [Client.ListGroupUsers]. A nil *ListGroupUsersParams sends none of them.
type ListGroupUsersParams struct {
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
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListGroupUsersParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
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
func (p *ListGroupUsersParams) withMarker(marker string) *ListGroupUsersParams {
	var out ListGroupUsersParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListGroupsParams are the optional filters and pagination controls for
// [Client.ListGroups]. A nil *ListGroupsParams sends none of them.
type ListGroupsParams struct {
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

	// Name filter by name (exact match or prefix with *)
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListGroupsParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
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
func (p *ListGroupsParams) withMarker(marker string) *ListGroupsParams {
	var out ListGroupsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListInvitationsParams are the optional filters and pagination controls for
// [Client.ListInvitations]. A nil *ListInvitationsParams sends none of them.
type ListInvitationsParams struct {
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
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListInvitationsParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
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
func (p *ListInvitationsParams) withMarker(marker string) *ListInvitationsParams {
	var out ListInvitationsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListOrganizationsParams are the optional filters and pagination controls for
// [Client.ListOrganizations]. A nil *ListOrganizationsParams sends none of them.
type ListOrganizationsParams struct {
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
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListOrganizationsParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
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
func (p *ListOrganizationsParams) withMarker(marker string) *ListOrganizationsParams {
	var out ListOrganizationsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListPoliciesParams are the optional filters and pagination controls for
// [Client.ListPolicies]. A nil *ListPoliciesParams sends none of them.
type ListPoliciesParams struct {
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

	// Name filter by name (exact match or prefix with *)
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListPoliciesParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
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

// ListPolicyGroupsParams are the optional filters and pagination controls for
// [Client.ListPolicyGroups]. A nil *ListPolicyGroupsParams sends none of them.
type ListPolicyGroupsParams struct {
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
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListPolicyGroupsParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
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
func (p *ListPolicyGroupsParams) withMarker(marker string) *ListPolicyGroupsParams {
	var out ListPolicyGroupsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListPolicyRolesParams are the optional filters and pagination controls for
// [Client.ListPolicyRoles]. A nil *ListPolicyRolesParams sends none of them.
type ListPolicyRolesParams struct {
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
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListPolicyRolesParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
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
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListPolicyServiceAccountsParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
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
func (p *ListPolicyServiceAccountsParams) withMarker(marker string) *ListPolicyServiceAccountsParams {
	var out ListPolicyServiceAccountsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListPolicyUsersParams are the optional filters and pagination controls for
// [Client.ListPolicyUsers]. A nil *ListPolicyUsersParams sends none of them.
type ListPolicyUsersParams struct {
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
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListPolicyUsersParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
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
func (p *ListPolicyUsersParams) withMarker(marker string) *ListPolicyUsersParams {
	var out ListPolicyUsersParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListRolesParams are the optional filters and pagination controls for
// [Client.ListRoles]. A nil *ListRolesParams sends none of them.
type ListRolesParams struct {
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

	// Name filter by name (exact match or prefix with *)
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListRolesParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
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

	// PrincipalID filter by principal ID (user or service account)
	PrincipalID string

	// PrincipalType filter by principal type. `assumed_role` selects the sessions role
	// chaining produces, where an existing assumed-role session assumed
	// another role.
	//
	//
	// One of: "user", "service_account", "assumed_role".
	PrincipalType string

	// RoleID filter by role ID
	RoleID string
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
	if p.Limit != 0 {
		q.Set("limit", strconv.Itoa(int(p.Limit)))
	}
	if p.Marker != "" {
		q.Set("marker", p.Marker)
	}
	if p.PrincipalID != "" {
		q.Set("principal_id", p.PrincipalID)
	}
	if p.PrincipalType != "" {
		q.Set("principal_type", p.PrincipalType)
	}
	if p.RoleID != "" {
		q.Set("role_id", p.RoleID)
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

// ListServiceAccountsParams are the optional filters and pagination controls for
// [Client.ListServiceAccounts]. A nil *ListServiceAccountsParams sends none of them.
type ListServiceAccountsParams struct {
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

	// Name filter by name (exact match or prefix with *)
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListServiceAccountsParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
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

// ListUsersParams are the optional filters and pagination controls for
// [Client.ListUsers]. A nil *ListUsersParams sends none of them.
type ListUsersParams struct {
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

	// Name filter by name (exact match or prefix with *)
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListUsersParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
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
func (p *ListUsersParams) withMarker(marker string) *ListUsersParams {
	var out ListUsersParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// AddServiceAccountToGroup adds service account to group.
//
// Add a service account to a group.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) AddServiceAccountToGroup(ctx context.Context, serviceAccountID string, body *ServiceAccountGroupAddRequest, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "addServiceAccountToGroup",
		Method:   "POST",
		Path:     "/v1/service-accounts/{service_account_id}/groups",
		PathArgs: []string{serviceAccountID},
		Body:     body,
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// AddUser adds user to organization.
//
// Invite a user to the organization by email. An invitation email is
// always sent and the user joins on accepting it — there is no path
// that adds someone without their consent, even when they already have a
// platform account. Inviting an existing member, or someone who already
// has a pending invitation, is rejected with 409.
//
// Optionally specify groups to add the user to on acceptance.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) AddUser(ctx context.Context, body *UserAddRequest, opts ...basaltic.RequestOption) (*UserAddResponse, error) {
	op := &basaltic.Operation{
		ID:     "addUser",
		Method: "POST",
		Path:   "/v1/users",
		Body:   body,
	}
	var out UserAddResponse
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// AddUserToGroup adds user to group.
//
// Add a user to a group.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) AddUserToGroup(ctx context.Context, userID string, body *UserGroupAddRequest, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "addUserToGroup",
		Method:   "POST",
		Path:     "/v1/users/{user_id}/groups",
		PathArgs: []string{userID},
		Body:     body,
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// AssumeRole assumes role.
//
// Assume an IAM role and receive temporary credentials. The caller must
// be allowed by the role's trust policy. Can be used by service accounts
// to get temporary credentials for a specific role.
//
// An optional `policy` scopes the session down to less than the role
// grants; see `SessionPolicyDocument`.
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
// `role_id` and `account_id` are read from the body like every other
// field.
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
// calling; it grants nothing. The role named in `role_id` is assumable
// only if its own trust policy admits this identity. Its `principals`
// must name the federation provider, written
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
// The credentials come back scoped to `account_id`, carrying the role's
// own permissions. There is no `policy` field here — unlike `POST
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

// AttachGroupPolicy attaches policy to group.
//
// Attach a policy to a group. All group members inherit this policy.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) AttachGroupPolicy(ctx context.Context, groupID string, body *PolicyAttachRequest, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "attachGroupPolicy",
		Method:   "POST",
		Path:     "/v1/groups/{group_id}/policies",
		PathArgs: []string{groupID},
		Body:     body,
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
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

// AttachUserPolicy attaches policy to user.
//
// Attach a policy to a user.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) AttachUserPolicy(ctx context.Context, userID string, body *PolicyAttachRequest, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "attachUserPolicy",
		Method:   "POST",
		Path:     "/v1/users/{user_id}/policies",
		PathArgs: []string{userID},
		Body:     body,
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// AuthorizeOAuthClient — Approve a CLI login and issue an authorization code.
//
// Approve a client to act as you, and receive the redirect that hands it
// an authorization code.
//
// **This is the console's endpoint, not a client's.** It is called by
// the Basaltic console's consent page on behalf of a signed-in user; the
// CLI never calls it. A CLI opens a browser at that page, and the page
// calls this. Anything driving it directly would need the user's console
// session, at which point it already has everything the code would
// grant.
//
// It is the half of the authorization-code flow that establishes WHO is
// approving. The user must already be signed in — including any second
// factor — and must be a member of the organization named in
// `organization_id`. The organization is explicit rather than inferred:
// a person in several has no single obvious answer, and choosing one for
// them would scope the resulting token to something they did not pick.
//
// Unlike the token endpoint, this answers in the usual API envelope. It
// is not part of the surface a third-party OAuth client talks to, so it
// follows the caller — and the caller is our own front end.
func (c *Client) AuthorizeOAuthClient(ctx context.Context, body *OAuthAuthorizeRequest, opts ...basaltic.RequestOption) (string, error) {
	op := &basaltic.Operation{
		ID:     "authorizeOAuthClient",
		Method: "POST",
		Path:   "/v1/oauth/authorize",
		Body:   body,
	}
	var out struct {
		RedirectTo string `json:"redirect_to"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return "", err
	}
	return out.RedirectTo, nil
}

// CancelInvitation cancels invitation.
//
// Cancel a pending invitation.
func (c *Client) CancelInvitation(ctx context.Context, invitationID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "cancelInvitation",
		Method:   "DELETE",
		Path:     "/v1/invitations/{invitation_id}",
		PathArgs: []string{invitationID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// CreateAccount creates account.
//
// Create a new account in the current organization. The caller chooses
// the `handle`, which is the account's public identifier and immutable
// once created; the platform assigns the internal `id`.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateAccount(ctx context.Context, body *CreateAccountRequest, opts ...basaltic.RequestOption) (*Account, error) {
	op := &basaltic.Operation{
		ID:     "createAccount",
		Method: "POST",
		Path:   "/v1/accounts",
		Body:   body,
	}
	var out struct {
		Account *Account `json:"account"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Account, nil
}

// CreateGroup creates group.
//
// Create a new group in the current organization.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateGroup(ctx context.Context, body *GroupCreateRequest, opts ...basaltic.RequestOption) (*Group, error) {
	op := &basaltic.Operation{
		ID:     "createGroup",
		Method: "POST",
		Path:   "/v1/groups",
		Body:   body,
	}
	var out struct {
		Group *Group `json:"group"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Group, nil
}

// CreatePolicy creates policy.
//
// Create a new policy in the current organization.
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
// Create a new role in the current organization.
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

// DeleteAccount deletes account.
//
// Account must not own any resources.
func (c *Client) DeleteAccount(ctx context.Context, accountID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteAccount",
		Method:   "DELETE",
		Path:     "/v1/accounts/{account_id}",
		PathArgs: []string{accountID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteGroup deletes group.
//
// Delete a group.
func (c *Client) DeleteGroup(ctx context.Context, groupID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteGroup",
		Method:   "DELETE",
		Path:     "/v1/groups/{group_id}",
		PathArgs: []string{groupID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteGroupInlinePolicy deletes a group's inline policy by name.
func (c *Client) DeleteGroupInlinePolicy(ctx context.Context, groupID string, policyName string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteGroupInlinePolicy",
		Method:   "DELETE",
		Path:     "/v1/groups/{group_id}/inline-policies/{policy_name}",
		PathArgs: []string{groupID, policyName},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteOrganization deletes organization.
//
// Delete an organization. This action is irreversible and will delete
// all resources associated with the organization. Only the organization
// owner can perform this action.
func (c *Client) DeleteOrganization(ctx context.Context, organizationID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteOrganization",
		Method:   "DELETE",
		Path:     "/v1/organizations/{organization_id}",
		PathArgs: []string{organizationID},
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

// DeleteUserInlinePolicy deletes a user's inline policy by name.
func (c *Client) DeleteUserInlinePolicy(ctx context.Context, userID string, policyName string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteUserInlinePolicy",
		Method:   "DELETE",
		Path:     "/v1/users/{user_id}/inline-policies/{policy_name}",
		PathArgs: []string{userID, policyName},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DetachGroupPolicy detaches policy from group.
//
// Detach a policy from a group.
func (c *Client) DetachGroupPolicy(ctx context.Context, groupID string, policyID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "detachGroupPolicy",
		Method:   "DELETE",
		Path:     "/v1/groups/{group_id}/policies/{policy_id}",
		PathArgs: []string{groupID, policyID},
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

// DetachUserPolicy detaches policy from user.
//
// Detach a policy from a user.
func (c *Client) DetachUserPolicy(ctx context.Context, userID string, policyID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "detachUserPolicy",
		Method:   "DELETE",
		Path:     "/v1/users/{user_id}/policies/{policy_id}",
		PathArgs: []string{userID, policyID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// GetAccount gets account.
func (c *Client) GetAccount(ctx context.Context, accountID string, opts ...basaltic.RequestOption) (*Account, error) {
	op := &basaltic.Operation{
		ID:       "getAccount",
		Method:   "GET",
		Path:     "/v1/accounts/{account_id}",
		PathArgs: []string{accountID},
	}
	var out struct {
		Account *Account `json:"account"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Account, nil
}

// GetGroup gets group.
//
// Get details of a specific group.
func (c *Client) GetGroup(ctx context.Context, groupID string, opts ...basaltic.RequestOption) (*Group, error) {
	op := &basaltic.Operation{
		ID:       "getGroup",
		Method:   "GET",
		Path:     "/v1/groups/{group_id}",
		PathArgs: []string{groupID},
	}
	var out struct {
		Group *Group `json:"group"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Group, nil
}

// GetGroupInlinePolicy gets a group's inline policy by name.
func (c *Client) GetGroupInlinePolicy(ctx context.Context, groupID string, policyName string, opts ...basaltic.RequestOption) (*InlinePolicy, error) {
	op := &basaltic.Operation{
		ID:       "getGroupInlinePolicy",
		Method:   "GET",
		Path:     "/v1/groups/{group_id}/inline-policies/{policy_name}",
		PathArgs: []string{groupID, policyName},
	}
	var out struct {
		InlinePolicy *InlinePolicy `json:"inline_policy"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.InlinePolicy, nil
}

// GetInvitation gets invitation.
//
// Get details of a specific invitation.
func (c *Client) GetInvitation(ctx context.Context, invitationID string, opts ...basaltic.RequestOption) (*Invitation, error) {
	op := &basaltic.Operation{
		ID:       "getInvitation",
		Method:   "GET",
		Path:     "/v1/invitations/{invitation_id}",
		PathArgs: []string{invitationID},
	}
	var out struct {
		Invitation *Invitation `json:"invitation"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Invitation, nil
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

// GetOrganization gets organization.
//
// Get details of a specific organization.
func (c *Client) GetOrganization(ctx context.Context, organizationID string, opts ...basaltic.RequestOption) (*Organization, error) {
	op := &basaltic.Operation{
		ID:       "getOrganization",
		Method:   "GET",
		Path:     "/v1/organizations/{organization_id}",
		PathArgs: []string{organizationID},
	}
	var out struct {
		Organization *Organization `json:"organization"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Organization, nil
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

// GetUser gets user.
//
// Get details of a specific user in the organization.
func (c *Client) GetUser(ctx context.Context, userID string, opts ...basaltic.RequestOption) (*User, error) {
	op := &basaltic.Operation{
		ID:       "getUser",
		Method:   "GET",
		Path:     "/v1/users/{user_id}",
		PathArgs: []string{userID},
	}
	var out struct {
		User *User `json:"user"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.User, nil
}

// GetUserInlinePolicy gets a user's inline policy by name.
func (c *Client) GetUserInlinePolicy(ctx context.Context, userID string, policyName string, opts ...basaltic.RequestOption) (*InlinePolicy, error) {
	op := &basaltic.Operation{
		ID:       "getUserInlinePolicy",
		Method:   "GET",
		Path:     "/v1/users/{user_id}/inline-policies/{policy_name}",
		PathArgs: []string{userID, policyName},
	}
	var out struct {
		InlinePolicy *InlinePolicy `json:"inline_policy"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.InlinePolicy, nil
}

// GetUserPermissionBoundary gets a user's permission boundary.
func (c *Client) GetUserPermissionBoundary(ctx context.Context, userID string, opts ...basaltic.RequestOption) (*PermissionBoundary, error) {
	op := &basaltic.Operation{
		ID:       "getUserPermissionBoundary",
		Method:   "GET",
		Path:     "/v1/users/{user_id}/permission-boundary",
		PathArgs: []string{userID},
	}
	var out struct {
		PermissionBoundary *PermissionBoundary `json:"permission_boundary"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.PermissionBoundary, nil
}

// ListAccounts lists accounts.
//
// List accounts in the current organization.
//
// Returns one page. Use ListAccountsAll to walk every page.
func (c *Client) ListAccounts(ctx context.Context, params *ListAccountsParams, opts ...basaltic.RequestOption) (*basaltic.Page[Account], error) {
	op := &basaltic.Operation{
		ID:     "listAccounts",
		Method: "GET",
		Path:   "/v1/accounts",
	}
	op.Query = params.query()
	var out struct {
		Items []Account `json:"accounts"`
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
	page := &basaltic.Page[Account]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListAccountsAll walks every page of ListAccounts, yielding one item at
// a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListAccountsAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListAccountsAll(ctx context.Context, params *ListAccountsParams, opts ...basaltic.RequestOption) iter.Seq2[Account, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Account], error) {
		return c.ListAccounts(ctx, params.withMarker(marker), opts...)
	})
}

// ListGroupInlinePolicies lists a group's inline policies.
func (c *Client) ListGroupInlinePolicies(ctx context.Context, groupID string, opts ...basaltic.RequestOption) (*basaltic.Page[InlinePolicy], error) {
	op := &basaltic.Operation{
		ID:       "listGroupInlinePolicies",
		Method:   "GET",
		Path:     "/v1/groups/{group_id}/inline-policies",
		PathArgs: []string{groupID},
	}
	var out struct {
		Items []InlinePolicy `json:"inline_policies"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[InlinePolicy]{Items: out.Items}
	return page, nil
}

// ListGroupPolicies lists group policies.
//
// List all policies attached to a group.
func (c *Client) ListGroupPolicies(ctx context.Context, groupID string, opts ...basaltic.RequestOption) (*basaltic.Page[Policy], error) {
	op := &basaltic.Operation{
		ID:       "listGroupPolicies",
		Method:   "GET",
		Path:     "/v1/groups/{group_id}/policies",
		PathArgs: []string{groupID},
	}
	var out struct {
		Items []Policy `json:"policies"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[Policy]{Items: out.Items}
	return page, nil
}

// ListGroupServiceAccounts lists group service accounts.
//
// List all service accounts in a group.
//
// Returns one page. Use ListGroupServiceAccountsAll to walk every page.
func (c *Client) ListGroupServiceAccounts(ctx context.Context, groupID string, params *ListGroupServiceAccountsParams, opts ...basaltic.RequestOption) (*basaltic.Page[GroupServiceAccount], error) {
	op := &basaltic.Operation{
		ID:       "listGroupServiceAccounts",
		Method:   "GET",
		Path:     "/v1/groups/{group_id}/service-accounts",
		PathArgs: []string{groupID},
	}
	op.Query = params.query()
	var out struct {
		Items []GroupServiceAccount `json:"service_accounts"`
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
	page := &basaltic.Page[GroupServiceAccount]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListGroupServiceAccountsAll walks every page of
// ListGroupServiceAccounts, yielding one item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListGroupServiceAccountsAll(ctx, groupID, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListGroupServiceAccountsAll(ctx context.Context, groupID string, params *ListGroupServiceAccountsParams, opts ...basaltic.RequestOption) iter.Seq2[GroupServiceAccount, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[GroupServiceAccount], error) {
		return c.ListGroupServiceAccounts(ctx, groupID, params.withMarker(marker), opts...)
	})
}

// ListGroupUsers lists group users.
//
// List all users in a group.
//
// Returns one page. Use ListGroupUsersAll to walk every page.
func (c *Client) ListGroupUsers(ctx context.Context, groupID string, params *ListGroupUsersParams, opts ...basaltic.RequestOption) (*basaltic.Page[GroupUser], error) {
	op := &basaltic.Operation{
		ID:       "listGroupUsers",
		Method:   "GET",
		Path:     "/v1/groups/{group_id}/users",
		PathArgs: []string{groupID},
	}
	op.Query = params.query()
	var out struct {
		Items []GroupUser `json:"users"`
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
	page := &basaltic.Page[GroupUser]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListGroupUsersAll walks every page of ListGroupUsers, yielding one
// item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListGroupUsersAll(ctx, groupID, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListGroupUsersAll(ctx context.Context, groupID string, params *ListGroupUsersParams, opts ...basaltic.RequestOption) iter.Seq2[GroupUser, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[GroupUser], error) {
		return c.ListGroupUsers(ctx, groupID, params.withMarker(marker), opts...)
	})
}

// ListGroups lists groups.
//
// List all groups in the current organization.
//
// Returns one page. Use ListGroupsAll to walk every page.
func (c *Client) ListGroups(ctx context.Context, params *ListGroupsParams, opts ...basaltic.RequestOption) (*basaltic.Page[Group], error) {
	op := &basaltic.Operation{
		ID:     "listGroups",
		Method: "GET",
		Path:   "/v1/groups",
	}
	op.Query = params.query()
	var out struct {
		Items []Group `json:"groups"`
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
	page := &basaltic.Page[Group]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListGroupsAll walks every page of ListGroups, yielding one item at a
// time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListGroupsAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListGroupsAll(ctx context.Context, params *ListGroupsParams, opts ...basaltic.RequestOption) iter.Seq2[Group, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Group], error) {
		return c.ListGroups(ctx, params.withMarker(marker), opts...)
	})
}

// ListInvitations lists invitations.
//
// List the current organization's invitations in every state —
// pending, accepted, expired and cancelled all come back; filter on
// `status` client-side. Every POST /v1/users creates one.
//
// Returns one page. Use ListInvitationsAll to walk every page.
func (c *Client) ListInvitations(ctx context.Context, params *ListInvitationsParams, opts ...basaltic.RequestOption) (*basaltic.Page[Invitation], error) {
	op := &basaltic.Operation{
		ID:     "listInvitations",
		Method: "GET",
		Path:   "/v1/invitations",
	}
	op.Query = params.query()
	var out struct {
		Items []Invitation `json:"invitations"`
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
	page := &basaltic.Page[Invitation]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListInvitationsAll walks every page of ListInvitations, yielding one
// item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListInvitationsAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListInvitationsAll(ctx context.Context, params *ListInvitationsParams, opts ...basaltic.RequestOption) iter.Seq2[Invitation, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Invitation], error) {
		return c.ListInvitations(ctx, params.withMarker(marker), opts...)
	})
}

// ListOrganizations lists organizations.
//
// List all organizations the current user belongs to.
//
// Returns one page. Use ListOrganizationsAll to walk every page.
func (c *Client) ListOrganizations(ctx context.Context, params *ListOrganizationsParams, opts ...basaltic.RequestOption) (*basaltic.Page[OrganizationWithMembership], error) {
	op := &basaltic.Operation{
		ID:     "listOrganizations",
		Method: "GET",
		Path:   "/v1/organizations",
	}
	op.Query = params.query()
	var out struct {
		Items []OrganizationWithMembership `json:"organizations"`
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
	page := &basaltic.Page[OrganizationWithMembership]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListOrganizationsAll walks every page of ListOrganizations, yielding
// one item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListOrganizationsAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListOrganizationsAll(ctx context.Context, params *ListOrganizationsParams, opts ...basaltic.RequestOption) iter.Seq2[OrganizationWithMembership, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[OrganizationWithMembership], error) {
		return c.ListOrganizations(ctx, params.withMarker(marker), opts...)
	})
}

// ListPolicies lists policies.
//
// List all policies in the current organization.
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

// ListPolicyGroups lists groups with policy.
//
// List all groups that have this policy attached.
//
// Returns one page. Use ListPolicyGroupsAll to walk every page.
func (c *Client) ListPolicyGroups(ctx context.Context, policyID string, params *ListPolicyGroupsParams, opts ...basaltic.RequestOption) (*basaltic.Page[Group], error) {
	op := &basaltic.Operation{
		ID:       "listPolicyGroups",
		Method:   "GET",
		Path:     "/v1/policies/{policy_id}/groups",
		PathArgs: []string{policyID},
	}
	op.Query = params.query()
	var out struct {
		Items []Group `json:"groups"`
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
	page := &basaltic.Page[Group]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListPolicyGroupsAll walks every page of ListPolicyGroups, yielding one
// item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListPolicyGroupsAll(ctx, policyID, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListPolicyGroupsAll(ctx context.Context, policyID string, params *ListPolicyGroupsParams, opts ...basaltic.RequestOption) iter.Seq2[Group, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Group], error) {
		return c.ListPolicyGroups(ctx, policyID, params.withMarker(marker), opts...)
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

// ListPolicyUsers lists users with policy.
//
// List all users that have this policy attached.
//
// Returns one page. Use ListPolicyUsersAll to walk every page.
func (c *Client) ListPolicyUsers(ctx context.Context, policyID string, params *ListPolicyUsersParams, opts ...basaltic.RequestOption) (*basaltic.Page[User], error) {
	op := &basaltic.Operation{
		ID:       "listPolicyUsers",
		Method:   "GET",
		Path:     "/v1/policies/{policy_id}/users",
		PathArgs: []string{policyID},
	}
	op.Query = params.query()
	var out struct {
		Items []User `json:"users"`
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
	page := &basaltic.Page[User]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListPolicyUsersAll walks every page of ListPolicyUsers, yielding one
// item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListPolicyUsersAll(ctx, policyID, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListPolicyUsersAll(ctx context.Context, policyID string, params *ListPolicyUsersParams, opts ...basaltic.RequestOption) iter.Seq2[User, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[User], error) {
		return c.ListPolicyUsers(ctx, policyID, params.withMarker(marker), opts...)
	})
}

// ListRegions lists regions.
//
// List all available regions. This endpoint is public and does not
// require authentication. Returns all regions with their availability
// status.
//
// Because it takes no credentials, requests are rate-limited per client
// IP.
//
// Sends no bearer token: the credentials in the request are the
// authentication.
func (c *Client) ListRegions(ctx context.Context, opts ...basaltic.RequestOption) (*ListRegionsResult, error) {
	op := &basaltic.Operation{
		ID:              "listRegions",
		Method:          "GET",
		Path:            "/v1/regions",
		Unauthenticated: true,
	}
	var out ListRegionsResult
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListRoleInlinePolicies lists a role's inline policies.
func (c *Client) ListRoleInlinePolicies(ctx context.Context, roleID string, opts ...basaltic.RequestOption) (*basaltic.Page[InlinePolicy], error) {
	op := &basaltic.Operation{
		ID:       "listRoleInlinePolicies",
		Method:   "GET",
		Path:     "/v1/roles/{role_id}/inline-policies",
		PathArgs: []string{roleID},
	}
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
func (c *Client) ListRolePolicies(ctx context.Context, roleID string, opts ...basaltic.RequestOption) (*basaltic.Page[Policy], error) {
	op := &basaltic.Operation{
		ID:       "listRolePolicies",
		Method:   "GET",
		Path:     "/v1/roles/{role_id}/policies",
		PathArgs: []string{roleID},
	}
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
// List all roles in the current organization.
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
// List active and recent STS (role assumption) sessions for the current
// organization. Requires `iam:ListSTSSessions` permission.
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
func (c *Client) ListServiceAccountCredentials(ctx context.Context, serviceAccountID string, opts ...basaltic.RequestOption) (*basaltic.Page[Credential], error) {
	op := &basaltic.Operation{
		ID:       "listServiceAccountCredentials",
		Method:   "GET",
		Path:     "/v1/service-accounts/{service_account_id}/credentials",
		PathArgs: []string{serviceAccountID},
	}
	var out struct {
		Items []Credential `json:"credentials"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[Credential]{Items: out.Items}
	return page, nil
}

// ListServiceAccountGroups lists service account groups.
//
// List all groups the service account belongs to.
func (c *Client) ListServiceAccountGroups(ctx context.Context, serviceAccountID string, opts ...basaltic.RequestOption) (*basaltic.Page[Group], error) {
	op := &basaltic.Operation{
		ID:       "listServiceAccountGroups",
		Method:   "GET",
		Path:     "/v1/service-accounts/{service_account_id}/groups",
		PathArgs: []string{serviceAccountID},
	}
	var out struct {
		Items []Group `json:"groups"`
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
	page := &basaltic.Page[Group]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListServiceAccountInlinePolicies lists a service account's inline policies.
func (c *Client) ListServiceAccountInlinePolicies(ctx context.Context, serviceAccountID string, opts ...basaltic.RequestOption) (*basaltic.Page[InlinePolicy], error) {
	op := &basaltic.Operation{
		ID:       "listServiceAccountInlinePolicies",
		Method:   "GET",
		Path:     "/v1/service-accounts/{service_account_id}/inline-policies",
		PathArgs: []string{serviceAccountID},
	}
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
func (c *Client) ListServiceAccountPolicies(ctx context.Context, serviceAccountID string, opts ...basaltic.RequestOption) (*basaltic.Page[Policy], error) {
	op := &basaltic.Operation{
		ID:       "listServiceAccountPolicies",
		Method:   "GET",
		Path:     "/v1/service-accounts/{service_account_id}/policies",
		PathArgs: []string{serviceAccountID},
	}
	var out struct {
		Items []Policy `json:"policies"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[Policy]{Items: out.Items}
	return page, nil
}

// ListServiceAccounts lists service accounts.
//
// List service accounts owned by the account selected via X-Account-Id
// within the current organization.
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

// ListUserGroups lists user groups.
//
// List all groups the user belongs to.
func (c *Client) ListUserGroups(ctx context.Context, userID string, opts ...basaltic.RequestOption) (*basaltic.Page[Group], error) {
	op := &basaltic.Operation{
		ID:       "listUserGroups",
		Method:   "GET",
		Path:     "/v1/users/{user_id}/groups",
		PathArgs: []string{userID},
	}
	var out struct {
		Items []Group `json:"groups"`
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
	page := &basaltic.Page[Group]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListUserInlinePolicies lists a user's inline policies.
func (c *Client) ListUserInlinePolicies(ctx context.Context, userID string, opts ...basaltic.RequestOption) (*basaltic.Page[InlinePolicy], error) {
	op := &basaltic.Operation{
		ID:       "listUserInlinePolicies",
		Method:   "GET",
		Path:     "/v1/users/{user_id}/inline-policies",
		PathArgs: []string{userID},
	}
	var out struct {
		Items []InlinePolicy `json:"inline_policies"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[InlinePolicy]{Items: out.Items}
	return page, nil
}

// ListUserPolicies lists user policies.
//
// List all policies attached to a user.
func (c *Client) ListUserPolicies(ctx context.Context, userID string, opts ...basaltic.RequestOption) (*basaltic.Page[Policy], error) {
	op := &basaltic.Operation{
		ID:       "listUserPolicies",
		Method:   "GET",
		Path:     "/v1/users/{user_id}/policies",
		PathArgs: []string{userID},
	}
	var out struct {
		Items []Policy `json:"policies"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[Policy]{Items: out.Items}
	return page, nil
}

// ListUsers lists users.
//
// List all users in the current organization.
//
// Returns one page. Use ListUsersAll to walk every page.
func (c *Client) ListUsers(ctx context.Context, params *ListUsersParams, opts ...basaltic.RequestOption) (*basaltic.Page[User], error) {
	op := &basaltic.Operation{
		ID:     "listUsers",
		Method: "GET",
		Path:   "/v1/users",
	}
	op.Query = params.query()
	var out struct {
		Items []User `json:"users"`
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
	page := &basaltic.Page[User]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListUsersAll walks every page of ListUsers, yielding one item at a
// time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListUsersAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListUsersAll(ctx context.Context, params *ListUsersParams, opts ...basaltic.RequestOption) iter.Seq2[User, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[User], error) {
		return c.ListUsers(ctx, params.withMarker(marker), opts...)
	})
}

// PutGroupInlinePolicy creates or replace a group's inline policy.
func (c *Client) PutGroupInlinePolicy(ctx context.Context, groupID string, policyName string, body *PutInlinePolicyRequest, opts ...basaltic.RequestOption) (*InlinePolicy, error) {
	op := &basaltic.Operation{
		ID:       "putGroupInlinePolicy",
		Method:   "PUT",
		Path:     "/v1/groups/{group_id}/inline-policies/{policy_name}",
		PathArgs: []string{groupID, policyName},
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

// PutUserInlinePolicy creates or replace a user's inline policy.
func (c *Client) PutUserInlinePolicy(ctx context.Context, userID string, policyName string, body *PutInlinePolicyRequest, opts ...basaltic.RequestOption) (*InlinePolicy, error) {
	op := &basaltic.Operation{
		ID:       "putUserInlinePolicy",
		Method:   "PUT",
		Path:     "/v1/users/{user_id}/inline-policies/{policy_name}",
		PathArgs: []string{userID, policyName},
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

// RemoveServiceAccountFromGroup removes service account from group.
//
// Remove a service account from a group.
func (c *Client) RemoveServiceAccountFromGroup(ctx context.Context, serviceAccountID string, groupID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "removeServiceAccountFromGroup",
		Method:   "DELETE",
		Path:     "/v1/service-accounts/{service_account_id}/groups/{group_id}",
		PathArgs: []string{serviceAccountID, groupID},
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

// RemoveUser removes user from organization.
//
// Remove a user from the organization.
func (c *Client) RemoveUser(ctx context.Context, userID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "removeUser",
		Method:   "DELETE",
		Path:     "/v1/users/{user_id}",
		PathArgs: []string{userID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// RemoveUserFromGroup removes user from group.
//
// Remove a user from a group.
func (c *Client) RemoveUserFromGroup(ctx context.Context, userID string, groupID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "removeUserFromGroup",
		Method:   "DELETE",
		Path:     "/v1/users/{user_id}/groups/{group_id}",
		PathArgs: []string{userID, groupID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// RemoveUserPermissionBoundary removes a user's permission boundary.
func (c *Client) RemoveUserPermissionBoundary(ctx context.Context, userID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "removeUserPermissionBoundary",
		Method:   "DELETE",
		Path:     "/v1/users/{user_id}/permission-boundary",
		PathArgs: []string{userID},
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
// Revoke an active STS session, invalidating its credentials. Requires
// `iam:RevokeSTSSession` permission.
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

// SetUserPermissionBoundary sets a user's permission boundary.
//
// Set the permission boundary — the boundary caps the principal's
// effective permissions to the intersection with its identity policies.
func (c *Client) SetUserPermissionBoundary(ctx context.Context, userID string, body *SetBoundaryRequest, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "setUserPermissionBoundary",
		Method:   "PUT",
		Path:     "/v1/users/{user_id}/permission-boundary",
		PathArgs: []string{userID},
		Body:     body,
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// UpdateAccount updates account.
func (c *Client) UpdateAccount(ctx context.Context, accountID string, body *UpdateAccountRequest, opts ...basaltic.RequestOption) (*Account, error) {
	op := &basaltic.Operation{
		ID:       "updateAccount",
		Method:   "PATCH",
		Path:     "/v1/accounts/{account_id}",
		PathArgs: []string{accountID},
		Body:     body,
	}
	var out struct {
		Account *Account `json:"account"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Account, nil
}

// UpdateGroup updates group.
//
// Update an existing group.
func (c *Client) UpdateGroup(ctx context.Context, groupID string, body *GroupUpdateRequest, opts ...basaltic.RequestOption) (*Group, error) {
	op := &basaltic.Operation{
		ID:       "updateGroup",
		Method:   "PATCH",
		Path:     "/v1/groups/{group_id}",
		PathArgs: []string{groupID},
		Body:     body,
	}
	var out struct {
		Group *Group `json:"group"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Group, nil
}

// UpdateOrganization updates organization.
//
// Update an existing organization.
func (c *Client) UpdateOrganization(ctx context.Context, organizationID string, body *OrganizationUpdateRequest, opts ...basaltic.RequestOption) (*Organization, error) {
	op := &basaltic.Operation{
		ID:       "updateOrganization",
		Method:   "PATCH",
		Path:     "/v1/organizations/{organization_id}",
		PathArgs: []string{organizationID},
		Body:     body,
	}
	var out struct {
		Organization *Organization `json:"organization"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Organization, nil
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
