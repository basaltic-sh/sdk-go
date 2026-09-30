// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package catalog

type ListRegionsResult struct {
	// Default accepting region code, or an empty string.
	Default string    `json:"default"`
	Regions []*Region `json:"regions"`
}

type Region struct {
	// AcceptingNewResources true only for active regions. Product services still enforce
	// eligibility, quotas, placement and capacity.
	AcceptingNewResources bool `json:"accepting_new_resources"`

	// Available compatibility alias for accepting_new_resources; not a live health
	// signal
	Available bool `json:"available"`

	// Code unique region code used in API calls and CRNs
	Code string `json:"code"`

	// ComingSoon compatibility flag that is true only when state is planned
	ComingSoon bool `json:"coming_soon"`

	// CountryCode ISO 3166-1 alpha-2 country code (used to display flag in UI)
	CountryCode string `json:"country_code"`

	// CRN platform-owned global region identity, using the immutable region
	// code.
	CRN string `json:"crn"`

	// Location geographic location of the region
	Location string `json:"location"`

	// Name human-readable region name
	Name string `json:"name"`

	// State catalog lifecycle, independent of health and capacity. Retired codes
	// are never reused.
	//
	// One of: "planned", "active", "restricted", "retiring", "retired".
	State string `json:"state"`
}
