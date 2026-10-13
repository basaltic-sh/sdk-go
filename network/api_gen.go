// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package network

import (
	"context"
	"iter"
	"net/url"
	"strconv"

	basaltic "github.com/basaltic-sh/sdk-go"
)

// ListEgressOnlyGatewayRoutesParams are the optional filters and pagination controls for
// [Client.ListEgressOnlyGatewayRoutes]. A nil *ListEgressOnlyGatewayRoutesParams sends none of them.
type ListEgressOnlyGatewayRoutesParams struct {
	// CRN Exact CRN, validated against the endpoint type, region and caller
	// account. Valid foreign or mismatched CRNs return an empty result;
	// malformed or flat child CRNs return 400. Filters are conjunctive.
	CRN   string
	Limit int

	// Marker resume token — the last id from the previous page.
	Marker string

	// Name exact resource name. This resource has no name, so a supplied name
	// returns an empty result.
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListEgressOnlyGatewayRoutesParams) query() url.Values {
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
func (p *ListEgressOnlyGatewayRoutesParams) withMarker(marker string) *ListEgressOnlyGatewayRoutesParams {
	var out ListEgressOnlyGatewayRoutesParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListEgressOnlyGatewaysParams are the optional filters and pagination controls for
// [Client.ListEgressOnlyGateways]. A nil *ListEgressOnlyGatewaysParams sends none of them.
type ListEgressOnlyGatewaysParams struct {
	// CRN Exact CRN, validated against the endpoint type, region and caller
	// account. Valid foreign or mismatched CRNs return an empty result;
	// malformed or flat child CRNs return 400. Filters are conjunctive.
	CRN   string
	Limit int

	// Marker resume token — the last id from the previous page.
	Marker string

	// Name exact resource name.
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListEgressOnlyGatewaysParams) query() url.Values {
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
func (p *ListEgressOnlyGatewaysParams) withMarker(marker string) *ListEgressOnlyGatewaysParams {
	var out ListEgressOnlyGatewaysParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListFloatingIPInterfacesParams are the optional filters and pagination controls for
// [Client.ListFloatingIPInterfaces]. A nil *ListFloatingIPInterfacesParams sends none of them.
type ListFloatingIPInterfacesParams struct {
	// CRN exact interface CRN.
	CRN    string
	Limit  int
	Marker string

	// Name interface names require subnet scope; this collection rejects bare
	// name filters. Use crn instead.
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListFloatingIPInterfacesParams) query() url.Values {
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
func (p *ListFloatingIPInterfacesParams) withMarker(marker string) *ListFloatingIPInterfacesParams {
	var out ListFloatingIPInterfacesParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListFloatingIPsParams are the optional filters and pagination controls for
// [Client.ListFloatingIPs]. A nil *ListFloatingIPsParams sends none of them.
type ListFloatingIPsParams struct {
	// AttachedTo exact attachment CRN: an interface in the interfaces collection, an
	// instance pool (including an empty pool), or a load balancer.
	// Interface membership matches even when the response attached_to is
	// null. Applied before pagination and intersected with other filters.
	// Malformed CRNs return 400; well-formed foreign or mismatched CRNs
	// return an empty page.
	AttachedTo string

	// CRN Exact CRN, validated against the endpoint type, region and caller
	// account. Valid foreign or mismatched CRNs return an empty result;
	// malformed or flat child CRNs return 400. Filters are conjunctive.
	CRN   string
	Limit int

	// Marker resume token — the last id from the previous page.
	Marker string

	// Name exact resource name. This resource has no name, so a supplied name
	// returns an empty result.
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListFloatingIPsParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.AttachedTo != "" {
		q.Set("attached_to", p.AttachedTo)
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
func (p *ListFloatingIPsParams) withMarker(marker string) *ListFloatingIPsParams {
	var out ListFloatingIPsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListInterfaceSecurityGroupsParams are the optional filters and pagination controls for
// [Client.ListInterfaceSecurityGroups]. A nil *ListInterfaceSecurityGroupsParams sends none of them.
type ListInterfaceSecurityGroupsParams struct {
	// CRN Exact CRN, validated against the endpoint type, region and caller
	// account. Valid foreign or mismatched CRNs return an empty result;
	// malformed or flat child CRNs return 400. Filters are conjunctive.
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

	// Name exact resource name.
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListInterfaceSecurityGroupsParams) query() url.Values {
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
func (p *ListInterfaceSecurityGroupsParams) withMarker(marker string) *ListInterfaceSecurityGroupsParams {
	var out ListInterfaceSecurityGroupsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListInterfacesParams are the optional filters and pagination controls for
// [Client.ListInterfaces]. A nil *ListInterfacesParams sends none of them.
type ListInterfacesParams struct {
	// CRN Exact CRN, validated against the endpoint type, region and caller
	// account. Valid foreign or mismatched CRNs return an empty result;
	// malformed or flat child CRNs return 400. Filters are conjunctive.
	CRN   string
	Limit int

	// Marker resume token — the last id from the previous page.
	Marker string

	// Name exact resource name. Requires the subnet filter.
	Name string

	// Subnet UUID or nested CRN; an exact bare name requires the vpc
	// filter.
	Subnet string

	// VPC UUID, CRN or exact account-scoped name. Resolved before subnet.
	VPC string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListInterfacesParams) query() url.Values {
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
	if p.Subnet != "" {
		q.Set("subnet", p.Subnet)
	}
	if p.VPC != "" {
		q.Set("vpc", p.VPC)
	}
	return q
}

// withMarker copies p with the pagination cursor replaced, leaving the
// caller's value untouched across pages.
func (p *ListInterfacesParams) withMarker(marker string) *ListInterfacesParams {
	var out ListInterfacesParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListInternetGatewayRoutesParams are the optional filters and pagination controls for
// [Client.ListInternetGatewayRoutes]. A nil *ListInternetGatewayRoutesParams sends none of them.
type ListInternetGatewayRoutesParams struct {
	// CRN Exact CRN, validated against the endpoint type, region and caller
	// account. Valid foreign or mismatched CRNs return an empty result;
	// malformed or flat child CRNs return 400. Filters are conjunctive.
	CRN   string
	Limit int

	// Marker resume token — the last id from the previous page.
	Marker string

	// Name exact resource name. This resource has no name, so a supplied name
	// returns an empty result.
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListInternetGatewayRoutesParams) query() url.Values {
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
func (p *ListInternetGatewayRoutesParams) withMarker(marker string) *ListInternetGatewayRoutesParams {
	var out ListInternetGatewayRoutesParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListInternetGatewaysParams are the optional filters and pagination controls for
// [Client.ListInternetGateways]. A nil *ListInternetGatewaysParams sends none of them.
type ListInternetGatewaysParams struct {
	// CRN Exact CRN, validated against the endpoint type, region and caller
	// account. Valid foreign or mismatched CRNs return an empty result;
	// malformed or flat child CRNs return 400. Filters are conjunctive.
	CRN   string
	Limit int

	// Marker resume token — the last id from the previous page.
	Marker string

	// Name exact resource name.
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListInternetGatewaysParams) query() url.Values {
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
func (p *ListInternetGatewaysParams) withMarker(marker string) *ListInternetGatewaysParams {
	var out ListInternetGatewaysParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListNATGatewayRoutesParams are the optional filters and pagination controls for
// [Client.ListNATGatewayRoutes]. A nil *ListNATGatewayRoutesParams sends none of them.
type ListNATGatewayRoutesParams struct {
	// CRN Exact CRN, validated against the endpoint type, region and caller
	// account. Valid foreign or mismatched CRNs return an empty result;
	// malformed or flat child CRNs return 400. Filters are conjunctive.
	CRN   string
	Limit int

	// Marker resume token — the last id from the previous page.
	Marker string

	// Name exact resource name. This resource has no name, so a supplied name
	// returns an empty result.
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListNATGatewayRoutesParams) query() url.Values {
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
func (p *ListNATGatewayRoutesParams) withMarker(marker string) *ListNATGatewayRoutesParams {
	var out ListNATGatewayRoutesParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListNATGatewaysParams are the optional filters and pagination controls for
// [Client.ListNATGateways]. A nil *ListNATGatewaysParams sends none of them.
type ListNATGatewaysParams struct {
	// CRN Exact CRN, validated against the endpoint type, region and caller
	// account. Valid foreign or mismatched CRNs return an empty result;
	// malformed or flat child CRNs return 400. Filters are conjunctive.
	CRN   string
	Limit int

	// Marker resume token — the last id from the previous page.
	Marker string

	// Name exact resource name.
	Name string

	// Subnet UUID or nested CRN; an exact bare name requires the vpc
	// filter.
	Subnet string

	// VPC UUID, CRN or exact account-scoped name. Resolved before subnet.
	VPC string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListNATGatewaysParams) query() url.Values {
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
	if p.Subnet != "" {
		q.Set("subnet", p.Subnet)
	}
	if p.VPC != "" {
		q.Set("vpc", p.VPC)
	}
	return q
}

// withMarker copies p with the pagination cursor replaced, leaving the
// caller's value untouched across pages.
func (p *ListNATGatewaysParams) withMarker(marker string) *ListNATGatewaysParams {
	var out ListNATGatewaysParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListRouteTablesParams are the optional filters and pagination controls for
// [Client.ListRouteTables]. A nil *ListRouteTablesParams sends none of them.
type ListRouteTablesParams struct {
	// CRN Exact CRN, validated against the endpoint type, region and caller
	// account. Valid foreign or mismatched CRNs return an empty result;
	// malformed or flat child CRNs return 400. Filters are conjunctive.
	CRN   string
	Limit int

	// Marker resume token — the last id from the previous page.
	Marker string

	// Name exact resource name. Requires the vpc filter.
	Name string

	// VPC UUID, CRN or exact name in the caller account.
	VPC string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListRouteTablesParams) query() url.Values {
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
	if p.VPC != "" {
		q.Set("vpc", p.VPC)
	}
	return q
}

// withMarker copies p with the pagination cursor replaced, leaving the
// caller's value untouched across pages.
func (p *ListRouteTablesParams) withMarker(marker string) *ListRouteTablesParams {
	var out ListRouteTablesParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListRoutesParams are the optional filters and pagination controls for
// [Client.ListRoutes]. A nil *ListRoutesParams sends none of them.
type ListRoutesParams struct {
	// CRN Exact CRN, validated against the endpoint type, region and caller
	// account. Valid foreign or mismatched CRNs return an empty result;
	// malformed or flat child CRNs return 400. Filters are conjunctive.
	CRN   string
	Limit int

	// Marker resume token — the last id from the previous page.
	Marker string

	// Name exact resource name. This resource has no name, so a supplied name
	// returns an empty result.
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListRoutesParams) query() url.Values {
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
func (p *ListRoutesParams) withMarker(marker string) *ListRoutesParams {
	var out ListRoutesParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListSecurityGroupRulesParams are the optional filters and pagination controls for
// [Client.ListSecurityGroupRules]. A nil *ListSecurityGroupRulesParams sends none of them.
type ListSecurityGroupRulesParams struct {
	// CRN Exact CRN, validated against the endpoint type, region and caller
	// account. Valid foreign or mismatched CRNs return an empty result;
	// malformed or flat child CRNs return 400. Filters are conjunctive.
	CRN   string
	Limit int

	// Marker resume token — the last id from the previous page.
	Marker string

	// Name exact resource name. This resource has no name, so a supplied name
	// returns an empty result.
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListSecurityGroupRulesParams) query() url.Values {
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
func (p *ListSecurityGroupRulesParams) withMarker(marker string) *ListSecurityGroupRulesParams {
	var out ListSecurityGroupRulesParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListSecurityGroupsParams are the optional filters and pagination controls for
// [Client.ListSecurityGroups]. A nil *ListSecurityGroupsParams sends none of them.
type ListSecurityGroupsParams struct {
	// CRN Exact CRN, validated against the endpoint type, region and caller
	// account. Valid foreign or mismatched CRNs return an empty result;
	// malformed or flat child CRNs return 400. Filters are conjunctive.
	CRN   string
	Limit int

	// Marker resume token — the last id from the previous page.
	Marker string

	// Name exact resource name.
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListSecurityGroupsParams) query() url.Values {
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
func (p *ListSecurityGroupsParams) withMarker(marker string) *ListSecurityGroupsParams {
	var out ListSecurityGroupsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListSubnetsParams are the optional filters and pagination controls for
// [Client.ListSubnets]. A nil *ListSubnetsParams sends none of them.
type ListSubnetsParams struct {
	// CRN Exact CRN, validated against the endpoint type, region and caller
	// account. Valid foreign or mismatched CRNs return an empty result;
	// malformed or flat child CRNs return 400. Filters are conjunctive.
	CRN   string
	Limit int

	// Marker resume token — the last id from the previous page.
	Marker string

	// Name exact resource name. Requires the vpc filter.
	Name string

	// VPC UUID, CRN or exact name in the caller account.
	VPC string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListSubnetsParams) query() url.Values {
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
	if p.VPC != "" {
		q.Set("vpc", p.VPC)
	}
	return q
}

// withMarker copies p with the pagination cursor replaced, leaving the
// caller's value untouched across pages.
func (p *ListSubnetsParams) withMarker(marker string) *ListSubnetsParams {
	var out ListSubnetsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListVPCsParams are the optional filters and pagination controls for
// [Client.ListVPCs]. A nil *ListVPCsParams sends none of them.
type ListVPCsParams struct {
	// CRN Exact CRN, validated against the endpoint type, region and caller
	// account. Valid foreign or mismatched CRNs return an empty result;
	// malformed or flat child CRNs return 400. Filters are conjunctive.
	CRN   string
	Limit int

	// Marker resume token — the last id from the previous page.
	Marker string

	// Name exact resource name.
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListVPCsParams) query() url.Values {
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
func (p *ListVPCsParams) withMarker(marker string) *ListVPCsParams {
	var out ListVPCsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// AttachFloatingIP attaches an interface to a floating IP.
//
// Adds the interface to the floating IP. Reattaching the same interface
// and address is idempotent.
//
// Multiple direct interfaces must belong to the same VPC and require
// anycast support in the region. Pool-owned and load-balancer-owned
// addresses are refused (409); manage them through their owner. A direct
// interface without a health check has health `unknown` and remains
// advertised.
//
// A public floating IP requires the interface subnet to route the family
// default (`0.0.0.0/0` or `::/0`) to an internet gateway. A private
// floating IP requires an interface in its allocation VPC. Each
// interface can receive one floating IP per family and visibility.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) AttachFloatingIP(ctx context.Context, floatingIPID string, body *AttachFloatingIPRequest, opts ...basaltic.RequestOption) (*FloatingIPInterface, error) {
	op := &basaltic.Operation{
		ID:       "attachFloatingIp",
		Method:   "POST",
		Path:     "/v1/floating-ips/{floating_ip_id}/interfaces",
		PathArgs: []string{floatingIPID},
		Body:     body,
	}
	var out struct {
		Interface *FloatingIPInterface `json:"interface"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Interface, nil
}

// AttachInterfaceSecurityGroup attaches a security group to an interface.
//
// Adds one security group without changing the other attached groups.
func (c *Client) AttachInterfaceSecurityGroup(ctx context.Context, interfaceID string, body *InterfaceSecurityGroupRequest, opts ...basaltic.RequestOption) (*SecurityGroup, error) {
	op := &basaltic.Operation{
		ID:       "attachInterfaceSecurityGroup",
		Method:   "POST",
		Path:     "/v1/interfaces/{interface_id}/security-groups",
		PathArgs: []string{interfaceID},
		Body:     body,
	}
	var out struct {
		SecurityGroup *SecurityGroup `json:"security_group"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.SecurityGroup, nil
}

// CreateEgressOnlyGateway creates egress-only gateway.
//
// Registers a VPC-scoped egress-only internet gateway. The target VPC
// must have an IPv6 CIDR. It becomes effective when a subnet points its
// route table's ::/0 at it (target_egress_only_gateway_id on the route).
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateEgressOnlyGateway(ctx context.Context, body *EgressOnlyGatewayCreateRequest, opts ...basaltic.RequestOption) (*EgressOnlyGateway, error) {
	op := &basaltic.Operation{
		ID:     "createEgressOnlyGateway",
		Method: "POST",
		Path:   "/v1/egress-only-gateways",
		Body:   body,
	}
	var out struct {
		EgressOnlyGateway *EgressOnlyGateway `json:"egress_only_gateway"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.EgressOnlyGateway, nil
}

// CreateFloatingIP allocates floating IP.
//
// Allocate the next free public address from the region's tenant pool of
// the requested family — `ipv4` unless the body says `ipv6`. The
// address is reserved against the caller's account but not yet bound to
// anything, and its family is fixed from here on.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateFloatingIP(ctx context.Context, body *FloatingIPCreateRequest, opts ...basaltic.RequestOption) (*FloatingIP, error) {
	op := &basaltic.Operation{
		ID:     "createFloatingIp",
		Method: "POST",
		Path:   "/v1/floating-ips",
		Body:   body,
	}
	var out struct {
		FloatingIP *FloatingIP `json:"floating_ip"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.FloatingIP, nil
}

// CreateInterface creates interface.
//
// Allocate a new NIC on a subnet, with the chosen MAC and IP. IP and MAC
// are auto-allocated if omitted.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateInterface(ctx context.Context, body *InterfaceCreateRequest, opts ...basaltic.RequestOption) (*Interface, error) {
	op := &basaltic.Operation{
		ID:     "createInterface",
		Method: "POST",
		Path:   "/v1/interfaces",
		Body:   body,
	}
	var out struct {
		Interface *Interface `json:"interface"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Interface, nil
}

// CreateInterfaceAddress creates interface address.
//
// Every enabled subnet family is assigned automatically. The current
// one-address-per-family limit rejects additional addresses with 409;
// IPv6 must first be enabled on the subnet.
func (c *Client) CreateInterfaceAddress(ctx context.Context, interfaceID string, body *AddressRequest, opts ...basaltic.RequestOption) (*InterfaceAddress, error) {
	op := &basaltic.Operation{
		ID:       "createInterfaceAddress",
		Method:   "POST",
		Path:     "/v1/interfaces/{interface_id}/addresses",
		PathArgs: []string{interfaceID},
		Body:     body,
	}
	var out struct {
		Address *InterfaceAddress `json:"address"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Address, nil
}

// CreateInterfacePrefix creates interface prefix.
func (c *Client) CreateInterfacePrefix(ctx context.Context, interfaceID string, body *CreateInterfacePrefixRequest, opts ...basaltic.RequestOption) (*RoutedPrefix, error) {
	op := &basaltic.Operation{
		ID:       "createInterfacePrefix",
		Method:   "POST",
		Path:     "/v1/interfaces/{interface_id}/prefixes",
		PathArgs: []string{interfaceID},
		Body:     body,
	}
	var out struct {
		Prefix *RoutedPrefix `json:"prefix"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Prefix, nil
}

// CreateInternetGateway creates internet gateway.
//
// Creates an internet gateway bound to the selected VPC. Each VPC can
// have one internet gateway. Add routes separately to send traffic
// through it.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateInternetGateway(ctx context.Context, body *InternetGatewayCreateRequest, opts ...basaltic.RequestOption) (*InternetGateway, error) {
	op := &basaltic.Operation{
		ID:     "createInternetGateway",
		Method: "POST",
		Path:   "/v1/internet-gateways",
		Body:   body,
	}
	var out struct {
		InternetGateway *InternetGateway `json:"internet_gateway"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.InternetGateway, nil
}

// CreateNATGateway creates NAT gateway.
//
// Allocates public IPv4 and, when the hosting subnet has IPv6, public
// IPv6 from the regional tenant pool. Enabling IPv6 on the hosting
// subnet later also assigns the gateway's public IPv6. Both addresses
// are stable until gateway deletion and count against their family's
// public IP quota. Allocation does not depend on routes. Subnets use the
// gateway for shared source NAT by routing 0.0.0.0/0 or ::/0 to it. IPv6
// source NAT supports ULA and global addresses. The target VPC must
// already have an Internet Gateway attached.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateNATGateway(ctx context.Context, body *NATGatewayCreateRequest, opts ...basaltic.RequestOption) (*NATGateway, error) {
	op := &basaltic.Operation{
		ID:     "createNATGateway",
		Method: "POST",
		Path:   "/v1/nat-gateways",
		Body:   body,
	}
	var out struct {
		NATGateway *NATGateway `json:"nat_gateway"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.NATGateway, nil
}

// CreatePrefixPool creates prefix pool.
func (c *Client) CreatePrefixPool(ctx context.Context, vpcID string, body *CreatePrefixPoolRequest, opts ...basaltic.RequestOption) (*PrefixPool, error) {
	op := &basaltic.Operation{
		ID:       "createPrefixPool",
		Method:   "POST",
		Path:     "/v1/vpcs/{vpc_id}/prefix-pools",
		PathArgs: []string{vpcID},
		Body:     body,
	}
	var out struct {
		PrefixPool *PrefixPool `json:"prefix_pool"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.PrefixPool, nil
}

// CreateRoute creates route.
//
// Install a static route on the route table.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateRoute(ctx context.Context, routeTableID string, body *RouteCreateRequest, opts ...basaltic.RequestOption) (*Route, error) {
	op := &basaltic.Operation{
		ID:       "createRoute",
		Method:   "POST",
		Path:     "/v1/route-tables/{route_table_id}/routes",
		PathArgs: []string{routeTableID},
		Body:     body,
	}
	var out struct {
		Route *Route `json:"route"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Route, nil
}

// CreateRouteTable creates route table.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateRouteTable(ctx context.Context, body *RouteTableCreateRequest, opts ...basaltic.RequestOption) (*RouteTable, error) {
	op := &basaltic.Operation{
		ID:     "createRouteTable",
		Method: "POST",
		Path:   "/v1/route-tables",
		Body:   body,
	}
	var out struct {
		RouteTable *RouteTable `json:"route_table"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.RouteTable, nil
}

// CreateSecurityGroup creates security group.
//
// Creates the security group with a default-deny policy in both
// directions, plus one default `egress all 0.0.0.0/0` allow rule
// (matches AWS — customer can delete to lock egress down).
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateSecurityGroup(ctx context.Context, body *SecurityGroupCreateRequest, opts ...basaltic.RequestOption) (*SecurityGroup, error) {
	op := &basaltic.Operation{
		ID:     "createSecurityGroup",
		Method: "POST",
		Path:   "/v1/security-groups",
		Body:   body,
	}
	var out struct {
		SecurityGroup *SecurityGroup `json:"security_group"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.SecurityGroup, nil
}

// CreateSecurityGroupRule creates security group rule.
//
// Appends one stateful allow rule to the security group. When multiple
// security groups apply to one interface, their allow rules are unioned.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateSecurityGroupRule(ctx context.Context, securityGroupID string, body *SecurityGroupRuleCreateRequest, opts ...basaltic.RequestOption) (*SecurityGroupRule, error) {
	op := &basaltic.Operation{
		ID:       "createSecurityGroupRule",
		Method:   "POST",
		Path:     "/v1/security-groups/{security_group_id}/rules",
		PathArgs: []string{securityGroupID},
		Body:     body,
	}
	var out struct {
		Rule *SecurityGroupRule `json:"rule"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Rule, nil
}

// CreateSubnet creates subnet.
//
// Carve a CIDR out of a VPC. The subnet network and its link to the VPC
// router are provisioned automatically.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateSubnet(ctx context.Context, body *SubnetCreateRequest, opts ...basaltic.RequestOption) (*Subnet, error) {
	op := &basaltic.Operation{
		ID:     "createSubnet",
		Method: "POST",
		Path:   "/v1/subnets",
		Body:   body,
	}
	var out struct {
		Subnet *Subnet `json:"subnet"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Subnet, nil
}

// CreateVPC creates VPC.
//
// Create a new VPC. The virtual router that backs it is provisioned
// automatically.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateVPC(ctx context.Context, body *VPCCreateRequest, opts ...basaltic.RequestOption) (*VPC, error) {
	op := &basaltic.Operation{
		ID:     "createVpc",
		Method: "POST",
		Path:   "/v1/vpcs",
		Body:   body,
	}
	var out struct {
		VPC *VPC `json:"vpc"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.VPC, nil
}

// DeleteEgressOnlyGateway deletes egress-only gateway.
//
// Drops the gateway. Refuses if any route still references it —
// rewrite those routes first (which also tears down the inbound drop +
// NAT).
func (c *Client) DeleteEgressOnlyGateway(ctx context.Context, egressOnlyGatewayID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteEgressOnlyGateway",
		Method:   "DELETE",
		Path:     "/v1/egress-only-gateways/{egress_only_gateway_id}",
		PathArgs: []string{egressOnlyGatewayID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteFloatingIP releases floating IP.
//
// Release the address. It must have no owner and an empty interfaces
// array; detach all bindings first.
func (c *Client) DeleteFloatingIP(ctx context.Context, floatingIPID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteFloatingIp",
		Method:   "DELETE",
		Path:     "/v1/floating-ips/{floating_ip_id}",
		PathArgs: []string{floatingIPID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteInterface deletes interface.
//
// Delete an interface. Refuses if the interface is attached.
func (c *Client) DeleteInterface(ctx context.Context, interfaceID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteInterface",
		Method:   "DELETE",
		Path:     "/v1/interfaces/{interface_id}",
		PathArgs: []string{interfaceID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteInterfaceAddress deletes interface address.
//
// Primary IPv4 and IPv6 addresses required by the subnet cannot be
// deleted individually (409).
func (c *Client) DeleteInterfaceAddress(ctx context.Context, interfaceID string, addressID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteInterfaceAddress",
		Method:   "DELETE",
		Path:     "/v1/interfaces/{interface_id}/addresses/{address_id}",
		PathArgs: []string{interfaceID, addressID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteInterfacePrefix deletes interface prefix.
func (c *Client) DeleteInterfacePrefix(ctx context.Context, interfaceID string, prefixID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteInterfacePrefix",
		Method:   "DELETE",
		Path:     "/v1/interfaces/{interface_id}/prefixes/{prefix_id}",
		PathArgs: []string{interfaceID, prefixID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteInternetGateway deletes internet gateway.
//
// Deletes the gateway and its VPC connection. Remove routes targeting
// the gateway first.
func (c *Client) DeleteInternetGateway(ctx context.Context, internetGatewayID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteInternetGateway",
		Method:   "DELETE",
		Path:     "/v1/internet-gateways/{internet_gateway_id}",
		PathArgs: []string{internetGatewayID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteNATGateway deletes NAT gateway.
//
// Removes the source NAT configuration and returns the EIP to the pool.
// Refuses if any route still references this NAT gateway.
func (c *Client) DeleteNATGateway(ctx context.Context, natGatewayID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteNATGateway",
		Method:   "DELETE",
		Path:     "/v1/nat-gateways/{nat_gateway_id}",
		PathArgs: []string{natGatewayID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeletePrefixPool deletes prefix pool.
func (c *Client) DeletePrefixPool(ctx context.Context, vpcID string, poolID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deletePrefixPool",
		Method:   "DELETE",
		Path:     "/v1/vpcs/{vpc_id}/prefix-pools/{pool_id}",
		PathArgs: []string{vpcID, poolID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteRoute deletes route.
func (c *Client) DeleteRoute(ctx context.Context, routeTableID string, routeID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteRoute",
		Method:   "DELETE",
		Path:     "/v1/route-tables/{route_table_id}/routes/{route_id}",
		PathArgs: []string{routeTableID, routeID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteRouteTable deletes route table.
//
// Drops a customer-created table. Refuses if any subnet still associates
// with it (re-associate them first) or if it's the VPC's `main` table
// (immutable).
func (c *Client) DeleteRouteTable(ctx context.Context, routeTableID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteRouteTable",
		Method:   "DELETE",
		Path:     "/v1/route-tables/{route_table_id}",
		PathArgs: []string{routeTableID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteSecurityGroup deletes security group.
//
// Refuses if any interface is still attached — detach first.
func (c *Client) DeleteSecurityGroup(ctx context.Context, securityGroupID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteSecurityGroup",
		Method:   "DELETE",
		Path:     "/v1/security-groups/{security_group_id}",
		PathArgs: []string{securityGroupID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteSecurityGroupRule deletes security group rule.
func (c *Client) DeleteSecurityGroupRule(ctx context.Context, securityGroupID string, ruleID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteSecurityGroupRule",
		Method:   "DELETE",
		Path:     "/v1/security-groups/{security_group_id}/rules/{rule_id}",
		PathArgs: []string{securityGroupID, ruleID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteSubnet deletes subnet.
//
// Delete a subnet. Its network and the link to the VPC router are torn
// down atomically.
func (c *Client) DeleteSubnet(ctx context.Context, subnetID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteSubnet",
		Method:   "DELETE",
		Path:     "/v1/subnets/{subnet_id}",
		PathArgs: []string{subnetID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteVPC deletes VPC.
//
// Delete a VPC. The VPC must have no subnets — delete them first.
func (c *Client) DeleteVPC(ctx context.Context, vpcID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteVpc",
		Method:   "DELETE",
		Path:     "/v1/vpcs/{vpc_id}",
		PathArgs: []string{vpcID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DetachFloatingIP detaches all direct interfaces from a floating IP.
//
// Remove all direct interface bindings. Pool and load balancer bindings
// must be removed through their owners.
func (c *Client) DetachFloatingIP(ctx context.Context, floatingIPID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "detachFloatingIp",
		Method:   "DELETE",
		Path:     "/v1/floating-ips/{floating_ip_id}/interfaces",
		PathArgs: []string{floatingIPID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DetachFloatingIPInterface detaches an interface from a floating IP.
//
// Remove one direct interface binding. Other interfaces continue
// receiving traffic.
func (c *Client) DetachFloatingIPInterface(ctx context.Context, floatingIPID string, interfaceID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "detachFloatingIpInterface",
		Method:   "DELETE",
		Path:     "/v1/floating-ips/{floating_ip_id}/interfaces/{interface_id}",
		PathArgs: []string{floatingIPID, interfaceID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DetachInterfaceSecurityGroup detaches a security group from an interface.
//
// Removes this attachment without changing the other attached groups.
func (c *Client) DetachInterfaceSecurityGroup(ctx context.Context, interfaceID string, securityGroupID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "detachInterfaceSecurityGroup",
		Method:   "DELETE",
		Path:     "/v1/interfaces/{interface_id}/security-groups/{security_group_id}",
		PathArgs: []string{interfaceID, securityGroupID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// EnableSubnetIPv6 enables IPv6 on a subnet.
//
// Assign an IPv6 range after creation. Existing ranges cannot be
// replaced.
func (c *Client) EnableSubnetIPv6(ctx context.Context, subnetID string, body *SubnetIPv6request, opts ...basaltic.RequestOption) (*Subnet, error) {
	op := &basaltic.Operation{
		ID:       "enableSubnetIPv6",
		Method:   "POST",
		Path:     "/v1/subnets/{subnet_id}/ipv6",
		PathArgs: []string{subnetID},
		Body:     body,
	}
	var out struct {
		Subnet *Subnet `json:"subnet"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Subnet, nil
}

// EnableVpciPv6 enables IPv6 on a vpc.
//
// Assign an IPv6 range after creation. Existing ranges cannot be
// replaced.
func (c *Client) EnableVpciPv6(ctx context.Context, vpcID string, body *VPCIPv6request, opts ...basaltic.RequestOption) (*VPC, error) {
	op := &basaltic.Operation{
		ID:       "enableVPCIPv6",
		Method:   "POST",
		Path:     "/v1/vpcs/{vpc_id}/ipv6",
		PathArgs: []string{vpcID},
		Body:     body,
	}
	var out struct {
		VPC *VPC `json:"vpc"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.VPC, nil
}

// GetEgressOnlyGateway gets egress-only gateway.
func (c *Client) GetEgressOnlyGateway(ctx context.Context, egressOnlyGatewayID string, opts ...basaltic.RequestOption) (*EgressOnlyGateway, error) {
	op := &basaltic.Operation{
		ID:       "getEgressOnlyGateway",
		Method:   "GET",
		Path:     "/v1/egress-only-gateways/{egress_only_gateway_id}",
		PathArgs: []string{egressOnlyGatewayID},
	}
	var out struct {
		EgressOnlyGateway *EgressOnlyGateway `json:"egress_only_gateway"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.EgressOnlyGateway, nil
}

// GetFloatingIP gets floating IP.
func (c *Client) GetFloatingIP(ctx context.Context, floatingIPID string, opts ...basaltic.RequestOption) (*FloatingIP, error) {
	op := &basaltic.Operation{
		ID:       "getFloatingIp",
		Method:   "GET",
		Path:     "/v1/floating-ips/{floating_ip_id}",
		PathArgs: []string{floatingIPID},
	}
	var out struct {
		FloatingIP *FloatingIP `json:"floating_ip"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.FloatingIP, nil
}

// GetInterface gets interface.
func (c *Client) GetInterface(ctx context.Context, interfaceID string, opts ...basaltic.RequestOption) (*Interface, error) {
	op := &basaltic.Operation{
		ID:       "getInterface",
		Method:   "GET",
		Path:     "/v1/interfaces/{interface_id}",
		PathArgs: []string{interfaceID},
	}
	var out struct {
		Interface *Interface `json:"interface"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Interface, nil
}

// GetInterfaceAddress gets interface address.
func (c *Client) GetInterfaceAddress(ctx context.Context, interfaceID string, addressID string, opts ...basaltic.RequestOption) (*InterfaceAddress, error) {
	op := &basaltic.Operation{
		ID:       "getInterfaceAddress",
		Method:   "GET",
		Path:     "/v1/interfaces/{interface_id}/addresses/{address_id}",
		PathArgs: []string{interfaceID, addressID},
	}
	var out struct {
		Address *InterfaceAddress `json:"address"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Address, nil
}

// GetInternetGateway gets internet gateway.
func (c *Client) GetInternetGateway(ctx context.Context, internetGatewayID string, opts ...basaltic.RequestOption) (*InternetGateway, error) {
	op := &basaltic.Operation{
		ID:       "getInternetGateway",
		Method:   "GET",
		Path:     "/v1/internet-gateways/{internet_gateway_id}",
		PathArgs: []string{internetGatewayID},
	}
	var out struct {
		InternetGateway *InternetGateway `json:"internet_gateway"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.InternetGateway, nil
}

// GetNATGateway gets NAT gateway.
func (c *Client) GetNATGateway(ctx context.Context, natGatewayID string, opts ...basaltic.RequestOption) (*NATGateway, error) {
	op := &basaltic.Operation{
		ID:       "getNATGateway",
		Method:   "GET",
		Path:     "/v1/nat-gateways/{nat_gateway_id}",
		PathArgs: []string{natGatewayID},
	}
	var out struct {
		NATGateway *NATGateway `json:"nat_gateway"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.NATGateway, nil
}

// GetRoute gets route.
func (c *Client) GetRoute(ctx context.Context, routeTableID string, routeID string, opts ...basaltic.RequestOption) (*Route, error) {
	op := &basaltic.Operation{
		ID:       "getRoute",
		Method:   "GET",
		Path:     "/v1/route-tables/{route_table_id}/routes/{route_id}",
		PathArgs: []string{routeTableID, routeID},
	}
	var out struct {
		Route *Route `json:"route"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Route, nil
}

// GetRouteTable gets route table.
func (c *Client) GetRouteTable(ctx context.Context, routeTableID string, opts ...basaltic.RequestOption) (*RouteTable, error) {
	op := &basaltic.Operation{
		ID:       "getRouteTable",
		Method:   "GET",
		Path:     "/v1/route-tables/{route_table_id}",
		PathArgs: []string{routeTableID},
	}
	var out struct {
		RouteTable *RouteTable `json:"route_table"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.RouteTable, nil
}

// GetSecurityGroup gets security group.
func (c *Client) GetSecurityGroup(ctx context.Context, securityGroupID string, opts ...basaltic.RequestOption) (*SecurityGroup, error) {
	op := &basaltic.Operation{
		ID:       "getSecurityGroup",
		Method:   "GET",
		Path:     "/v1/security-groups/{security_group_id}",
		PathArgs: []string{securityGroupID},
	}
	var out struct {
		SecurityGroup *SecurityGroup `json:"security_group"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.SecurityGroup, nil
}

// GetSecurityGroupRule gets security group rule.
func (c *Client) GetSecurityGroupRule(ctx context.Context, securityGroupID string, ruleID string, opts ...basaltic.RequestOption) (*SecurityGroupRule, error) {
	op := &basaltic.Operation{
		ID:       "getSecurityGroupRule",
		Method:   "GET",
		Path:     "/v1/security-groups/{security_group_id}/rules/{rule_id}",
		PathArgs: []string{securityGroupID, ruleID},
	}
	var out struct {
		Rule *SecurityGroupRule `json:"rule"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Rule, nil
}

// GetSubnet gets subnet.
func (c *Client) GetSubnet(ctx context.Context, subnetID string, opts ...basaltic.RequestOption) (*Subnet, error) {
	op := &basaltic.Operation{
		ID:       "getSubnet",
		Method:   "GET",
		Path:     "/v1/subnets/{subnet_id}",
		PathArgs: []string{subnetID},
	}
	var out struct {
		Subnet *Subnet `json:"subnet"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Subnet, nil
}

// GetVPC gets VPC.
func (c *Client) GetVPC(ctx context.Context, vpcID string, opts ...basaltic.RequestOption) (*VPC, error) {
	op := &basaltic.Operation{
		ID:       "getVpc",
		Method:   "GET",
		Path:     "/v1/vpcs/{vpc_id}",
		PathArgs: []string{vpcID},
	}
	var out struct {
		VPC *VPC `json:"vpc"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.VPC, nil
}

// ListEgressOnlyGatewayRoutes lists egress-only gateway routes.
//
// List routes targeting this egress-only gateway, newest first, with
// route-table identity. Requires read access to the gateway. Only routes
// in tables where the caller has network:ListRoutes permission are
// returned; pagination applies to those visible routes.
//
// Returns one page. Use ListEgressOnlyGatewayRoutesAll to walk every page.
func (c *Client) ListEgressOnlyGatewayRoutes(ctx context.Context, egressOnlyGatewayID string, params *ListEgressOnlyGatewayRoutesParams, opts ...basaltic.RequestOption) (*basaltic.Page[GatewayRoute], error) {
	op := &basaltic.Operation{
		ID:       "listEgressOnlyGatewayRoutes",
		Method:   "GET",
		Path:     "/v1/egress-only-gateways/{egress_only_gateway_id}/routes",
		PathArgs: []string{egressOnlyGatewayID},
	}
	op.Query = params.query()
	var out struct {
		Items []GatewayRoute `json:"routes"`
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
	page := &basaltic.Page[GatewayRoute]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListEgressOnlyGatewayRoutesAll walks every page of
// ListEgressOnlyGatewayRoutes, yielding one item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListEgressOnlyGatewayRoutesAll(ctx, egressOnlyGatewayID, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListEgressOnlyGatewayRoutesAll(ctx context.Context, egressOnlyGatewayID string, params *ListEgressOnlyGatewayRoutesParams, opts ...basaltic.RequestOption) iter.Seq2[GatewayRoute, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[GatewayRoute], error) {
		return c.ListEgressOnlyGatewayRoutes(ctx, egressOnlyGatewayID, params.withMarker(marker), opts...)
	})
}

// ListEgressOnlyGateways lists egress-only gateways.
//
// Returns one page. Use ListEgressOnlyGatewaysAll to walk every page.
func (c *Client) ListEgressOnlyGateways(ctx context.Context, params *ListEgressOnlyGatewaysParams, opts ...basaltic.RequestOption) (*basaltic.Page[EgressOnlyGateway], error) {
	op := &basaltic.Operation{
		ID:     "listEgressOnlyGateways",
		Method: "GET",
		Path:   "/v1/egress-only-gateways",
	}
	op.Query = params.query()
	var out struct {
		Items []EgressOnlyGateway `json:"egress_only_gateways"`
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
	page := &basaltic.Page[EgressOnlyGateway]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListEgressOnlyGatewaysAll walks every page of ListEgressOnlyGateways,
// yielding one item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListEgressOnlyGatewaysAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListEgressOnlyGatewaysAll(ctx context.Context, params *ListEgressOnlyGatewaysParams, opts ...basaltic.RequestOption) iter.Seq2[EgressOnlyGateway, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[EgressOnlyGateway], error) {
		return c.ListEgressOnlyGateways(ctx, params.withMarker(marker), opts...)
	})
}

// ListFloatingIPInterfaces lists a floating IP’s interfaces.
//
// List attached interfaces using the canonical interface contract, with
// health and reason for this binding. Results are paginated by interface
// ID. Reading the floating IP authorizes this related collection.
//
// Returns one page. Use ListFloatingIPInterfacesAll to walk every page.
func (c *Client) ListFloatingIPInterfaces(ctx context.Context, floatingIPID string, params *ListFloatingIPInterfacesParams, opts ...basaltic.RequestOption) (*basaltic.Page[FloatingIPInterface], error) {
	op := &basaltic.Operation{
		ID:       "listFloatingIpInterfaces",
		Method:   "GET",
		Path:     "/v1/floating-ips/{floating_ip_id}/interfaces",
		PathArgs: []string{floatingIPID},
	}
	op.Query = params.query()
	var out struct {
		Items []FloatingIPInterface `json:"interfaces"`
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
	page := &basaltic.Page[FloatingIPInterface]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListFloatingIPInterfacesAll walks every page of
// ListFloatingIPInterfaces, yielding one item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListFloatingIPInterfacesAll(ctx, floatingIPID, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListFloatingIPInterfacesAll(ctx context.Context, floatingIPID string, params *ListFloatingIPInterfacesParams, opts ...basaltic.RequestOption) iter.Seq2[FloatingIPInterface, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[FloatingIPInterface], error) {
		return c.ListFloatingIPInterfaces(ctx, floatingIPID, params.withMarker(marker), opts...)
	})
}

// ListFloatingIPs lists floating IPs.
//
// List floating IPs owned by the caller's account.
//
// Returns one page. Use ListFloatingIPsAll to walk every page.
func (c *Client) ListFloatingIPs(ctx context.Context, params *ListFloatingIPsParams, opts ...basaltic.RequestOption) (*basaltic.Page[FloatingIP], error) {
	op := &basaltic.Operation{
		ID:     "listFloatingIps",
		Method: "GET",
		Path:   "/v1/floating-ips",
	}
	op.Query = params.query()
	var out struct {
		Items []FloatingIP `json:"floating_ips"`
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
	page := &basaltic.Page[FloatingIP]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListFloatingIPsAll walks every page of ListFloatingIPs, yielding one
// item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListFloatingIPsAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListFloatingIPsAll(ctx context.Context, params *ListFloatingIPsParams, opts ...basaltic.RequestOption) iter.Seq2[FloatingIP, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[FloatingIP], error) {
		return c.ListFloatingIPs(ctx, params.withMarker(marker), opts...)
	})
}

// ListInterfaceAddresses lists interface addresses.
func (c *Client) ListInterfaceAddresses(ctx context.Context, interfaceID string, opts ...basaltic.RequestOption) ([]*InterfaceAddress, error) {
	op := &basaltic.Operation{
		ID:       "listInterfaceAddresses",
		Method:   "GET",
		Path:     "/v1/interfaces/{interface_id}/addresses",
		PathArgs: []string{interfaceID},
	}
	var out struct {
		Addresses []*InterfaceAddress `json:"addresses"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Addresses, nil
}

// ListInterfacePrefixes lists interface prefixes.
func (c *Client) ListInterfacePrefixes(ctx context.Context, interfaceID string, opts ...basaltic.RequestOption) (*ListInterfacePrefixesResult, error) {
	op := &basaltic.Operation{
		ID:       "listInterfacePrefixes",
		Method:   "GET",
		Path:     "/v1/interfaces/{interface_id}/prefixes",
		PathArgs: []string{interfaceID},
	}
	var out ListInterfacePrefixesResult
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListInterfaceSecurityGroups lists an interface’s security groups.
//
// Returns one page. Use ListInterfaceSecurityGroupsAll to walk every page.
func (c *Client) ListInterfaceSecurityGroups(ctx context.Context, interfaceID string, params *ListInterfaceSecurityGroupsParams, opts ...basaltic.RequestOption) (*basaltic.Page[SecurityGroup], error) {
	op := &basaltic.Operation{
		ID:       "listInterfaceSecurityGroups",
		Method:   "GET",
		Path:     "/v1/interfaces/{interface_id}/security-groups",
		PathArgs: []string{interfaceID},
	}
	op.Query = params.query()
	var out struct {
		Items []SecurityGroup `json:"security_groups"`
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
	page := &basaltic.Page[SecurityGroup]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListInterfaceSecurityGroupsAll walks every page of
// ListInterfaceSecurityGroups, yielding one item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListInterfaceSecurityGroupsAll(ctx, interfaceID, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListInterfaceSecurityGroupsAll(ctx context.Context, interfaceID string, params *ListInterfaceSecurityGroupsParams, opts ...basaltic.RequestOption) iter.Seq2[SecurityGroup, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[SecurityGroup], error) {
		return c.ListInterfaceSecurityGroups(ctx, interfaceID, params.withMarker(marker), opts...)
	})
}

// ListInterfaces lists interfaces.
//
// List interfaces owned by the caller's account. Filter by `subnet` to
// narrow to one subnet, or by `vpc` to span every subnet in one VPC.
// Both filters apply; the subnet must belong to the VPC.
//
// Returns one page. Use ListInterfacesAll to walk every page.
func (c *Client) ListInterfaces(ctx context.Context, params *ListInterfacesParams, opts ...basaltic.RequestOption) (*basaltic.Page[Interface], error) {
	op := &basaltic.Operation{
		ID:     "listInterfaces",
		Method: "GET",
		Path:   "/v1/interfaces",
	}
	op.Query = params.query()
	var out struct {
		Items []Interface `json:"interfaces"`
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
	page := &basaltic.Page[Interface]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListInterfacesAll walks every page of ListInterfaces, yielding one
// item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListInterfacesAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListInterfacesAll(ctx context.Context, params *ListInterfacesParams, opts ...basaltic.RequestOption) iter.Seq2[Interface, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Interface], error) {
		return c.ListInterfaces(ctx, params.withMarker(marker), opts...)
	})
}

// ListInternetGatewayRoutes lists internet gateway routes.
//
// List routes targeting this internet gateway, newest first, with
// route-table identity. Requires read access to the gateway. Only routes
// in tables where the caller has network:ListRoutes permission are
// returned; pagination applies to those visible routes.
//
// Returns one page. Use ListInternetGatewayRoutesAll to walk every page.
func (c *Client) ListInternetGatewayRoutes(ctx context.Context, internetGatewayID string, params *ListInternetGatewayRoutesParams, opts ...basaltic.RequestOption) (*basaltic.Page[GatewayRoute], error) {
	op := &basaltic.Operation{
		ID:       "listInternetGatewayRoutes",
		Method:   "GET",
		Path:     "/v1/internet-gateways/{internet_gateway_id}/routes",
		PathArgs: []string{internetGatewayID},
	}
	op.Query = params.query()
	var out struct {
		Items []GatewayRoute `json:"routes"`
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
	page := &basaltic.Page[GatewayRoute]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListInternetGatewayRoutesAll walks every page of
// ListInternetGatewayRoutes, yielding one item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListInternetGatewayRoutesAll(ctx, internetGatewayID, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListInternetGatewayRoutesAll(ctx context.Context, internetGatewayID string, params *ListInternetGatewayRoutesParams, opts ...basaltic.RequestOption) iter.Seq2[GatewayRoute, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[GatewayRoute], error) {
		return c.ListInternetGatewayRoutes(ctx, internetGatewayID, params.withMarker(marker), opts...)
	})
}

// ListInternetGateways lists internet gateways.
//
// Returns one page. Use ListInternetGatewaysAll to walk every page.
func (c *Client) ListInternetGateways(ctx context.Context, params *ListInternetGatewaysParams, opts ...basaltic.RequestOption) (*basaltic.Page[InternetGateway], error) {
	op := &basaltic.Operation{
		ID:     "listInternetGateways",
		Method: "GET",
		Path:   "/v1/internet-gateways",
	}
	op.Query = params.query()
	var out struct {
		Items []InternetGateway `json:"internet_gateways"`
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
	page := &basaltic.Page[InternetGateway]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListInternetGatewaysAll walks every page of ListInternetGateways,
// yielding one item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListInternetGatewaysAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListInternetGatewaysAll(ctx context.Context, params *ListInternetGatewaysParams, opts ...basaltic.RequestOption) iter.Seq2[InternetGateway, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[InternetGateway], error) {
		return c.ListInternetGateways(ctx, params.withMarker(marker), opts...)
	})
}

// ListNATGatewayRoutes lists NAT gateway routes.
//
// List routes targeting this NAT gateway, newest first, with route-table
// identity. Requires read access to the gateway. Only routes in tables
// where the caller has network:ListRoutes permission are returned;
// pagination applies to those visible routes.
//
// Returns one page. Use ListNATGatewayRoutesAll to walk every page.
func (c *Client) ListNATGatewayRoutes(ctx context.Context, natGatewayID string, params *ListNATGatewayRoutesParams, opts ...basaltic.RequestOption) (*basaltic.Page[GatewayRoute], error) {
	op := &basaltic.Operation{
		ID:       "listNATGatewayRoutes",
		Method:   "GET",
		Path:     "/v1/nat-gateways/{nat_gateway_id}/routes",
		PathArgs: []string{natGatewayID},
	}
	op.Query = params.query()
	var out struct {
		Items []GatewayRoute `json:"routes"`
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
	page := &basaltic.Page[GatewayRoute]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListNATGatewayRoutesAll walks every page of ListNATGatewayRoutes,
// yielding one item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListNATGatewayRoutesAll(ctx, natGatewayID, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListNATGatewayRoutesAll(ctx context.Context, natGatewayID string, params *ListNATGatewayRoutesParams, opts ...basaltic.RequestOption) iter.Seq2[GatewayRoute, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[GatewayRoute], error) {
		return c.ListNATGatewayRoutes(ctx, natGatewayID, params.withMarker(marker), opts...)
	})
}

// ListNATGateways lists NAT gateways.
//
// List NAT gateways owned by the caller's account. Filter by `subnet` to
// narrow to the gateway living in one subnet, or by `vpc` to span the
// VPC. Both filters apply; the subnet must belong to the VPC.
//
// Returns one page. Use ListNATGatewaysAll to walk every page.
func (c *Client) ListNATGateways(ctx context.Context, params *ListNATGatewaysParams, opts ...basaltic.RequestOption) (*basaltic.Page[NATGateway], error) {
	op := &basaltic.Operation{
		ID:     "listNATGateways",
		Method: "GET",
		Path:   "/v1/nat-gateways",
	}
	op.Query = params.query()
	var out struct {
		Items []NATGateway `json:"nat_gateways"`
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
	page := &basaltic.Page[NATGateway]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListNATGatewaysAll walks every page of ListNATGateways, yielding one
// item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListNATGatewaysAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListNATGatewaysAll(ctx context.Context, params *ListNATGatewaysParams, opts ...basaltic.RequestOption) iter.Seq2[NATGateway, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[NATGateway], error) {
		return c.ListNATGateways(ctx, params.withMarker(marker), opts...)
	})
}

// ListPrefixPools lists prefix pools.
func (c *Client) ListPrefixPools(ctx context.Context, vpcID string, opts ...basaltic.RequestOption) (*ListPrefixPoolsResult, error) {
	op := &basaltic.Operation{
		ID:       "listPrefixPools",
		Method:   "GET",
		Path:     "/v1/vpcs/{vpc_id}/prefix-pools",
		PathArgs: []string{vpcID},
	}
	var out ListPrefixPoolsResult
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListRouteTables lists route tables.
//
// List route tables owned by the caller's account. Filter by vpc to
// narrow to one VPC.
//
// Returns one page. Use ListRouteTablesAll to walk every page.
func (c *Client) ListRouteTables(ctx context.Context, params *ListRouteTablesParams, opts ...basaltic.RequestOption) (*basaltic.Page[RouteTable], error) {
	op := &basaltic.Operation{
		ID:     "listRouteTables",
		Method: "GET",
		Path:   "/v1/route-tables",
	}
	op.Query = params.query()
	var out struct {
		Items []RouteTable `json:"route_tables"`
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
	page := &basaltic.Page[RouteTable]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListRouteTablesAll walks every page of ListRouteTables, yielding one
// item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListRouteTablesAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListRouteTablesAll(ctx context.Context, params *ListRouteTablesParams, opts ...basaltic.RequestOption) iter.Seq2[RouteTable, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[RouteTable], error) {
		return c.ListRouteTables(ctx, params.withMarker(marker), opts...)
	})
}

// ListRoutes lists routes.
//
// List static routes installed on the given route table.
//
// Returns one page. Use ListRoutesAll to walk every page.
func (c *Client) ListRoutes(ctx context.Context, routeTableID string, params *ListRoutesParams, opts ...basaltic.RequestOption) (*basaltic.Page[Route], error) {
	op := &basaltic.Operation{
		ID:       "listRoutes",
		Method:   "GET",
		Path:     "/v1/route-tables/{route_table_id}/routes",
		PathArgs: []string{routeTableID},
	}
	op.Query = params.query()
	var out struct {
		Items []Route `json:"routes"`
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
	page := &basaltic.Page[Route]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListRoutesAll walks every page of ListRoutes, yielding one item at a
// time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListRoutesAll(ctx, routeTableID, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListRoutesAll(ctx context.Context, routeTableID string, params *ListRoutesParams, opts ...basaltic.RequestOption) iter.Seq2[Route, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Route], error) {
		return c.ListRoutes(ctx, routeTableID, params.withMarker(marker), opts...)
	})
}

// ListSecurityGroupRules lists security group rules.
//
// Returns one page. Use ListSecurityGroupRulesAll to walk every page.
func (c *Client) ListSecurityGroupRules(ctx context.Context, securityGroupID string, params *ListSecurityGroupRulesParams, opts ...basaltic.RequestOption) (*basaltic.Page[SecurityGroupRule], error) {
	op := &basaltic.Operation{
		ID:       "listSecurityGroupRules",
		Method:   "GET",
		Path:     "/v1/security-groups/{security_group_id}/rules",
		PathArgs: []string{securityGroupID},
	}
	op.Query = params.query()
	var out struct {
		Items []SecurityGroupRule `json:"rules"`
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
	page := &basaltic.Page[SecurityGroupRule]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListSecurityGroupRulesAll walks every page of ListSecurityGroupRules,
// yielding one item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListSecurityGroupRulesAll(ctx, securityGroupID, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListSecurityGroupRulesAll(ctx context.Context, securityGroupID string, params *ListSecurityGroupRulesParams, opts ...basaltic.RequestOption) iter.Seq2[SecurityGroupRule, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[SecurityGroupRule], error) {
		return c.ListSecurityGroupRules(ctx, securityGroupID, params.withMarker(marker), opts...)
	})
}

// ListSecurityGroups lists security groups.
//
// Returns one page. Use ListSecurityGroupsAll to walk every page.
func (c *Client) ListSecurityGroups(ctx context.Context, params *ListSecurityGroupsParams, opts ...basaltic.RequestOption) (*basaltic.Page[SecurityGroup], error) {
	op := &basaltic.Operation{
		ID:     "listSecurityGroups",
		Method: "GET",
		Path:   "/v1/security-groups",
	}
	op.Query = params.query()
	var out struct {
		Items []SecurityGroup `json:"security_groups"`
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
	page := &basaltic.Page[SecurityGroup]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListSecurityGroupsAll walks every page of ListSecurityGroups, yielding
// one item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListSecurityGroupsAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListSecurityGroupsAll(ctx context.Context, params *ListSecurityGroupsParams, opts ...basaltic.RequestOption) iter.Seq2[SecurityGroup, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[SecurityGroup], error) {
		return c.ListSecurityGroups(ctx, params.withMarker(marker), opts...)
	})
}

// ListSubnets lists subnets.
//
// List subnets owned by the caller's account. Filter by vpc to narrow to
// one VPC.
//
// Returns one page. Use ListSubnetsAll to walk every page.
func (c *Client) ListSubnets(ctx context.Context, params *ListSubnetsParams, opts ...basaltic.RequestOption) (*basaltic.Page[Subnet], error) {
	op := &basaltic.Operation{
		ID:     "listSubnets",
		Method: "GET",
		Path:   "/v1/subnets",
	}
	op.Query = params.query()
	var out struct {
		Items []Subnet `json:"subnets"`
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
	page := &basaltic.Page[Subnet]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListSubnetsAll walks every page of ListSubnets, yielding one item at a
// time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListSubnetsAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListSubnetsAll(ctx context.Context, params *ListSubnetsParams, opts ...basaltic.RequestOption) iter.Seq2[Subnet, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Subnet], error) {
		return c.ListSubnets(ctx, params.withMarker(marker), opts...)
	})
}

// ListVPCs lists VPCs.
//
// List all VPCs owned by the caller's account.
//
// Returns one page. Use ListVPCsAll to walk every page.
func (c *Client) ListVPCs(ctx context.Context, params *ListVPCsParams, opts ...basaltic.RequestOption) (*basaltic.Page[VPC], error) {
	op := &basaltic.Operation{
		ID:     "listVpcs",
		Method: "GET",
		Path:   "/v1/vpcs",
	}
	op.Query = params.query()
	var out struct {
		Items []VPC `json:"vpcs"`
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
	page := &basaltic.Page[VPC]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListVPCsAll walks every page of ListVPCs, yielding one item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListVPCsAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListVPCsAll(ctx context.Context, params *ListVPCsParams, opts ...basaltic.RequestOption) iter.Seq2[VPC, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[VPC], error) {
		return c.ListVPCs(ctx, params.withMarker(marker), opts...)
	})
}

// UpdateEgressOnlyGateway updates egress-only gateway.
//
// Update description / tags. Name + VPC are immutable.
func (c *Client) UpdateEgressOnlyGateway(ctx context.Context, egressOnlyGatewayID string, body *EgressOnlyGatewayUpdateRequest, opts ...basaltic.RequestOption) (*EgressOnlyGateway, error) {
	op := &basaltic.Operation{
		ID:       "updateEgressOnlyGateway",
		Method:   "PATCH",
		Path:     "/v1/egress-only-gateways/{egress_only_gateway_id}",
		PathArgs: []string{egressOnlyGatewayID},
		Body:     body,
	}
	var out struct {
		EgressOnlyGateway *EgressOnlyGateway `json:"egress_only_gateway"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.EgressOnlyGateway, nil
}

// UpdateFloatingIP updates floating IP.
//
// Update description / tags. IP address is immutable.
func (c *Client) UpdateFloatingIP(ctx context.Context, floatingIPID string, body *FloatingIPUpdateRequest, opts ...basaltic.RequestOption) (*FloatingIP, error) {
	op := &basaltic.Operation{
		ID:       "updateFloatingIp",
		Method:   "PATCH",
		Path:     "/v1/floating-ips/{floating_ip_id}",
		PathArgs: []string{floatingIPID},
		Body:     body,
	}
	var out struct {
		FloatingIP *FloatingIP `json:"floating_ip"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.FloatingIP, nil
}

// UpdateInterface updates interface.
//
// Update mutable fields (description, tags). MAC/IP/subnet are
// immutable.
func (c *Client) UpdateInterface(ctx context.Context, interfaceID string, body *InterfaceUpdateRequest, opts ...basaltic.RequestOption) (*Interface, error) {
	op := &basaltic.Operation{
		ID:       "updateInterface",
		Method:   "PATCH",
		Path:     "/v1/interfaces/{interface_id}",
		PathArgs: []string{interfaceID},
		Body:     body,
	}
	var out struct {
		Interface *Interface `json:"interface"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Interface, nil
}

// UpdateInternetGateway updates internet gateway.
//
// Update description / tags. Name is immutable.
func (c *Client) UpdateInternetGateway(ctx context.Context, internetGatewayID string, body *InternetGatewayUpdateRequest, opts ...basaltic.RequestOption) (*InternetGateway, error) {
	op := &basaltic.Operation{
		ID:       "updateInternetGateway",
		Method:   "PATCH",
		Path:     "/v1/internet-gateways/{internet_gateway_id}",
		PathArgs: []string{internetGatewayID},
		Body:     body,
	}
	var out struct {
		InternetGateway *InternetGateway `json:"internet_gateway"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.InternetGateway, nil
}

// UpdateNATGateway updates NAT gateway.
//
// Update description / tags. Name + EIP + subnet are immutable.
func (c *Client) UpdateNATGateway(ctx context.Context, natGatewayID string, body *NATGatewayUpdateRequest, opts ...basaltic.RequestOption) (*NATGateway, error) {
	op := &basaltic.Operation{
		ID:       "updateNATGateway",
		Method:   "PATCH",
		Path:     "/v1/nat-gateways/{nat_gateway_id}",
		PathArgs: []string{natGatewayID},
		Body:     body,
	}
	var out struct {
		NATGateway *NATGateway `json:"nat_gateway"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.NATGateway, nil
}

// UpdateRoute updates route.
//
// Update description / tags. Destination + next-hop are immutable.
func (c *Client) UpdateRoute(ctx context.Context, routeTableID string, routeID string, body *RouteUpdateRequest, opts ...basaltic.RequestOption) (*Route, error) {
	op := &basaltic.Operation{
		ID:       "updateRoute",
		Method:   "PATCH",
		Path:     "/v1/route-tables/{route_table_id}/routes/{route_id}",
		PathArgs: []string{routeTableID, routeID},
		Body:     body,
	}
	var out struct {
		Route *Route `json:"route"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Route, nil
}

// UpdateRouteTable updates route table.
//
// Update description / tags. Name is immutable.
func (c *Client) UpdateRouteTable(ctx context.Context, routeTableID string, body *RouteTableUpdateRequest, opts ...basaltic.RequestOption) (*RouteTable, error) {
	op := &basaltic.Operation{
		ID:       "updateRouteTable",
		Method:   "PATCH",
		Path:     "/v1/route-tables/{route_table_id}",
		PathArgs: []string{routeTableID},
		Body:     body,
	}
	var out struct {
		RouteTable *RouteTable `json:"route_table"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.RouteTable, nil
}

// UpdateSecurityGroup updates security group.
//
// Update description / tags. Name is immutable.
func (c *Client) UpdateSecurityGroup(ctx context.Context, securityGroupID string, body *SecurityGroupUpdateRequest, opts ...basaltic.RequestOption) (*SecurityGroup, error) {
	op := &basaltic.Operation{
		ID:       "updateSecurityGroup",
		Method:   "PATCH",
		Path:     "/v1/security-groups/{security_group_id}",
		PathArgs: []string{securityGroupID},
		Body:     body,
	}
	var out struct {
		SecurityGroup *SecurityGroup `json:"security_group"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.SecurityGroup, nil
}

// UpdateSubnet updates subnet.
//
// Update mutable fields (description, tags). CIDR is immutable.
func (c *Client) UpdateSubnet(ctx context.Context, subnetID string, body *SubnetUpdateRequest, opts ...basaltic.RequestOption) (*Subnet, error) {
	op := &basaltic.Operation{
		ID:       "updateSubnet",
		Method:   "PATCH",
		Path:     "/v1/subnets/{subnet_id}",
		PathArgs: []string{subnetID},
		Body:     body,
	}
	var out struct {
		Subnet *Subnet `json:"subnet"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Subnet, nil
}

// UpdateVPC updates VPC.
//
// Update mutable fields (description, tags). Name and CIDR are
// immutable.
func (c *Client) UpdateVPC(ctx context.Context, vpcID string, body *VPCUpdateRequest, opts ...basaltic.RequestOption) (*VPC, error) {
	op := &basaltic.Operation{
		ID:       "updateVpc",
		Method:   "PATCH",
		Path:     "/v1/vpcs/{vpc_id}",
		PathArgs: []string{vpcID},
		Body:     body,
	}
	var out struct {
		VPC *VPC `json:"vpc"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.VPC, nil
}

// GetEgressOnlyGatewayByReference fetches one egress only gateway by an
// id, a CRN or a name.
//
// The reference is classified by its syntax alone, exactly as the
// platform does (see [basaltic.ParseReference]): an id is fetched with
// [Client.GetEgressOnlyGateway]; a CRN or a name goes to
// [Client.ListEgressOnlyGateways] as an exact filter, together with any
// filters already set on scope, which may be nil. A miss is a not-found
// error for the kind the string was read as — no other kind is tried
// — and more than one match is a [basaltic.AmbiguousReferenceError].
func (c *Client) GetEgressOnlyGatewayByReference(ctx context.Context, ref string, scope *ListEgressOnlyGatewaysParams, opts ...basaltic.RequestOption) (*EgressOnlyGateway, error) {
	return basaltic.ResolveByReference(ctx, ref, "egress-only-gateway", "listEgressOnlyGateways", true,
		func(ctx context.Context, refID string) (*EgressOnlyGateway, error) {
			return c.GetEgressOnlyGateway(ctx, refID, opts...)
		},
		func(ctx context.Context, refName, refCRN string) (*basaltic.Page[EgressOnlyGateway], error) {
			var p ListEgressOnlyGatewaysParams
			if scope != nil {
				p = *scope
			}
			p.Name = refName
			p.CRN = refCRN
			p.Limit = 2
			return c.ListEgressOnlyGateways(ctx, &p, opts...)
		})
}

// GetFloatingIPByReference fetches one floating ip by an id, a CRN or a
// name.
//
// The reference is classified by its syntax alone, exactly as the
// platform does (see [basaltic.ParseReference]): an id is fetched with
// [Client.GetFloatingIP]; a CRN or a name goes to
// [Client.ListFloatingIPs] as an exact filter, together with any filters
// already set on scope, which may be nil. A miss is a not-found error
// for the kind the string was read as — no other kind is tried — and
// more than one match is a [basaltic.AmbiguousReferenceError].
//
// A name is unique only within its parent; fix it on scope (AttachedTo)
// or the lookup can match more than one.
func (c *Client) GetFloatingIPByReference(ctx context.Context, ref string, scope *ListFloatingIPsParams, opts ...basaltic.RequestOption) (*FloatingIP, error) {
	return basaltic.ResolveByReference(ctx, ref, "floating-ip", "listFloatingIps", true,
		func(ctx context.Context, refID string) (*FloatingIP, error) {
			return c.GetFloatingIP(ctx, refID, opts...)
		},
		func(ctx context.Context, refName, refCRN string) (*basaltic.Page[FloatingIP], error) {
			var p ListFloatingIPsParams
			if scope != nil {
				p = *scope
			}
			p.Name = refName
			p.CRN = refCRN
			p.Limit = 2
			return c.ListFloatingIPs(ctx, &p, opts...)
		})
}

// GetInterfaceByReference fetches one interface by an id, a CRN or a
// name.
//
// The reference is classified by its syntax alone, exactly as the
// platform does (see [basaltic.ParseReference]): an id is fetched with
// [Client.GetInterface]; a CRN or a name goes to [Client.ListInterfaces]
// as an exact filter, together with any filters already set on scope,
// which may be nil. A miss is a not-found error for the kind the string
// was read as — no other kind is tried — and more than one match is
// a [basaltic.AmbiguousReferenceError].
//
// A name is unique only within its parent; fix it on scope (Subnet, VPC)
// or the lookup can match more than one.
func (c *Client) GetInterfaceByReference(ctx context.Context, ref string, scope *ListInterfacesParams, opts ...basaltic.RequestOption) (*Interface, error) {
	return basaltic.ResolveByReference(ctx, ref, "interface", "listInterfaces", true,
		func(ctx context.Context, refID string) (*Interface, error) {
			return c.GetInterface(ctx, refID, opts...)
		},
		func(ctx context.Context, refName, refCRN string) (*basaltic.Page[Interface], error) {
			var p ListInterfacesParams
			if scope != nil {
				p = *scope
			}
			p.Name = refName
			p.CRN = refCRN
			p.Limit = 2
			return c.ListInterfaces(ctx, &p, opts...)
		})
}

// GetInternetGatewayByReference fetches one internet gateway by an id, a
// CRN or a name.
//
// The reference is classified by its syntax alone, exactly as the
// platform does (see [basaltic.ParseReference]): an id is fetched with
// [Client.GetInternetGateway]; a CRN or a name goes to
// [Client.ListInternetGateways] as an exact filter, together with any
// filters already set on scope, which may be nil. A miss is a not-found
// error for the kind the string was read as — no other kind is tried
// — and more than one match is a [basaltic.AmbiguousReferenceError].
func (c *Client) GetInternetGatewayByReference(ctx context.Context, ref string, scope *ListInternetGatewaysParams, opts ...basaltic.RequestOption) (*InternetGateway, error) {
	return basaltic.ResolveByReference(ctx, ref, "internet-gateway", "listInternetGateways", true,
		func(ctx context.Context, refID string) (*InternetGateway, error) {
			return c.GetInternetGateway(ctx, refID, opts...)
		},
		func(ctx context.Context, refName, refCRN string) (*basaltic.Page[InternetGateway], error) {
			var p ListInternetGatewaysParams
			if scope != nil {
				p = *scope
			}
			p.Name = refName
			p.CRN = refCRN
			p.Limit = 2
			return c.ListInternetGateways(ctx, &p, opts...)
		})
}

// GetNATGatewayByReference fetches one nat gateway by an id, a CRN or a
// name.
//
// The reference is classified by its syntax alone, exactly as the
// platform does (see [basaltic.ParseReference]): an id is fetched with
// [Client.GetNATGateway]; a CRN or a name goes to
// [Client.ListNATGateways] as an exact filter, together with any filters
// already set on scope, which may be nil. A miss is a not-found error
// for the kind the string was read as — no other kind is tried — and
// more than one match is a [basaltic.AmbiguousReferenceError].
//
// A name is unique only within its parent; fix it on scope (Subnet, VPC)
// or the lookup can match more than one.
func (c *Client) GetNATGatewayByReference(ctx context.Context, ref string, scope *ListNATGatewaysParams, opts ...basaltic.RequestOption) (*NATGateway, error) {
	return basaltic.ResolveByReference(ctx, ref, "nat-gateway", "listNATGateways", true,
		func(ctx context.Context, refID string) (*NATGateway, error) {
			return c.GetNATGateway(ctx, refID, opts...)
		},
		func(ctx context.Context, refName, refCRN string) (*basaltic.Page[NATGateway], error) {
			var p ListNATGatewaysParams
			if scope != nil {
				p = *scope
			}
			p.Name = refName
			p.CRN = refCRN
			p.Limit = 2
			return c.ListNATGateways(ctx, &p, opts...)
		})
}

// GetRouteByReference fetches one route by an id, a CRN or a name.
//
// The reference is classified by its syntax alone, exactly as the
// platform does (see [basaltic.ParseReference]): an id is fetched with
// [Client.GetRoute]; a CRN or a name goes to [Client.ListRoutes] as an
// exact filter, together with any filters already set on scope, which
// may be nil. A miss is a not-found error for the kind the string was
// read as — no other kind is tried — and more than one match is a
// [basaltic.AmbiguousReferenceError].
func (c *Client) GetRouteByReference(ctx context.Context, routeTableID string, ref string, scope *ListRoutesParams, opts ...basaltic.RequestOption) (*Route, error) {
	return basaltic.ResolveByReference(ctx, ref, "route", "listRoutes", true,
		func(ctx context.Context, refID string) (*Route, error) {
			return c.GetRoute(ctx, routeTableID, refID, opts...)
		},
		func(ctx context.Context, refName, refCRN string) (*basaltic.Page[Route], error) {
			var p ListRoutesParams
			if scope != nil {
				p = *scope
			}
			p.Name = refName
			p.CRN = refCRN
			p.Limit = 2
			return c.ListRoutes(ctx, routeTableID, &p, opts...)
		})
}

// GetRouteTableByReference fetches one route table by an id, a CRN or a
// name.
//
// The reference is classified by its syntax alone, exactly as the
// platform does (see [basaltic.ParseReference]): an id is fetched with
// [Client.GetRouteTable]; a CRN or a name goes to
// [Client.ListRouteTables] as an exact filter, together with any filters
// already set on scope, which may be nil. A miss is a not-found error
// for the kind the string was read as — no other kind is tried — and
// more than one match is a [basaltic.AmbiguousReferenceError].
//
// A name is unique only within its parent; fix it on scope (VPC) or the
// lookup can match more than one.
func (c *Client) GetRouteTableByReference(ctx context.Context, ref string, scope *ListRouteTablesParams, opts ...basaltic.RequestOption) (*RouteTable, error) {
	return basaltic.ResolveByReference(ctx, ref, "route-table", "listRouteTables", true,
		func(ctx context.Context, refID string) (*RouteTable, error) {
			return c.GetRouteTable(ctx, refID, opts...)
		},
		func(ctx context.Context, refName, refCRN string) (*basaltic.Page[RouteTable], error) {
			var p ListRouteTablesParams
			if scope != nil {
				p = *scope
			}
			p.Name = refName
			p.CRN = refCRN
			p.Limit = 2
			return c.ListRouteTables(ctx, &p, opts...)
		})
}

// GetSecurityGroupByReference fetches one security group by an id, a CRN
// or a name.
//
// The reference is classified by its syntax alone, exactly as the
// platform does (see [basaltic.ParseReference]): an id is fetched with
// [Client.GetSecurityGroup]; a CRN or a name goes to
// [Client.ListSecurityGroups] as an exact filter, together with any
// filters already set on scope, which may be nil. A miss is a not-found
// error for the kind the string was read as — no other kind is tried
// — and more than one match is a [basaltic.AmbiguousReferenceError].
func (c *Client) GetSecurityGroupByReference(ctx context.Context, ref string, scope *ListSecurityGroupsParams, opts ...basaltic.RequestOption) (*SecurityGroup, error) {
	return basaltic.ResolveByReference(ctx, ref, "security-group", "listSecurityGroups", true,
		func(ctx context.Context, refID string) (*SecurityGroup, error) {
			return c.GetSecurityGroup(ctx, refID, opts...)
		},
		func(ctx context.Context, refName, refCRN string) (*basaltic.Page[SecurityGroup], error) {
			var p ListSecurityGroupsParams
			if scope != nil {
				p = *scope
			}
			p.Name = refName
			p.CRN = refCRN
			p.Limit = 2
			return c.ListSecurityGroups(ctx, &p, opts...)
		})
}

// GetSecurityGroupRuleByReference fetches one security group rule by an
// id, a CRN or a name.
//
// The reference is classified by its syntax alone, exactly as the
// platform does (see [basaltic.ParseReference]): an id is fetched with
// [Client.GetSecurityGroupRule]; a CRN or a name goes to
// [Client.ListSecurityGroupRules] as an exact filter, together with any
// filters already set on scope, which may be nil. A miss is a not-found
// error for the kind the string was read as — no other kind is tried
// — and more than one match is a [basaltic.AmbiguousReferenceError].
func (c *Client) GetSecurityGroupRuleByReference(ctx context.Context, securityGroupID string, ref string, scope *ListSecurityGroupRulesParams, opts ...basaltic.RequestOption) (*SecurityGroupRule, error) {
	return basaltic.ResolveByReference(ctx, ref, "security-group-rule", "listSecurityGroupRules", true,
		func(ctx context.Context, refID string) (*SecurityGroupRule, error) {
			return c.GetSecurityGroupRule(ctx, securityGroupID, refID, opts...)
		},
		func(ctx context.Context, refName, refCRN string) (*basaltic.Page[SecurityGroupRule], error) {
			var p ListSecurityGroupRulesParams
			if scope != nil {
				p = *scope
			}
			p.Name = refName
			p.CRN = refCRN
			p.Limit = 2
			return c.ListSecurityGroupRules(ctx, securityGroupID, &p, opts...)
		})
}

// GetSubnetByReference fetches one subnet by an id, a CRN or a name.
//
// The reference is classified by its syntax alone, exactly as the
// platform does (see [basaltic.ParseReference]): an id is fetched with
// [Client.GetSubnet]; a CRN or a name goes to [Client.ListSubnets] as an
// exact filter, together with any filters already set on scope, which
// may be nil. A miss is a not-found error for the kind the string was
// read as — no other kind is tried — and more than one match is a
// [basaltic.AmbiguousReferenceError].
//
// A name is unique only within its parent; fix it on scope (VPC) or the
// lookup can match more than one.
func (c *Client) GetSubnetByReference(ctx context.Context, ref string, scope *ListSubnetsParams, opts ...basaltic.RequestOption) (*Subnet, error) {
	return basaltic.ResolveByReference(ctx, ref, "subnet", "listSubnets", true,
		func(ctx context.Context, refID string) (*Subnet, error) {
			return c.GetSubnet(ctx, refID, opts...)
		},
		func(ctx context.Context, refName, refCRN string) (*basaltic.Page[Subnet], error) {
			var p ListSubnetsParams
			if scope != nil {
				p = *scope
			}
			p.Name = refName
			p.CRN = refCRN
			p.Limit = 2
			return c.ListSubnets(ctx, &p, opts...)
		})
}

// GetVPCByReference fetches one vpc by an id, a CRN or a name.
//
// The reference is classified by its syntax alone, exactly as the
// platform does (see [basaltic.ParseReference]): an id is fetched with
// [Client.GetVPC]; a CRN or a name goes to [Client.ListVPCs] as an exact
// filter, together with any filters already set on scope, which may be
// nil. A miss is a not-found error for the kind the string was read as
// — no other kind is tried — and more than one match is a
// [basaltic.AmbiguousReferenceError].
func (c *Client) GetVPCByReference(ctx context.Context, ref string, scope *ListVPCsParams, opts ...basaltic.RequestOption) (*VPC, error) {
	return basaltic.ResolveByReference(ctx, ref, "vpc", "listVpcs", true,
		func(ctx context.Context, refID string) (*VPC, error) {
			return c.GetVPC(ctx, refID, opts...)
		},
		func(ctx context.Context, refName, refCRN string) (*basaltic.Page[VPC], error) {
			var p ListVPCsParams
			if scope != nil {
				p = *scope
			}
			p.Name = refName
			p.CRN = refCRN
			p.Limit = 2
			return c.ListVPCs(ctx, &p, opts...)
		})
}
