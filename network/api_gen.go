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

// ListEgressOnlyGatewaysParams are the optional filters and pagination controls for
// [Client.ListEgressOnlyGateways]. A nil *ListEgressOnlyGatewaysParams sends none of them.
type ListEgressOnlyGatewaysParams struct {
	Limit int

	// Marker resume token — the last id from the previous page.
	Marker string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListEgressOnlyGatewaysParams) query() url.Values {
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
func (p *ListEgressOnlyGatewaysParams) withMarker(marker string) *ListEgressOnlyGatewaysParams {
	var out ListEgressOnlyGatewaysParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListFloatingIPsParams are the optional filters and pagination controls for
// [Client.ListFloatingIPs]. A nil *ListFloatingIPsParams sends none of them.
type ListFloatingIPsParams struct {
	Limit int

	// Marker resume token — the last id from the previous page.
	Marker string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListFloatingIPsParams) query() url.Values {
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
func (p *ListFloatingIPsParams) withMarker(marker string) *ListFloatingIPsParams {
	var out ListFloatingIPsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListInterfacesParams are the optional filters and pagination controls for
// [Client.ListInterfaces]. A nil *ListInterfacesParams sends none of them.
type ListInterfacesParams struct {
	Limit int

	// Marker resume token — the last id from the previous page.
	Marker   string
	SubnetID string
	VPCID    string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListInterfacesParams) query() url.Values {
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
	if p.SubnetID != "" {
		q.Set("subnet_id", p.SubnetID)
	}
	if p.VPCID != "" {
		q.Set("vpc_id", p.VPCID)
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

// ListInternetGatewaysParams are the optional filters and pagination controls for
// [Client.ListInternetGateways]. A nil *ListInternetGatewaysParams sends none of them.
type ListInternetGatewaysParams struct {
	Limit int

	// Marker resume token — the last id from the previous page.
	Marker string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListInternetGatewaysParams) query() url.Values {
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
func (p *ListInternetGatewaysParams) withMarker(marker string) *ListInternetGatewaysParams {
	var out ListInternetGatewaysParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListNATGatewaysParams are the optional filters and pagination controls for
// [Client.ListNATGateways]. A nil *ListNATGatewaysParams sends none of them.
type ListNATGatewaysParams struct {
	Limit int

	// Marker resume token — the last id from the previous page.
	Marker   string
	SubnetID string
	VPCID    string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListNATGatewaysParams) query() url.Values {
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
	if p.SubnetID != "" {
		q.Set("subnet_id", p.SubnetID)
	}
	if p.VPCID != "" {
		q.Set("vpc_id", p.VPCID)
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
	Limit int

	// Marker resume token — the last id from the previous page.
	Marker string

	// VPCID filter by VPC ID
	VPCID string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListRouteTablesParams) query() url.Values {
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
	if p.VPCID != "" {
		q.Set("vpc_id", p.VPCID)
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
	Limit int

	// Marker resume token — the last id from the previous page.
	Marker string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListRoutesParams) query() url.Values {
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
	Limit int

	// Marker resume token — the last id from the previous page.
	Marker string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListSecurityGroupRulesParams) query() url.Values {
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
	Limit int

	// Marker resume token — the last id from the previous page.
	Marker string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListSecurityGroupsParams) query() url.Values {
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
	Limit int

	// Marker resume token — the last id from the previous page.
	Marker string

	// VPCID filter by VPC ID
	VPCID string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListSubnetsParams) query() url.Values {
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
	if p.VPCID != "" {
		q.Set("vpc_id", p.VPCID)
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
	Limit int

	// Marker resume token — the last id from the previous page.
	Marker string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListVPCsParams) query() url.Values {
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
func (p *ListVPCsParams) withMarker(marker string) *ListVPCsParams {
	var out ListVPCsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// AttachFloatingIP attaches a floating IP to an interface.
//
// Binds the floating IP to the given interface and routes the public IP
// to the instance behind it. Idempotent re-attach to the same iface is a
// no-op.
//
// SEVERAL INTERFACES MAKE IT AN ANYCAST ADDRESS. Attaching a second
// interface is allowed: every member answers on the same public address,
// and a connection is delivered to one of them. Members may share a
// hypervisor — the host splits connections across the members it holds
// and the edge splits across hypervisors — so this is a capacity
// decision, not a correctness one. Losing a member ends the connections
// it was serving; the address itself keeps working from the others.
//
// Every member must be in the SAME VPC (409 otherwise): one router
// carries the address, so a member in another VPC would have no return
// path.
//
// Members are not health-checked. `members[].health` reads `unknown` for
// instance NICs, and an instance whose application has died keeps its
// share of the connections until it is detached.
//
// An instance pool can own the address instead (`POST
// /v1/instance-pools/{pool_id}/floating-ips`), which keeps the member
// set in step with the pool as it scales rather than leaving you to
// re-attach by hand. An address a pool already owns is refused here —
// attach and detach it at the pool. A floating IP bound to a load
// balancer / email sender can't take NIC members either (409).
//
// The interface's subnet must already route `0.0.0.0/0` to an internet
// gateway, or the address would be handed back unreachable (400).
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) AttachFloatingIP(ctx context.Context, floatingIPID string, body *AttachFloatingIPRequest, opts ...basaltic.RequestOption) (*FloatingIP, error) {
	op := &basaltic.Operation{
		ID:       "attachFloatingIp",
		Method:   "POST",
		Path:     "/v1/floating-ips/{floating_ip_id}/attach",
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

// AttachInternetGateway attaches internet gateway to a VPC.
//
// Allocates an address from the regional external pool and connects the
// VPC's router to the regional external network. AWS rule: each VPC can
// have at most one IGW attached at a time.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) AttachInternetGateway(ctx context.Context, internetGatewayID string, body *InternetGatewayAttachRequest, opts ...basaltic.RequestOption) (*InternetGateway, error) {
	op := &basaltic.Operation{
		ID:       "attachInternetGateway",
		Method:   "POST",
		Path:     "/v1/internet-gateways/{internet_gateway_id}/attach",
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
// Allocate the next free public IPv4 from the region's configured pool.
// The IP is reserved against the caller's account but not yet bound to
// anything.
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

// CreateInternetGateway creates internet gateway.
//
// Creates an IGW in the detached state. Use POST /attach to hook it up
// to a VPC.
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
// Allocates an EIP from the regional public pool and configures source
// NAT for the VPC's private subnets through the internet gateway's
// uplink. The target VPC must already have an Internet Gateway attached.
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
// Return the IP to the pool. Refuses if the floating IP is still
// attached to an interface — detach first.
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

// DeleteInternetGateway deletes internet gateway.
//
// Refuses if still attached or if any route references it.
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

// DetachFloatingIP detaches a floating IP.
//
// Stops routing the public IP and clears the binding. Idempotent —
// detaching an already-detached floating IP (or a NIC that isn't a
// member) returns 200 with the unchanged row.
//
// An address an instance pool owns is refused (409): its members are
// that pool's live replicas, so removing one here would be reinstated by
// the pool's next reconcile. Detach it at `DELETE
// /v1/instance-pools/{pool_id}/floating-ips/{floating_ip_id}`.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) DetachFloatingIP(ctx context.Context, floatingIPID string, body *DetachFloatingIPRequest, opts ...basaltic.RequestOption) (*FloatingIP, error) {
	op := &basaltic.Operation{
		ID:       "detachFloatingIp",
		Method:   "POST",
		Path:     "/v1/floating-ips/{floating_ip_id}/detach",
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

// DetachInternetGateway detaches internet gateway from its VPC.
//
// Disconnects the VPC's router from the external network and returns the
// external IP to the pool. Refuses if any route still references the
// IGW.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) DetachInternetGateway(ctx context.Context, internetGatewayID string, opts ...basaltic.RequestOption) (*InternetGateway, error) {
	op := &basaltic.Operation{
		ID:       "detachInternetGateway",
		Method:   "POST",
		Path:     "/v1/internet-gateways/{internet_gateway_id}/detach",
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

// ListInterfaceSecurityGroups lists interface security-group membership.
func (c *Client) ListInterfaceSecurityGroups(ctx context.Context, interfaceID string, opts ...basaltic.RequestOption) (*basaltic.Page[string], error) {
	op := &basaltic.Operation{
		ID:       "listInterfaceSecurityGroups",
		Method:   "GET",
		Path:     "/v1/interfaces/{interface_id}/security-groups",
		PathArgs: []string{interfaceID},
	}
	var out struct {
		Items []string `json:"security_group_ids"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[string]{Items: out.Items}
	return page, nil
}

// ListInterfaces lists interfaces.
//
// List interfaces owned by the caller's account. Filter by `subnet_id`
// to narrow to one subnet, or by `vpc_id` to span every subnet in one
// VPC. subnet_id wins if both are supplied.
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

// ListNATGateways lists NAT gateways.
//
// List NAT gateways owned by the caller's account. Filter by `subnet_id`
// to narrow to the gateway living in one subnet, or by `vpc_id` to span
// the VPC. subnet_id wins if both are supplied.
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

// ListRouteTables lists route tables.
//
// List route tables owned by the caller's account. Filter by vpc_id to
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
// List subnets owned by the caller's account. Filter by vpc_id to narrow
// to one VPC.
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

// SetInterfaceSecurityGroups sets interface security-group membership.
//
// Atomically replaces the interface's SG membership with the supplied
// list. Adds and removes are diffed against current state. Each ID must
// reference an SG owned by the caller.
func (c *Client) SetInterfaceSecurityGroups(ctx context.Context, interfaceID string, body *InterfaceSecurityGroupsRequest, opts ...basaltic.RequestOption) (*basaltic.Page[string], error) {
	op := &basaltic.Operation{
		ID:       "setInterfaceSecurityGroups",
		Method:   "PUT",
		Path:     "/v1/interfaces/{interface_id}/security-groups",
		PathArgs: []string{interfaceID},
		Body:     body,
	}
	var out struct {
		Items []string `json:"security_group_ids"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[string]{Items: out.Items}
	return page, nil
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
