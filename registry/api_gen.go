// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package registry

import (
	"context"
	"iter"
	"net/url"
	"strconv"

	basaltic "github.com/basaltic-sh/sdk-go"
)

// DeleteRepositoryParams are the optional filters and pagination controls for
// [Client.DeleteRepository]. A nil *DeleteRepositoryParams sends none of them.
type DeleteRepositoryParams struct {
	// Force delete the repository even though it still contains images.
	Force *bool
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *DeleteRepositoryParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.Force != nil {
		q.Set("force", strconv.FormatBool(*p.Force))
	}
	return q
}

// DescribeImagesParams are the optional filters and pagination controls for
// [Client.DescribeImages]. A nil *DescribeImagesParams sends none of them.
type DescribeImagesParams struct {
	// Limit maximum items to return.
	Limit int

	// Marker id of the last item on the previous page.
	Marker string

	// TagPrefix return only images carrying a tag with this prefix.
	TagPrefix string

	// TagStatus restrict to one population. `untagged` finds the residue of moved
	// tags, which is usually what fills a repository up.
	//
	//
	// One of: "tagged", "untagged", "any".
	TagStatus string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *DescribeImagesParams) query() url.Values {
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
	if p.TagPrefix != "" {
		q.Set("tag_prefix", p.TagPrefix)
	}
	if p.TagStatus != "" {
		q.Set("tag_status", p.TagStatus)
	}
	return q
}

// withMarker copies p with the pagination cursor replaced, leaving the
// caller's value untouched across pages.
func (p *DescribeImagesParams) withMarker(marker string) *DescribeImagesParams {
	var out DescribeImagesParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// DescribePullThroughCacheRulesParams are the optional filters and pagination controls for
// [Client.DescribePullThroughCacheRules]. A nil *DescribePullThroughCacheRulesParams sends none of them.
type DescribePullThroughCacheRulesParams struct {
	// CRN exact pull-through-cache-rule CRN, intersected with `name`.
	// Malformed CRN syntax returns 400 INVALID_INPUT. A valid CRN for a
	// different service, type, region or account returns an empty list.
	// Slash-bearing prefixes remain part of the opaque rule name.
	CRN string

	// Name exact, account-scoped immutable repository name or cache-rule
	// prefix. Slash-bearing names are accepted. Intersects with `crn`
	// before pagination.
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *DescribePullThroughCacheRulesParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.CRN != "" {
		q.Set("crn", p.CRN)
	}
	if p.Name != "" {
		q.Set("name", p.Name)
	}
	return q
}

// DescribeRepositoriesParams are the optional filters and pagination controls for
// [Client.DescribeRepositories]. A nil *DescribeRepositoriesParams sends none of them.
type DescribeRepositoriesParams struct {
	// CRN exact repository CRN, intersected with `name` before pagination.
	// Malformed CRN syntax returns 400 INVALID_INPUT. A valid CRN for a
	// different service, type, region or account returns an empty page.
	CRN string

	// Limit maximum items to return.
	Limit int

	// Marker id of the last item on the previous page.
	Marker string

	// Name exact, account-scoped immutable repository name or cache-rule
	// prefix. Slash-bearing names are accepted. Intersects with `crn`
	// before pagination.
	Name string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *DescribeRepositoriesParams) query() url.Values {
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
func (p *DescribeRepositoriesParams) withMarker(marker string) *DescribeRepositoriesParams {
	var out DescribeRepositoriesParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// ListTagsParams are the optional filters and pagination controls for
// [Client.ListTags]. A nil *ListTagsParams sends none of them.
type ListTagsParams struct {
	// Limit maximum items to return.
	Limit int

	// Marker name of the last tag on the previous page. Tags are ordered
	// lexically, so the cursor is a name rather than an id.
	Marker string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListTagsParams) query() url.Values {
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
func (p *ListTagsParams) withMarker(marker string) *ListTagsParams {
	var out ListTagsParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// BatchDeleteImage deletes images.
//
// Deletes by digest and/or by tag, and the two are not the same
// operation.
//
// A **digest** removes the manifest and every tag pointing at it. The
// response reports which tags went with it, because a caller who meant
// to remove one tag needs to see the rest go.
//
// A **tag** removes only the pointer. The manifest survives, untagged,
// because another tag may still name it.
//
// Partial success is reported rather than failing the call. A cleanup
// script naturally names digests that are already gone, and turning that
// into a whole-batch failure would make the script's second run fail on
// the work its first run completed.
func (c *Client) BatchDeleteImage(ctx context.Context, repositoryID string, body *BatchDeleteImageRequest, opts ...basaltic.RequestOption) (*BatchDeleteResult, error) {
	op := &basaltic.Operation{
		ID:       "batchDeleteImage",
		Method:   "POST",
		Path:     "/v1/repositories/{repository_id}/images/batch-delete",
		PathArgs: []string{repositoryID},
		Body:     body,
	}
	var out BatchDeleteResult
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// CreatePullThroughCacheRule creates a pull-through cache rule.
//
// Maps a prefix in this account onto an upstream registry, so a fleet
// pulls public base images through its own registry.
//
// Set `prefix` to `hub`, `upstream_registry` to `docker.io`, and
// `upstream_prefix` to `library` to pull Docker Hub official images:
//
// ```bash
// docker pull cr.sa-saopaulo-1.basaltic.cloud/my-account/hub/nginx:1.27
// ```
//
// **What this is for.** Docker Hub rate-limits by source IP, and a fleet
// behind one NAT is one IP: a rolling deploy across two hundred nodes is
// two hundred pulls of the same base image against a shared budget, and
// the failure — every node's image pull backing off at once —
// arrives as an outage in the middle of a deploy. One fetch, cached,
// takes the budget off the critical path. Setting `upstream_username`
// and `upstream_secret` helps even for a public upstream, because an
// authenticated Docker Hub pull draws on a much larger budget than an
// anonymous one.
//
// **Nothing is mirrored now.** A rule is a mapping, not an instruction
// to copy a registry. The repository under it and its content are
// created by the first pull that misses, so a rule nobody pulls through
// costs a row.
//
// **Who can trigger a fetch.** A cache miss spends the account's storage
// quota and pulls bytes from a third party, so it needs
// `registry:BatchImportUpstreamImage` on top of the pull actions. A
// fleet given only pull consumes what is already cached and gets a `404`
// on anything else — which is how you keep a production cluster from
// reaching the internet while CI keeps the mirror warm.
//
// **The prefix must own its namespace.** Creating a rule over a prefix
// the account already has repositories under is refused: those
// repositories would stay un-mirrored while new names under the same
// prefix mirrored, and nothing in a reference would say which applied.
//
// Requires `registry:CreatePullThroughCacheRule`.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreatePullThroughCacheRule(ctx context.Context, body *CreatePullThroughCacheRuleRequest, opts ...basaltic.RequestOption) (*PullThroughCacheRule, error) {
	op := &basaltic.Operation{
		ID:     "createPullThroughCacheRule",
		Method: "POST",
		Path:   "/v1/pull-through-cache-rules",
		Body:   body,
	}
	var out struct {
		PullThroughCacheRule *PullThroughCacheRule `json:"pull_through_cache_rule"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.PullThroughCacheRule, nil
}

// CreateRepository creates a repository.
//
// A repository must exist before anything can be pushed to it. There is
// no create-on-push: a repository carries quota, resource tags, a
// tag-mutability setting, and possibly a resource policy, and conjuring
// one from a typo in a `docker push` would produce an unlabelled,
// unpoliced repository nobody meant to make.
//
// The name is immutable. It is the CRN every IAM policy names and the
// reference every deployment manifest pulls, so renaming would silently
// invalidate both.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateRepository(ctx context.Context, body *CreateRepositoryRequest, opts ...basaltic.RequestOption) (*RepositoryResponse, error) {
	op := &basaltic.Operation{
		ID:     "createRepository",
		Method: "POST",
		Path:   "/v1/repositories",
		Body:   body,
	}
	var out RepositoryResponse
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// DeleteEncryptionConfiguration deletes encryption configuration.
//
// Restore the implicit platform key for future writes. Requires
// registry:PutEncryptionConfiguration. Existing blobs retain their
// encryption and ordinary garbage collection.
func (c *Client) DeleteEncryptionConfiguration(ctx context.Context, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:     "deleteEncryptionConfiguration",
		Method: "DELETE",
		Path:   "/v1/encryption-configuration",
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteLifecyclePolicy deletes the lifecycle policy.
//
// Stops automatic expiry. Images already expired are not restored —
// there is nothing to restore them from.
func (c *Client) DeleteLifecyclePolicy(ctx context.Context, repositoryID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteLifecyclePolicy",
		Method:   "DELETE",
		Path:     "/v1/repositories/{repository_id}/lifecycle-policy",
		PathArgs: []string{repositoryID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeletePullThroughCacheRule deletes a pull-through cache rule.
//
// Removes the mapping. The repositories it created and everything
// mirrored into them stay exactly where they are: they hold real
// pullable images a fleet may be depending on right now, and deleting a
// mapping is a statement about where NEW content comes from. They keep
// serving what they have and stop refreshing — which also makes this
// the way to freeze a mirror.
//
// Requires `registry:DeletePullThroughCacheRule`.
func (c *Client) DeletePullThroughCacheRule(ctx context.Context, prefix string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deletePullThroughCacheRule",
		Method:   "DELETE",
		Path:     "/v1/pull-through-cache-rules/{prefix}",
		PathArgs: []string{prefix},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteRepository deletes a repository.
//
// Removes the repository and every image in it.
//
// `force` is required when the repository still holds images. It does
// NOT override `deletion_protection` — that has to be cleared first,
// which is the whole point of a flag someone deliberately set.
//
// The layers those images referenced are not freed immediately. Whether
// a layer is still needed is a question about every other repository in
// the account, and it is answered by the background sweep.
func (c *Client) DeleteRepository(ctx context.Context, repositoryID string, params *DeleteRepositoryParams, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteRepository",
		Method:   "DELETE",
		Path:     "/v1/repositories/{repository_id}",
		PathArgs: []string{repositoryID},
	}
	op.Query = params.query()
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DeleteRepositoryPolicy deletes the repository policy.
//
// Closes the repository to everyone outside the owning account.
func (c *Client) DeleteRepositoryPolicy(ctx context.Context, repositoryID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteRepositoryPolicy",
		Method:   "DELETE",
		Path:     "/v1/repositories/{repository_id}/policy",
		PathArgs: []string{repositoryID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// DescribeImage describes an image.
//
// Resolves a tag or a digest to one image, including its raw manifest
// document.
//
// The `manifest` field is the parsed JSON, for convenience. It is NOT
// suitable for verifying the digest: the digest covers the exact bytes
// the client pushed, and a round-trip through a JSON encoder reorders
// keys. Fetch the manifest from `/v2` if you need to re-hash it.
func (c *Client) DescribeImage(ctx context.Context, repositoryID string, imageReference string, opts ...basaltic.RequestOption) (*ImageResponse, error) {
	op := &basaltic.Operation{
		ID:       "describeImage",
		Method:   "GET",
		Path:     "/v1/repositories/{repository_id}/images/{image_reference}",
		PathArgs: []string{repositoryID, imageReference},
	}
	var out ImageResponse
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// DescribeImages lists images.
//
// Every manifest in the repository with the tags pointing at it.
//
// Filter with `tag_status=untagged` to find the residue of moved tags
// — the images that accumulate when `:latest` is pushed repeatedly,
// and usually the answer to "why is this repository so large".
//
// Returns one page. Use DescribeImagesAll to walk every page.
func (c *Client) DescribeImages(ctx context.Context, repositoryID string, params *DescribeImagesParams, opts ...basaltic.RequestOption) (*basaltic.Page[Image], error) {
	op := &basaltic.Operation{
		ID:       "describeImages",
		Method:   "GET",
		Path:     "/v1/repositories/{repository_id}/images",
		PathArgs: []string{repositoryID},
	}
	op.Query = params.query()
	var out struct {
		Items []Image `json:"images"`
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
	page := &basaltic.Page[Image]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// DescribeImagesAll walks every page of DescribeImages, yielding one
// item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.DescribeImagesAll(ctx, repositoryID, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) DescribeImagesAll(ctx context.Context, repositoryID string, params *DescribeImagesParams, opts ...basaltic.RequestOption) iter.Seq2[Image, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Image], error) {
		return c.DescribeImages(ctx, repositoryID, params.withMarker(marker), opts...)
	})
}

// DescribePullThroughCacheRule gets a pull-through cache rule.
//
// Requires `registry:DescribePullThroughCacheRules`.
func (c *Client) DescribePullThroughCacheRule(ctx context.Context, prefix string, opts ...basaltic.RequestOption) (*PullThroughCacheRule, error) {
	op := &basaltic.Operation{
		ID:       "describePullThroughCacheRule",
		Method:   "GET",
		Path:     "/v1/pull-through-cache-rules/{prefix}",
		PathArgs: []string{prefix},
	}
	var out struct {
		PullThroughCacheRule *PullThroughCacheRule `json:"pull_through_cache_rule"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.PullThroughCacheRule, nil
}

// DescribePullThroughCacheRules lists pull-through cache rules.
//
// Requires `registry:DescribePullThroughCacheRules`.
//
// The upstream credential is never returned. A rule reports only
// `has_credential`, so rotation is a write of the new value rather than
// a read-modify-write.
func (c *Client) DescribePullThroughCacheRules(ctx context.Context, params *DescribePullThroughCacheRulesParams, opts ...basaltic.RequestOption) (*basaltic.Page[PullThroughCacheRule], error) {
	op := &basaltic.Operation{
		ID:     "describePullThroughCacheRules",
		Method: "GET",
		Path:   "/v1/pull-through-cache-rules",
	}
	op.Query = params.query()
	var out struct {
		Items []PullThroughCacheRule `json:"pull_through_cache_rules"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	page := &basaltic.Page[PullThroughCacheRule]{Items: out.Items}
	return page, nil
}

// DescribeRepositories lists repositories.
//
// Returns one page. Use DescribeRepositoriesAll to walk every page.
func (c *Client) DescribeRepositories(ctx context.Context, params *DescribeRepositoriesParams, opts ...basaltic.RequestOption) (*basaltic.Page[Repository], error) {
	op := &basaltic.Operation{
		ID:     "describeRepositories",
		Method: "GET",
		Path:   "/v1/repositories",
	}
	op.Query = params.query()
	var out struct {
		Items []Repository `json:"repositories"`
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
	page := &basaltic.Page[Repository]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// DescribeRepositoriesAll walks every page of DescribeRepositories,
// yielding one item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.DescribeRepositoriesAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) DescribeRepositoriesAll(ctx context.Context, params *DescribeRepositoriesParams, opts ...basaltic.RequestOption) iter.Seq2[Repository, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Repository], error) {
		return c.DescribeRepositories(ctx, params.withMarker(marker), opts...)
	})
}

// DescribeRepository describes a repository.
func (c *Client) DescribeRepository(ctx context.Context, repositoryID string, opts ...basaltic.RequestOption) (*RepositoryResponse, error) {
	op := &basaltic.Operation{
		ID:       "describeRepository",
		Method:   "GET",
		Path:     "/v1/repositories/{repository_id}",
		PathArgs: []string{repositoryID},
	}
	var out RepositoryResponse
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetAuthorizationToken gets a docker login credential.
//
// Mints the short-lived password `docker login` consumes.
//
// A POST rather than a GET because the response is a bearer secret: a
// GET would land in browser history, proxy access logs, and any
// intermediary that records URLs.
//
// The credential identifies the principal and nothing more. It names no
// repository and grants no access — the docker daemon exchanges it at
// `/v2/token` for per-repository bearer tokens, and every authorization
// decision is made there, against the policies as they stand at that
// moment.
//
// Requires `registry:GetAuthorizationToken`. That is a separate
// permission from any repository action: a principal may be allowed to
// pull through this API without being allowed to hand a docker daemon a
// credential.
func (c *Client) GetAuthorizationToken(ctx context.Context, opts ...basaltic.RequestOption) (*GetAuthorizationTokenAuthorizationData, error) {
	op := &basaltic.Operation{
		ID:     "getAuthorizationToken",
		Method: "POST",
		Path:   "/v1/authorization-token",
	}
	var out struct {
		AuthorizationData *GetAuthorizationTokenAuthorizationData `json:"authorization_data"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.AuthorizationData, nil
}

// GetEncryptionConfiguration gets the encryption configuration.
//
// Reports which key seals the account's layer bytes.
//
// Account-wide, not per-repository. Layer bytes are content-addressed
// and deduplicated across every repository in the account, so a layer
// two repositories share has exactly one copy — and therefore exactly
// one key. A per-repository setting would have no answer for which key
// that is, and both resolutions are bad: fork the blob namespace and
// lose deduplication, or seal under a key the customer did not choose.
//
// Requires `registry:GetEncryptionConfiguration`.
func (c *Client) GetEncryptionConfiguration(ctx context.Context, opts ...basaltic.RequestOption) (*EncryptionConfiguration, error) {
	op := &basaltic.Operation{
		ID:     "getEncryptionConfiguration",
		Method: "GET",
		Path:   "/v1/encryption-configuration",
	}
	var out struct {
		EncryptionConfiguration *EncryptionConfiguration `json:"encryption_configuration"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.EncryptionConfiguration, nil
}

// GetLifecyclePolicy gets the lifecycle policy.
func (c *Client) GetLifecyclePolicy(ctx context.Context, repositoryID string, opts ...basaltic.RequestOption) (*GetLifecyclePolicyPolicy, error) {
	op := &basaltic.Operation{
		ID:       "getLifecyclePolicy",
		Method:   "GET",
		Path:     "/v1/repositories/{repository_id}/lifecycle-policy",
		PathArgs: []string{repositoryID},
	}
	var out struct {
		Policy *GetLifecyclePolicyPolicy `json:"policy"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Policy, nil
}

// GetRepositoryPolicy gets the repository policy.
func (c *Client) GetRepositoryPolicy(ctx context.Context, repositoryID string, opts ...basaltic.RequestOption) (*GetRepositoryPolicyPolicy, error) {
	op := &basaltic.Operation{
		ID:       "getRepositoryPolicy",
		Method:   "GET",
		Path:     "/v1/repositories/{repository_id}/policy",
		PathArgs: []string{repositoryID},
	}
	var out struct {
		Policy *GetRepositoryPolicyPolicy `json:"policy"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Policy, nil
}

// GetRetainedRegistryPresence — Check retained registry data presence.
//
// Requires registry:GetEncryptionConfiguration. Reports unfinished
// uploads and blobs awaiting ordinary garbage collection in the selected
// account, without returning content or bypassing retention.
func (c *Client) GetRetainedRegistryPresence(ctx context.Context, opts ...basaltic.RequestOption) (bool, error) {
	op := &basaltic.Operation{
		ID:     "getRetainedRegistryPresence",
		Method: "GET",
		Path:   "/v1/encryption-configuration/retained-data",
	}
	var out struct {
		HasResources bool `json:"has_resources"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return false, err
	}
	return out.HasResources, nil
}

// ListTags lists tags.
//
// Ordered lexically, matching `GET /v2/{name}/tags/list`. The pagination
// marker is therefore a tag NAME, not a uuid.
//
// Returns one page. Use ListTagsAll to walk every page.
func (c *Client) ListTags(ctx context.Context, repositoryID string, params *ListTagsParams, opts ...basaltic.RequestOption) (*basaltic.Page[Tag], error) {
	op := &basaltic.Operation{
		ID:       "listTags",
		Method:   "GET",
		Path:     "/v1/repositories/{repository_id}/tags",
		PathArgs: []string{repositoryID},
	}
	op.Query = params.query()
	var out struct {
		Items []Tag `json:"tags"`
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
	page := &basaltic.Page[Tag]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListTagsAll walks every page of ListTags, yielding one item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListTagsAll(ctx, repositoryID, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListTagsAll(ctx context.Context, repositoryID string, params *ListTagsParams, opts ...basaltic.RequestOption) iter.Seq2[Tag, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Tag], error) {
		return c.ListTags(ctx, repositoryID, params.withMarker(marker), opts...)
	})
}

// PutEncryptionConfiguration sets the encryption configuration.
//
// Chooses the key that wraps the data key of every layer pushed AFTER
// this call.
//
// Existing layers are not re-encrypted and not re-wrapped. Each records
// the key it was sealed under and stays readable through it — so a
// customer who rotates away from a key they believe compromised should
// expect the old layers to become unreadable when that key is destroyed.
// That is the point of rotating, and re-wrapping an account's entire
// layer set is not something a synchronous API call can honestly
// promise.
//
// The key is proved usable here rather than on the first push. A
// disabled key, a destroyed key, or a key in another account is rejected
// with `400`, because a registry whose pushes fail with a KMS error is
// indistinguishable from one that is down.
//
// Deduplication is unaffected either way: a blob's digest is computed
// over its plaintext as the bytes stream in, before anything is sealed.
//
// Requires `registry:PutEncryptionConfiguration`.
func (c *Client) PutEncryptionConfiguration(ctx context.Context, body *EncryptionConfigurationRequest, opts ...basaltic.RequestOption) (*EncryptionConfiguration, error) {
	op := &basaltic.Operation{
		ID:     "putEncryptionConfiguration",
		Method: "PUT",
		Path:   "/v1/encryption-configuration",
		Body:   body,
	}
	var out struct {
		EncryptionConfiguration *EncryptionConfiguration `json:"encryption_configuration"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.EncryptionConfiguration, nil
}

// PutLifecyclePolicy sets the lifecycle policy.
//
// Rules are applied in `priority` order and an image is claimed by the
// FIRST rule that matches it. That is what makes "keep the last 10
// `prod-*` images (priority 1), expire everything older than 90 days
// (priority 100)" mean what it reads as, rather than having the second
// rule also delete the images the first one kept.
//
// Three constraints exist to stop a policy from quietly meaning
// something other than it appears to:
//
//   - At most one rule may select `tag_status: any`, and it must carry
//     the highest priority number. A catch-all at priority 1 would
//     claim every image before any specific rule ran, leaving the
//     careful rules below it silently dead.
//   - A `tagged` rule must list a prefix or a pattern. Without one it
//     is `any` in disguise, and would escape the constraint above.
//   - A tag pattern of exactly `*` is refused for the same reason.
//
// Installing a policy resets `last_evaluated_at`: rules that have just
// been rewritten have never run in their current form.
func (c *Client) PutLifecyclePolicy(ctx context.Context, repositoryID string, body *LifecyclePolicyRequest, opts ...basaltic.RequestOption) (*PutLifecyclePolicyPolicy, error) {
	op := &basaltic.Operation{
		ID:       "putLifecyclePolicy",
		Method:   "PUT",
		Path:     "/v1/repositories/{repository_id}/lifecycle-policy",
		PathArgs: []string{repositoryID},
		Body:     body,
	}
	var out struct {
		Policy *PutLifecyclePolicyPolicy `json:"policy"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Policy, nil
}

// SetRepositoryPolicy sets the repository policy.
//
// Grants access to principals outside the owning account. With principal
// `*` this makes the repository publicly pullable, which is how a
// base-image repository is shared.
//
// Replaced wholesale, never merged: a caller who removes a statement
// from a policy they fetched expects the grant gone.
//
// Every statement must name a principal, must use only `registry:`
// actions, and must name this repository's own CRN (or `*`, which means
// the same thing here). `not_actions` and `not_resources` are refused
// — a negated set cannot be bounded to one repository by inspection.
//
// Cross-account access requires the caller's own IAM policy to allow the
// action as well. A resource policy widens what this repository permits;
// it cannot exceed what the caller's organization permits.
func (c *Client) SetRepositoryPolicy(ctx context.Context, repositoryID string, body *RepositoryPolicyRequest, opts ...basaltic.RequestOption) (*SetRepositoryPolicyPolicy, error) {
	op := &basaltic.Operation{
		ID:       "setRepositoryPolicy",
		Method:   "PUT",
		Path:     "/v1/repositories/{repository_id}/policy",
		PathArgs: []string{repositoryID},
		Body:     body,
	}
	var out struct {
		Policy *SetRepositoryPolicyPolicy `json:"policy"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Policy, nil
}

// StartLifecyclePolicyPreview — Preview a lifecycle policy.
//
// Reports what a policy would expire, without expiring anything.
//
// A POST despite being read-only, because it takes a candidate rule set
// in the body — this is how you see what a policy would delete before
// installing it. Send an empty body to preview the installed policy,
// answering "what happens on the next pass".
//
// `size_bytes` counts only layers that no SURVIVING image in this
// repository still references, and counts each once. Summing the matched
// images' sizes would double-count every shared base layer and report
// several times the truth.
func (c *Client) StartLifecyclePolicyPreview(ctx context.Context, repositoryID string, body *LifecyclePolicyRequest, opts ...basaltic.RequestOption) (*StartLifecyclePolicyPreviewPreview, error) {
	op := &basaltic.Operation{
		ID:       "startLifecyclePolicyPreview",
		Method:   "POST",
		Path:     "/v1/repositories/{repository_id}/lifecycle-policy/preview",
		PathArgs: []string{repositoryID},
		Body:     body,
	}
	var out struct {
		Preview *StartLifecyclePolicyPreviewPreview `json:"preview"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Preview, nil
}

// UpdatePullThroughCacheRule updates a pull-through cache rule.
//
// Repoints a rule or rotates its credential.
//
// The prefix itself is immutable: it is the rule's identity, the CRN a
// policy names, and the leading path of every repository already
// mirrored through it. Moving it would orphan those repositories from
// the rule that created them.
//
// Omitting a field leaves it alone; sending an empty `upstream_username`
// and `upstream_secret` together clears the credential and makes the
// rule anonymous. Content already mirrored is not re-fetched when the
// upstream changes — it is content-addressed and still valid.
//
// Requires `registry:UpdatePullThroughCacheRule`.
func (c *Client) UpdatePullThroughCacheRule(ctx context.Context, prefix string, body *UpdatePullThroughCacheRuleRequest, opts ...basaltic.RequestOption) (*PullThroughCacheRule, error) {
	op := &basaltic.Operation{
		ID:       "updatePullThroughCacheRule",
		Method:   "PATCH",
		Path:     "/v1/pull-through-cache-rules/{prefix}",
		PathArgs: []string{prefix},
		Body:     body,
	}
	var out struct {
		PullThroughCacheRule *PullThroughCacheRule `json:"pull_through_cache_rule"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.PullThroughCacheRule, nil
}

// UpdateRepository updates repository settings.
//
// The name is immutable; everything else here is not.
//
// `tag_mutability` may be changed in either direction. Turning
// immutability on does not rewrite history — tags already moved stay
// where they are — it only constrains what happens next.
//
// A request that changes ONLY `tags` is authorized as
// `registry:TagResource` rather than `registry:UpdateRepository`, so a
// principal allowed to label resources for cost allocation does not
// thereby become able to turn off tag immutability.
func (c *Client) UpdateRepository(ctx context.Context, repositoryID string, body *UpdateRepositoryRequest, opts ...basaltic.RequestOption) (*RepositoryResponse, error) {
	op := &basaltic.Operation{
		ID:       "updateRepository",
		Method:   "PATCH",
		Path:     "/v1/repositories/{repository_id}",
		PathArgs: []string{repositoryID},
		Body:     body,
	}
	var out RepositoryResponse
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// DescribePullThroughCacheRuleByReference fetches one pull through cache
// rule by an id, a CRN or a name.
//
// The reference is classified by its syntax alone, exactly as the
// platform does (see [basaltic.ParseReference]): an id is fetched with
// [Client.DescribePullThroughCacheRule]; a CRN or a name goes to
// [Client.DescribePullThroughCacheRules] as an exact filter, together
// with any filters already set on scope, which may be nil. A miss is a
// not-found error for the kind the string was read as — no other kind
// is tried — and more than one match is a
// [basaltic.AmbiguousReferenceError].
func (c *Client) DescribePullThroughCacheRuleByReference(ctx context.Context, ref string, scope *DescribePullThroughCacheRulesParams, opts ...basaltic.RequestOption) (*PullThroughCacheRule, error) {
	return basaltic.ResolveByReference(ctx, ref, "pull-through-cache-rule", "describePullThroughCacheRules", true,
		func(ctx context.Context, refID string) (*PullThroughCacheRule, error) {
			return c.DescribePullThroughCacheRule(ctx, refID, opts...)
		},
		func(ctx context.Context, refName, refCRN string) (*basaltic.Page[PullThroughCacheRule], error) {
			var p DescribePullThroughCacheRulesParams
			if scope != nil {
				p = *scope
			}
			p.Name = refName
			p.CRN = refCRN
			return c.DescribePullThroughCacheRules(ctx, &p, opts...)
		})
}
