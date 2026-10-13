// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package catalog

type ListRegionsResult struct {
	// Default accepting region code, or an empty string.
	Default string          `json:"default"`
	Meta    *PaginationMeta `json:"meta"`
	Regions []*Region       `json:"regions"`
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

type Region struct {
	// Code unique region code used in API calls and CRNs
	Code string `json:"code"`

	// CountryCode ISO 3166-1 alpha-2 country code (used to display flag in UI)
	CountryCode string `json:"country_code"`

	// CRN platform-owned global region identity, using the immutable region
	// code.
	CRN string `json:"crn"`

	// Location geographic location of the region
	Location string `json:"location"`

	// Name human-readable region name
	Name string `json:"name"`

	// State region lifecycle and availability. Only active regions accept new
	// resources; planned, restricted, retiring, and retired regions do
	// not. This does not guarantee live health, quota, or capacity.
	//
	// One of: "planned", "active", "restricted", "retiring", "retired".
	State string `json:"state"`
}
