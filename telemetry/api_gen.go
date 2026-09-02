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

// ListLogGroupsParams are the optional filters and pagination controls for
// [Client.ListLogGroups]. A nil *ListLogGroupsParams sends none of them.
type ListLogGroupsParams struct {
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

	// Name exact-match name lookup (single-result shortcut)
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListLogGroupsParams) query() url.Values {
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
	End   string
	Start string
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
	if p.Start != "" {
		q.Set("start", p.Start)
	}
	return q
}

// ListMetricSeriesParams are the optional filters and pagination controls for
// [Client.ListMetricSeries]. A nil *ListMetricSeriesParams sends none of them.
type ListMetricSeriesParams struct {
	End    string
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
	if p.Metric != "" {
		q.Set("metric", p.Metric)
	}
	if p.Start != "" {
		q.Set("start", p.Start)
	}
	return q
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

	// LogGroup filter by parent log group name
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

// GetTrace gets all spans for a trace.
//
// Returns every span belonging to `trace_id`, ordered by `start_time`
// ASC for direct waterfall rendering. Requires `telemetry:ReadTraces`.
func (c *Client) GetTrace(ctx context.Context, traceID string, opts ...basaltic.RequestOption) (*TraceResponse, error) {
	op := &basaltic.Operation{
		ID:       "getTrace",
		Method:   "GET",
		Path:     "/v1/traces/{trace_id}",
		PathArgs: []string{traceID},
	}
	var out TraceResponse
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetTraceSettings gets the caller account's trace settings.
//
// Returns the trace retention + KMS association for the active
// X-Account-Id. When no settings row has been written for the account
// yet, the response carries the in-memory defaults (retention_days=null
// → never-expire is *not* the default; ingest stamps 30 days when both
// DB row and PUT body are empty). Requires `telemetry:GetTraceSettings`.
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
	var out IngestResult
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
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
	var out IngestResult
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListLogGroups lists log groups (or look up one by name).
//
// Without `name`: returns the org's log groups, newest-first,
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

// ListMetricNames lists the distinct metric names emitted in a time window.
//
// Returns the distinct metric names the tenant has emitted in [start,
// end] — the metric explorer's catalogue. GET + POST are equivalent;
// POST lets large param sets ride the body. Requires
// `telemetry:ReadMetrics`.
func (c *Client) ListMetricNames(ctx context.Context, params *ListMetricNamesParams, opts ...basaltic.RequestOption) (*ListMetricNamesResult, error) {
	op := &basaltic.Operation{
		ID:     "listMetricNames",
		Method: "GET",
		Path:   "/v1/metrics/names",
	}
	op.Query = params.query()
	var out ListMetricNamesResult
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListMetricNamesPost lists the distinct metric names emitted in a time window (form body).
//
// Same as GET but accepts form-encoded body.
func (c *Client) ListMetricNamesPost(ctx context.Context, body io.Reader, opts ...basaltic.RequestOption) (*ListMetricNamesPostResult, error) {
	op := &basaltic.Operation{
		ID:          "listMetricNamesPost",
		Method:      "POST",
		Path:        "/v1/metrics/names",
		Stream:      body,
		ContentType: "application/x-www-form-urlencoded",
	}
	var out ListMetricNamesPostResult
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListMetricSeries lists distinct label sets for a metric.
//
// Returns distinct label sets for `metric` in [start, end] — the
// discovery surface for custom dashboards. Required: `metric`, `start`,
// `end`.
//
// Requires `telemetry:ReadMetrics`.
func (c *Client) ListMetricSeries(ctx context.Context, params *ListMetricSeriesParams, opts ...basaltic.RequestOption) (map[string]any, error) {
	op := &basaltic.Operation{
		ID:     "listMetricSeries",
		Method: "GET",
		Path:   "/v1/metrics/series",
	}
	op.Query = params.query()
	var out map[string]any
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out, nil
}

// ListMetricSeriesPost lists distinct label sets for a metric (form body).
//
// Same as GET but accepts form-encoded body.
func (c *Client) ListMetricSeriesPost(ctx context.Context, body io.Reader, opts ...basaltic.RequestOption) (map[string]any, error) {
	op := &basaltic.Operation{
		ID:          "listMetricSeriesPost",
		Method:      "POST",
		Path:        "/v1/metrics/series",
		Stream:      body,
		ContentType: "application/x-www-form-urlencoded",
	}
	var out map[string]any
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out, nil
}

// PutTraceSettings updates the caller account's trace settings.
//
// Upserts the trace retention + KMS key. PUT semantics — the body is
// the full intent, fields not supplied default to unchanged-from-empty
// (i.e. retention_days=null means never-expire, kms_key_crn omitted
// means plaintext at-rest). Set `clear_retention=true` to switch to
// never-expire when the row already has a bounded value. Requires
// `telemetry:PutTraceSettings`.
func (c *Client) PutTraceSettings(ctx context.Context, body *UpdateTraceSettingsRequest, opts ...basaltic.RequestOption) (*TraceSettings, error) {
	op := &basaltic.Operation{
		ID:     "putTraceSettings",
		Method: "PUT",
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

// QueryMetricsInstant — Instant structured metric query.
//
// Instant query over ClickHouse. Required: `metric` (name) and `agg`
// (avg|sum|min|max|count|last|rate|increase). Optional `match[]` label
// matchers, `by[]` group-by labels, `step` (lookback window, default
// 5m), and `time` (evaluation instant). Returns a Prometheus-compatible
// vector envelope.
//
// Requires `telemetry:ReadMetrics`.
func (c *Client) QueryMetricsInstant(ctx context.Context, params *QueryMetricsInstantParams, opts ...basaltic.RequestOption) (map[string]any, error) {
	op := &basaltic.Operation{
		ID:     "queryMetricsInstant",
		Method: "GET",
		Path:   "/v1/metrics/query",
	}
	op.Query = params.query()
	var out map[string]any
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out, nil
}

// QueryMetricsInstantPost — Instant structured metric query (form body).
//
// Same as GET but accepts form-encoded body for long matcher sets.
func (c *Client) QueryMetricsInstantPost(ctx context.Context, body io.Reader, opts ...basaltic.RequestOption) (map[string]any, error) {
	op := &basaltic.Operation{
		ID:          "queryMetricsInstantPost",
		Method:      "POST",
		Path:        "/v1/metrics/query",
		Stream:      body,
		ContentType: "application/x-www-form-urlencoded",
	}
	var out map[string]any
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out, nil
}

// QueryMetricsRange — Range structured metric query.
//
// Range variant returning a Prometheus matrix. Required: `metric`,
// `agg`, `start`, `end`. Optional `match[]`, `by[]`, `step` (bucket
// width; repository defaults missing step to 60s). Window max 31d;
// result capped at 11000 points.
//
// Requires `telemetry:ReadMetrics`.
func (c *Client) QueryMetricsRange(ctx context.Context, params *QueryMetricsRangeParams, opts ...basaltic.RequestOption) (map[string]any, error) {
	op := &basaltic.Operation{
		ID:     "queryMetricsRange",
		Method: "GET",
		Path:   "/v1/metrics/query_range",
	}
	op.Query = params.query()
	var out map[string]any
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out, nil
}

// QueryMetricsRangePost — Range structured metric query (form body).
//
// Same as GET but accepts form-encoded body for long matcher sets.
func (c *Client) QueryMetricsRangePost(ctx context.Context, body io.Reader, opts ...basaltic.RequestOption) (map[string]any, error) {
	op := &basaltic.Operation{
		ID:          "queryMetricsRangePost",
		Method:      "POST",
		Path:        "/v1/metrics/query_range",
		Stream:      body,
		ContentType: "application/x-www-form-urlencoded",
	}
	var out map[string]any
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out, nil
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
