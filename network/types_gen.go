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
	// Interface UUID or nested CRN. Bare names have no subnet scope and
	// are rejected.
	//
	// Required.
	Interface string `json:"interface"`
}

type DetachFloatingIPRequest struct {
	// Interface UUID or nested CRN; bare names, null and empty references
	// are rejected. Omitting the field clears the binding. Naming a NIC
	// that is not a member is a no-op.
	Interface *string `json:"interface,omitempty"`
}

// EgressOnlyGateway The IPv6 analogue of a NAT gateway, and its inverse: a subnet whose
// route table points ::/0 at one gets OUTBOUND v6 (plus the return
// traffic of its own flows), but the internet can never initiate an
// inbound connection — a platform-band drop enforces that regardless
// of the tenant's security groups. It owns no address (v6 has no NAT)
// and reuses the VPC's internet gateway for the L3 uplink, so the VPC
// must have an IGW attached. One per VPC.
type EgressOnlyGateway struct {
	CreatedAt   time.Time `json:"created_at"`
	CRN         string    `json:"crn"`
	Description string    `json:"description,omitempty"`
	ID          string    `json:"id"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	Name      string            `json:"name"`
	Tags      map[string]string `json:"tags"`
	UpdatedAt time.Time         `json:"updated_at"`
	VPC       *VPC              `json:"vpc"`
}

type EgressOnlyGatewayCreateRequest struct {
	Description *string `json:"description,omitempty"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	//
	// Required.
	Name string            `json:"name"`
	Tags map[string]string `json:"tags,omitempty"`

	// VPC UUID, CRN or exact name in the caller account.
	//
	// Required.
	VPC string `json:"vpc"`
}

type EgressOnlyGatewayUpdateRequest struct {
	Description *string           `json:"description,omitempty"`
	Tags        map[string]string `json:"tags,omitempty"`
}

type FloatingIP struct {
	// AttachedTo Canonical CRN of the bound interface, instance pool, or load
	// balancer; null when unattached. A pool-owned address names its pool
	// even when the pool has zero members. Only pool-owned addresses may
	// have multiple NIC members. Manage their bindings through the
	// instance pool floating IP endpoints; direct attach and detach are
	// refused.
	AttachedTo  string    `json:"attached_to"`
	CreatedAt   time.Time `json:"created_at"`
	CRN         string    `json:"crn"`
	Description string    `json:"description,omitempty"`
	Family      IPFamily  `json:"family"`

	// HealthCheck the readiness check applied to this address's members. Absent when
	// none is configured. See `FloatingIpHealthCheck`.
	HealthCheck *FloatingIPHealthCheck `json:"health_check,omitempty"`
	ID          string                 `json:"id"`

	// IPAddress the public address, in the family it was allocated in.
	IPAddress string `json:"ip_address"`

	// Members the floating IP's bindings. A floating IP fronts 0 members
	// (allocated, unattached), 1 member (the everyday case), or N members
	// for an instance pool — an anycast floating IP, where one public IP
	// is delivered to N VM NICs across hosts (each advertised as a /32
	// from the host holding it).
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
	Description *string `json:"description,omitempty"`

	// Family which family to allocate in. Fixed for the life of the address —
	// it decides the pool the address comes from, the quota it counts
	// against (`floating_ips_v4` or `floating_ips_v6`) and the SKU it
	// bills as. Omitted means `ipv4`.
	Family *IPFamily `json:"family,omitempty"`

	// HealthCheck an optional readiness check for the address's members. Omitted means
	// none — the address behaves exactly as an ordinary floating IP.
	HealthCheck *FloatingIPHealthCheck `json:"health_check,omitempty"`
	Tags        map[string]string      `json:"tags,omitempty"`
}

// FloatingIPHealthCheck a readiness check for a shared (anycast) floating IP's members — the
// same vocabulary as a load balancer target group's health check, one
// you already know. The platform checks each member's private address
// within the VPC. A member that fails stops receiving traffic through
// the floating IP and returns when it passes again. If EVERY member
// fails, the whole address goes dark — a misconfigured check is a
// visible outage you caused, not the platform quietly advertising
// something it believes is down.
//
// The check is on the address, not per member: members are
// interchangeable backends, and a pool derives them. An address with no
// check behaves exactly as before — liveness only for pool members,
// always-advertised for hand-attached ones.
type FloatingIPHealthCheck struct {
	// HealthyThreshold consecutive passes before a member flips healthy.
	HealthyThreshold int `json:"healthy_threshold"`
	IntervalSec      int `json:"interval_sec"`

	// Matcher HTTP status or range that counts as passing; ignored for tcp.
	Matcher string `json:"matcher,omitempty"`

	// Path HTTP path probed; ignored for tcp.
	Path string `json:"path,omitempty"`

	// Port probed on the member.
	Port int `json:"port"`

	// Protocol `tcp` opens a connection; `http`/`https` issue a GET and match the
	// status against `matcher`. There is no `udp`: a readiness probe needs
	// an answer — check a udp service on a tcp health port instead.
	//
	// One of: "tcp", "http", "https".
	Protocol string `json:"protocol"`

	// TimeoutSec per-probe timeout; must be less than interval_sec.
	TimeoutSec int `json:"timeout_sec"`

	// UnhealthyThreshold consecutive failures before a member flips unhealthy.
	UnhealthyThreshold int `json:"unhealthy_threshold"`
}

// FloatingIPMember one binding of a floating IP.
type FloatingIPMember struct {
	CreatedAt time.Time `json:"created_at"`

	// Health what the platform knows about this member.
	//
	// `unknown` — nobody is checking. A member you attached yourself
	// with no health check on the address reads this: you chose the moment
	// of attach, and the platform has no signal about what runs inside the
	// instance. It is advertised.
	//
	// `healthy` — the platform has evidence this member is up (and, if a
	// health check is configured on the address, that the check is
	// passing).
	//
	// `unhealthy` — the platform is waiting for that evidence and has
	// not seen it, or a configured check is failing. The member keeps its
	// place on the address and receives no traffic until it recovers.
	//
	// Without a health check this is liveness only — `healthy` means the
	// guest came up, not that your service is listening. Configure
	// `health_check` on the floating IP to add readiness on top of that.
	//
	// One of: "unknown", "healthy", "unhealthy".
	Health string `json:"health"`

	// Interface Bound NIC summary; null for a load balancer binding named by
	// attached_to.
	Interface *FloatingIPMemberInterface `json:"interface"`

	// Reason why the member reads the `health` it does — so you can tell "your
	// service is not answering" from "the guest has not booted yet".
	//
	// `unprobed` — nobody is checking (no health check, hand-attached).
	// `booting` — the platform has not yet seen the guest come up.
	// `probe_failed` — the configured health check is failing. `passing`
	// — the guest is up and, if a check is configured, it passes.
	//
	// One of: "unprobed", "booting", "probe_failed", "passing".
	Reason string `json:"reason"`
}

// FloatingIPMemberInterface Bound NIC summary; null for a load balancer binding named by
// attached_to.
type FloatingIPMemberInterface struct {
	CRN string `json:"crn"`
	ID  string `json:"id"`

	// Instance owning instance; null when the interface has no owning instance.
	Instance *FloatingIPMemberInterfaceInstance `json:"instance"`
}

// FloatingIPMemberInterfaceInstance owning instance; null when the interface has no owning instance.
type FloatingIPMemberInterfaceInstance struct {
	CRN  string `json:"crn"`
	ID   string `json:"id"`
	Name string `json:"name"`
}

type FloatingIPUpdateRequest struct {
	Description *string `json:"description,omitempty"`

	// HealthCheck set (an object) or clear (null) the address's readiness check. Omit
	// the field to leave it unchanged.
	HealthCheck *FloatingIPHealthCheck `json:"health_check,omitempty"`
	Tags        map[string]string      `json:"tags,omitempty"`
}

// IPFamily the address family of a public address. A floating IP is the same
// resource in either family — allocated, attached to one or more
// members, advertised from their chassis — and the family is a
// property of the address rather than a different product. What changes
// with it:
//
//   - **Pool.** A `ipv4` address comes from the region's tenant IPv4 block,
//     a `ipv6` one from its tenant IPv6 block.
//   - **Attach.** A `ipv6` address can only be attached to an interface
//     that has an IPv6 address — one on a dual-stack subnet — and the
//     subnet needs a `::/0` route to an internet gateway, the way a
//     `ipv4` one needs `0.0.0.0/0`. An interface holds at most one
//     floating IP of each family; a v4 and a v6 on the same interface is
//     fine.
//   - **Identity.** While a `ipv6` floating IP is attached, it is the
//     interface's public IPv6 identity: the interface's own address stops
//     being reachable from the internet and comes back when the floating
//     IP is detached. That is the same rule a `ipv4` floating IP has
//     always had, applied to a family whose addresses are public to begin
//     with.
type IPFamily string

// Values IPFamily accepts.
const (
	IPFamilyIPv4 IPFamily = "ipv4"
	IPFamilyIPv6 IPFamily = "ipv6"
)

type Interface struct {
	// AttachedTo UUID of the instance holding this interface, including stopped
	// instances. Null when no instance NIC binding exists. Deletion is
	// refused while bound; floating IP attachment is tracked separately.
	AttachedTo  string    `json:"attached_to,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	CRN         string    `json:"crn"`
	Description string    `json:"description,omitempty"`
	ID          string    `json:"id"`
	IPAddress   string    `json:"ip_address"`

	// IPv6Address the interface's /128, auto-assigned when its subnet is dual-stack.
	IPv6Address string `json:"ipv6_address,omitempty"`
	MAC         string `json:"mac"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	Name      string            `json:"name"`
	Subnet    *Subnet           `json:"subnet"`
	Tags      map[string]string `json:"tags"`
	UpdatedAt time.Time         `json:"updated_at"`
}

type InterfaceCreateRequest struct {
	Description *string `json:"description,omitempty"`

	// IPAddress defaults to the next free address in the subnet
	IPAddress *string `json:"ip_address,omitempty"`

	// MAC defaults to a fresh locally-administered EUI-48
	MAC *string `json:"mac,omitempty"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	//
	// Required.
	Name string `json:"name"`

	// Subnet UUID or nested CRN (vpc/<vpc>/subnet/<subnet>). A bare name
	// requires an explicit VPC filter; create requests without a VPC do
	// not accept bare names.
	//
	// Required.
	Subnet string            `json:"subnet"`
	Tags   map[string]string `json:"tags,omitempty"`
}

type InterfaceSecurityGroupsRequest struct {
	// SecurityGroups Security-group UUIDs, CRNs or account-scoped names. All entries
	// resolve before replacement; duplicate canonical IDs collapse to one
	// membership. An empty array removes all groups.
	//
	// Required.
	SecurityGroups []string `json:"security_groups"`
}

type InterfaceUpdateRequest struct {
	Description *string           `json:"description,omitempty"`
	Tags        map[string]string `json:"tags,omitempty"`
}

type InternetGateway struct {
	// AttachedVPCID VPC the IGW is currently attached to (null when detached).
	AttachedVPCID string    `json:"attached_vpc_id,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
	CRN           string    `json:"crn"`
	Description   string    `json:"description,omitempty"`
	ID            string    `json:"id"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	Name      string            `json:"name"`
	Tags      map[string]string `json:"tags"`
	UpdatedAt time.Time         `json:"updated_at"`
}

type InternetGatewayAttachRequest struct {
	// VPC UUID, CRN or exact name in the caller account.
	//
	// Required.
	VPC string `json:"vpc"`
}

type InternetGatewayCreateRequest struct {
	Description *string `json:"description,omitempty"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	//
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

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	Name      string            `json:"name"`
	Subnet    *Subnet           `json:"subnet"`
	Tags      map[string]string `json:"tags"`
	UpdatedAt time.Time         `json:"updated_at"`
}

type NATGatewayCreateRequest struct {
	Description *string `json:"description,omitempty"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	//
	// Required.
	Name string `json:"name"`

	// Subnet UUID or nested CRN (vpc/<vpc>/subnet/<subnet>). A bare name
	// requires an explicit VPC filter; create requests without a VPC do
	// not accept bare names.
	//
	// Required.
	Subnet string            `json:"subnet"`
	Tags   map[string]string `json:"tags,omitempty"`
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

// RouteCreateRequest exactly one of target_ip / target_internet_gateway /
// target_nat_gateway / target_egress_only_gateway (and future
// target_*_id fields) must be set.
type RouteCreateRequest struct {
	Description *string `json:"description,omitempty"`

	// Required.
	Destination string            `json:"destination"`
	Tags        map[string]string `json:"tags,omitempty"`

	// TargetEgressOnlyGateway Gateway UUID, CRN or exact account-scoped name. Must belong to the
	// route table VPC and match the destination address family. Exactly
	// one route target is required.
	TargetEgressOnlyGateway *string `json:"target_egress_only_gateway,omitempty"`

	// TargetInternetGateway Gateway UUID, CRN or exact account-scoped name. Must belong to the
	// route table VPC and match the destination address family. Exactly
	// one route target is required.
	TargetInternetGateway *string `json:"target_internet_gateway,omitempty"`

	// TargetIP unicast next hop inside this VPC's CIDR (same IP family as
	// destination). Not for internet egress — use a gateway target id
	// instead.
	TargetIP *string `json:"target_ip,omitempty"`

	// TargetNATGateway Gateway UUID, CRN or exact account-scoped name. Must belong to the
	// route table VPC and match the destination address family. Exactly
	// one route target is required.
	TargetNATGateway *string `json:"target_nat_gateway,omitempty"`
}

type RouteTable struct {
	CreatedAt   time.Time `json:"created_at"`
	CRN         string    `json:"crn"`
	Description string    `json:"description,omitempty"`
	ID          string    `json:"id"`

	// IsMain true for the per-VPC default table. The main table is created
	// automatically and can't be deleted. Subnets that don't specify a
	// route_table_id at create time land here.
	IsMain bool `json:"is_main"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	Name      string            `json:"name"`
	Tags      map[string]string `json:"tags"`
	UpdatedAt time.Time         `json:"updated_at"`
	VPC       *VPC              `json:"vpc"`
}

type RouteTableCreateRequest struct {
	Description *string `json:"description,omitempty"`

	// Name 1-63 chars, lowercase alphanumeric + hyphen. `main` is reserved.
	// Resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	//
	// Required.
	Name string            `json:"name"`
	Tags map[string]string `json:"tags,omitempty"`

	// VPC UUID, CRN or exact name in the caller account.
	//
	// Required.
	VPC string `json:"vpc"`
}

// RouteTableSummary route table used by a subnet, without repeating its VPC. Null when the
// non-owning lookup no longer resolves, for example during concurrent
// reassociation and deletion of the former table. Deleting a table still
// associated with subnets is refused.
type RouteTableSummary struct {
	CRN  string `json:"crn"`
	ID   string `json:"id"`
	Name string `json:"name"`
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
	CreatedAt   time.Time `json:"created_at"`
	CRN         string    `json:"crn"`
	Description string    `json:"description,omitempty"`
	ID          string    `json:"id"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	Name      string            `json:"name"`
	Tags      map[string]string `json:"tags"`
	UpdatedAt time.Time         `json:"updated_at"`
}

type SecurityGroupCreateRequest struct {
	Description *string `json:"description,omitempty"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	//
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
	Protocol   SecurityGroupRuleProtocol `json:"protocol"`
	SourceCIDR *string                   `json:"source_cidr,omitempty"`

	// SourceSecurityGroup Security-group UUID, CRN or exact account-scoped name. Mutually
	// exclusive with source_cidr.
	SourceSecurityGroup *string `json:"source_security_group,omitempty"`
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

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	Name       string             `json:"name"`
	RouteTable *RouteTableSummary `json:"route_table"`
	Tags       map[string]string  `json:"tags"`
	UpdatedAt  time.Time          `json:"updated_at"`
	VPC        *VPC               `json:"vpc"`
}

type SubnetCreateRequest struct {
	// AssignIPv6CIDR allocate the lowest free IPv6 /64 inside the VPC's IPv6 CIDR.
	// Requires a VPC created with IPv6. Mutually exclusive with a nonempty
	// cidr_v6. Returns 409 when the VPC has no free IPv6 /64s.
	AssignIPv6CIDR *bool `json:"assign_ipv6_cidr,omitempty"`

	// Required.
	CIDR string `json:"cidr"`

	// CIDRV6 makes the subnet dual-stack. A /64 inside the VPC's IPv6 CIDR (the
	// VPC must have been created with assign_ipv6_cidr). Mutually
	// exclusive with assign_ipv6_cidr=true. Omit both fields for a v4-only
	// subnet.
	CIDRV6      *string `json:"cidr_v6,omitempty"`
	Description *string `json:"description,omitempty"`

	// GatewayIP defaults to the first usable host in the CIDR
	GatewayIP *string `json:"gateway_ip,omitempty"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	//
	// Required.
	Name string `json:"name"`

	// RouteTable Route-table UUID, nested CRN or exact name within the subnet VPC. On
	// PATCH the owned path subnet supplies the VPC. Omission on create
	// selects main; an empty reference is invalid.
	RouteTable *string           `json:"route_table,omitempty"`
	Tags       map[string]string `json:"tags,omitempty"`

	// VPC UUID, CRN or exact name in the caller account.
	//
	// Required.
	VPC string `json:"vpc"`
}

type SubnetUpdateRequest struct {
	Description *string `json:"description,omitempty"`

	// RouteTable Route-table UUID, nested CRN or exact name within the subnet VPC. On
	// PATCH the owned path subnet supplies the VPC. Omission on create
	// selects main; an empty reference is invalid.
	RouteTable *string           `json:"route_table,omitempty"`
	Tags       map[string]string `json:"tags,omitempty"`
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

	// Name 1-63 chars, lowercase alphanumeric + hyphen Resource names must not
	// start with the literal crn: prefix or be UUIDs (canonical, compact,
	// braced, or urn:uuid: forms, in either case).
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

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	//
	// Required.
	Name string            `json:"name"`
	Tags map[string]string `json:"tags,omitempty"`
}

type VPCUpdateRequest struct {
	Description *string           `json:"description,omitempty"`
	Tags        map[string]string `json:"tags,omitempty"`
}
