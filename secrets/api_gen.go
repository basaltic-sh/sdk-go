// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package secrets

import (
	"context"
	"iter"
	"net/url"
	"strconv"

	basaltic "github.com/basaltic-sh/sdk-go"
)

// GetSecretValueParams are the optional filters and pagination controls for
// [Client.GetSecretValue]. A nil *GetSecretValueParams sends none of them.
type GetSecretValueParams struct {
	// Version specific version to read. Omit for current.
	Version int
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *GetSecretValueParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.Version != 0 {
		q.Set("version", strconv.Itoa(int(p.Version)))
	}
	return q
}

// ListSecretsParams are the optional filters and pagination controls for
// [Client.ListSecrets]. A nil *ListSecretsParams sends none of them.
type ListSecretsParams struct {
	// IncludeDeleted include secrets in the recovery window.
	IncludeDeleted *bool

	// Limit maximum items to return. A larger value is clamped to the maximum
	// rather than rejected, so page until `meta.has_more` is false.
	Limit  int
	Marker string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListSecretsParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.IncludeDeleted != nil {
		q.Set("include_deleted", strconv.FormatBool(*p.IncludeDeleted))
	}
	if p.Limit != 0 {
		q.Set("limit", strconv.Itoa(int(p.Limit)))
	}
	if p.Marker != "" {
		q.Set("marker", p.Marker)
	}
	return q
}

// withMarker copies p with the pagination cursor replaced, leaving the
// caller's value untouched across pages.
func (p *ListSecretsParams) withMarker(marker string) *ListSecretsParams {
	var out ListSecretsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListVersionsParams are the optional filters and pagination controls for
// [Client.ListVersions]. A nil *ListVersionsParams sends none of them.
type ListVersionsParams struct {
	// Limit maximum items to return. A larger value is clamped to the maximum
	// rather than rejected, so page until `meta.has_more` is false.
	Limit  int
	Marker string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListVersionsParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.Limit != 0 {
		q.Set("limit", strconv.Itoa(int(p.Limit)))
	}
	if p.Marker != "" {
		q.Set("marker", p.Marker)
	}
	return q
}

// withMarker copies p with the pagination cursor replaced, leaving the
// caller's value untouched across pages.
func (p *ListVersionsParams) withMarker(marker string) *ListVersionsParams {
	var out ListVersionsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// CreateSecret creates a new secret with an initial value.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateSecret(ctx context.Context, body *CreateSecretRequest, opts ...basaltic.RequestOption) (*Secret, error) {
	op := &basaltic.Operation{
		ID:     "createSecret",
		Method: "POST",
		Path:   "/v1/secrets",
		Body:   body,
	}
	var out struct {
		Secret *Secret `json:"secret"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Secret, nil
}

// DeleteSecret schedules deletion (soft delete with recovery window).
func (c *Client) DeleteSecret(ctx context.Context, secretID string, body *DeleteSecretRequest, opts ...basaltic.RequestOption) (*Secret, error) {
	op := &basaltic.Operation{
		ID:       "deleteSecret",
		Method:   "DELETE",
		Path:     "/v1/secrets/{secret_id}",
		PathArgs: []string{secretID},
		Body:     body,
	}
	var out struct {
		Secret *Secret `json:"secret"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Secret, nil
}

// DescribeSecret describes a secret (no value).
func (c *Client) DescribeSecret(ctx context.Context, secretID string, opts ...basaltic.RequestOption) (*Secret, error) {
	op := &basaltic.Operation{
		ID:       "describeSecret",
		Method:   "GET",
		Path:     "/v1/secrets/{secret_id}",
		PathArgs: []string{secretID},
	}
	var out struct {
		Secret *Secret `json:"secret"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Secret, nil
}

// GetSecretValue reads the current value (or a specific version).
func (c *Client) GetSecretValue(ctx context.Context, secretID string, params *GetSecretValueParams, opts ...basaltic.RequestOption) (*SecretValue, error) {
	op := &basaltic.Operation{
		ID:       "getSecretValue",
		Method:   "GET",
		Path:     "/v1/secrets/{secret_id}/value",
		PathArgs: []string{secretID},
	}
	op.Query = params.query()
	var out struct {
		Secret *SecretValue `json:"secret"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Secret, nil
}

// ListSecrets lists secrets.
//
// Returns one page. Use ListSecretsAll to walk every page.
func (c *Client) ListSecrets(ctx context.Context, params *ListSecretsParams, opts ...basaltic.RequestOption) (*basaltic.Page[Secret], error) {
	op := &basaltic.Operation{
		ID:     "listSecrets",
		Method: "GET",
		Path:   "/v1/secrets",
	}
	op.Query = params.query()
	var out struct {
		Items []Secret `json:"secrets"`
		Meta  *struct {
			Total   int    `json:"total"`
			Limit   int    `json:"limit"`
			Marker  string `json:"marker"`
			HasMore bool   `json:"has_more"`
		} `json:"meta"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[Secret]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListSecretsAll walks every page of ListSecrets, yielding one item at a
// time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListSecretsAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListSecretsAll(ctx context.Context, params *ListSecretsParams, opts ...basaltic.RequestOption) iter.Seq2[Secret, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Secret], error) {
		return c.ListSecrets(ctx, params.withMarker(marker), opts...)
	})
}

// ListVersions lists versions.
//
// Returns one page. Use ListVersionsAll to walk every page.
func (c *Client) ListVersions(ctx context.Context, secretID string, params *ListVersionsParams, opts ...basaltic.RequestOption) (*basaltic.Page[SecretVersion], error) {
	op := &basaltic.Operation{
		ID:       "listVersions",
		Method:   "GET",
		Path:     "/v1/secrets/{secret_id}/versions",
		PathArgs: []string{secretID},
	}
	op.Query = params.query()
	var out struct {
		Items []SecretVersion `json:"versions"`
		Meta  *struct {
			Total   int    `json:"total"`
			Limit   int    `json:"limit"`
			Marker  string `json:"marker"`
			HasMore bool   `json:"has_more"`
		} `json:"meta"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[SecretVersion]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListVersionsAll walks every page of ListVersions, yielding one item at
// a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListVersionsAll(ctx, secretID, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListVersionsAll(ctx context.Context, secretID string, params *ListVersionsParams, opts ...basaltic.RequestOption) iter.Seq2[SecretVersion, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[SecretVersion], error) {
		return c.ListVersions(ctx, secretID, params.withMarker(marker), opts...)
	})
}

// PutSecretValue stores a new version (becomes current).
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) PutSecretValue(ctx context.Context, secretID string, body *PutSecretValueRequest, opts ...basaltic.RequestOption) (*SecretVersion, error) {
	op := &basaltic.Operation{
		ID:       "putSecretValue",
		Method:   "POST",
		Path:     "/v1/secrets/{secret_id}/value",
		PathArgs: []string{secretID},
		Body:     body,
	}
	var out struct {
		Version *SecretVersion `json:"version"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Version, nil
}

// RestoreSecret restores a secret from the recovery window.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) RestoreSecret(ctx context.Context, secretID string, opts ...basaltic.RequestOption) (*Secret, error) {
	op := &basaltic.Operation{
		ID:       "restoreSecret",
		Method:   "POST",
		Path:     "/v1/secrets/{secret_id}/restore",
		PathArgs: []string{secretID},
	}
	var out struct {
		Secret *Secret `json:"secret"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Secret, nil
}

// UpdateSecret updates mutable metadata.
func (c *Client) UpdateSecret(ctx context.Context, secretID string, body *UpdateSecretRequest, opts ...basaltic.RequestOption) (*Secret, error) {
	op := &basaltic.Operation{
		ID:       "updateSecret",
		Method:   "PATCH",
		Path:     "/v1/secrets/{secret_id}",
		PathArgs: []string{secretID},
		Body:     body,
	}
	var out struct {
		Secret *Secret `json:"secret"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Secret, nil
}
