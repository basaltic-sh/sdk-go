// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package storage

import (
	"time"
)

type Bucket struct {
	ACL       string    `json:"acl"`
	CreatedAt time.Time `json:"created_at"`
	CRN       string    `json:"crn"`

	// DeletionProtection when true, DeleteBucket schedules deletion instead of removing the
	// bucket immediately.
	DeletionProtection bool `json:"deletion_protection"`

	// DeletionRecoveryDays recovery window, in days, applied when a deletion is scheduled. 0
	// means the service default is used.
	DeletionRecoveryDays int    `json:"deletion_recovery_days"`
	ID                   string `json:"id"`
	Name                 string `json:"name"`

	// ScheduledDeletionAt deadline a protected bucket is deleted at. Present only while a
	// deletion is pending — POST /v1/buckets/{bucket}/restore before it
	// passes to cancel.
	ScheduledDeletionAt time.Time `json:"scheduled_deletion_at,omitempty"`

	// Versioning state. `suspended` means versioning was on and was turned
	// off — existing versions are kept, new writes stop creating them
	// — which is distinct from `disabled`, a bucket that never had it
	// enabled. The S3-compatible endpoint spells the same states `Enabled`
	// / `Suspended` in its XML, and reports `disabled` by omitting the
	// element.
	//
	// One of: "disabled", "enabled", "suspended".
	Versioning string `json:"versioning"`
}

// BucketPolicy pkg/policy.PolicyDocument serialization. See the IAM policy docs for
// the statement shape; here we just declare it as an opaque object so
// the spec doesn't have to track schema changes inside the policy
// engine.
type BucketPolicy = map[string]any

type CORSConfig struct {
	Rules []*CORSRule `json:"rules"`
}

type CORSRule struct {
	AllowedHeaders []string `json:"allowed_headers,omitempty"`
	AllowedMethods []string `json:"allowed_methods"`
	AllowedOrigins []string `json:"allowed_origins"`
	ExposeHeaders  []string `json:"expose_headers,omitempty"`
	ID             string   `json:"id,omitempty"`
	MaxAgeSeconds  int      `json:"max_age_seconds,omitempty"`
}

type CompleteMultipartUploadRequest struct {
	// Parts every part the assembled object is made of, in ascending part_number
	// order. Each etag must match the one that part's upload returned.
	//
	// Required.
	Parts []*CompleteMultipartUploadRequestPart `json:"parts"`
}

type CompleteMultipartUploadRequestPart struct {
	// Required.
	Etag string `json:"etag"`

	// Required.
	PartNumber int `json:"part_number"`
}

type CompleteMultipartUploadResponse struct {
	Etag         string `json:"etag"`
	Size         int64  `json:"size"`
	StorageClass string `json:"storage_class,omitempty"`
	VersionID    string `json:"version_id,omitempty"`
}

// CreatableVolumeTypeName the tiers a new volume may be provisioned on. `hdd` is a valid stored
// tier but its pool is reserved for cold storage, so creating a block
// volume on it is rejected.
type CreatableVolumeTypeName string

// Values CreatableVolumeTypeName accepts.
const (
	CreatableVolumeTypeNameSSD  CreatableVolumeTypeName = "ssd"
	CreatableVolumeTypeNameNVMe CreatableVolumeTypeName = "nvme"
)

type CreateBucketRequest struct {
	// Required.
	Name string `json:"name"`

	// ObjectLockEnabled when true, enables S3 Object Lock on the bucket at creation time and
	// turns versioning on. Object Lock cannot be enabled later.
	ObjectLockEnabled *bool `json:"object_lock_enabled,omitempty"`
}

type EncryptionConfig struct {
	Rules []*EncryptionRule `json:"rules"`
}

type EncryptionRule struct {
	BucketKeyEnabled bool                   `json:"bucket_key_enabled,omitempty"`
	Default          *EncryptionRuleDefault `json:"default,omitempty"`
}

type EncryptionRuleDefault struct {
	KMSMasterKeyID string `json:"kms_master_key_id,omitempty"`
	SseAlgorithm   string `json:"sse_algorithm"`
}

type InitiateMultipartUploadRequest struct {
	ContentType *string `json:"content_type,omitempty"`

	// Required.
	Key          string            `json:"key"`
	Metadata     map[string]string `json:"metadata,omitempty"`
	StorageClass *string           `json:"storage_class,omitempty"`
}

type LifecycleConfig struct {
	Rules []*LifecycleRule `json:"rules"`
}

type LifecycleRule struct {
	AbortIncompleteMultipartUpload *LifecycleRuleAbortIncompleteMultipartUpload `json:"abort_incomplete_multipart_upload,omitempty"`
	Expiration                     *LifecycleRuleExpiration                     `json:"expiration,omitempty"`
	Filter                         *LifecycleRuleFilter                         `json:"filter,omitempty"`
	ID                             string                                       `json:"id,omitempty"`
	NoncurrentVersionExpiration    *LifecycleRuleNoncurrentVersionExpiration    `json:"noncurrent_version_expiration,omitempty"`

	// Status `disabled` keeps the rule in the configuration but skips it during
	// evaluation. The S3-compatible endpoint spells the same states
	// `Enabled` / `Disabled` in its XML.
	//
	// One of: "enabled", "disabled".
	Status string `json:"status"`

	// Transition move matching objects to another storage class once they are old
	// enough. The object keeps its identity — same key, same version id,
	// same last-modified — and only its bytes move between pools, so
	// pairing a transition with an expiration works: the expiry clock is
	// not restarted by the move. Exactly one of `days` or `date`. When the
	// rule also has an `expiration`, the transition must come strictly
	// first, otherwise the object would be deleted before it ever moved
	// and the rule is rejected. Only one transition per rule: the platform
	// serves two classes, so a second has nowhere to go. The S3-compatible
	// endpoint accepts a single-element `<Transition>` list and rejects
	// longer ones rather than silently applying the first.
	Transition *LifecycleRuleTransition `json:"transition,omitempty"`
}

type LifecycleRuleAbortIncompleteMultipartUpload struct {
	DaysAfterInitiation int `json:"days_after_initiation,omitempty"`
}

type LifecycleRuleExpiration struct {
	Date time.Time `json:"date,omitempty"`
	Days int       `json:"days,omitempty"`
}

type LifecycleRuleFilter struct {
	Prefix string `json:"prefix,omitempty"`
}

type LifecycleRuleNoncurrentVersionExpiration struct {
	NewerNoncurrentVersions int `json:"newer_noncurrent_versions,omitempty"`
	NoncurrentDays          int `json:"noncurrent_days,omitempty"`
}

// LifecycleRuleTransition move matching objects to another storage class once they are old
// enough. The object keeps its identity — same key, same version id,
// same last-modified — and only its bytes move between pools, so
// pairing a transition with an expiration works: the expiry clock is not
// restarted by the move. Exactly one of `days` or `date`. When the rule
// also has an `expiration`, the transition must come strictly first,
// otherwise the object would be deleted before it ever moved and the
// rule is rejected. Only one transition per rule: the platform serves
// two classes, so a second has nowhere to go. The S3-compatible endpoint
// accepts a single-element `<Transition>` list and rejects longer ones
// rather than silently applying the first.
type LifecycleRuleTransition struct {
	// Date absolute cut-off; fires on the next sweep after this instant.
	Date time.Time `json:"date,omitempty"`

	// Days after the object's last-modified time.
	Days int `json:"days,omitempty"`

	// StorageClass destination class. Keeps the S3 standard's casing rather than the
	// platform's lowercase convention, because the class vocabulary is
	// S3's. Unsupported values are rejected — a transition to a class
	// that does not exist would silently never run.
	//
	// One of: "STANDARD", "COLD".
	StorageClass string `json:"storage_class"`
}

type ListObjectVersionsResponse struct {
	IsTruncated         bool             `json:"is_truncated"`
	NextKeyMarker       string           `json:"next_key_marker,omitempty"`
	NextVersionIDMarker string           `json:"next_version_id_marker,omitempty"`
	Versions            []*ObjectVersion `json:"versions"`
}

type ListPartsResponse struct {
	Parts  []*MultipartPart `json:"parts"`
	Upload *MultipartUpload `json:"upload"`
}

type MultipartPart struct {
	Etag       string `json:"etag"`
	PartNumber int    `json:"part_number"`
	Size       int64  `json:"size"`
}

type MultipartUpload struct {
	Bucket       string            `json:"bucket"`
	ContentType  string            `json:"content_type,omitempty"`
	CreatedAt    time.Time         `json:"created_at"`
	Key          string            `json:"key"`
	Metadata     map[string]string `json:"metadata,omitempty"`
	StorageClass string            `json:"storage_class,omitempty"`
	UploadID     string            `json:"upload_id"`
}

type ObjectEntry struct {
	ContentType  string    `json:"content_type,omitempty"`
	Etag         string    `json:"etag"`
	Key          string    `json:"key"`
	LastModified time.Time `json:"last_modified"`
	Size         int64     `json:"size"`
}

type ObjectListResponse struct {
	CommonPrefixes []string        `json:"common_prefixes,omitempty"`
	IsTruncated    bool            `json:"is_truncated,omitempty"`
	Meta           *PaginationMeta `json:"meta,omitempty"`
	Objects        []*ObjectEntry  `json:"objects,omitempty"`
}

type ObjectLockConfig struct {
	ObjectLockEnabled string                `json:"object_lock_enabled,omitempty"`
	Rule              *ObjectLockConfigRule `json:"rule,omitempty"`
}

type ObjectLockConfigRule struct {
	DefaultRetention *ObjectLockConfigRuleDefaultRetention `json:"default_retention,omitempty"`
}

type ObjectLockConfigRuleDefaultRetention struct {
	Days int `json:"days,omitempty"`

	// One of: "GOVERNANCE", "COMPLIANCE".
	Mode  string `json:"mode,omitempty"`
	Years int    `json:"years,omitempty"`
}

type ObjectVersion struct {
	ContentType    string    `json:"content_type,omitempty"`
	Etag           string    `json:"etag"`
	IsDeleteMarker bool      `json:"is_delete_marker"`
	IsLatest       bool      `json:"is_latest"`
	Key            string    `json:"key"`
	LastModified   time.Time `json:"last_modified"`
	Size           int64     `json:"size"`
	StorageClass   string    `json:"storage_class,omitempty"`
	VersionID      string    `json:"version_id"`
}

type PaginationMeta struct {
	// HasMore whether there are more items
	HasMore bool `json:"has_more,omitempty"`

	// Limit number of items per page
	Limit int `json:"limit,omitempty"`

	// Marker opaque cursor for the next page. Pass it back as the `marker` query
	// parameter; treat it as a token, not a value to parse.
	Marker string `json:"marker,omitempty"`

	// Total number of items
	Total int `json:"total,omitempty"`
}

type PutBucketCORSRequest struct {
	// Required.
	CORS *CORSConfig `json:"cors"`
}

type PutBucketDeletionProtectionRequest struct {
	// Enabled when true, DeleteBucket schedules deletion instead of removing
	// immediately.
	//
	// Required.
	Enabled bool `json:"enabled"`

	// RecoveryDays recovery window when enabling protection. 0 uses the service
	// default; values are clamped to [1,30].
	RecoveryDays *int `json:"recovery_days,omitempty"`
}

type PutBucketEncryptionRequest struct {
	// Required.
	Encryption *EncryptionConfig `json:"encryption"`
}

type PutBucketLifecycleRequest struct {
	// Required.
	Lifecycle *LifecycleConfig `json:"lifecycle"`
}

type PutBucketObjectLockRequest struct {
	// Required.
	ObjectLock *ObjectLockConfig `json:"object_lock"`
}

type PutBucketPolicyRequest struct {
	// Required.
	Document BucketPolicy `json:"document"`
}

type PutBucketTaggingRequest struct {
	// Required.
	Tags TagSet `json:"tags"`
}

type PutBucketVersioningRequest struct {
	// One of: "enabled", "suspended".
	//
	// Required.
	Status string `json:"status"`
}

type PutObjectResponse struct {
	Etag string `json:"etag"`
	Key  string `json:"key"`
	Size int64  `json:"size"`
}

type Snapshot struct {
	CreatedAt time.Time `json:"created_at,omitempty"`

	// CRN Cloud Resource Name (name-based, region-scoped).
	CRN         string `json:"crn,omitempty"`
	Description string `json:"description,omitempty"`

	// ErrorMessage last failure reason. Empty unless status=error.
	ErrorMessage string `json:"error_message,omitempty"`
	ID           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`

	// SizeGB frozen size of the source volume at the time the snapshot was taken
	// — the volume may have been extended since.
	SizeGB int `json:"size_gb,omitempty"`

	// SnapshotPolicyID the snapshot policy that took this snapshot. Absent when a person
	// did. This is also what retention matches on, so its presence is what
	// makes a snapshot eligible for automatic deletion — snapshots taken
	// by hand are never reaped.
	SnapshotPolicyID string         `json:"snapshot_policy_id,omitempty"`
	Status           SnapshotStatus `json:"status,omitempty"`
	Tags             Tags           `json:"tags,omitempty"`
	UpdatedAt        time.Time      `json:"updated_at,omitempty"`

	// VolumeID source volume the snapshot was taken from.
	VolumeID string `json:"volume_id,omitempty"`
}

type SnapshotCreateRequest struct {
	Description *string `json:"description,omitempty"`

	// Required.
	Name string `json:"name"`
	Tags Tags   `json:"tags,omitempty"`

	// Required.
	VolumeID string `json:"volume_id"`
}

// SnapshotIntervalMinutes minutes between snapshots — a minimum gap, not an exact cadence. A
// periodic pass takes whatever has come due and re-bases each policy's
// next run off the moment it ran, so a snapshot lands at or after
// `interval_minutes` and never before, and can land a minute or two
// later when the pass is busy. A window the pass misses costs one
// snapshot rather than producing a catch-up burst afterwards.
//
// The floor is one minute, because that pass is what evaluates the
// schedule and nothing finer can be honoured; the ceiling is 30 days.
// Sub-hourly intervals multiply Ceph snapshot churn and count against
// the `snapshots` quota, so pick the largest interval that meets your
// recovery point objective.
type SnapshotIntervalMinutes = int

// SnapshotPolicy a schedule attached to one volume: take a snapshot every
// `interval_minutes`, then keep at most `retention_count` of the
// snapshots this policy created.
//
// Retention only ever deletes snapshots the policy itself created (those
// carrying its `snapshot_policy_id`) — a snapshot taken by hand is
// never reaped. It also never deletes a snapshot something depends on: a
// snapshot a volume was created from, including a restore still in
// progress, is skipped and re-examined later.
type SnapshotPolicy struct {
	CreatedAt time.Time `json:"created_at,omitempty"`

	// CRN Cloud Resource Name (name-based, region-scoped).
	CRN         string `json:"crn,omitempty"`
	Description string `json:"description,omitempty"`

	// Enabled disabling pauses the whole policy — no scheduled snapshots and no
	// retention. A paused schedule that kept reaping would delete history
	// while you were looking at it.
	Enabled         bool                    `json:"enabled,omitempty"`
	ID              string                  `json:"id,omitempty"`
	IntervalMinutes SnapshotIntervalMinutes `json:"interval_minutes,omitempty"`

	// LastError why the most recent fire produced no snapshot (quota exhausted,
	// volume mid-extend, …). Absent when the last fire succeeded.
	LastError string `json:"last_error,omitempty"`

	// LastRunAt when the policy last fired. Absent until the first fire.
	LastRunAt time.Time `json:"last_run_at,omitempty"`
	Name      string    `json:"name,omitempty"`

	// NextRunAt when the next snapshot is due. Re-stamped to `now +
	// interval_minutes` each time the policy fires — never to `previous
	// + interval` — so a window missed while the region was busy costs
	// one snapshot, not one per window missed.
	NextRunAt      time.Time              `json:"next_run_at,omitempty"`
	RetentionCount SnapshotRetentionCount `json:"retention_count,omitempty"`
	RetentionDays  SnapshotRetentionDays  `json:"retention_days,omitempty"`
	Tags           Tags                   `json:"tags,omitempty"`
	UpdatedAt      time.Time              `json:"updated_at,omitempty"`

	// VolumeID the volume this schedule is attached to. A volume has at most one
	// policy.
	VolumeID string `json:"volume_id,omitempty"`
}

type SnapshotPolicyCreateRequest struct {
	Description *string `json:"description,omitempty"`

	// Enabled defaults to true. Set false to attach a paused schedule. Pausing
	// stops the whole policy — no snapshots are taken and none are
	// deleted, because a paused schedule that kept reaping would delete
	// history while you were looking at it.
	Enabled *bool `json:"enabled,omitempty"`

	// Required.
	IntervalMinutes SnapshotIntervalMinutes `json:"interval_minutes"`

	// Name unique within the account — it names the policy in its CRN.
	// Scheduled snapshots are named `<policy>-<UTC timestamp>`.
	//
	// Required.
	Name string `json:"name"`

	// Required.
	RetentionCount SnapshotRetentionCount `json:"retention_count"`
	RetentionDays  *SnapshotRetentionDays `json:"retention_days,omitempty"`
	Tags           Tags                   `json:"tags,omitempty"`

	// Required.
	VolumeID string `json:"volume_id"`
}

// SnapshotPolicyUpdateRequest PATCH semantics — an omitted field keeps its current value. Changing
// `interval_minutes` re-bases the next run off now, so shortening a
// daily schedule to hourly takes effect within the hour.
type SnapshotPolicyUpdateRequest struct {
	Description *string `json:"description,omitempty"`

	// Enabled false pauses the policy, true resumes it. Pausing stops the whole
	// policy — no snapshots are taken and none are deleted, so a paused
	// schedule cannot lose you history. Resuming applies the retention
	// window again on the next run, so anything sitting outside it by then
	// — because you lowered `retention_count` while paused, say — is
	// reaped on that run.
	Enabled         *bool                    `json:"enabled,omitempty"`
	IntervalMinutes *SnapshotIntervalMinutes `json:"interval_minutes,omitempty"`
	Name            *string                  `json:"name,omitempty"`
	RetentionCount  *SnapshotRetentionCount  `json:"retention_count,omitempty"`
	RetentionDays   *SnapshotRetentionDays   `json:"retention_days,omitempty"`
	Tags            Tags                     `json:"tags,omitempty"`
}

// SnapshotRetentionCount how many of this policy's snapshots to keep. When a fire takes the
// count past this, the oldest go first.
type SnapshotRetentionCount = int

// SnapshotRetentionDays optional age bound, applied on top of `retention_count`: a snapshot
// outside EITHER window is reaped. 0 means no age bound. The single
// newest snapshot is exempt from the age bound, so a volume that could
// not be snapshotted for longer than the window never loses its whole
// history.
type SnapshotRetentionDays = int

type SnapshotStatus string

// Values SnapshotStatus accepts.
const (
	SnapshotStatusCreating  SnapshotStatus = "creating"
	SnapshotStatusAvailable SnapshotStatus = "available"
	SnapshotStatusDeleting  SnapshotStatus = "deleting"
	SnapshotStatusError     SnapshotStatus = "error"
)

type SnapshotUpdateRequest struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
	Tags        Tags    `json:"tags,omitempty"`
}

type TagSet = map[string]string

type Tags = map[string]string

type UploadPartResponse struct {
	Etag       string `json:"etag"`
	PartNumber int    `json:"part_number"`
	Size       int64  `json:"size"`
}

type Volume struct {
	// Bootable set at create time when the volume is provisioned from a bootable
	// image. Immutable after creation.
	Bootable  bool      `json:"bootable,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`

	// CRN Cloud Resource Name (name-based, region-scoped).
	CRN         string `json:"crn,omitempty"`
	Description string `json:"description,omitempty"`

	// ErrorMessage last failure reason. Empty unless status=error.
	ErrorMessage string `json:"error_message,omitempty"`
	ID           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	SizeGB       int    `json:"size_gb,omitempty"`

	// SourceImageID image the volume was cloned from, when applicable.
	SourceImageID string `json:"source_image_id,omitempty"`

	// SourceSnapshotID snapshot the volume was cloned from (restore path), when applicable.
	SourceSnapshotID string         `json:"source_snapshot_id,omitempty"`
	Status           VolumeStatus   `json:"status,omitempty"`
	Tags             Tags           `json:"tags,omitempty"`
	UpdatedAt        time.Time      `json:"updated_at,omitempty"`
	VolumeType       VolumeTypeName `json:"volume_type,omitempty"`
}

type VolumeCreateRequest struct {
	Bootable    *bool   `json:"bootable,omitempty"`
	Description *string `json:"description,omitempty"`

	// Required.
	Name string `json:"name"`

	// Required.
	SizeGB int `json:"size_gb"`

	// SourceImageID reserved: when supplied, the volume will be provisioned as a clone
	// of the referenced image's base snapshot. Currently recorded on the
	// row but not yet wired to the clone path.
	SourceImageID *string `json:"source_image_id,omitempty"`

	// SourceSnapshotID clone the new volume from an existing snapshot (restore). Mutually
	// exclusive with source_image_id; size_gb must be at least the
	// snapshot's frozen size.
	SourceSnapshotID *string `json:"source_snapshot_id,omitempty"`
	Tags             Tags    `json:"tags,omitempty"`

	// Required.
	VolumeType CreatableVolumeTypeName `json:"volume_type"`
}

type VolumeExtendRequest struct {
	// NewSizeGB new size in GB. Must be strictly greater than the current size.
	//
	// Required.
	NewSizeGB int `json:"new_size_gb"`
}

type VolumeStatus string

// Values VolumeStatus accepts.
const (
	VolumeStatusCreating  VolumeStatus = "creating"
	VolumeStatusAvailable VolumeStatus = "available"
	VolumeStatusInUse     VolumeStatus = "in_use"
	VolumeStatusExtending VolumeStatus = "extending"
	VolumeStatusDeleting  VolumeStatus = "deleting"
	VolumeStatusError     VolumeStatus = "error"
)

type VolumeType struct {
	Description string `json:"description,omitempty"`

	// ID the type token (matches `volume_type` on a Volume).
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// VolumeTypeName customer-facing storage tier a volume is stored on. Mapped per-region
// to a Ceph pool by operator config (the mapping is opaque to API
// callers).
type VolumeTypeName string

// Values VolumeTypeName accepts.
const (
	VolumeTypeNameHDD  VolumeTypeName = "hdd"
	VolumeTypeNameSSD  VolumeTypeName = "ssd"
	VolumeTypeNameNVMe VolumeTypeName = "nvme"
)

type VolumeUpdateRequest struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
	Tags        Tags    `json:"tags,omitempty"`
}
