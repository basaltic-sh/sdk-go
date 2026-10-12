// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package catalog

import (
	"context"
	"net/url"
	"strconv"

	basaltic "github.com/basaltic-sh/sdk-go"
)

// ListCitiesParams are the optional filters and pagination controls for
// [Client.ListCities]. A nil *ListCitiesParams sends none of them.
type ListCitiesParams struct {
	// Country UUID, CRN, or ISO code. Required unless `state` is a UUID or
	// CRN.
	Country string
	Limit   int

	// Q name substring, compared without case or accents.
	Q string

	// State UUID or CRN, or its code together with `country`.
	State string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListCitiesParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.Country != "" {
		q.Set("country", p.Country)
	}
	if p.Limit != 0 {
		q.Set("limit", strconv.Itoa(int(p.Limit)))
	}
	if p.Q != "" {
		q.Set("q", p.Q)
	}
	if p.State != "" {
		q.Set("state", p.State)
	}
	return q
}

// ListRegionsParams are the optional filters and pagination controls for
// [Client.ListRegions]. A nil *ListRegionsParams sends none of them.
type ListRegionsParams struct {
	// CRN exact returned CRN, combined with name using AND before pagination.
	// Malformed or empty CRNs return 400; valid mismatched or foreign CRNs
	// return an empty page. Resources without a CRN never match.
	// Organization-scoped CRNs use the authenticated organization.
	CRN string

	// Name exact resource name, combined with crn using AND before pagination.
	// Empty values are filters. Resources without a name never match.
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListRegionsParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.CRN != "" {
		q.Set("crn", p.CRN)
	}
	if p.Name != "" {
		q.Set("name", p.Name)
	}
	return q
}

// ListStatesParams are the optional filters and pagination controls for
// [Client.ListStates]. A nil *ListStatesParams sends none of them.
type ListStatesParams struct {
	// Country UUID, CRN, or ISO code.
	Country string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListStatesParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.Country != "" {
		q.Set("country", p.Country)
	}
	return q
}

// LookupPostalCodeParams are the optional filters and pagination controls for
// [Client.LookupPostalCode]. A nil *LookupPostalCodeParams sends none of them.
type LookupPostalCodeParams struct {
	// Country UUID, CRN, or ISO code.
	Country string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *LookupPostalCodeParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.Country != "" {
		q.Set("country", p.Country)
	}
	return q
}

// GetCountry gets a country and its address rules.
//
// Return one country with the rules a billing profile follows there:
// customer types, accepted tax identifier kinds, and the address fields
// in form order with their labels, masks and patterns. No authentication
// is required; responses may be cached privately for an hour.
//
// Sends no bearer token: the credentials in the request are the
// authentication.
func (c *Client) GetCountry(ctx context.Context, country string, opts ...basaltic.RequestOption) (*CountryDetail, error) {
	op := &basaltic.Operation{
		ID:              "getCountry",
		Method:          "GET",
		Path:            "/v1/countries/{country}",
		PathArgs:        []string{country},
		Unauthenticated: true,
	}
	var out struct {
		Country *CountryDetail `json:"country"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Country, nil
}

// GetRegion gets a region.
//
// Return one published region by its immutable code. Unknown and
// unpublished regions both return 404. This public endpoint is
// rate-limited per client IP; successful responses may be cached
// privately for 60 seconds.
//
// Sends no bearer token: the credentials in the request are the
// authentication.
func (c *Client) GetRegion(ctx context.Context, code string, opts ...basaltic.RequestOption) (*Region, error) {
	op := &basaltic.Operation{
		ID:              "getRegion",
		Method:          "GET",
		Path:            "/v1/regions/{code}",
		PathArgs:        []string{code},
		Unauthenticated: true,
	}
	var out Region
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListCities lists cities.
//
// List cities of a state, or of a country without states, ordered by
// name. Only countries whose rules list cities have any. `q` matches a
// substring of the name ignoring case and accents. When more rows match
// than `limit`, `meta.has_more` is true. No authentication is required.
//
// Sends no bearer token: the credentials in the request are the
// authentication.
func (c *Client) ListCities(ctx context.Context, params *ListCitiesParams, opts ...basaltic.RequestOption) (*ListCitiesResult, error) {
	op := &basaltic.Operation{
		ID:              "listCities",
		Method:          "GET",
		Path:            "/v1/cities",
		Unauthenticated: true,
	}
	op.Query = params.query()
	var out ListCitiesResult
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListCountries lists countries.
//
// List every ISO 3166-1 country, ordered by English name. No
// authentication is required; requests are rate-limited per client IP.
// Reference geography changes only with a platform release, so responses
// may be cached privately for an hour. Fetch one country for its address
// and tax identifier rules.
//
// Sends no bearer token: the credentials in the request are the
// authentication.
func (c *Client) ListCountries(ctx context.Context, opts ...basaltic.RequestOption) (*ListCountriesResult, error) {
	op := &basaltic.Operation{
		ID:              "listCountries",
		Method:          "GET",
		Path:            "/v1/countries",
		Unauthenticated: true,
	}
	var out ListCountriesResult
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListRegions lists regions.
//
// List published regions, including planned and retiring locations. No
// authentication is required. Requests are rate-limited per client IP.
// Exact name and CRN filters combine with AND; an empty name matches
// nothing and an empty or malformed CRN returns 400. The default is the
// configured accepting region if it is in the result, otherwise the
// first accepting region by code, or an empty string when none accepts
// new resources.
//
// Successful responses may be cached privately for 60 seconds. Discovery
// does not grant provisioning permissions or guarantee capacity or
// health. A catalog outage or retirement does not delete resources or
// change their existing placement.
//
// Sends no bearer token: the credentials in the request are the
// authentication.
func (c *Client) ListRegions(ctx context.Context, params *ListRegionsParams, opts ...basaltic.RequestOption) (*ListRegionsResult, error) {
	op := &basaltic.Operation{
		ID:              "listRegions",
		Method:          "GET",
		Path:            "/v1/regions",
		Unauthenticated: true,
	}
	op.Query = params.query()
	var out ListRegionsResult
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListStates lists a country's states.
//
// List the first-level subdivisions of a country, ordered by name. The
// list is empty for a country whose rules type the state as text or do
// not ask for it. No authentication is required.
//
// Sends no bearer token: the credentials in the request are the
// authentication.
func (c *Client) ListStates(ctx context.Context, params *ListStatesParams, opts ...basaltic.RequestOption) (*ListStatesResult, error) {
	op := &basaltic.Operation{
		ID:              "listStates",
		Method:          "GET",
		Path:            "/v1/states",
		Unauthenticated: true,
	}
	op.Query = params.query()
	var out ListStatesResult
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// LookupPostalCode — Look up a postal code.
//
// Resolve a postal code through public providers and match the answer
// against the catalog. Each country has an ordered list of providers; an
// error or a miss from one falls through to the next, so one provider's
// outage does not stop a lookup. Brazil asks ViaCEP, BrasilAPI and
// OpenCEP; about seventy other countries ask Zippopotam.us. Country
// detail reports which countries are covered in `postal_lookup`.
//
// The answer only pre-fills a form: state and city come back as catalog
// references where the country lists them, and the street as text. A
// lookup that cannot answer never blocks saving a billing profile. No
// authentication is required; requests are rate-limited per client IP
// and successful answers may be cached privately for a day.
//
// Sends no bearer token: the credentials in the request are the
// authentication.
func (c *Client) LookupPostalCode(ctx context.Context, code string, params *LookupPostalCodeParams, opts ...basaltic.RequestOption) (*PostalCode, error) {
	op := &basaltic.Operation{
		ID:              "lookupPostalCode",
		Method:          "GET",
		Path:            "/v1/postal-codes/{code}",
		PathArgs:        []string{code},
		Unauthenticated: true,
	}
	op.Query = params.query()
	var out struct {
		PostalCode *PostalCode `json:"postal_code"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.PostalCode, nil
}
