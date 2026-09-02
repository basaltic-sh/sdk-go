// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package quota

import (
	"context"
	"net/url"

	basaltic "github.com/basaltic-sh/sdk-go"
)

// ListQuotasParams are the optional filters and pagination controls for
// [Client.ListQuotas]. A nil *ListQuotasParams sends none of them.
type ListQuotasParams struct {
	// Region to filter regional quotas by. Omit for global only.
	Region string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListQuotasParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.Region != "" {
		q.Set("region", p.Region)
	}
	return q
}

// ListQuotas lists quotas.
//
// List quotas for the caller's organization, joined with current usage
// and any per-org overrides. Filter by region with `?region=`; omitting
// it returns global-scoped quotas only.
//
// Per-resource quotas (e.g. `domains_per_certificate`) are NOT returned
// here — they're per-parent caps, surfaced only as validation errors
// at the relevant resource API. Requires `quota:GetQuotas` permission.
func (c *Client) ListQuotas(ctx context.Context, params *ListQuotasParams, opts ...basaltic.RequestOption) (*basaltic.Page[QuotaItem], error) {
	op := &basaltic.Operation{
		ID:     "listQuotas",
		Method: "GET",
		Path:   "/v1/quotas",
	}
	op.Query = params.query()
	var out struct {
		Items []QuotaItem `json:"quotas"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[QuotaItem]{Items: out.Items}
	return page, nil
}
