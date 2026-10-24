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
	// Port per-target backend port override. When omitted, the target inherits
	// the target group port. Pool-derived targets always inherit the group
	// port.
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

	// One of: "pending", "disabled", "stable", "scaling", "waiting", "warming_up", "metrics_unavailable", "stabilizing", "cooldown", "draining", "scale_in_blocked".
	Status string `json:"status"`
}

type AutoscalingStatusHistoryItem struct {
	At     time.Time `json:"at"`
	From   int       `json:"from"`
	Reason string    `json:"reason"`
	To     int       `json:"to"`
}

type Certificate struct {
	// CertificatePEM PEM-encoded leaf certificate. Empty until active.
	CertificatePEM string `json:"certificate_pem,omitempty"`

	// ChainPEM PEM-encoded intermediate chain.
	ChainPEM string `json:"chain_pem,omitempty"`

	// Challenges Per-domain CNAME delegation state — one entry per domain on the
	// cert. While status=pending_dns issuance waits for every challenge's
	// `verified` to flip true; use `expected_cname` and `our_dns` to tell
	// which records need to be added at the registrar.
	Challenges []*CertificateChallenge `json:"challenges,omitempty"`

	// CRN name-based, so an IAM policy can wildcard a naming convention
	// (`crn:certificate::my-account:certificate/prod-*`). The region slot
	// is empty for compatibility; certificate storage and KMS material are
	// regional.
	CRN       string    `json:"crn,omitempty"`
	Domains   []string  `json:"domains,omitempty"`
	ExpiresAt time.Time `json:"expires_at,omitempty"`

	// Faults active faults, ordered newest first. Empty when healthy. Renewal
	// failures are warnings while valid certificate material still serves.
	// Codes: CERTIFICATE_ISSUANCE_START_FAILED (issuance could not be
	// scheduled), CERTIFICATE_ISSUANCE_FAILED (the signing request
	// failed), CERTIFICATE_RENEWAL_FAILED (renewal did not complete; a
	// warning while valid material is still serving, an error once it has
	// expired), CERTIFICATE_REVOCATION_FAILED (revocation did not
	// complete).
	Faults []*Fault `json:"faults"`

	// Fingerprint Hex SHA-256 of the leaf's DER — the certificate's material
	// version. Changes on every rotation; consumers use it to know when to
	// re-fetch material and to verify they fetched the intended
	// generation. Empty until the cert has a leaf.
	Fingerprint  string                  `json:"fingerprint,omitempty"`
	ID           string                  `json:"id,omitempty"`
	IssuedAt     time.Time               `json:"issued_at,omitempty"`
	KeyAlgorithm CertificateKeyAlgorithm `json:"key_algorithm,omitempty"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	Name   string            `json:"name,omitempty"`
	Source CertificateSource `json:"source,omitempty"`
	Status CertificateStatus `json:"status,omitempty"`
	Tags   Tags              `json:"tags,omitempty"`
}

type CertificateChallenge struct {
	// CnameRecordName Full LHS of the CNAME record the customer needs to add. For wildcard
	// SANs this is the parent name (`_acme-challenge.example.com.`), not
	// the literal SAN — wildcards validate at their parent under RFC
	// 8555 §8.4.
	CnameRecordName string `json:"cname_record_name,omitempty"`

	// Domain the cert SAN this challenge belongs to (as the customer wrote it).
	Domain string `json:"domain,omitempty"`

	// ExpectedCname Target FQDN (RHS of the CNAME). Hosted in the platform's validation
	// zone, where the per-order TXT is published during issuance.
	ExpectedCname string `json:"expected_cname,omitempty"`

	// Faults active verification warnings only; empty when healthy. Successful
	// verification resolves verification faults while retaining history.
	// verified records a successful observation and remains true if a
	// later renewal observes a DNS failure. A CNAME mismatch records
	// CERTIFICATE_DNS_VERIFICATION_FAILED as a warning; verified does not
	// move. Internal history uses the parent CRN followed by
	// /challenge/<stored-uuid>; no separate endpoint is exposed.
	Faults []*Fault `json:"faults"`

	// OurDNS true when the domain is hosted on the platform DNS service and the
	// CNAME was created automatically. False means the customer owns the
	// zone and must add the CNAME at their registrar.
	OurDNS bool `json:"our_dns,omitempty"`

	// Verified true once the CNAME has resolved to expected_cname; cert issuance
	// only proceeds when every challenge is verified.
	Verified   bool      `json:"verified,omitempty"`
	VerifiedAt time.Time `json:"verified_at,omitempty"`
}

type CertificateKeyAlgorithm string

// Values CertificateKeyAlgorithm accepts.
const (
	CertificateKeyAlgorithmEcdsaP256 CertificateKeyAlgorithm = "ecdsa-p256"
	CertificateKeyAlgorithmEcdsaP384 CertificateKeyAlgorithm = "ecdsa-p384"
	CertificateKeyAlgorithmRSA2048   CertificateKeyAlgorithm = "rsa-2048"
	CertificateKeyAlgorithmRSA4096   CertificateKeyAlgorithm = "rsa-4096"
)

type CertificateSource string

// Values CertificateSource accepts.
const (
	CertificateSourceAcme     CertificateSource = "acme"
	CertificateSourceUploaded CertificateSource = "uploaded"
)

// CertificateStatus error if and only if an active error-severity fault exists.
type CertificateStatus string

// Values CertificateStatus accepts.
const (
	CertificateStatusPendingDNS CertificateStatus = "pending_dns"
	CertificateStatusPending    CertificateStatus = "pending"
	CertificateStatusActive     CertificateStatus = "active"
	CertificateStatusError      CertificateStatus = "error"
	CertificateStatusExpired    CertificateStatus = "expired"
	CertificateStatusRevoked    CertificateStatus = "revoked"
)

// CreateListenerRequest target group relationships accept UUID, CRN or exact account-scoped
// names in this region. HTTPS listeners require >=1 certificate, named
// by CRN, UUID or exact account-scoped name in `certificates`; the first
// entry becomes the default (the fallback when SNI doesn't match). No
// key material is accepted — the agent fetches it from the certificate
// service against the CRN.
type CreateListenerRequest struct {
	Certificates       []string `json:"certificates,omitempty"`
	DefaultTargetGroup *string  `json:"default_target_group,omitempty"`

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
// by syntax. VPC, subnets and security groups must belong to the caller
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

	// FloatingIPs existing free floating IPs from this account and region, at most one
	// per family and visibility (private/public, IPv4/IPv6). Private
	// addresses must belong to any selected subnet. Missing private
	// families are allocated automatically from the first selected subnet
	// for each family enabled there. Public addresses are optional and
	// require a matching-family default route to an internet gateway in
	// every selected subnet; NAT and egress-only gateways do not qualify.
	// Public IPv6 requires every selected subnet to support IPv6.
	// Pool-owned or attached addresses are unavailable. On deletion,
	// supplied addresses are detached and retained; automatic private
	// allocations are released.
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

	// Subnets one or more subnet selectors in the same VPC, account and region.
	// Each instance has one interface; new instances are placed in the
	// least-populated eligible subnet. The first subnet supplies automatic
	// private floating IP allocations. Placement is fixed at creation.
	//
	// Required.
	Subnets []string `json:"subnets"`
	Tags    Tags     `json:"tags,omitempty"`

	// One of: "application", "network".
	//
	// Required.
	Type string `json:"type"`

	// VPC the LB will live in. Every selected subnet must belong to this
	// VPC.
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

// CurrentState where an instance actually is. The transitional states live here, not
// on desired_state — nobody asks for `stopping`.
type CurrentState string

// Values CurrentState accepts.
const (
	CurrentStatePending   CurrentState = "pending"
	CurrentStateBuilding  CurrentState = "building"
	CurrentStateRunning   CurrentState = "running"
	CurrentStateStopping  CurrentState = "stopping"
	CurrentStateStopped   CurrentState = "stopped"
	CurrentStateRebooting CurrentState = "rebooting"
	CurrentStateMigrating CurrentState = "migrating"
	CurrentStateDeleting  CurrentState = "deleting"
	CurrentStateDeleted   CurrentState = "deleted"
	CurrentStateError     CurrentState = "error"
	CurrentStateCrashed   CurrentState = "crashed"
	CurrentStatePaused    CurrentState = "paused"
	CurrentStateSuspended CurrentState = "suspended"
)

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

type FloatingIPRef struct {
	Address string `json:"address"`
	CRN     string `json:"crn"`
	ID      string `json:"id"`

	// One of: "public", "private".
	Visibility string `json:"visibility"`
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

// IntervalSec zero uses the default.
type IntervalSec = int

type Listener struct {
	CreatedAt time.Time `json:"created_at"`

	// CRN Parent-scoped CRN with immutable load balancer name and child UUID
	// components.
	CRN string `json:"crn"`

	// DefaultCertificate default certificate for an HTTPS listener, used when no SNI matches
	// or the client sends none. Null when absent. List all attached
	// certificates through the listener certificates endpoint.
	DefaultCertificate *ResourceRef `json:"default_certificate"`

	// DefaultTargetGroup target group used when no routing rule matches, or null when none is
	// configured.
	DefaultTargetGroup *ResourceRef `json:"default_target_group"`

	// Exposure Which LB addresses this listener binds. 'public_only' and 'both'
	// require the LB to carry a floating IP; if the FIP is detached later,
	// the listener is dropped until a FIP is re-attached or exposure is
	// flipped to private_only.
	//
	// One of: "public_only", "private_only", "both".
	Exposure string `json:"exposure"`
	ID       string `json:"id"`
	Port     int    `json:"port"`

	// One of: "http", "https", "tcp", "udp".
	Protocol  string    `json:"protocol"`
	Tags      Tags      `json:"tags"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoadBalancer struct {
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

	// Flavor regional loadbalancer-family compute flavor assigned to replicas.
	Flavor *ResourceRef `json:"flavor"`

	// FloatingIPs public and private floating IP references, including automatic
	// private allocations. Each includes id, crn, visibility, and address.
	// Infer IPv4 or IPv6 from the address.
	FloatingIPs []*FloatingIPRef `json:"floating_ips"`
	ID          string           `json:"id"`

	// MaxCount upper capacity bound including rollout surge.
	MaxCount int `json:"max_count"`

	// MinCount lower capacity bound.
	MinCount int `json:"min_count"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	Name string `json:"name"`

	// ReplicaCount deprecated alias of desired_count.
	ReplicaCount int `json:"replica_count"`

	// RolloutSurge temporary extra capacity within max_count; does not change
	// desired_count.
	RolloutSurge bool `json:"rollout_surge,omitempty"`

	// One of: "provisioning", "active", "error", "deleting".
	Status string `json:"status"`

	// Subnets eligible subnets for member placement. Each instance has one
	// interface, and new instances use the least-populated eligible
	// subnet.
	Subnets []*ResourceRef `json:"subnets"`
	Tags    Tags           `json:"tags"`

	// Type ALB-shape (L7) vs NLB-shape (L4)
	//
	// One of: "application", "network".
	Type      string    `json:"type"`
	UpdatedAt time.Time `json:"updated_at"`
}

// LoadBalancerInstance a compute instance backing the load balancer, extended with proxy
// health and retirement details. During a resize, compare each
// instance's flavor.id with the load balancer's flavor.id. An extra
// instance may be present while the replacement starts serving and the
// old instance drains.
type LoadBalancerInstance struct {
	AgentVersion string    `json:"agent_version,omitempty"`
	CreatedAt    time.Time `json:"created_at"`

	// CRN Cloud Resource Name
	CRN string `json:"crn"`

	// CurrentState where the instance actually is. Read this one to answer "is it up"
	// — the transitional states live here, not on desired_state, because
	// nobody asks for `stopping`.
	//
	// desired_state=running with current_state=stopped is an instance that
	// was asked to start and has not come up yet.
	CurrentState CurrentState `json:"current_state"`
	Description  string       `json:"description,omitempty"`

	// DesiredState what was asked for. Only three values, because there are only three
	// things you can ask an instance to be: Create/Start/Reboot ask for
	// running, Stop for stopped, Delete for deleted.
	//
	// One of: "running", "stopped", "deleted".
	DesiredState string `json:"desired_state"`

	// Faults active faults ordered by last_at descending, then internal history
	// id descending for a stable tie-breaker. Healthy resources return [].
	Faults []*Fault `json:"faults"`

	// Flavor identity. Every instance has a flavor; query the flavor
	// resource for CPU and RAM.
	Flavor *ResourceRef `json:"flavor"`

	// IAMRole Attached IAM role identity, or null when no role is attached.
	// Deleted roles retain their identity without name.
	IAMRole *ResourceRef `json:"iam_role"`
	ID      string       `json:"id"`

	// Image pinned source image identity, or null for a volume-only boot.
	Image *ResourceRef `json:"image"`

	// LastSeen omitted when this instance has not reported yet
	LastSeen   time.Time  `json:"last_seen,omitempty"`
	LaunchedAt *time.Time `json:"launched_at,omitempty"`
	Metadata   Metadata   `json:"metadata,omitempty"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	Name string `json:"name"`

	// ProxyOk whether the proxy reported ready at the last health report
	ProxyOk    bool        `json:"proxy_ok"`
	Retirement *Retirement `json:"retirement,omitempty"`

	// Status proxy health: initializing before the first report, healthy when the
	// last report says the proxy is serving, unhealthy when it is down,
	// and draining when the instance is retiring. See last_seen for the
	// age of the last observation.
	//
	// One of: "initializing", "healthy", "unhealthy", "draining".
	Status string `json:"status"`
	Tags   Tags   `json:"tags,omitempty"`

	// TaskState in-flight transition, if any; null when settled.
	TaskState    *string    `json:"task_state,omitempty"`
	TerminatedAt *time.Time `json:"terminated_at,omitempty"`
	UpdatedAt    time.Time  `json:"updated_at"`

	// UserData base64-encoded cloud-init user-data supplied at launch.
	UserData string `json:"user_data,omitempty"`
}

// Matcher HTTP status codes from 100 to 599; comma-separated codes or inclusive
// ranges. Empty uses 200.
type Matcher = string

type Metadata = map[string]string

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

type ResourceRef struct {
	CRN string `json:"crn"`
	ID  string `json:"id"`

	// Name display label at response time — the resource's name where it has
	// one, otherwise its human label (an invoice number, an account
	// display name). Absent when the target is gone. Not a lookup key.
	Name string `json:"name,omitempty"`
}

type Retirement struct {
	AgentAcknowledgedAt time.Time `json:"agent_acknowledged_at,omitempty"`
	DrainSeconds        int       `json:"drain_seconds"`

	// DrainUntil earliest deletion time; absent while withdrawal is pending.
	DrainUntil  time.Time `json:"drain_until,omitempty"`
	RequestedAt time.Time `json:"requested_at"`
}

type Rule struct {
	Conditions []*RuleCondition `json:"conditions"`
	CreatedAt  time.Time        `json:"created_at"`

	// CRN Parent-scoped CRN with immutable load balancer name and child UUID
	// components.
	CRN         string       `json:"crn"`
	ID          string       `json:"id"`
	Priority    int          `json:"priority"`
	TargetGroup *ResourceRef `json:"target_group"`
	UpdatedAt   time.Time    `json:"updated_at"`
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

type Tags = map[string]string

// Target a registered backend. Exactly one of instance or address is returned.
type Target struct {
	// Address Canonical IPv4 or IPv6 address for an IP target. Omitted for
	// instance targets.
	Address   string    `json:"address,omitempty"`
	CreatedAt time.Time `json:"created_at"`

	// One of: "initial", "healthy", "unhealthy", "draining".
	Health string `json:"health"`
	ID     string `json:"id"`

	// Instance compute instance backing an instance target or a pool-derived
	// target. Omitted for IP targets.
	Instance   *ResourceRef `json:"instance,omitempty"`
	LastSeenAt time.Time    `json:"last_seen_at,omitempty"`

	// Port per-target backend port override. When omitted, the target inherits
	// the target group port. Pool-derived targets always inherit the group
	// port.
	Port      int       `json:"port,omitempty"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TargetGroup struct {
	CreatedAt   time.Time    `json:"created_at"`
	CRN         string       `json:"crn"`
	HealthCheck *HealthCheck `json:"health_check"`
	ID          string       `json:"id"`

	// InstancePool compute instance pool backing the group. Omitted for static target
	// groups.
	InstancePool *ResourceRef `json:"instance_pool,omitempty"`

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
	// instance_pool, so scaling the pool moves the backends with it.
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

// UpdateListenerRequest patch a listener. Use default_certificate to select an already
// attached certificate as the default. Certificate renewals are picked
// up automatically; manage attachments through the listener's
// certificates collection.
//
// Omit default_target_group to leave it unchanged, send a string
// reference to replace it, or send null to clear it.
type UpdateListenerRequest struct {
	// DefaultCertificate UUID, CRN, or exact name of a certificate already attached to this
	// listener. Makes it the default and demotes the previous default
	// atomically. Does not attach new certificates.
	DefaultCertificate *string `json:"default_certificate,omitempty"`

	// DefaultTargetGroup UUID, CRN or exact account-scoped target group name. Null clears the
	// default; omission leaves it unchanged.
	DefaultTargetGroup *string `json:"default_target_group,omitempty"`

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
	// /v1/load-balancers/{id}/instances to watch: a replica has been
	// replaced when its id changes, and the resize is done when every
	// flavor.id there matches the requested flavor.
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
