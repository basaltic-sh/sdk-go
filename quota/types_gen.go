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
