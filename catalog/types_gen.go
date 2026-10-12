// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package catalog

type AddressField struct {
	Example string `json:"example,omitempty"`

	// One of: "street_name", "street_number", "complement", "neighborhood", "city", "state", "postal_code".
	Key   string `json:"key"`
	Label string `json:"label"`

	// Mask display mask: `9` a digit, `A` a letter, `*` a letter or digit, `?`
	// makes the previous symbol optional, anything else is a literal.
	// Values are stored without the literals.
	Mask string `json:"mask,omitempty"`

	// Max maximum length in characters.
	Max int `json:"max,omitempty"`

	// Pattern regular expression the stored value must match in full.
	Pattern string `json:"pattern,omitempty"`

	// RequireMetadata for a listed city, an internal attribute the city must carry for the
	// profile to be complete.
	RequireMetadata string `json:"require_metadata,omitempty"`
	Required        bool   `json:"required"`

	// Source for state and city only. `list` means pick a catalog row and send
	// its selector; `free` means type the value under the address key.
	//
	// One of: "list", "free".
	Source string `json:"source,omitempty"`
}

type City struct {
	Country *ResourceRef `json:"country"`
	CRN     string       `json:"crn"`
	ID      string       `json:"id"`
	Name    string       `json:"name"`

	// State null for countries without states.
	State *ResourceRef `json:"state"`
}

type Country struct {
	// CallingCode international calling code without the plus sign. Several countries
	// share one (1 covers the United States and Canada). Empty when none
	// is assigned.
	CallingCode string `json:"calling_code"`

	// Code ISO 3166-1 alpha-2 code. Also accepted wherever a country selector
	// is.
	Code string `json:"code"`
	CRN  string `json:"crn"`

	// ID stable in every environment.
	ID string `json:"id"`

	// Iso3 ISO 3166-1 alpha-3 code.
	Iso3 string `json:"iso3"`

	// Name english short name. Clients localize it from `code`.
	Name string `json:"name"`
}

type CountryDetail struct {
	// CallingCode international calling code without the plus sign. Several countries
	// share one (1 covers the United States and Canada). Empty when none
	// is assigned.
	CallingCode string `json:"calling_code"`

	// Code ISO 3166-1 alpha-2 code. Also accepted wherever a country selector
	// is.
	Code string `json:"code"`
	CRN  string `json:"crn"`

	// ID stable in every environment.
	ID string `json:"id"`

	// Iso3 ISO 3166-1 alpha-3 code.
	Iso3 string `json:"iso3"`

	// Name english short name. Clients localize it from `code`.
	Name string `json:"name"`

	// PostalLookup true when `GET /v1/postal-codes/{code}` can resolve this country's
	// postal codes.
	PostalLookup bool          `json:"postal_lookup"`
	Rules        *CountryRules `json:"rules"`
}

// CountryRules how a billing profile is entered and validated for the country. Every
// part is present: defaults fill whatever the country does not
// configure.
type CountryRules struct {
	CustomerTypes []string `json:"customer_types"`

	// Fields address fields in form order. A key that is absent is not asked for
	// in this country.
	Fields []*AddressField `json:"fields"`

	// NameLabels label for the recipient name, keyed by customer type.
	NameLabels map[string]string `json:"name_labels"`

	// TaxIDs accepted tax identifier kinds, keyed by customer type.
	TaxIDs map[string][]*TaxIDRule `json:"tax_ids"`
}

type GeoPage struct {
	// HasMore true when more rows matched than were returned. Narrow with `q`.
	HasMore bool `json:"has_more"`
	Limit   int  `json:"limit"`

	// Total items in this response.
	Total int `json:"total"`
}

type ListCitiesResult struct {
	Cities []*City  `json:"cities"`
	Meta   *GeoPage `json:"meta"`
}

type ListCountriesResult struct {
	Countries []*Country `json:"countries"`
	Meta      *GeoPage   `json:"meta"`
}

type ListRegionsResult struct {
	// Default accepting region code, or an empty string.
	Default string    `json:"default"`
	Regions []*Region `json:"regions"`
}

type ListStatesResult struct {
	Meta   *GeoPage `json:"meta"`
	States []*State `json:"states"`
}

type PostalCode struct {
	// Address other address values the provider knew, limited to keys the country
	// uses. State and city appear here as text where the country types
	// them.
	Address map[string]string `json:"address"`

	// City the resolved catalog city, when the country lists cities and the
	// answer matched one.
	City *ResourceRef `json:"city"`

	// Code the postal code as stored, without mask literals.
	Code    string       `json:"code"`
	Country *ResourceRef `json:"country"`

	// Source the provider that answered.
	Source string `json:"source"`

	// State the resolved catalog state, when the country lists states and the
	// answer matched one.
	State *ResourceRef `json:"state"`
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

type ResourceRef struct {
	CRN string `json:"crn"`
	ID  string `json:"id"`

	// Name display label at response time — the resource's name where it has
	// one, otherwise its human label (an invoice number, an account
	// display name). Absent when the target is gone. Not a lookup key.
	Name string `json:"name,omitempty"`
}

type State struct {
	// Code ISO 3166-2 subdivision code without the country prefix, where one
	// exists.
	Code    string       `json:"code"`
	Country *ResourceRef `json:"country"`
	CRN     string       `json:"crn"`
	ID      string       `json:"id"`
	Name    string       `json:"name"`
}

type TaxIDRule struct {
	Example string `json:"example,omitempty"`

	// Kind identifier kind, returned as `tax_id_kind` on the billing profile.
	// `other` is free text in a country without curated rules.
	Kind    string `json:"kind"`
	Label   string `json:"label"`
	Mask    string `json:"mask,omitempty"`
	Pattern string `json:"pattern,omitempty"`

	// Required when true, `no_tax_id_reason` is not accepted instead.
	Required bool `json:"required"`

	// Validator named check-digit algorithm the server applies.
	//
	// One of: "br_cpf", "br_cnpj", "ar_cuit", "cl_rut".
	Validator string `json:"validator,omitempty"`
}
