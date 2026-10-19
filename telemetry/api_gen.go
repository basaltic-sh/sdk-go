// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package telemetry

import (
	"context"
	"io"
	"iter"
	"net/url"
	"strconv"
	"time"

	basaltic "github.com/basaltic-sh/sdk-go"
)

// ListAlertsParams are the optional filters and pagination controls for
// [Client.ListAlerts]. A nil *ListAlertsParams sends none of them.
type ListAlertsParams struct {
	// CRN exact canonical alert CRN, intersected with name and state.
	// Empty/malformed CRNs return 400; valid foreign CRNs return an empty
	// page.
	CRN string

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

	// Name exact name filter; an explicitly empty name matches nothing.
	Name string

	// State filter effective state. Disabled takes precedence over silenced;
	// otherwise uses evaluation state.
	//
	// One of: "ok", "pending", "firing", "no_data", "error", "silenced",
	// "disabled".
	State string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListAlertsParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.CRN != "" {
		q.Set("crn", p.CRN)
	}
	if p.Limit != 0 {
		q.Set("limit", strconv.Itoa(int(p.Limit)))
	}
	if p.Marker != "" {
		q.Set("marker", p.Marker)
	}
	if p.Name != "" {
		q.Set("name", p.Name)
	}
	if p.State != "" {
		q.Set("state", p.State)
	}
	return q
}

// withMarker copies p with the pagination cursor replaced, leaving the
// caller's value untouched across pages.
func (p *ListAlertsParams) withMarker(marker string) *ListAlertsParams {
	var out ListAlertsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListLogGroupsParams are the optional filters and pagination controls for
// [Client.ListLogGroups]. A nil *ListLogGroupsParams sends none of them.
type ListLogGroupsParams struct {
	// CRN exact telemetry/log-group CRN filter in the authenticated account
	// and serving region. Malformed or empty CRNs return 400. Valid
	// foreign or mismatched CRNs return an empty page. Intersected with
	// name and pagination.
	CRN string

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

	// Name exact name filter, intersected with crn and pagination.
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListLogGroupsParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.CRN != "" {
		q.Set("crn", p.CRN)
	}
	if p.Limit != 0 {
		q.Set("limit", strconv.Itoa(int(p.Limit)))
	}
	if p.Marker != "" {
		q.Set("marker", p.Marker)
	}
	if p.Name != "" {
		q.Set("name", p.Name)
	}
	return q
}

// withMarker copies p with the pagination cursor replaced, leaving the
// caller's value untouched across pages.
func (p *ListLogGroupsParams) withMarker(marker string) *ListLogGroupsParams {
	var out ListLogGroupsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListMetricNamesParams are the optional filters and pagination controls for
// [Client.ListMetricNames]. A nil *ListMetricNamesParams sends none of them.
type ListMetricNamesParams struct {
	End string

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
	Start  string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListMetricNamesParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.End != "" {
		q.Set("end", p.End)
	}
	if p.Limit != 0 {
		q.Set("limit", strconv.Itoa(int(p.Limit)))
	}
	if p.Marker != "" {
		q.Set("marker", p.Marker)
	}
	if p.Start != "" {
		q.Set("start", p.Start)
	}
	return q
}

// withMarker copies p with the pagination cursor replaced, leaving the
// caller's value untouched across pages.
func (p *ListMetricNamesParams) withMarker(marker string) *ListMetricNamesParams {
	var out ListMetricNamesParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListMetricSeriesParams are the optional filters and pagination controls for
// [Client.ListMetricSeries]. A nil *ListMetricSeriesParams sends none of them.
type ListMetricSeriesParams struct {
	End string

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
	Metric string
	Start  string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListMetricSeriesParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.End != "" {
		q.Set("end", p.End)
	}
	if p.Limit != 0 {
		q.Set("limit", strconv.Itoa(int(p.Limit)))
	}
	if p.Marker != "" {
		q.Set("marker", p.Marker)
	}
	if p.Metric != "" {
		q.Set("metric", p.Metric)
	}
	if p.Start != "" {
		q.Set("start", p.Start)
	}
	return q
}

// withMarker copies p with the pagination cursor replaced, leaving the
// caller's value untouched across pages.
func (p *ListMetricSeriesParams) withMarker(marker string) *ListMetricSeriesParams {
	var out ListMetricSeriesParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListTraceSpansParams are the optional filters and pagination controls for
// [Client.ListTraceSpans]. A nil *ListTraceSpansParams sends none of them.
type ListTraceSpansParams struct {
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
func (p *ListTraceSpansParams) query() url.Values {
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
func (p *ListTraceSpansParams) withMarker(marker string) *ListTraceSpansParams {
	var out ListTraceSpansParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// QueryMetricsInstantParams are the optional filters and pagination controls for
// [Client.QueryMetricsInstant]. A nil *QueryMetricsInstantParams sends none of them.
type QueryMetricsInstantParams struct {
	// Agg aggregation
	//
	// One of: "avg", "sum", "min", "max", "count", "last", "rate",
	// "increase".
	Agg string

	// By group-by label names
	By []string

	// Match label matchers (e.g. job="api")
	Match []string

	// Metric name
	Metric string

	// Step lookback window duration (e.g. "15s", "5m"). Defaults to 5m.
	Step string

	// Time evaluation timestamp (RFC 3339 or unix seconds). Defaults to now.
	Time string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *QueryMetricsInstantParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.Agg != "" {
		q.Set("agg", p.Agg)
	}
	for _, v := range p.By {
		q.Add("by[]", v)
	}
	for _, v := range p.Match {
		q.Add("match[]", v)
	}
	if p.Metric != "" {
		q.Set("metric", p.Metric)
	}
	if p.Step != "" {
		q.Set("step", p.Step)
	}
	if p.Time != "" {
		q.Set("time", p.Time)
	}
	return q
}

// QueryMetricsRangeParams are the optional filters and pagination controls for
// [Client.QueryMetricsRange]. A nil *QueryMetricsRangeParams sends none of them.
type QueryMetricsRangeParams struct {
	// Agg one of: "avg", "sum", "min", "max", "count", "last", "rate",
	// "increase".
	Agg    string
	By     []string
	End    string
	Match  []string
	Metric string
	Start  string

	// Step bucket width duration (e.g. "15s", "1m")
	Step string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *QueryMetricsRangeParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.Agg != "" {
		q.Set("agg", p.Agg)
	}
	for _, v := range p.By {
		q.Add("by[]", v)
	}
	if p.End != "" {
		q.Set("end", p.End)
	}
	for _, v := range p.Match {
		q.Add("match[]", v)
	}
	if p.Metric != "" {
		q.Set("metric", p.Metric)
	}
	if p.Start != "" {
		q.Set("start", p.Start)
	}
	if p.Step != "" {
		q.Set("step", p.Step)
	}
	return q
}

// SearchLogsParams are the optional filters and pagination controls for
// [Client.SearchLogs]. A nil *SearchLogsParams sends none of them.
type SearchLogsParams struct {
	// From lower bound on timestamp (inclusive)
	From time.Time

	// Limit maximum number of items to return. A value above the maximum is
	// clamped to it rather than rejected, so a page shorter than the one
	// you asked for is normal — page until `meta.has_more` is false, not
	// until a page looks short.
	Limit int

	// LogGroup filter by log-group UUID, CRN or exact name in the authenticated
	// account. CRNs must match the serving region and account handle;
	// invalid references never fall back to names.
	LogGroup string

	// LogStream filter by stream name within the group
	LogStream string

	// Marker opaque pagination cursor. Echo back the `meta.marker` value from the
	// previous page to fetch the next one; do not construct or parse it.
	// The token's internal form varies by endpoint (a resource ID, a
	// timestamp, …) and is not guaranteed stable across releases.
	Marker string

	// MinSeverity filter to records with severity >= this band
	//
	// One of: "TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL".
	MinSeverity string

	// Q case-insensitive substring on the log body
	Q string

	// Region filter by emitter region
	Region string

	// To upper bound on timestamp (exclusive); must be after from and within
	// 31d
	To time.Time

	// TraceID filter to records carrying this W3C trace id
	TraceID string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *SearchLogsParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if !p.From.IsZero() {
		q.Set("from", p.From.UTC().Format(time.RFC3339))
	}
	if p.Limit != 0 {
		q.Set("limit", strconv.Itoa(int(p.Limit)))
	}
	if p.LogGroup != "" {
		q.Set("log_group", p.LogGroup)
	}
	if p.LogStream != "" {
		q.Set("log_stream", p.LogStream)
	}
	if p.Marker != "" {
		q.Set("marker", p.Marker)
	}
	if p.MinSeverity != "" {
		q.Set("min_severity", p.MinSeverity)
	}
	if p.Q != "" {
		q.Set("q", p.Q)
	}
	if p.Region != "" {
		q.Set("region", p.Region)
	}
	if !p.To.IsZero() {
		q.Set("to", p.To.UTC().Format(time.RFC3339))
	}
	if p.TraceID != "" {
		q.Set("trace_id", p.TraceID)
	}
	return q
}

// withMarker copies p with the pagination cursor replaced, leaving the
// caller's value untouched across pages.
func (p *SearchLogsParams) withMarker(marker string) *SearchLogsParams {
	var out SearchLogsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// SearchTracesParams are the optional filters and pagination controls for
// [Client.SearchTraces]. A nil *SearchTracesParams sends none of them.
type SearchTracesParams struct {
	// From lower bound on span start_time (inclusive)
	From time.Time

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

	// MinDurationMs filter to traces whose end-to-end duration is at least this many ms
	MinDurationMs float64

	// Operation filter by span name (operation)
	Operation string

	// Service filter by service.name
	Service string

	// StatusCode filter by status code
	//
	// One of: "UNSET", "OK", "ERROR".
	StatusCode string

	// To upper bound on span start_time (exclusive); must be after from and
	// within 31d
	To time.Time
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *SearchTracesParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if !p.From.IsZero() {
		q.Set("from", p.From.UTC().Format(time.RFC3339))
	}
	if p.Limit != 0 {
		q.Set("limit", strconv.Itoa(int(p.Limit)))
	}
	if p.Marker != "" {
		q.Set("marker", p.Marker)
	}
	if p.MinDurationMs != 0 {
		q.Set("min_duration_ms", strconv.FormatFloat(p.MinDurationMs, 'f', -1, 64))
	}
	if p.Operation != "" {
		q.Set("operation", p.Operation)
	}
	if p.Service != "" {
		q.Set("service", p.Service)
	}
	if p.StatusCode != "" {
		q.Set("status_code", p.StatusCode)
	}
	if !p.To.IsZero() {
		q.Set("to", p.To.UTC().Format(time.RFC3339))
	}
	return q
}

// withMarker copies p with the pagination cursor replaced, leaving the
// caller's value untouched across pages.
func (p *SearchTracesParams) withMarker(marker string) *SearchTracesParams {
	var out SearchTracesParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// CreateAlert creates a metric alert.
//
// Creates a durable metric read and notification publish delegation.
// Requires telemetry:CreateAlert on the named alert,
// telemetry:ReadMetrics for the account, and notifications:Publish on
// the selected topic. Firing and resolved transitions use the topic’s
// confirmed email, webhook and queue subscriptions. Events are stored
// atomically with state changes and retried with stable message IDs.
// Missing data and evaluation errors are visible states, not recovery
// notifications.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateAlert(ctx context.Context, body *CreateAlertRequest, opts ...basaltic.RequestOption) (*Alert, error) {
	op := &basaltic.Operation{
		ID:     "createAlert",
		Method: "POST",
		Path:   "/v1/alerts",
		Body:   body,
	}
	var out struct {
		Alert *Alert `json:"alert"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Alert, nil
}

// CreateLogGroup creates a log group.
//
// Register a new log group in the caller's org. Name is the
// customer-visible identifier (1–512 chars of A-Za-z0-9_./#-) and is
// immutable. Requires `telemetry:CreateLogGroup`.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateLogGroup(ctx context.Context, body *CreateLogGroupRequest, opts ...basaltic.RequestOption) (*LogGroup, error) {
	op := &basaltic.Operation{
		ID:     "createLogGroup",
		Method: "POST",
		Path:   "/v1/log-groups",
		Body:   body,
	}
	var out struct {
		LogGroup *LogGroup `json:"log_group"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.LogGroup, nil
}

// DeleteAlert deletes a metric alert.
//
// Deletes the alert and its unpublished notification events. Already
// accepted notification deliveries retain the Notifications service
// lifecycle.
func (c *Client) DeleteAlert(ctx context.Context, id string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteAlert",
		Method:   "DELETE",
		Path:     "/v1/alerts/{id}",
		PathArgs: []string{id},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteLogGroup deletes a log group.
//
// Removes the group's admin row. Existing log records keep their
// embedded log_group_id reference but the joined name no longer
// resolves. The log store's retention policy still expires them on their
// own schedule. Requires `telemetry:DeleteLogGroup`.
func (c *Client) DeleteLogGroup(ctx context.Context, id string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteLogGroup",
		Method:   "DELETE",
		Path:     "/v1/log-groups/{id}",
		PathArgs: []string{id},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteTraceSettings deletes trace settings.
//
// Restore implicit defaults. Requires telemetry:PutTraceSettings. Stored
// data retains its original retention and encryption.
func (c *Client) DeleteTraceSettings(ctx context.Context, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:     "deleteTraceSettings",
		Method: "DELETE",
		Path:   "/v1/trace-settings",
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// GetAlert gets a metric alert.
func (c *Client) GetAlert(ctx context.Context, id string, opts ...basaltic.RequestOption) (*Alert, error) {
	op := &basaltic.Operation{
		ID:       "getAlert",
		Method:   "GET",
		Path:     "/v1/alerts/{id}",
		PathArgs: []string{id},
	}
	var out struct {
		Alert *Alert `json:"alert"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Alert, nil
}

// GetLog gets a single log record by id.
func (c *Client) GetLog(ctx context.Context, logID string, opts ...basaltic.RequestOption) (*LogRecord, error) {
	op := &basaltic.Operation{
		ID:       "getLog",
		Method:   "GET",
		Path:     "/v1/logs/{log_id}",
		PathArgs: []string{logID},
	}
	var out struct {
		Log *LogRecord `json:"log"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Log, nil
}

// GetLogGroup gets a log group by id.
func (c *Client) GetLogGroup(ctx context.Context, id string, opts ...basaltic.RequestOption) (*LogGroup, error) {
	op := &basaltic.Operation{
		ID:       "getLogGroup",
		Method:   "GET",
		Path:     "/v1/log-groups/{id}",
		PathArgs: []string{id},
	}
	var out struct {
		LogGroup *LogGroup `json:"log_group"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.LogGroup, nil
}

// GetRetainedTelemetryPresence — Check retained telemetry presence.
//
// Reports whether logs, spans, metrics or retention rollups remain in
// this account and region, including expired data awaiting reclamation.
// The blocks_deletion flag distinguishes log/span storage and encryption
// dependencies from metric history, which can expire normally after
// account soft-deletion. Requires telemetry:GetRetainedData.
func (c *Client) GetRetainedTelemetryPresence(ctx context.Context, opts ...basaltic.RequestOption) (*GetRetainedTelemetryPresenceRetainedData, error) {
	op := &basaltic.Operation{
		ID:     "getRetainedTelemetryPresence",
		Method: "GET",
		Path:   "/v1/retained-data",
	}
	var out struct {
		RetainedData *GetRetainedTelemetryPresenceRetainedData `json:"retained_data"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.RetainedData, nil
}

// GetTrace gets a trace.
//
// Returns the same trace summary as the trace list. Requires
// telemetry:ReadTraces.
func (c *Client) GetTrace(ctx context.Context, traceID string, opts ...basaltic.RequestOption) (*TraceSummary, error) {
	op := &basaltic.Operation{
		ID:       "getTrace",
		Method:   "GET",
		Path:     "/v1/traces/{trace_id}",
		PathArgs: []string{traceID},
	}
	var out struct {
		Trace *TraceSummary `json:"trace"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Trace, nil
}

// GetTraceSettings gets the caller account's trace settings.
//
// Returns effective trace retention and encryption settings for the
// active account. Defaults to 30 days and no KMS key. Requires
// telemetry:GetTraceSettings.
func (c *Client) GetTraceSettings(ctx context.Context, opts ...basaltic.RequestOption) (*TraceSettings, error) {
	op := &basaltic.Operation{
		ID:     "getTraceSettings",
		Method: "GET",
		Path:   "/v1/trace-settings",
	}
	var out struct {
		TraceSettings *TraceSettings `json:"trace_settings"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.TraceSettings, nil
}

// IngestLogs ingests a batch of log records.
//
// Append one or more log records to the org's telemetry store.
// Per-record validation failures are reported in the response body's
// `rejected` field; the request itself returns 202 whenever the batch
// was accepted at the wire level. Requires `telemetry:WriteLogs`
// permission.
func (c *Client) IngestLogs(ctx context.Context, body *IngestRequest, opts ...basaltic.RequestOption) (*IngestResult, error) {
	op := &basaltic.Operation{
		ID:     "ingestLogs",
		Method: "POST",
		Path:   "/v1/logs",
		Body:   body,
	}
	var out struct {
		Ingestion *IngestResult `json:"ingestion"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Ingestion, nil
}

// IngestSpans ingests a batch of trace spans.
//
// Append one or more spans to the org's trace store. Per-record
// validation failures are reported in the response body's `rejected`
// field; the batch keeps flowing. Requires `telemetry:WriteSpans`.
func (c *Client) IngestSpans(ctx context.Context, body *IngestSpansRequest, opts ...basaltic.RequestOption) (*IngestResult, error) {
	op := &basaltic.Operation{
		ID:     "ingestSpans",
		Method: "POST",
		Path:   "/v1/spans",
		Body:   body,
	}
	var out struct {
		Ingestion *IngestResult `json:"ingestion"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Ingestion, nil
}

// ListAlerts lists metric alerts.
//
// Account-scoped keyset pagination by UUID, ascending. Markers are
// opaque and bound to the account and filters. No tenant identity is
// included in the response.
//
// Returns one page. Use ListAlertsAll to walk every page.
func (c *Client) ListAlerts(ctx context.Context, params *ListAlertsParams, opts ...basaltic.RequestOption) (*basaltic.Page[Alert], error) {
	op := &basaltic.Operation{
		ID:     "listAlerts",
		Method: "GET",
		Path:   "/v1/alerts",
	}
	op.Query = params.query()
	var out struct {
		Items []Alert `json:"alerts"`
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
	page := &basaltic.Page[Alert]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListAlertsAll walks every page of ListAlerts, yielding one item at a
// time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListAlertsAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListAlertsAll(ctx context.Context, params *ListAlertsParams, opts ...basaltic.RequestOption) iter.Seq2[Alert, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Alert], error) {
		return c.ListAlerts(ctx, params.withMarker(marker), opts...)
	})
}

// ListLogGroups lists log groups (or look up one by name).
//
// Without `name`: returns the account's log groups, newest-first,
// keyset-paginated. With `name`: returns at most one log group matching
// that exact name (empty list if no match). Requires
// `telemetry:DescribeLogGroups`.
//
// Returns one page. Use ListLogGroupsAll to walk every page.
func (c *Client) ListLogGroups(ctx context.Context, params *ListLogGroupsParams, opts ...basaltic.RequestOption) (*basaltic.Page[LogGroup], error) {
	op := &basaltic.Operation{
		ID:     "listLogGroups",
		Method: "GET",
		Path:   "/v1/log-groups",
	}
	op.Query = params.query()
	var out struct {
		Items []LogGroup `json:"log_groups"`
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
	page := &basaltic.Page[LogGroup]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListLogGroupsAll walks every page of ListLogGroups, yielding one item
// at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListLogGroupsAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListLogGroupsAll(ctx context.Context, params *ListLogGroupsParams, opts ...basaltic.RequestOption) iter.Seq2[LogGroup, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[LogGroup], error) {
		return c.ListLogGroups(ctx, params.withMarker(marker), opts...)
	})
}

// ListMetricNames lists metric names.
//
// Returns metric names in alphabetical order for the time window. Follow
// meta.marker while meta.has_more is true. Requires
// telemetry:ReadMetrics.
//
// Returns one page. Use ListMetricNamesAll to walk every page.
func (c *Client) ListMetricNames(ctx context.Context, params *ListMetricNamesParams, opts ...basaltic.RequestOption) (*basaltic.Page[string], error) {
	op := &basaltic.Operation{
		ID:     "listMetricNames",
		Method: "GET",
		Path:   "/v1/metrics/names",
	}
	op.Query = params.query()
	var out struct {
		Items []string `json:"metric_names"`
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
	page := &basaltic.Page[string]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListMetricNamesAll walks every page of ListMetricNames, yielding one
// item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListMetricNamesAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListMetricNamesAll(ctx context.Context, params *ListMetricNamesParams, opts ...basaltic.RequestOption) iter.Seq2[string, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[string], error) {
		return c.ListMetricNames(ctx, params.withMarker(marker), opts...)
	})
}

// ListMetricNamesPost lists metric names (JSON).
//
// Same operation as GET, with a JSON body. Arrays use match and by
// instead of repeated query parameters.
func (c *Client) ListMetricNamesPost(ctx context.Context, body *ListMetricNamesPostRequest, opts ...basaltic.RequestOption) (*basaltic.Page[string], error) {
	op := &basaltic.Operation{
		ID:     "listMetricNamesPost",
		Method: "POST",
		Path:   "/v1/metrics/names",
		Body:   body,
	}
	var out struct {
		Items []string `json:"metric_names"`
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
	page := &basaltic.Page[string]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListMetricSeries lists metric series.
//
// Returns distinct label sets for one metric in the time window. Follow
// meta.marker while meta.has_more is true. Requires
// telemetry:ReadMetrics.
//
// Returns one page. Use ListMetricSeriesAll to walk every page.
func (c *Client) ListMetricSeries(ctx context.Context, params *ListMetricSeriesParams, opts ...basaltic.RequestOption) (*basaltic.Page[map[string]string], error) {
	op := &basaltic.Operation{
		ID:     "listMetricSeries",
		Method: "GET",
		Path:   "/v1/metrics/series",
	}
	op.Query = params.query()
	var out struct {
		Items []map[string]string `json:"series"`
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
	page := &basaltic.Page[map[string]string]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListMetricSeriesAll walks every page of ListMetricSeries, yielding one
// item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListMetricSeriesAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListMetricSeriesAll(ctx context.Context, params *ListMetricSeriesParams, opts ...basaltic.RequestOption) iter.Seq2[map[string]string, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[map[string]string], error) {
		return c.ListMetricSeries(ctx, params.withMarker(marker), opts...)
	})
}

// ListMetricSeriesPost lists metric series (JSON).
//
// Same operation as GET, with a JSON body. Arrays use match and by
// instead of repeated query parameters.
func (c *Client) ListMetricSeriesPost(ctx context.Context, body *ListMetricSeriesPostRequest, opts ...basaltic.RequestOption) (*basaltic.Page[map[string]string], error) {
	op := &basaltic.Operation{
		ID:     "listMetricSeriesPost",
		Method: "POST",
		Path:   "/v1/metrics/series",
		Body:   body,
	}
	var out struct {
		Items []map[string]string `json:"series"`
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
	page := &basaltic.Page[map[string]string]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListTraceSpans lists spans in a trace.
//
// Returns spans ordered by start_time and span_id, with cursor
// pagination. Requires telemetry:ReadTraces.
//
// Returns one page. Use ListTraceSpansAll to walk every page.
func (c *Client) ListTraceSpans(ctx context.Context, traceID string, params *ListTraceSpansParams, opts ...basaltic.RequestOption) (*basaltic.Page[Span], error) {
	op := &basaltic.Operation{
		ID:       "listTraceSpans",
		Method:   "GET",
		Path:     "/v1/traces/{trace_id}/spans",
		PathArgs: []string{traceID},
	}
	op.Query = params.query()
	var out struct {
		Items []Span `json:"spans"`
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
	page := &basaltic.Page[Span]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListTraceSpansAll walks every page of ListTraceSpans, yielding one
// item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListTraceSpansAll(ctx, traceID, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListTraceSpansAll(ctx context.Context, traceID string, params *ListTraceSpansParams, opts ...basaltic.RequestOption) iter.Seq2[Span, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Span], error) {
		return c.ListTraceSpans(ctx, traceID, params.withMarker(marker), opts...)
	})
}

// QueryMetricsInstant — Run an instant metric query.
//
// Returns the latest aggregate point per series. Select metric, optional
// agg (default avg), match and by labels. Default lookback is 5m.
// Requires telemetry:ReadMetrics.
func (c *Client) QueryMetricsInstant(ctx context.Context, params *QueryMetricsInstantParams, opts ...basaltic.RequestOption) (*MetricQueryResult, error) {
	op := &basaltic.Operation{
		ID:     "queryMetricsInstant",
		Method: "GET",
		Path:   "/v1/metrics/query",
	}
	op.Query = params.query()
	var out struct {
		QueryResult *MetricQueryResult `json:"query_result"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.QueryResult, nil
}

// QueryMetricsInstantPost — Run an instant metric query (JSON).
//
// Same operation as GET, with a JSON body. Arrays use match and by
// instead of repeated query parameters.
func (c *Client) QueryMetricsInstantPost(ctx context.Context, body *QueryMetricsInstantPostRequest, opts ...basaltic.RequestOption) (*MetricQueryResult, error) {
	op := &basaltic.Operation{
		ID:     "queryMetricsInstantPost",
		Method: "POST",
		Path:   "/v1/metrics/query",
		Body:   body,
	}
	var out struct {
		QueryResult *MetricQueryResult `json:"query_result"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.QueryResult, nil
}

// QueryMetricsRange — Run a range metric query.
//
// Returns aggregate points per series between start and end. The window
// is limited to 31 days and 11000 buckets per series; step defaults to
// 60s. Requires telemetry:ReadMetrics.
func (c *Client) QueryMetricsRange(ctx context.Context, params *QueryMetricsRangeParams, opts ...basaltic.RequestOption) (*MetricQueryResult, error) {
	op := &basaltic.Operation{
		ID:     "queryMetricsRange",
		Method: "GET",
		Path:   "/v1/metrics/query-range",
	}
	op.Query = params.query()
	var out struct {
		QueryResult *MetricQueryResult `json:"query_result"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.QueryResult, nil
}

// QueryMetricsRangePost — Run a range metric query (JSON).
//
// Same operation as GET, with a JSON body. Arrays use match and by
// instead of repeated query parameters.
func (c *Client) QueryMetricsRangePost(ctx context.Context, body *QueryMetricsRangePostRequest, opts ...basaltic.RequestOption) (*MetricQueryResult, error) {
	op := &basaltic.Operation{
		ID:     "queryMetricsRangePost",
		Method: "POST",
		Path:   "/v1/metrics/query-range",
		Body:   body,
	}
	var out struct {
		QueryResult *MetricQueryResult `json:"query_result"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.QueryResult, nil
}

// SearchLogs searches log records.
//
// Returns logs for the current organization, newest-first, with keyset
// pagination on (timestamp, id). Free-text body search is
// case-insensitive substring; per-attribute filters use the
// `attr.<key>=<value>` query-param convention. `from` and `to` are
// required; the window must be at most 31 days. Requires
// `telemetry:ReadLogs` permission.
//
// Returns one page. Use SearchLogsAll to walk every page.
func (c *Client) SearchLogs(ctx context.Context, params *SearchLogsParams, opts ...basaltic.RequestOption) (*basaltic.Page[LogRecord], error) {
	op := &basaltic.Operation{
		ID:     "searchLogs",
		Method: "GET",
		Path:   "/v1/logs",
	}
	op.Query = params.query()
	var out struct {
		Items []LogRecord `json:"logs"`
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
	page := &basaltic.Page[LogRecord]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// SearchLogsAll walks every page of SearchLogs, yielding one item at a
// time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.SearchLogsAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) SearchLogsAll(ctx context.Context, params *SearchLogsParams, opts ...basaltic.RequestOption) iter.Seq2[LogRecord, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[LogRecord], error) {
		return c.SearchLogs(ctx, params.withMarker(marker), opts...)
	})
}

// SearchTraces lists traces.
//
// Returns one summary per distinct trace_id, newest-first. Keyset
// pagination on (start_time, trace_id). `from` and `to` are required;
// the window must be at most 31 days. Requires `telemetry:ReadTraces`.
//
// Returns one page. Use SearchTracesAll to walk every page.
func (c *Client) SearchTraces(ctx context.Context, params *SearchTracesParams, opts ...basaltic.RequestOption) (*basaltic.Page[TraceSummary], error) {
	op := &basaltic.Operation{
		ID:     "searchTraces",
		Method: "GET",
		Path:   "/v1/traces",
	}
	op.Query = params.query()
	var out struct {
		Items []TraceSummary `json:"traces"`
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
	page := &basaltic.Page[TraceSummary]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// SearchTracesAll walks every page of SearchTraces, yielding one item at
// a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.SearchTracesAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) SearchTracesAll(ctx context.Context, params *SearchTracesParams, opts ...basaltic.RequestOption) iter.Seq2[TraceSummary, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[TraceSummary], error) {
		return c.SearchTraces(ctx, params.withMarker(marker), opts...)
	})
}

// UpdateAlert updates a metric alert.
//
// Requires telemetry:UpdateAlert. Changing the rule/topic or enabling an
// alert additionally renews telemetry:ReadMetrics and
// notifications:Publish authorization. enabled=false stops evaluation;
// silenced_until mutes notification delivery while evaluation continues.
// Changing the rule or enabled flag resets pending/incident state. Name
// is immutable.
func (c *Client) UpdateAlert(ctx context.Context, id string, body *UpdateAlertRequest, opts ...basaltic.RequestOption) (*Alert, error) {
	op := &basaltic.Operation{
		ID:       "updateAlert",
		Method:   "PATCH",
		Path:     "/v1/alerts/{id}",
		PathArgs: []string{id},
		Body:     body,
	}
	var out struct {
		Alert *Alert `json:"alert"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Alert, nil
}

// UpdateLogGroup updates a log group.
//
// Mutate description, retention policy, KMS association, or tags. Name
// is immutable (rename = recreate). Requires `telemetry:UpdateLogGroup`.
func (c *Client) UpdateLogGroup(ctx context.Context, id string, body *UpdateLogGroupRequest, opts ...basaltic.RequestOption) (*LogGroup, error) {
	op := &basaltic.Operation{
		ID:       "updateLogGroup",
		Method:   "PATCH",
		Path:     "/v1/log-groups/{id}",
		PathArgs: []string{id},
		Body:     body,
	}
	var out struct {
		LogGroup *LogGroup `json:"log_group"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.LogGroup, nil
}

// UpdateTraceSettings updates the caller account's trace settings.
//
// Updates supplied fields only. Null retention restores 30 days; null
// kms_key clears encryption for future spans. Stored spans keep their
// original retention and encryption. Requires
// telemetry:PutTraceSettings.
func (c *Client) UpdateTraceSettings(ctx context.Context, body *UpdateTraceSettingsRequest, opts ...basaltic.RequestOption) (*TraceSettings, error) {
	op := &basaltic.Operation{
		ID:     "updateTraceSettings",
		Method: "PATCH",
		Path:   "/v1/trace-settings",
		Body:   body,
	}
	var out struct {
		TraceSettings *TraceSettings `json:"trace_settings"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.TraceSettings, nil
}

// WriteMetrics — Prometheus remote_write ingest.
//
// Accepts a snappy-compressed protobuf WriteRequest body, the canonical
// Prometheus remote_write wire format. The service decodes, stamps
// `basaltic_account_id` + `basaltic_org_id` labels on every timeseries,
// and persists samples in ClickHouse. Caller-supplied tenant labels are
// stripped before stamping — producers can't spoof their account.
//
// Requires `telemetry:WriteMetrics`. Returns 204 on success (Prometheus
// spec); 4xx with a JSON error envelope otherwise.
func (c *Client) WriteMetrics(ctx context.Context, body io.Reader, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:          "writeMetrics",
		Method:      "POST",
		Path:        "/v1/metrics/write",
		Stream:      body,
		ContentType: "application/x-protobuf",
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// GetAlertByReference fetches one alert by an id, a CRN or a name.
//
// The reference is classified by its syntax alone, exactly as the
// platform does (see [basaltic.ParseReference]): an id is fetched with
// [Client.GetAlert]; a CRN or a name goes to [Client.ListAlerts] as an
// exact filter, together with any filters already set on scope, which
// may be nil. A miss is a not-found error for the kind the string was
// read as — no other kind is tried — and more than one match is a
// [basaltic.AmbiguousReferenceError].
func (c *Client) GetAlertByReference(ctx context.Context, ref string, scope *ListAlertsParams, opts ...basaltic.RequestOption) (*Alert, error) {
	return basaltic.ResolveByReference(ctx, ref, "alert", "listAlerts", true,
		func(ctx context.Context, refID string) (*Alert, error) {
			return c.GetAlert(ctx, refID, opts...)
		},
		func(ctx context.Context, refName, refCRN string) (*basaltic.Page[Alert], error) {
			var p ListAlertsParams
			if scope != nil {
				p = *scope
			}
			p.Name = refName
			p.CRN = refCRN
			p.Limit = 2
			return c.ListAlerts(ctx, &p, opts...)
		})
}

// GetLogGroupByReference fetches one log group by an id, a CRN or a
// name.
//
// The reference is classified by its syntax alone, exactly as the
// platform does (see [basaltic.ParseReference]): an id is fetched with
// [Client.GetLogGroup]; a CRN or a name goes to [Client.ListLogGroups]
// as an exact filter, together with any filters already set on scope,
// which may be nil. A miss is a not-found error for the kind the string
// was read as — no other kind is tried — and more than one match is
// a [basaltic.AmbiguousReferenceError].
func (c *Client) GetLogGroupByReference(ctx context.Context, ref string, scope *ListLogGroupsParams, opts ...basaltic.RequestOption) (*LogGroup, error) {
	return basaltic.ResolveByReference(ctx, ref, "log-group", "listLogGroups", true,
		func(ctx context.Context, refID string) (*LogGroup, error) {
			return c.GetLogGroup(ctx, refID, opts...)
		},
		func(ctx context.Context, refName, refCRN string) (*basaltic.Page[LogGroup], error) {
			var p ListLogGroupsParams
			if scope != nil {
				p = *scope
			}
			p.Name = refName
			p.CRN = refCRN
			p.Limit = 2
			return c.ListLogGroups(ctx, &p, opts...)
		})
}
