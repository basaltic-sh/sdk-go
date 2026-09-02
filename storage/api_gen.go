// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package storage

import (
	"context"
	"io"
	"iter"
	"net/url"
	"strconv"

	basaltic "github.com/basaltic-sh/sdk-go"
)

// ListBucketsParams are the optional filters and pagination controls for
// [Client.ListBuckets]. A nil *ListBucketsParams sends none of them.
type ListBucketsParams struct {
	Limit int

	// Marker resume token — the last bucket name from the previous page.
	Marker string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListBucketsParams) query() url.Values {
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
func (p *ListBucketsParams) withMarker(marker string) *ListBucketsParams {
	var out ListBucketsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListMultipartUploadsParams are the optional filters and pagination controls for
// [Client.ListMultipartUploads]. A nil *ListMultipartUploadsParams sends none of them.
type ListMultipartUploadsParams struct {
	MaxUploads int
	Prefix     string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListMultipartUploadsParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.MaxUploads != 0 {
		q.Set("max_uploads", strconv.Itoa(int(p.MaxUploads)))
	}
	if p.Prefix != "" {
		q.Set("prefix", p.Prefix)
	}
	return q
}

// ListObjectVersionsParams are the optional filters and pagination controls for
// [Client.ListObjectVersions]. A nil *ListObjectVersionsParams sends none of them.
type ListObjectVersionsParams struct {
	KeyMarker       string
	MaxKeys         int
	Prefix          string
	VersionIDMarker string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListObjectVersionsParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.KeyMarker != "" {
		q.Set("key_marker", p.KeyMarker)
	}
	if p.MaxKeys != 0 {
		q.Set("max_keys", strconv.Itoa(int(p.MaxKeys)))
	}
	if p.Prefix != "" {
		q.Set("prefix", p.Prefix)
	}
	if p.VersionIDMarker != "" {
		q.Set("version_id_marker", p.VersionIDMarker)
	}
	return q
}

// ListObjectsParams are the optional filters and pagination controls for
// [Client.ListObjects]. A nil *ListObjectsParams sends none of them.
type ListObjectsParams struct {
	Delimiter string
	Marker    string
	MaxKeys   int
	Prefix    string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListObjectsParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.Delimiter != "" {
		q.Set("delimiter", p.Delimiter)
	}
	if p.Marker != "" {
		q.Set("marker", p.Marker)
	}
	if p.MaxKeys != 0 {
		q.Set("max_keys", strconv.Itoa(int(p.MaxKeys)))
	}
	if p.Prefix != "" {
		q.Set("prefix", p.Prefix)
	}
	return q
}

// ListSnapshotPoliciesParams are the optional filters and pagination controls for
// [Client.ListSnapshotPolicies]. A nil *ListSnapshotPoliciesParams sends none of them.
type ListSnapshotPoliciesParams struct {
	// Enabled narrow to enabled (or paused) policies.
	Enabled *bool
	Limit   int

	// Marker resume token — the last policy id from the previous page.
	Marker string

	// Name case-insensitive substring match on the policy name.
	Name string

	// VolumeID narrow the listing to the policy attached to one volume.
	VolumeID string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListSnapshotPoliciesParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.Enabled != nil {
		q.Set("enabled", strconv.FormatBool(*p.Enabled))
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
	if p.VolumeID != "" {
		q.Set("volume_id", p.VolumeID)
	}
	return q
}

// withMarker copies p with the pagination cursor replaced, leaving the
// caller's value untouched across pages.
func (p *ListSnapshotPoliciesParams) withMarker(marker string) *ListSnapshotPoliciesParams {
	var out ListSnapshotPoliciesParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListSnapshotsParams are the optional filters and pagination controls for
// [Client.ListSnapshots]. A nil *ListSnapshotsParams sends none of them.
type ListSnapshotsParams struct {
	Limit int

	// Marker resume token — the last snapshot id from the previous page.
	Marker string

	// Name case-insensitive substring match on the snapshot name.
	Name   string
	Status SnapshotStatus

	// VolumeID narrow the listing to snapshots of one volume.
	VolumeID string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListSnapshotsParams) query() url.Values {
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
	if p.Name != "" {
		q.Set("name", p.Name)
	}
	if p.Status != "" {
		q.Set("status", string(p.Status))
	}
	if p.VolumeID != "" {
		q.Set("volume_id", p.VolumeID)
	}
	return q
}

// withMarker copies p with the pagination cursor replaced, leaving the
// caller's value untouched across pages.
func (p *ListSnapshotsParams) withMarker(marker string) *ListSnapshotsParams {
	var out ListSnapshotsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListVolumesParams are the optional filters and pagination controls for
// [Client.ListVolumes]. A nil *ListVolumesParams sends none of them.
type ListVolumesParams struct {
	Limit int

	// Marker resume token — the last volume id from the previous page.
	Marker string

	// Name case-insensitive substring match on the volume name.
	Name   string
	Status VolumeStatus
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListVolumesParams) query() url.Values {
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
	if p.Name != "" {
		q.Set("name", p.Name)
	}
	if p.Status != "" {
		q.Set("status", string(p.Status))
	}
	return q
}

// withMarker copies p with the pagination cursor replaced, leaving the
// caller's value untouched across pages.
func (p *ListVolumesParams) withMarker(marker string) *ListVolumesParams {
	var out ListVolumesParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// AbortMultipartUpload aborts a multipart upload.
func (c *Client) AbortMultipartUpload(ctx context.Context, bucket string, uploadID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "abortMultipartUpload",
		Method:   "DELETE",
		Path:     "/v1/buckets/{bucket}/multipart-uploads/{upload_id}",
		PathArgs: []string{bucket, uploadID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// CompleteMultipartUpload completes a multipart upload.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CompleteMultipartUpload(ctx context.Context, bucket string, uploadID string, body *CompleteMultipartUploadRequest, opts ...basaltic.RequestOption) (*CompleteMultipartUploadResponse, error) {
	op := &basaltic.Operation{
		ID:       "completeMultipartUpload",
		Method:   "POST",
		Path:     "/v1/buckets/{bucket}/multipart-uploads/{upload_id}/complete",
		PathArgs: []string{bucket, uploadID},
		Body:     body,
	}
	var out CompleteMultipartUploadResponse
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// CreateBucket creates bucket.
//
// Provision a new bucket. The name must be unique within the region (S3
// rule); 3-63 lowercase characters, digits, and hyphens; no
// leading/trailing hyphen.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateBucket(ctx context.Context, body *CreateBucketRequest, opts ...basaltic.RequestOption) (*Bucket, error) {
	op := &basaltic.Operation{
		ID:     "createBucket",
		Method: "POST",
		Path:   "/v1/buckets",
		Body:   body,
	}
	var out struct {
		Bucket *Bucket `json:"bucket"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Bucket, nil
}

// CreateSnapshot creates snapshot.
//
// Take a point-in-time snapshot of a volume. The row is persisted in
// status `creating` and the snapshot is provisioned asynchronously; poll
// GET until status changes to `available` (success) or `error`. Allowed
// against volumes in `available` or `in_use` (in-use volumes can be
// snapshotted).
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateSnapshot(ctx context.Context, body *SnapshotCreateRequest, opts ...basaltic.RequestOption) (*Snapshot, error) {
	op := &basaltic.Operation{
		ID:     "createSnapshot",
		Method: "POST",
		Path:   "/v1/snapshots",
		Body:   body,
	}
	var out struct {
		Snapshot *Snapshot `json:"snapshot"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Snapshot, nil
}

// CreateSnapshotPolicy creates snapshot policy.
//
// Attach a snapshot schedule to a volume. A volume has at most one
// policy; attaching a second is a 409.
//
// The first snapshot lands one `interval_minutes` from now — attaching
// a schedule is not itself a request for a snapshot. Use POST
// /v1/snapshots for one now.
//
// What retention can delete, since a schedule is also an automatic
// deleter: only the snapshots this policy itself took. A snapshot taken
// by hand is never reaped, and neither is one something depends on — a
// snapshot a volume was created from, including a restore still running,
// is skipped and looked at again later. Pausing the policy stops the
// deleting as well as the taking, and deleting the policy keeps every
// snapshot it already took.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateSnapshotPolicy(ctx context.Context, body *SnapshotPolicyCreateRequest, opts ...basaltic.RequestOption) (*SnapshotPolicy, error) {
	op := &basaltic.Operation{
		ID:     "createSnapshotPolicy",
		Method: "POST",
		Path:   "/v1/snapshot-policies",
		Body:   body,
	}
	var out struct {
		SnapshotPolicy *SnapshotPolicy `json:"snapshot_policy"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.SnapshotPolicy, nil
}

// CreateVolume creates volume.
//
// Accept a volume creation request. The row is persisted in status
// `creating` and provisioned asynchronously — poll GET until status
// changes to `available` (success) or `error` (provisioning failed; see
// `error_message`).
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateVolume(ctx context.Context, body *VolumeCreateRequest, opts ...basaltic.RequestOption) (*Volume, error) {
	op := &basaltic.Operation{
		ID:     "createVolume",
		Method: "POST",
		Path:   "/v1/volumes",
		Body:   body,
	}
	var out struct {
		Volume *Volume `json:"volume"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Volume, nil
}

// DeleteBucket deletes bucket.
//
// Delete an empty bucket. Returns 409 if any objects or in-flight
// multipart uploads remain.
func (c *Client) DeleteBucket(ctx context.Context, bucket string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteBucket",
		Method:   "DELETE",
		Path:     "/v1/buckets/{bucket}",
		PathArgs: []string{bucket},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteBucketCORS deletes bucket CORS configuration.
func (c *Client) DeleteBucketCORS(ctx context.Context, bucket string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteBucketCORS",
		Method:   "DELETE",
		Path:     "/v1/buckets/{bucket}/cors",
		PathArgs: []string{bucket},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteBucketEncryption deletes bucket encryption configuration.
func (c *Client) DeleteBucketEncryption(ctx context.Context, bucket string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteBucketEncryption",
		Method:   "DELETE",
		Path:     "/v1/buckets/{bucket}/encryption",
		PathArgs: []string{bucket},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteBucketLifecycle deletes bucket lifecycle configuration.
func (c *Client) DeleteBucketLifecycle(ctx context.Context, bucket string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteBucketLifecycle",
		Method:   "DELETE",
		Path:     "/v1/buckets/{bucket}/lifecycle",
		PathArgs: []string{bucket},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteBucketObjectLock deletes bucket object-lock configuration.
func (c *Client) DeleteBucketObjectLock(ctx context.Context, bucket string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteBucketObjectLock",
		Method:   "DELETE",
		Path:     "/v1/buckets/{bucket}/object-lock",
		PathArgs: []string{bucket},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteBucketPolicy deletes bucket policy.
func (c *Client) DeleteBucketPolicy(ctx context.Context, bucket string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteBucketPolicy",
		Method:   "DELETE",
		Path:     "/v1/buckets/{bucket}/policy",
		PathArgs: []string{bucket},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteBucketTagging deletes bucket tag set.
func (c *Client) DeleteBucketTagging(ctx context.Context, bucket string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteBucketTagging",
		Method:   "DELETE",
		Path:     "/v1/buckets/{bucket}/tagging",
		PathArgs: []string{bucket},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteObject deletes object.
//
// Delete an object (or a specific `?versionId`). Per-object subresources
// ride as query parameters on this method — `?tagging` deletes the
// object's tag set. `X-Amz-Bypass- Governance-Retention: true` overrides
// a GOVERNANCE lock and also requires
// `storage:BypassGovernanceRetention` on the object.
func (c *Client) DeleteObject(ctx context.Context, bucket string, key string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteObject",
		Method:   "DELETE",
		Path:     "/v1/buckets/{bucket}/objects/{key}",
		PathArgs: []string{bucket, key},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteSnapshot deletes snapshot.
//
// Accept a snapshot delete request. The row is flipped to `deleting` and
// the teardown is provisioned asynchronously. Poll GET until the row
// disappears (deleted) or changes to `error`.
func (c *Client) DeleteSnapshot(ctx context.Context, snapshotID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteSnapshot",
		Method:   "DELETE",
		Path:     "/v1/snapshots/{snapshot_id}",
		PathArgs: []string{snapshotID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteSnapshotPolicy deletes snapshot policy.
//
// Detach the schedule. The snapshots it already created are kept —
// they become ordinary snapshots you own outright, and are never reaped
// again. Deleting a schedule is a scheduling decision, not a request to
// destroy history; delete the snapshots themselves if that is what you
// want.
//
// Deleting the volume also removes its policy.
func (c *Client) DeleteSnapshotPolicy(ctx context.Context, policyID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteSnapshotPolicy",
		Method:   "DELETE",
		Path:     "/v1/snapshot-policies/{policy_id}",
		PathArgs: []string{policyID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteVolume deletes volume.
//
// Accept a volume delete request. The row is flipped to `deleting` and
// the teardown is provisioned asynchronously. Poll GET until the row
// disappears (deleted) or changes to `error` (teardown failed). Blocked
// on `in_use` (attached) and on transient statuses (`creating`,
// `extending`, `deleting`).
func (c *Client) DeleteVolume(ctx context.Context, volumeID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteVolume",
		Method:   "DELETE",
		Path:     "/v1/volumes/{volume_id}",
		PathArgs: []string{volumeID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// ExtendVolume extends volume.
//
// Accept a volume extend request. The row is flipped to `extending` and
// the resize is provisioned asynchronously. Poll GET until status
// changes back to `available` (with the new size) or `error`. The resize
// is online — once the volume is bigger, the guest sees the new size
// on its next rescan (rescans aren't triggered from here; that's the
// compute service's job).
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) ExtendVolume(ctx context.Context, volumeID string, body *VolumeExtendRequest, opts ...basaltic.RequestOption) (*Volume, error) {
	op := &basaltic.Operation{
		ID:       "extendVolume",
		Method:   "POST",
		Path:     "/v1/volumes/{volume_id}/extend",
		PathArgs: []string{volumeID},
		Body:     body,
	}
	var out struct {
		Volume *Volume `json:"volume"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Volume, nil
}

// GetBucketCORS gets bucket CORS configuration.
func (c *Client) GetBucketCORS(ctx context.Context, bucket string, opts ...basaltic.RequestOption) (*CORSConfig, error) {
	op := &basaltic.Operation{
		ID:       "getBucketCORS",
		Method:   "GET",
		Path:     "/v1/buckets/{bucket}/cors",
		PathArgs: []string{bucket},
	}
	var out struct {
		CORS *CORSConfig `json:"cors"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.CORS, nil
}

// GetBucketEncryption gets bucket encryption configuration.
func (c *Client) GetBucketEncryption(ctx context.Context, bucket string, opts ...basaltic.RequestOption) (*EncryptionConfig, error) {
	op := &basaltic.Operation{
		ID:       "getBucketEncryption",
		Method:   "GET",
		Path:     "/v1/buckets/{bucket}/encryption",
		PathArgs: []string{bucket},
	}
	var out struct {
		Encryption *EncryptionConfig `json:"encryption"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Encryption, nil
}

// GetBucketLifecycle gets bucket lifecycle configuration.
func (c *Client) GetBucketLifecycle(ctx context.Context, bucket string, opts ...basaltic.RequestOption) (*LifecycleConfig, error) {
	op := &basaltic.Operation{
		ID:       "getBucketLifecycle",
		Method:   "GET",
		Path:     "/v1/buckets/{bucket}/lifecycle",
		PathArgs: []string{bucket},
	}
	var out struct {
		Lifecycle *LifecycleConfig `json:"lifecycle"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Lifecycle, nil
}

// GetBucketObjectLock gets bucket object-lock configuration.
func (c *Client) GetBucketObjectLock(ctx context.Context, bucket string, opts ...basaltic.RequestOption) (*ObjectLockConfig, error) {
	op := &basaltic.Operation{
		ID:       "getBucketObjectLock",
		Method:   "GET",
		Path:     "/v1/buckets/{bucket}/object-lock",
		PathArgs: []string{bucket},
	}
	var out struct {
		ObjectLock *ObjectLockConfig `json:"object_lock"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.ObjectLock, nil
}

// GetBucketPolicy gets bucket policy.
func (c *Client) GetBucketPolicy(ctx context.Context, bucket string, opts ...basaltic.RequestOption) (BucketPolicy, error) {
	op := &basaltic.Operation{
		ID:       "getBucketPolicy",
		Method:   "GET",
		Path:     "/v1/buckets/{bucket}/policy",
		PathArgs: []string{bucket},
	}
	var out struct {
		Document BucketPolicy `json:"document"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Document, nil
}

// GetBucketTagging gets bucket tag set.
func (c *Client) GetBucketTagging(ctx context.Context, bucket string, opts ...basaltic.RequestOption) (TagSet, error) {
	op := &basaltic.Operation{
		ID:       "getBucketTagging",
		Method:   "GET",
		Path:     "/v1/buckets/{bucket}/tagging",
		PathArgs: []string{bucket},
	}
	var out struct {
		Tags TagSet `json:"tags"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Tags, nil
}

// GetBucketVersioning gets bucket versioning state.
//
// Returns the versioning state. A bucket that was never configured
// reports `Disabled` (not a 404).
func (c *Client) GetBucketVersioning(ctx context.Context, bucket string, opts ...basaltic.RequestOption) (string, error) {
	op := &basaltic.Operation{
		ID:       "getBucketVersioning",
		Method:   "GET",
		Path:     "/v1/buckets/{bucket}/versioning",
		PathArgs: []string{bucket},
	}
	var out struct {
		Status string `json:"status"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return "", err
	}
	return out.Status, nil
}

// GetObject downloads object.
//
// Download an object's bytes. Per-object subresources ride as query
// parameters — `?tagging` returns the object's tag set as JSON instead
// of the object body.
//
// The caller must close the returned reader.
func (c *Client) GetObject(ctx context.Context, bucket string, key string, opts ...basaltic.RequestOption) (io.ReadCloser, error) {
	op := &basaltic.Operation{
		ID:       "getObject",
		Method:   "GET",
		Path:     "/v1/buckets/{bucket}/objects/{key}",
		PathArgs: []string{bucket, key},
	}
	stream, _, err := c.rt.DoStream(ctx, op, opts...)
	if err != nil {
		return nil, err
	}
	return stream, nil
}

// GetSnapshot gets snapshot.
func (c *Client) GetSnapshot(ctx context.Context, snapshotID string, opts ...basaltic.RequestOption) (*Snapshot, error) {
	op := &basaltic.Operation{
		ID:       "getSnapshot",
		Method:   "GET",
		Path:     "/v1/snapshots/{snapshot_id}",
		PathArgs: []string{snapshotID},
	}
	var out struct {
		Snapshot *Snapshot `json:"snapshot"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Snapshot, nil
}

// GetSnapshotPolicy gets snapshot policy.
//
// Read one policy's schedule, retention window and run state.
//
// `last_error` is the field to check on a schedule that has stopped
// producing snapshots: it carries the reason the most recent run took
// none (quota exhausted, volume mid-extend, …) and is empty after a
// run that succeeded. `next_run_at` and `last_run_at` say where in the
// cycle the policy is.
func (c *Client) GetSnapshotPolicy(ctx context.Context, policyID string, opts ...basaltic.RequestOption) (*SnapshotPolicy, error) {
	op := &basaltic.Operation{
		ID:       "getSnapshotPolicy",
		Method:   "GET",
		Path:     "/v1/snapshot-policies/{policy_id}",
		PathArgs: []string{policyID},
	}
	var out struct {
		SnapshotPolicy *SnapshotPolicy `json:"snapshot_policy"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.SnapshotPolicy, nil
}

// GetVolume gets volume.
func (c *Client) GetVolume(ctx context.Context, volumeID string, opts ...basaltic.RequestOption) (*Volume, error) {
	op := &basaltic.Operation{
		ID:       "getVolume",
		Method:   "GET",
		Path:     "/v1/volumes/{volume_id}",
		PathArgs: []string{volumeID},
	}
	var out struct {
		Volume *Volume `json:"volume"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Volume, nil
}

// HeadBucket heads bucket.
func (c *Client) HeadBucket(ctx context.Context, bucket string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "headBucket",
		Method:   "HEAD",
		Path:     "/v1/buckets/{bucket}",
		PathArgs: []string{bucket},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// HeadObject heads object.
func (c *Client) HeadObject(ctx context.Context, bucket string, key string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "headObject",
		Method:   "HEAD",
		Path:     "/v1/buckets/{bucket}/objects/{key}",
		PathArgs: []string{bucket, key},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// InitiateMultipartUpload initiates a multipart upload.
//
// Stage an object as up to 10000 parts and assemble them in one step.
// Every part but the last must be at least 5 MiB, and no part may exceed
// 5 GiB — so an upload's parts can hold far more than any object a
// caller would send through it.
//
// The assembled object is written out in full when the upload is
// completed, and that write happens inside the completing request:
// budget for it at roughly 85 MiB/s, and note that the public edge ends
// any request at one hour. Objects beyond a few hundred GiB need a
// completion path this API does not offer yet.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) InitiateMultipartUpload(ctx context.Context, bucket string, body *InitiateMultipartUploadRequest, opts ...basaltic.RequestOption) (*MultipartUpload, error) {
	op := &basaltic.Operation{
		ID:       "initiateMultipartUpload",
		Method:   "POST",
		Path:     "/v1/buckets/{bucket}/multipart-uploads",
		PathArgs: []string{bucket},
		Body:     body,
	}
	var out struct {
		Upload *MultipartUpload `json:"upload"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Upload, nil
}

// ListBuckets lists buckets.
//
// List buckets owned by the requesting account, ordered by name
// ascending. Keyset-paginated on the bucket name (globally unique) —
// pass the last name from the previous page as `marker` to fetch the
// next.
//
// Returns one page. Use ListBucketsAll to walk every page.
func (c *Client) ListBuckets(ctx context.Context, params *ListBucketsParams, opts ...basaltic.RequestOption) (*basaltic.Page[Bucket], error) {
	op := &basaltic.Operation{
		ID:     "listBuckets",
		Method: "GET",
		Path:   "/v1/buckets",
	}
	op.Query = params.query()
	var out struct {
		Items []Bucket `json:"buckets"`
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
	page := &basaltic.Page[Bucket]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListBucketsAll walks every page of ListBuckets, yielding one item at a
// time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListBucketsAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListBucketsAll(ctx context.Context, params *ListBucketsParams, opts ...basaltic.RequestOption) iter.Seq2[Bucket, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Bucket], error) {
		return c.ListBuckets(ctx, params.withMarker(marker), opts...)
	})
}

// ListMultipartUploads lists in-flight multipart uploads.
func (c *Client) ListMultipartUploads(ctx context.Context, bucket string, params *ListMultipartUploadsParams, opts ...basaltic.RequestOption) (*basaltic.Page[MultipartUpload], error) {
	op := &basaltic.Operation{
		ID:       "listMultipartUploads",
		Method:   "GET",
		Path:     "/v1/buckets/{bucket}/multipart-uploads",
		PathArgs: []string{bucket},
	}
	op.Query = params.query()
	var out struct {
		Items []MultipartUpload `json:"uploads"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[MultipartUpload]{Items: out.Items}
	return page, nil
}

// ListObjectVersions lists object versions.
func (c *Client) ListObjectVersions(ctx context.Context, bucket string, params *ListObjectVersionsParams, opts ...basaltic.RequestOption) (*ListObjectVersionsResponse, error) {
	op := &basaltic.Operation{
		ID:       "listObjectVersions",
		Method:   "GET",
		Path:     "/v1/buckets/{bucket}/object-versions",
		PathArgs: []string{bucket},
	}
	op.Query = params.query()
	var out ListObjectVersionsResponse
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListObjects lists objects.
//
// List objects in the bucket. Supports prefix/delimiter for "folder"
// semantics and keyset pagination via `marker`.
func (c *Client) ListObjects(ctx context.Context, bucket string, params *ListObjectsParams, opts ...basaltic.RequestOption) (*ObjectListResponse, error) {
	op := &basaltic.Operation{
		ID:       "listObjects",
		Method:   "GET",
		Path:     "/v1/buckets/{bucket}/objects",
		PathArgs: []string{bucket},
	}
	op.Query = params.query()
	var out ObjectListResponse
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListParts lists uploaded parts.
func (c *Client) ListParts(ctx context.Context, bucket string, uploadID string, opts ...basaltic.RequestOption) (*ListPartsResponse, error) {
	op := &basaltic.Operation{
		ID:       "listParts",
		Method:   "GET",
		Path:     "/v1/buckets/{bucket}/multipart-uploads/{upload_id}/parts",
		PathArgs: []string{bucket, uploadID},
	}
	var out ListPartsResponse
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListSnapshotPolicies lists snapshot policies.
//
// List snapshot policies owned by the requesting organization in the
// target region, newest-first. Keyset-paginated by policy id (UUIDv7).
//
// Returns one page. Use ListSnapshotPoliciesAll to walk every page.
func (c *Client) ListSnapshotPolicies(ctx context.Context, params *ListSnapshotPoliciesParams, opts ...basaltic.RequestOption) (*basaltic.Page[SnapshotPolicy], error) {
	op := &basaltic.Operation{
		ID:     "listSnapshotPolicies",
		Method: "GET",
		Path:   "/v1/snapshot-policies",
	}
	op.Query = params.query()
	var out struct {
		Items []SnapshotPolicy `json:"snapshot_policies"`
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
	page := &basaltic.Page[SnapshotPolicy]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListSnapshotPoliciesAll walks every page of ListSnapshotPolicies,
// yielding one item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListSnapshotPoliciesAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListSnapshotPoliciesAll(ctx context.Context, params *ListSnapshotPoliciesParams, opts ...basaltic.RequestOption) iter.Seq2[SnapshotPolicy, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[SnapshotPolicy], error) {
		return c.ListSnapshotPolicies(ctx, params.withMarker(marker), opts...)
	})
}

// ListSnapshots lists snapshots.
//
// List snapshots owned by the requesting organization in the target
// region, newest-first. Optional `volume_id` narrows to one volume's
// snapshots. Keyset-paginated by snapshot id (UUIDv7).
//
// Returns one page. Use ListSnapshotsAll to walk every page.
func (c *Client) ListSnapshots(ctx context.Context, params *ListSnapshotsParams, opts ...basaltic.RequestOption) (*basaltic.Page[Snapshot], error) {
	op := &basaltic.Operation{
		ID:     "listSnapshots",
		Method: "GET",
		Path:   "/v1/snapshots",
	}
	op.Query = params.query()
	var out struct {
		Items []Snapshot `json:"snapshots"`
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
	page := &basaltic.Page[Snapshot]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListSnapshotsAll walks every page of ListSnapshots, yielding one item
// at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListSnapshotsAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListSnapshotsAll(ctx context.Context, params *ListSnapshotsParams, opts ...basaltic.RequestOption) iter.Seq2[Snapshot, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Snapshot], error) {
		return c.ListSnapshots(ctx, params.withMarker(marker), opts...)
	})
}

// ListVolumeTypes lists volume types.
func (c *Client) ListVolumeTypes(ctx context.Context, opts ...basaltic.RequestOption) (*basaltic.Page[VolumeType], error) {
	op := &basaltic.Operation{
		ID:     "listVolumeTypes",
		Method: "GET",
		Path:   "/v1/volume-types",
	}
	var out struct {
		Items []VolumeType `json:"volume_types"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[VolumeType]{Items: out.Items}
	return page, nil
}

// ListVolumes lists volumes.
//
// List volumes owned by the requesting organization in the target
// region, newest-first. Keyset-paginated by volume id (UUIDv7 sorts by
// creation time) — pass the last id from the previous page as `marker`
// to fetch the next.
//
// Returns one page. Use ListVolumesAll to walk every page.
func (c *Client) ListVolumes(ctx context.Context, params *ListVolumesParams, opts ...basaltic.RequestOption) (*basaltic.Page[Volume], error) {
	op := &basaltic.Operation{
		ID:     "listVolumes",
		Method: "GET",
		Path:   "/v1/volumes",
	}
	op.Query = params.query()
	var out struct {
		Items []Volume `json:"volumes"`
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
	page := &basaltic.Page[Volume]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListVolumesAll walks every page of ListVolumes, yielding one item at a
// time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListVolumesAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListVolumesAll(ctx context.Context, params *ListVolumesParams, opts ...basaltic.RequestOption) iter.Seq2[Volume, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Volume], error) {
		return c.ListVolumes(ctx, params.withMarker(marker), opts...)
	})
}

// PutBucketCORS puts bucket CORS configuration.
func (c *Client) PutBucketCORS(ctx context.Context, bucket string, body *PutBucketCORSRequest, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "putBucketCORS",
		Method:   "PUT",
		Path:     "/v1/buckets/{bucket}/cors",
		PathArgs: []string{bucket},
		Body:     body,
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// PutBucketDeletionProtection sets bucket deletion protection.
//
// Toggle deletion protection. When enabling, `recovery_days` sets the
// scheduled-deletion window (clamped to [1,30]; 0 uses the service
// default).
func (c *Client) PutBucketDeletionProtection(ctx context.Context, bucket string, body *PutBucketDeletionProtectionRequest, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "putBucketDeletionProtection",
		Method:   "PUT",
		Path:     "/v1/buckets/{bucket}/deletion-protection",
		PathArgs: []string{bucket},
		Body:     body,
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// PutBucketEncryption puts bucket encryption configuration.
//
// Only the AES256 SSE algorithm is supported.
func (c *Client) PutBucketEncryption(ctx context.Context, bucket string, body *PutBucketEncryptionRequest, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "putBucketEncryption",
		Method:   "PUT",
		Path:     "/v1/buckets/{bucket}/encryption",
		PathArgs: []string{bucket},
		Body:     body,
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// PutBucketLifecycle puts bucket lifecycle configuration.
func (c *Client) PutBucketLifecycle(ctx context.Context, bucket string, body *PutBucketLifecycleRequest, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "putBucketLifecycle",
		Method:   "PUT",
		Path:     "/v1/buckets/{bucket}/lifecycle",
		PathArgs: []string{bucket},
		Body:     body,
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// PutBucketObjectLock puts bucket object-lock configuration.
func (c *Client) PutBucketObjectLock(ctx context.Context, bucket string, body *PutBucketObjectLockRequest, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "putBucketObjectLock",
		Method:   "PUT",
		Path:     "/v1/buckets/{bucket}/object-lock",
		PathArgs: []string{bucket},
		Body:     body,
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// PutBucketPolicy puts bucket policy.
func (c *Client) PutBucketPolicy(ctx context.Context, bucket string, body *PutBucketPolicyRequest, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "putBucketPolicy",
		Method:   "PUT",
		Path:     "/v1/buckets/{bucket}/policy",
		PathArgs: []string{bucket},
		Body:     body,
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// PutBucketTagging puts bucket tag set.
func (c *Client) PutBucketTagging(ctx context.Context, bucket string, body *PutBucketTaggingRequest, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "putBucketTagging",
		Method:   "PUT",
		Path:     "/v1/buckets/{bucket}/tagging",
		PathArgs: []string{bucket},
		Body:     body,
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// PutBucketVersioning sets bucket versioning state.
func (c *Client) PutBucketVersioning(ctx context.Context, bucket string, body *PutBucketVersioningRequest, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "putBucketVersioning",
		Method:   "PUT",
		Path:     "/v1/buckets/{bucket}/versioning",
		PathArgs: []string{bucket},
		Body:     body,
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// PutObject uploads object.
//
// Upload an object. Per-object subresources ride as query parameters on
// this method — `?tagging` replaces the object's tag set (JSON body)
// and returns 204 with no body.
//
// A single upload carries at most 5 GiB. Anything larger is a multipart
// upload: initiate one under `/v1/buckets/{bucket}/multipart-uploads`,
// send the bytes as parts, and complete it. A body over the limit is
// refused with 413 whether it declares its size or streams past it.
func (c *Client) PutObject(ctx context.Context, bucket string, key string, body io.Reader, opts ...basaltic.RequestOption) (*PutObjectResponse, error) {
	op := &basaltic.Operation{
		ID:          "putObject",
		Method:      "PUT",
		Path:        "/v1/buckets/{bucket}/objects/{key}",
		PathArgs:    []string{bucket, key},
		Stream:      body,
		ContentType: "application/octet-stream",
	}
	var out PutObjectResponse
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// RestoreBucket restores a bucket pending deletion.
//
// Cancels a scheduled deletion during the recovery window (deletion
// protection was on when DeleteBucket was called).
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) RestoreBucket(ctx context.Context, bucket string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "restoreBucket",
		Method:   "POST",
		Path:     "/v1/buckets/{bucket}/restore",
		PathArgs: []string{bucket},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// UpdateSnapshot updates snapshot metadata.
//
// Update the snapshot's name or description.
func (c *Client) UpdateSnapshot(ctx context.Context, snapshotID string, body *SnapshotUpdateRequest, opts ...basaltic.RequestOption) (*Snapshot, error) {
	op := &basaltic.Operation{
		ID:       "updateSnapshot",
		Method:   "PATCH",
		Path:     "/v1/snapshots/{snapshot_id}",
		PathArgs: []string{snapshotID},
		Body:     body,
	}
	var out struct {
		Snapshot *Snapshot `json:"snapshot"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Snapshot, nil
}

// UpdateSnapshotPolicy updates snapshot policy.
//
// Change the schedule, the retention window, or pause the policy.
// Changing `interval_minutes` re-bases the next run off now.
func (c *Client) UpdateSnapshotPolicy(ctx context.Context, policyID string, body *SnapshotPolicyUpdateRequest, opts ...basaltic.RequestOption) (*SnapshotPolicy, error) {
	op := &basaltic.Operation{
		ID:       "updateSnapshotPolicy",
		Method:   "PATCH",
		Path:     "/v1/snapshot-policies/{policy_id}",
		PathArgs: []string{policyID},
		Body:     body,
	}
	var out struct {
		SnapshotPolicy *SnapshotPolicy `json:"snapshot_policy"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.SnapshotPolicy, nil
}

// UpdateVolume updates volume metadata.
//
// Update the volume's name or description.
func (c *Client) UpdateVolume(ctx context.Context, volumeID string, body *VolumeUpdateRequest, opts ...basaltic.RequestOption) (*Volume, error) {
	op := &basaltic.Operation{
		ID:       "updateVolume",
		Method:   "PATCH",
		Path:     "/v1/volumes/{volume_id}",
		PathArgs: []string{volumeID},
		Body:     body,
	}
	var out struct {
		Volume *Volume `json:"volume"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Volume, nil
}

// UploadPart uploads a part.
//
// Store one part of an in-flight upload. A part holds at most 5 GiB;
// every part except the last one named at completion must be at least 5
// MiB, which completion is where that floor is checked.
func (c *Client) UploadPart(ctx context.Context, bucket string, uploadID string, partNumber string, body io.Reader, opts ...basaltic.RequestOption) (*UploadPartResponse, error) {
	op := &basaltic.Operation{
		ID:          "uploadPart",
		Method:      "PUT",
		Path:        "/v1/buckets/{bucket}/multipart-uploads/{upload_id}/parts/{part_number}",
		PathArgs:    []string{bucket, uploadID, partNumber},
		Stream:      body,
		ContentType: "application/octet-stream",
	}
	var out UploadPartResponse
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}
