// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package telemetry

import (
	"time"
)

type CreateLogGroupRequest struct {
	Description *string `json:"description,omitempty"`
	KMSKeyCRN   *string `json:"kms_key_crn,omitempty"`

	// Required.
	Name string `json:"name"`

	// RetentionDays 1..3650, or omit for never expire
	RetentionDays *int              `json:"retention_days,omitempty"`
	Tags          map[string]string `json:"tags,omitempty"`
}

type IngestRecord struct {
	Attributes map[string]string `json:"attributes,omitempty"`

	// Required.
	Body string `json:"body"`

	// LogGroup name of an existing log group. Rejected if not registered.
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
	Accepted int      `json:"accepted,omitempty"`
	Errors   []string `json:"errors,omitempty"`
	Rejected int      `json:"rejected,omitempty"`
}

type IngestSpansRequest struct {
	// Required.
	Spans []*SpanIngestRecord `json:"spans"`
}

type ListMetricNamesPostResult struct {
	Data   []string `json:"data,omitempty"`
	Status string   `json:"status,omitempty"`
}

type ListMetricNamesResult struct {
	Data   []string `json:"data,omitempty"`
	Status string   `json:"status,omitempty"`
}

type LogGroup struct {
	// AccountID owning account — the tenant fence for this group
	AccountID string    `json:"account_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`

	// CRN of the log group (used as the IAM resource ARN)
	CRN         string `json:"crn,omitempty"`
	Description string `json:"description,omitempty"`
	ID          string `json:"id,omitempty"`

	// KMSKeyCRN KMS key associated with the group (empty = plaintext)
	KMSKeyCRN string `json:"kms_key_crn,omitempty"`

	// Name 1..512 chars, [A-Za-z0-9_./#-]. Leading "/" not allowed.
	Name string `json:"name,omitempty"`

	// RetentionDays 1..3650 days, or null for never expire
	RetentionDays int               `json:"retention_days,omitempty"`
	Tags          map[string]string `json:"tags,omitempty"`
	UpdatedAt     time.Time         `json:"updated_at,omitempty"`
}

type LogRecord struct {
	// AccountID owning account, empty for org-scoped events
	AccountID  string            `json:"account_id,omitempty"`
	Attributes map[string]string `json:"attributes,omitempty"`
	Body       string            `json:"body,omitempty"`
	ID         string            `json:"id,omitempty"`

	// LogGroup name of the parent log group
	LogGroup   string `json:"log_group,omitempty"`
	LogGroupID string `json:"log_group_id,omitempty"`

	// LogStream within-group stream name (typically the producer host)
	LogStream      string            `json:"log_stream,omitempty"`
	OrganizationID string            `json:"organization_id,omitempty"`
	Resource       map[string]string `json:"resource,omitempty"`

	// One of: "TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL".
	Severity string `json:"severity,omitempty"`

	// SeverityNumber OTel numeric severity band (1..24)
	SeverityNumber int    `json:"severity_number,omitempty"`
	SpanID         string `json:"span_id,omitempty"`

	// Timestamp producer-stamped event time
	Timestamp time.Time `json:"timestamp,omitempty"`
	TraceID   string    `json:"trace_id,omitempty"`
}

type Span struct {
	Attributes map[string]string `json:"attributes,omitempty"`
	DurationMs float64           `json:"duration_ms,omitempty"`
	EndTime    time.Time         `json:"end_time,omitempty"`
	Events     []*SpanEvent      `json:"events,omitempty"`

	// One of: "INTERNAL", "SERVER", "CLIENT", "PRODUCER", "CONSUMER".
	Kind         string            `json:"kind,omitempty"`
	Links        []*SpanLink       `json:"links,omitempty"`
	Name         string            `json:"name,omitempty"`
	ParentSpanID string            `json:"parent_span_id,omitempty"`
	Resource     map[string]string `json:"resource,omitempty"`
	ServiceName  string            `json:"service_name,omitempty"`
	SpanID       string            `json:"span_id,omitempty"`
	StartTime    time.Time         `json:"start_time,omitempty"`

	// One of: "UNSET", "OK", "ERROR".
	StatusCode    string `json:"status_code,omitempty"`
	StatusMessage string `json:"status_message,omitempty"`
	TraceID       string `json:"trace_id,omitempty"`
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

type TraceResponse struct {
	Spans   []*Span `json:"spans,omitempty"`
	TraceID string  `json:"trace_id,omitempty"`
}

type TraceSettings struct {
	AccountID string    `json:"account_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`

	// KMSKeyCRN KMS key CRN that span bodies are envelope-encrypted under (empty =
	// plaintext)
	KMSKeyCRN string `json:"kms_key_crn,omitempty"`

	// RetentionDays 1..3650, or null for never-expire
	RetentionDays int       `json:"retention_days,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}

type TraceSummary struct {
	DurationMs   float64   `json:"duration_ms,omitempty"`
	ErrorCount   int       `json:"error_count,omitempty"`
	RootName     string    `json:"root_name,omitempty"`
	RootService  string    `json:"root_service,omitempty"`
	RootSpanID   string    `json:"root_span_id,omitempty"`
	ServiceCount int       `json:"service_count,omitempty"`
	SpanCount    int       `json:"span_count,omitempty"`
	StartTime    time.Time `json:"start_time,omitempty"`
	TraceID      string    `json:"trace_id,omitempty"`
}

type UpdateLogGroupRequest struct {
	// ClearRetention when true, sets retention to never-expire (ignores retention_days)
	ClearRetention *bool   `json:"clear_retention,omitempty"`
	Description    *string `json:"description,omitempty"`

	// KMSKeyCRN pass an empty string to disassociate
	KMSKeyCRN *string `json:"kms_key_crn,omitempty"`

	// RetentionDays 1..3650; pass clear_retention to switch to never expire
	RetentionDays *int              `json:"retention_days,omitempty"`
	Tags          map[string]string `json:"tags,omitempty"`
}

type UpdateTraceSettingsRequest struct {
	// ClearRetention when true, sets retention to never-expire (ignores retention_days)
	ClearRetention *bool `json:"clear_retention,omitempty"`

	// KMSKeyCRN pass an empty string to disassociate
	KMSKeyCRN *string `json:"kms_key_crn,omitempty"`

	// RetentionDays 1..3650; pass clear_retention=true to switch to never-expire
	RetentionDays *int `json:"retention_days,omitempty"`
}
