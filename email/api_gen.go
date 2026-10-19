// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package email

import (
	"context"
	"io"
	"iter"
	"net/url"
	"strconv"
	"time"

	basaltic "github.com/basaltic-sh/sdk-go"
)

// ListConfigurationSetsParams are the optional filters and pagination controls for
// [Client.ListConfigurationSets]. A nil *ListConfigurationSetsParams sends none of them.
type ListConfigurationSetsParams struct {
	// CRN Exact CRN filter. Malformed CRNs return 400; foreign or mismatched
	// CRNs return an empty page.
	CRN    string
	Limit  int
	Marker string

	// Name exact immutable account-scoped name, combined with all other
	// filters.
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListConfigurationSetsParams) query() url.Values {
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
func (p *ListConfigurationSetsParams) withMarker(marker string) *ListConfigurationSetsParams {
	var out ListConfigurationSetsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListDedicatedIPsParams are the optional filters and pagination controls for
// [Client.ListDedicatedIPs]. A nil *ListDedicatedIPsParams sends none of them.
type ListDedicatedIPsParams struct {
	// CRN exact dedicated-IP CRN, filtered before pagination. Foreign or
	// mismatched CRNs return an empty page.
	CRN    string
	IPPool string
	Limit  int
	Marker string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListDedicatedIPsParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.CRN != "" {
		q.Set("crn", p.CRN)
	}
	if p.IPPool != "" {
		q.Set("ip_pool", p.IPPool)
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
func (p *ListDedicatedIPsParams) withMarker(marker string) *ListDedicatedIPsParams {
	var out ListDedicatedIPsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListDmarcReportsParams are the optional filters and pagination controls for
// [Client.ListDmarcReports]. A nil *ListDmarcReportsParams sends none of them.
type ListDmarcReportsParams struct {
	From   time.Time
	Limit  int
	Marker string
	To     time.Time
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListDmarcReportsParams) query() url.Values {
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
	if !p.To.IsZero() {
		q.Set("to", p.To.UTC().Format(time.RFC3339))
	}
	return q
}

// withMarker copies p with the pagination cursor replaced, leaving the
// caller's value untouched across pages.
func (p *ListDmarcReportsParams) withMarker(marker string) *ListDmarcReportsParams {
	var out ListDmarcReportsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListIPPoolsParams are the optional filters and pagination controls for
// [Client.ListIPPools]. A nil *ListIPPoolsParams sends none of them.
type ListIPPoolsParams struct {
	// CRN Exact CRN filter. Malformed CRNs return 400; foreign or mismatched
	// CRNs return an empty page.
	CRN    string
	Limit  int
	Marker string

	// Name exact immutable account-scoped name, combined with all other
	// filters.
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListIPPoolsParams) query() url.Values {
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
func (p *ListIPPoolsParams) withMarker(marker string) *ListIPPoolsParams {
	var out ListIPPoolsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListIdentitiesParams are the optional filters and pagination controls for
// [Client.ListIdentities]. A nil *ListIdentitiesParams sends none of them.
type ListIdentitiesParams struct {
	// CRN Exact CRN filter. Malformed CRNs return 400; foreign or mismatched
	// CRNs return an empty page.
	CRN    string
	Limit  int
	Marker string

	// Name exact immutable account-scoped name, combined with all other
	// filters.
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListIdentitiesParams) query() url.Values {
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
func (p *ListIdentitiesParams) withMarker(marker string) *ListIdentitiesParams {
	var out ListIdentitiesParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListReceivedMessagesParams are the optional filters and pagination controls for
// [Client.ListReceivedMessages]. A nil *ListReceivedMessagesParams sends none of them.
type ListReceivedMessagesParams struct {
	// Identity receiving identity UUID, CRN or name.
	Identity  string
	Limit     int
	Marker    string
	Recipient string

	// ReviewReason one of: "clear", "spam", "sender_unverified", "malware".
	ReviewReason string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListReceivedMessagesParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.Identity != "" {
		q.Set("identity", p.Identity)
	}
	if p.Limit != 0 {
		q.Set("limit", strconv.Itoa(int(p.Limit)))
	}
	if p.Marker != "" {
		q.Set("marker", p.Marker)
	}
	if p.Recipient != "" {
		q.Set("recipient", p.Recipient)
	}
	if p.ReviewReason != "" {
		q.Set("review_reason", p.ReviewReason)
	}
	return q
}

// withMarker copies p with the pagination cursor replaced, leaving the
// caller's value untouched across pages.
func (p *ListReceivedMessagesParams) withMarker(marker string) *ListReceivedMessagesParams {
	var out ListReceivedMessagesParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListRecipientRulesParams are the optional filters and pagination controls for
// [Client.ListRecipientRules]. A nil *ListRecipientRulesParams sends none of them.
type ListRecipientRulesParams struct {
	Limit  int
	Marker string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListRecipientRulesParams) query() url.Values {
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
func (p *ListRecipientRulesParams) withMarker(marker string) *ListRecipientRulesParams {
	var out ListRecipientRulesParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListRejectionsParams are the optional filters and pagination controls for
// [Client.ListRejections]. A nil *ListRejectionsParams sends none of them.
type ListRejectionsParams struct {
	// Before RFC 3339 timestamp — return only rejections older than this. Omit
	// to start from the newest. This is what `meta.marker` carries.
	Before time.Time

	// CRN exact rejection event CRN. Foreign scope or a missing event returns
	// an empty page.
	CRN string

	// Identity UUID, CRN or account-unique name; missing or foreign
	// references return empty results.
	Identity string
	Limit    int

	// Reason narrow to one rejection reason.
	//
	// One of: "ALL_RECIPIENTS_SUPPRESSED", "IDENTITY_SUSPENDED",
	// "IDENTITY_NOT_VERIFIED", "FROM_DOMAIN_MISMATCH", "RATE_LIMIT".
	Reason string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListRejectionsParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if !p.Before.IsZero() {
		q.Set("before", p.Before.UTC().Format(time.RFC3339))
	}
	if p.CRN != "" {
		q.Set("crn", p.CRN)
	}
	if p.Identity != "" {
		q.Set("identity", p.Identity)
	}
	if p.Limit != 0 {
		q.Set("limit", strconv.Itoa(int(p.Limit)))
	}
	if p.Reason != "" {
		q.Set("reason", p.Reason)
	}
	return q
}

// ListSentMessagesParams are the optional filters and pagination controls for
// [Client.ListSentMessages]. A nil *ListSentMessagesParams sends none of them.
type ListSentMessagesParams struct {
	// CRN Exact CRN filter. Malformed CRNs return 400; foreign or mismatched
	// CRNs return an empty page.
	CRN string

	// Identity UUID, CRN or account-unique name; missing or foreign
	// references return empty results.
	Identity  string
	Limit     int
	Marker    string
	Recipient string
	Status    string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListSentMessagesParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.CRN != "" {
		q.Set("crn", p.CRN)
	}
	if p.Identity != "" {
		q.Set("identity", p.Identity)
	}
	if p.Limit != 0 {
		q.Set("limit", strconv.Itoa(int(p.Limit)))
	}
	if p.Marker != "" {
		q.Set("marker", p.Marker)
	}
	if p.Recipient != "" {
		q.Set("recipient", p.Recipient)
	}
	if p.Status != "" {
		q.Set("status", p.Status)
	}
	return q
}

// withMarker copies p with the pagination cursor replaced, leaving the
// caller's value untouched across pages.
func (p *ListSentMessagesParams) withMarker(marker string) *ListSentMessagesParams {
	var out ListSentMessagesParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListSuppressionsParams are the optional filters and pagination controls for
// [Client.ListSuppressions]. A nil *ListSuppressionsParams sends none of them.
type ListSuppressionsParams struct {
	// CRN Exact CRN filter. Malformed CRNs return 400; foreign or mismatched
	// CRNs return an empty page.
	CRN    string
	Limit  int
	Marker string

	// Reason return only entries put on the list this way. An unknown value is
	// not rejected — it simply matches nothing, so the page comes back
	// empty rather than 400.
	//
	//
	// One of: "auto_bounce", "auto_complaint", "manual".
	Reason string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListSuppressionsParams) query() url.Values {
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
	if p.Reason != "" {
		q.Set("reason", p.Reason)
	}
	return q
}

// withMarker copies p with the pagination cursor replaced, leaving the
// caller's value untouched across pages.
func (p *ListSuppressionsParams) withMarker(marker string) *ListSuppressionsParams {
	var out ListSuppressionsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListTemplatesParams are the optional filters and pagination controls for
// [Client.ListTemplates]. A nil *ListTemplatesParams sends none of them.
type ListTemplatesParams struct {
	// CRN Exact CRN filter. Malformed CRNs return 400; foreign or mismatched
	// CRNs return an empty page.
	CRN    string
	Limit  int
	Marker string

	// Name exact immutable account-scoped name, combined with all other
	// filters.
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListTemplatesParams) query() url.Values {
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
func (p *ListTemplatesParams) withMarker(marker string) *ListTemplatesParams {
	var out ListTemplatesParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// AddSuppressionOverride adds a suppression override address.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) AddSuppressionOverride(ctx context.Context, configurationSetID string, body *SuppressionOverrideRequest, opts ...basaltic.RequestOption) (*ConfigurationSet, error) {
	op := &basaltic.Operation{
		ID:       "addSuppressionOverride",
		Method:   "POST",
		Path:     "/v1/configuration-sets/{configuration_set_id}/suppression-overrides",
		PathArgs: []string{configurationSetID},
		Body:     body,
	}
	var out struct {
		ConfigurationSet *ConfigurationSet `json:"configuration_set"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.ConfigurationSet, nil
}

// BeginDkimRotation begins DKIM key rotation.
//
// Mints a pending DKIM keypair and returns dns_records including the new
// dkim_pending TXT the customer must publish before
// FinalizeDKIMRotation.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) BeginDkimRotation(ctx context.Context, identityID string, opts ...basaltic.RequestOption) (*Identity, error) {
	op := &basaltic.Operation{
		ID:       "beginDKIMRotation",
		Method:   "POST",
		Path:     "/v1/identities/{identity_id}/dkim/rotate",
		PathArgs: []string{identityID},
	}
	var out struct {
		Identity *Identity `json:"identity"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Identity, nil
}

// CancelDkimRotation cancels in-progress DKIM rotation.
func (c *Client) CancelDkimRotation(ctx context.Context, identityID string, opts ...basaltic.RequestOption) (*Identity, error) {
	op := &basaltic.Operation{
		ID:       "cancelDKIMRotation",
		Method:   "DELETE",
		Path:     "/v1/identities/{identity_id}/dkim/rotate",
		PathArgs: []string{identityID},
	}
	var out struct {
		Identity *Identity `json:"identity"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Identity, nil
}

// CreateConfigurationSet creates a configuration set.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateConfigurationSet(ctx context.Context, body *CreateConfigurationSetRequest, opts ...basaltic.RequestOption) (*ConfigurationSet, error) {
	op := &basaltic.Operation{
		ID:     "createConfigurationSet",
		Method: "POST",
		Path:   "/v1/configuration-sets",
		Body:   body,
	}
	var out struct {
		ConfigurationSet *ConfigurationSet `json:"configuration_set"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.ConfigurationSet, nil
}

// CreateIPPool creates an IP pool.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateIPPool(ctx context.Context, body *CreateIPPoolRequest, opts ...basaltic.RequestOption) (*IPPool, error) {
	op := &basaltic.Operation{
		ID:     "createIPPool",
		Method: "POST",
		Path:   "/v1/ip-pools",
		Body:   body,
	}
	var out struct {
		IPPool *IPPool `json:"ip_pool"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.IPPool, nil
}

// CreateIdentity creates a sender identity.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateIdentity(ctx context.Context, body *CreateIdentityRequest, opts ...basaltic.RequestOption) (*Identity, error) {
	op := &basaltic.Operation{
		ID:     "createIdentity",
		Method: "POST",
		Path:   "/v1/identities",
		Body:   body,
	}
	var out struct {
		Identity *Identity `json:"identity"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Identity, nil
}

// CreateRecipientRule creates recipient rule.
func (c *Client) CreateRecipientRule(ctx context.Context, identityID string, body *CreateRecipientRuleRequest, opts ...basaltic.RequestOption) (*RecipientRule, error) {
	op := &basaltic.Operation{
		ID:       "createRecipientRule",
		Method:   "POST",
		Path:     "/v1/identities/{identity_id}/recipient-rules",
		PathArgs: []string{identityID},
		Body:     body,
	}
	var out struct {
		Rule *RecipientRule `json:"rule"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Rule, nil
}

// CreateSuppression adds a suppression entry.
//
// Create a manual suppression. An existing address returns 409 Conflict;
// use PATCH on its UUID to change description or expiry.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateSuppression(ctx context.Context, body *CreateSuppressionRequest, opts ...basaltic.RequestOption) (*Suppression, error) {
	op := &basaltic.Operation{
		ID:     "createSuppression",
		Method: "POST",
		Path:   "/v1/suppressions",
		Body:   body,
	}
	var out struct {
		Suppression *Suppression `json:"suppression"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Suppression, nil
}

// CreateTemplate creates a template.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateTemplate(ctx context.Context, body *CreateTemplateRequest, opts ...basaltic.RequestOption) (*Template, error) {
	op := &basaltic.Operation{
		ID:     "createTemplate",
		Method: "POST",
		Path:   "/v1/templates",
		Body:   body,
	}
	var out struct {
		Template *Template `json:"template"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Template, nil
}

// DeleteConfigurationSet deletes a configuration set.
func (c *Client) DeleteConfigurationSet(ctx context.Context, configurationSetID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteConfigurationSet",
		Method:   "DELETE",
		Path:     "/v1/configuration-sets/{configuration_set_id}",
		PathArgs: []string{configurationSetID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteIPPool deletes an IP pool.
func (c *Client) DeleteIPPool(ctx context.Context, ipPoolID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteIPPool",
		Method:   "DELETE",
		Path:     "/v1/ip-pools/{ip_pool_id}",
		PathArgs: []string{ipPoolID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteIdentity deletes an identity.
//
// Starts the teardown: the DKIM key or keys are freed, the hosted
// Easy-DKIM TXT and the DMARC authorization are withdrawn, quota is
// released, and the identity is removed last.
//
// Asynchronous, which is why this answers `202` rather than `204`. The
// identity is usually still readable for a moment after this returns —
// poll until a describe answers `404` rather than treating the response
// as proof the deletion has already happened.
func (c *Client) DeleteIdentity(ctx context.Context, identityID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteIdentity",
		Method:   "DELETE",
		Path:     "/v1/identities/{identity_id}",
		PathArgs: []string{identityID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteRecipientRule deletes recipient rule.
func (c *Client) DeleteRecipientRule(ctx context.Context, identityID string, ruleID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteRecipientRule",
		Method:   "DELETE",
		Path:     "/v1/identities/{identity_id}/recipient-rules/{rule_id}",
		PathArgs: []string{identityID, ruleID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteSuppression deletes suppression.
func (c *Client) DeleteSuppression(ctx context.Context, suppressionID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteSuppression",
		Method:   "DELETE",
		Path:     "/v1/suppressions/{suppression_id}",
		PathArgs: []string{suppressionID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteTemplate deletes a template.
func (c *Client) DeleteTemplate(ctx context.Context, templateID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteTemplate",
		Method:   "DELETE",
		Path:     "/v1/templates/{template_id}",
		PathArgs: []string{templateID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DescribeConfigurationSet describes a configuration set.
func (c *Client) DescribeConfigurationSet(ctx context.Context, configurationSetID string, opts ...basaltic.RequestOption) (*ConfigurationSet, error) {
	op := &basaltic.Operation{
		ID:       "describeConfigurationSet",
		Method:   "GET",
		Path:     "/v1/configuration-sets/{configuration_set_id}",
		PathArgs: []string{configurationSetID},
	}
	var out struct {
		ConfigurationSet *ConfigurationSet `json:"configuration_set"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.ConfigurationSet, nil
}

// DescribeDedicatedIP describes a dedicated IP.
func (c *Client) DescribeDedicatedIP(ctx context.Context, dedicatedIPID string, opts ...basaltic.RequestOption) (*DedicatedIP, error) {
	op := &basaltic.Operation{
		ID:       "describeDedicatedIP",
		Method:   "GET",
		Path:     "/v1/dedicated-ips/{dedicated_ip_id}",
		PathArgs: []string{dedicatedIPID},
	}
	var out struct {
		DedicatedIP *DedicatedIP `json:"dedicated_ip"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.DedicatedIP, nil
}

// DescribeIPPool describes an IP pool.
func (c *Client) DescribeIPPool(ctx context.Context, ipPoolID string, opts ...basaltic.RequestOption) (*IPPool, error) {
	op := &basaltic.Operation{
		ID:       "describeIPPool",
		Method:   "GET",
		Path:     "/v1/ip-pools/{ip_pool_id}",
		PathArgs: []string{ipPoolID},
	}
	var out struct {
		IPPool *IPPool `json:"ip_pool"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.IPPool, nil
}

// DescribeIdentity describes an identity.
func (c *Client) DescribeIdentity(ctx context.Context, identityID string, opts ...basaltic.RequestOption) (*Identity, error) {
	op := &basaltic.Operation{
		ID:       "describeIdentity",
		Method:   "GET",
		Path:     "/v1/identities/{identity_id}",
		PathArgs: []string{identityID},
	}
	var out struct {
		Identity *Identity `json:"identity"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Identity, nil
}

// DescribeReceivedMessage gets received message metadata.
func (c *Client) DescribeReceivedMessage(ctx context.Context, messageID string, opts ...basaltic.RequestOption) (*ReceivedMessage, error) {
	op := &basaltic.Operation{
		ID:       "describeReceivedMessage",
		Method:   "GET",
		Path:     "/v1/received-messages/{message_id}",
		PathArgs: []string{messageID},
	}
	var out struct {
		Message *ReceivedMessage `json:"message"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Message, nil
}

// DescribeRecipientRule gets recipient rule.
func (c *Client) DescribeRecipientRule(ctx context.Context, identityID string, ruleID string, opts ...basaltic.RequestOption) (*RecipientRule, error) {
	op := &basaltic.Operation{
		ID:       "describeRecipientRule",
		Method:   "GET",
		Path:     "/v1/identities/{identity_id}/recipient-rules/{rule_id}",
		PathArgs: []string{identityID, ruleID},
	}
	var out struct {
		Rule *RecipientRule `json:"rule"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Rule, nil
}

// DescribeSentMessage describes a sent message.
func (c *Client) DescribeSentMessage(ctx context.Context, messageID string, opts ...basaltic.RequestOption) (*Message, error) {
	op := &basaltic.Operation{
		ID:       "describeSentMessage",
		Method:   "GET",
		Path:     "/v1/sent-messages/{message_id}",
		PathArgs: []string{messageID},
	}
	var out struct {
		Message *Message `json:"message"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Message, nil
}

// DescribeSuppression gets suppression.
func (c *Client) DescribeSuppression(ctx context.Context, suppressionID string, opts ...basaltic.RequestOption) (*Suppression, error) {
	op := &basaltic.Operation{
		ID:       "describeSuppression",
		Method:   "GET",
		Path:     "/v1/suppressions/{suppression_id}",
		PathArgs: []string{suppressionID},
	}
	var out struct {
		Suppression *Suppression `json:"suppression"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Suppression, nil
}

// DescribeTemplate describes a template.
func (c *Client) DescribeTemplate(ctx context.Context, templateID string, opts ...basaltic.RequestOption) (*Template, error) {
	op := &basaltic.Operation{
		ID:       "describeTemplate",
		Method:   "GET",
		Path:     "/v1/templates/{template_id}",
		PathArgs: []string{templateID},
	}
	var out struct {
		Template *Template `json:"template"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Template, nil
}

// DownloadReceivedAttachment downloads received attachment.
//
// The caller must close the returned reader.
func (c *Client) DownloadReceivedAttachment(ctx context.Context, messageID string, attachmentID string, opts ...basaltic.RequestOption) (io.ReadCloser, error) {
	op := &basaltic.Operation{
		ID:       "downloadReceivedAttachment",
		Method:   "GET",
		Path:     "/v1/received-messages/{message_id}/attachments/{attachment_id}",
		PathArgs: []string{messageID, attachmentID},
	}
	stream, _, err := c.rt.DoStream(ctx, op, opts...)
	if err != nil {
		return nil, err
	}
	return stream, nil
}

// DownloadReceivedMessage downloads received message.
//
// The caller must close the returned reader.
func (c *Client) DownloadReceivedMessage(ctx context.Context, messageID string, opts ...basaltic.RequestOption) (io.ReadCloser, error) {
	op := &basaltic.Operation{
		ID:       "downloadReceivedMessage",
		Method:   "GET",
		Path:     "/v1/received-messages/{message_id}/raw",
		PathArgs: []string{messageID},
	}
	stream, _, err := c.rt.DoStream(ctx, op, opts...)
	if err != nil {
		return nil, err
	}
	return stream, nil
}

// DownloadSentAttachment downloads sent attachment.
//
// The caller must close the returned reader.
func (c *Client) DownloadSentAttachment(ctx context.Context, messageID string, attachmentID string, opts ...basaltic.RequestOption) (io.ReadCloser, error) {
	op := &basaltic.Operation{
		ID:       "downloadSentAttachment",
		Method:   "GET",
		Path:     "/v1/sent-messages/{message_id}/attachments/{attachment_id}",
		PathArgs: []string{messageID, attachmentID},
	}
	stream, _, err := c.rt.DoStream(ctx, op, opts...)
	if err != nil {
		return nil, err
	}
	return stream, nil
}

// DownloadSentMessage downloads sent message.
//
// The caller must close the returned reader.
func (c *Client) DownloadSentMessage(ctx context.Context, messageID string, opts ...basaltic.RequestOption) (io.ReadCloser, error) {
	op := &basaltic.Operation{
		ID:       "downloadSentMessage",
		Method:   "GET",
		Path:     "/v1/sent-messages/{message_id}/raw",
		PathArgs: []string{messageID},
	}
	stream, _, err := c.rt.DoStream(ctx, op, opts...)
	if err != nil {
		return nil, err
	}
	return stream, nil
}

// FinalizeDkimRotation finalizes DKIM key rotation.
//
// Verifies the pending DKIM TXT is published and flips the active
// selector. The previous selector enters an overlap window.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) FinalizeDkimRotation(ctx context.Context, identityID string, opts ...basaltic.RequestOption) (*Identity, error) {
	op := &basaltic.Operation{
		ID:       "finalizeDKIMRotation",
		Method:   "POST",
		Path:     "/v1/identities/{identity_id}/dkim/rotate/finalize",
		PathArgs: []string{identityID},
	}
	var out struct {
		Identity *Identity `json:"identity"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Identity, nil
}

// GetReceivedMessageContent gets received message content.
//
// Returns full decoded text and HTML, plus attachment metadata. Content
// expires after 30 days. Malware content is inaccessible.
func (c *Client) GetReceivedMessageContent(ctx context.Context, messageID string, opts ...basaltic.RequestOption) (*MessageContent, error) {
	op := &basaltic.Operation{
		ID:       "getReceivedMessageContent",
		Method:   "GET",
		Path:     "/v1/received-messages/{message_id}/content",
		PathArgs: []string{messageID},
	}
	var out struct {
		Content *MessageContent `json:"content"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Content, nil
}

// GetReceivingConfiguration gets receiving configuration.
//
// Returns receiving state and the selected attachment bucket. DNS
// records are returned on the identity. Message content is retained for
// 30 days; retention is fixed. Organization quotas control traffic and
// message size.
func (c *Client) GetReceivingConfiguration(ctx context.Context, identityID string, opts ...basaltic.RequestOption) (*ReceivingConfiguration, error) {
	op := &basaltic.Operation{
		ID:       "getReceivingConfiguration",
		Method:   "GET",
		Path:     "/v1/identities/{identity_id}/receiving",
		PathArgs: []string{identityID},
	}
	var out struct {
		Receiving *ReceivingConfiguration `json:"receiving"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Receiving, nil
}

// GetSentMessageContent gets sent message content.
//
// Returns full decoded text and HTML, plus attachment metadata. Content
// expires after 30 days. Malware content is inaccessible.
func (c *Client) GetSentMessageContent(ctx context.Context, messageID string, opts ...basaltic.RequestOption) (*MessageContent, error) {
	op := &basaltic.Operation{
		ID:       "getSentMessageContent",
		Method:   "GET",
		Path:     "/v1/sent-messages/{message_id}/content",
		PathArgs: []string{messageID},
	}
	var out struct {
		Content *MessageContent `json:"content"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Content, nil
}

// GetSentMessageEngagement — Sent message engagement.
func (c *Client) GetSentMessageEngagement(ctx context.Context, messageID string, opts ...basaltic.RequestOption) (*MessageEngagement, error) {
	op := &basaltic.Operation{
		ID:       "getSentMessageEngagement",
		Method:   "GET",
		Path:     "/v1/sent-messages/{message_id}/engagement",
		PathArgs: []string{messageID},
	}
	var out struct {
		Engagement *MessageEngagement `json:"engagement"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Engagement, nil
}

// ListConfigurationSets lists configuration sets.
//
// Returns one page. Use ListConfigurationSetsAll to walk every page.
func (c *Client) ListConfigurationSets(ctx context.Context, params *ListConfigurationSetsParams, opts ...basaltic.RequestOption) (*basaltic.Page[ConfigurationSet], error) {
	op := &basaltic.Operation{
		ID:     "listConfigurationSets",
		Method: "GET",
		Path:   "/v1/configuration-sets",
	}
	op.Query = params.query()
	var out struct {
		Items []ConfigurationSet `json:"configuration_sets"`
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
	page := &basaltic.Page[ConfigurationSet]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListConfigurationSetsAll walks every page of ListConfigurationSets,
// yielding one item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListConfigurationSetsAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListConfigurationSetsAll(ctx context.Context, params *ListConfigurationSetsParams, opts ...basaltic.RequestOption) iter.Seq2[ConfigurationSet, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[ConfigurationSet], error) {
		return c.ListConfigurationSets(ctx, params.withMarker(marker), opts...)
	})
}

// ListDedicatedIPs lists dedicated IPs.
//
// Returns one page. Use ListDedicatedIPsAll to walk every page.
func (c *Client) ListDedicatedIPs(ctx context.Context, params *ListDedicatedIPsParams, opts ...basaltic.RequestOption) (*basaltic.Page[DedicatedIP], error) {
	op := &basaltic.Operation{
		ID:     "listDedicatedIPs",
		Method: "GET",
		Path:   "/v1/dedicated-ips",
	}
	op.Query = params.query()
	var out struct {
		Items []DedicatedIP `json:"dedicated_ips"`
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
	page := &basaltic.Page[DedicatedIP]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListDedicatedIPsAll walks every page of ListDedicatedIPs, yielding one
// item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListDedicatedIPsAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListDedicatedIPsAll(ctx context.Context, params *ListDedicatedIPsParams, opts ...basaltic.RequestOption) iter.Seq2[DedicatedIP, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[DedicatedIP], error) {
		return c.ListDedicatedIPs(ctx, params.withMarker(marker), opts...)
	})
}

// ListDmarcReports lists DMARC report records.
//
// Diagnostic report records, retained for 180 days. Filter by the
// reported period end; default is the last seven days, maximum interval
// 180 days. Exact retransmissions are deduplicated by reporter, report,
// interval and source/authentication attributes. Late reports may appear
// on later reads. Use Telemetry metrics for aggregate charts.
//
// Returns one page. Use ListDmarcReportsAll to walk every page.
func (c *Client) ListDmarcReports(ctx context.Context, identityID string, params *ListDmarcReportsParams, opts ...basaltic.RequestOption) (*basaltic.Page[DmarcReportRecord], error) {
	op := &basaltic.Operation{
		ID:       "listDMARCReports",
		Method:   "GET",
		Path:     "/v1/identities/{identity_id}/dmarc-reports",
		PathArgs: []string{identityID},
	}
	op.Query = params.query()
	var out struct {
		Items []DmarcReportRecord `json:"records"`
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
	page := &basaltic.Page[DmarcReportRecord]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListDmarcReportsAll walks every page of ListDmarcReports, yielding one
// item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListDmarcReportsAll(ctx, identityID, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListDmarcReportsAll(ctx context.Context, identityID string, params *ListDmarcReportsParams, opts ...basaltic.RequestOption) iter.Seq2[DmarcReportRecord, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[DmarcReportRecord], error) {
		return c.ListDmarcReports(ctx, identityID, params.withMarker(marker), opts...)
	})
}

// ListIPPools lists IP pools.
//
// Returns one page. Use ListIPPoolsAll to walk every page.
func (c *Client) ListIPPools(ctx context.Context, params *ListIPPoolsParams, opts ...basaltic.RequestOption) (*basaltic.Page[IPPool], error) {
	op := &basaltic.Operation{
		ID:     "listIPPools",
		Method: "GET",
		Path:   "/v1/ip-pools",
	}
	op.Query = params.query()
	var out struct {
		Items []IPPool `json:"ip_pools"`
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
	page := &basaltic.Page[IPPool]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListIPPoolsAll walks every page of ListIPPools, yielding one item at a
// time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListIPPoolsAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListIPPoolsAll(ctx context.Context, params *ListIPPoolsParams, opts ...basaltic.RequestOption) iter.Seq2[IPPool, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[IPPool], error) {
		return c.ListIPPools(ctx, params.withMarker(marker), opts...)
	})
}

// ListIdentities lists sender identities.
//
// Returns one page. Use ListIdentitiesAll to walk every page.
func (c *Client) ListIdentities(ctx context.Context, params *ListIdentitiesParams, opts ...basaltic.RequestOption) (*basaltic.Page[Identity], error) {
	op := &basaltic.Operation{
		ID:     "listIdentities",
		Method: "GET",
		Path:   "/v1/identities",
	}
	op.Query = params.query()
	var out struct {
		Items []Identity `json:"identities"`
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
	page := &basaltic.Page[Identity]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListIdentitiesAll walks every page of ListIdentities, yielding one
// item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListIdentitiesAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListIdentitiesAll(ctx context.Context, params *ListIdentitiesParams, opts ...basaltic.RequestOption) iter.Seq2[Identity, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Identity], error) {
		return c.ListIdentities(ctx, params.withMarker(marker), opts...)
	})
}

// ListReceivedMessages lists received messages.
//
// Account-scoped metadata, ordered by received_at descending, then ID
// descending. Marker is the last ID of the previous page. Filters
// combine before pagination. No body or attachment bytes are included.
//
// Returns one page. Use ListReceivedMessagesAll to walk every page.
func (c *Client) ListReceivedMessages(ctx context.Context, params *ListReceivedMessagesParams, opts ...basaltic.RequestOption) (*basaltic.Page[ReceivedMessage], error) {
	op := &basaltic.Operation{
		ID:     "listReceivedMessages",
		Method: "GET",
		Path:   "/v1/received-messages",
	}
	op.Query = params.query()
	var out struct {
		Items []ReceivedMessage `json:"messages"`
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
	page := &basaltic.Page[ReceivedMessage]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListReceivedMessagesAll walks every page of ListReceivedMessages,
// yielding one item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListReceivedMessagesAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListReceivedMessagesAll(ctx context.Context, params *ListReceivedMessagesParams, opts ...basaltic.RequestOption) iter.Seq2[ReceivedMessage, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[ReceivedMessage], error) {
		return c.ListReceivedMessages(ctx, params.withMarker(marker), opts...)
	})
}

// ListRecipientRules lists recipient rules.
//
// Returns one page. Use ListRecipientRulesAll to walk every page.
func (c *Client) ListRecipientRules(ctx context.Context, identityID string, params *ListRecipientRulesParams, opts ...basaltic.RequestOption) (*basaltic.Page[RecipientRule], error) {
	op := &basaltic.Operation{
		ID:       "listRecipientRules",
		Method:   "GET",
		Path:     "/v1/identities/{identity_id}/recipient-rules",
		PathArgs: []string{identityID},
	}
	op.Query = params.query()
	var out struct {
		Items []RecipientRule `json:"rules"`
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
	page := &basaltic.Page[RecipientRule]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListRecipientRulesAll walks every page of ListRecipientRules, yielding
// one item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListRecipientRulesAll(ctx, identityID, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListRecipientRulesAll(ctx context.Context, identityID string, params *ListRecipientRulesParams, opts ...basaltic.RequestOption) iter.Seq2[RecipientRule, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[RecipientRule], error) {
		return c.ListRecipientRules(ctx, identityID, params.withMarker(marker), opts...)
	})
}

// ListRejections lists API-side send rejections.
//
// Newest first, keyset-paginated on `rejected_at` rather than on an id
// — the rejection log lives in ClickHouse, which has no id index to
// walk.
//
// `meta.marker` therefore carries a timestamp, and it goes back as
// `before`, not as `marker`: this endpoint has no `marker` parameter.
func (c *Client) ListRejections(ctx context.Context, params *ListRejectionsParams, opts ...basaltic.RequestOption) (*basaltic.Page[Rejection], error) {
	op := &basaltic.Operation{
		ID:     "listRejections",
		Method: "GET",
		Path:   "/v1/rejections",
	}
	op.Query = params.query()
	var out struct {
		Items []Rejection `json:"rejections"`
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
	page := &basaltic.Page[Rejection]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListSentMessages lists sent messages.
//
// Returns one page. Use ListSentMessagesAll to walk every page.
func (c *Client) ListSentMessages(ctx context.Context, params *ListSentMessagesParams, opts ...basaltic.RequestOption) (*basaltic.Page[Message], error) {
	op := &basaltic.Operation{
		ID:     "listSentMessages",
		Method: "GET",
		Path:   "/v1/sent-messages",
	}
	op.Query = params.query()
	var out struct {
		Items []Message `json:"messages"`
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
	page := &basaltic.Page[Message]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListSentMessagesAll walks every page of ListSentMessages, yielding one
// item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListSentMessagesAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListSentMessagesAll(ctx context.Context, params *ListSentMessagesParams, opts ...basaltic.RequestOption) iter.Seq2[Message, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Message], error) {
		return c.ListSentMessages(ctx, params.withMarker(marker), opts...)
	})
}

// ListSuppressions lists suppression entries.
//
// Returns one page. Use ListSuppressionsAll to walk every page.
func (c *Client) ListSuppressions(ctx context.Context, params *ListSuppressionsParams, opts ...basaltic.RequestOption) (*basaltic.Page[Suppression], error) {
	op := &basaltic.Operation{
		ID:     "listSuppressions",
		Method: "GET",
		Path:   "/v1/suppressions",
	}
	op.Query = params.query()
	var out struct {
		Items []Suppression `json:"suppressions"`
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
	page := &basaltic.Page[Suppression]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListSuppressionsAll walks every page of ListSuppressions, yielding one
// item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListSuppressionsAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListSuppressionsAll(ctx context.Context, params *ListSuppressionsParams, opts ...basaltic.RequestOption) iter.Seq2[Suppression, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Suppression], error) {
		return c.ListSuppressions(ctx, params.withMarker(marker), opts...)
	})
}

// ListTemplates lists templates.
//
// Returns one page. Use ListTemplatesAll to walk every page.
func (c *Client) ListTemplates(ctx context.Context, params *ListTemplatesParams, opts ...basaltic.RequestOption) (*basaltic.Page[Template], error) {
	op := &basaltic.Operation{
		ID:     "listTemplates",
		Method: "GET",
		Path:   "/v1/templates",
	}
	op.Query = params.query()
	var out struct {
		Items []Template `json:"templates"`
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
	page := &basaltic.Page[Template]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListTemplatesAll walks every page of ListTemplates, yielding one item
// at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListTemplatesAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListTemplatesAll(ctx context.Context, params *ListTemplatesParams, opts ...basaltic.RequestOption) iter.Seq2[Template, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Template], error) {
		return c.ListTemplates(ctx, params.withMarker(marker), opts...)
	})
}

// ReleaseDedicatedIP releases a dedicated IP.
func (c *Client) ReleaseDedicatedIP(ctx context.Context, dedicatedIPID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "releaseDedicatedIP",
		Method:   "DELETE",
		Path:     "/v1/dedicated-ips/{dedicated_ip_id}",
		PathArgs: []string{dedicatedIPID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// RemoveSuppressionOverride removes a suppression override address.
func (c *Client) RemoveSuppressionOverride(ctx context.Context, configurationSetID string, address string, opts ...basaltic.RequestOption) (*ConfigurationSet, error) {
	op := &basaltic.Operation{
		ID:       "removeSuppressionOverride",
		Method:   "DELETE",
		Path:     "/v1/configuration-sets/{configuration_set_id}/suppression-overrides/{address}",
		PathArgs: []string{configurationSetID, address},
	}
	var out struct {
		ConfigurationSet *ConfigurationSet `json:"configuration_set"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.ConfigurationSet, nil
}

// RequestDedicatedIP requests a dedicated IP.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) RequestDedicatedIP(ctx context.Context, body *RequestDedicatedIPRequest, opts ...basaltic.RequestOption) (*DedicatedIP, error) {
	op := &basaltic.Operation{
		ID:     "requestDedicatedIP",
		Method: "POST",
		Path:   "/v1/dedicated-ips",
		Body:   body,
	}
	var out struct {
		DedicatedIP *DedicatedIP `json:"dedicated_ip"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.DedicatedIP, nil
}

// SendMessage sends a message.
//
// Accept a message using explicit text/HTML, raw MIME, or template
// content. A successful response means queued, not delivered.
// Organization quotas apply to MIME size and recipients. Template
// content also requires email:SendTemplatedMessage on the template. Use
// Idempotency-Key for safe retries.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) SendMessage(ctx context.Context, body *SendMessageRequest, opts ...basaltic.RequestOption) (*Message, error) {
	op := &basaltic.Operation{
		ID:     "sendMessage",
		Method: "POST",
		Path:   "/v1/sent-messages",
		Body:   body,
	}
	var out struct {
		Message *Message `json:"message"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Message, nil
}

// SendMessageBatch sends message batch.
//
// Accept up to 50 messages. All identity/template permissions and
// content-mode shapes are checked before enqueueing; denial rejects the
// entire request. Sending restrictions, quota, suppression, rendering
// and storage failures are returned per entry, in input order. queued
// means accepted, not delivered. Use one Idempotency-Key for the entire
// batch; retry rejected entries separately with a new key.
func (c *Client) SendMessageBatch(ctx context.Context, body *BatchSendRequest, opts ...basaltic.RequestOption) (*basaltic.Page[BatchSendResult], error) {
	op := &basaltic.Operation{
		ID:     "sendMessageBatch",
		Method: "POST",
		Path:   "/v1/sent-messages/batch",
		Body:   body,
	}
	var out struct {
		Items []BatchSendResult `json:"results"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[BatchSendResult]{Items: out.Items}
	return page, nil
}

// SetBimi sets BIMI logo / VMC URLs.
//
// Update only the supplied settings. Omitted fields remain unchanged;
// null clears the supplied setting.
func (c *Client) SetBimi(ctx context.Context, identityID string, body *SetBimiRequest, opts ...basaltic.RequestOption) (*Identity, error) {
	op := &basaltic.Operation{
		ID:       "setBIMI",
		Method:   "PATCH",
		Path:     "/v1/identities/{identity_id}/bimi",
		PathArgs: []string{identityID},
		Body:     body,
	}
	var out struct {
		Identity *Identity `json:"identity"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Identity, nil
}

// SetMailFromDomain sets custom MAIL FROM domain.
//
// Update only the supplied settings. Omitted fields remain unchanged;
// null clears the supplied setting.
func (c *Client) SetMailFromDomain(ctx context.Context, identityID string, body *SetMailFromDomainRequest, opts ...basaltic.RequestOption) (*Identity, error) {
	op := &basaltic.Operation{
		ID:       "setMailFromDomain",
		Method:   "PATCH",
		Path:     "/v1/identities/{identity_id}/mail-from-domain",
		PathArgs: []string{identityID},
		Body:     body,
	}
	var out struct {
		Identity *Identity `json:"identity"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Identity, nil
}

// UpdateConfigurationSet updates a configuration set.
func (c *Client) UpdateConfigurationSet(ctx context.Context, configurationSetID string, body *UpdateConfigurationSetRequest, opts ...basaltic.RequestOption) (*ConfigurationSet, error) {
	op := &basaltic.Operation{
		ID:       "updateConfigurationSet",
		Method:   "PATCH",
		Path:     "/v1/configuration-sets/{configuration_set_id}",
		PathArgs: []string{configurationSetID},
		Body:     body,
	}
	var out struct {
		ConfigurationSet *ConfigurationSet `json:"configuration_set"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.ConfigurationSet, nil
}

// UpdateDedicatedIP updates a dedicated IP.
func (c *Client) UpdateDedicatedIP(ctx context.Context, dedicatedIPID string, body *UpdateDedicatedIPRequest, opts ...basaltic.RequestOption) (*DedicatedIP, error) {
	op := &basaltic.Operation{
		ID:       "updateDedicatedIP",
		Method:   "PATCH",
		Path:     "/v1/dedicated-ips/{dedicated_ip_id}",
		PathArgs: []string{dedicatedIPID},
		Body:     body,
	}
	var out struct {
		DedicatedIP *DedicatedIP `json:"dedicated_ip"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.DedicatedIP, nil
}

// UpdateIPPool updates an IP pool.
func (c *Client) UpdateIPPool(ctx context.Context, ipPoolID string, body *UpdateIPPoolRequest, opts ...basaltic.RequestOption) (*IPPool, error) {
	op := &basaltic.Operation{
		ID:       "updateIPPool",
		Method:   "PATCH",
		Path:     "/v1/ip-pools/{ip_pool_id}",
		PathArgs: []string{ipPoolID},
		Body:     body,
	}
	var out struct {
		IPPool *IPPool `json:"ip_pool"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.IPPool, nil
}

// UpdateIdentity updates identity.
//
// Replace identity tags. Omission preserves the existing tags; an empty
// object removes all tags.
func (c *Client) UpdateIdentity(ctx context.Context, identityID string, body *UpdateIdentityRequest, opts ...basaltic.RequestOption) (*Identity, error) {
	op := &basaltic.Operation{
		ID:       "updateIdentity",
		Method:   "PATCH",
		Path:     "/v1/identities/{identity_id}",
		PathArgs: []string{identityID},
		Body:     body,
	}
	var out struct {
		Identity *Identity `json:"identity"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Identity, nil
}

// UpdateReceivingConfiguration updates receiving configuration.
//
// Update enabled state or the attachment bucket. Omitted fields are
// preserved. attachments_bucket: null clears the bucket only when
// receiving is disabled. Enabling requires a verified identity and a
// bucket with delegated GetObject and PutObject access.
func (c *Client) UpdateReceivingConfiguration(ctx context.Context, identityID string, body *UpdateReceivingConfigurationRequest, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "updateReceivingConfiguration",
		Method:   "PATCH",
		Path:     "/v1/identities/{identity_id}/receiving",
		PathArgs: []string{identityID},
		Body:     body,
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// UpdateRecipientRule updates recipient rule.
func (c *Client) UpdateRecipientRule(ctx context.Context, identityID string, ruleID string, body *UpdateRecipientRuleRequest, opts ...basaltic.RequestOption) (*RecipientRule, error) {
	op := &basaltic.Operation{
		ID:       "updateRecipientRule",
		Method:   "PATCH",
		Path:     "/v1/identities/{identity_id}/recipient-rules/{rule_id}",
		PathArgs: []string{identityID, ruleID},
		Body:     body,
	}
	var out struct {
		Rule *RecipientRule `json:"rule"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Rule, nil
}

// UpdateSuppression updates suppression.
func (c *Client) UpdateSuppression(ctx context.Context, suppressionID string, body *UpdateSuppressionRequest, opts ...basaltic.RequestOption) (*Suppression, error) {
	op := &basaltic.Operation{
		ID:       "updateSuppression",
		Method:   "PATCH",
		Path:     "/v1/suppressions/{suppression_id}",
		PathArgs: []string{suppressionID},
		Body:     body,
	}
	var out struct {
		Suppression *Suppression `json:"suppression"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Suppression, nil
}

// UpdateTemplate updates a template.
func (c *Client) UpdateTemplate(ctx context.Context, templateID string, body *UpdateTemplateRequest, opts ...basaltic.RequestOption) (*Template, error) {
	op := &basaltic.Operation{
		ID:       "updateTemplate",
		Method:   "PATCH",
		Path:     "/v1/templates/{template_id}",
		PathArgs: []string{templateID},
		Body:     body,
	}
	var out struct {
		Template *Template `json:"template"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Template, nil
}

// VerifyBimi verifies BIMI logo (+ optional VMC).
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) VerifyBimi(ctx context.Context, identityID string, opts ...basaltic.RequestOption) (*Identity, error) {
	op := &basaltic.Operation{
		ID:       "verifyBIMI",
		Method:   "POST",
		Path:     "/v1/identities/{identity_id}/bimi/verify",
		PathArgs: []string{identityID},
	}
	var out struct {
		Identity *Identity `json:"identity"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Identity, nil
}

// VerifyIdentity triggers DNS verification.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) VerifyIdentity(ctx context.Context, identityID string, opts ...basaltic.RequestOption) (*Identity, error) {
	op := &basaltic.Operation{
		ID:       "verifyIdentity",
		Method:   "POST",
		Path:     "/v1/identities/{identity_id}/verify",
		PathArgs: []string{identityID},
	}
	var out struct {
		Identity *Identity `json:"identity"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Identity, nil
}

// VerifyMailFromDomain verifies custom MAIL FROM domain MX.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) VerifyMailFromDomain(ctx context.Context, identityID string, opts ...basaltic.RequestOption) (*Identity, error) {
	op := &basaltic.Operation{
		ID:       "verifyMailFromDomain",
		Method:   "POST",
		Path:     "/v1/identities/{identity_id}/mail-from-domain/verify",
		PathArgs: []string{identityID},
	}
	var out struct {
		Identity *Identity `json:"identity"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Identity, nil
}

// DescribeConfigurationSetByReference fetches one configuration set by
// an id, a CRN or a name.
//
// The reference is classified by its syntax alone, exactly as the
// platform does (see [basaltic.ParseReference]): an id is fetched with
// [Client.DescribeConfigurationSet]; a CRN or a name goes to
// [Client.ListConfigurationSets] as an exact filter, together with any
// filters already set on scope, which may be nil. A miss is a not-found
// error for the kind the string was read as — no other kind is tried
// — and more than one match is a [basaltic.AmbiguousReferenceError].
func (c *Client) DescribeConfigurationSetByReference(ctx context.Context, ref string, scope *ListConfigurationSetsParams, opts ...basaltic.RequestOption) (*ConfigurationSet, error) {
	return basaltic.ResolveByReference(ctx, ref, "configuration-set", "listConfigurationSets", true,
		func(ctx context.Context, refID string) (*ConfigurationSet, error) {
			return c.DescribeConfigurationSet(ctx, refID, opts...)
		},
		func(ctx context.Context, refName, refCRN string) (*basaltic.Page[ConfigurationSet], error) {
			var p ListConfigurationSetsParams
			if scope != nil {
				p = *scope
			}
			p.Name = refName
			p.CRN = refCRN
			p.Limit = 2
			return c.ListConfigurationSets(ctx, &p, opts...)
		})
}

// DescribeDedicatedIPByReference fetches one dedicated ip by an id or a
// CRN.
//
// The reference is classified by its syntax alone, exactly as the
// platform does (see [basaltic.ParseReference]): an id is fetched with
// [Client.DescribeDedicatedIP]; a CRN or a name goes to
// [Client.ListDedicatedIPs] as an exact filter, together with any
// filters already set on scope, which may be nil. A miss is a not-found
// error for the kind the string was read as — no other kind is tried
// — and more than one match is a [basaltic.AmbiguousReferenceError].
//
// This resource has no name; a name reference is refused before any request.
func (c *Client) DescribeDedicatedIPByReference(ctx context.Context, ref string, scope *ListDedicatedIPsParams, opts ...basaltic.RequestOption) (*DedicatedIP, error) {
	return basaltic.ResolveByReference(ctx, ref, "dedicated-ip", "listDedicatedIPs", false,
		func(ctx context.Context, refID string) (*DedicatedIP, error) {
			return c.DescribeDedicatedIP(ctx, refID, opts...)
		},
		func(ctx context.Context, refName, refCRN string) (*basaltic.Page[DedicatedIP], error) {
			var p ListDedicatedIPsParams
			if scope != nil {
				p = *scope
			}
			_ = refName
			p.CRN = refCRN
			p.Limit = 2
			return c.ListDedicatedIPs(ctx, &p, opts...)
		})
}

// DescribeIPPoolByReference fetches one ip pool by an id, a CRN or a
// name.
//
// The reference is classified by its syntax alone, exactly as the
// platform does (see [basaltic.ParseReference]): an id is fetched with
// [Client.DescribeIPPool]; a CRN or a name goes to [Client.ListIPPools]
// as an exact filter, together with any filters already set on scope,
// which may be nil. A miss is a not-found error for the kind the string
// was read as — no other kind is tried — and more than one match is
// a [basaltic.AmbiguousReferenceError].
func (c *Client) DescribeIPPoolByReference(ctx context.Context, ref string, scope *ListIPPoolsParams, opts ...basaltic.RequestOption) (*IPPool, error) {
	return basaltic.ResolveByReference(ctx, ref, "ip-pool", "listIPPools", true,
		func(ctx context.Context, refID string) (*IPPool, error) {
			return c.DescribeIPPool(ctx, refID, opts...)
		},
		func(ctx context.Context, refName, refCRN string) (*basaltic.Page[IPPool], error) {
			var p ListIPPoolsParams
			if scope != nil {
				p = *scope
			}
			p.Name = refName
			p.CRN = refCRN
			p.Limit = 2
			return c.ListIPPools(ctx, &p, opts...)
		})
}

// DescribeIdentityByReference fetches one identity by an id, a CRN or a
// name.
//
// The reference is classified by its syntax alone, exactly as the
// platform does (see [basaltic.ParseReference]): an id is fetched with
// [Client.DescribeIdentity]; a CRN or a name goes to
// [Client.ListIdentities] as an exact filter, together with any filters
// already set on scope, which may be nil. A miss is a not-found error
// for the kind the string was read as — no other kind is tried — and
// more than one match is a [basaltic.AmbiguousReferenceError].
func (c *Client) DescribeIdentityByReference(ctx context.Context, ref string, scope *ListIdentitiesParams, opts ...basaltic.RequestOption) (*Identity, error) {
	return basaltic.ResolveByReference(ctx, ref, "identity", "listIdentities", true,
		func(ctx context.Context, refID string) (*Identity, error) {
			return c.DescribeIdentity(ctx, refID, opts...)
		},
		func(ctx context.Context, refName, refCRN string) (*basaltic.Page[Identity], error) {
			var p ListIdentitiesParams
			if scope != nil {
				p = *scope
			}
			p.Name = refName
			p.CRN = refCRN
			p.Limit = 2
			return c.ListIdentities(ctx, &p, opts...)
		})
}

// DescribeSentMessageByReference fetches one sent message by an id or a
// CRN.
//
// The reference is classified by its syntax alone, exactly as the
// platform does (see [basaltic.ParseReference]): an id is fetched with
// [Client.DescribeSentMessage]; a CRN or a name goes to
// [Client.ListSentMessages] as an exact filter, together with any
// filters already set on scope, which may be nil. A miss is a not-found
// error for the kind the string was read as — no other kind is tried
// — and more than one match is a [basaltic.AmbiguousReferenceError].
//
// This resource has no name; a name reference is refused before any request.
//
// A name is unique only within its parent; fix it on scope (Identity) or
// the lookup can match more than one.
func (c *Client) DescribeSentMessageByReference(ctx context.Context, ref string, scope *ListSentMessagesParams, opts ...basaltic.RequestOption) (*Message, error) {
	return basaltic.ResolveByReference(ctx, ref, "sent-message", "listSentMessages", false,
		func(ctx context.Context, refID string) (*Message, error) {
			return c.DescribeSentMessage(ctx, refID, opts...)
		},
		func(ctx context.Context, refName, refCRN string) (*basaltic.Page[Message], error) {
			var p ListSentMessagesParams
			if scope != nil {
				p = *scope
			}
			_ = refName
			p.CRN = refCRN
			p.Limit = 2
			return c.ListSentMessages(ctx, &p, opts...)
		})
}

// DescribeSuppressionByReference fetches one suppression by an id or a
// CRN.
//
// The reference is classified by its syntax alone, exactly as the
// platform does (see [basaltic.ParseReference]): an id is fetched with
// [Client.DescribeSuppression]; a CRN or a name goes to
// [Client.ListSuppressions] as an exact filter, together with any
// filters already set on scope, which may be nil. A miss is a not-found
// error for the kind the string was read as — no other kind is tried
// — and more than one match is a [basaltic.AmbiguousReferenceError].
//
// This resource has no name; a name reference is refused before any request.
func (c *Client) DescribeSuppressionByReference(ctx context.Context, ref string, scope *ListSuppressionsParams, opts ...basaltic.RequestOption) (*Suppression, error) {
	return basaltic.ResolveByReference(ctx, ref, "suppression", "listSuppressions", false,
		func(ctx context.Context, refID string) (*Suppression, error) {
			return c.DescribeSuppression(ctx, refID, opts...)
		},
		func(ctx context.Context, refName, refCRN string) (*basaltic.Page[Suppression], error) {
			var p ListSuppressionsParams
			if scope != nil {
				p = *scope
			}
			_ = refName
			p.CRN = refCRN
			p.Limit = 2
			return c.ListSuppressions(ctx, &p, opts...)
		})
}

// DescribeTemplateByReference fetches one template by an id, a CRN or a
// name.
//
// The reference is classified by its syntax alone, exactly as the
// platform does (see [basaltic.ParseReference]): an id is fetched with
// [Client.DescribeTemplate]; a CRN or a name goes to
// [Client.ListTemplates] as an exact filter, together with any filters
// already set on scope, which may be nil. A miss is a not-found error
// for the kind the string was read as — no other kind is tried — and
// more than one match is a [basaltic.AmbiguousReferenceError].
func (c *Client) DescribeTemplateByReference(ctx context.Context, ref string, scope *ListTemplatesParams, opts ...basaltic.RequestOption) (*Template, error) {
	return basaltic.ResolveByReference(ctx, ref, "template", "listTemplates", true,
		func(ctx context.Context, refID string) (*Template, error) {
			return c.DescribeTemplate(ctx, refID, opts...)
		},
		func(ctx context.Context, refName, refCRN string) (*basaltic.Page[Template], error) {
			var p ListTemplatesParams
			if scope != nil {
				p = *scope
			}
			p.Name = refName
			p.CRN = refCRN
			p.Limit = 2
			return c.ListTemplates(ctx, &p, opts...)
		})
}
