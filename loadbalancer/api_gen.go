// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package loadbalancer

import (
	"context"
	"iter"
	"net/url"
	"strconv"

	basaltic "github.com/basaltic-sh/sdk-go"
)

// ListLoadBalancersParams are the optional filters and pagination controls for
// [Client.ListLoadBalancers]. A nil *ListLoadBalancersParams sends none of them.
type ListLoadBalancersParams struct {
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

	// Status one of: "provisioning", "active", "error", "deleting".
	Status string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListLoadBalancersParams) query() url.Values {
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
	if p.Status != "" {
		q.Set("status", p.Status)
	}
	return q
}

// withMarker copies p with the pagination cursor replaced, leaving the
// caller's value untouched across pages.
func (p *ListLoadBalancersParams) withMarker(marker string) *ListLoadBalancersParams {
	var out ListLoadBalancersParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListTargetGroupsParams are the optional filters and pagination controls for
// [Client.ListTargetGroups]. A nil *ListTargetGroupsParams sends none of them.
type ListTargetGroupsParams struct {
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

	// Protocol one of: "http", "https", "tcp", "udp".
	Protocol string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListTargetGroupsParams) query() url.Values {
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
	if p.Protocol != "" {
		q.Set("protocol", p.Protocol)
	}
	return q
}

// withMarker copies p with the pagination cursor replaced, leaving the
// caller's value untouched across pages.
func (p *ListTargetGroupsParams) withMarker(marker string) *ListTargetGroupsParams {
	var out ListTargetGroupsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// AttachListenerCertificate attaches an additional certificate to an HTTPS listener.
//
// Multi-SNI: the listener routes per-connection by matching the client's
// SNI against each attached cert's SAN. Pass is_default=true to make
// this the fallback cert (used when no SNI matches or the client omits
// SNI entirely); whatever was default before is demoted in the same
// transaction.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) AttachListenerCertificate(ctx context.Context, id string, listenerID string, body *AttachListenerCertificateRequest, opts ...basaltic.RequestOption) (*Listener, error) {
	op := &basaltic.Operation{
		ID:       "attachListenerCertificate",
		Method:   "POST",
		Path:     "/v1/load-balancers/{id}/listeners/{listener_id}/certificates",
		PathArgs: []string{id, listenerID},
		Body:     body,
	}
	var out struct {
		Listener *Listener `json:"listener"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Listener, nil
}

// AttachTarget attaches a target to this group.
//
// `target_ref` must match the group's `target_type`: an IP address for
// `ip`, a compute instance id for `instance`. A `function` group takes
// no targets yet, and a `pool`-mode group takes its backends from the
// instance pool — attaching to either is an error rather than a row
// nothing will ever route to.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) AttachTarget(ctx context.Context, id string, body *AttachTargetRequest, opts ...basaltic.RequestOption) (*Target, error) {
	op := &basaltic.Operation{
		ID:       "attachTarget",
		Method:   "POST",
		Path:     "/v1/target-groups/{id}/targets",
		PathArgs: []string{id},
		Body:     body,
	}
	var out struct {
		Target *Target `json:"target"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Target, nil
}

// CreateListener creates a listener on this load balancer.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateListener(ctx context.Context, id string, body *CreateListenerRequest, opts ...basaltic.RequestOption) (*Listener, error) {
	op := &basaltic.Operation{
		ID:       "createListener",
		Method:   "POST",
		Path:     "/v1/load-balancers/{id}/listeners",
		PathArgs: []string{id},
		Body:     body,
	}
	var out struct {
		Listener *Listener `json:"listener"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Listener, nil
}

// CreateLoadBalancer creates a load balancer.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateLoadBalancer(ctx context.Context, body *CreateLoadBalancerRequest, opts ...basaltic.RequestOption) (*LoadBalancer, error) {
	op := &basaltic.Operation{
		ID:     "createLoadBalancer",
		Method: "POST",
		Path:   "/v1/load-balancers",
		Body:   body,
	}
	var out struct {
		LoadBalancer *LoadBalancer `json:"load_balancer"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.LoadBalancer, nil
}

// CreateRule creates a routing rule on this listener (HTTP/HTTPS only).
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateRule(ctx context.Context, id string, listenerID string, body *CreateRuleRequest, opts ...basaltic.RequestOption) (*Rule, error) {
	op := &basaltic.Operation{
		ID:       "createRule",
		Method:   "POST",
		Path:     "/v1/load-balancers/{id}/listeners/{listener_id}/rules",
		PathArgs: []string{id, listenerID},
		Body:     body,
	}
	var out struct {
		Rule *Rule `json:"rule"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Rule, nil
}

// CreateTargetGroup creates a target group.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateTargetGroup(ctx context.Context, body *CreateTargetGroupRequest, opts ...basaltic.RequestOption) (*TargetGroup, error) {
	op := &basaltic.Operation{
		ID:     "createTargetGroup",
		Method: "POST",
		Path:   "/v1/target-groups",
		Body:   body,
	}
	var out struct {
		TargetGroup *TargetGroup `json:"target_group"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.TargetGroup, nil
}

// DeleteListener deletes a listener.
func (c *Client) DeleteListener(ctx context.Context, id string, listenerID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteListener",
		Method:   "DELETE",
		Path:     "/v1/load-balancers/{id}/listeners/{listener_id}",
		PathArgs: []string{id, listenerID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteLoadBalancer deletes a load balancer.
//
// Request teardown. Returns 202 Accepted — the delete is async: the
// load balancer transitions to `deleting` and stays readable while the
// teardown saga releases the OVN rows, the replica pool, the replica IAM
// role, and the VIP reservation, deleting the row last. Poll the load
// balancer until it answers 404 rather than treating this response as
// proof it is gone.
func (c *Client) DeleteLoadBalancer(ctx context.Context, id string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteLoadBalancer",
		Method:   "DELETE",
		Path:     "/v1/load-balancers/{id}",
		PathArgs: []string{id},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteRule deletes a rule (superseded).
//
// Deletes a rule without naming the listener that owns it — the same
// resource `DELETE
// /v1/load-balancers/{id}/listeners/{listener_id}/rules/{rule_id}`
// addresses, and every other rule verb is listener-nested. Use the
// nested form; this one remains for clients already on it.
//
// Because the path carries no listener, the service has to scan the load
// balancer's listeners to prove the rule belongs to it before deleting.
func (c *Client) DeleteRule(ctx context.Context, id string, ruleID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteRule",
		Method:   "DELETE",
		Path:     "/v1/load-balancers/{id}/rules/{rule_id}",
		PathArgs: []string{id, ruleID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteRuleInListener deletes a routing rule.
//
// Deletes a rule addressed under the listener that owns it. Prefer this
// over `DELETE /v1/load-balancers/{id}/rules/{rule_id}`, which addresses
// the same resource without naming its listener.
func (c *Client) DeleteRuleInListener(ctx context.Context, id string, listenerID string, ruleID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteRuleInListener",
		Method:   "DELETE",
		Path:     "/v1/load-balancers/{id}/listeners/{listener_id}/rules/{rule_id}",
		PathArgs: []string{id, listenerID, ruleID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteTargetGroup deletes a target group.
func (c *Client) DeleteTargetGroup(ctx context.Context, id string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteTargetGroup",
		Method:   "DELETE",
		Path:     "/v1/target-groups/{id}",
		PathArgs: []string{id},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DetachListenerCertificate detaches a certificate from an HTTPS listener.
//
// Refuses to remove the last cert on an HTTPS listener, or to remove the
// currently-default cert while siblings still exist — promote a
// replacement first via PATCH listener.
//
// A certificate CRN ends in `certificate/<name>`, so the slash has to be
// percent-encoded (`%2F`) to keep the CRN inside one path segment —
// sending it raw addresses a different, non-existent route.
func (c *Client) DetachListenerCertificate(ctx context.Context, id string, listenerID string, certificateCRN string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "detachListenerCertificate",
		Method:   "DELETE",
		Path:     "/v1/load-balancers/{id}/listeners/{listener_id}/certificates/{certificate_crn}",
		PathArgs: []string{id, listenerID, certificateCRN},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DetachTarget detaches a target.
func (c *Client) DetachTarget(ctx context.Context, id string, targetID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "detachTarget",
		Method:   "DELETE",
		Path:     "/v1/target-groups/{id}/targets/{target_id}",
		PathArgs: []string{id, targetID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// GetListener gets a listener.
func (c *Client) GetListener(ctx context.Context, id string, listenerID string, opts ...basaltic.RequestOption) (*Listener, error) {
	op := &basaltic.Operation{
		ID:       "getListener",
		Method:   "GET",
		Path:     "/v1/load-balancers/{id}/listeners/{listener_id}",
		PathArgs: []string{id, listenerID},
	}
	var out struct {
		Listener *Listener `json:"listener"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Listener, nil
}

// GetLoadBalancer gets a load balancer.
func (c *Client) GetLoadBalancer(ctx context.Context, id string, opts ...basaltic.RequestOption) (*LoadBalancer, error) {
	op := &basaltic.Operation{
		ID:       "getLoadBalancer",
		Method:   "GET",
		Path:     "/v1/load-balancers/{id}",
		PathArgs: []string{id},
	}
	var out struct {
		LoadBalancer *LoadBalancer `json:"load_balancer"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.LoadBalancer, nil
}

// GetRule gets a routing rule.
func (c *Client) GetRule(ctx context.Context, id string, listenerID string, ruleID string, opts ...basaltic.RequestOption) (*Rule, error) {
	op := &basaltic.Operation{
		ID:       "getRule",
		Method:   "GET",
		Path:     "/v1/load-balancers/{id}/listeners/{listener_id}/rules/{rule_id}",
		PathArgs: []string{id, listenerID, ruleID},
	}
	var out struct {
		Rule *Rule `json:"rule"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Rule, nil
}

// GetTarget gets a target.
func (c *Client) GetTarget(ctx context.Context, id string, targetID string, opts ...basaltic.RequestOption) (*Target, error) {
	op := &basaltic.Operation{
		ID:       "getTarget",
		Method:   "GET",
		Path:     "/v1/target-groups/{id}/targets/{target_id}",
		PathArgs: []string{id, targetID},
	}
	var out struct {
		Target *Target `json:"target"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Target, nil
}

// GetTargetGroup gets a target group.
func (c *Client) GetTargetGroup(ctx context.Context, id string, opts ...basaltic.RequestOption) (*TargetGroup, error) {
	op := &basaltic.Operation{
		ID:       "getTargetGroup",
		Method:   "GET",
		Path:     "/v1/target-groups/{id}",
		PathArgs: []string{id},
	}
	var out struct {
		TargetGroup *TargetGroup `json:"target_group"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.TargetGroup, nil
}

// ListListeners lists this load balancer's listeners.
func (c *Client) ListListeners(ctx context.Context, id string, opts ...basaltic.RequestOption) (*basaltic.Page[Listener], error) {
	op := &basaltic.Operation{
		ID:       "listListeners",
		Method:   "GET",
		Path:     "/v1/load-balancers/{id}/listeners",
		PathArgs: []string{id},
	}
	var out struct {
		Items []Listener `json:"listeners"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[Listener]{Items: out.Items}
	return page, nil
}

// ListLoadBalancerReplicas lists the LB's instance replicas with live health.
//
// Returns one entry per compute instance backing the LB. The per-replica
// `proxy_ok` + `last_seen` overlay is refreshed on every replica health
// report (typically every 15s) and is absent when a replica hasn't yet
// reported (boot still in flight).
func (c *Client) ListLoadBalancerReplicas(ctx context.Context, id string, opts ...basaltic.RequestOption) (*basaltic.Page[LoadBalancerReplica], error) {
	op := &basaltic.Operation{
		ID:       "listLoadBalancerReplicas",
		Method:   "GET",
		Path:     "/v1/load-balancers/{id}/replicas",
		PathArgs: []string{id},
	}
	var out struct {
		Items []LoadBalancerReplica `json:"replicas"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[LoadBalancerReplica]{Items: out.Items}
	return page, nil
}

// ListLoadBalancers lists load balancers.
//
// Returns one page. Use ListLoadBalancersAll to walk every page.
func (c *Client) ListLoadBalancers(ctx context.Context, params *ListLoadBalancersParams, opts ...basaltic.RequestOption) (*basaltic.Page[LoadBalancer], error) {
	op := &basaltic.Operation{
		ID:     "listLoadBalancers",
		Method: "GET",
		Path:   "/v1/load-balancers",
	}
	op.Query = params.query()
	var out struct {
		Items []LoadBalancer `json:"load_balancers"`
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
	page := &basaltic.Page[LoadBalancer]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListLoadBalancersAll walks every page of ListLoadBalancers, yielding
// one item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListLoadBalancersAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListLoadBalancersAll(ctx context.Context, params *ListLoadBalancersParams, opts ...basaltic.RequestOption) iter.Seq2[LoadBalancer, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[LoadBalancer], error) {
		return c.ListLoadBalancers(ctx, params.withMarker(marker), opts...)
	})
}

// ListRules lists this listener's rules.
func (c *Client) ListRules(ctx context.Context, id string, listenerID string, opts ...basaltic.RequestOption) (*basaltic.Page[Rule], error) {
	op := &basaltic.Operation{
		ID:       "listRules",
		Method:   "GET",
		Path:     "/v1/load-balancers/{id}/listeners/{listener_id}/rules",
		PathArgs: []string{id, listenerID},
	}
	var out struct {
		Items []Rule `json:"rules"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[Rule]{Items: out.Items}
	return page, nil
}

// ListTargetGroups lists target groups.
//
// Returns one page. Use ListTargetGroupsAll to walk every page.
func (c *Client) ListTargetGroups(ctx context.Context, params *ListTargetGroupsParams, opts ...basaltic.RequestOption) (*basaltic.Page[TargetGroup], error) {
	op := &basaltic.Operation{
		ID:     "listTargetGroups",
		Method: "GET",
		Path:   "/v1/target-groups",
	}
	op.Query = params.query()
	var out struct {
		Items []TargetGroup `json:"target_groups"`
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
	page := &basaltic.Page[TargetGroup]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListTargetGroupsAll walks every page of ListTargetGroups, yielding one
// item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListTargetGroupsAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListTargetGroupsAll(ctx context.Context, params *ListTargetGroupsParams, opts ...basaltic.RequestOption) iter.Seq2[TargetGroup, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[TargetGroup], error) {
		return c.ListTargetGroups(ctx, params.withMarker(marker), opts...)
	})
}

// ListTargets lists targets in this group.
func (c *Client) ListTargets(ctx context.Context, id string, opts ...basaltic.RequestOption) (*basaltic.Page[Target], error) {
	op := &basaltic.Operation{
		ID:       "listTargets",
		Method:   "GET",
		Path:     "/v1/target-groups/{id}/targets",
		PathArgs: []string{id},
	}
	var out struct {
		Items []Target `json:"targets"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[Target]{Items: out.Items}
	return page, nil
}

// UpdateListener patches a listener (rotate cert, change default target group).
func (c *Client) UpdateListener(ctx context.Context, id string, listenerID string, body *UpdateListenerRequest, opts ...basaltic.RequestOption) (*Listener, error) {
	op := &basaltic.Operation{
		ID:       "updateListener",
		Method:   "PATCH",
		Path:     "/v1/load-balancers/{id}/listeners/{listener_id}",
		PathArgs: []string{id, listenerID},
		Body:     body,
	}
	var out struct {
		Listener *Listener `json:"listener"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Listener, nil
}

// UpdateLoadBalancer renames, scale, or resize a load balancer.
func (c *Client) UpdateLoadBalancer(ctx context.Context, id string, body *UpdateLoadBalancerRequest, opts ...basaltic.RequestOption) (*LoadBalancer, error) {
	op := &basaltic.Operation{
		ID:       "updateLoadBalancer",
		Method:   "PATCH",
		Path:     "/v1/load-balancers/{id}",
		PathArgs: []string{id},
		Body:     body,
	}
	var out struct {
		LoadBalancer *LoadBalancer `json:"load_balancer"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.LoadBalancer, nil
}

// UpdateRule updates a routing rule (full replace).
func (c *Client) UpdateRule(ctx context.Context, id string, listenerID string, ruleID string, body *UpdateRuleRequest, opts ...basaltic.RequestOption) (*Rule, error) {
	op := &basaltic.Operation{
		ID:       "updateRule",
		Method:   "PATCH",
		Path:     "/v1/load-balancers/{id}/listeners/{listener_id}/rules/{rule_id}",
		PathArgs: []string{id, listenerID, ruleID},
		Body:     body,
	}
	var out struct {
		Rule *Rule `json:"rule"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Rule, nil
}

// UpdateTargetGroup renames a target group or retune its health check, framing, or
// stickiness.
func (c *Client) UpdateTargetGroup(ctx context.Context, id string, body *UpdateTargetGroupRequest, opts ...basaltic.RequestOption) (*TargetGroup, error) {
	op := &basaltic.Operation{
		ID:       "updateTargetGroup",
		Method:   "PATCH",
		Path:     "/v1/target-groups/{id}",
		PathArgs: []string{id},
		Body:     body,
	}
	var out struct {
		TargetGroup *TargetGroup `json:"target_group"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.TargetGroup, nil
}
