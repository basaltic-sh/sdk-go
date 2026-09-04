// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package dns

import (
	"time"
)

type Record struct {
	// CRN Cloud Resource Name.
	CRN string `json:"crn,omitempty"`
	ID  string `json:"id,omitempty"`

	// Managed true for platform-managed records (SOA, apex NS, and the DNSSEC
	// set). Managed records are read-only — they cannot be updated or
	// deleted through the API.
	Managed bool `json:"managed,omitempty"`

	// Name record name (FQDN).
	Name string `json:"name,omitempty"`
	TTL  int    `json:"ttl,omitempty"`

	// Type record type. Customer records use one of the creatable RecordType
	// values; reads also surface the platform-managed apex/DNSSEC types
	// (SOA, NS, DNSKEY, DS, NSEC, RRSIG).
	Type   string         `json:"type,omitempty"`
	Values []*RecordValue `json:"values,omitempty"`
	ZoneID string         `json:"zone_id,omitempty"`
}

type RecordCreateRequest struct {
	// Name record name (FQDN).
	//
	// Required.
	Name string `json:"name"`
	TTL  *int   `json:"ttl,omitempty"`

	// Required.
	Type RecordType `json:"type"`

	// Required.
	Values []*RecordValue `json:"values"`
}

// RecordType curated subset of supported DNS record types. SOA and DNSSEC-managed
// types (DNSKEY, DS, NSEC*, RRSIG, …) are managed by the platform and
// not creatable through the API.
//
// Every type here has a wire representation. ALIAS is not offered: it
// resolves only if the authoritative server flattens it to A/AAAA, and
// ours does not. Point a zone apex at a target with A/AAAA records
// carrying its addresses.
type RecordType string

// Values RecordType accepts.
const (
	RecordTypeA          RecordType = "A"
	RecordTypeAaaa       RecordType = "AAAA"
	RecordTypeAfsdb      RecordType = "AFSDB"
	RecordTypeApl        RecordType = "APL"
	RecordTypeCaa        RecordType = "CAA"
	RecordTypeCert       RecordType = "CERT"
	RecordTypeCname      RecordType = "CNAME"
	RecordTypeCsync      RecordType = "CSYNC"
	RecordTypeDhcid      RecordType = "DHCID"
	RecordTypeDname      RecordType = "DNAME"
	RecordTypeEui48      RecordType = "EUI48"
	RecordTypeEui64      RecordType = "EUI64"
	RecordTypeHinfo      RecordType = "HINFO"
	RecordTypeHTTPS      RecordType = "HTTPS"
	RecordTypeIpseckey   RecordType = "IPSECKEY"
	RecordTypeKx         RecordType = "KX"
	RecordTypeL32        RecordType = "L32"
	RecordTypeL64        RecordType = "L64"
	RecordTypeLoc        RecordType = "LOC"
	RecordTypeLp         RecordType = "LP"
	RecordTypeMx         RecordType = "MX"
	RecordTypeNaptr      RecordType = "NAPTR"
	RecordTypeNid        RecordType = "NID"
	RecordTypeNs         RecordType = "NS"
	RecordTypeOpenpgpkey RecordType = "OPENPGPKEY"
	RecordTypePtr        RecordType = "PTR"
	RecordTypeRkey       RecordType = "RKEY"
	RecordTypeRp         RecordType = "RP"
	RecordTypeSmimea     RecordType = "SMIMEA"
	RecordTypeSpf        RecordType = "SPF"
	RecordTypeSrv        RecordType = "SRV"
	RecordTypeSshfp      RecordType = "SSHFP"
	RecordTypeSvcb       RecordType = "SVCB"
	RecordTypeTlsa       RecordType = "TLSA"
	RecordTypeTxt        RecordType = "TXT"
	RecordTypeURI        RecordType = "URI"
)

type RecordUpdateRequest struct {
	TTL    *int           `json:"ttl,omitempty"`
	Values []*RecordValue `json:"values,omitempty"`
}

type RecordValue struct {
	// Content RDATA — wire representation per record type.
	Content string `json:"content"`

	// Disabled must be `false`. `true` is **refused**.
	//
	// There is nowhere to keep a value that is not served, so a disabled
	// value used to be accepted, echoed back in the response, and then
	// dropped — the staged value was gone by the next read, with the
	// write having reported success. Refusing is the honest version.
	//
	// The field remains on the schema because the console and CLI send it
	// on every value; only `true` is rejected. To take a value out of an
	// RRset, remove it from `values`.
	Disabled bool `json:"disabled,omitempty"`
}

type Soa struct {
	// AdminEmail SOA `rname` — the platform's DNS-operations contact. Stamped
	// server-side; not customer-configurable.
	AdminEmail string `json:"admin_email,omitempty"`
	Expire     int    `json:"expire,omitempty"`

	// Minimum NXDOMAIN cache TTL.
	Minimum int `json:"minimum,omitempty"`

	// PrimaryNs SOA `mname` — first authoritative nameserver for the zone.
	PrimaryNs string `json:"primary_ns,omitempty"`
	Refresh   int    `json:"refresh,omitempty"`
	Retry     int    `json:"retry,omitempty"`
}

type Tags = map[string]string

type VPCAssociationAccepted struct {
	VPCID  string `json:"vpc_id"`
	ZoneID string `json:"zone_id"`
}

type VPCAssociationRequest struct {
	// VPCID VPC to associate with this private zone.
	//
	// Required.
	VPCID string `json:"vpc_id"`
}

type Zone struct {
	// CRN Cloud Resource Name.
	CRN string `json:"crn,omitempty"`

	// Description free text set at creation. Read-only after that — `tags` are the
	// mutable metadata and the only thing `PATCH` accepts.
	Description string      `json:"description,omitempty"`
	DNSSEC      *ZoneDNSSEC `json:"dnssec,omitempty"`
	ID          string      `json:"id,omitempty"`

	// Name Zone FQDN.
	Name string `json:"name,omitempty"`

	// Nameservers authoritative nameservers for the zone — the apex NS set. Copy
	// this list verbatim into your registrar's nameserver configuration to
	// delegate the zone to the platform.
	//
	// These names are unique to THIS zone: each carries a per-zone label,
	// which is what makes the delegation double as the ownership proof
	// (see `ownership`). Two zones for the same domain get different
	// names, and the one the registrar points at is the one that serves.
	// Use them exactly as written — Basaltic's bare nameserver names
	// will not verify the zone.
	Nameservers []string       `json:"nameservers,omitempty"`
	Ownership   *ZoneOwnership `json:"ownership,omitempty"`
	Soa         *Soa           `json:"soa,omitempty"`
	Tags        Tags           `json:"tags,omitempty"`

	// Visibility `public` zones answer on the internet-facing nameservers; `private`
	// zones answer only inside the VPCs associated with them (see the
	// vpc-associations endpoints). Fixed at creation.
	//
	// One of: "public", "private".
	Visibility string `json:"visibility,omitempty"`
}

type ZoneCreateRequest struct {
	// Description free-form note stored with the zone. Not echoed back on the Zone
	// object.
	Description *string `json:"description,omitempty"`

	// DNSSEC sign the zone with DNSSEC. On unless you say otherwise, and almost
	// every zone should leave it on.
	//
	// **Turn it off only if this domain is served by another DNS provider
	// at the same time as us.** A signed zone puts our DS record at the
	// parent, and that DS covers only the answers WE sign — so a
	// validating resolver that happens to ask the other provider gets a
	// signature it cannot verify and fails the lookup. Roughly half your
	// queries, unpredictably, which is worse than either provider on its
	// own. Unsigned is the only configuration that works for that setup
	// today.
	//
	// Fixed at creation. Turning signing off later breaks the domain until
	// the DS is withdrawn at the registrar and that withdrawal has
	// propagated, which is a sequence this API cannot drive for you.
	DNSSEC *bool `json:"dnssec,omitempty"`

	// ImportExistingRecords read the domain's records from the nameservers that serve it TODAY
	// and copy them into this zone, before you move the delegation here.
	//
	// Worth asking for when you are migrating a live domain. The
	// delegation is the ownership proof, so the moment you point your
	// registrar at this zone is the moment we start answering for it —
	// and an empty zone answers with nothing, which takes the site and the
	// mail down until you have retyped everything.
	//
	// Runs in the background; the zone is created immediately. Poll GET
	// /v1/zones/{zone_id}/record-import for the outcome.
	//
	// Best effort, and the result says how good it was. A zone transfer is
	// exhaustive and almost always refused; the fallback queries a list of
	// common names and cannot find a record it did not think to ask for.
	// Check `record_import.complete` before you switch your old provider
	// off.
	//
	// Records you have already created are never overwritten, and records
	// this platform manages itself — the SOA, the DNSSEC chain, the
	// zone's nameservers — are never imported.
	ImportExistingRecords *bool `json:"import_existing_records,omitempty"`

	// Name Zone FQDN.
	//
	// Required.
	Name string `json:"name"`
	Tags Tags   `json:"tags,omitempty"`

	// Visibility `private` restricts the zone to the VPCs named in `vpc_ids` and
	// requires at least one; `public` (the default) rejects `vpc_ids`
	// outright rather than ignoring them. Cannot be changed afterwards.
	//
	// One of: "public", "private".
	Visibility *string `json:"visibility,omitempty"`

	// VPCIDs VPCs the zone resolves in. Required when visibility=private,
	// rejected when visibility=public. More can be associated later via
	// POST /v1/zones/{zone_id}/vpc-associations.
	VPCIDs []string `json:"vpc_ids,omitempty"`
}

// ZoneDNSSEC DNSSEC signing state. Present once the signer has bootstrapped the
// zone, which is always-on for zones created through this API; absent on
// a zone that is not signed.
//
// `ds_records` is what the customer pastes at the parent registrar to
// complete the chain of trust. `rdata` is the zone-file form of the same
// record, for registrars that take one string.
type ZoneDNSSEC struct {
	// Algorithm DNSSEC algorithm number. 13 = ECDSA P-256 SHA-256.
	Algorithm int             `json:"algorithm,omitempty"`
	DsRecords []*ZoneDsRecord `json:"ds_records,omitempty"`
	Enabled   bool            `json:"enabled"`

	// KskKeyTag key tag of the key-signing key — matches the DS records below.
	KskKeyTag int `json:"ksk_key_tag,omitempty"`

	// ZskKeyTag key tag of the zone-signing key.
	ZskKeyTag int `json:"zsk_key_tag,omitempty"`
}

// ZoneDsRecord One DS record to publish at the parent zone.
type ZoneDsRecord struct {
	Algorithm int    `json:"algorithm"`
	Digest    string `json:"digest"`

	// DigestType 2 = SHA-256.
	DigestType int `json:"digest_type"`
	KeyTag     int `json:"key_tag"`

	// Rdata full zone-file form — `<key_tag> <algorithm> <digest_type>
	// <digest>`.
	Rdata string `json:"rdata"`
}

type ZoneImportRequest struct {
	// ZoneFile the zone file, as text. `$ORIGIN`, `$TTL`, `$GENERATE`, relative
	// names and parenthesised multi-line records are all honoured;
	// `$INCLUDE` is refused, because the path it names would be read on
	// our filesystem rather than yours.
	//
	// Required.
	ZoneFile string `json:"zone_file"`
}

// ZoneImportResult what the import did. `records_created` plus `records_replaced` is
// every RRset taken from the file; anything in `skipped` is still in the
// file but not in the zone.
type ZoneImportResult struct {
	// RecordsByType Imported RRset count per record type.
	RecordsByType map[string]int `json:"records_by_type,omitempty"`

	// RecordsCreated RRsets in the file that the zone did not already have.
	RecordsCreated int `json:"records_created,omitempty"`

	// RecordsReplaced RRsets that existed at the same name and type and were replaced
	// wholesale by the file's values.
	RecordsReplaced int                  `json:"records_replaced,omitempty"`
	Skipped         []*ZoneImportSkipped `json:"skipped,omitempty"`

	// Warnings records that were imported, but not exactly as written — an RRset
	// the file gave more than one TTL, for instance.
	Warnings []string `json:"warnings,omitempty"`
}

// ZoneImportSkipped One RRset in the file that was not imported, and why.
type ZoneImportSkipped struct {
	Name   string `json:"name,omitempty"`
	Reason string `json:"reason,omitempty"`
	Type   string `json:"type,omitempty"`
}

// ZoneOwnership domain-ownership proof. A zone with no existing zone above it on the
// platform does not resolve until `verified` is true — set the zone's
// `nameservers` at your registrar, then POST
// /v1/zones/{zone_id}/verify-ownership.
//
// There is no record to publish. The zone's nameservers carry a label
// unique to it, so delegating the domain to them is itself the proof:
// only whoever controls the domain at its registrar can do it, and the
// names cannot be inherited from a zone that used to hold this domain.
//
// Use the zone's own `nameservers` and nothing else. Basaltic's bare
// nameserver names prove nothing on their own — anyone can point a
// domain at those — so a delegation naming them leaves the zone
// unverified.
//
// Other providers' nameservers may sit alongside yours if you host the
// domain at two providers. What does NOT work is delegating to two
// Basaltic zones at once: while a domain names another zone's
// nameservers as well as this one's, neither is served, and the remedy
// is to remove the ones this zone does not list. That is the state to
// expect midway through moving a domain between accounts.
//
// A zone created underneath a zone you already own inherits its parent's
// proof and is verified on creation.
//
// The proof is a lease, not a deed: it is re-confirmed periodically for
// as long as the zone is served, so a domain that lapses or changes
// hands stops being served from here. `checked_at` is the last
// confirmation; `recheck_deadline` appears only while re-confirmation is
// failing and is the field to render a warning from.
type ZoneOwnership struct {
	// CheckedAt when the proof was last re-confirmed. `verified_at` keeps meaning
	// "first demonstrated" and does not move; this does, on every pass
	// that passes. Absent until the zone has been re-confirmed at least
	// once.
	CheckedAt time.Time `json:"checked_at,omitempty"`

	// RecheckDeadline Present ONLY while the periodic re-proof is failing: the instant the
	// zone stops answering unless it passes again. Absent means healthy.
	//
	// Each pass re-checks a failing zone, so reaching this date takes
	// sustained failure, not one bad afternoon — a transient resolver
	// problem clears itself on the next pass. To clear it deliberately,
	// point the domain's delegation back at this zone's `nameservers` and
	// POST /v1/zones/{zone_id}/verify-ownership.
	//
	// The organization's owner is emailed when this date is set, and again
	// if the zone does stop resolving. A zone is never taken off the air
	// before that message has gone out, so the date here is the earliest
	// the zone can stop answering and never the only warning.
	RecheckDeadline time.Time `json:"recheck_deadline,omitempty"`

	// Verified whether the zone has proved ownership. Unverified zones do not
	// resolve.
	Verified bool `json:"verified,omitempty"`

	// VerifiedAt when ownership was first proved. Absent while unverified.
	VerifiedAt time.Time `json:"verified_at,omitempty"`
}

// ZoneRecordImport what came of `import_existing_records` — reading the domain's
// records off the provider it is moving away from.
//
// Its own resource, at GET /v1/zones/{zone_id}/record-import, because it
// describes one event that happened once when the zone was created and
// never changes afterwards. 404 when the zone was created without asking
// for an import, which is the ordinary case and not an error.
//
// A failed import is never a failed zone. The zone exists, resolves once
// delegated, and can be filled in by hand or from a zone file; only the
// convenience did not happen.
type ZoneRecordImport struct {
	// Complete true for the two sources that enumerate the zone — `axfr` and
	// `nsec-walk` — and false for `query`.
	//
	// **This is the field to read before switching your old provider
	// off.** A false here means records may exist that we did not find,
	// not that none do.
	Complete bool `json:"complete,omitempty"`

	// Error present only when state is `failed`, and says what could not be
	// done.
	Error string `json:"error,omitempty"`

	// Found record sets the scan turned up.
	Found int `json:"found,omitempty"`

	// Imported record sets actually written. Lower than `found` for records you had
	// already created — yours win — and for the ones this platform
	// manages itself.
	Imported int `json:"imported,omitempty"`

	// Notes what could not be established, and what was deliberately not
	// imported.
	Notes []string `json:"notes,omitempty"`

	// Source how the records were found, in descending order of how much the
	// result is worth.
	//
	// `axfr` is a zone transfer: the whole zone, exactly. `nsec-walk`
	// follows the zone's own DNSSEC NSEC chain, which names every record
	// set in it — also exact, and available on signed zones whose
	// provider refuses transfers. `query` is a list of common names, which
	// finds what it thought to ask for and cannot know what it missed.
	//
	// Absent while pending.
	//
	// One of: "axfr", "nsec-walk", "query".
	Source string `json:"source,omitempty"`

	// State `pending` while the background job runs. `complete` means the scan
	// ran and what it found was applied — not that everything the domain
	// has is now here; see `complete`.
	//
	// `pending` is bounded. A scan is capped well below the point at which
	// this stops reporting it, so a job that dies without recording an
	// outcome is reported as `failed` rather than staying `pending` for
	// the life of the zone. There is no state that means "still importing"
	// after that bound, and nothing you can poll for that would ever
	// change.
	//
	// One of: "pending", "complete", "failed".
	State     string    `json:"state,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// ZoneUpdateRequest tags are the only mutable field on a zone: its name, visibility and
// DNSSEC posture are fixed when it is created, and its records are their
// own sub-resource.
type ZoneUpdateRequest struct {
	Tags map[string]string `json:"tags,omitempty"`
}
