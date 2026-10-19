// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package email

import (
	"time"
)

// Attachment an immutable attachment reference. id is stable within the message.
// version_id is the Storage object version; sha256 pins decoded bytes
// even in an unversioned bucket. Bucket lifecycle controls object
// retention independently of the 30-day Email content lifetime. Legacy
// attachments can have a null bucket and remain downloadable by id until
// message expiry.
type Attachment struct {
	Bucket      *ResourceRef `json:"bucket"`
	ContentType string       `json:"content_type"`
	ID          string       `json:"id"`
	Key         string       `json:"key"`
	Name        string       `json:"name"`

	// Sha256 SHA-256 of the decoded attachment bytes.
	Sha256 string `json:"sha256"`
	Size   int    `json:"size"`

	// VersionID Opaque Storage object version token, when versioning is enabled.
	VersionID string `json:"version_id,omitempty"`
}

type AttachmentInput struct {
	// Bucket same-account, same-region bucket name, UUID or CRN.
	//
	// Required.
	Bucket string `json:"bucket"`

	// ContentType MIME type; defaults to the object content type.
	ContentType *string `json:"content_type,omitempty"`

	// Required.
	Key string `json:"key"`

	// Name attachment filename.
	//
	// Required.
	Name string `json:"name"`

	// VersionID optional opaque Storage object version token.
	VersionID *string `json:"version_id,omitempty"`
}

type BatchSendRequest struct {
	// Required.
	Messages []*SendMessageRequest `json:"messages"`
}

type BatchSendResult struct {
	Error   *Error       `json:"error,omitempty"`
	Index   int          `json:"index"`
	Message *ResourceRef `json:"message,omitempty"`

	// One of: "queued", "rejected".
	Status string `json:"status"`
}

type ConfigurationSet struct {
	CreatedAt time.Time    `json:"created_at,omitempty"`
	CRN       string       `json:"crn"`
	ID        string       `json:"id"`
	IPPool    *ResourceRef `json:"ip_pool,omitempty"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	Name                 string    `json:"name"`
	SendingEnabled       bool      `json:"sending_enabled"`
	SuppressionOverrides []string  `json:"suppression_overrides"`
	Tags                 Tags      `json:"tags,omitempty"`
	TrackingClick        bool      `json:"tracking_click,omitempty"`
	TrackingOpen         bool      `json:"tracking_open,omitempty"`
	UpdatedAt            time.Time `json:"updated_at,omitempty"`
}

type CreateConfigurationSetRequest struct {
	// IPPool UUID, CRN or account-unique name. Omit to preserve; null detaches
	// the pool.
	IPPool *string `json:"ip_pool,omitempty"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	//
	// Required.
	Name                 string   `json:"name"`
	SendingEnabled       *bool    `json:"sending_enabled,omitempty"`
	SuppressionOverrides []string `json:"suppression_overrides,omitempty"`
	Tags                 Tags     `json:"tags,omitempty"`
	TrackingClick        *bool    `json:"tracking_click,omitempty"`
	TrackingOpen         *bool    `json:"tracking_open,omitempty"`
}

type CreateIPPoolRequest struct {
	Description *string `json:"description,omitempty"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	//
	// Required.
	Name string `json:"name"`
}

type CreateIdentityRequest struct {
	// Name sender domain (FQDN, lowercase). Resource names must not start with
	// the literal crn: prefix or be UUIDs (canonical, compact, braced, or
	// urn:uuid: forms, in either case).
	//
	// Required.
	Name string `json:"name"`
	Tags Tags   `json:"tags,omitempty"`
}

type CreateRecipientRuleRequest struct {
	Enabled *bool `json:"enabled,omitempty"`

	// LocalPart required for exact and prefix; empty or omitted for all.
	LocalPart *string `json:"local_part,omitempty"`

	// MatchType exact local part, local-part prefix, or catch-all for this identity.
	// Exact matches take precedence, then the longest prefix, then
	// catch-all.
	//
	// One of: "exact", "prefix", "all".
	//
	// Required.
	MatchType string `json:"match_type"`

	// Topic notifications topic name, UUID or CRN. Requires
	// notifications:Publish on the same-account topic.
	//
	// Required.
	Topic string `json:"topic"`
}

type CreateSuppressionRequest struct {
	// Required.
	Address     string     `json:"address"`
	Description *string    `json:"description,omitempty"`
	ExpiresAt   *time.Time `json:"expires_at,omitempty"`
}

type CreateTemplateRequest struct {
	// Required.
	BodyTemplate string `json:"body_template"`

	// One of: "text", "html".
	BodyType *string `json:"body_type,omitempty"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	//
	// Required.
	Name            string  `json:"name"`
	SubjectTemplate *string `json:"subject_template,omitempty"`
}

type DNSRecord struct {
	// AlternativeGroup records sharing this value are alternatives. Publish exactly one
	// record from a required alternative group; TXT and CNAME cannot
	// coexist at the same name.
	AlternativeGroup string `json:"alternative_group,omitempty"`

	// One of: "sending", "receiving".
	Group string `json:"group"`

	// Name FQDN to publish the record at.
	Name string `json:"name"`

	// One of: "ownership", "dkim", "dkim_cname", "dkim_pending", "dkim_pending_cname", "dkim_rotated_out", "spf", "dmarc", "return_path", "mail_from_mx", "mail_from_spf", "bimi", "receiving_mx".
	Purpose string `json:"purpose"`

	// Required for this sending/receiving function. For an
	// alternative_group, one alternative satisfies the requirement.
	Required bool `json:"required"`

	// One of: "TXT", "CNAME", "MX".
	Type string `json:"type"`

	// Value complete copyable record data: unquoted TXT content, CNAME target,
	// or MX priority followed by host.
	Value string `json:"value"`
}

type DedicatedIP struct {
	Address   string       `json:"address"`
	CreatedAt time.Time    `json:"created_at"`
	CRN       string       `json:"crn"`
	ID        string       `json:"id"`
	IPPool    *ResourceRef `json:"ip_pool"`
	UpdatedAt time.Time    `json:"updated_at"`

	// WarmupPercentage progress computed by the warmup worker; cannot be edited.
	WarmupPercentage int        `json:"warmup_percentage"`
	WarmupStartedAt  *time.Time `json:"warmup_started_at,omitempty"`

	// One of: "new", "warming", "warm".
	WarmupStatus string `json:"warmup_status"`
}

type DmarcReportRecord struct {
	Disposition  string    `json:"disposition"`
	DkimAligned  bool      `json:"dkim_aligned"`
	DkimDomain   string    `json:"dkim_domain"`
	DkimResult   string    `json:"dkim_result"`
	HeaderFrom   string    `json:"header_from"`
	MessageCount int       `json:"message_count"`
	PeriodEnd    time.Time `json:"period_end"`
	PeriodStart  time.Time `json:"period_start"`
	ReceivedAt   time.Time `json:"received_at"`
	ReportID     string    `json:"report_id"`
	Reporter     string    `json:"reporter"`
	SourceIP     string    `json:"source_ip"`
	SpfAligned   bool      `json:"spf_aligned"`
	SpfDomain    string    `json:"spf_domain"`
	SpfResult    string    `json:"spf_result"`
}

type Error struct {
	// Code error code identifying the type of error
	Code string `json:"code"`

	// Message human-readable error message
	Message string `json:"message"`

	// Params optional non-sensitive values for interpolation of localized errors,
	// keyed by error code. Never present for server errors.
	Params map[string]any `json:"params,omitempty"`

	// RequestID Request ID for debugging
	RequestID string `json:"request_id"`
}

type IPPool struct {
	CreatedAt   time.Time `json:"created_at"`
	CRN         string    `json:"crn"`
	Description string    `json:"description,omitempty"`
	ID          string    `json:"id"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Identity struct {
	BimiLogoURL string `json:"bimi_logo_url,omitempty"`

	// One of: "pending", "verified", "failed".
	BimiStatus          string       `json:"bimi_status,omitempty"`
	BimiVerifiedAt      *time.Time   `json:"bimi_verified_at,omitempty"`
	BimiVmcURL          string       `json:"bimi_vmc_url,omitempty"`
	CreatedAt           time.Time    `json:"created_at"`
	CRN                 string       `json:"crn"`
	DkimOldExpiresAt    *time.Time   `json:"dkim_old_expires_at,omitempty"`
	DkimPendingSelector string       `json:"dkim_pending_selector,omitempty"`
	DkimSelector        string       `json:"dkim_selector"`
	DkimSelectorOld     string       `json:"dkim_selector_old,omitempty"`
	DNSRecords          []*DNSRecord `json:"dns_records,omitempty"`
	ID                  string       `json:"id"`
	MailFromDomain      string       `json:"mail_from_domain,omitempty"`

	// One of: "pending", "verified", "failed".
	MailFromStatus     string     `json:"mail_from_status,omitempty"`
	MailFromVerifiedAt *time.Time `json:"mail_from_verified_at,omitempty"`

	// Name sender domain. Resource names must not start with the literal crn:
	// prefix or be UUIDs (canonical, compact, braced, or urn:uuid: forms,
	// in either case).
	Name string `json:"name"`

	// One of: "pending", "verified", "failed", "suspended".
	Status          string     `json:"status"`
	SuspendedAt     *time.Time `json:"suspended_at,omitempty"`
	SuspendedReason string     `json:"suspended_reason,omitempty"`
	Tags            Tags       `json:"tags,omitempty"`

	// One of: "domain".
	Type       string     `json:"type"`
	UpdatedAt  time.Time  `json:"updated_at"`
	VerifiedAt *time.Time `json:"verified_at,omitempty"`
}

type Message struct {
	Attachments  []*Attachment `json:"attachments"`
	CRN          string        `json:"crn"`
	ErrorMessage string        `json:"error_message,omitempty"`

	// ExpiresAt message content expires 30 days after acceptance.
	ExpiresAt    time.Time    `json:"expires_at"`
	FailedAt     *time.Time   `json:"failed_at,omitempty"`
	From         string       `json:"from"`
	ID           string       `json:"id"`
	Identity     *ResourceRef `json:"identity"`
	OutboundIP   string       `json:"outbound_ip,omitempty"`
	QueuedAt     time.Time    `json:"queued_at"`
	SentAt       *time.Time   `json:"sent_at,omitempty"`
	SMTPResponse string       `json:"smtp_response,omitempty"`

	// One of: "queued", "sending", "sent", "failed", "bounced", "complained".
	Status  string   `json:"status"`
	Subject string   `json:"subject,omitempty"`
	To      []string `json:"to"`
}

type MessageContent struct {
	Attachments []*Attachment `json:"attachments"`
	From        string        `json:"from"`

	// HTML full decoded HTML. Untrusted content: sanitize before rendering; do
	// not load remote resources automatically.
	HTML string `json:"html"`

	// RfcMessageID RFC Message-ID header, distinct from the Email resource UUID.
	RfcMessageID string `json:"rfc_message_id,omitempty"`
	Subject      string `json:"subject"`

	// Text full decoded plain-text body.
	Text string   `json:"text"`
	To   []string `json:"to"`
}

type MessageEngagement struct {
	ClickCount   int64      `json:"click_count"`
	FirstClickAt *time.Time `json:"first_click_at,omitempty"`
	FirstOpenAt  *time.Time `json:"first_open_at,omitempty"`
	GeneratedAt  time.Time  `json:"generated_at"`
	LastClickAt  *time.Time `json:"last_click_at,omitempty"`
	LastOpenAt   *time.Time `json:"last_open_at,omitempty"`
	OpenCount    int64      `json:"open_count"`
	UniqueClicks int64      `json:"unique_clicks"`
	UniqueOpens  int64      `json:"unique_opens"`
}

type ReceivedMessage struct {
	Attachments      []*Attachment `json:"attachments"`
	Authenticated    bool          `json:"authenticated"`
	ContentAvailable bool          `json:"content_available"`
	CRN              string        `json:"crn"`
	ExpiresAt        time.Time     `json:"expires_at"`
	From             string        `json:"from"`
	ID               string        `json:"id"`

	// Identity receiving identity; null only when the original identity no longer
	// exists.
	Identity *ResourceRef `json:"identity"`

	// NotificationError safe publication failure detail. Retried automatically while content
	// is retained.
	NotificationError *string `json:"notification_error"`

	// One of: "pending", "retrying", "published".
	NotificationStatus string     `json:"notification_status"`
	PublishedAt        *time.Time `json:"published_at"`
	ReceivedAt         time.Time  `json:"received_at"`
	Recipient          string     `json:"recipient"`

	// One of: "", "spam", "sender_unverified", "malware".
	ReviewReason string `json:"review_reason"`

	// RfcMessageID RFC Message-ID header, distinct from the Email resource UUID.
	RfcMessageID string `json:"rfc_message_id,omitempty"`
	Subject      string `json:"subject"`
}

type ReceivingConfiguration struct {
	AttachmentsBucket *ResourceRef `json:"attachments_bucket"`
	Enabled           bool         `json:"enabled"`
}

type RecipientRule struct {
	CRN       string `json:"crn"`
	Enabled   bool   `json:"enabled"`
	ID        string `json:"id"`
	LocalPart string `json:"local_part"`

	// MatchType exact local part, local-part prefix, or catch-all for this identity.
	// Exact matches take precedence, then the longest prefix, then
	// catch-all.
	//
	// One of: "exact", "prefix", "all".
	MatchType string       `json:"match_type"`
	Topic     *ResourceRef `json:"topic"`
}

type Rejection struct {
	CRN      string       `json:"crn"`
	Detail   string       `json:"detail,omitempty"`
	From     string       `json:"from,omitempty"`
	ID       string       `json:"id"`
	Identity *ResourceRef `json:"identity,omitempty"`

	// One of: "ALL_RECIPIENTS_SUPPRESSED", "IDENTITY_SUSPENDED", "IDENTITY_NOT_VERIFIED", "FROM_DOMAIN_MISMATCH", "RATE_LIMIT".
	Reason     string    `json:"reason"`
	RejectedAt time.Time `json:"rejected_at"`
	To         []string  `json:"to"`
}

type RequestDedicatedIPRequest struct {
	// IPPool UUID, CRN or account-unique pool name; empty leaves unassigned.
	IPPool *string `json:"ip_pool,omitempty"`
}

type ResourceRef struct {
	CRN string `json:"crn"`
	ID  string `json:"id"`

	// Name display label at response time — the resource's name where it has
	// one, otherwise its human label (an invoice number, an account
	// display name). Absent when the target is gone. Not a lookup key.
	Name string `json:"name,omitempty"`
}

// SendContent exactly one content mode: text and/or html; raw; or template with
// optional template_data. Text is never guessed to be HTML.
type SendContent struct {
	HTML *string `json:"html,omitempty"`

	// Raw Complete RFC 5322 message. Its From header must belong to the
	// authorized identity domain.
	Raw *string `json:"raw,omitempty"`

	// Template UUID, CRN or name; requires email:SendTemplatedMessage on
	// that template.
	Template     *string           `json:"template,omitempty"`
	TemplateData map[string]string `json:"template_data,omitempty"`
	Text         *string           `json:"text,omitempty"`
}

type SendMessageRequest struct {
	// Attachments bucket objects to attach. GetObject permission is required; version
	// and checksum are pinned at acceptance.
	Attachments []*AttachmentInput `json:"attachments,omitempty"`

	// AttachmentsBucket Bucket UUID, CRN or name for attachments embedded in raw MIME.
	// Objects are written under email/{identity_id}/sent/{message_id}/.
	// Separate attachments entries reference existing bucket objects.
	AttachmentsBucket *string `json:"attachments_bucket,omitempty"`

	// ConfigurationSet UUID, CRN or account-unique name.
	ConfigurationSet *string `json:"configuration_set,omitempty"`

	// Required.
	Content *SendContent `json:"content"`

	// Required.
	From string `json:"from"`

	// Identity UUID, CRN or account-unique name.
	//
	// Required.
	Identity string `json:"identity"`

	// Subject for text/HTML content. For raw or template content, the
	// subject is supplied by that content.
	Subject *string `json:"subject,omitempty"`

	// Required.
	To []string `json:"to"`
}

// SetBimiRequest Omitted URLs are preserved. Null logo_url clears BIMI, including its
// certificate; null vmc_url clears only the certificate. A certificate
// requires a logo.
type SetBimiRequest struct {
	// LogoURL HTTPS SVG Tiny PS URL; null clears BIMI.
	LogoURL *string `json:"logo_url,omitempty"`

	// VmcURL Optional HTTPS VMC URL; omission preserves it and null removes the
	// certificate.
	VmcURL *string `json:"vmc_url,omitempty"`
}

type SetMailFromDomainRequest struct {
	// Domain mail-from domain; null clears the binding.
	Domain *string `json:"domain,omitempty"`
}

type Suppression struct {
	Address     string     `json:"address"`
	CreatedAt   time.Time  `json:"created_at"`
	CRN         string     `json:"crn"`
	Description string     `json:"description,omitempty"`
	ExpiresAt   *time.Time `json:"expires_at,omitempty"`
	ID          string     `json:"id"`

	// One of: "auto_bounce", "auto_complaint", "manual".
	Reason        string       `json:"reason"`
	SourceMessage *ResourceRef `json:"source_message,omitempty"`
	UpdatedAt     time.Time    `json:"updated_at"`
}

type SuppressionOverrideRequest struct {
	// Required.
	Address string `json:"address"`
}

type Tags = map[string]string

type Template struct {
	BodyTemplate string `json:"body_template"`

	// One of: "text", "html".
	BodyType  string    `json:"body_type"`
	CreatedAt time.Time `json:"created_at"`
	CRN       string    `json:"crn"`
	ID        string    `json:"id"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	Name            string    `json:"name"`
	SubjectTemplate string    `json:"subject_template"`
	UpdatedAt       time.Time `json:"updated_at"`
	Variables       []string  `json:"variables"`
}

type UpdateConfigurationSetRequest struct {
	// IPPool UUID, CRN or account-unique name. Omit to preserve; null detaches
	// the pool.
	IPPool               *string  `json:"ip_pool,omitempty"`
	SendingEnabled       *bool    `json:"sending_enabled,omitempty"`
	SuppressionOverrides []string `json:"suppression_overrides,omitempty"`
	Tags                 Tags     `json:"tags,omitempty"`
	TrackingClick        *bool    `json:"tracking_click,omitempty"`
	TrackingOpen         *bool    `json:"tracking_open,omitempty"`
}

type UpdateDedicatedIPRequest struct {
	// IPPool UUID, CRN or account-unique pool name. Null detaches; omission
	// leaves unchanged.
	IPPool *string `json:"ip_pool,omitempty"`

	// One of: "new", "warming", "warm".
	WarmupStatus *string `json:"warmup_status,omitempty"`
}

type UpdateIPPoolRequest struct {
	Description *string `json:"description,omitempty"`
}

type UpdateIdentityRequest struct {
	Tags Tags `json:"tags,omitempty"`
}

type UpdateReceivingConfigurationRequest struct {
	// AttachmentsBucket same-account, same-region bucket name, UUID or CRN. Omission
	// preserves the current bucket; null clears it while receiving is
	// disabled. Saving requires storage PutObject and GetObject permission
	// under email/{identity_id}/ in this bucket. Changing the bucket
	// affects new messages only.
	AttachmentsBucket *string `json:"attachments_bucket,omitempty"`
	Enabled           *bool   `json:"enabled,omitempty"`
}

type UpdateRecipientRuleRequest struct {
	Enabled *bool `json:"enabled,omitempty"`

	// LocalPart required for exact and prefix; empty or omitted for all.
	LocalPart *string `json:"local_part,omitempty"`

	// MatchType exact local part, local-part prefix, or catch-all for this identity.
	// Exact matches take precedence, then the longest prefix, then
	// catch-all.
	//
	// One of: "exact", "prefix", "all".
	MatchType *string `json:"match_type,omitempty"`

	// Topic notifications topic name, UUID or CRN. Requires
	// notifications:Publish on the same-account topic.
	Topic *string `json:"topic,omitempty"`
}

type UpdateSuppressionRequest struct {
	Description *string `json:"description,omitempty"`

	// ExpiresAt null makes the suppression permanent; omission preserves its expiry.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// UpdateTemplateRequest omitted fields are preserved. An empty subject clears it.
// body_template must be nonempty and body_type must be text or html.
// Null values are rejected.
type UpdateTemplateRequest struct {
	BodyTemplate *string `json:"body_template,omitempty"`

	// One of: "text", "html".
	BodyType        *string `json:"body_type,omitempty"`
	SubjectTemplate *string `json:"subject_template,omitempty"`
}
