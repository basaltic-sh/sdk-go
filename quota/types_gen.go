// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package quota

type QuotaItem struct {
	// Available limit - in_use - reserved, floored at 0; -1 if unlimited.
	Available   int    `json:"available"`
	Description string `json:"description"`

	// InUse resources currently consuming this quota. Zero for per_resource
	// quotas (no running counter — the cap applies per parent resource).
	InUse int `json:"in_use"`

	// IsDefault true if the limit comes from the system default; false if there is a
	// per-org override.
	IsDefault bool `json:"is_default"`

	// Limit effective limit (-1 = unlimited). Override if set, otherwise
	// default.
	Limit int `json:"limit"`

	// Reserved resources temporarily reserved during async creation. Zero for
	// per_resource quotas.
	Reserved int `json:"reserved"`

	// ResourceType the resource the cap applies to. See migrations/quota/001_quotas.sql
	// for the seeded list (instances, vcpus, ram_mb, volumes, …,
	// domains_per_certificate, records_per_zone, …).
	ResourceType string `json:"resource_type"`

	// One of: "regional", "global", "per_resource".
	Scope string `json:"scope"`

	// Service the service that owns the quota. Together with resource_type it
	// identifies the quota — resource_type alone is not unique (e.g.
	// both compute and database have an `instances` quota).
	Service string `json:"service"`
}

type QuotaListResponse struct {
	// Policies non-reservation policies for the organization in the requested
	// region. Currently contains the four Email size and recipient
	// policies. Empty when region is omitted. Older servers may omit this
	// additive field.
	Policies []*QuotaPolicy `json:"policies,omitempty"`
	Quotas   []*QuotaItem   `json:"quotas"`
}

// QuotaPolicy an organization policy enforced in the specified region, shared across
// its accounts. Size applies per complete encoded sending message;
// recipient rates count accepted recipients over a rolling window, not
// API calls or a calendar day. This object deliberately has no in_use,
// reserved or available counters.
type QuotaPolicy struct {
	Description string `json:"description"`

	// Limit effective policy value after regional and organization overrides; -1
	// means unlimited by this policy and 0 permits none. The separate
	// Email message ceiling of 52428800 bytes (50 MiB) applies even when
	// the size policy is unlimited or larger than the ceiling.
	Limit int `json:"limit"`

	// LimitSource platform default, organization-wide override, or override for the
	// requested region.
	//
	// One of: "default", "organization", "region".
	LimitSource string `json:"limit_source"`
	Region      string `json:"region"`

	// One of: "message_size_bytes", "sent_recipients_per_second", "sent_recipients_per_day", "received_recipients_per_second".
	ResourceType string `json:"resource_type"`

	// Scope the policy applies to the caller's organization in region, even if
	// its value comes from an organization-wide override.
	//
	// One of: "regional".
	Scope string `json:"scope"`

	// One of: "email".
	Service string `json:"service"`

	// Unit bytes per complete encoded sending message, or accepted recipients
	// per rolling window.
	//
	// One of: "bytes", "recipients".
	Unit string `json:"unit"`

	// WindowSeconds rolling window length for recipient policies. Omitted for message
	// size; 86400 means rolling 24 hours.
	//
	// One of: "1", "86400".
	WindowSeconds int `json:"window_seconds,omitempty"`
}
