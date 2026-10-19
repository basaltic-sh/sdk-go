// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package telemetry

import (
	"time"
)

type Alert struct {
	CreatedAt time.Time `json:"created_at"`
	CRN       string    `json:"crn"`

	// DeliveryError notification remains queued for retry; metric evaluation continues.
	DeliveryError   string           `json:"delivery_error,omitempty"`
	Enabled         bool             `json:"enabled"`
	EvaluatedAt     time.Time        `json:"evaluated_at,omitempty"`
	EvaluationError string           `json:"evaluation_error,omitempty"`
	Expression      *AlertExpression `json:"expression"`

	// ForSeconds continuous breach duration before firing; 0 fires immediately.
	// Missing data, query failures and scheduling gaps reset a pending
	// duration.
	ForSeconds int      `json:"for_seconds"`
	ID         string   `json:"id"`
	LastValue  *float64 `json:"last_value"`
	Name       string   `json:"name"`

	// One of: "gt", "gte", "lt", "lte".
	Operator     string    `json:"operator"`
	PendingSince time.Time `json:"pending_since,omitempty"`

	// SilencedUntil mute notifications until this instant (at most 30 days ahead).
	// Evaluation continues; null removes silence. An active alert notifies
	// on the next evaluation after silence expires.
	SilencedUntil *time.Time `json:"silenced_until"`

	// State evaluation state, independent of enabled and silenced_until. Missing
	// data/errors do not resolve an open incident.
	//
	// One of: "ok", "pending", "firing", "no_data", "error".
	State          string       `json:"state"`
	StateChangedAt time.Time    `json:"state_changed_at,omitempty"`
	Threshold      float64      `json:"threshold"`
	Topic          *ResourceRef `json:"topic"`
	UpdatedAt      time.Time    `json:"updated_at"`

	// WindowSeconds whole-minute window. Evaluated every minute against the last
	// complete UTC-aligned window.
	WindowSeconds int `json:"window_seconds"`
}

type AlertExpression struct {
	// One of: "avg", "min", "max", "sum", "last", "count", "rate", "increase".
	Agg string `json:"agg"`

	// By grouping labels; the resulting series are summed to match the
	// console metric card.
	By          []string    `json:"by,omitempty"`
	Denominator *AlertQuery `json:"denominator,omitempty"`

	// Match label matchers (label=value, label!=value, label=~regex,
	// label!~regex). Tenant scope is always applied independently.
	Match  []string `json:"match,omitempty"`
	Metric string   `json:"metric"`

	// Scale multiplier after optional division, e.g. 100 for a percentage or
	// 1/1048576 for MiB/s.
	Scale float64 `json:"scale"`

	// Unit display unit for the scaled threshold and value.
	Unit *string `json:"unit,omitempty"`
}

type AlertQuery struct {
	// One of: "avg", "min", "max", "sum", "last", "count", "rate", "increase".
	Agg string `json:"agg"`

	// By grouping labels; the resulting series are summed to match the
	// console metric card.
	By []string `json:"by,omitempty"`

	// Match label matchers (label=value, label!=value, label=~regex,
	// label!~regex). Tenant scope is always applied independently.
	Match  []string `json:"match,omitempty"`
	Metric string   `json:"metric"`
}

type CreateAlertRequest struct {
	Enabled *bool `json:"enabled,omitempty"`

	// Required.
	Expression *AlertExpression `json:"expression"`

	// ForSeconds continuous breach duration before firing; 0 fires immediately.
	// Missing data, query failures and scheduling gaps reset a pending
	// duration.
	//
	// Required.
	ForSeconds int `json:"for_seconds"`

	// Name immutable account-unique name. UUID names and the crn: prefix are
	// reserved.
	//
	// Required.
	Name string `json:"name"`

	// One of: "gt", "gte", "lt", "lte".
	//
	// Required.
	Operator string `json:"operator"`

	// SilencedUntil mute notifications until this instant (at most 30 days ahead).
	// Evaluation continues; null removes silence. An active alert notifies
	// on the next evaluation after silence expires.
	SilencedUntil *time.Time `json:"silenced_until,omitempty"`

	// Required.
	Threshold float64 `json:"threshold"`

	// Topic required same-account, same-region notification topic reference:
	// UUID, canonical CRN or exact name, classified by syntax. Multiple
	// alerts may share a topic. Binding requires notifications:Publish and
	// respects explicit topic-policy denies.
	//
	// Required.
	Topic string `json:"topic"`

	// WindowSeconds whole-minute window. Evaluated every minute against the last
	// complete UTC-aligned window.
	//
	// Required.
	WindowSeconds int `json:"window_seconds"`
}

type CreateLogGroupRequest struct {
	Description *string `json:"description,omitempty"`

	// KMSKey KMS key UUID, CRN or exact name in the authenticated account and
	// serving region. The resolved UUID is pinned; deleting a key and
	// reusing its name never retargets existing data. An empty string
	// selects plaintext.
	KMSKey *string `json:"kms_key,omitempty"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	//
	// Required.
	Name string `json:"name"`

	// RetentionDays retention in days. Omit or use null for no expiry.
	//
	// One of: "1", "3", "5", "7", "14", "30", "60", "90", "120", "150", "180", "365", "400", "545", "731", "1096", "1827", "2192", "2557", "2922", "3288", "3653", "<nil>".
	RetentionDays *int              `json:"retention_days,omitempty"`
	Tags          map[string]string `json:"tags,omitempty"`
}

type GetRetainedTelemetryPresenceRetainedData struct {
	// BlocksDeletion true when retained logs, spans or their storage rollups still block
	// account deletion. Metric samples alone do not block account
	// soft-deletion and expire through normal retention. Clients talking
	// to older servers that omit this field must fall back to has_data.
	BlocksDeletion bool `json:"blocks_deletion,omitempty"`
	HasData        bool `json:"has_data"`
}

type IngestError struct {
	Code string `json:"code"`

	// Index zero-based input record index.
	Index   int    `json:"index"`
	Message string `json:"message"`
}

type IngestRecord struct {
	Attributes map[string]string `json:"attributes,omitempty"`

	// Required.
	Body string `json:"body"`

	// LogGroup reference to an existing log group by UUID, CRN or exact name in the
	// authenticated account. CRNs must identify telemetry/log-group in the
	// serving region and account handle. Classification is by syntax with
	// no lookup fallback; slash-bearing names are preserved.
	//
	// Required.
	LogGroup string `json:"log_group"`

	// LogStream within-group stream name (free-form, customer choice)
	//
	// Required.
	LogStream string            `json:"log_stream"`
	Resource  map[string]string `json:"resource,omitempty"`

	// Severity one of TRACE/DEBUG/INFO/WARN/ERROR/FATAL (defaults to INFO)
	Severity *string `json:"severity,omitempty"`
	SpanID   *string `json:"span_id,omitempty"`

	// Timestamp optional — defaults to ingest time
	Timestamp *time.Time `json:"timestamp,omitempty"`
	TraceID   *string    `json:"trace_id,omitempty"`
}

type IngestRequest struct {
	// Required.
	Logs []*IngestRecord `json:"logs"`
}

type IngestResult struct {
	Accepted int            `json:"accepted"`
	Errors   []*IngestError `json:"errors,omitempty"`
	Rejected int            `json:"rejected"`
}

type IngestSpansRequest struct {
	// Required.
	Spans []*SpanIngestRecord `json:"spans"`
}

type ListMetricNamesPostRequest struct {
	// Required.
	End    string  `json:"end"`
	Limit  *int    `json:"limit,omitempty"`
	Marker *string `json:"marker,omitempty"`

	// Required.
	Start string `json:"start"`
}

type ListMetricSeriesPostRequest struct {
	// Required.
	End    string  `json:"end"`
	Limit  *int    `json:"limit,omitempty"`
	Marker *string `json:"marker,omitempty"`

	// Required.
	Metric string `json:"metric"`

	// Required.
	Start string `json:"start"`
}

type LogGroup struct {
	CreatedAt time.Time `json:"created_at"`

	// CRN of the log group (used as the IAM resource ARN)
	CRN         string       `json:"crn"`
	Description string       `json:"description,omitempty"`
	ID          string       `json:"id"`
	KMSKey      *ResourceRef `json:"kms_key,omitempty"`

	// Name 1..512 chars, [A-Za-z0-9_./#-]. Leading "/" not allowed. Resource
	// names must not start with the literal crn: prefix or be UUIDs
	// (canonical, compact, braced, or urn:uuid: forms, in either case).
	Name string `json:"name"`

	// RetentionDays retention in days, or null for no expiry.
	//
	// One of: "1", "3", "5", "7", "14", "30", "60", "90", "120", "150", "180", "365", "400", "545", "731", "1096", "1827", "2192", "2557", "2922", "3288", "3653", "<nil>".
	RetentionDays *int              `json:"retention_days"`
	Tags          map[string]string `json:"tags"`
	UpdatedAt     time.Time         `json:"updated_at"`
}

type LogRecord struct {
	Attributes map[string]string `json:"attributes,omitempty"`
	Body       string            `json:"body"`
	ID         string            `json:"id"`
	LogGroup   *ResourceRef      `json:"log_group"`

	// LogStream within-group stream name (typically the producer host)
	LogStream string            `json:"log_stream"`
	Resource  map[string]string `json:"resource,omitempty"`

	// One of: "TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL".
	Severity string `json:"severity"`

	// SeverityNumber OTel numeric severity band (1..24)
	SeverityNumber int    `json:"severity_number"`
	SpanID         string `json:"span_id,omitempty"`

	// Timestamp producer-stamped event time
	Timestamp time.Time `json:"timestamp"`
	TraceID   string    `json:"trace_id,omitempty"`
}

type MetricPoint struct {
	// Timestamp unix timestamp in seconds, including fractional seconds.
	Timestamp float64 `json:"timestamp"`

	// Value float value as text, including NaN and infinities.
	Value string `json:"value"`
}

type MetricQueryResult struct {
	Metric string               `json:"metric"`
	Series []*MetricQuerySeries `json:"series"`
}

type MetricQuerySeries struct {
	Labels map[string]string `json:"labels"`
	Points []*MetricPoint    `json:"points"`
}

type QueryMetricsInstantPostRequest struct {
	// One of: "avg", "sum", "min", "max", "count", "last", "rate", "increase".
	Agg   *string  `json:"agg,omitempty"`
	By    []string `json:"by,omitempty"`
	Match []string `json:"match,omitempty"`

	// Required.
	Metric string  `json:"metric"`
	Step   *string `json:"step,omitempty"`
	Time   *string `json:"time,omitempty"`
}

type QueryMetricsRangePostRequest struct {
	// One of: "avg", "sum", "min", "max", "count", "last", "rate", "increase".
	Agg *string  `json:"agg,omitempty"`
	By  []string `json:"by,omitempty"`

	// Required.
	End   string   `json:"end"`
	Match []string `json:"match,omitempty"`

	// Required.
	Metric string `json:"metric"`

	// Required.
	Start string  `json:"start"`
	Step  *string `json:"step,omitempty"`
}

type ResourceRef struct {
	CRN string `json:"crn"`
	ID  string `json:"id"`

	// Name display label at response time — the resource's name where it has
	// one, otherwise its human label (an invoice number, an account
	// display name). Absent when the target is gone. Not a lookup key.
	Name string `json:"name,omitempty"`
}

type Span struct {
	Attributes map[string]string `json:"attributes,omitempty"`
	DurationMs float64           `json:"duration_ms"`
	EndTime    time.Time         `json:"end_time"`
	Events     []*SpanEvent      `json:"events,omitempty"`

	// One of: "INTERNAL", "SERVER", "CLIENT", "PRODUCER", "CONSUMER".
	Kind         string            `json:"kind"`
	Links        []*SpanLink       `json:"links,omitempty"`
	Name         string            `json:"name"`
	ParentSpanID string            `json:"parent_span_id,omitempty"`
	Resource     map[string]string `json:"resource,omitempty"`
	ServiceName  string            `json:"service_name"`
	SpanID       string            `json:"span_id"`
	StartTime    time.Time         `json:"start_time"`

	// One of: "UNSET", "OK", "ERROR".
	StatusCode    string `json:"status_code"`
	StatusMessage string `json:"status_message,omitempty"`
	TraceID       string `json:"trace_id"`
}

type SpanEvent struct {
	Attributes map[string]string `json:"attributes,omitempty"`

	// Required.
	Name      string     `json:"name"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

type SpanIngestRecord struct {
	Attributes map[string]string `json:"attributes,omitempty"`

	// Required.
	EndTime time.Time    `json:"end_time"`
	Events  []*SpanEvent `json:"events,omitempty"`

	// Kind defaults to INTERNAL
	//
	// One of: "INTERNAL", "SERVER", "CLIENT", "PRODUCER", "CONSUMER".
	Kind  *string     `json:"kind,omitempty"`
	Links []*SpanLink `json:"links,omitempty"`

	// Name operation name (e.g. "GET /api/users")
	//
	// Required.
	Name string `json:"name"`

	// ParentSpanID 16 lower-hex chars or empty for root
	ParentSpanID *string           `json:"parent_span_id,omitempty"`
	Resource     map[string]string `json:"resource,omitempty"`

	// ServiceName defaults to resource[service.name]
	ServiceName *string `json:"service_name,omitempty"`

	// SpanID 16 lower-hex chars (OTel SpanId)
	//
	// Required.
	SpanID string `json:"span_id"`

	// Required.
	StartTime time.Time `json:"start_time"`

	// One of: "UNSET", "OK", "ERROR".
	StatusCode    *string `json:"status_code,omitempty"`
	StatusMessage *string `json:"status_message,omitempty"`

	// TraceID 32 lower-hex chars (OTel TraceId)
	//
	// Required.
	TraceID string `json:"trace_id"`
}

type SpanLink struct {
	// Required.
	SpanID string `json:"span_id"`

	// Required.
	TraceID string `json:"trace_id"`
}

type TraceSettings struct {
	CreatedAt time.Time    `json:"created_at,omitempty"`
	KMSKey    *ResourceRef `json:"kms_key,omitempty"`

	// RetentionDays effective retention in days. Defaults to 30; traces always expire.
	//
	// One of: "1", "3", "5", "7", "14", "30", "60", "90", "120", "150", "180", "365", "400", "545", "731", "1096", "1827", "2192", "2557", "2922", "3288", "3653".
	RetentionDays int       `json:"retention_days"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}

type TraceSummary struct {
	DurationMs   float64   `json:"duration_ms"`
	ErrorCount   int       `json:"error_count"`
	RootName     string    `json:"root_name"`
	RootService  string    `json:"root_service"`
	RootSpanID   string    `json:"root_span_id"`
	ServiceCount int       `json:"service_count"`
	SpanCount    int       `json:"span_count"`
	StartTime    time.Time `json:"start_time"`
	TraceID      string    `json:"trace_id"`
}

// UpdateAlertRequest only supplied fields change. The expression is replaced as a whole.
// Null is allowed only for silenced_until. Name is immutable. Rule edits
// and enable changes reset evaluation; edits discard events queued for
// the previous configuration.
type UpdateAlertRequest struct {
	Enabled    *bool            `json:"enabled,omitempty"`
	Expression *AlertExpression `json:"expression,omitempty"`

	// ForSeconds continuous breach duration before firing; 0 fires immediately.
	// Missing data, query failures and scheduling gaps reset a pending
	// duration.
	ForSeconds *int `json:"for_seconds,omitempty"`

	// One of: "gt", "gte", "lt", "lte".
	Operator *string `json:"operator,omitempty"`

	// SilencedUntil mute notifications until this instant (at most 30 days ahead).
	// Evaluation continues; null removes silence. An active alert notifies
	// on the next evaluation after silence expires.
	SilencedUntil *time.Time `json:"silenced_until,omitempty"`
	Threshold     *float64   `json:"threshold,omitempty"`

	// Topic required same-account, same-region notification topic reference:
	// UUID, canonical CRN or exact name, classified by syntax. Multiple
	// alerts may share a topic. Binding requires notifications:Publish and
	// respects explicit topic-policy denies.
	Topic *string `json:"topic,omitempty"`

	// WindowSeconds whole-minute window. Evaluated every minute against the last
	// complete UTC-aligned window.
	WindowSeconds *int `json:"window_seconds,omitempty"`
}

type UpdateLogGroupRequest struct {
	Description *string `json:"description,omitempty"`

	// KMSKey KMS key UUID, CRN or exact name in this account and region. Omit to
	// preserve the binding; null clears encryption for future records.
	KMSKey *string `json:"kms_key,omitempty"`

	// RetentionDays omit to preserve retention. Null selects no expiry.
	//
	// One of: "1", "3", "5", "7", "14", "30", "60", "90", "120", "150", "180", "365", "400", "545", "731", "1096", "1827", "2192", "2557", "2922", "3288", "3653", "<nil>".
	RetentionDays *int              `json:"retention_days,omitempty"`
	Tags          map[string]string `json:"tags,omitempty"`
}

type UpdateTraceSettingsRequest struct {
	// KMSKey KMS key UUID, CRN or exact name in this account and region. Omit to
	// preserve the binding; null clears encryption for future records.
	KMSKey *string `json:"kms_key,omitempty"`

	// RetentionDays omit to preserve retention. Null restores the default 30 days;
	// traces always expire.
	//
	// One of: "1", "3", "5", "7", "14", "30", "60", "90", "120", "150", "180", "365", "400", "545", "731", "1096", "1827", "2192", "2557", "2922", "3288", "3653", "<nil>".
	RetentionDays *int `json:"retention_days,omitempty"`
}
