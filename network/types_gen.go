// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package network

import (
	"time"
)

type AttachFloatingIPRequest struct {
	// Required.
	InterfaceID string `json:"interface_id"`
}

type DetachFloatingIPRequest struct {
	// InterfaceID the member to remove. Naming one member of a shared address removes
	// just that member and leaves the address serving from the rest;
	// omitting it detaches every member. On an address with a single
	// member the two are the same thing. Naming a NIC that is not a member
	// is a no-op.
	InterfaceID *string `json:"interface_id,omitempty"`
}

// EgressOnlyGateway The IPv6 analogue of a NAT gateway, and its inverse: a subnet whose
// route table points ::/0 at one gets OUTBOUND v6 (plus the return
// traffic of its own flows), but the internet can never initiate an
// inbound connection — a platform-band drop enforces that regardless
// of the tenant's security groups. It owns no address (v6 has no NAT)
// and reuses the VPC's internet gateway for the L3 uplink, so the VPC
// must have an IGW attached. One per VPC.
type EgressOnlyGateway struct {
	CreatedAt   time.Time         `json:"created_at"`
	CRN         string            `json:"crn"`
	Description string            `json:"description,omitempty"`
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Tags        map[string]string `json:"tags"`
	UpdatedAt   time.Time         `json:"updated_at"`
	VPCID       string            `json:"vpc_id"`
}

type EgressOnlyGatewayCreateRequest struct {
	Description *string `json:"description,omitempty"`

	// Required.
	Name string            `json:"name"`
	Tags map[string]string `json:"tags,omitempty"`

	// VPCID The VPC must have an IPv6 CIDR (an egress-only gateway only routes
	// v6).
	//
	// Required.
	VPCID string `json:"vpc_id"`
}

type EgressOnlyGatewayUpdateRequest struct {
	Description *string           `json:"description,omitempty"`
	Tags        map[string]string `json:"tags,omitempty"`
}

type FloatingIP struct {
	// AttachedToInterfaceID legacy single-binding field: the sole interface this floating IP is
	// bound to, or null when unattached OR when it has more than one
	// member (an anycast floating IP). `members` is authoritative.
	AttachedToInterfaceID string    `json:"attached_to_interface_id,omitempty"`
	CreatedAt             time.Time `json:"created_at"`
	CRN                   string    `json:"crn"`
	Description           string    `json:"description,omitempty"`
	ID                    string    `json:"id"`

	// InstancePoolID the instance pool this address belongs to, or null for an ordinary
	// floating IP.
	//
	// A pool's address is the only one that can have more than one member.
	// Its members are the pool's live replicas — one per hypervisor,
	// maintained by the pool as it scales — so `attach` and `detach` on
	// this floating IP are refused: use `POST
	// /v1/instance-pools/{pool_id}/floating-ips` and `DELETE
	// /v1/instance-pools/{pool_id}/floating-ips/{floating_ip_id}`.
	InstancePoolID string `json:"instance_pool_id,omitempty"`
	IPAddress      string `json:"ip_address"`

	// Members the floating IP's bindings. A floating IP fronts 0 members
	// (allocated, unattached), 1 member (the everyday case), or N members
	// — an anycast floating IP, where one public IP is delivered to N VM
	// NICs across hosts (each advertised as a /32 from the host holding
	// it).
	//
	// Members may share a hypervisor. Two of them on one host used to mean
	// one served and the other was silently dark; a member's forwarding
	// rule now names the member, and the host splits connections across
	// the members it holds, so where the members sit is a capacity
	// decision rather than a correctness one. An instance pool's address
	// takes its members from the pool's live replicas — every one of
	// them — so a scale-out joins and a scale-in leaves without a
	// per-replica attach.
	//
	// With more than one member ONE member serves each connection, chosen
	// by hashing the flow's addresses and ports, and every packet of that
	// connection goes to the same one. The members are separate instances
	// that share nothing, so this spreads connections and survives the
	// loss of a host — it is not a load balancer: nothing checks whether
	// the service inside the instance is up, and connections in progress
	// to a member that goes away are not moved, they end.
	//
	// A POOL's address is the exception, and only for booting. A replica
	// joins the address as soon as it is placed, but does not receive
	// traffic until it has reached the instance metadata service —
	// evidence that the guest booted, rather than that its virtual machine
	// was started. Until then it is a member with `health` `unhealthy`. A
	// replica whose image never contacts the metadata service is admitted
	// anyway after a few minutes, so an unusual image delays traffic
	// rather than never getting it.
	Members   []*FloatingIPMember `json:"members"`
	Tags      map[string]string   `json:"tags"`
	UpdatedAt time.Time           `json:"updated_at"`
}

type FloatingIPCreateRequest struct {
	Description *string           `json:"description,omitempty"`
	Tags        map[string]string `json:"tags,omitempty"`
}

// FloatingIPMember one binding of a floating IP.
type FloatingIPMember struct {
	CreatedAt time.Time `json:"created_at"`

	// Health what the platform knows about this member.
	//
	// `unknown` — nobody is checking. Every member you attached yourself
	// reads this: you chose the moment of attach, and the platform has no
	// signal about what runs inside the instance. It is advertised.
	//
	// `healthy` — the platform has evidence this member is up. An
	// instance pool's replica reads this once it has reached the instance
	// metadata service.
	//
	// `unhealthy` — the platform is waiting for that evidence and has
	// not seen it. The member keeps its place on the address and receives
	// no traffic until it does. A pool replica reads this while it is
	// still booting.
	//
	// This is liveness, not readiness: `healthy` means the guest came up,
	// not that your service is listening on it.
	//
	// One of: "unknown", "healthy", "unhealthy".
	Health string `json:"health"`

	// InterfaceID the bound interface (instance_nic floating IPs).
	InterfaceID string `json:"interface_id,omitempty"`

	// ResourceID the bound resource id (lb / email_sender floating IPs).
	ResourceID string `json:"resource_id,omitempty"`
}

type FloatingIPUpdateRequest struct {
	Description *string           `json:"description,omitempty"`
	Tags        map[string]string `json:"tags,omitempty"`
}

type Interface struct {
	// AttachedTo Resource CRN this interface is attached to (e.g. a compute
	// instance). Null until VM-attach lands with the compute rewrite.
	AttachedTo  string    `json:"attached_to,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	CRN         string    `json:"crn"`
	Description string    `json:"description,omitempty"`
	ID          string    `json:"id"`
	IPAddress   string    `json:"ip_address"`

	// IPv6Address the interface's /128, auto-assigned when its subnet is dual-stack.
	IPv6Address string            `json:"ipv6_address,omitempty"`
	MAC         string            `json:"mac"`
	Name        string            `json:"name"`
	SubnetID    string            `json:"subnet_id"`
	Tags        map[string]string `json:"tags"`
	UpdatedAt   time.Time         `json:"updated_at"`
	VPCID       string            `json:"vpc_id"`
}

type InterfaceCreateRequest struct {
	Description *string `json:"description,omitempty"`

	// IPAddress defaults to the next free address in the subnet
	IPAddress *string `json:"ip_address,omitempty"`

	// MAC defaults to a fresh locally-administered EUI-48
	MAC *string `json:"mac,omitempty"`

	// Required.
	Name string `json:"name"`

	// Required.
	SubnetID string            `json:"subnet_id"`
	Tags     map[string]string `json:"tags,omitempty"`
}

type InterfaceSecurityGroupsRequest struct {
	// Required.
	SecurityGroupIDs []string `json:"security_group_ids"`
}

type InterfaceUpdateRequest struct {
	Description *string           `json:"description,omitempty"`
	Tags        map[string]string `json:"tags,omitempty"`
}

type InternetGateway struct {
	// AttachedVPCID VPC the IGW is currently attached to (null when detached).
	AttachedVPCID string            `json:"attached_vpc_id,omitempty"`
	CreatedAt     time.Time         `json:"created_at"`
	CRN           string            `json:"crn"`
	Description   string            `json:"description,omitempty"`
	ID            string            `json:"id"`
	Name          string            `json:"name"`
	Tags          map[string]string `json:"tags"`
	UpdatedAt     time.Time         `json:"updated_at"`
}

type InternetGatewayAttachRequest struct {
	// VPCID VPC to attach this IGW to. At most one IGW per VPC.
	//
	// Required.
	VPCID string `json:"vpc_id"`
}

type InternetGatewayCreateRequest struct {
	Description *string `json:"description,omitempty"`

	// Required.
	Name string            `json:"name"`
	Tags map[string]string `json:"tags,omitempty"`
}

type InternetGatewayUpdateRequest struct {
	Description *string           `json:"description,omitempty"`
	Tags        map[string]string `json:"tags,omitempty"`
}

type NATGateway struct {
	CreatedAt   time.Time `json:"created_at"`
	CRN         string    `json:"crn"`
	Description string    `json:"description,omitempty"`

	// ExternalIP Public IP allocated from the regional pool. Stable for the NAT GW's
	// lifetime.
	ExternalIP string `json:"external_ip"`
	ID         string `json:"id"`
	Name       string `json:"name"`

	// SubnetID subnet the NAT GW lives in. Subnet delete is blocked while occupied.
	SubnetID  string            `json:"subnet_id"`
	Tags      map[string]string `json:"tags"`
	UpdatedAt time.Time         `json:"updated_at"`

	// VPCID VPC of the parent subnet (denormalised for convenience).
	VPCID string `json:"vpc_id"`
}

type NATGatewayCreateRequest struct {
	Description *string `json:"description,omitempty"`

	// Required.
	Name string `json:"name"`

	// SubnetID subnet the NAT GW lives in. The parent VPC must already have an
	// Internet Gateway attached — the NAT GW reuses that uplink.
	//
	// Required.
	SubnetID string            `json:"subnet_id"`
	Tags     map[string]string `json:"tags,omitempty"`
}

type NATGatewayUpdateRequest struct {
	Description *string           `json:"description,omitempty"`
	Tags        map[string]string `json:"tags,omitempty"`
}

type Route struct {
	CreatedAt    time.Time         `json:"created_at"`
	CRN          string            `json:"crn"`
	Description  string            `json:"description,omitempty"`
	Destination  string            `json:"destination"`
	ID           string            `json:"id"`
	RouteTableID string            `json:"route_table_id"`
	Tags         map[string]string `json:"tags"`

	// TargetEgressOnlyGatewayID set when target_type=egress_only_gateway (IPv6 only). Gives the
	// subnet outbound v6 with the internet unable to initiate inbound.
	TargetEgressOnlyGatewayID string `json:"target_egress_only_gateway_id,omitempty"`

	// TargetInternetGatewayID set when target_type=internet_gateway.
	TargetInternetGatewayID string `json:"target_internet_gateway_id,omitempty"`

	// TargetIP set when target_type=ip. Mutex with the target_*_id fields. Must be
	// a unicast address inside this VPC's CIDR (same IP family as
	// destination); internet egress uses target_internet_gateway_id /
	// target_nat_gateway_id.
	TargetIP string `json:"target_ip,omitempty"`

	// TargetNATGatewayID set when target_type=nat_gateway.
	TargetNATGatewayID string          `json:"target_nat_gateway_id,omitempty"`
	TargetType         RouteTargetType `json:"target_type"`
	UpdatedAt          time.Time       `json:"updated_at"`
}

// RouteCreateRequest exactly one of target_ip / target_internet_gateway_id /
// target_nat_gateway_id / target_egress_only_gateway_id (and future
// target_*_id fields) must be set.
type RouteCreateRequest struct {
	Description *string `json:"description,omitempty"`

	// Required.
	Destination string            `json:"destination"`
	Tags        map[string]string `json:"tags,omitempty"`

	// TargetEgressOnlyGatewayID IPv6-only. Outbound v6 with no internet-initiated inbound.
	TargetEgressOnlyGatewayID *string `json:"target_egress_only_gateway_id,omitempty"`
	TargetInternetGatewayID   *string `json:"target_internet_gateway_id,omitempty"`

	// TargetIP unicast next hop inside this VPC's CIDR (same IP family as
	// destination). Not for internet egress — use a gateway target id
	// instead.
	TargetIP           *string `json:"target_ip,omitempty"`
	TargetNATGatewayID *string `json:"target_nat_gateway_id,omitempty"`
}

type RouteTable struct {
	CreatedAt   time.Time `json:"created_at"`
	CRN         string    `json:"crn"`
	Description string    `json:"description,omitempty"`
	ID          string    `json:"id"`

	// IsMain true for the per-VPC default table. The main table is created
	// automatically and can't be deleted. Subnets that don't specify a
	// route_table_id at create time land here.
	IsMain    bool              `json:"is_main"`
	Name      string            `json:"name"`
	Tags      map[string]string `json:"tags"`
	UpdatedAt time.Time         `json:"updated_at"`
	VPCID     string            `json:"vpc_id"`
}

type RouteTableCreateRequest struct {
	Description *string `json:"description,omitempty"`

	// Name 1-63 chars, lowercase alphanumeric + hyphen. `main` is reserved.
	//
	// Required.
	Name string            `json:"name"`
	Tags map[string]string `json:"tags,omitempty"`

	// VPCID VPC this table belongs to.
	//
	// Required.
	VPCID string `json:"vpc_id"`
}

type RouteTableUpdateRequest struct {
	Description *string           `json:"description,omitempty"`
	Tags        map[string]string `json:"tags,omitempty"`
}

// RouteTargetType discriminator for the route's target. New target types (interface,
// vpc_peering, …) extend this enum and add their own target_* field.
// Mirrors AWS one-field-per-target style — exactly one target_*
// property must be set on create.
type RouteTargetType string

// Values RouteTargetType accepts.
const (
	RouteTargetTypeIP                RouteTargetType = "ip"
	RouteTargetTypeInternetGateway   RouteTargetType = "internet_gateway"
	RouteTargetTypeNATGateway        RouteTargetType = "nat_gateway"
	RouteTargetTypeEgressOnlyGateway RouteTargetType = "egress_only_gateway"
)

type RouteUpdateRequest struct {
	Description *string           `json:"description,omitempty"`
	Tags        map[string]string `json:"tags,omitempty"`
}

type SecurityGroup struct {
	CreatedAt   time.Time         `json:"created_at"`
	CRN         string            `json:"crn"`
	Description string            `json:"description,omitempty"`
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Tags        map[string]string `json:"tags"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

type SecurityGroupCreateRequest struct {
	Description *string `json:"description,omitempty"`

	// Required.
	Name string            `json:"name"`
	Tags map[string]string `json:"tags,omitempty"`
}

type SecurityGroupRule struct {
	CreatedAt   time.Time                  `json:"created_at"`
	Description string                     `json:"description,omitempty"`
	Direction   SecurityGroupRuleDirection `json:"direction"`
	Ethertype   SecurityGroupRuleEthertype `json:"ethertype"`
	ID          string                     `json:"id"`
	PortMax     int                        `json:"port_max,omitempty"`

	// PortMin required when protocol is tcp/udp; ignored otherwise.
	PortMin         int                       `json:"port_min,omitempty"`
	Protocol        SecurityGroupRuleProtocol `json:"protocol"`
	SecurityGroupID string                    `json:"security_group_id"`

	// SourceCIDR source (ingress) or destination (egress) CIDR. Must match the rule's
	// ethertype. Mutually exclusive with source_security_group_id.
	SourceCIDR string `json:"source_cidr,omitempty"`

	// SourceSecurityGroupID source (ingress) or destination (egress) is "any workload in this
	// SG". Traffic is matched by membership in the named security group.
	// Mutually exclusive with source_cidr.
	SourceSecurityGroupID string `json:"source_security_group_id,omitempty"`
}

type SecurityGroupRuleCreateRequest struct {
	Description *string `json:"description,omitempty"`

	// Required.
	Direction SecurityGroupRuleDirection  `json:"direction"`
	Ethertype *SecurityGroupRuleEthertype `json:"ethertype,omitempty"`
	PortMax   *int                        `json:"port_max,omitempty"`
	PortMin   *int                        `json:"port_min,omitempty"`

	// Required.
	Protocol              SecurityGroupRuleProtocol `json:"protocol"`
	SourceCIDR            *string                   `json:"source_cidr,omitempty"`
	SourceSecurityGroupID *string                   `json:"source_security_group_id,omitempty"`
}

type SecurityGroupRuleDirection string

// Values SecurityGroupRuleDirection accepts.
const (
	SecurityGroupRuleDirectionIngress SecurityGroupRuleDirection = "ingress"
	SecurityGroupRuleDirectionEgress  SecurityGroupRuleDirection = "egress"
)

type SecurityGroupRuleEthertype string

// Values SecurityGroupRuleEthertype accepts.
const (
	SecurityGroupRuleEthertypeIPv4 SecurityGroupRuleEthertype = "ipv4"
	SecurityGroupRuleEthertypeIPv6 SecurityGroupRuleEthertype = "ipv6"
)

type SecurityGroupRuleProtocol string

// Values SecurityGroupRuleProtocol accepts.
const (
	SecurityGroupRuleProtocolTcp  SecurityGroupRuleProtocol = "tcp"
	SecurityGroupRuleProtocolUdp  SecurityGroupRuleProtocol = "udp"
	SecurityGroupRuleProtocolICMP SecurityGroupRuleProtocol = "icmp"
	SecurityGroupRuleProtocolAll  SecurityGroupRuleProtocol = "all"
)

type SecurityGroupUpdateRequest struct {
	Description *string           `json:"description,omitempty"`
	Tags        map[string]string `json:"tags,omitempty"`
}

type Subnet struct {
	CIDR string `json:"cidr"`

	// CIDRV6 the dual-stack IPv6 /64, if the subnet is v6-enabled. Its presence
	// (vs the v4 cidr) is how a client tells the subnet's families apart.
	CIDRV6      string    `json:"cidr_v6,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	CRN         string    `json:"crn"`
	Description string    `json:"description,omitempty"`
	GatewayIP   string    `json:"gateway_ip"`
	GatewayIPV6 string    `json:"gateway_ip_v6,omitempty"`
	ID          string    `json:"id"`
	Name        string    `json:"name"`

	// RouteTableID route table this subnet uses. Determines public/private semantics
	// — a subnet is "public" if its route table has a 0.0.0.0/0 route
	// pointing at an internet gateway, "private" otherwise. Defaults to
	// the VPC's main table.
	RouteTableID string            `json:"route_table_id"`
	Tags         map[string]string `json:"tags"`
	UpdatedAt    time.Time         `json:"updated_at"`
	VPCID        string            `json:"vpc_id"`
}

type SubnetCreateRequest struct {
	// Required.
	CIDR string `json:"cidr"`

	// CIDRV6 makes the subnet dual-stack. A /64 inside the VPC's IPv6 CIDR (the
	// VPC must have been created with assign_ipv6_cidr). Omit for a
	// v4-only subnet.
	CIDRV6      *string `json:"cidr_v6,omitempty"`
	Description *string `json:"description,omitempty"`

	// GatewayIP defaults to the first usable host in the CIDR
	GatewayIP *string `json:"gateway_ip,omitempty"`

	// Required.
	Name string `json:"name"`

	// RouteTableID defaults to the VPC's main route table
	RouteTableID *string           `json:"route_table_id,omitempty"`
	Tags         map[string]string `json:"tags,omitempty"`

	// Required.
	VPCID string `json:"vpc_id"`
}

type SubnetUpdateRequest struct {
	Description *string `json:"description,omitempty"`

	// RouteTableID re-associate the subnet with a different route table
	RouteTableID *string           `json:"route_table_id,omitempty"`
	Tags         map[string]string `json:"tags,omitempty"`
}

type VPC struct {
	// CIDRV4 IPv4 CIDR block carved up by subnets. Must be private (RFC 1918):
	// within 10.0.0.0/8, 172.16.0.0/12 or 192.168.0.0/16. Immutable after
	// create.
	CIDRV4 string `json:"cidr_v4"`

	// CIDRV6 the globally-routable /60 delegated from the region's IPv6 pool when
	// the VPC was created with assign_ipv6_cidr; null for v4-only VPCs.
	// Immutable after create.
	CIDRV6    string    `json:"cidr_v6,omitempty"`
	CreatedAt time.Time `json:"created_at"`

	// CRN Cloud Resource Name (name-based, region+account-scoped).
	CRN         string `json:"crn"`
	Description string `json:"description,omitempty"`
	ID          string `json:"id"`

	// Name 1-63 chars, lowercase alphanumeric + hyphen
	Name      string            `json:"name"`
	Tags      map[string]string `json:"tags"`
	UpdatedAt time.Time         `json:"updated_at"`
}

type VPCCreateRequest struct {
	// AssignIPv6CIDR request a globally-routable /60 delegated from the region's IPv6
	// pool — the only way a VPC gets an IPv6 prefix; the region must
	// have IPv6 enabled. The delegated prefix is returned as the VPC's
	// cidr_v6.
	AssignIPv6CIDR *bool `json:"assign_ipv6_cidr,omitempty"`

	// CIDRV4 must be private (RFC 1918): within 10.0.0.0/8, 172.16.0.0/12 or
	// 192.168.0.0/16.
	//
	// Required.
	CIDRV4      string  `json:"cidr_v4"`
	Description *string `json:"description,omitempty"`

	// Required.
	Name string            `json:"name"`
	Tags map[string]string `json:"tags,omitempty"`
}

type VPCUpdateRequest struct {
	Description *string           `json:"description,omitempty"`
	Tags        map[string]string `json:"tags,omitempty"`
}
