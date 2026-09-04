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
	// CertificateCRN the certificate to serve, by CRN. The listener stores a reference
	// — no key material is sent here, and the replicas fetch it from the
	// certificate service under their own identity.
	//
	// Required.
	CertificateCRN string `json:"certificate_crn"`

	// IsDefault when true, demote whatever's currently default and promote this cert
	// in the same transaction.
	IsDefault *bool `json:"is_default,omitempty"`
}

type AttachTargetRequest struct {
	Port *int `json:"port,omitempty"`

	// TargetRef must match the group's target_type: an IP address for `ip`, a
	// compute instance id for `instance`. An `ip` ref has to be a routable
	// unicast address — loopback, link-local (including the
	// 169.254.169.254 metadata endpoint), multicast, and unspecified
	// addresses are rejected.
	//
	// Required.
	TargetRef string `json:"target_ref"`
}

type CreateListenerCertificate struct {
	// CertificateCRN cert reference stored alongside the PEM material so rotation flows
	// can map back to the upstream cert id.
	//
	// Required.
	CertificateCRN string `json:"certificate_crn"`

	// CertificatePEM leaf cert in PEM.
	//
	// Required.
	CertificatePEM string `json:"certificate_pem"`

	// ChainPEM intermediate chain in PEM (concatenated). Optional.
	ChainPEM *string `json:"chain_pem,omitempty"`

	// PrivateKeyPEM private key in PEM. Sensitive — handled the same way as the cert.
	//
	// Required.
	PrivateKeyPEM string `json:"private_key_pem"`
}

// CreateListenerRequest HTTPS listeners require >=1 certificate. Pass them via `certificates`;
// the first entry becomes the default (the fallback when SNI doesn't
// match). The legacy singular fields (certificate_crn + certificate_pem
// + chain_pem + private_key_pem) are still accepted for back-compat and
// fold into a one-element certificates list.
type CreateListenerRequest struct {
	// CertificateCRN deprecated — use certificates[]
	CertificateCRN *string `json:"certificate_crn,omitempty"`

	// CertificatePEM deprecated — use certificates[]
	CertificatePEM *string                      `json:"certificate_pem,omitempty"`
	Certificates   []*CreateListenerCertificate `json:"certificates,omitempty"`

	// ChainPEM deprecated — use certificates[]
	ChainPEM             *string `json:"chain_pem,omitempty"`
	DefaultTargetGroupID *string `json:"default_target_group_id,omitempty"`

	// Exposure Which LB addresses are bound. Defaults to 'both'; pick private_only
	// when the LB has no FIP yet.
	//
	// One of: "public_only", "private_only", "both".
	Exposure *string `json:"exposure,omitempty"`

	// Required.
	Port int `json:"port"`

	// PrivateKeyPEM deprecated — use certificates[]
	PrivateKeyPEM *string `json:"private_key_pem,omitempty"`

	// One of: "http", "https", "tcp", "udp".
	//
	// Required.
	Protocol string `json:"protocol"`
	Tags     Tags   `json:"tags,omitempty"`
}

type CreateLoadBalancerRequest struct {
	// FlavorID compute flavor for each LB instance.
	//
	// Required.
	FlavorID string `json:"flavor_id"`

	// FloatingIPID Optional FIP attached on create for public exposure.
	FloatingIPID *string `json:"floating_ip_id,omitempty"`

	// KeyNames platform-operator break-glass only. Stamps SSH keypairs onto the
	// replica VMs, which run the platform's own envoy and lbaas-agent; a
	// tenant reaches the load balancer over its VIP, never over SSH.
	// Accepted only from the platform account and only with
	// `loadbalancer:StampBreakGlassKeys` — any other account setting it
	// is rejected with 400 `INVALID_INPUT`.
	KeyNames []string `json:"key_names,omitempty"`

	// Name 1..127 chars of [A-Za-z0-9._-]
	//
	// Required.
	Name string `json:"name"`

	// ReplicaCount number of LB compute instances. Defaults to 1; pick >=2 for HA.
	ReplicaCount *int `json:"replica_count,omitempty"`

	// SecurityGroupIDs security groups attached to every replica NIC (AWS ALB shape). A VPC
	// NIC with no security group denies all data traffic, so the listener
	// port(s) must be opened by a security group listed here. Re-applied
	// to replacement replicas. The LB's own control-plane path (agent
	// config + heartbeat via the metadata endpoint) is always-allowed and
	// needs none.
	//
	// Required.
	SecurityGroupIDs []string `json:"security_group_ids"`

	// SubnetID subnet the LB instances attach to. The virtual IP is allocated from
	// this subnet.
	//
	// Required.
	SubnetID string `json:"subnet_id"`
	Tags     Tags   `json:"tags,omitempty"`

	// One of: "application", "network".
	//
	// Required.
	Type string `json:"type"`

	// VPCID VPC the LB will live in. Must match subnet_id's VPC.
	//
	// Required.
	VPCID string `json:"vpc_id"`
}

type CreateRuleRequest struct {
	// Required.
	Conditions []*RuleCondition `json:"conditions"`

	// Required.
	Priority int `json:"priority"`

	// Required.
	TargetGroupID string `json:"target_group_id"`
}

type CreateTargetGroupRequest struct {
	HealthCheck *HealthCheck `json:"health_check,omitempty"`

	// InstancePoolID compute instance pool to draw backends from. Required when
	// target_mode=pool and must belong to the calling account; ignored
	// otherwise.
	InstancePoolID *string `json:"instance_pool_id,omitempty"`

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
	// instance_pool_id; the group is forced to target_type=instance, and
	// attaching targets to it is rejected.
	//
	// One of: "static", "pool".
	TargetMode *string `json:"target_mode,omitempty"`

	// One of: "ip", "instance", "function".
	TargetType *string `json:"target_type,omitempty"`
}

type HealthCheck struct {
	HealthyThreshold int    `json:"healthy_threshold,omitempty"`
	IntervalSec      int    `json:"interval_sec,omitempty"`
	Matcher          string `json:"matcher,omitempty"`
	Path             string `json:"path,omitempty"`
	Port             int    `json:"port,omitempty"`

	// Protocol defaults to the target group protocol when omitted
	//
	// One of: "http", "https", "tcp", "udp".
	Protocol           string `json:"protocol,omitempty"`
	TimeoutSec         int    `json:"timeout_sec,omitempty"`
	UnhealthyThreshold int    `json:"unhealthy_threshold,omitempty"`
}

type Listener struct {
	// Certificates HTTPS listeners only. One entry per attached certificate; the right
	// cert is picked per-connection by matching the client's SNI against
	// each cert's SAN. The entry flagged is_default serves traffic that
	// doesn't match any other SNI (or clients that omit SNI). PEM material
	// is NOT echoed — the listener stores its own copy fetched at attach
	// time.
	Certificates         []*ListenerCertificate `json:"certificates,omitempty"`
	CreatedAt            time.Time              `json:"created_at"`
	DefaultTargetGroupID string                 `json:"default_target_group_id,omitempty"`

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
	AccountID string    `json:"account_id"`
	CreatedAt time.Time `json:"created_at"`

	// CRN IAM resource CRN
	CRN string `json:"crn"`

	// DNSName convenience hostname auto-published for the load balancer,
	// `{name}.{account-handle}.lb.{region}.{base-domain}`. Resolves to the
	// floating IP on an internet-facing LB and to the private VIP
	// otherwise. Omitted in regions where auto-DNS is not configured —
	// the VIP and FIP stay authoritative either way.
	DNSName      string `json:"dns_name,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`

	// FlavorID compute flavor each LB instance runs on. Must be a
	// loadbalancer-family flavor.
	FlavorID string `json:"flavor_id"`

	// FloatingIPID Optional FIP attached for public exposure. NULL ⇒ private-only LB.
	FloatingIPID string `json:"floating_ip_id,omitempty"`
	ID           string `json:"id"`
	Name         string `json:"name"`

	// PublicVipV6 Public IPv6 GUA for the load balancer — the v6 analogue of a
	// floating IP (IPv6 has no NAT, so this address is itself the public
	// ingress, anycast-advertised). Set on an internet-facing LB in a
	// dual-stack subnet.
	PublicVipV6 string `json:"public_vip_v6,omitempty"`

	// ReplicaCount number of LB compute instances. >=2 for HA.
	ReplicaCount int `json:"replica_count"`

	// One of: "provisioning", "active", "error", "deleting".
	Status string `json:"status"`

	// SubnetID subnet hosting the LB's compute instances and VIP.
	SubnetID string `json:"subnet_id"`
	Tags     Tags   `json:"tags"`

	// Type ALB-shape (L7) vs NLB-shape (L4)
	//
	// One of: "application", "network".
	Type      string    `json:"type"`
	UpdatedAt time.Time `json:"updated_at"`

	// VipV4 Virtual IP for the load balancer; traffic is distributed to backends
	// per connection.
	VipV4 string `json:"vip_v4,omitempty"`

	// VipV6 Internal IPv6 VIP (set when the subnet is dual-stack).
	VipV6 string `json:"vip_v6,omitempty"`

	// VPCID VPC the LB lives in.
	VPCID string `json:"vpc_id"`
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
	ProxyOk      bool `json:"proxy_ok"`
	ReplicaIndex int  `json:"replica_index"`

	// Status the liveness view folded into one word: 'initializing' (the agent
	// has never reported — boot still in flight), 'healthy'
	// (heartbeating and the proxy is serving), 'unhealthy' (heartbeating
	// but the proxy is down).
	//
	// One of: "initializing", "healthy", "unhealthy".
	Status string `json:"status"`
}

type Rule struct {
	Conditions    []*RuleCondition `json:"conditions"`
	CreatedAt     time.Time        `json:"created_at"`
	ID            string           `json:"id"`
	ListenerID    string           `json:"listener_id"`
	Priority      int              `json:"priority"`
	TargetGroupID string           `json:"target_group_id"`
	UpdatedAt     time.Time        `json:"updated_at"`
}

type RuleCondition struct {
	// One of: "host", "path", "header", "query", "method".
	Field string `json:"field"`

	// Name header or query key name
	Name string `json:"name,omitempty"`

	// One of: "exact", "prefix", "glob", "regex".
	Op     string   `json:"op"`
	Values []string `json:"values"`
}

// SessionAffinity sticky sessions: pin a client to one backend in the group instead of
// balancing each request independently. Off unless you ask for it.
//
// The data plane hashes the key consistently, so adding or losing a
// backend only moves the clients that backend was serving.
type SessionAffinity struct {
	// CookieName cookie the load balancer sets and hashes. `type=cookie` only; the
	// field is dropped for the other types.
	CookieName string `json:"cookie_name,omitempty"`

	// DurationSec how long that cookie lives, up to 7 days. `type=cookie` only.
	DurationSec int `json:"duration_sec,omitempty"`

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
	Name           string `json:"name"`
	Port           int    `json:"port"`

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

// UpdateListenerRequest patch a listener. Certificate rotation requires *all* of
// certificate_crn + certificate_pem + private_key_pem to be set in the
// same request; chain_pem is optional. Set
// clear_default_target_group=true to remove the default; otherwise
// omitting default_target_group_id leaves it unchanged.
type UpdateListenerRequest struct {
	CertificateCRN          *string `json:"certificate_crn,omitempty"`
	CertificatePEM          *string `json:"certificate_pem,omitempty"`
	ChainPEM                *string `json:"chain_pem,omitempty"`
	ClearDefaultTargetGroup *bool   `json:"clear_default_target_group,omitempty"`
	DefaultTargetGroupID    *string `json:"default_target_group_id,omitempty"`

	// Exposure mutate which addresses are bound. Omit to leave unchanged.
	//
	// One of: "public_only", "private_only", "both".
	Exposure      *string `json:"exposure,omitempty"`
	PrivateKeyPEM *string `json:"private_key_pem,omitempty"`
	Tags          Tags    `json:"tags,omitempty"`
}

type UpdateLoadBalancerRequest struct {
	// FlavorID resize each replica to a different compute flavor. Must be a
	// loadbalancer-family flavor.
	//
	// A running instance cannot change size in place, so the request
	// records the new size and returns; the replicas already up are then
	// replaced one at a time in the background. The load balancer
	// temporarily runs one replica over replica_count while it does: the
	// extra replica comes up on the new flavor and starts serving before
	// any replica on the old one is retired, so the number serving never
	// drops below replica_count — a resize does not cost you capacity,
	// at any replica count.
	//
	// Expect it to take several minutes, and poll GET
	// /v1/load-balancers/{id}/replicas to watch: a replica has been
	// replaced when its instance_id changes, and the resize is done when
	// every flavor_id there matches this one.
	//
	// The one exception is a load balancer already at the maximum of 10
	// replicas, which has nowhere to grow. There the replicas are replaced
	// in place and 9 serve while each replacement boots.
	//
	// Rejected up front if the account does not have the compute quota for
	// the replacement replica, so a resize cannot half-apply and leave the
	// load balancer short.
	FlavorID *string `json:"flavor_id,omitempty"`
	Name     *string `json:"name,omitempty"`

	// ReplicaCount resize the set of load balancer instances. Scale-out provisions the
	// new replicas in sequence; scale-in removes the highest-indexed
	// replicas best-effort. 1..10.
	ReplicaCount *int `json:"replica_count,omitempty"`
	Tags         Tags `json:"tags,omitempty"`
}

// UpdateRuleRequest full replace of the rule — priority, conditions, and target group
// must all be supplied (same shape as create).
type UpdateRuleRequest struct {
	// Required.
	Conditions []*RuleCondition `json:"conditions"`

	// Required.
	Priority int `json:"priority"`

	// Required.
	TargetGroupID string `json:"target_group_id"`
}

type UpdateTargetGroupRequest struct {
	HealthCheck *HealthCheck `json:"health_check,omitempty"`
	Name        *string      `json:"name,omitempty"`

	// ProxyProtocol Toggle PROXY v2 framing on upstream connections. Omitting the field
	// leaves the current setting; setting true/false flips it explicitly.
	ProxyProtocol *bool `json:"proxy_protocol,omitempty"`

	// SessionAffinity replaces the stickiness config. Omitting the field leaves it alone;
	// turning it off is an explicit `{"type": "none"}`.
	SessionAffinity *SessionAffinity `json:"session_affinity,omitempty"`
	Tags            Tags             `json:"tags,omitempty"`
}
