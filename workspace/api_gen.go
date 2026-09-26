// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package workspace

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
func (p *ListAccountsParams) query() url.Values {
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
func (p *ListAccountsParams) withMarker(marker string) *ListAccountsParams {
	var out ListAccountsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListGroupInlinePoliciesParams are the optional filters and pagination controls for
// [Client.ListGroupInlinePolicies]. A nil *ListGroupInlinePoliciesParams sends none of them.
type ListGroupInlinePoliciesParams struct {
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
func (p *ListGroupInlinePoliciesParams) query() url.Values {
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

// ListGroupPoliciesParams are the optional filters and pagination controls for
// [Client.ListGroupPolicies]. A nil *ListGroupPoliciesParams sends none of them.
type ListGroupPoliciesParams struct {
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
func (p *ListGroupPoliciesParams) query() url.Values {
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

// ListGroupUsersParams are the optional filters and pagination controls for
// [Client.ListGroupUsers]. A nil *ListGroupUsersParams sends none of them.
type ListGroupUsersParams struct {
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
func (p *ListGroupUsersParams) query() url.Values {
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
func (p *ListGroupsParams) query() url.Values {
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
func (p *ListInvitationsParams) query() url.Values {
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
func (p *ListOrganizationsParams) query() url.Values {
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

// ListPolicyGroupsParams are the optional filters and pagination controls for
// [Client.ListPolicyGroups]. A nil *ListPolicyGroupsParams sends none of them.
type ListPolicyGroupsParams struct {
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
func (p *ListPolicyGroupsParams) query() url.Values {
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

// ListPolicyUsersParams are the optional filters and pagination controls for
// [Client.ListPolicyUsers]. A nil *ListPolicyUsersParams sends none of them.
type ListPolicyUsersParams struct {
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
func (p *ListPolicyUsersParams) query() url.Values {
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
func (p *ListPolicyUsersParams) withMarker(marker string) *ListPolicyUsersParams {
	var out ListPolicyUsersParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
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

// ListUserGroupsParams are the optional filters and pagination controls for
// [Client.ListUserGroups]. A nil *ListUserGroupsParams sends none of them.
type ListUserGroupsParams struct {
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
func (p *ListUserGroupsParams) query() url.Values {
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

// ListUserInlinePoliciesParams are the optional filters and pagination controls for
// [Client.ListUserInlinePolicies]. A nil *ListUserInlinePoliciesParams sends none of them.
type ListUserInlinePoliciesParams struct {
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
func (p *ListUserInlinePoliciesParams) query() url.Values {
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

// ListUserPoliciesParams are the optional filters and pagination controls for
// [Client.ListUserPolicies]. A nil *ListUserPoliciesParams sends none of them.
type ListUserPoliciesParams struct {
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
func (p *ListUserPoliciesParams) query() url.Values {
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

// ListUsersParams are the optional filters and pagination controls for
// [Client.ListUsers]. A nil *ListUsersParams sends none of them.
type ListUsersParams struct {
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
func (p *ListUsersParams) query() url.Values {
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
func (p *ListUsersParams) withMarker(marker string) *ListUsersParams {
	var out ListUsersParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
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

// AssignAccountRole assigns account role.
//
// Assign an account role to a user or group in the organization. The
// assignment grants permission to request role assumption, not direct
// resource access. The role trust policy must independently authorize
// the user. Changing a group assignment affects its human members.
func (c *Client) AssignAccountRole(ctx context.Context, accountID string, body *AccountRoleAssignmentCreateRequest, opts ...basaltic.RequestOption) (*AccountRoleAssignment, error) {
	op := &basaltic.Operation{
		ID:       "assignAccountRole",
		Method:   "POST",
		Path:     "/v1/accounts/{account_id}/role-assignments",
		PathArgs: []string{accountID},
		Body:     body,
	}
	var out struct {
		RoleAssignment *AccountRoleAssignment `json:"role_assignment"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.RoleAssignment, nil
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
// Manage the organization policies delegated to this account identity.
// The identity must belong to the authenticated organization. Changing
// attachments requires organization policy-assignment permission and
// authority to manage the target identity.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) AttachRolePolicy(ctx context.Context, roleID string, body *OrganizationPolicyAttachRequest, opts ...basaltic.RequestOption) error {
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
// Manage the organization policies delegated to this account identity.
// The identity must belong to the authenticated organization. Changing
// attachments requires organization policy-assignment permission and
// authority to manage the target identity.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) AttachServiceAccountPolicy(ctx context.Context, serviceAccountID string, body *OrganizationPolicyAttachRequest, opts ...basaltic.RequestOption) error {
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
// Manage the organization policies delegated to this account identity.
// The identity must belong to the authenticated organization. Changing
// attachments requires organization policy-assignment permission and
// authority to manage the target identity.
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
// Manage the organization policies delegated to this account identity.
// The identity must belong to the authenticated organization. Changing
// attachments requires organization policy-assignment permission and
// authority to manage the target identity.
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

// GetAccountResources — Check account resource presence.
//
// Requires workspace:GetAccount in the owning organization. Checks every
// configured region and the global database, including retained storage
// and hidden rows. A failed check returns an error; it never reports an
// empty account. Telemetry data retention is exposed separately by each
// regional telemetry API.
func (c *Client) GetAccountResources(ctx context.Context, accountID string, opts ...basaltic.RequestOption) (bool, error) {
	op := &basaltic.Operation{
		ID:       "getAccountResources",
		Method:   "GET",
		Path:     "/v1/accounts/{account_id}/resources",
		PathArgs: []string{accountID},
	}
	var out struct {
		HasResources bool `json:"has_resources"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return false, err
	}
	return out.HasResources, nil
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

// ListAccountRoleAssignments lists account role assignments.
//
// List users and groups assigned to roles in this account. An assignment
// supplies source permission to assume its role; the role trust policy
// must still admit the caller.
func (c *Client) ListAccountRoleAssignments(ctx context.Context, accountID string, opts ...basaltic.RequestOption) (*basaltic.Page[AccountRoleAssignment], error) {
	op := &basaltic.Operation{
		ID:       "listAccountRoleAssignments",
		Method:   "GET",
		Path:     "/v1/accounts/{account_id}/role-assignments",
		PathArgs: []string{accountID},
	}
	var out struct {
		Items []AccountRoleAssignment `json:"role_assignments"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[AccountRoleAssignment]{Items: out.Items}
	return page, nil
}

// ListAccountRoles lists available account roles.
//
// List the signed-in human user's effective role assignments across
// accounts in the current organization, including group assignments. A
// listed role still requires successful trust evaluation when assumed.
// Personal authentication is required.
func (c *Client) ListAccountRoles(ctx context.Context, opts ...basaltic.RequestOption) (*basaltic.Page[AccountRole], error) {
	op := &basaltic.Operation{
		ID:     "listAccountRoles",
		Method: "GET",
		Path:   "/v1/account-roles",
	}
	var out struct {
		Items []AccountRole `json:"account_roles"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[AccountRole]{Items: out.Items}
	return page, nil
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
func (c *Client) ListGroupInlinePolicies(ctx context.Context, groupID string, params *ListGroupInlinePoliciesParams, opts ...basaltic.RequestOption) (*basaltic.Page[InlinePolicy], error) {
	op := &basaltic.Operation{
		ID:       "listGroupInlinePolicies",
		Method:   "GET",
		Path:     "/v1/groups/{group_id}/inline-policies",
		PathArgs: []string{groupID},
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

// ListGroupPolicies lists group policies.
//
// List all policies attached to a group.
func (c *Client) ListGroupPolicies(ctx context.Context, groupID string, params *ListGroupPoliciesParams, opts ...basaltic.RequestOption) (*basaltic.Page[Policy], error) {
	op := &basaltic.Operation{
		ID:       "listGroupPolicies",
		Method:   "GET",
		Path:     "/v1/groups/{group_id}/policies",
		PathArgs: []string{groupID},
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
func (c *Client) ListPolicyRoles(ctx context.Context, policyID string, params *ListPolicyRolesParams, opts ...basaltic.RequestOption) (*basaltic.Page[AccountPrincipalReference], error) {
	op := &basaltic.Operation{
		ID:       "listPolicyRoles",
		Method:   "GET",
		Path:     "/v1/policies/{policy_id}/roles",
		PathArgs: []string{policyID},
	}
	op.Query = params.query()
	var out struct {
		Items []AccountPrincipalReference `json:"roles"`
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
	page := &basaltic.Page[AccountPrincipalReference]{Items: out.Items}
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
func (c *Client) ListPolicyRolesAll(ctx context.Context, policyID string, params *ListPolicyRolesParams, opts ...basaltic.RequestOption) iter.Seq2[AccountPrincipalReference, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[AccountPrincipalReference], error) {
		return c.ListPolicyRoles(ctx, policyID, params.withMarker(marker), opts...)
	})
}

// ListPolicyServiceAccounts lists service accounts with policy.
//
// List all service accounts that have this policy attached.
//
// Returns one page. Use ListPolicyServiceAccountsAll to walk every page.
func (c *Client) ListPolicyServiceAccounts(ctx context.Context, policyID string, params *ListPolicyServiceAccountsParams, opts ...basaltic.RequestOption) (*basaltic.Page[AccountPrincipalReference], error) {
	op := &basaltic.Operation{
		ID:       "listPolicyServiceAccounts",
		Method:   "GET",
		Path:     "/v1/policies/{policy_id}/service-accounts",
		PathArgs: []string{policyID},
	}
	op.Query = params.query()
	var out struct {
		Items []AccountPrincipalReference `json:"service_accounts"`
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
	page := &basaltic.Page[AccountPrincipalReference]{Items: out.Items}
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
func (c *Client) ListPolicyServiceAccountsAll(ctx context.Context, policyID string, params *ListPolicyServiceAccountsParams, opts ...basaltic.RequestOption) iter.Seq2[AccountPrincipalReference, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[AccountPrincipalReference], error) {
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

// ListRolePolicies lists role policies.
//
// Manage the organization policies delegated to this account identity.
// The identity must belong to the authenticated organization. Changing
// attachments requires organization policy-assignment permission and
// authority to manage the target identity.
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

// ListServiceAccountPolicies lists service account policies.
//
// Manage the organization policies delegated to this account identity.
// The identity must belong to the authenticated organization. Changing
// attachments requires organization policy-assignment permission and
// authority to manage the target identity.
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

// ListUserGroups lists user groups.
//
// List all groups the user belongs to.
func (c *Client) ListUserGroups(ctx context.Context, userID string, params *ListUserGroupsParams, opts ...basaltic.RequestOption) (*basaltic.Page[Group], error) {
	op := &basaltic.Operation{
		ID:       "listUserGroups",
		Method:   "GET",
		Path:     "/v1/users/{user_id}/groups",
		PathArgs: []string{userID},
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

// ListUserInlinePolicies lists a user's inline policies.
func (c *Client) ListUserInlinePolicies(ctx context.Context, userID string, params *ListUserInlinePoliciesParams, opts ...basaltic.RequestOption) (*basaltic.Page[InlinePolicy], error) {
	op := &basaltic.Operation{
		ID:       "listUserInlinePolicies",
		Method:   "GET",
		Path:     "/v1/users/{user_id}/inline-policies",
		PathArgs: []string{userID},
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

// ListUserPolicies lists user policies.
//
// List all policies attached to a user.
func (c *Client) ListUserPolicies(ctx context.Context, userID string, params *ListUserPoliciesParams, opts ...basaltic.RequestOption) (*basaltic.Page[Policy], error) {
	op := &basaltic.Operation{
		ID:       "listUserPolicies",
		Method:   "GET",
		Path:     "/v1/users/{user_id}/policies",
		PathArgs: []string{userID},
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

// RemoveAccountRoleAssignment removes account role assignment.
//
// Remove the user or group assignment to this account role.
func (c *Client) RemoveAccountRoleAssignment(ctx context.Context, accountID string, assignmentID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "removeAccountRoleAssignment",
		Method:   "DELETE",
		Path:     "/v1/accounts/{account_id}/role-assignments/{assignment_id}",
		PathArgs: []string{accountID, assignmentID},
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

// GetAccountByReference fetches one account by an id, a CRN or a name.
//
// The reference is classified by its syntax alone, exactly as the
// platform does (see [basaltic.ParseReference]): an id is fetched with
// [Client.GetAccount]; a CRN or a name goes to [Client.ListAccounts] as
// an exact filter, together with any filters already set on scope, which
// may be nil. A miss is a not-found error for the kind the string was
// read as — no other kind is tried — and more than one match is a
// [basaltic.AmbiguousReferenceError].
func (c *Client) GetAccountByReference(ctx context.Context, ref string, scope *ListAccountsParams, opts ...basaltic.RequestOption) (*Account, error) {
	return basaltic.ResolveByReference(ctx, ref, "account", "listAccounts", true,
		func(ctx context.Context, refID string) (*Account, error) {
			return c.GetAccount(ctx, refID, opts...)
		},
		func(ctx context.Context, refName, refCRN string) (*basaltic.Page[Account], error) {
			var p ListAccountsParams
			if scope != nil {
				p = *scope
			}
			p.Name = refName
			p.CRN = refCRN
			p.Limit = 2
			return c.ListAccounts(ctx, &p, opts...)
		})
}

// GetGroupByReference fetches one group by an id, a CRN or a name.
//
// The reference is classified by its syntax alone, exactly as the
// platform does (see [basaltic.ParseReference]): an id is fetched with
// [Client.GetGroup]; a CRN or a name goes to [Client.ListGroups] as an
// exact filter, together with any filters already set on scope, which
// may be nil. A miss is a not-found error for the kind the string was
// read as — no other kind is tried — and more than one match is a
// [basaltic.AmbiguousReferenceError].
func (c *Client) GetGroupByReference(ctx context.Context, ref string, scope *ListGroupsParams, opts ...basaltic.RequestOption) (*Group, error) {
	return basaltic.ResolveByReference(ctx, ref, "group", "listGroups", true,
		func(ctx context.Context, refID string) (*Group, error) {
			return c.GetGroup(ctx, refID, opts...)
		},
		func(ctx context.Context, refName, refCRN string) (*basaltic.Page[Group], error) {
			var p ListGroupsParams
			if scope != nil {
				p = *scope
			}
			p.Name = refName
			p.CRN = refCRN
			p.Limit = 2
			return c.ListGroups(ctx, &p, opts...)
		})
}

// GetInvitationByReference fetches one invitation by an id, a CRN or a
// name.
//
// The reference is classified by its syntax alone, exactly as the
// platform does (see [basaltic.ParseReference]): an id is fetched with
// [Client.GetInvitation]; a CRN or a name goes to
// [Client.ListInvitations] as an exact filter, together with any filters
// already set on scope, which may be nil. A miss is a not-found error
// for the kind the string was read as — no other kind is tried — and
// more than one match is a [basaltic.AmbiguousReferenceError].
func (c *Client) GetInvitationByReference(ctx context.Context, ref string, scope *ListInvitationsParams, opts ...basaltic.RequestOption) (*Invitation, error) {
	return basaltic.ResolveByReference(ctx, ref, "invitation", "listInvitations", true,
		func(ctx context.Context, refID string) (*Invitation, error) {
			return c.GetInvitation(ctx, refID, opts...)
		},
		func(ctx context.Context, refName, refCRN string) (*basaltic.Page[Invitation], error) {
			var p ListInvitationsParams
			if scope != nil {
				p = *scope
			}
			p.Name = refName
			p.CRN = refCRN
			p.Limit = 2
			return c.ListInvitations(ctx, &p, opts...)
		})
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

// GetUserByReference fetches one user by an id, a CRN or a name.
//
// The reference is classified by its syntax alone, exactly as the
// platform does (see [basaltic.ParseReference]): an id is fetched with
// [Client.GetUser]; a CRN or a name goes to [Client.ListUsers] as an
// exact filter, together with any filters already set on scope, which
// may be nil. A miss is a not-found error for the kind the string was
// read as — no other kind is tried — and more than one match is a
// [basaltic.AmbiguousReferenceError].
func (c *Client) GetUserByReference(ctx context.Context, ref string, scope *ListUsersParams, opts ...basaltic.RequestOption) (*User, error) {
	return basaltic.ResolveByReference(ctx, ref, "user", "listUsers", true,
		func(ctx context.Context, refID string) (*User, error) {
			return c.GetUser(ctx, refID, opts...)
		},
		func(ctx context.Context, refName, refCRN string) (*basaltic.Page[User], error) {
			var p ListUsersParams
			if scope != nil {
				p = *scope
			}
			p.Name = refName
			p.CRN = refCRN
			p.Limit = 2
			return c.ListUsers(ctx, &p, opts...)
		})
}
