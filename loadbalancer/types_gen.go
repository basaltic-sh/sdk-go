// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package loadbalancer

import (
	"time"
)

type AttachListenerCertificateRequest struct {
	// Certificate the certificate to serve, by CRN, UUID or exact account-scoped name.
	// Certificate CRNs require an empty region. The listener stores a
	// reference — no key material is sent here, and the replicas fetch
	// it from the certificate service under their own identity.
	//
	// Required.
	Certificate string `json:"certificate"`

	// IsDefault when true, demote whatever's currently default and promote this cert
	// in the same transaction.
	IsDefault *bool `json:"is_default,omitempty"`
}

// AttachTargetRequest target is a literal IP for IP groups or an instance UUID, CRN or exact
// account-scoped name for instance groups.
type AttachTargetRequest struct {
	Port *int `json:"port,omitempty"`

	// Target must match the group's target_type: an IP address for `ip`, a
	// compute instance UUID, CRN or exact account-scoped name for
	// `instance`. An `ip` ref has to be a routable unicast address —
	// loopback, link-local (including the 169.254.169.254 metadata
	// endpoint), multicast, and unspecified addresses are rejected.
	//
	// Required.
	Target string `json:"target"`
}

// AutoscalingPolicy target tracking shared by instance pools and load balancers. Updates
// replace the policy. Set enabled=false to retain settings and use
// manual sizing. Each metric recommends a desired count; the largest
// recommendation wins. Missing, stale or incomplete observations prevent
// scale-in but do not block scale-out recommended by another valid
// metric. Decisions obey the resource's min_count/max_count, warmup,
// cooldown, stabilization, step limits and quotas. State survives
// controller restarts. Active policies own desired_count; manual changes
// are accepted and automatic evaluation resumes after cooldown. Custom
// telemetry requires telemetry:ReadMetrics in the same account.
type AutoscalingPolicy struct {
	CooldownSeconds *int `json:"cooldown_seconds,omitempty"`

	// DrainSeconds grace period after route withdrawal and proxy acknowledgements,
	// before deleting a retiring member. Long-lived TCP/UDP sessions may
	// end at the deadline; arbitrary application shutdown hooks are not
	// supported.
	DrainSeconds                  *int             `json:"drain_seconds,omitempty"`
	Enabled                       bool             `json:"enabled"`
	MaxScaleInStep                *int             `json:"max_scale_in_step,omitempty"`
	MaxScaleOutStep               *int             `json:"max_scale_out_step,omitempty"`
	Metrics                       []*ScalingMetric `json:"metrics"`
	ScaleDownStabilizationSeconds *int             `json:"scale_down_stabilization_seconds,omitempty"`
	WarmupSeconds                 *int             `json:"warmup_seconds,omitempty"`
}

type AutoscalingStatus struct {
	EvaluatedAt  time.Time                       `json:"evaluated_at,omitempty"`
	History      []*AutoscalingStatusHistoryItem `json:"history"`
	LastScaledAt time.Time                       `json:"last_scaled_at,omitempty"`
	Reason       string                          `json:"reason"`

	// One of: "pending", "disabled", "stable", "scaling", "waiting", "warming_up", "metrics_unavailable", "stabilizing", "cooldown", "draining".
	Status string `json:"status"`
}

type AutoscalingStatusHistoryItem struct {
	At     time.Time `json:"at"`
	From   int       `json:"from"`
	Reason string    `json:"reason"`
	To     int       `json:"to"`
}

type CreateListenerCertificate struct {
	// Certificate CRN, UUID or exact name in the caller's account.
	// Certificate CRNs require an empty region. No key material is
	// accepted.
	//
	// Required.
	Certificate string `json:"certificate"`
}

// CreateListenerRequest target group relationships accept UUID, CRN or exact account-scoped
// names in this region. HTTPS listeners require >=1 certificate, named
// by CRN, UUID or exact account-scoped name in `certificates`; the first
// entry becomes the default (the fallback when SNI doesn't match). No
// key material is accepted — the agent fetches it from the certificate
// service against the CRN.
type CreateListenerRequest struct {
	Certificates       []*CreateListenerCertificate `json:"certificates,omitempty"`
	DefaultTargetGroup *string                      `json:"default_target_group,omitempty"`

	// Exposure Which LB addresses are bound. Defaults to 'both'; pick private_only
	// when the LB has no FIP yet.
	//
	// One of: "public_only", "private_only", "both".
	Exposure *string `json:"exposure,omitempty"`

	// Required.
	Port int `json:"port"`

	// One of: "http", "https", "tcp", "udp".
	//
	// Required.
	Protocol string `json:"protocol"`
	Tags     Tags   `json:"tags,omitempty"`
}

// CreateLoadBalancerRequest relationships accept a UUID, CRN or exact immutable name, classified
// by syntax. VPC, subnet and security groups must belong to the caller
// account in this region. Subnet names are scoped by vpc. Flavor is a
// regional catalog reference. Floating IP references accept UUID or CRN
// only. All references resolve before writes. Addresses are fixed at
// creation; updates cannot replace them.
type CreateLoadBalancerRequest struct {
	Autoscaling *AutoscalingPolicy `json:"autoscaling,omitempty"`

	// DesiredCount steady target within min_count and max_count.
	DesiredCount *int `json:"desired_count,omitempty"`

	// Flavor compute flavor for each LB instance.
	//
	// Required.
	Flavor string `json:"flavor"`

	// FloatingIP Public IPv4 shorthand. Cannot be combined with floating_ips. Does
	// not allocate public IPv6.
	FloatingIP *string `json:"floating_ip,omitempty"`

	// FloatingIPs existing free floating IPs from this account and region, at most one
	// per family and visibility (private/public, IPv4/IPv6). Private
	// addresses must belong to the selected subnet. Missing private
	// families are allocated automatically for each family enabled on that
	// subnet. Public addresses are optional and require a matching-family
	// default route to an internet gateway; NAT and egress-only gateways
	// do not qualify. IPv6 requires an IPv6-enabled subnet. Pool-owned or
	// attached addresses are unavailable. On deletion, supplied addresses
	// are detached and retained; automatic private allocations are
	// released. Cannot be combined with floating_ip.
	FloatingIPs []string `json:"floating_ips,omitempty"`

	// MaxCount upper capacity bound including rollout surge.
	MaxCount *int `json:"max_count,omitempty"`

	// MinCount lower capacity bound.
	MinCount *int `json:"min_count,omitempty"`

	// Name 1..127 chars of [A-Za-z0-9._-] Resource names must not start with
	// the literal crn: prefix or be UUIDs (canonical, compact, braced, or
	// urn:uuid: forms, in either case).
	//
	// Required.
	Name string `json:"name"`

	// ReplicaCount deprecated input alias of desired_count; send only one. Desired
	// defaults to 1. Omitted bounds default to desired.
	ReplicaCount *int `json:"replica_count,omitempty"`

	// SecurityGroups security groups attached to every replica NIC (AWS ALB shape). A VPC
	// NIC with no security group denies all data traffic, so the listener
	// port(s) must be opened by a security group listed here. Re-applied
	// to replacement replicas. The LB's own control-plane path (agent
	// config + heartbeat via the metadata endpoint) is always-allowed and
	// needs none.
	//
	// Required.
	SecurityGroups []string `json:"security_groups"`

	// Subnet the LB instances attach to. The virtual IP is allocated from
	// this subnet.
	//
	// Required.
	Subnet string `json:"subnet"`
	Tags   Tags   `json:"tags,omitempty"`

	// One of: "application", "network".
	//
	// Required.
	Type string `json:"type"`

	// VPC the LB will live in. Must match subnet's VPC.
	//
	// Required.
	VPC string `json:"vpc"`
}

// CreateRuleRequest target group relationships accept UUID, CRN or exact account-scoped
// names in this region.
type CreateRuleRequest struct {
	// Required.
	Conditions []*RuleCondition `json:"conditions"`

	// Required.
	Priority int `json:"priority"`

	// Required.
	TargetGroup string `json:"target_group"`
}

// CreateTargetGroupRequest instance_pool accepts a UUID, CRN or exact account-scoped name in this
// region and requires pool target mode.
type CreateTargetGroupRequest struct {
	HealthCheck *HealthCheck `json:"health_check,omitempty"`

	// InstancePool compute instance pool to draw backends from. Required when
	// target_mode=pool and must belong to the calling account; ignored
	// otherwise.
	InstancePool *string `json:"instance_pool,omitempty"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	//
	// Required.
	Name string `json:"name"`

	// Required.
	Port int `json:"port"`

	// One of: "http", "https", "tcp", "udp".
	//
	// Required.
	Protocol        string           `json:"protocol"`
	ProxyProtocol   *bool            `json:"proxy_protocol,omitempty"`
	SessionAffinity *SessionAffinity `json:"session_affinity,omitempty"`
	Tags            Tags             `json:"tags,omitempty"`

	// TargetMode `static` (the default) takes the backends you attach as targets.
	// `pool` takes them from a compute instance pool and requires
	// instance_pool; the group is forced to target_type=instance, and
	// attaching targets to it is rejected.
	//
	// One of: "static", "pool".
	TargetMode *string `json:"target_mode,omitempty"`

	// One of: "ip", "instance", "function".
	TargetType *string `json:"target_type,omitempty"`
}

type Fault struct {
	// Code stable machine-readable code owned by the reporting operation.
	Code string `json:"code"`

	// Details structured context; legacy strings are preserved in legacy_text.
	Details map[string]any `json:"details"`

	// FirstAt first observation in this active occurrence series.
	FirstAt time.Time `json:"first_at"`

	// LastAt latest observation in this active occurrence series.
	LastAt      time.Time `json:"last_at"`
	Message     string    `json:"message"`
	Occurrences int       `json:"occurrences"`

	// One of: "error", "warning".
	Severity string `json:"severity"`
}

type FloatingIP struct {
	// Address allocated public or private address.
	Address string `json:"address"`

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
	Members []*FloatingIPMember `json:"members"`

	// SubnetID allocation subnet for private floating IPs.
	SubnetID  string            `json:"subnet_id,omitempty"`
	Tags      map[string]string `json:"tags"`
	UpdatedAt time.Time         `json:"updated_at"`

	// One of: "public", "private".
	Visibility string `json:"visibility"`

	// VPCID Allocation VPC for private floating IPs.
	VPCID string `json:"vpc_id,omitempty"`
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
	// AddressID target child address on the member interface.
	AddressID string    `json:"address_id,omitempty"`
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

// HealthCheck active probes are enabled by default. Set enabled=false to stop probes
// while retaining their configuration. An explicitly enabled HTTP/HTTPS
// check requires a non-empty path starting with /. TCP checks connect to
// the check port without an HTTP path. UDP groups default to TCP connect
// probes.
type HealthCheck struct {
	Enabled *bool `json:"enabled,omitempty"`

	// HealthyThreshold zero uses the default.
	HealthyThreshold *int `json:"healthy_threshold,omitempty"`

	// IntervalSec zero uses the default.
	IntervalSec *int `json:"interval_sec,omitempty"`

	// Matcher HTTP status codes from 100 to 599; comma-separated codes or
	// inclusive ranges. Empty uses 200.
	Matcher *string `json:"matcher,omitempty"`
	Path    *string `json:"path,omitempty"`
	Port    *int    `json:"port,omitempty"`

	// Protocol omitted or empty uses the target group protocol. UDP uses a TCP
	// connect probe. HTTPS probes use TLS.
	//
	// One of: "http", "https", "tcp", "udp", "".
	Protocol *string `json:"protocol,omitempty"`

	// TimeoutSec zero uses the default.
	TimeoutSec *int `json:"timeout_sec,omitempty"`

	// UnhealthyThreshold zero uses the default.
	UnhealthyThreshold *int `json:"unhealthy_threshold,omitempty"`
}

// HealthCheckPatch merge supplied fields into the existing check. Omitted fields are
// preserved; null or an empty object leaves the check unchanged. Set
// enabled=false to disable probes without clearing settings, and
// enabled=true to enable the saved configuration. Set port=0 to use each
// target's traffic port; zero timing/thresholds and an empty matcher
// reset their defaults. An empty protocol inherits the target group
// protocol. Clearing the path of an explicitly enabled HTTP/HTTPS check
// is invalid.
type HealthCheckPatch struct {
	Enabled          *bool             `json:"enabled,omitempty"`
	HealthyThreshold *HealthyThreshold `json:"healthy_threshold,omitempty"`
	IntervalSec      *IntervalSec      `json:"interval_sec,omitempty"`
	Matcher          *Matcher          `json:"matcher,omitempty"`
	Path             *Path             `json:"path,omitempty"`

	// Port zero removes the probe port override.
	Port               *int                `json:"port,omitempty"`
	Protocol           *Protocol           `json:"protocol,omitempty"`
	TimeoutSec         *TimeoutSec         `json:"timeout_sec,omitempty"`
	UnhealthyThreshold *UnhealthyThreshold `json:"unhealthy_threshold,omitempty"`
}

// HealthyThreshold zero uses the default.
type HealthyThreshold = int

// IPFamily The IP address family. Floating IPs support either family and may have
// public or private visibility. Public floating IPs allocate from the
// region's public address pool; private floating IPs allocate from their
// selected subnet's range for that family.
//
// Attaching a floating IP to an interface requires an address of the
// same family on that interface. IPv6 does not require IPv4 to be
// enabled on the subnet. Internet reachability also depends on routes
// and security rules.
//
// Attaching an IPv6 floating IP does not disable the interface's native
// globally routable IPv6 address. Both addresses remain reachable when
// routing and security rules permit, and replies to incoming connections
// retain the address that received the connection. A private IPv6
// address does not become directly internet-routable by attaching a
// floating IP.
type IPFamily string

// Values IPFamily accepts.
const (
	IPFamilyIPv4 IPFamily = "ipv4"
	IPFamilyIPv6 IPFamily = "ipv6"
)

// IntervalSec zero uses the default.
type IntervalSec = int

type Listener struct {
	// Certificates HTTPS listeners only. One entry per attached certificate; the right
	// cert is picked per-connection by matching the client's SNI against
	// each cert's SAN. The entry flagged is_default serves traffic that
	// doesn't match any other SNI (or clients that omit SNI). PEM material
	// is NOT echoed — the listener stores its own copy fetched at attach
	// time.
	Certificates []*ListenerCertificate `json:"certificates,omitempty"`
	CreatedAt    time.Time              `json:"created_at"`

	// CRN Parent-scoped CRN with immutable load balancer name and child UUID
	// components.
	CRN                  string `json:"crn"`
	DefaultTargetGroupID string `json:"default_target_group_id,omitempty"`

	// Exposure Which LB addresses this listener binds. 'public_only' and 'both'
	// require the LB to carry a floating IP; if the FIP is detached later,
	// the listener is dropped until a FIP is re-attached or exposure is
	// flipped to private_only.
	//
	// One of: "public_only", "private_only", "both".
	Exposure       string `json:"exposure"`
	ID             string `json:"id"`
	LoadBalancerID string `json:"load_balancer_id"`
	Port           int    `json:"port"`

	// One of: "http", "https", "tcp", "udp".
	Protocol  string    `json:"protocol"`
	Tags      Tags      `json:"tags"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ListenerCertificate struct {
	CertificateCRN string    `json:"certificate_crn"`
	CreatedAt      time.Time `json:"created_at"`
	ID             string    `json:"id"`
	IsDefault      bool      `json:"is_default"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type LoadBalancer struct {
	AccountID         string             `json:"account_id"`
	Autoscaling       *AutoscalingPolicy `json:"autoscaling,omitempty"`
	AutoscalingStatus *AutoscalingStatus `json:"autoscaling_status,omitempty"`
	CreatedAt         time.Time          `json:"created_at"`

	// CRN IAM resource CRN
	CRN string `json:"crn"`

	// DesiredCount steady target within min_count and max_count.
	DesiredCount int `json:"desired_count"`

	// DNSName convenience hostname auto-published for the load balancer,
	// `{name}.{account-handle}.lb.{region}.{base-domain}`. Resolves to the
	// floating IP on an internet-facing LB and to the private VIP
	// otherwise. Omitted in regions where auto-DNS is not configured —
	// the VIP and FIP stay authoritative either way.
	DNSName string `json:"dns_name,omitempty"`

	// Faults active faults; status is error exactly when an active error fault
	// remains.
	Faults []*Fault `json:"faults"`

	// FlavorID compute flavor each LB instance runs on. Must be a
	// loadbalancer-family flavor.
	FlavorID string `json:"flavor_id"`

	// FloatingIPID optional public IPv4 floating IP. Public IPv6 is independent.
	FloatingIPID string `json:"floating_ip_id,omitempty"`

	// FloatingIPs all attached public and private floating IPs, including automatic
	// private allocations.
	FloatingIPs []*FloatingIP `json:"floating_ips"`
	ID          string        `json:"id"`

	// InternalIPv4 Virtual IP for the load balancer; traffic is distributed to backends
	// per connection.
	InternalIPv4 string `json:"internal_ipv4,omitempty"`

	// InternalIPv6 Internal IPv6 VIP (set when the subnet is dual-stack).
	InternalIPv6 string `json:"internal_ipv6,omitempty"`

	// MaxCount upper capacity bound including rollout surge.
	MaxCount int `json:"max_count"`

	// MinCount lower capacity bound.
	MinCount int `json:"min_count"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	Name string `json:"name"`

	// PublicIPv6 explicitly selected public IPv6 floating IP, translated to replica
	// IPv6 addresses in a GUA or ULA subnet.
	PublicIPv6 string `json:"public_ipv6,omitempty"`

	// ReplicaCount deprecated alias of desired_count.
	ReplicaCount int `json:"replica_count"`

	// RolloutSurge temporary extra capacity within max_count; does not change
	// desired_count.
	RolloutSurge bool `json:"rollout_surge,omitempty"`

	// One of: "provisioning", "active", "error", "deleting".
	Status string `json:"status"`

	// Subnet placement; null when the referenced subnet no longer exists.
	Subnet *Subnet `json:"subnet"`
	Tags   Tags    `json:"tags"`

	// Type ALB-shape (L7) vs NLB-shape (L4)
	//
	// One of: "application", "network".
	Type      string    `json:"type"`
	UpdatedAt time.Time `json:"updated_at"`
}

// LoadBalancerReplica one compute instance backing the load balancer. The bookkeeping fields
// (instance_id, replica_index, created_at, flavor_id) are persisted; the
// liveness overlay (last_seen, proxy_ok, agent_version) refreshes on
// each replica health report.
//
// Replicas are managed instances, so the compute instance endpoints do
// not return them — this is where you watch them. During a resize the
// replicas turn over one at a time: a replica has been replaced when its
// instance_id changes, and the resize is complete when every flavor_id
// here matches the load balancer's. Expect one more replica than
// replica_count to be listed part-way through, which is the resize
// keeping capacity up rather than a replica leaking; it goes away when
// the last old one does.
type LoadBalancerReplica struct {
	AgentVersion string    `json:"agent_version,omitempty"`
	CreatedAt    time.Time `json:"created_at"`

	// FlavorID the size this replica actually booted on. Matches the load
	// balancer's flavor_id except mid-resize, when the replicas not yet
	// replaced still report the old one.
	FlavorID   string `json:"flavor_id"`
	InstanceID string `json:"instance_id"`

	// LastSeen omitted when this replica hasn't reported yet
	LastSeen time.Time `json:"last_seen,omitempty"`

	// ProxyOk whether the proxy reported ready at the last health report
	ProxyOk      bool        `json:"proxy_ok"`
	ReplicaIndex int         `json:"replica_index"`
	Retirement   *Retirement `json:"retirement,omitempty"`

	// Status the liveness view folded into one word: 'initializing' (the agent
	// has never reported — boot still in flight), 'healthy'
	// (heartbeating and the proxy is serving), 'unhealthy' (heartbeating
	// but the proxy is down).
	//
	// One of: "initializing", "healthy", "unhealthy", "draining".
	Status string `json:"status"`
}

// Matcher HTTP status codes from 100 to 599; comma-separated codes or inclusive
// ranges. Empty uses 200.
type Matcher = string

type Path = string

// Protocol omitted or empty uses the target group protocol. UDP uses a TCP
// connect probe. HTTPS probes use TLS.
type Protocol string

// Values Protocol accepts.
const (
	ProtocolHTTP  Protocol = "http"
	ProtocolHTTPS Protocol = "https"
	ProtocolTcp   Protocol = "tcp"
	ProtocolUdp   Protocol = "udp"
	ProtocolField Protocol = ""
)

type Retirement struct {
	AgentAcknowledgedAt time.Time `json:"agent_acknowledged_at,omitempty"`
	DrainSeconds        int       `json:"drain_seconds"`

	// DrainUntil earliest deletion time; absent while withdrawal is pending.
	DrainUntil  time.Time `json:"drain_until,omitempty"`
	RequestedAt time.Time `json:"requested_at"`
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

type Rule struct {
	Conditions []*RuleCondition `json:"conditions"`
	CreatedAt  time.Time        `json:"created_at"`

	// CRN Parent-scoped CRN with immutable load balancer name and child UUID
	// components.
	CRN           string    `json:"crn"`
	ID            string    `json:"id"`
	ListenerID    string    `json:"listener_id"`
	Priority      int       `json:"priority"`
	TargetGroupID string    `json:"target_group_id"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type RuleCondition struct {
	// One of: "host", "path", "header", "query", "method".
	Field string `json:"field"`

	// Name header or query key name
	Name *string `json:"name,omitempty"`

	// One of: "exact", "prefix", "glob", "regex".
	Op     string   `json:"op"`
	Values []string `json:"values"`
}

// ScalingMetric CPU uses source=cpu, target_type=utilization and a percentage target
// <=100. Utilization is CPU seconds per second divided by allocated
// vCPUs across all ready members. Enabled CPU scaling requires
// min_count>=1. Other selector and aggregation fields are not allowed
// for CPU.
//
// Custom demand uses source=telemetry and target_type=average_value.
// Metric name and labels select series in the resource's account,
// organization and region. Temporal aggregation is applied within each
// series before combining series; repeated gauge samples are never
// summed as extra demand. The desired count is ceil(combined value /
// target_value): 750 pending jobs at a target of 100 per instance
// recommends 8. Custom demand can scale a customer pool from 0.
// Producers must publish fresh zeroes for idle queues; absent data is
// not zero.
type ScalingMetric struct {
	// ExpectedSeries exact expected cardinality; incomplete or ambiguous selectors are
	// unavailable.
	ExpectedSeries *int `json:"expected_series,omitempty"`

	// Labels exact-match labels; tenancy labels and __name__ cannot be supplied.
	Labels map[string]string `json:"labels,omitempty"`

	// MaxAgeSeconds actual newest observation age per series; must not exceed
	// window_seconds. Defaults to the smaller of 90 and the window.
	MaxAgeSeconds *int    `json:"max_age_seconds,omitempty"`
	Name          *string `json:"name,omitempty"`

	// SampleAggregation use last for queue gauges; rate for monotonically increasing
	// counters, with reset handling.
	//
	// One of: "last", "avg", "max", "rate".
	SampleAggregation *string `json:"sample_aggregation,omitempty"`

	// One of: "sum", "avg", "max".
	SeriesAggregation *string `json:"series_aggregation,omitempty"`

	// One of: "cpu", "telemetry".
	Source string `json:"source"`

	// One of: "utilization", "average_value".
	TargetType    string  `json:"target_type"`
	TargetValue   float64 `json:"target_value"`
	WindowSeconds *int    `json:"window_seconds,omitempty"`
}

// SessionAffinity sticky sessions: pin a client to one backend in the group instead of
// balancing each request independently. Off unless you ask for it.
//
// The data plane hashes the key consistently, so adding or losing a
// backend only moves the clients that backend was serving.
type SessionAffinity struct {
	// CookieName cookie the load balancer sets and hashes. `type=cookie` only; the
	// field is dropped for the other types.
	CookieName *string `json:"cookie_name,omitempty"`

	// DurationSec how long that cookie lives, up to 7 days. `type=cookie` only.
	DurationSec *int `json:"duration_sec,omitempty"`

	// Type `none` balances every request. `cookie` sets an opaque cookie on the
	// first response and routes every later request carrying it to the
	// same backend — http and https groups only. `source_ip` hashes the
	// client address, works on every protocol and is the only option for
	// tcp and udp, but a NAT gateway makes every client behind it a single
	// key.
	//
	// One of: "none", "cookie", "source_ip".
	Type string `json:"type"`
}

type Subnet struct {
	CIDRIPv4 string `json:"cidr_ipv4"`

	// CIDRIPv6 the dual-stack IPv6 /64, if the subnet is v6-enabled. Its presence
	// (vs the v4 cidr_ipv4) is how a client tells the subnet's families
	// apart.
	CIDRIPv6    string    `json:"cidr_ipv6,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	CRN         string    `json:"crn"`
	Description string    `json:"description,omitempty"`
	GatewayIPv4 string    `json:"gateway_ipv4"`
	GatewayIPv6 string    `json:"gateway_ipv6,omitempty"`
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

type Tags = map[string]string

type Target struct {
	CreatedAt time.Time `json:"created_at"`

	// One of: "initial", "healthy", "unhealthy", "draining".
	Health        string    `json:"health"`
	ID            string    `json:"id"`
	LastSeenAt    time.Time `json:"last_seen_at,omitempty"`
	Port          int       `json:"port,omitempty"`
	TargetGroupID string    `json:"target_group_id"`

	// TargetRef IP address (target_type=ip) or compute instance id
	// (target_type=instance). Stored in canonical form, so the spelling
	// here may differ from the one you sent.
	TargetRef string    `json:"target_ref"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TargetGroup struct {
	AccountID   string       `json:"account_id"`
	CreatedAt   time.Time    `json:"created_at"`
	CRN         string       `json:"crn"`
	HealthCheck *HealthCheck `json:"health_check"`
	ID          string       `json:"id"`

	// InstancePoolID compute instance pool backing the group. Set iff target_mode=pool.
	InstancePoolID string `json:"instance_pool_id,omitempty"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	Name string `json:"name"`
	Port int    `json:"port"`

	// One of: "http", "https", "tcp", "udp".
	Protocol string `json:"protocol"`

	// ProxyProtocol when true, upstream connections are wrapped in the PROXY v2 header
	// so backends see the original client IP + port. HTTP backends already
	// get X-Forwarded-For; PROXY is the right pick for TCP/UDP target
	// groups or HTTP backends that prefer the framed envelope.
	ProxyProtocol   bool             `json:"proxy_protocol"`
	SessionAffinity *SessionAffinity `json:"session_affinity"`
	Tags            Tags             `json:"tags"`

	// TargetMode where the backend set comes from. `static` routes to the targets
	// attached via POST /v1/target-groups/{id}/targets; `pool` resolves
	// live instance addresses from the compute instance pool named by
	// instance_pool_id, so scaling the pool moves the backends with it.
	//
	// One of: "static", "pool".
	TargetMode string `json:"target_mode"`

	// One of: "ip", "instance", "function".
	TargetType string    `json:"target_type"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// TimeoutSec zero uses the default.
type TimeoutSec = int

// UnhealthyThreshold zero uses the default.
type UnhealthyThreshold = int

// UpdateListenerRequest patch a listener.
//
// `certificate` accepts a CRN, UUID or exact account-scoped name and
// names a certificate ALREADY attached to this listener and re-stamps
// it, which makes the agent re-fetch material — the on-demand rotation
// trigger. It does not attach: an unattached CRN is refused, and the
// attach-certificate endpoint is what adds one. No key material is
// accepted here.
//
// Set clear_default_target_group=true to remove the default; otherwise
// omitting default_target_group leaves it unchanged.
type UpdateListenerRequest struct {
	Certificate             *string `json:"certificate,omitempty"`
	ClearDefaultTargetGroup *bool   `json:"clear_default_target_group,omitempty"`
	DefaultTargetGroup      *string `json:"default_target_group,omitempty"`

	// Exposure mutate which addresses are bound. Omit to leave unchanged.
	//
	// One of: "public_only", "private_only", "both".
	Exposure *string `json:"exposure,omitempty"`
	Tags     Tags    `json:"tags,omitempty"`
}

// UpdateLoadBalancerRequest names are fixed at creation because they form the CRN used by IAM
// policies. Sending name in an update, including an unchanged, empty or
// null value, returns a validation error.
type UpdateLoadBalancerRequest struct {
	Autoscaling *AutoscalingPolicy `json:"autoscaling,omitempty"`

	// DesiredCount steady target within min_count and max_count.
	DesiredCount *int `json:"desired_count,omitempty"`

	// Flavor resize each replica to a different compute flavor. Must be a
	// loadbalancer-family flavor.
	//
	// A running instance cannot change size in place, so the request
	// records the new size and returns; the replicas already up are then
	// replaced one at a time in the background. The load balancer
	// temporarily runs one replica over desired_count, within max_count
	// while it does: the extra replica comes up on the new flavor and
	// starts serving before any replica on the old one is retired, so the
	// number serving never drops below desired_count — a resize does not
	// cost you capacity, at any replica count.
	//
	// Expect it to take several minutes, and poll GET
	// /v1/load-balancers/{id}/replicas to watch: a replica has been
	// replaced when its instance_id changes, and the resize is done when
	// every flavor there matches this one.
	//
	// A resize requires max_count above desired_count for surge headroom.
	// A rollout waits if headroom is removed while it is in progress.
	//
	// Rejected up front if the account does not have the compute quota for
	// the replacement replica, so a resize cannot half-apply and leave the
	// load balancer short.
	Flavor *string `json:"flavor,omitempty"`

	// MaxCount upper capacity bound including rollout surge.
	MaxCount *int `json:"max_count,omitempty"`

	// MinCount lower capacity bound.
	MinCount *int `json:"min_count,omitempty"`

	// ReplicaCount deprecated alias of desired_count; send only one. Bounds are
	// preserved. With desired_count omitted, it is clamped into the
	// resulting bounds. Scale-in withdraws and drains members before
	// deletion.
	ReplicaCount *int `json:"replica_count,omitempty"`
	Tags         Tags `json:"tags,omitempty"`
}

// UpdateRuleRequest target group relationships accept UUID, CRN or exact account-scoped
// names in this region. Full replace of the rule — priority,
// conditions, and target group must all be supplied (same shape as
// create).
type UpdateRuleRequest struct {
	// Required.
	Conditions []*RuleCondition `json:"conditions"`

	// Required.
	Priority int `json:"priority"`

	// Required.
	TargetGroup string `json:"target_group"`
}

// UpdateTargetGroupRequest names are fixed at creation because they form the CRN used by IAM
// policies. Sending name in an update, including an unchanged, empty or
// null value, returns a validation error.
type UpdateTargetGroupRequest struct {
	HealthCheck *HealthCheckPatch `json:"health_check,omitempty"`

	// ProxyProtocol Toggle PROXY v2 framing on upstream connections. Omitting the field
	// leaves the current setting; setting true/false flips it explicitly.
	ProxyProtocol *bool `json:"proxy_protocol,omitempty"`

	// SessionAffinity replaces the stickiness config. Omitting the field leaves it alone;
	// turning it off is an explicit `{"type": "none"}`.
	SessionAffinity *SessionAffinity `json:"session_affinity,omitempty"`
	Tags            Tags             `json:"tags,omitempty"`
}

type VPC struct {
	// CIDRIPv4 IPv4 CIDR block carved up by subnets. Must be private (RFC 1918):
	// within 10.0.0.0/8, 172.16.0.0/12 or 192.168.0.0/16. Immutable after
	// create.
	CIDRIPv4 string `json:"cidr_ipv4"`

	// CIDRIPv6 associated regional GUA or private ULA prefix.
	CIDRIPv6  string    `json:"cidr_ipv6,omitempty"`
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
