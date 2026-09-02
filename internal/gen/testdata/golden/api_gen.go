// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package widget

import (
	"context"
	"io"
	"iter"
	"net/url"
	"strconv"

	basaltic "github.com/basaltic-sh/sdk-go"
)

// ListWidgetsParams are the optional filters and pagination controls for
// [Client.ListWidgets]. A nil *ListWidgetsParams sends none of them.
type ListWidgetsParams struct {
	// IncludeRetired include retired widgets.
	IncludeRetired *bool

	// Limit maximum number of items to return.
	Limit int

	// Marker opaque pagination cursor.
	Marker string

	// WidgetState filter by lifecycle state.
	WidgetState WidgetState
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListWidgetsParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.IncludeRetired != nil {
		q.Set("include_retired", strconv.FormatBool(*p.IncludeRetired))
	}
	if p.Limit != 0 {
		q.Set("limit", strconv.Itoa(int(p.Limit)))
	}
	if p.Marker != "" {
		q.Set("marker", p.Marker)
	}
	if p.WidgetState != "" {
		q.Set("widget_state", string(p.WidgetState))
	}
	return q
}

// withMarker copies p with the pagination cursor replaced, leaving the
// caller's value untouched across pages.
func (p *ListWidgetsParams) withMarker(marker string) *ListWidgetsParams {
	var out ListWidgetsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// CreateWidget creates a widget.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateWidget(ctx context.Context, body *WidgetCreateRequest, opts ...basaltic.RequestOption) (*Widget, error) {
	op := &basaltic.Operation{
		ID:     "createWidget",
		Method: "POST",
		Path:   "/v1/widgets",
		Body:   body,
	}
	var out struct {
		Widget *Widget `json:"widget"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Widget, nil
}

// DeleteWidget deletes a widget.
func (c *Client) DeleteWidget(ctx context.Context, widgetID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteWidget",
		Method:   "DELETE",
		Path:     "/v1/widgets/{widget_id}",
		PathArgs: []string{widgetID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DownloadBlueprint downloads a widget blueprint.
//
// The caller must close the returned reader.
func (c *Client) DownloadBlueprint(ctx context.Context, widgetID string, opts ...basaltic.RequestOption) (io.ReadCloser, error) {
	op := &basaltic.Operation{
		ID:       "downloadBlueprint",
		Method:   "GET",
		Path:     "/v1/widgets/{widget_id}/blueprint",
		PathArgs: []string{widgetID},
	}
	stream, _, err := c.rt.DoStream(ctx, op, opts...)
	if err != nil {
		return nil, err
	}
	return stream, nil
}

// GetToken exchanges credentials for a token.
//
// Sends no bearer token: the credentials in the request are the
// authentication.
func (c *Client) GetToken(ctx context.Context, body *GetTokenRequest, opts ...basaltic.RequestOption) (*GetTokenResult, error) {
	op := &basaltic.Operation{
		ID:              "getToken",
		Method:          "POST",
		Path:            "/v1/token",
		Body:            body,
		Unauthenticated: true,
	}
	var out GetTokenResult
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetWidget gets a widget.
func (c *Client) GetWidget(ctx context.Context, widgetID string, opts ...basaltic.RequestOption) (*Widget, error) {
	op := &basaltic.Operation{
		ID:       "getWidget",
		Method:   "GET",
		Path:     "/v1/widgets/{widget_id}",
		PathArgs: []string{widgetID},
	}
	var out struct {
		Widget *Widget `json:"widget"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Widget, nil
}

// ListWidgets lists widgets.
//
// List every widget in the account.
//
// Returns one page. Use ListWidgetsAll to walk every page.
func (c *Client) ListWidgets(ctx context.Context, params *ListWidgetsParams, opts ...basaltic.RequestOption) (*basaltic.Page[Widget], error) {
	op := &basaltic.Operation{
		ID:     "listWidgets",
		Method: "GET",
		Path:   "/v1/widgets",
	}
	op.Query = params.query()
	var out struct {
		Items []Widget `json:"widgets"`
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
	page := &basaltic.Page[Widget]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListWidgetsAll walks every page of ListWidgets, yielding one item at a
// time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListWidgetsAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListWidgetsAll(ctx context.Context, params *ListWidgetsParams, opts ...basaltic.RequestOption) iter.Seq2[Widget, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Widget], error) {
		return c.ListWidgets(ctx, params.withMarker(marker), opts...)
	})
}
