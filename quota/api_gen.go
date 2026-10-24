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
	// Region for organization resource quotas and Email policies. Omit for
	// global resource quotas only.
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
// it returns global-scoped resource quotas and an empty `policies`
// array.
//
// With `?region=`, the additive `policies` array also exposes Email
// message-size and rolling recipient limits for the organization in that
// region. A regional override takes precedence over an organization-wide
// override, then the platform default. Policies include units,
// rolling-window length where applicable, and the limit source. They do
// not include reservation or usage counters; a size or rate limit must
// not be interpreted as resource capacity. The existing `quotas` array
// and its fields are unchanged.
//
// Other per-resource quotas (e.g. `domains_per_certificate`) are not
// returned. The Email size policy applies to the complete encoded
// sending message, including attachments. The independent 50 MiB service
// ceiling still applies even if the policy is unlimited. Requires
// `quota:GetQuotas` permission.
func (c *Client) ListQuotas(ctx context.Context, params *ListQuotasParams, opts ...basaltic.RequestOption) (*QuotaListResponse, error) {
	op := &basaltic.Operation{
		ID:     "listQuotas",
		Method: "GET",
		Path:   "/v1/quotas",
	}
	op.Query = params.query()
	var out QuotaListResponse
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}
