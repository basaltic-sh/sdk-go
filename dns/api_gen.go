// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package dns

import (
	"context"
	"io"
	"iter"
	"net/url"
	"strconv"

	basaltic "github.com/basaltic-sh/sdk-go"
)

// ListRecordsParams are the optional filters and pagination controls for
// [Client.ListRecords]. A nil *ListRecordsParams sends none of them.
type ListRecordsParams struct {
	// IncludeManaged include the platform-stamped rows (SOA, apex NS, and the DNSSEC set)
	// alongside your own. Off by default — they cannot be edited or
	// deleted, and on a signed zone there are more of them than there are
	// of yours.
	IncludeManaged *bool
	Limit          int

	// Marker resume token — the last record id from the previous page.
	Marker string

	// Name substring match on the record name.
	Name string

	// Type exact record type to filter by (e.g. `A`, `MX`).
	Type string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListRecordsParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.IncludeManaged != nil {
		q.Set("include_managed", strconv.FormatBool(*p.IncludeManaged))
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
	if p.Type != "" {
		q.Set("type", p.Type)
	}
	return q
}

// withMarker copies p with the pagination cursor replaced, leaving the
// caller's value untouched across pages.
func (p *ListRecordsParams) withMarker(marker string) *ListRecordsParams {
	var out ListRecordsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListZonesParams are the optional filters and pagination controls for
// [Client.ListZones]. A nil *ListZonesParams sends none of them.
type ListZonesParams struct {
	Limit int

	// Marker resume token — the last zone id from the previous page.
	Marker string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListZonesParams) query() url.Values {
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
func (p *ListZonesParams) withMarker(marker string) *ListZonesParams {
	var out ListZonesParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// AssociateZoneVPC associates a VPC with a private zone.
//
// Refused on public zones.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) AssociateZoneVPC(ctx context.Context, zoneID string, body *VPCAssociationRequest, opts ...basaltic.RequestOption) (*VPCAssociationAccepted, error) {
	op := &basaltic.Operation{
		ID:       "associateZoneVPC",
		Method:   "POST",
		Path:     "/v1/zones/{zone_id}/vpc-associations",
		PathArgs: []string{zoneID},
		Body:     body,
	}
	var out VPCAssociationAccepted
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// CreateRecord creates record.
//
// Create a new record in the zone.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateRecord(ctx context.Context, zoneID string, body *RecordCreateRequest, opts ...basaltic.RequestOption) (*Record, error) {
	op := &basaltic.Operation{
		ID:       "createRecord",
		Method:   "POST",
		Path:     "/v1/zones/{zone_id}/records",
		PathArgs: []string{zoneID},
		Body:     body,
	}
	var out struct {
		Record *Record `json:"record"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Record, nil
}

// CreateZone creates zone.
//
// Create a new DNS zone. The SOA and NS records are generated
// automatically against the platform nameserver list; the zone is
// queryable immediately on success.
//
// Defaults to a public zone. Pass `visibility: private` plus at least
// one `vpc_ids` entry for a zone that resolves only inside those VPCs
// — a private zone with no VPC, or a public zone carrying `vpc_ids`,
// is rejected with 400 rather than silently coerced.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateZone(ctx context.Context, body *ZoneCreateRequest, opts ...basaltic.RequestOption) (*Zone, error) {
	op := &basaltic.Operation{
		ID:     "createZone",
		Method: "POST",
		Path:   "/v1/zones",
		Body:   body,
	}
	var out struct {
		Zone *Zone `json:"zone"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Zone, nil
}

// DeleteRecord deletes record.
//
// Delete a record.
func (c *Client) DeleteRecord(ctx context.Context, zoneID string, recordID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteRecord",
		Method:   "DELETE",
		Path:     "/v1/zones/{zone_id}/records/{record_id}",
		PathArgs: []string{zoneID, recordID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteZone deletes zone.
//
// Start deleting a DNS zone and every record it owns.
//
// **Accepted, not done.** The teardown runs in the background: it frees
// the zone's DNSSEC key material, releases the delegating DS at its
// parent, deletes the records, and drops the zone itself last. Nothing
// is left half-removed — but the zone still exists when this call
// returns.
//
// Poll `GET /v1/zones/{zone_id}` until it answers 404. That matters if
// you intend to re-create the same name: until the teardown finishes,
// creating it again answers 409, because the old zone is still there.
//
// Idempotent. Deleting a zone whose teardown is already running is
// accepted and does not start a second one.
func (c *Client) DeleteZone(ctx context.Context, zoneID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteZone",
		Method:   "DELETE",
		Path:     "/v1/zones/{zone_id}",
		PathArgs: []string{zoneID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteZoneRecordImport — Discard the record-import outcome.
//
// Forget what came of `import_existing_records`. **The imported records
// stay** — by the time you read this they are ordinary records in the
// zone, and this only discards the note about where they came from.
//
// The outcome describes something that happened once and never changes,
// so the console keeps showing it — including the warning that a
// probe-based import may not have found everything. That warning is
// worth reading once and is noise afterwards, so this is how you put it
// away.
//
// Answers `204` whether or not there was an outcome to discard: the
// request asks for a state, and that state is reached either way.
func (c *Client) DeleteZoneRecordImport(ctx context.Context, zoneID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteZoneRecordImport",
		Method:   "DELETE",
		Path:     "/v1/zones/{zone_id}/record-import",
		PathArgs: []string{zoneID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DissociateZoneVPC dissociates a VPC from a private zone.
func (c *Client) DissociateZoneVPC(ctx context.Context, zoneID string, vpcID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "dissociateZoneVPC",
		Method:   "DELETE",
		Path:     "/v1/zones/{zone_id}/vpc-associations/{vpc_id}",
		PathArgs: []string{zoneID, vpcID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// ExportZoneFile exports the zone as a zone file.
//
// Returns the zone as an RFC 1035 master file — the BIND format every
// provider reads. The inverse of `/import`.
//
// Worth keeping before a bulk change: the import **merges and never
// deletes**, so without a file taken beforehand there is no way back to
// the previous state.
//
// **The file cannot describe the whole zone**, and says so in its own
// header: private-zone VPC associations have no zone-file
// representation. Re-importing restores the records; those have to be
// set up again.
//
// The DNSSEC records are omitted — they sign keys held by this
// platform and mean nothing elsewhere. The SOA and apex NS are included,
// as in any zone file; an importer replaces them with its own.
//
// The caller must close the returned reader.
func (c *Client) ExportZoneFile(ctx context.Context, zoneID string, opts ...basaltic.RequestOption) (io.ReadCloser, error) {
	op := &basaltic.Operation{
		ID:       "exportZoneFile",
		Method:   "GET",
		Path:     "/v1/zones/{zone_id}/export",
		PathArgs: []string{zoneID},
	}
	stream, _, err := c.rt.DoStream(ctx, op, opts...)
	if err != nil {
		return nil, err
	}
	return stream, nil
}

// GetRecord gets record.
//
// Get a specific record by ID.
func (c *Client) GetRecord(ctx context.Context, zoneID string, recordID string, opts ...basaltic.RequestOption) (*Record, error) {
	op := &basaltic.Operation{
		ID:       "getRecord",
		Method:   "GET",
		Path:     "/v1/zones/{zone_id}/records/{record_id}",
		PathArgs: []string{zoneID, recordID},
	}
	var out struct {
		Record *Record `json:"record"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Record, nil
}

// GetZone gets zone.
//
// Get details of a specific DNS zone.
func (c *Client) GetZone(ctx context.Context, zoneID string, opts ...basaltic.RequestOption) (*Zone, error) {
	op := &basaltic.Operation{
		ID:       "getZone",
		Method:   "GET",
		Path:     "/v1/zones/{zone_id}",
		PathArgs: []string{zoneID},
	}
	var out struct {
		Zone *Zone `json:"zone"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Zone, nil
}

// GetZoneRecordImport gets the record-import outcome.
//
// What came of `import_existing_records` — reading the domain's
// records off the provider it was moving away from.
//
// Its own resource rather than a field on the zone, because it describes
// one event that happened once when the zone was created and never
// changes.
//
// Poll this after creating a zone with `import_existing_records: true`:
// `state` is `pending` while the background job runs. **Read `complete`
// before you switch your old provider off** — see the schema.
func (c *Client) GetZoneRecordImport(ctx context.Context, zoneID string, opts ...basaltic.RequestOption) (*ZoneRecordImport, error) {
	op := &basaltic.Operation{
		ID:       "getZoneRecordImport",
		Method:   "GET",
		Path:     "/v1/zones/{zone_id}/record-import",
		PathArgs: []string{zoneID},
	}
	var out struct {
		RecordImport *ZoneRecordImport `json:"record_import"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.RecordImport, nil
}

// ImportZoneFile imports a zone file.
//
// Load an RFC 1035 master file — a BIND zone file, the format every
// DNS provider exports — into an existing zone.
//
// **Merges, never replaces.** An RRset in the file replaces the RRset at
// the same name and type wholesale; an RRset the file does not mention
// is left exactly as it is. So an import cannot delete a record, and
// importing the same file twice is a no-op rather than an error —
// which is what makes it safe to retry.
//
// **All or nothing.** The whole file is validated before any record is
// written: every record has to pass the same type, name, TTL and rdata
// checks `createRecord` applies, the same CNAME-coexistence rule, and
// the same per-zone and per-org record quota. One bad record refuses the
// file and the zone is untouched.
//
// **What is skipped rather than imported.** The SOA and the apex NS set
// are the platform's: the zone is served by our nameservers, and the
// apex NS is what the parent delegates to. The DNSSEC records (`RRSIG`,
// `DNSKEY`, `DS`, `NSEC`, `NSEC3`, `NSEC3PARAM`, `CDS`, `CDNSKEY`) are
// generated by our signer from keys your file cannot know. Records the
// file places outside this zone are skipped too. All of them come back
// in `skipped` with a reason — a file exported from a signed zone
// carries every one, so refusing the file over them would make
// export/import unusable.
//
// The file may be sent as JSON (`zone_file`) or as the raw request body
// under any other content type, so `curl --data-binary @db.example.com`
// works directly. Maximum 1 MiB.
//
// On success the zone's SOA serial advances once for the whole import
// and the change propagates to the authoritative nameservers.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) ImportZoneFile(ctx context.Context, zoneID string, body *ZoneImportRequest, opts ...basaltic.RequestOption) (*ZoneImportResult, error) {
	op := &basaltic.Operation{
		ID:       "importZoneFile",
		Method:   "POST",
		Path:     "/v1/zones/{zone_id}/import",
		PathArgs: []string{zoneID},
		Body:     body,
	}
	var out struct {
		Import *ZoneImportResult `json:"import"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Import, nil
}

// ListRecords lists records.
//
// List records in a DNS zone, newest-RRset-first. Supports filtering by
// `type` and a substring match on `name`. Keyset- paginated by record id
// (UUIDv7) — pass the last id from the previous page as `marker`.
//
// **Returns your records only.** The rows the platform stamps for itself
// — the SOA, the apex NS, and the DNSSEC set (`DNSKEY`, `DS`,
// `NSEC`/`NSEC3`, `NSEC3PARAM`, `RRSIG`, `CDS`, `CDNSKEY`) — are left
// out unless you ask for them. On a signed zone they outnumber customer
// records, none of them can be edited or deleted, and the DNSSEC
// material you might actually want is on the zone itself under `dnssec`.
//
// Pass `include_managed=true` to get the zone exactly as it is served,
// which is what a zone-diffing tool wants.
//
// Returns one page. Use ListRecordsAll to walk every page.
func (c *Client) ListRecords(ctx context.Context, zoneID string, params *ListRecordsParams, opts ...basaltic.RequestOption) (*basaltic.Page[Record], error) {
	op := &basaltic.Operation{
		ID:       "listRecords",
		Method:   "GET",
		Path:     "/v1/zones/{zone_id}/records",
		PathArgs: []string{zoneID},
	}
	op.Query = params.query()
	var out struct {
		Items []Record `json:"records"`
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
	page := &basaltic.Page[Record]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListRecordsAll walks every page of ListRecords, yielding one item at a
// time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListRecordsAll(ctx, zoneID, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListRecordsAll(ctx context.Context, zoneID string, params *ListRecordsParams, opts ...basaltic.RequestOption) iter.Seq2[Record, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Record], error) {
		return c.ListRecords(ctx, zoneID, params.withMarker(marker), opts...)
	})
}

// ListZoneVPCAssociations lists VPC associations.
//
// Returns VPC ids attached to a private zone. Empty for public zones or
// unattached private zones.
func (c *Client) ListZoneVPCAssociations(ctx context.Context, zoneID string, opts ...basaltic.RequestOption) (*basaltic.Page[string], error) {
	op := &basaltic.Operation{
		ID:       "listZoneVPCAssociations",
		Method:   "GET",
		Path:     "/v1/zones/{zone_id}/vpc-associations",
		PathArgs: []string{zoneID},
	}
	var out struct {
		Items []string `json:"vpc_ids"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[string]{Items: out.Items}
	return page, nil
}

// ListZones lists zones.
//
// List DNS zones owned by the current organization, newest first.
// Keyset-paginated by zone id (UUIDv7 sorts by creation time) — pass
// the last id from the previous page as `marker` to fetch the next.
//
// Returns one page. Use ListZonesAll to walk every page.
func (c *Client) ListZones(ctx context.Context, params *ListZonesParams, opts ...basaltic.RequestOption) (*basaltic.Page[Zone], error) {
	op := &basaltic.Operation{
		ID:     "listZones",
		Method: "GET",
		Path:   "/v1/zones",
	}
	op.Query = params.query()
	var out struct {
		Items []Zone `json:"zones"`
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
	page := &basaltic.Page[Zone]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListZonesAll walks every page of ListZones, yielding one item at a
// time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListZonesAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListZonesAll(ctx context.Context, params *ListZonesParams, opts ...basaltic.RequestOption) iter.Seq2[Zone, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Zone], error) {
		return c.ListZones(ctx, params.withMarker(marker), opts...)
	})
}

// UpdateRecord updates record.
//
// Update a record's TTL or value set. The value list is replaced
// wholesale.
func (c *Client) UpdateRecord(ctx context.Context, zoneID string, recordID string, body *RecordUpdateRequest, opts ...basaltic.RequestOption) (*Record, error) {
	op := &basaltic.Operation{
		ID:       "updateRecord",
		Method:   "PATCH",
		Path:     "/v1/zones/{zone_id}/records/{record_id}",
		PathArgs: []string{zoneID, recordID},
		Body:     body,
	}
	var out struct {
		Record *Record `json:"record"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Record, nil
}

// UpdateZone updates zone.
//
// Updates a zone's tags. Omitted tags are left alone; a tags map
// replaces the zone's tag set, so an empty object clears it.
//
// Tags gate policy as well as label the zone, so this is authorized
// against both the tags being requested and the tags the zone already
// carries — a principal fenced to one of them is refused by the other.
func (c *Client) UpdateZone(ctx context.Context, zoneID string, body *ZoneUpdateRequest, opts ...basaltic.RequestOption) (*Zone, error) {
	op := &basaltic.Operation{
		ID:       "updateZone",
		Method:   "PATCH",
		Path:     "/v1/zones/{zone_id}",
		PathArgs: []string{zoneID},
		Body:     body,
	}
	var out struct {
		Zone *Zone `json:"zone"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Zone, nil
}

// VerifyZoneOwnership verifies zone ownership.
//
// Check the domain's delegation and, if it names this zone, start
// serving it.
//
// A zone created with no existing zone above it on the platform does not
// resolve until ownership is proved. Set the zone's `nameservers` at
// your registrar, then call this. There is no record to publish: those
// names carry a label unique to this zone, so delegating to them
// requires control of the domain at its registrar — which is what
// stops someone claiming a subdomain of a domain they do not own and
// sitting in your subtree, and what stops an abandoned delegation from
// being redeemed by whoever next creates the name.
//
// A zone created underneath a zone you already own inherits its proof
// and is served immediately; calling this on one returns success with
// nothing to do. Idempotent, so a client may poll it.
//
// This is also the call that clears a failing re-proof. Ownership is
// re-confirmed periodically, and a zone whose re-confirmation is failing
// carries an `ownership.recheck_deadline`; on such a zone this runs the
// check for real and clears the deadline when it passes. So unlike the
// already-verified case above, a verified zone can still answer 400
// here.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) VerifyZoneOwnership(ctx context.Context, zoneID string, opts ...basaltic.RequestOption) (*Zone, error) {
	op := &basaltic.Operation{
		ID:       "verifyZoneOwnership",
		Method:   "POST",
		Path:     "/v1/zones/{zone_id}/verify-ownership",
		PathArgs: []string{zoneID},
	}
	var out struct {
		Zone *Zone `json:"zone"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Zone, nil
}
