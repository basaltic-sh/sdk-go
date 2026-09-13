// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package compute

import (
	"time"
)

type AttachInstanceNICAttachment struct {
	BootIndex   int    `json:"boot_index,omitempty"`
	InterfaceID string `json:"interface_id,omitempty"`
	IP          string `json:"ip,omitempty"`
	MAC         string `json:"mac,omitempty"`
}

// AttachInstanceNICRequest exactly one of subnet (provision a fresh NIC) or interface (attach an
// existing standalone interface) must be set.
type AttachInstanceNICRequest struct {
	// Interface existing standalone interface UUID or complete VPC/subnet/interface
	// CRN. Bare names lack the subnet parent and are rejected. It keeps
	// its address, MAC, and security groups; detach returns it to
	// standalone instead of destroying it.
	Interface      *string  `json:"interface,omitempty"`
	IPAddress      *string  `json:"ip_address,omitempty"`
	MAC            *string  `json:"mac,omitempty"`
	SecurityGroups []string `json:"security_groups,omitempty"`

	// Subnet UUID or complete VPC/subnet CRN; a bare name requires a VPC
	// parent and is rejected here.
	Subnet *string `json:"subnet,omitempty"`
}

type AttachInstanceVolumeAttachment struct {
	Device     string `json:"device,omitempty"`
	InstanceID string `json:"instance_id,omitempty"`
	VolumeID   string `json:"volume_id,omitempty"`
}

type AttachInstanceVolumeRequest struct {
	// Device optional device-name override; auto-picks the next free slot
	// (vdb/vdc/…) when omitted.
	Device *string `json:"device,omitempty"`

	// Fstype filesystem the in-guest agent formats the disk with, and only when
	// `mount_path` is set and the disk is blank. Rejected with 400 if it
	// is neither value.
	//
	// One of: "ext4", "xfs".
	Fstype *string `json:"fstype,omitempty"`

	// MountPath when set, the in-guest agent formats the disk (only if blank) and
	// mounts it at this path. Empty attaches the block device only.
	MountPath *string `json:"mount_path,omitempty"`

	// Volume account-scoped volume reference (UUID, CRN or exact name).
	//
	// Required.
	Volume string `json:"volume"`
}

type CreateKeypairKeypair struct {
	CreatedAt time.Time `json:"created_at,omitempty"`

	// CRN Cloud Resource Name
	CRN         string `json:"crn,omitempty"`
	Fingerprint string `json:"fingerprint,omitempty"`
	ID          string `json:"id,omitempty"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	Name string `json:"name,omitempty"`

	// PrivateKey private key (only returned when keypair is generated)
	PrivateKey string `json:"private_key,omitempty"`

	// PublicKey SSH public key
	PublicKey string `json:"public_key,omitempty"`
	Tags      Tags   `json:"tags,omitempty"`
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

// Flavor a compute size (vCPU + RAM). A flavor carries no disk size — the
// boot disk is a customer volume sized at launch, floored by the image's
// min_disk_gb.
type Flavor struct {
	// Class host-pool routing. "shared" oversubscribes CPU for higher density;
	// "dedicated" pins each vCPU 1:1 to a physical core.
	//
	// One of: "shared", "dedicated".
	Class     string    `json:"class,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`

	// CRN Cloud Resource Name
	CRN         string `json:"crn,omitempty"`
	Description string `json:"description,omitempty"`

	// Family which product can book the flavor. "general" flavors are for regular
	// instances and instance pools; "loadbalancer" and "database" flavors
	// are reserved for the managed products (their nodes are platform-
	// operated and priced accordingly) and cannot be used for regular
	// instances.
	//
	// One of: "general", "loadbalancer", "database".
	Family string `json:"family,omitempty"`
	ID     string `json:"id,omitempty"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	Name string `json:"name,omitempty"`

	// RAMMB RAM in MB
	RAMMB int `json:"ram_mb,omitempty"`

	// One of: "active", "disabled".
	Status    string    `json:"status,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// Vcpus number of virtual CPUs
	Vcpus int `json:"vcpus,omitempty"`
}

type FloatingIP struct {
	// AttachedToInterfaceID legacy single-binding field: the sole interface this floating IP is
	// bound to, or null when unattached OR when it has more than one
	// member (an anycast floating IP). `members` is authoritative.
	AttachedToInterfaceID string    `json:"attached_to_interface_id,omitempty"`
	CreatedAt             time.Time `json:"created_at"`
	CRN                   string    `json:"crn"`
	Description           string    `json:"description,omitempty"`
	Family                IPFamily  `json:"family"`
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

	// IPAddress the public address, in the family it was allocated in.
	IPAddress string `json:"ip_address"`

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

type GetConsoleOutputResult struct {
	// Output the transcript as plain text, newlines included. Empty when the
	// instance has not booted yet.
	Output string `json:"output"`

	// Truncated true when the transcript was longer than the requested size and its
	// BEGINNING was dropped to fit. The end is always kept.
	Truncated bool `json:"truncated"`
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

type Image struct {
	Architecture string            `json:"architecture"`
	Attributes   map[string]string `json:"attributes,omitempty"`
	CreatedAt    time.Time         `json:"created_at"`
	CRN          string            `json:"crn"`
	Description  string            `json:"description,omitempty"`

	// EOLDate the day this image's OS release stops receiving free security
	// updates for a default install. Absent when nobody has recorded one
	// — which means unknown, not "supported indefinitely".
	//
	// Platform images are withdrawn from the catalog a grace period after
	// this date. They stay bootable by id until then, and the date is
	// published well ahead of it so you can plan the move.
	EOLDate string `json:"eol_date,omitempty"`

	// One of: "qcow2", "raw", "vmdk", "vhd", "vhdx", "vdi".
	Format string `json:"format"`
	ID     string `json:"id"`

	// ImportError for a qcow2 upload that failed conversion (status=error), the
	// reason. Absent otherwise.
	ImportError string `json:"import_error,omitempty"`

	// IsCurrent whether this is the version resolve-by-name returns for its (name,
	// architecture) — i.e. the name's current tag target.
	IsCurrent bool `json:"is_current,omitempty"`
	MinDiskGB int  `json:"min_disk_gb,omitempty"`
	MinRAMMB  int  `json:"min_ram_mb,omitempty"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	Name      string `json:"name"`
	OS        string `json:"os,omitempty"`
	OSVersion string `json:"os_version,omitempty"`
	SizeBytes int64  `json:"size_bytes,omitempty"`

	// One of: "pending", "importing", "active", "error", "hidden".
	Status    string    `json:"status"`
	Tags      Tags      `json:"tags,omitempty"`
	UpdatedAt time.Time `json:"updated_at"`

	// Version the build's identity within its name. Unique there: a name is a
	// movable tag, so it can't also be what tells two builds apart.
	// Stamped as a UTC timestamp when the uploader didn't choose one.
	Version string `json:"version"`

	// One of: "public", "private".
	Visibility string `json:"visibility"`
}

type ImageCreateRequest struct {
	Architecture *string           `json:"architecture,omitempty"`
	Attributes   map[string]string `json:"attributes,omitempty"`

	// Current make this the current version for its (name, architecture) once
	// active.
	Current     *bool   `json:"current,omitempty"`
	Description *string `json:"description,omitempty"`

	// EOLDate the day this release stops receiving free security updates. Omit it
	// and the image inherits the date the name's current version carries,
	// so re-publishing a tag can't quietly stop tracking its release.
	EOLDate *string `json:"eol_date,omitempty"`

	// Format source disk format at source_url; converted to the raw base on
	// import.
	//
	// One of: "qcow2", "raw", "vmdk", "vhd", "vhdx", "vdi".
	//
	// Required.
	Format    string `json:"format"`
	MinDiskGB *int   `json:"min_disk_gb,omitempty"`
	MinRAMMB  *int   `json:"min_ram_mb,omitempty"`

	// Name movable tag name (e.g. debian-13); the new image becomes its current
	// version. Names are shared across a tag's builds — what identifies
	// one build is `version`. Resource names must not start with the
	// literal crn: prefix or be UUIDs (canonical, compact, braced, or
	// urn:uuid: forms, in either case).
	//
	// Required.
	Name      string  `json:"name"`
	OS        *string `json:"os,omitempty"`
	OSVersion *string `json:"os_version,omitempty"`

	// SourceURL presigned https GET URL to the disk in an object store you control.
	// Fetched once by the import worker (which rejects private/link-local
	// targets). Not retained after import.
	//
	// Required.
	SourceURL string `json:"source_url"`
	Tags      Tags   `json:"tags,omitempty"`

	// Version identifies this build within `name`, and must be unique there —
	// re-publishing a version that a tag already carries is a 409. Omit it
	// and the server stamps a UTC timestamp, so every build is addressable
	// as `name:version` whether or not you labelled it.
	Version *string `json:"version,omitempty"`

	// One of: "public", "private".
	Visibility *string `json:"visibility,omitempty"`
}

type ImageUpdateRequest struct {
	Attributes map[string]string `json:"attributes,omitempty"`

	// Current switch the resolve-by-name pointer for this image's name. true
	// promotes this version to current (the switch / rollback action) and
	// demotes whatever else was current for the same (name, architecture);
	// false clears the pointer. Only active images can be made current.
	Current     *bool   `json:"current,omitempty"`
	Description *string `json:"description,omitempty"`

	// EOLDate set the release's end-of-life date. An explicit null clears it;
	// omitting the field leaves it unchanged. Clearing matters because the
	// catalog withdraws platform images on this date — one recorded by
	// mistake has to be removable.
	EOLDate *string `json:"eol_date,omitempty"`
	Tags    Tags    `json:"tags,omitempty"`

	// One of: "public", "private".
	Visibility *string `json:"visibility,omitempty"`
}

type Instance struct {
	CreatedAt time.Time `json:"created_at,omitempty"`

	// CRN Cloud Resource Name
	CRN string `json:"crn,omitempty"`

	// CurrentState where the instance actually is. Read this one to answer "is it up"
	// — the transitional states live here, not on desired_state, because
	// nobody asks for `stopping`.
	//
	// desired_state=running with current_state=stopped is an instance that
	// was asked to start and has not come up yet.
	CurrentState CurrentState `json:"current_state,omitempty"`
	Description  string       `json:"description,omitempty"`

	// DesiredState what was asked for. Only three values, because there are only three
	// things you can ask an instance to be: Create/Start/Reboot ask for
	// running, Stop for stopped, Delete for deleted.
	//
	// One of: "running", "stopped", "deleted".
	DesiredState string         `json:"desired_state,omitempty"`
	Fault        *InstanceFault `json:"fault,omitempty"`

	// Flavor resolved flavor (compute size) the instance runs on. Omitted if the
	// referenced flavor row has been retired.
	Flavor *Flavor `json:"flavor,omitempty"`

	// IAMRoleID IAM role the instance can assume via IMDS, if attached.
	IAMRoleID string `json:"iam_role_id,omitempty"`
	ID        string `json:"id,omitempty"`

	// Image resolved source image the instance booted from. Omitted for a
	// volume-only boot or if the referenced image row is gone.
	Image *Image `json:"image,omitempty"`

	// Keypairs SSH keypairs baked into the instance at launch. These are embedded
	// — unlike attached volumes, NICs, and security groups, which each
	// have their own list endpoint and so are not duplicated here.
	Keypairs   []*Keypair `json:"keypairs,omitempty"`
	LaunchedAt time.Time  `json:"launched_at,omitempty"`
	Metadata   Metadata   `json:"metadata,omitempty"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	Name string `json:"name,omitempty"`

	// PrimaryIP the primary NIC's IPv4 address, resolved at read time.
	PrimaryIP string `json:"primary_ip,omitempty"`

	// PrimaryIPv6 the primary NIC's public IPv6 address, when dual-stack.
	PrimaryIPv6 string `json:"primary_ipv6,omitempty"`

	// PublicIP the floating IP attached to the PRIMARY NIC, when any.
	//
	// It reports that one interface, so it is empty for an instance whose
	// public address sits on a secondary NIC — which is exactly what
	// `networks[].assign_public_ip` makes possible. An empty `public_ip`
	// is therefore not evidence that an instance has no public address,
	// and polling this field will never surface one.
	//
	// `GET /v1/instances/{instance_id}/nics` is the read that covers every
	// interface: each NIC carries its own `public_ip` and
	// `floating_ip_id`.
	PublicIP string `json:"public_ip,omitempty"`
	Tags     Tags   `json:"tags,omitempty"`

	// TaskState in-flight transition, if any; null when settled.
	TaskState    string    `json:"task_state,omitempty"`
	TerminatedAt time.Time `json:"terminated_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`

	// UserData base64-encoded cloud-init user-data supplied at launch.
	UserData string `json:"user_data,omitempty"`
}

type InstanceCreateRequest struct {
	// Architecture for bare image names; a CRN pins its architecture.
	Architecture *string `json:"architecture,omitempty"`
	Description  *string `json:"description,omitempty"`

	// Flavor regional flavor reference (UUID, CRN or exact name)
	//
	// Required.
	Flavor string `json:"flavor"`

	// IAMRole attach this organization-scoped IAM role by UUID, CRN or exact name.
	// The role's trust policy must permit `crn:compute:*:*:instance/*` (or
	// the specific instance CRN). The instance's IMDS endpoint
	// (169.254.169.254) mints short-lived STS credentials for this role
	// from inside the VM.
	IAMRole *string `json:"iam_role,omitempty"`

	// Image to clone the boot disk from (required if not booting from
	// volume). Four forms are accepted: a complete
	// image/name/architecture/arch/version/version CRN; an image id;
	// `name:version`, which pins one build and is how you opt out of the
	// tag moving under you; or a bare `name`, which follows the tag to
	// whichever build is current when the instance is created.
	Image    *string  `json:"image,omitempty"`
	Keypairs []string `json:"keypairs,omitempty"`
	Metadata Metadata `json:"metadata,omitempty"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	//
	// Required.
	Name string `json:"name"`

	// Networks interfaces to attach, at least one. index 0 is the primary NIC.
	//
	// Required because an instance with no interface boots with no network
	// at all, and nothing inside it can add one afterwards.
	//
	// Required.
	Networks []*NetworkConfig `json:"networks"`
	Tags     Tags             `json:"tags,omitempty"`

	// UserData base64-encoded user data (cloud-init)
	UserData []byte `json:"user_data,omitempty"`

	// Volumes disks created and bound with the instance, the boot disk included
	// — mark it with `boot: true`. At most one entry may.
	//
	// Omit the boot entry to take the image's minimum size and the
	// region's default tier.
	Volumes []*InstanceVolume `json:"volumes,omitempty"`
}

type InstanceFault struct {
	At time.Time `json:"at,omitempty"`

	// Code short machine-readable failure code.
	Code    string `json:"code,omitempty"`
	Details string `json:"details,omitempty"`
	Message string `json:"message,omitempty"`
}

// InstancePool a launch template plus a desired count. Creating a pool spawns
// desired_count instances; a reconciler converges member_count toward
// desired_count as it changes. member_count is how many members the pool
// holds; live_count is how many of them are running. A pool carries two
// tag sets and they answer different questions. `tags` labels the pool
// resource — that is what an IAM condition reads as
// `basalt:ResourceTag/<key>` and what a cost report groups by, and it
// reaches no instance. `template.tags` is the set stamped on every
// replica the pool launches.
type InstancePool struct {
	// CRN Cloud Resource Name. This is the value an IAM policy statement must
	// name to scope a permission to this pool alone; a policy written
	// against anything else will not match.
	CRN          string `json:"crn,omitempty"`
	Description  string `json:"description,omitempty"`
	DesiredCount int    `json:"desired_count,omitempty"`

	// ErrorMessage the last failure the reconciler recorded, cleared when the pool
	// reaches its target. Set alongside status `error`, and left in place
	// through a later resize — a pool that failed to spawn and is being
	// scaled again has not yet proved the failure is behind it.
	ErrorMessage string `json:"error_message,omitempty"`
	ID           string `json:"id,omitempty"`

	// LiveCount how many members are UP — bound instances whose current_state is
	// `running`.
	LiveCount int    `json:"live_count,omitempty"`
	ManagedBy string `json:"managed_by,omitempty"`
	MaxCount  int    `json:"max_count,omitempty"`

	// MemberCount how many instances the pool holds, running or not. This is what the
	// reconciler converges toward desired_count and what `status`
	// reflects, so member_count == desired_count with live_count below it
	// means the pool has the members it was asked for and some of them are
	// not up.
	MemberCount int `json:"member_count,omitempty"`
	MinCount    int `json:"min_count,omitempty"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	Name string `json:"name,omitempty"`

	// RefreshInProgress true while a rolling replacement requested through POST
	// /v1/instance-pools/{pool_id}/refresh is still running. It clears
	// itself once every member is on the current template. The pool reads
	// `scaling` for the duration, since it runs one instance over its
	// target while a replacement comes up.
	RefreshInProgress bool `json:"refresh_in_progress,omitempty"`

	// StaleInstanceCount how many members were launched from a template other than the pool's
	// current one — that is, how many a refresh would replace. Non-zero
	// after editing `template` and before refreshing, which is the signal
	// that a template change has not been rolled out yet.
	StaleInstanceCount int `json:"stale_instance_count,omitempty"`

	// Status where the pool is against its target.
	//
	// `active` means member_count == desired_count — the pool holds the
	// members it was asked for. It is not a claim that all of them are up;
	// read live_count for that.
	//
	// `scaling` means it does not, and the reconciler is converging it:
	// after a create, after a desired_count change, and for the length of
	// an instance refresh, which runs the pool one instance over its
	// target while a replacement comes up.
	//
	// `error` carries the last failure in error_message and is still
	// reconciled — the pool keeps being retried. `deleting` is a
	// teardown in flight.
	//
	// One of: "active", "scaling", "error", "deleting".
	Status string `json:"status,omitempty"`

	// Tags labels on the POOL itself, for IAM conditions
	// (`basalt:ResourceTag/<key>`) and cost attribution. They are attached
	// to nothing else: no instance the pool launches carries them. The
	// tags a replica is launched with are `template.tags`. Unlike the
	// other top-level fields beside this one, `tags` is not a projection
	// of the launch template — it is the pool's own set, and PATCHable
	// on its own.
	Tags Tags `json:"tags,omitempty"`

	// Template the pool's launch config, in the shape instance create takes. The
	// only place it appears: a flat copy of it beside this was two
	// spellings of one thing, and two spellings drift.
	Template *InstancePoolTemplate `json:"template,omitempty"`
}

// InstancePoolCreateRequest `template` is the launch config — the same fields instance create
// takes. The pool's own fields — name, description, tags, sizing —
// stay at the top level, because they describe the pool rather than the
// instances in it. A flavor and a primary subnet are required, as
// `template.flavor` + `template.networks[0].subnet`.
type InstancePoolCreateRequest struct {
	Description  *string `json:"description,omitempty"`
	DesiredCount *int    `json:"desired_count,omitempty"`
	MaxCount     *int    `json:"max_count,omitempty"`
	MinCount     *int    `json:"min_count,omitempty"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	//
	// Required.
	Name string `json:"name"`

	// Tags labels on the pool resource, for IAM conditions
	// (`basalt:RequestTag/<key>` here, `basalt:ResourceTag/<key>` on later
	// operations) and cost attribution. They are not propagated to the
	// instances the pool launches; `template.tags` is that set. A pool
	// field, sent beside `template`. Replica tags are only reachable
	// through `template.tags`.
	Tags Tags `json:"tags,omitempty"`

	// Required.
	Template *InstancePoolTemplateRequest `json:"template"`
}

type InstancePoolFloatingIPAttachRequest struct {
	// FloatingIP an account-scoped floating IP UUID or CRN (bare names are not
	// accepted), currently attached to nothing. This binds it to the pool;
	// it does not allocate one.
	//
	// Required.
	FloatingIP string `json:"floating_ip"`
}

// InstancePoolTemplate the pool's launch config, in the shape a standalone instance create
// takes: same field names, same types, same meanings, so a client that
// can build an instance can build a pool of them without a second,
// narrower contract to learn. It belongs to the pool. There is no
// separate launch-template resource to create, version or share between
// pools. Networking is one ordered `networks` list, index 0 being the
// primary NIC, replacing the flat `subnet_id` + `extra_nics` split.
// `ip_address` and `mac` are part of that shared NIC shape but are
// refused here: every replica launches from this one template, so a
// fixed address would have the second replica ask for one the first
// already holds.
type InstancePoolTemplate struct {
	FlavorID string `json:"flavor_id,omitempty"`

	// IAMRoleID IAM role attached to every replica, reachable from its IMDS
	// endpoint.
	IAMRoleID string `json:"iam_role_id,omitempty"`

	// ImageID image to clone each replica's boot disk from. Accepts the same four
	// forms instance create does: an architecture-qualified CRN, an image
	// id, `name:version`, or a bare `name`. Unlike instance create, the
	// reference is resolved ONCE, when the pool is created, and the
	// resulting image id is what every replica boots — including
	// replacements spawned months later. A tag re-resolved per replica
	// would let a heal boot a newer build than its siblings, and a pool
	// whose members are quietly not identical is the premise of the
	// primitive breaking silently. To move a pool to a new build, change
	// the template.
	ImageID  string   `json:"image_id,omitempty"`
	KeyNames []string `json:"key_names,omitempty"`
	Metadata Metadata `json:"metadata,omitempty"`

	// Networks per-replica interfaces. Index 0 is the primary NIC and is required;
	// the rest are extras.
	Networks []*NetworkConfigResponse `json:"networks,omitempty"`

	// Tags stamped on every instance this template launches. These are the
	// replicas' tags, not the pool's — the pool's own labels are the
	// top-level `tags`, and the two are independent. Changing them affects
	// FUTURE launches only. The instances already running keep the tags
	// they were launched with, so between the change and a refresh the
	// pool holds members carrying two different tag sets;
	// `stale_instance_count` is how many are still on the old one. POST
	// /v1/instance-pools/{pool_id}/refresh rolls them onto the current
	// template.
	Tags Tags `json:"tags,omitempty"`

	// UserData base64-encoded user data (cloud-init), stamped on every replica.
	UserData []byte `json:"user_data,omitempty"`

	// Volumes per-replica disks, the boot disk included — mark it with `boot:
	// true`. Same shape as instance create.
	Volumes []*InstanceVolume `json:"volumes,omitempty"`
}

// InstancePoolTemplateRequest the pool's launch config, in the shape a standalone instance create
// takes: same field names, same types, same meanings, so a client that
// can build an instance can build a pool of them without a second,
// narrower contract to learn. It belongs to the pool. There is no
// separate launch-template resource to create, version or share between
// pools. Networking is one ordered `networks` list, index 0 being the
// primary NIC, replacing the flat `subnet` + `extra_nics` split.
// `ip_address` and `mac` are part of that shared NIC shape but are
// refused here: every replica launches from this one template, so a
// fixed address would have the second replica ask for one the first
// already holds.
type InstancePoolTemplateRequest struct {
	Architecture *string `json:"architecture,omitempty"`

	// Required.
	Flavor string `json:"flavor"`

	// IAMRole IAM role attached to every replica, reachable from its IMDS
	// endpoint.
	IAMRole *string `json:"iam_role,omitempty"`

	// Image to clone each replica's boot disk from. Accepts the same four
	// forms instance create does: an architecture-qualified CRN, an image
	// id, `name:version`, or a bare `name`. Unlike instance create, the
	// reference is resolved ONCE, when the pool is created, and the
	// resulting image id is what every replica boots — including
	// replacements spawned months later. A tag re-resolved per replica
	// would let a heal boot a newer build than its siblings, and a pool
	// whose members are quietly not identical is the premise of the
	// primitive breaking silently. To move a pool to a new build, change
	// the template.
	Image    *string  `json:"image,omitempty"`
	Keypairs []string `json:"keypairs,omitempty"`
	Metadata Metadata `json:"metadata,omitempty"`

	// Networks per-replica interfaces. Index 0 is the primary NIC and is required;
	// the rest are extras.
	//
	// Required.
	Networks []*NetworkConfig `json:"networks"`

	// Tags stamped on every instance this template launches. These are the
	// replicas' tags, not the pool's — the pool's own labels are the
	// top-level `tags`, and the two are independent. Changing them affects
	// FUTURE launches only. The instances already running keep the tags
	// they were launched with, so between the change and a refresh the
	// pool holds members carrying two different tag sets;
	// `stale_instance_count` is how many are still on the old one. POST
	// /v1/instance-pools/{pool_id}/refresh rolls them onto the current
	// template.
	Tags Tags `json:"tags,omitempty"`

	// UserData base64-encoded user data (cloud-init), stamped on every replica.
	UserData []byte `json:"user_data,omitempty"`

	// Volumes per-replica disks, the boot disk included — mark it with `boot:
	// true`. Same shape as instance create.
	Volumes []*InstanceVolume `json:"volumes,omitempty"`
}

// InstancePoolUpdateRequest every field is optional; an omitted one is left alone, and sending
// none of them is a 400 rather than a silent no-op.
//
// `min_count` and `max_count` are fixed at create and are not patchable.
//
// The two tag sets move independently. `tags` relabels the pool itself
// and takes effect immediately, touching no instance. `template.tags`
// — like the rest of `template` — does NOT touch the instances
// already running: a live VM cannot change flavor, tier, subnet or its
// tags in place. It changes what the pool launches NEXT, so until you
// roll the pool it holds members carrying two different tag sets, and
// `stale_instance_count` is how many are on the older one. Bring them
// onto the current template with POST
// /v1/instance-pools/{pool_id}/refresh.
type InstancePoolUpdateRequest struct {
	// DesiredCount new target size, bounded by the pool's min_count/max_count and the
	// hard platform cap of 100.
	DesiredCount *int `json:"desired_count,omitempty"`

	// Tags REPLACES the pool's labels: the map you send becomes the whole set,
	// an empty object clears them, and omitting the field leaves them
	// alone. Replacement rather than a merge because a merge leaves no way
	// to say a key should be removed. These label the pool, not its
	// instances. To change what future replicas are tagged with, send
	// `template.tags`.
	Tags Tags `json:"tags,omitempty"`

	// Template replaces the launch config WHOLESALE — the object you send is what
	// the pool launches next, and anything you leave out is cleared rather
	// than kept. Replacement rather than a deep merge so a shorter
	// `networks` or `volumes` cannot be read as a truncation and silently
	// drop an interface or a disk.
	Template *InstancePoolTemplateRequest `json:"template,omitempty"`
}

type InstanceRebootRequest struct {
	// Hard force a power cycle (destroy + start, equivalent to a reset button)
	// instead of the default ACPI graceful reboot the guest can act on.
	Hard *bool `json:"hard,omitempty"`
}

type InstanceUpdateRequest struct {
	Description *string  `json:"description,omitempty"`
	Metadata    Metadata `json:"metadata,omitempty"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	Name *string `json:"name,omitempty"`
	Tags Tags    `json:"tags,omitempty"`
}

// InstanceVolume one disk created with the instance. `boot: true` marks the one cloned
// from image_id; every other entry is a blank volume the in-guest agent
// formats and mounts.
type InstanceVolume struct {
	// Boot marks the boot disk. It takes no mount_path or fstype — both come
	// from the image — and sending either is refused rather than
	// ignored.
	Boot bool `json:"boot,omitempty"`

	// DeleteOnTermination destroyed with the instance unless set false.
	DeleteOnTermination bool `json:"delete_on_termination,omitempty"`

	// Fstype filesystem the in-guest agent formats the volume with.
	Fstype    string `json:"fstype,omitempty"`
	MountPath string `json:"mount_path,omitempty"`
	SizeGB    int    `json:"size_gb"`

	// VolumeType tier; omitted = the region default.
	VolumeType string `json:"volume_type,omitempty"`
}

type Keypair struct {
	CreatedAt time.Time `json:"created_at,omitempty"`

	// CRN Cloud Resource Name
	CRN         string `json:"crn,omitempty"`
	Fingerprint string `json:"fingerprint,omitempty"`
	ID          string `json:"id,omitempty"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	Name string `json:"name,omitempty"`

	// PublicKey SSH public key
	PublicKey string `json:"public_key,omitempty"`
	Tags      Tags   `json:"tags,omitempty"`
}

type KeypairCreateRequest struct {
	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	//
	// Required.
	Name string `json:"name"`

	// PublicKey SSH public key (if not provided, a new keypair will be generated)
	PublicKey *string `json:"public_key,omitempty"`
	Tags      Tags    `json:"tags,omitempty"`
}

type ListInstanceNiCsNIC struct {
	BootIndex int `json:"boot_index,omitempty"`

	// External a customer-attached standalone interface — detach unbinds it
	// instead of destroying it.
	External bool `json:"external,omitempty"`

	// FloatingIPID the floating IP resource behind `public_ip`, so detaching or
	// releasing it needs no lookup.
	FloatingIPID string `json:"floating_ip_id,omitempty"`
	InterfaceID  string `json:"interface_id,omitempty"`
	IPAddress    string `json:"ip_address,omitempty"`
	IPv6Address  string `json:"ipv6_address,omitempty"`
	MAC          string `json:"mac,omitempty"`
	Name         string `json:"name,omitempty"`

	// Primary the lowest-boot-index NIC — the one carrying the guest's default
	// and metadata routes.
	Primary bool `json:"primary,omitempty"`

	// PublicIP the floating IP NATed to THIS interface, absent when it has none. A
	// floating IP fronting several NICs at once reports the same address
	// on each of them.
	PublicIP string `json:"public_ip,omitempty"`
	SubnetID string `json:"subnet_id,omitempty"`
	VPCID    string `json:"vpc_id,omitempty"`
}

type ListInstanceVolumesAttachment struct {
	BootIndex           int    `json:"boot_index,omitempty"`
	Bootable            bool   `json:"bootable,omitempty"`
	DeleteOnTermination bool   `json:"delete_on_termination,omitempty"`
	Device              string `json:"device,omitempty"`

	// Fstype filesystem the in-guest agent formatted the volume with, or absent
	// when the attachment did not name one and the agent used the ext4
	// default. Every write path — instance create, pool template and
	// attach — refuses anything else, so this is the whole set the field
	// can hold.
	//
	// One of: "ext4", "xfs".
	Fstype     string       `json:"fstype,omitempty"`
	Mount      *VolumeMount `json:"mount,omitempty"`
	MountPath  string       `json:"mount_path,omitempty"`
	Name       string       `json:"name,omitempty"`
	SizeGB     int          `json:"size_gb,omitempty"`
	Status     string       `json:"status,omitempty"`
	VolumeID   string       `json:"volume_id,omitempty"`
	VolumeType string       `json:"volume_type,omitempty"`
}

type Metadata = map[string]string

type NetworkConfig struct {
	// AssignPublicIP allocate a floating IP and attach it to THIS interface once it
	// exists. Per NIC, so a secondary interface can carry the public
	// address while the primary stays private, and an instance with
	// several public interfaces gets one address each.
	//
	// Each address is a separate floating-IP allocation: it counts against
	// the account's floating_ips quota and is billed like any other. It is
	// released when the instance is torn down — an address you allocated
	// yourself and attached to the same NIC is not, and survives the
	// instance.
	//
	// The interface's subnet must already route 0.0.0.0/0 to an internet
	// gateway. Without that the address would be silently unreachable, so
	// the launch fails instead.
	AssignPublicIP *bool `json:"assign_public_ip,omitempty"`

	// IPAddress optional fixed IP. Must be in the subnet's CIDR and not currently
	// allocated to another interface. An address is picked automatically
	// when omitted.
	IPAddress *string `json:"ip_address,omitempty"`

	// MAC Optional MAC address. Must be locally-administered (`X2:`, `X6:`,
	// `XA:`, `XE:` in the first octet). Generated when omitted.
	MAC *string `json:"mac,omitempty"`

	// SecurityGroups account-scoped security group references (UUID, CRN or name) to
	// attach to this NIC. Each must be owned by the same account. Empty
	// list = no per-NIC ACLs (the platform's default-allow stays in
	// force).
	SecurityGroups []string `json:"security_groups,omitempty"`

	// Subnet UUID or complete VPC/subnet CRN. Bare names require a VPC
	// parent and are rejected here.
	//
	// Required.
	Subnet string `json:"subnet"`
}

type NetworkConfigResponse struct {
	// AssignPublicIP allocate a floating IP and attach it to THIS interface once it
	// exists. Per NIC, so a secondary interface can carry the public
	// address while the primary stays private, and an instance with
	// several public interfaces gets one address each.
	//
	// Each address is a separate floating-IP allocation: it counts against
	// the account's floating_ips_v4 quota and is billed like any other. It
	// is released when the instance is torn down — an address you
	// allocated yourself and attached to the same NIC is not, and survives
	// the instance.
	//
	// The interface's subnet must already route 0.0.0.0/0 to an internet
	// gateway. Without that the address would be silently unreachable, so
	// the launch fails instead.
	AssignPublicIP bool `json:"assign_public_ip,omitempty"`

	// IPAddress optional fixed IP. Must be in the subnet's CIDR and not currently
	// allocated to another interface. An address is picked automatically
	// when omitted.
	IPAddress string `json:"ip_address,omitempty"`

	// MAC Optional MAC address. Must be locally-administered (`X2:`, `X6:`,
	// `XA:`, `XE:` in the first octet). Generated when omitted.
	MAC string `json:"mac,omitempty"`

	// SecurityGroupIDs account-scoped security group references (UUID, CRN or name) to
	// attach to this NIC. Each must be owned by the same account. Empty
	// list = no per-NIC ACLs (the platform's default-allow stays in
	// force).
	SecurityGroupIDs []string `json:"security_group_ids,omitempty"`

	// SubnetID Subnet UUID or complete VPC/subnet CRN. Bare names require a VPC
	// parent and are rejected here.
	SubnetID string `json:"subnet_id"`
}

type PoolInstance struct {
	CreatedAt  time.Time `json:"created_at,omitempty"`
	ID         string    `json:"id,omitempty"`
	InstanceID string    `json:"instance_id,omitempty"`
	PoolID     string    `json:"pool_id,omitempty"`

	// SequenceNum stable per-instance index within the pool.
	SequenceNum int `json:"sequence_num,omitempty"`
}

type ReinstallInstanceRequest struct {
	// Image replacement image reference (UUID, architecture-qualified CRN, name
	// or name:version). Omit to reinstall from the instance's current
	// image.
	Image *string `json:"image,omitempty"`

	// SizeGB replacement boot disk size; omitted = the image's min_disk_gb. Must
	// be within the volume size range (1..16384) and at least the image's
	// min_disk_gb.
	SizeGB *int `json:"size_gb,omitempty"`

	// VolumeType replacement boot disk tier; omitted = the region default.
	//
	// One of: "ssd", "nvme".
	VolumeType *string `json:"volume_type,omitempty"`
}

type ResizeInstanceRequest struct {
	// Flavor regional flavor reference (UUID, CRN or exact name) to resize to.
	//
	// Required.
	Flavor string `json:"flavor"`
}

// SerialConsoleTicket a one-shot credential for opening a serial console from a browser.
// Pass `ticket` as a query parameter on the WebSocket upgrade.
type SerialConsoleTicket struct {
	ExpiresAt time.Time `json:"expires_at"`

	// ExpiresIn seconds until it expires.
	ExpiresIn int `json:"expires_in"`

	// Ticket the credential. Opaque — do not parse it. Good for one instance
	// and one minute; mint a new one per connection rather than storing
	// it.
	Ticket string `json:"ticket"`
}

type Tags = map[string]string

type UpdateInstanceVolumeAttachmentRequest struct {
	// Required.
	DeleteOnTermination bool `json:"delete_on_termination"`
}

// VolumeMount what the in-guest agent reported about this attachment.
//
// Present only on an attachment that carries a `mount_path`. A volume
// without one is attached as a bare block device the tenant owns:
// nothing reconciles it, so nothing reports on it and a state for it
// would be a claim about something no one is watching.
//
// The agent is installed on first boot and nothing upgrades it
// afterwards, so an instance older than this feature reports nothing and
// every one of its volumes reads `unknown`. Treat `unknown` as "no
// information", never as healthy.
type VolumeMount struct {
	// Code why the volume is in this state. Absent when there is nothing to
	// say.
	//
	// Independent of `state` rather than something only a failure carries:
	// `fstab_write_failed` accompanies a **mounted** volume whose fstab
	// entry could not be written, which works now and will be gone after
	// the next reboot.
	//
	// The commonest one to act on is `signatures_no_filesystem` — the
	// disk carries a partition table or other signatures but no mountable
	// filesystem, so the agent will not format it, because formatting
	// would destroy what is there. A volume cloned from a boot disk and
	// attached with a `mount_path` lands here. Partition and format it
	// inside the guest, or attach it without a `mount_path` and mount it
	// yourself.
	//
	// `unknown_error` is a code this platform does not recognise, reported
	// by a guest agent newer than the region.
	//
	// One of: "unsafe_serial", "device_absent", "probe_failed", "signatures_no_filesystem", "unsupported_fstype", "mkfs_failed", "unsafe_mount_path", "mkdir_failed", "mount_failed", "fstab_write_failed", "unknown_error".
	Code string `json:"code,omitempty"`

	// Message human-readable detail from inside the guest — the failing
	// command's output, the partition table type it found. Free text
	// originating in the customer's own VM: sanitised and capped, but
	// display it as text, never as markup.
	Message string `json:"message,omitempty"`

	// ReportedAt when the guest agent last reported, whether or not anything had
	// changed. A `reported_at` far in the past means the agent has stopped
	// talking to us, and the state beside it is what it last said rather
	// than what is true now. Absent when `state` is `unknown`.
	ReportedAt time.Time `json:"reported_at,omitempty"`

	// Since when the volume entered this condition. Absent when `state` is
	// `unknown`.
	Since time.Time `json:"since,omitempty"`

	// State - `unknown` — no guest agent has ever reported on this volume.
	//   The agent may predate this feature, may have been removed
	//   (which is supported), or the guest may never have booted.
	// - `pending` — the agent cannot mount it yet and expects that to
	//   change. The ordinary state for the first seconds after an
	//   attach, while the hot-plugged disk appears in the guest.
	// - `mounted` — mounted at `mount_path`. May still carry a `code`.
	// - `failed` — it will not mount until something changes. Either
	//   the agent reported a refusal that waiting cannot fix, or it
	//   has been unable to make progress for long enough that waiting
	//   is no longer the explanation. `code` says which.
	//
	// One of: "unknown", "pending", "mounted", "failed".
	State string `json:"state"`
}
