// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package audit

import (
	"context"
	"iter"
	"net/url"
	"strconv"
	"time"

	basaltic "github.com/basaltic-sh/sdk-go"
)

// ListAuditLogsParams are the optional filters and pagination controls for
// [Client.ListAuditLogs]. A nil *ListAuditLogsParams sends none of them.
type ListAuditLogsParams struct {
	// Action filter by action (exact match or prefix with wildcard, e.g.,
	// "iam.*")
	Action string

	// ActorID filter by actor ID (user or service account)
	ActorID string

	// ActorType filter by actor type
	//
	// One of: "user", "service_account", "system".
	ActorType string

	// From filter logs from this timestamp (inclusive)
	From time.Time

	// IPAddress filter by IP address
	IPAddress string

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

	// ResourceID filter by resource ID
	ResourceID string

	// ResourceType filter by resource type
	ResourceType string

	// Status filter by status
	//
	// One of: "success", "failure", "denied".
	Status string

	// To filter logs until this timestamp (exclusive)
	To time.Time
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListAuditLogsParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.Action != "" {
		q.Set("action", p.Action)
	}
	if p.ActorID != "" {
		q.Set("actor_id", p.ActorID)
	}
	if p.ActorType != "" {
		q.Set("actor_type", p.ActorType)
	}
	if !p.From.IsZero() {
		q.Set("from", p.From.UTC().Format(time.RFC3339))
	}
	if p.IPAddress != "" {
		q.Set("ip_address", p.IPAddress)
	}
	if p.Limit != 0 {
		q.Set("limit", strconv.Itoa(int(p.Limit)))
	}
	if p.Marker != "" {
		q.Set("marker", p.Marker)
	}
	if p.ResourceID != "" {
		q.Set("resource_id", p.ResourceID)
	}
	if p.ResourceType != "" {
		q.Set("resource_type", p.ResourceType)
	}
	if p.Status != "" {
		q.Set("status", p.Status)
	}
	if !p.To.IsZero() {
		q.Set("to", p.To.UTC().Format(time.RFC3339))
	}
	return q
}

// withMarker copies p with the pagination cursor replaced, leaving the
// caller's value untouched across pages.
func (p *ListAuditLogsParams) withMarker(marker string) *ListAuditLogsParams {
	var out ListAuditLogsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// GetAuditLog gets audit log entry.
//
// Get a specific audit log entry by ID. Requires `audit:GetLog`
// permission.
func (c *Client) GetAuditLog(ctx context.Context, logID string, opts ...basaltic.RequestOption) (*AuditLog, error) {
	op := &basaltic.Operation{
		ID:       "getAuditLog",
		Method:   "GET",
		Path:     "/v1/audit-logs/{log_id}",
		PathArgs: []string{logID},
	}
	var out struct {
		AuditLog *AuditLog `json:"audit_log"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.AuditLog, nil
}

// ListAuditLogs lists audit logs.
//
// List audit logs for the current organization with filtering and
// pagination. Requires `audit:ListLogs` permission.
//
// Returns one page. Use ListAuditLogsAll to walk every page.
func (c *Client) ListAuditLogs(ctx context.Context, params *ListAuditLogsParams, opts ...basaltic.RequestOption) (*basaltic.Page[AuditLog], error) {
	op := &basaltic.Operation{
		ID:     "listAuditLogs",
		Method: "GET",
		Path:   "/v1/audit-logs",
	}
	op.Query = params.query()
	var out struct {
		Items []AuditLog `json:"audit_logs"`
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
	page := &basaltic.Page[AuditLog]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListAuditLogsAll walks every page of ListAuditLogs, yielding one item
// at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListAuditLogsAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListAuditLogsAll(ctx context.Context, params *ListAuditLogsParams, opts ...basaltic.RequestOption) iter.Seq2[AuditLog, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[AuditLog], error) {
		return c.ListAuditLogs(ctx, params.withMarker(marker), opts...)
	})
}
