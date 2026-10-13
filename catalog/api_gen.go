// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package catalog

import (
	"context"
	"net/url"

	basaltic "github.com/basaltic-sh/sdk-go"
)

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
