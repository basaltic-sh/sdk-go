// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package database

import (
	"time"
)

// AddReplicaRequest add a read replica to the cluster. The body is empty today; future
// revisions carry replica options (read-only, sync, witness).
type AddReplicaRequest = map[string]any

type Backup struct {
	// BackendID the engine backend's own name for this backup — for postgres, the
	// pgbackrest label. It is what a restore passes to select THIS backup
	// rather than the stanza's most recent. Absent on failed runs and on
	// backups taken before the platform recorded labels.
	BackendID         string    `json:"backend_id,omitempty"`
	ClusterID         string    `json:"cluster_id"`
	EarliestRestoreAt time.Time `json:"earliest_restore_at,omitempty"`

	// EndLsn Postgres WAL LSN at backup end.
	EndLsn     string    `json:"end_lsn,omitempty"`
	Fault      *Fault    `json:"fault,omitempty"`
	FinishedAt time.Time `json:"finished_at,omitempty"`
	ID         string    `json:"id"`

	// Kind backup kind. base and incremental are what a customer requests; wal
	// and snapshot are catalogue entries the platform writes.
	//
	// One of: "base", "incremental", "wal", "snapshot".
	Kind            string    `json:"kind"`
	LatestRestoreAt time.Time `json:"latest_restore_at,omitempty"`

	// ObjectURI Object-storage URI of the backup artifact.
	ObjectURI string `json:"object_uri"`

	// Restorable whether this backup can be chosen by name. False means it can only
	// be restored while it is the cluster's most recent succeeded backup
	// — a backup picker should offer it accordingly.
	Restorable bool      `json:"restorable,omitempty"`
	SizeBytes  int64     `json:"size_bytes,omitempty"`
	StartedAt  time.Time `json:"started_at"`

	// Status backup status (e.g. running, succeeded, failed).
	Status string `json:"status"`
}

type Cluster struct {
	// AdminSecretID secrets service id holding the admin password.
	AdminSecretID string `json:"admin_secret_id,omitempty"`

	// AdminUser bootstrap admin/superuser role name.
	AdminUser string `json:"admin_user,omitempty"`

	// AssignPublicIP whether the cluster's endpoints are backed by floating IPs and
	// reachable from the internet. False means every endpoint resolves to
	// a member's in-VPC address and the cluster is reachable only from the
	// VPC.
	AssignPublicIP bool      `json:"assign_public_ip"`
	CreatedAt      time.Time `json:"created_at"`

	// CRN IAM resource CRN
	CRN             string      `json:"crn"`
	DefaultDatabase string      `json:"default_database,omitempty"`
	DeletedAt       time.Time   `json:"deleted_at,omitempty"`
	Description     string      `json:"description,omitempty"`
	Endpoints       []*Endpoint `json:"endpoints,omitempty"`

	// One of: "postgres", "valkey".
	EngineType string `json:"engine_type"`

	// EngineVersion e.g. '17' for postgres
	EngineVersion string `json:"engine_version"`
	Fault         *Fault `json:"fault,omitempty"`

	// FlavorID compute flavor each cluster member runs on. Must be a
	// database-family flavor.
	FlavorID string `json:"flavor_id"`
	ID       string `json:"id"`

	// InstanceCount node count. Postgres: 1 = single-node; 2+ adds Patroni HA members
	// (max 10). Valkey: 1 = single-node; 3 = primary + replicas with
	// Sentinel (2 is rejected; max 3). Does not guarantee a completed
	// failover drill for every topology.
	InstanceCount    int         `json:"instance_count"`
	Instances        []*Instance `json:"instances,omitempty"`
	Metadata         Metadata    `json:"metadata,omitempty"`
	Name             string      `json:"name"`
	ParameterGroupID string      `json:"parameter_group_id,omitempty"`

	// PatroniManaged whether postgres here runs under Patroni. Not the same question as
	// instance_count >= 2: a cluster converted from single-node is
	// Patroni-managed with one member — able to take a replica, not yet
	// HA. Adding a replica requires this, not the node count.
	PatroniManaged bool `json:"patroni_managed,omitempty"`

	// Status lifecycle status of the cluster. `converting` is a single-node
	// postgres cluster being adopted into Patroni — the node restarts
	// under a new supervisor, so replica adds, failovers and further
	// conversions are refused while it holds. `restoring` is a cluster
	// whose data is being overwritten in place: it is neither active nor
	// building, and must not be read as serving current data.
	//
	// One of: "pending", "building", "active", "modifying", "converting", "restoring", "failing-over", "deleting", "deleted", "error".
	Status    string    `json:"status"`
	StorageGB int       `json:"storage_gb"`
	Tags      Tags      `json:"tags"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateClusterRequest struct {
	// AdminUser optional; defaults to "admin". Postgres-safe identifier only.
	AdminUser *string `json:"admin_user,omitempty"`

	// AssignPublicIP endpoint exposure. True (the default) allocates a floating IP per
	// endpoint, and requires the cluster's subnet to carry a default route
	// (0.0.0.0/0) to an internet gateway — a public cluster on a subnet
	// without one is rejected before anything is provisioned. False
	// allocates no floating IP: endpoints resolve to member addresses
	// inside the VPC, which is how a cluster is placed on a private
	// subnet. IPv6 reachability is unaffected either way — it follows
	// the subnet's ::/0 route, as it does for every other resource.
	AssignPublicIP *bool `json:"assign_public_ip,omitempty"`

	// DefaultDatabase optional; defaults to "default". Postgres-safe identifier only.
	DefaultDatabase *string `json:"default_database,omitempty"`
	Description     *string `json:"description,omitempty"`

	// One of: "postgres", "valkey".
	//
	// Required.
	EngineType string `json:"engine_type"`

	// EngineVersion e.g. '17' for postgres. Defaults to the engine's current default
	// when omitted.
	EngineVersion *string `json:"engine_version,omitempty"`

	// FlavorID compute flavor each cluster member runs on. Must be a
	// database-family flavor.
	//
	// Required.
	FlavorID string `json:"flavor_id"`

	// InstanceCount node count. Postgres: 1 = single-node; 2+ adds Patroni HA members
	// (max 10). Valkey: 1 = single-node; 3 = primary + replicas with
	// Sentinel (2 is rejected; max 3).
	InstanceCount *int `json:"instance_count,omitempty"`

	// KeyNames platform-operator break-glass only. Stamps SSH keypairs onto the
	// cluster's managed VMs, which run the platform's own software; a
	// tenant reaches its database over the cluster endpoint, never over
	// SSH. Accepted only from the platform account and only with
	// `database:StampBreakGlassKeys` — any other account setting it is
	// rejected with 400 `INVALID_INPUT`.
	KeyNames []string `json:"key_names,omitempty"`
	Metadata Metadata `json:"metadata,omitempty"`

	// Required.
	Name string `json:"name"`

	// Networks at least one network interface is required.
	//
	// Required.
	Networks []*NICRequest `json:"networks"`

	// ParameterGroupID binds the cluster to a parameter group whose settings its members
	// boot with and converge to. Omitted means engine defaults. The group
	// must match the cluster's engine type and major version.
	ParameterGroupID *string      `json:"parameter_group_id,omitempty"`
	RestoreFrom      *RestoreFrom `json:"restore_from,omitempty"`

	// Required.
	StorageGB int  `json:"storage_gb"`
	Tags      Tags `json:"tags,omitempty"`
}

type CreateDBUserRequest struct {
	// Required.
	Name string `json:"name"`

	// Permissions engine-specific grant/permission map.
	Permissions map[string]any `json:"permissions,omitempty"`
}

type CreateDatabaseRequest struct {
	CharacterSet *string `json:"character_set,omitempty"`
	Collation    *string `json:"collation,omitempty"`

	// Required.
	Name string `json:"name"`
}

type CreateParameterGroupRequest struct {
	Description *string `json:"description,omitempty"`

	// One of: "postgres", "valkey".
	//
	// Required.
	EngineType string `json:"engine_type"`

	// EngineVersion defaults to the engine current default when omitted.
	EngineVersion *string `json:"engine_version,omitempty"`

	// Required.
	Name string `json:"name"`

	// Params engine settings this group applies. Validated against the engine's
	// tunable-parameter catalog — GET
	// /v1/engines/{engine_type}/parameters.
	Params map[string]string `json:"params,omitempty"`
	Tags   Tags              `json:"tags,omitempty"`
}

type DBUser struct {
	ClusterID string    `json:"cluster_id"`
	CreatedAt time.Time `json:"created_at"`
	ID        string    `json:"id"`
	IsAdmin   bool      `json:"is_admin"`
	Name      string    `json:"name"`

	// Password the generated password. Returned ONCE — only on create and
	// rotate-password responses. List/Get never include it; capture it
	// when you see it.
	Password string `json:"password,omitempty"`

	// PasswordSecretID secrets service id holding the user password.
	PasswordSecretID string `json:"password_secret_id"`

	// Permissions engine-specific grant/permission map.
	Permissions map[string]any `json:"permissions"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

type Database struct {
	CharacterSet string    `json:"character_set,omitempty"`
	ClusterID    string    `json:"cluster_id"`
	Collation    string    `json:"collation,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	ID           string    `json:"id"`
	Name         string    `json:"name"`
}

// Endpoint one writer/reader/admin connection endpoint for a cluster.
type Endpoint struct {
	DNSName string `json:"dns_name"`

	// FloatingIPID the endpoint's public floating IP. Absent on clusters created with
	// assign_public_ip=false, whose endpoints resolve to a member's in-VPC
	// address instead. Connect by dns_name in both cases.
	FloatingIPID string `json:"floating_ip_id,omitempty"`
	ID           string `json:"id"`

	// IPAddress the address `dns_name` currently answers with — the floating IP on
	// a public endpoint, or the serving member's in-VPC address on a
	// private one. Absent while it cannot be resolved, e.g. an endpoint
	// whose member is still provisioning. It is resolved per request and a
	// private endpoint's changes on failover, so connect by `dns_name` and
	// use this for the things a name cannot express, such as a
	// security-group rule.
	IPAddress string `json:"ip_address,omitempty"`

	// Kind which role this endpoint targets.
	//
	// One of: "writer", "reader", "admin".
	Kind string `json:"kind"`
	Port int    `json:"port"`
}

// Engine one supported engine and the major versions a tenant may select.
type Engine struct {
	// DefaultVersion major applied when a create request omits engine_version.
	DefaultVersion string   `json:"default_version"`
	EngineValue    string   `json:"engine"`
	Versions       []string `json:"versions"`
}

// EngineParameter one setting a parameter group for this engine may carry.
type EngineParameter struct {
	// Apply what it costs to put the setting into effect. reload lands without
	// dropping connections; restart takes effect only when the engine next
	// restarts — the platform does not bounce a cluster for a config
	// edit.
	//
	// One of: "reload", "restart".
	Apply string `json:"apply"`

	// Kind the value grammar this parameter accepts.
	//
	// One of: "integer", "memory", "duration", "float", "boolean", "enum", "pattern".
	Kind string `json:"kind"`

	// Max upper bound, in the parameter base unit.
	Max float64 `json:"max,omitempty"`

	// Min lower bound, in the parameter base unit.
	Min  float64 `json:"min,omitempty"`
	Name string  `json:"name"`

	// Pattern regexp the value must match, for pattern parameters.
	Pattern string `json:"pattern,omitempty"`

	// Values the legal settings, for enum and boolean parameters.
	Values []string `json:"values,omitempty"`
}

type EngineParameterListResponse struct {
	Engine     string             `json:"engine"`
	Parameters []*EngineParameter `json:"parameters"`
}

// FailoverRequest trigger a planned HA switchover. Postgres uses Patroni; Valkey uses
// FAILOVER TO (when target_member is set) or SENTINEL FAILOVER (when
// empty). Leave target_member empty to let the engine pick a candidate.
type FailoverRequest struct {
	// TargetMember optional member name to promote.
	TargetMember *string `json:"target_member,omitempty"`
}

// Fault error detail attached to a resource in status=error.
type Fault struct {
	At      time.Time `json:"at"`
	Code    string    `json:"code"`
	Details string    `json:"details,omitempty"`
	Message string    `json:"message"`
}

// Instance One VM-shaped member of a cluster.
type Instance struct {
	// AgentVersion version of the in-VM management agent this member reported in its
	// last heartbeat. Absent when the member has not reported one — it
	// has not beat yet, or it runs an agent released before the field
	// existed. Absent means unknown: never assume it matches another
	// member's.
	AgentVersion      string `json:"agent_version,omitempty"`
	BootVolumeID      string `json:"boot_volume_id,omitempty"`
	ComputeInstanceID string `json:"compute_instance_id,omitempty"`
	ID                string `json:"id"`

	// MemberName cluster member name — pass as failover target_member to promote
	// this node.
	MemberName string `json:"member_name,omitempty"`

	// Role member role within the cluster.
	//
	// One of: "primary", "replica".
	Role   string `json:"role"`
	Status string `json:"status"`
}

type Metadata = map[string]string

// NICRequest a network interface attachment for the cluster's members.
type NICRequest struct {
	// IPAddress optional fixed IP; auto-allocated from the subnet when omitted.
	IPAddress *string `json:"ip_address,omitempty"`

	// MAC optional fixed MAC; auto-generated when omitted.
	MAC              *string  `json:"mac,omitempty"`
	SecurityGroupIDs []string `json:"security_group_ids,omitempty"`

	// Required.
	SubnetID string `json:"subnet_id"`
}

type ParameterGroup struct {
	AccountID   string    `json:"account_id"`
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description,omitempty"`

	// One of: "postgres", "valkey".
	EngineType    string `json:"engine_type"`
	EngineVersion string `json:"engine_version"`
	ID            string `json:"id"`
	Name          string `json:"name"`

	// Params engine settings this group applies. Every name must be tunable for
	// the group's engine and version — see GET
	// /v1/engines/{engine_type}/parameters for the list, each entry's
	// value grammar and bounds, and whether applying it needs a restart.
	// Settings the platform owns (postgres archive_command or wal_level,
	// valkey bind or dir) are not tunable and are rejected.
	Params map[string]string `json:"params"`

	// SystemManaged true for the platform-provided default groups (immutable).
	SystemManaged bool      `json:"system_managed"`
	Tags          Tags      `json:"tags"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// RequestBackupRequest trigger a manual backup. Empty kind defaults to "base".
type RequestBackupRequest struct {
	// Kind backup kind. Defaults to base — a full snapshot, which needs no
	// earlier backup to build on. incremental records only what changed
	// since the last one. pgbackrest's own names for these two, full and
	// incr, are still accepted here and on the list filter, but the backup
	// is stored and returned as base or incremental.
	//
	// One of: "base", "incremental", "full", "incr".
	Kind *string `json:"kind,omitempty"`
}

// RestoreClusterRequest restore a backup INTO THIS CLUSTER, overwriting its data. Destructive
// and irreversible.
//
// A pre-restore backup is taken automatically before anything is
// overwritten, so there is a way back from a mistaken restore.
type RestoreClusterRequest struct {
	// BackupID the backup to restore. May be a backup of THIS cluster (the common
	// "undo a bad migration" case) or of another cluster in the same
	// account, in which case this cluster is granted read on that backup
	// bucket for the duration and it is revoked afterwards. Must satisfy
	// the same rules as a create-time restore — see
	// restore_from.backup_id.
	//
	// Required.
	BackupID string `json:"backup_id"`

	// Confirm must equal the cluster's name. Restoring overwrites live data and
	// cannot be undone, so consent has to be something that cannot arrive
	// by accident — a retried request, a stale tab, a script looping
	// over ids.
	//
	// Required.
	Confirm string `json:"confirm"`

	// RecoveryTargetTime RFC3339 timestamp for point-in-time recovery — WAL is replayed
	// forward from the backup until this time, then stops.
	RecoveryTargetTime *time.Time `json:"recovery_target_time,omitempty"`
}

// RestoreFrom bootstrap the new cluster from an existing backup instead of a fresh
// initdb. The source backup must belong to the same account and be in
// status=succeeded. Single-node only in the current pass.
type RestoreFrom struct {
	// BackupID any succeeded backup of the source cluster whose `restorable` is
	// true. A backup with `restorable: false` predates the platform
	// recording pgbackrest labels, so it can only be restored while it is
	// the source cluster's most recent succeeded backup, and is refused
	// otherwise — pgbackrest would silently restore the newest instead.
	//
	// Required.
	BackupID string `json:"backup_id"`

	// RecoveryTargetTime RFC3339 timestamp for point-in-time recovery — pgbackrest replays
	// WAL forward from the base backup until this time, then stops. Empty
	// = restore to the latest available WAL.
	RecoveryTargetTime *time.Time `json:"recovery_target_time,omitempty"`
}

type Tags = map[string]string

// UpdateClusterRequest description, tags and the parameter-group binding are editable. The
// topology (engine, flavor, storage, instance count) is immutable
// through this path. Omitted fields are left unchanged.
type UpdateClusterRequest struct {
	Description *string `json:"description,omitempty"`

	// ParameterGroupID re-binds the cluster to a different parameter group; its members
	// converge on the new settings. An empty string clears the binding and
	// returns the cluster to engine defaults. Omitted leaves it unchanged.
	ParameterGroupID *string `json:"parameter_group_id,omitempty"`
	Tags             Tags    `json:"tags,omitempty"`
}

// UpdateParameterGroupRequest omitted fields are left unchanged. Editing params REPLACES the whole
// set and is republished to every cluster bound to this group, whose
// members converge on it — a setting dropped from the group is
// reverted on them.
type UpdateParameterGroupRequest struct {
	Description *string           `json:"description,omitempty"`
	Params      map[string]string `json:"params,omitempty"`
	Tags        Tags              `json:"tags,omitempty"`
}
