// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package billing

import (
	"context"
	"io"
	"iter"
	"net/url"
	"strconv"
	"time"

	basaltic "github.com/basaltic-sh/sdk-go"
)

// ListCreditsParams are the optional filters and pagination controls for
// [Client.ListCredits]. A nil *ListCreditsParams sends none of them.
type ListCreditsParams struct {
	// Limit maximum number of items to return. A value above the maximum is
	// clamped to it rather than rejected, so a page shorter than the one
	// you asked for is normal — page until `meta.has_more` is false, not
	// until a page looks short.
	Limit int

	// Marker opaque pagination cursor. Echo back the `meta.marker` value from the
	// previous page to fetch the next one; do not construct or parse it.
	// The token's internal form varies by endpoint (a resource ID, a
	// timestamp, …) and is not guaranteed stable across releases.
	Marker string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListCreditsParams) query() url.Values {
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
func (p *ListCreditsParams) withMarker(marker string) *ListCreditsParams {
	var out ListCreditsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListInvoicesParams are the optional filters and pagination controls for
// [Client.ListInvoices]. A nil *ListInvoicesParams sends none of them.
type ListInvoicesParams struct {
	// Limit maximum number of items to return. A value above the maximum is
	// clamped to it rather than rejected, so a page shorter than the one
	// you asked for is normal — page until `meta.has_more` is false, not
	// until a page looks short.
	Limit int

	// Marker opaque pagination cursor. Echo back the `meta.marker` value from the
	// previous page to fetch the next one; do not construct or parse it.
	// The token's internal form varies by endpoint (a resource ID, a
	// timestamp, …) and is not guaranteed stable across releases.
	Marker string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListInvoicesParams) query() url.Values {
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
func (p *ListInvoicesParams) withMarker(marker string) *ListInvoicesParams {
	var out ListInvoicesParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListPaymentsParams are the optional filters and pagination controls for
// [Client.ListPayments]. A nil *ListPaymentsParams sends none of them.
type ListPaymentsParams struct {
	// Limit maximum number of items to return. A value above the maximum is
	// clamped to it rather than rejected, so a page shorter than the one
	// you asked for is normal — page until `meta.has_more` is false, not
	// until a page looks short.
	Limit int

	// Marker opaque pagination cursor. Echo back the `meta.marker` value from the
	// previous page to fetch the next one; do not construct or parse it.
	// The token's internal form varies by endpoint (a resource ID, a
	// timestamp, …) and is not guaranteed stable across releases.
	Marker string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListPaymentsParams) query() url.Values {
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
func (p *ListPaymentsParams) withMarker(marker string) *ListPaymentsParams {
	var out ListPaymentsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListPricesParams are the optional filters and pagination controls for
// [Client.ListPrices]. A nil *ListPricesParams sends none of them.
type ListPricesParams struct {
	// At read the catalog as of this instant instead of now, for showing a
	// historical price. RFC3339.
	At time.Time

	// Family Only SKUs whose `metadata.family` matches — how the managed
	// products are separated from the general compute flavors.
	Family       string
	ResourceType string

	// Service Only SKUs billed by this service.
	Service string

	// Sku exactly one SKU.
	Sku string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListPricesParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if !p.At.IsZero() {
		q.Set("at", p.At.UTC().Format(time.RFC3339))
	}
	if p.Family != "" {
		q.Set("family", p.Family)
	}
	if p.ResourceType != "" {
		q.Set("resource_type", p.ResourceType)
	}
	if p.Service != "" {
		q.Set("service", p.Service)
	}
	if p.Sku != "" {
		q.Set("sku", p.Sku)
	}
	return q
}

// ListTransactionsParams are the optional filters and pagination controls for
// [Client.ListTransactions]. A nil *ListTransactionsParams sends none of them.
type ListTransactionsParams struct {
	// Limit maximum number of items to return. A value above the maximum is
	// clamped to it rather than rejected, so a page shorter than the one
	// you asked for is normal — page until `meta.has_more` is false, not
	// until a page looks short.
	Limit int

	// Marker opaque pagination cursor. Echo back the `meta.marker` value from the
	// previous page to fetch the next one; do not construct or parse it.
	// The token's internal form varies by endpoint (a resource ID, a
	// timestamp, …) and is not guaranteed stable across releases.
	Marker string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListTransactionsParams) query() url.Values {
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
func (p *ListTransactionsParams) withMarker(marker string) *ListTransactionsParams {
	var out ListTransactionsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// GetCurrentUsage gets month-to-date usage total.
func (c *Client) GetCurrentUsage(ctx context.Context, opts ...basaltic.RequestOption) (*CurrentUsage, error) {
	op := &basaltic.Operation{
		ID:     "getCurrentUsage",
		Method: "GET",
		Path:   "/v1/usage",
	}
	var out CurrentUsage
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetInvoice gets an invoice with its line items.
func (c *Client) GetInvoice(ctx context.Context, invoiceID string, opts ...basaltic.RequestOption) (*Invoice, error) {
	op := &basaltic.Operation{
		ID:       "getInvoice",
		Method:   "GET",
		Path:     "/v1/invoices/{invoice_id}",
		PathArgs: []string{invoiceID},
	}
	var out Invoice
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetInvoicePDF downloads an invoice as a PDF statement.
//
// Rendered on demand from the invoice's current state — there is no
// stored artifact to age out of sync with its status.
//
// The caller must close the returned reader.
func (c *Client) GetInvoicePDF(ctx context.Context, invoiceID string, opts ...basaltic.RequestOption) (io.ReadCloser, error) {
	op := &basaltic.Operation{
		ID:       "getInvoicePdf",
		Method:   "GET",
		Path:     "/v1/invoices/{invoice_id}/pdf",
		PathArgs: []string{invoiceID},
	}
	stream, _, err := c.rt.DoStream(ctx, op, opts...)
	if err != nil {
		return nil, err
	}
	return stream, nil
}

// ListCredits lists credit grants.
//
// Returns one page. Use ListCreditsAll to walk every page.
func (c *Client) ListCredits(ctx context.Context, params *ListCreditsParams, opts ...basaltic.RequestOption) (*basaltic.Page[Credit], error) {
	op := &basaltic.Operation{
		ID:     "listCredits",
		Method: "GET",
		Path:   "/v1/credits",
	}
	op.Query = params.query()
	var out struct {
		Items []Credit `json:"credits"`
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
	page := &basaltic.Page[Credit]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListCreditsAll walks every page of ListCredits, yielding one item at a
// time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListCreditsAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListCreditsAll(ctx context.Context, params *ListCreditsParams, opts ...basaltic.RequestOption) iter.Seq2[Credit, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Credit], error) {
		return c.ListCredits(ctx, params.withMarker(marker), opts...)
	})
}

// ListInvoices lists invoices.
//
// Returns one page. Use ListInvoicesAll to walk every page.
func (c *Client) ListInvoices(ctx context.Context, params *ListInvoicesParams, opts ...basaltic.RequestOption) (*basaltic.Page[Invoice], error) {
	op := &basaltic.Operation{
		ID:     "listInvoices",
		Method: "GET",
		Path:   "/v1/invoices",
	}
	op.Query = params.query()
	var out struct {
		Items []Invoice `json:"invoices"`
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
	page := &basaltic.Page[Invoice]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListInvoicesAll walks every page of ListInvoices, yielding one item at
// a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListInvoicesAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListInvoicesAll(ctx context.Context, params *ListInvoicesParams, opts ...basaltic.RequestOption) iter.Seq2[Invoice, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Invoice], error) {
		return c.ListInvoices(ctx, params.withMarker(marker), opts...)
	})
}

// ListPayments lists invoice payments.
//
// Returns one page. Use ListPaymentsAll to walk every page.
func (c *Client) ListPayments(ctx context.Context, params *ListPaymentsParams, opts ...basaltic.RequestOption) (*basaltic.Page[Payment], error) {
	op := &basaltic.Operation{
		ID:     "listPayments",
		Method: "GET",
		Path:   "/v1/payments",
	}
	op.Query = params.query()
	var out struct {
		Items []Payment `json:"payments"`
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
	page := &basaltic.Page[Payment]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListPaymentsAll walks every page of ListPayments, yielding one item at
// a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListPaymentsAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListPaymentsAll(ctx context.Context, params *ListPaymentsParams, opts ...basaltic.RequestOption) iter.Seq2[Payment, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Payment], error) {
		return c.ListPayments(ctx, params.withMarker(marker), opts...)
	})
}

// ListPrices lists catalog prices.
//
// The public price catalog — every rate the platform charges,
// effective now (or at `at`). This endpoint is public and takes no
// credentials: the rates are identical for every caller, with no
// account-specific discounts or committed-use terms, so there is nothing
// tenant-scoped to protect. It exists so the marketing site and the
// console read prices from billing instead of mirroring them in source,
// where they drift every time a migration reprices.
//
// Because it takes no credentials, requests are rate-limited per client
// IP. Responses carry a short public `Cache-Control` — the catalog
// changes on a migration, not on a request.
//
// Sends no bearer token: the credentials in the request are the
// authentication.
func (c *Client) ListPrices(ctx context.Context, params *ListPricesParams, opts ...basaltic.RequestOption) (*PriceListResponse, error) {
	op := &basaltic.Operation{
		ID:              "listPrices",
		Method:          "GET",
		Path:            "/v1/prices",
		Unauthenticated: true,
	}
	op.Query = params.query()
	var out PriceListResponse
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListTransactions lists ledger transactions.
//
// Returns one page. Use ListTransactionsAll to walk every page.
func (c *Client) ListTransactions(ctx context.Context, params *ListTransactionsParams, opts ...basaltic.RequestOption) (*basaltic.Page[Transaction], error) {
	op := &basaltic.Operation{
		ID:     "listTransactions",
		Method: "GET",
		Path:   "/v1/transactions",
	}
	op.Query = params.query()
	var out struct {
		Items []Transaction `json:"transactions"`
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
	page := &basaltic.Page[Transaction]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListTransactionsAll walks every page of ListTransactions, yielding one
// item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListTransactionsAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListTransactionsAll(ctx context.Context, params *ListTransactionsParams, opts ...basaltic.RequestOption) iter.Seq2[Transaction, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Transaction], error) {
		return c.ListTransactions(ctx, params.withMarker(marker), opts...)
	})
}
