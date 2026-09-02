// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package compute

import (
	"context"
	"io"
	"iter"
	"net/url"
	"strconv"

	basaltic "github.com/basaltic-sh/sdk-go"
)

// GetConsoleOutputParams are the optional filters and pagination controls for
// [Client.GetConsoleOutput]. A nil *GetConsoleOutputParams sends none of them.
type GetConsoleOutputParams struct {
	// MaxBytes return at most this many bytes from the END of the transcript. A
	// ceiling you may lower, not raise: values above the 65536-byte
	// maximum are clamped to it.
	MaxBytes int
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *GetConsoleOutputParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.MaxBytes != 0 {
		q.Set("max_bytes", strconv.Itoa(int(p.MaxBytes)))
	}
	return q
}

// ListFlavorsParams are the optional filters and pagination controls for
// [Client.ListFlavors]. A nil *ListFlavorsParams sends none of them.
type ListFlavorsParams struct {
	// Family filter by product family. Load-balancer and database create flows
	// should list their own family; regular instances use "general".
	//
	// One of: "general", "loadbalancer", "database".
	Family string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListFlavorsParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.Family != "" {
		q.Set("family", p.Family)
	}
	return q
}

// ListImagesParams are the optional filters and pagination controls for
// [Client.ListImages]. A nil *ListImagesParams sends none of them.
type ListImagesParams struct {
	// AllVersions include builds a newer version has superseded. Off by default, when
	// each tag contributes only the build worth launching.
	AllVersions  *bool
	Architecture string

	// Limit maximum number of items to return. A value above the maximum is
	// clamped to it rather than rejected, so a page shorter than the one
	// you asked for is normal — page until `meta.has_more` is false, not
	// until a page looks short.
	Limit int

	// Marker opaque pagination cursor. Echo back the `meta.marker` value from the
	// previous page to fetch the next one; do not construct or parse it.
	// The token's internal form varies by endpoint (a resource ID, a
	// timestamp, …) and is not guaranteed stable across releases.
	Marker string

	// Name substring match on name
	Name string
	OS   string

	// Status one of: "pending", "importing", "active", "error", "hidden".
	Status string

	// Visibility one of: "public", "private".
	Visibility string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListImagesParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.AllVersions != nil {
		q.Set("all_versions", strconv.FormatBool(*p.AllVersions))
	}
	if p.Architecture != "" {
		q.Set("architecture", p.Architecture)
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
	if p.OS != "" {
		q.Set("os", p.OS)
	}
	if p.Status != "" {
		q.Set("status", p.Status)
	}
	if p.Visibility != "" {
		q.Set("visibility", p.Visibility)
	}
	return q
}

// withMarker copies p with the pagination cursor replaced, leaving the
// caller's value untouched across pages.
func (p *ListImagesParams) withMarker(marker string) *ListImagesParams {
	var out ListImagesParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListInstancePoolsParams are the optional filters and pagination controls for
// [Client.ListInstancePools]. A nil *ListInstancePoolsParams sends none of them.
type ListInstancePoolsParams struct {
	// Limit maximum number of items to return. A value above the maximum is
	// clamped to it rather than rejected, so a page shorter than the one
	// you asked for is normal — page until `meta.has_more` is false, not
	// until a page looks short.
	Limit int

	// Marker opaque pagination cursor. Echo back the `meta.marker` value from the
	// previous page to fetch the next one; do not construct or parse it.
	// The token's internal form varies by endpoint (a resource ID, a
	// timestamp, …) and is not guaranteed stable across releases.
	Marker string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListInstancePoolsParams) query() url.Values {
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
func (p *ListInstancePoolsParams) withMarker(marker string) *ListInstancePoolsParams {
	var out ListInstancePoolsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListInstancesParams are the optional filters and pagination controls for
// [Client.ListInstances]. A nil *ListInstancesParams sends none of them.
type ListInstancesParams struct {
	// FlavorID filter by flavor ID
	FlavorID string

	// ImageID filter by image ID
	ImageID string

	// Limit maximum number of items to return. A value above the maximum is
	// clamped to it rather than rejected, so a page shorter than the one
	// you asked for is normal — page until `meta.has_more` is false, not
	// until a page looks short.
	Limit int

	// Marker opaque pagination cursor. Echo back the `meta.marker` value from the
	// previous page to fetch the next one; do not construct or parse it.
	// The token's internal form varies by endpoint (a resource ID, a
	// timestamp, …) and is not guaranteed stable across releases.
	Marker string

	// Name filter by name (exact match or prefix with *)
	Name string

	// VMState filter by lifecycle state.
	VMState VMState
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListInstancesParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.FlavorID != "" {
		q.Set("flavor_id", p.FlavorID)
	}
	if p.ImageID != "" {
		q.Set("image_id", p.ImageID)
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
	if p.VMState != "" {
		q.Set("vm_state", string(p.VMState))
	}
	return q
}

// withMarker copies p with the pagination cursor replaced, leaving the
// caller's value untouched across pages.
func (p *ListInstancesParams) withMarker(marker string) *ListInstancesParams {
	var out ListInstancesParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListKeypairsParams are the optional filters and pagination controls for
// [Client.ListKeypairs]. A nil *ListKeypairsParams sends none of them.
type ListKeypairsParams struct {
	// Limit maximum number of items to return. A value above the maximum is
	// clamped to it rather than rejected, so a page shorter than the one
	// you asked for is normal — page until `meta.has_more` is false, not
	// until a page looks short.
	Limit int

	// Marker opaque pagination cursor. Echo back the `meta.marker` value from the
	// previous page to fetch the next one; do not construct or parse it.
	// The token's internal form varies by endpoint (a resource ID, a
	// timestamp, …) and is not guaranteed stable across releases.
	Marker string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListKeypairsParams) query() url.Values {
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
func (p *ListKeypairsParams) withMarker(marker string) *ListKeypairsParams {
	var out ListKeypairsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// StartSerialConsoleParams are the optional filters and pagination controls for
// [Client.StartSerialConsole]. A nil *StartSerialConsoleParams sends none of them.
type StartSerialConsoleParams struct {
	// BacklogBytes replay this many bytes of already-written output before live output
	// begins, so attaching to a quiet guest shows why it is quiet instead
	// of an empty screen. 0 disables replay. Values above 65536 are
	// clamped. Replay and live output are the same stream read forward, so
	// nothing is lost or duplicated at the join — including output the
	// guest produces while you are connecting.
	BacklogBytes int
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *StartSerialConsoleParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.BacklogBytes != 0 {
		q.Set("backlog_bytes", strconv.Itoa(int(p.BacklogBytes)))
	}
	return q
}

// AttachInstanceNIC attaches a NIC to a running instance.
//
// Provisions a new network interface in the given subnet (address
// allocation + optional security-group attachments). The device is
// hot-plugged into the instance asynchronously — the instance must be
// running or stopped.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) AttachInstanceNIC(ctx context.Context, instanceID string, body *AttachInstanceNICRequest, opts ...basaltic.RequestOption) (*AttachInstanceNICAttachment, error) {
	op := &basaltic.Operation{
		ID:       "attachInstanceNIC",
		Method:   "POST",
		Path:     "/v1/instances/{instance_id}/nics",
		PathArgs: []string{instanceID},
		Body:     body,
	}
	var out struct {
		Attachment *AttachInstanceNICAttachment `json:"attachment"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Attachment, nil
}

// AttachInstancePoolFloatingIP gives the pool a shared public address.
//
// Binds an already-allocated floating IP of yours to the pool. One
// public IP, answered by every replica — an anycast address — as
// opposed to `template.assign_public_ip`, which gives each replica its
// own.
//
// THE OPERATION NAMES THE POOL because membership is then maintained for
// you: the address's members are the pool's live replicas — every one
// of them, co-resident ones included — so a scale-out joins, a
// scale-in leaves, and a replaced member is swapped, with no per-replica
// attach to make. The same address can be built by hand at `POST
// /v1/floating-ips/{floating_ip_id}/attach`, one interface at a time;
// what the pool adds is that nobody has to keep it in step.
//
// The floating IP must be unattached, and the pool's subnet must already
// route `0.0.0.0/0` to an internet gateway. Idempotent: re-attaching the
// same address to the same pool returns it unchanged.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) AttachInstancePoolFloatingIP(ctx context.Context, poolID string, body *InstancePoolFloatingIPAttachRequest, opts ...basaltic.RequestOption) (*FloatingIP, error) {
	op := &basaltic.Operation{
		ID:       "attachInstancePoolFloatingIp",
		Method:   "POST",
		Path:     "/v1/instance-pools/{pool_id}/floating-ips",
		PathArgs: []string{poolID},
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

// AttachInstanceVolume attaches a data volume to an instance.
//
// Attach an existing, available storage volume to the instance and
// republish the desired spec so the on-host compute-agent hot-plugs the
// RBD disk into the running domain. The volume must be in the same
// account and in status `available`.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) AttachInstanceVolume(ctx context.Context, instanceID string, body *AttachInstanceVolumeRequest, opts ...basaltic.RequestOption) (*AttachInstanceVolumeAttachment, error) {
	op := &basaltic.Operation{
		ID:       "attachInstanceVolume",
		Method:   "POST",
		Path:     "/v1/instances/{instance_id}/volumes",
		PathArgs: []string{instanceID},
		Body:     body,
	}
	var out struct {
		Attachment *AttachInstanceVolumeAttachment `json:"attachment"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Attachment, nil
}

// CreateImage imports an image from an object URL.
//
// Registers an image for import from a presigned object URL — no image
// bytes flow through this API. Upload your disk to any bucket you
// control (our object store, AWS S3, MinIO, …) using a robust
// multipart S3 client, then pass a presigned GET URL as `source_url`.
//
// The import runs in the background: the response is 202 with
// status=importing, and a worker fetches the URL, converts it to the raw
// base (qcow2 / raw / vmdk / vhd / vhdx / vdi are accepted), and imports
// it into Ceph. The row flips to active (or error, with import_error
// set) once it finishes — poll GET /v1/images/{image_id} for the
// status.
//
// A name behaves like a movable tag: by default the new image becomes
// the "current" version for its (name, architecture), so future launches
// of that name boot the new bits. Older versions stay bootable by id and
// by `name:version`. Pass `current: false` to stage a version without
// switching, then promote it later with PATCH /v1/images/{image_id}
// (current=true).
//
// `version` identifies the build within the name and must be unique
// there; omit it and the server stamps a UTC timestamp. Publishing twice
// under one name with the same version is a 409, not a second anonymous
// build.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateImage(ctx context.Context, body *ImageCreateRequest, opts ...basaltic.RequestOption) (*Image, error) {
	op := &basaltic.Operation{
		ID:     "createImage",
		Method: "POST",
		Path:   "/v1/images",
		Body:   body,
	}
	var out struct {
		Image *Image `json:"image"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Image, nil
}

// CreateInstance creates instance.
//
// Create a new compute instance.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateInstance(ctx context.Context, body *InstanceCreateRequest, opts ...basaltic.RequestOption) (*Instance, error) {
	op := &basaltic.Operation{
		ID:     "createInstance",
		Method: "POST",
		Path:   "/v1/instances",
		Body:   body,
	}
	var out struct {
		Instance *Instance `json:"instance"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Instance, nil
}

// CreateInstancePool creates an instance pool.
//
// Create a pool from a launch template and a desired count. The
// desired_count instances are spawned synchronously; a background
// reconciler then converges member_count toward desired_count as it's
// changed. min_count/max_count default to desired_count.
//
// The top-level `tags` label the pool itself; `template.tags` are
// stamped on every instance it launches.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateInstancePool(ctx context.Context, body *InstancePoolCreateRequest, opts ...basaltic.RequestOption) (*InstancePool, error) {
	op := &basaltic.Operation{
		ID:     "createInstancePool",
		Method: "POST",
		Path:   "/v1/instance-pools",
		Body:   body,
	}
	var out struct {
		InstancePool *InstancePool `json:"instance_pool"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.InstancePool, nil
}

// CreateKeypair creates keypair.
//
// Create a new SSH keypair. If `public_key` is provided, it will be
// imported. If not provided, a new keypair will be generated and the
// private key returned.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateKeypair(ctx context.Context, body *KeypairCreateRequest, opts ...basaltic.RequestOption) (*CreateKeypairKeypair, error) {
	op := &basaltic.Operation{
		ID:     "createKeypair",
		Method: "POST",
		Path:   "/v1/keypairs",
		Body:   body,
	}
	var out struct {
		Keypair *CreateKeypairKeypair `json:"keypair"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Keypair, nil
}

// DeleteImage deletes (hide) an image.
//
// Soft-delete: the catalog row is flipped to status=hidden so in-flight
// clones can still complete. The underlying image data is reclaimed as
// soon as nothing is cloned from it — immediately when the image has
// no running instances, and by a background sweep otherwise.
//
// An image the catalog withdrew at end-of-life keeps its data on
// purpose, so deleting it by id is how you ask for that data to go.
// Deleting an already-deleted image is a 404.
func (c *Client) DeleteImage(ctx context.Context, imageID string, opts ...basaltic.RequestOption) (*Image, error) {
	op := &basaltic.Operation{
		ID:       "deleteImage",
		Method:   "DELETE",
		Path:     "/v1/images/{image_id}",
		PathArgs: []string{imageID},
	}
	var out struct {
		Image *Image `json:"image"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Image, nil
}

// DeleteInstance deletes instance.
//
// Request instance deletion. Returns 202 Accepted — the delete is
// async: the instance transitions to `deleting` and the on-host
// compute-agent tears the domain down before it reaches `deleted`.
func (c *Client) DeleteInstance(ctx context.Context, instanceID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteInstance",
		Method:   "DELETE",
		Path:     "/v1/instances/{instance_id}",
		PathArgs: []string{instanceID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteInstancePool deletes an instance pool.
//
// Tear down every instance the pool owns and drop the pool. Idempotent.
func (c *Client) DeleteInstancePool(ctx context.Context, poolID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteInstancePool",
		Method:   "DELETE",
		Path:     "/v1/instance-pools/{pool_id}",
		PathArgs: []string{poolID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteKeypair deletes keypair.
//
// Delete a keypair.
func (c *Client) DeleteKeypair(ctx context.Context, keypairID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteKeypair",
		Method:   "DELETE",
		Path:     "/v1/keypairs/{keypair_id}",
		PathArgs: []string{keypairID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DetachInstanceNIC detaches a NIC from a running instance.
//
// Hot-unplugs the network interface from the instance, then tears down
// the interface, its address allocation, and any security-group
// bindings. Refuses to detach the only NIC on the instance.
func (c *Client) DetachInstanceNIC(ctx context.Context, instanceID string, interfaceID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "detachInstanceNIC",
		Method:   "DELETE",
		Path:     "/v1/instances/{instance_id}/nics/{interface_id}",
		PathArgs: []string{instanceID, interfaceID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DetachInstancePoolFloatingIP takes a shared address off the pool.
//
// Removes the address from the pool and stops routing it. The ADDRESS IS
// NOT RELEASED — you allocated it, it stays yours, unattached, to
// reuse or release with `DELETE /v1/floating-ips/{floating_ip_id}`.
//
// Idempotent: detaching an address the pool does not hold returns 204.
func (c *Client) DetachInstancePoolFloatingIP(ctx context.Context, poolID string, floatingIPID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "detachInstancePoolFloatingIp",
		Method:   "DELETE",
		Path:     "/v1/instance-pools/{pool_id}/floating-ips/{floating_ip_id}",
		PathArgs: []string{poolID, floatingIPID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DetachInstanceVolume detaches a data volume from an instance.
//
// Remove the volume from the desired spec; the compute-agent hot-unplugs
// the RBD disk, then compute clears the attachment.
func (c *Client) DetachInstanceVolume(ctx context.Context, instanceID string, volumeID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "detachInstanceVolume",
		Method:   "DELETE",
		Path:     "/v1/instances/{instance_id}/volumes/{volume_id}",
		PathArgs: []string{instanceID, volumeID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// GetConsoleOutput gets the instance's serial console output.
//
// Returns the guest's serial console transcript — what it wrote to
// ttyS0 during its current boot.
//
// This is the diagnostic for an instance that never came up, which is
// exactly when SSH cannot answer: a bad fstab, a wrong kernel, a
// security group that locked you out, a cloud-init failure. Nothing has
// to be installed in the guest for it to work.
//
// The transcript covers the CURRENT boot only — it is reset each time
// the instance starts — so a crashed guest's last transcript is gone
// once it restarts. An instance that has never booted returns an empty
// output rather than an error.
func (c *Client) GetConsoleOutput(ctx context.Context, instanceID string, params *GetConsoleOutputParams, opts ...basaltic.RequestOption) (*GetConsoleOutputResult, error) {
	op := &basaltic.Operation{
		ID:       "getConsoleOutput",
		Method:   "GET",
		Path:     "/v1/instances/{instance_id}/console/output",
		PathArgs: []string{instanceID},
	}
	op.Query = params.query()
	var out GetConsoleOutputResult
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetConsoleScreenshot captures the instance's display.
//
// Returns a still image of what the instance's screen is showing right
// now.
//
// The counterpart to console output, for everything a serial transcript
// cannot reach: a guest stuck in its boot manager, sitting at a GRUB
// prompt, panicking before serial init, or booted from an image whose
// kernel was never told to log to ttyS0. In those the transcript is
// empty and the screen holds the whole answer.
//
// A still, not an interactive session — there is no remote desktop.
// Requires the instance to be running; a stopped instance has no display
// to capture and returns 409 rather than a blank frame.
//
// The caller must close the returned reader.
func (c *Client) GetConsoleScreenshot(ctx context.Context, instanceID string, opts ...basaltic.RequestOption) (io.ReadCloser, error) {
	op := &basaltic.Operation{
		ID:       "getConsoleScreenshot",
		Method:   "GET",
		Path:     "/v1/instances/{instance_id}/console/screenshot",
		PathArgs: []string{instanceID},
	}
	stream, _, err := c.rt.DoStream(ctx, op, opts...)
	if err != nil {
		return nil, err
	}
	return stream, nil
}

// GetFlavor gets flavor.
//
// Get details of a specific flavor.
func (c *Client) GetFlavor(ctx context.Context, flavorID string, opts ...basaltic.RequestOption) (*Flavor, error) {
	op := &basaltic.Operation{
		ID:       "getFlavor",
		Method:   "GET",
		Path:     "/v1/flavors/{flavor_id}",
		PathArgs: []string{flavorID},
	}
	var out struct {
		Flavor *Flavor `json:"flavor"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Flavor, nil
}

// GetImage gets an image.
func (c *Client) GetImage(ctx context.Context, imageID string, opts ...basaltic.RequestOption) (*Image, error) {
	op := &basaltic.Operation{
		ID:       "getImage",
		Method:   "GET",
		Path:     "/v1/images/{image_id}",
		PathArgs: []string{imageID},
	}
	var out struct {
		Image *Image `json:"image"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Image, nil
}

// GetInstance gets instance.
//
// Get details of a specific instance.
func (c *Client) GetInstance(ctx context.Context, instanceID string, opts ...basaltic.RequestOption) (*Instance, error) {
	op := &basaltic.Operation{
		ID:       "getInstance",
		Method:   "GET",
		Path:     "/v1/instances/{instance_id}",
		PathArgs: []string{instanceID},
	}
	var out struct {
		Instance *Instance `json:"instance"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Instance, nil
}

// GetInstancePool gets an instance pool.
func (c *Client) GetInstancePool(ctx context.Context, poolID string, opts ...basaltic.RequestOption) (*InstancePool, error) {
	op := &basaltic.Operation{
		ID:       "getInstancePool",
		Method:   "GET",
		Path:     "/v1/instance-pools/{pool_id}",
		PathArgs: []string{poolID},
	}
	var out struct {
		InstancePool *InstancePool `json:"instance_pool"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.InstancePool, nil
}

// GetKeypair gets keypair.
//
// Get details of a specific keypair.
func (c *Client) GetKeypair(ctx context.Context, keypairID string, opts ...basaltic.RequestOption) (*Keypair, error) {
	op := &basaltic.Operation{
		ID:       "getKeypair",
		Method:   "GET",
		Path:     "/v1/keypairs/{keypair_id}",
		PathArgs: []string{keypairID},
	}
	var out struct {
		Keypair *Keypair `json:"keypair"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Keypair, nil
}

// ListFlavors lists flavors.
//
// List all available instance flavors. The catalog is small and comes
// back in one shot — this endpoint is not paginated.
func (c *Client) ListFlavors(ctx context.Context, params *ListFlavorsParams, opts ...basaltic.RequestOption) (*basaltic.Page[Flavor], error) {
	op := &basaltic.Operation{
		ID:     "listFlavors",
		Method: "GET",
		Path:   "/v1/flavors",
	}
	op.Query = params.query()
	var out struct {
		Items []Flavor `json:"flavors"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[Flavor]{Items: out.Items}
	return page, nil
}

// ListImages lists images.
//
// List images visible to the requesting account: public platform images
// plus any images the account owns. Private platform images are staging
// and visible only to the platform account.
//
// One row per tag. A build drops out of this listing once a *newer*
// build holds its name — superseded history, which stays bootable by
// id and by `name:version`. Anything still importing or errored stays
// listed whatever its age, and so does a version staged with `current:
// false` that is newer than the current one. Pass `all_versions=true`
// for a tag's whole history.
//
// Returns one page. Use ListImagesAll to walk every page.
func (c *Client) ListImages(ctx context.Context, params *ListImagesParams, opts ...basaltic.RequestOption) (*basaltic.Page[Image], error) {
	op := &basaltic.Operation{
		ID:     "listImages",
		Method: "GET",
		Path:   "/v1/images",
	}
	op.Query = params.query()
	var out struct {
		Items []Image `json:"images"`
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
	page := &basaltic.Page[Image]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListImagesAll walks every page of ListImages, yielding one item at a
// time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListImagesAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListImagesAll(ctx context.Context, params *ListImagesParams, opts ...basaltic.RequestOption) iter.Seq2[Image, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Image], error) {
		return c.ListImages(ctx, params.withMarker(marker), opts...)
	})
}

// ListInstanceNiCs lists the instance's network interfaces.
//
// Returns the instance's NIC bindings ordered by boot index (primary
// first), each resolved with the interface's current MAC/IP/subnet
// detail and with the floating IP attached to that interface.
//
// This is the read that answers "what are this instance's public
// addresses". The instance's own `public_ip` reports the primary NIC
// alone, so an address that `networks[].assign_public_ip` put on a
// secondary NIC appears here and nowhere else.
func (c *Client) ListInstanceNiCs(ctx context.Context, instanceID string, opts ...basaltic.RequestOption) ([]*ListInstanceNiCsNIC, error) {
	op := &basaltic.Operation{
		ID:       "listInstanceNICs",
		Method:   "GET",
		Path:     "/v1/instances/{instance_id}/nics",
		PathArgs: []string{instanceID},
	}
	var out struct {
		NICs []*ListInstanceNiCsNIC `json:"nics"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.NICs, nil
}

// ListInstancePoolFloatingIPs lists the pool's shared public addresses.
//
// The addresses the WHOLE pool answers on. Not the per-replica addresses
// `template.assign_public_ip` allocates — those belong to the replica
// and are read from the instance.
func (c *Client) ListInstancePoolFloatingIPs(ctx context.Context, poolID string, opts ...basaltic.RequestOption) (*basaltic.Page[FloatingIP], error) {
	op := &basaltic.Operation{
		ID:       "listInstancePoolFloatingIps",
		Method:   "GET",
		Path:     "/v1/instance-pools/{pool_id}/floating-ips",
		PathArgs: []string{poolID},
	}
	var out struct {
		Items []FloatingIP `json:"floating_ips"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[FloatingIP]{Items: out.Items}
	return page, nil
}

// ListInstancePools lists instance pools.
//
// List the account's customer-managed instance pools. Pools created
// internally by other services (e.g. the load balancer's amphora fleet)
// are filtered out.
//
// Returns one page. Use ListInstancePoolsAll to walk every page.
func (c *Client) ListInstancePools(ctx context.Context, params *ListInstancePoolsParams, opts ...basaltic.RequestOption) (*basaltic.Page[InstancePool], error) {
	op := &basaltic.Operation{
		ID:     "listInstancePools",
		Method: "GET",
		Path:   "/v1/instance-pools",
	}
	op.Query = params.query()
	var out struct {
		Items []InstancePool `json:"instance_pools"`
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
	page := &basaltic.Page[InstancePool]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListInstancePoolsAll walks every page of ListInstancePools, yielding
// one item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListInstancePoolsAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListInstancePoolsAll(ctx context.Context, params *ListInstancePoolsParams, opts ...basaltic.RequestOption) iter.Seq2[InstancePool, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[InstancePool], error) {
		return c.ListInstancePools(ctx, params.withMarker(marker), opts...)
	})
}

// ListInstanceVolumes lists the instance's attached volumes.
//
// Returns the instance's volume bindings ordered by boot index (boot
// disk first), each resolved with the volume's current
// name/type/size/status.
func (c *Client) ListInstanceVolumes(ctx context.Context, instanceID string, opts ...basaltic.RequestOption) ([]*ListInstanceVolumesAttachment, error) {
	op := &basaltic.Operation{
		ID:       "listInstanceVolumes",
		Method:   "GET",
		Path:     "/v1/instances/{instance_id}/volumes",
		PathArgs: []string{instanceID},
	}
	var out struct {
		Attachments []*ListInstanceVolumesAttachment `json:"attachments"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Attachments, nil
}

// ListInstances lists instances.
//
// List all instances in the current organization.
//
// Returns one page. Use ListInstancesAll to walk every page.
func (c *Client) ListInstances(ctx context.Context, params *ListInstancesParams, opts ...basaltic.RequestOption) (*basaltic.Page[Instance], error) {
	op := &basaltic.Operation{
		ID:     "listInstances",
		Method: "GET",
		Path:   "/v1/instances",
	}
	op.Query = params.query()
	var out struct {
		Items []Instance `json:"instances"`
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
	page := &basaltic.Page[Instance]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListInstancesAll walks every page of ListInstances, yielding one item
// at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListInstancesAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListInstancesAll(ctx context.Context, params *ListInstancesParams, opts ...basaltic.RequestOption) iter.Seq2[Instance, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Instance], error) {
		return c.ListInstances(ctx, params.withMarker(marker), opts...)
	})
}

// ListKeypairs lists keypairs.
//
// List all SSH keypairs.
//
// Returns one page. Use ListKeypairsAll to walk every page.
func (c *Client) ListKeypairs(ctx context.Context, params *ListKeypairsParams, opts ...basaltic.RequestOption) (*basaltic.Page[Keypair], error) {
	op := &basaltic.Operation{
		ID:     "listKeypairs",
		Method: "GET",
		Path:   "/v1/keypairs",
	}
	op.Query = params.query()
	var out struct {
		Items []Keypair `json:"keypairs"`
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
	page := &basaltic.Page[Keypair]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListKeypairsAll walks every page of ListKeypairs, yielding one item at
// a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListKeypairsAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListKeypairsAll(ctx context.Context, params *ListKeypairsParams, opts ...basaltic.RequestOption) iter.Seq2[Keypair, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Keypair], error) {
		return c.ListKeypairs(ctx, params.withMarker(marker), opts...)
	})
}

// ListPoolInstances lists a pool's instance bindings.
//
// The live (pool, instance) bindings, each with its stable sequence
// number.
func (c *Client) ListPoolInstances(ctx context.Context, poolID string, opts ...basaltic.RequestOption) (*basaltic.Page[PoolInstance], error) {
	op := &basaltic.Operation{
		ID:       "listPoolInstances",
		Method:   "GET",
		Path:     "/v1/instance-pools/{pool_id}/instances",
		PathArgs: []string{poolID},
	}
	var out struct {
		Items []PoolInstance `json:"instances"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[PoolInstance]{Items: out.Items}
	return page, nil
}

// RebootInstance reboots instance.
//
// Reboot an instance (soft or hard).
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) RebootInstance(ctx context.Context, instanceID string, body *InstanceRebootRequest, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "rebootInstance",
		Method:   "POST",
		Path:     "/v1/instances/{instance_id}/reboot",
		PathArgs: []string{instanceID},
		Body:     body,
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// RefreshInstancePool rolls every member onto the pool's current launch template.
//
// Starts a rolling replacement of every instance not launched from the
// pool's current template, including any that predate template tracking.
//
// Asynchronous, and deliberately so: each replacement is a VM boot, and
// a request that waited would time out long before a pool of any size
// finished. The reconciler replaces ONE member per pass, and only once
// the pool is back at size with every member running — so a template
// that does not boot stalls the roll with the pool intact instead of
// walking it down one instance at a time.
//
// Capacity does not dip. The pool runs one instance over its target for
// the duration so a replacement is serving before anything is retired,
// unless it is already at max_count, where it replaces in place instead.
//
// Idempotent: asking again while a roll is running is accepted and does
// not restart it. Watch `refresh_in_progress` and `stale_instance_count`
// on the pool for progress.
func (c *Client) RefreshInstancePool(ctx context.Context, poolID string, opts ...basaltic.RequestOption) (*InstancePool, error) {
	op := &basaltic.Operation{
		ID:       "refreshInstancePool",
		Method:   "POST",
		Path:     "/v1/instance-pools/{pool_id}/refresh",
		PathArgs: []string{poolID},
	}
	var out struct {
		InstancePool *InstancePool `json:"instance_pool"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.InstancePool, nil
}

// ReinstallInstance reinstalls instance.
//
// Re-image a STOPPED instance's boot volume from an image (the current
// one, or a new image_id), keeping the instance's identity — id, name,
// IPs, keypairs, and cloud-init seed. The replacement is sized and
// tiered by size_gb and volume_type, defaulting to the image's
// min_disk_gb on the region default tier. The old boot volume is
// deleted; attached data volumes are untouched. The new OS is applied
// when the compute-agent redefines the domain on the next start.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) ReinstallInstance(ctx context.Context, instanceID string, body *ReinstallInstanceRequest, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "reinstallInstance",
		Method:   "POST",
		Path:     "/v1/instances/{instance_id}/reinstall",
		PathArgs: []string{instanceID},
		Body:     body,
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// ResizeInstance resizes instance.
//
// Change a STOPPED instance's flavor (vCPU/RAM). The new size is
// published to the desired spec now and materializes when the on-host
// compute-agent redefines the domain on the next start — so resize the
// instance, then start it. The instance must be stopped (libvirt can't
// change a running domain's max vcpus/memory in place).
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) ResizeInstance(ctx context.Context, instanceID string, body *ResizeInstanceRequest, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "resizeInstance",
		Method:   "POST",
		Path:     "/v1/instances/{instance_id}/resize",
		PathArgs: []string{instanceID},
		Body:     body,
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// StartInstance starts instance.
//
// Start a stopped instance.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) StartInstance(ctx context.Context, instanceID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "startInstance",
		Method:   "POST",
		Path:     "/v1/instances/{instance_id}/start",
		PathArgs: []string{instanceID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// StartSerialConsole opens an interactive serial console.
//
// Upgrades to a WebSocket carrying an interactive session on the
// instance's serial port — the equivalent of a crash cart, and the
// only way in when the network is broken: a bad fstab, a wrong kernel, a
// security group that locked you out, a full disk.
//
// Raw bytes in binary frames, both directions. It is a terminal, not a
// protocol: send keystrokes, receive whatever the guest prints. Point a
// terminal emulator at it.
//
// This drops you at the guest's OWN login prompt. It is not a backdoor
// — the guest's credentials are still required and nothing here grants
// access past what the guest itself allows.
//
// Authentication is on the upgrade request, signed or
// cookie-authenticated like any other call, so no separate token step is
// needed.
//
// One session per instance: opening a second disconnects the first,
// rather than interleaving two people's keystrokes into the same
// terminal. A session ends after 15 minutes idle, or 4 hours regardless,
// and the close frame carries the reason.
//
// Requires the instance to be running.
func (c *Client) StartSerialConsole(ctx context.Context, instanceID string, params *StartSerialConsoleParams, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "startSerialConsole",
		Method:   "GET",
		Path:     "/v1/instances/{instance_id}/console/serial",
		PathArgs: []string{instanceID},
	}
	op.Query = params.query()
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// StopInstance stops instance.
//
// Stop a running instance.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) StopInstance(ctx context.Context, instanceID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "stopInstance",
		Method:   "POST",
		Path:     "/v1/instances/{instance_id}/stop",
		PathArgs: []string{instanceID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// UpdateImage updates an image's metadata.
func (c *Client) UpdateImage(ctx context.Context, imageID string, body *ImageUpdateRequest, opts ...basaltic.RequestOption) (*Image, error) {
	op := &basaltic.Operation{
		ID:       "updateImage",
		Method:   "PATCH",
		Path:     "/v1/images/{image_id}",
		PathArgs: []string{imageID},
		Body:     body,
	}
	var out struct {
		Image *Image `json:"image"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Image, nil
}

// UpdateInstance updates instance.
//
// Update an instance's name, description, or metadata.
func (c *Client) UpdateInstance(ctx context.Context, instanceID string, body *InstanceUpdateRequest, opts ...basaltic.RequestOption) (*Instance, error) {
	op := &basaltic.Operation{
		ID:       "updateInstance",
		Method:   "PATCH",
		Path:     "/v1/instances/{instance_id}",
		PathArgs: []string{instanceID},
		Body:     body,
	}
	var out struct {
		Instance *Instance `json:"instance"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Instance, nil
}

// UpdateInstancePool updates an instance pool's size, tags or launch template.
//
// Change desired_count (bounded by the pool's min_count/max_count), the
// pool's own tags, and/or the launch template. The reconciler scales the
// live instance set to match a new desired_count. min_count and
// max_count are fixed at create.
//
// `tags` relabels the POOL and nothing else: it takes effect
// immediately, no instance is touched, and the new set is what a later
// IAM condition reads as `basalt:ResourceTag/<key>`. It replaces the
// whole set — an empty object clears it, an omitted field leaves it
// alone.
//
// A new `template` replaces the stored one wholesale and changes what
// the pool launches NEXT; the instances already running keep what they
// booted with, because a live VM cannot change flavor, tier, subnet or
// tags in place. So a `template.tags` edit leaves the pool holding
// members with two different tag sets until it is rolled. The pool
// reports `stale_instance_count` — how many members are on the old
// template — and POST /v1/instance-pools/{pool_id}/refresh rolls them.
func (c *Client) UpdateInstancePool(ctx context.Context, poolID string, body *InstancePoolUpdateRequest, opts ...basaltic.RequestOption) (*InstancePool, error) {
	op := &basaltic.Operation{
		ID:       "updateInstancePool",
		Method:   "PATCH",
		Path:     "/v1/instance-pools/{pool_id}",
		PathArgs: []string{poolID},
		Body:     body,
	}
	var out struct {
		InstancePool *InstancePool `json:"instance_pool"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.InstancePool, nil
}

// UpdateInstanceVolumeAttachment updates a volume attachment's settings.
//
// Mutates binding-level settings — today just delete_on_termination,
// which controls whether the volume is destroyed with the instance.
func (c *Client) UpdateInstanceVolumeAttachment(ctx context.Context, instanceID string, volumeID string, body *UpdateInstanceVolumeAttachmentRequest, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "updateInstanceVolumeAttachment",
		Method:   "PATCH",
		Path:     "/v1/instances/{instance_id}/volumes/{volume_id}",
		PathArgs: []string{instanceID, volumeID},
		Body:     body,
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}
