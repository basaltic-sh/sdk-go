// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package registry

import (
	"time"
)

// BatchDeleteImageRequest at least one entry, at most 100 across both lists. Digests and tags
// may be mixed, which is what a cleanup script naturally produces.
type BatchDeleteImageRequest struct {
	// Digests removes the manifest AND every tag pointing at it.
	Digests []string `json:"digests,omitempty"`

	// Tags removes only the tag. The manifest survives, untagged.
	Tags []string `json:"tags,omitempty"`
}

type BatchDeleteResult struct {
	Deleted []*BatchDeleteResultDeletedItem `json:"deleted"`
	Failed  []*BatchDeleteResultFailedItem  `json:"failed"`
}

type BatchDeleteResultDeletedItem struct {
	Digest string `json:"digest"`

	// Tags the tags the delete took with it. Surfaced because deleting by
	// digest silently removes every tag pointing at it, and a caller who
	// meant to remove one needs to see the rest go.
	Tags []string `json:"tags,omitempty"`
}

type BatchDeleteResultFailedItem struct {
	Code   string `json:"code"`
	Digest string `json:"digest,omitempty"`
	Reason string `json:"reason"`
	Tag    string `json:"tag,omitempty"`
}

type CreatePullThroughCacheRuleRequest struct {
	// Prefix one or two lowercase repository-name components. It becomes the
	// leading path of every repository the rule creates, so it has to be a
	// name a client can pull.
	//
	// Required.
	Prefix string `json:"prefix"`

	// UpstreamPrefix optional path prepended upstream.
	UpstreamPrefix *string `json:"upstream_prefix,omitempty"`

	// UpstreamRegistry a hostname with an optional port. `https://` is accepted and
	// stripped; anything else is refused, because the connection can carry
	// a credential.
	//
	// Required.
	UpstreamRegistry string `json:"upstream_registry"`

	// UpstreamSecret the password or token. Write-only — sealed at rest and never
	// returned. Required with `upstream_username`, and refused without it:
	// a half-specified credential produces an anonymous pull that silently
	// draws on the small budget.
	UpstreamSecret *string `json:"upstream_secret,omitempty"`

	// UpstreamUsername Optional. Worth setting even against a public registry: an
	// authenticated Docker Hub pull draws on a far larger rate-limit
	// budget than an anonymous one.
	UpstreamUsername *string `json:"upstream_username,omitempty"`
}

type CreateRepositoryRequest struct {
	DeletionProtection *bool `json:"deletion_protection,omitempty"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	//
	// Required.
	Name string `json:"name"`

	// One of: "mutable", "immutable".
	TagMutability *string           `json:"tag_mutability,omitempty"`
	Tags          map[string]string `json:"tags,omitempty"`
}

type EncryptionConfiguration struct {
	// EncryptionType `platform_managed` wraps each blob's data key under a platform key
	// the customer does not manage. `kms` wraps it under a key in the
	// account's own KMS, so disabling that key stops the account's layers
	// from being readable.
	//
	// Reported as a named type rather than a boolean so a future third
	// posture has somewhere to go.
	//
	// One of: "platform_managed", "kms".
	EncryptionType string `json:"encryption_type"`

	// KMSKey customer-managed key reference. Omitted under platform encryption. A
	// deleted key retains its ID and CRN without a name; existing layers
	// still require that key.
	KMSKey *ResourceRef `json:"kms_key,omitempty"`

	// UpdatedAt when the setting last changed. Absent on an account that has never
	// configured encryption — there is no record, and reporting the
	// current time would read as a change that never happened.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type EncryptionConfigurationRequest struct {
	// EncryptionType Required. Inferring it from whether `kms_key` is present would make
	// "move me back to the platform key" indistinguishable from a client
	// that forgot the field.
	//
	// One of: "platform_managed", "kms".
	//
	// Required.
	EncryptionType string `json:"encryption_type"`

	// KMSKey KMS key UUID, CRN or account-scoped name. Required with `kms`, and
	// refused with `platform_managed` rather than ignored — the two
	// together are contradictory, and guessing which was meant could leave
	// content under a key the caller did not intend.
	KMSKey *string `json:"kms_key,omitempty"`
}

type GetAuthorizationTokenAuthorizationData struct {
	// AuthorizationToken `base64("<username>:<password>")` — the single opaque blob
	// ECR-shaped tooling expects.
	AuthorizationToken string    `json:"authorization_token"`
	Endpoint           string    `json:"endpoint"`
	ExpiresAt          time.Time `json:"expires_at"`

	// Password the credential itself, for `docker login --password-stdin`.
	Password string `json:"password"`

	// Username fixed, and ignored on the way in: the token identifies the
	// principal. Present so `docker login -u` has something to pass.
	Username string `json:"username"`
}

type GetLifecyclePolicyPolicy struct {
	// LastEvaluatedAt null for a policy the worker has not reached yet. Surfaced so a
	// customer whose images are not expiring can tell "never ran" from
	// "ran and matched nothing".
	LastEvaluatedAt *time.Time       `json:"last_evaluated_at,omitempty"`
	Rules           []*LifecycleRule `json:"rules"`
	UpdatedAt       time.Time        `json:"updated_at"`
}

type GetRepositoryPolicyPolicy struct {
	Document  *RepositoryPolicyDocument `json:"document"`
	UpdatedAt time.Time                 `json:"updated_at,omitempty"`
}

type Image struct {
	// ArtifactType what kind of thing this is. Distinguishes a container image from a
	// Helm chart, an SBOM, or a signature — all of which are stored as
	// manifests and would otherwise be indistinguishable.
	ArtifactType string `json:"artifact_type,omitempty"`

	// ConfigDigest absent on a multi-platform index, which has no config of its own.
	ConfigDigest string `json:"config_digest,omitempty"`

	// Digest the sha256 of the manifest document's exact bytes.
	Digest string `json:"digest"`

	// ID the image's row id — the pagination cursor for the list endpoint.
	ID string `json:"id"`

	// IsIndex true for a multi-platform image index or Docker manifest list.
	IsIndex bool `json:"is_index"`

	// LastPulledAt coarsened to a minute. A fleet of 500 nodes pulling one image in a
	// few seconds would otherwise be 500 row updates contending on one
	// page, and the lifecycle rule reading this measures in days.
	LastPulledAt *time.Time `json:"last_pulled_at,omitempty"`

	// ManifestSizeBytes the manifest document itself.
	ManifestSizeBytes int64 `json:"manifest_size_bytes,omitempty"`

	// MediaType The OCI or Docker v2 media type. Both vocabularies are stored as
	// pushed: containerd emits OCI, older `docker push` emits Docker v2,
	// and they are the same documents with different labels.
	MediaType string    `json:"media_type"`
	PushedAt  time.Time `json:"pushed_at"`

	// SizeBytes the sum of the distinct blobs this manifest references — what a
	// customer means by "how big is this image".
	SizeBytes int64 `json:"size_bytes"`

	// Subject the digest this manifest is *about*, set on a signature, SBOM, or
	// attestation. What the OCI referrers API resolves.
	Subject string `json:"subject,omitempty"`

	// Tags every tag pointing at this manifest. Empty for an untagged image.
	Tags []string `json:"tags"`
}

type ImageResponse struct {
	Image *Image `json:"image"`

	// Manifest the manifest document, parsed. NOT suitable for verifying the digest
	// — that covers the exact bytes pushed, and a round-trip through a
	// JSON encoder reorders keys. Fetch from `/v2` to re-hash.
	Manifest map[string]any `json:"manifest,omitempty"`
}

type LifecyclePolicyRequest struct {
	// Required.
	Rules []*LifecycleRule `json:"rules"`
}

type LifecycleRule struct {
	// Action the only action there is. A field rather than an implication so the
	// document has room for a future one without a policy written today
	// meaning something different tomorrow.
	//
	// One of: "expire".
	Action      string  `json:"action"`
	Description *string `json:"description,omitempty"`

	// Priority evaluation order, ascending, unique within the policy. An image is
	// claimed by the first matching rule.
	Priority  int                 `json:"priority"`
	Selection *LifecycleSelection `json:"selection"`
}

type LifecycleSelection struct {
	CountNumber int `json:"count_number"`

	// CountType `image_count_more_than` keeps the N newest matching images.
	// `since_image_pushed` expires those older than N days.
	// `since_image_pulled` expires those not pulled in N days, measuring a
	// never-pulled image from its push time — otherwise a never-pulled
	// image would be immortal, which is the opposite of what the rule is
	// for.
	//
	// One of: "image_count_more_than", "since_image_pushed", "since_image_pulled".
	CountType string `json:"count_type"`

	// CountUnit required by the two `since_*` types, forbidden by the other.
	//
	// One of: "days".
	CountUnit *string `json:"count_unit,omitempty"`

	// TagPatterns a glob whose only wildcard is `*`. A bare `*` is refused.
	TagPatterns []string `json:"tag_patterns,omitempty"`
	TagPrefixes []string `json:"tag_prefixes,omitempty"`

	// TagStatus `tagged` requires a prefix or pattern. `any` may appear on at most
	// one rule and must carry the highest priority number.
	//
	// One of: "tagged", "untagged", "any".
	TagStatus string `json:"tag_status"`
}

type PolicyCondition struct {
	Key      string `json:"key"`
	Operator string `json:"operator"`

	// One of: "for_all_values", "for_any_value".
	SetOperator *string  `json:"set_operator,omitempty"`
	Values      []string `json:"values"`
}

type PullThroughCacheRule struct {
	CreatedAt time.Time `json:"created_at"`
	CRN       string    `json:"crn"`

	// HasCredential whether a secret is stored. The secret itself is never returned;
	// rotation is a write of the new value.
	HasCredential bool `json:"has_credential"`

	// Prefix the leading path component(s) of the repository names this rule
	// claims.
	Prefix    string    `json:"prefix"`
	UpdatedAt time.Time `json:"updated_at"`

	// UpstreamPrefix prepended to the rest of the repository name upstream. `library` is
	// what makes `hub/nginx` resolve to Docker Hub's `library/nginx`.
	UpstreamPrefix string `json:"upstream_prefix,omitempty"`

	// UpstreamRegistry the upstream host. `docker.io` is stored as `registry-1.docker.io`,
	// which is the host that actually speaks `/v2` — the docker daemon
	// makes the same rewrite internally.
	UpstreamRegistry string `json:"upstream_registry"`
	UpstreamUsername string `json:"upstream_username,omitempty"`

	// URIPrefix the pullable reference prefix the rule creates. Returned rather than
	// assembled by the client, for the same reason a repository's `uri`
	// is: the account handle's position in the path is the only thing
	// separating two tenants' namespaces.
	URIPrefix string `json:"uri_prefix"`
}

type PutLifecyclePolicyPolicy struct {
	// LastEvaluatedAt null for a policy the worker has not reached yet. Surfaced so a
	// customer whose images are not expiring can tell "never ran" from
	// "ran and matched nothing".
	LastEvaluatedAt *time.Time       `json:"last_evaluated_at,omitempty"`
	Rules           []*LifecycleRule `json:"rules"`
	UpdatedAt       time.Time        `json:"updated_at"`
}

type Repository struct {
	CreatedAt time.Time `json:"created_at"`

	// CRN The IAM resource name. Built from the repository NAME, not its id,
	// so a policy naming a repository keeps applying after a delete and
	// re-create under the same name — which matters here more than
	// elsewhere, because the name is what every deployment manifest pulls.
	CRN string `json:"crn"`

	// DeletionProtection refuses `DeleteRepository` while set, even for an empty repository
	// and even with `force`.
	DeletionProtection bool   `json:"deletion_protection"`
	ID                 string `json:"id"`

	// Name lowercase alphanumeric path components separated by a single period,
	// one or two underscores, or hyphens, joined by slashes. This is the
	// OCI grammar — anything else is not pushable, so a laxer rule would
	// only produce repositories nobody can use. Resource names must not
	// start with the literal crn: prefix or be UUIDs (canonical, compact,
	// braced, or urn:uuid: forms, in either case).
	Name  string           `json:"name"`
	Stats *RepositoryStats `json:"stats,omitempty"`

	// TagMutability `immutable` refuses to move a tag that already resolves, so a
	// deployed digest cannot be swapped under a running fleet. Pushing a
	// tag that does not exist yet still works — the guarantee is that a
	// tag never changes meaning, not that the tag set is frozen. A repeat
	// push of the identical manifest is always allowed.
	//
	// One of: "mutable", "immutable".
	TagMutability string `json:"tag_mutability"`

	// Tags AWS-style resource tags. Feed the `basalt:ResourceTag/*` policy
	// conditions.
	Tags      map[string]string `json:"tags,omitempty"`
	UpdatedAt time.Time         `json:"updated_at"`

	// URI the pullable reference prefix. Returned rather than left for clients
	// to assemble: the account handle's position in the path is the only
	// thing separating two tenants' repositories of the same name.
	URI string `json:"uri"`
}

// RepositoryPolicyDocument a resource-based policy. Unlike an IAM policy, every statement must
// name a principal — the document is attached to a resource, so "who"
// cannot be implied by attachment.
type RepositoryPolicyDocument struct {
	Statements []*RepositoryPolicyDocumentStatement `json:"statements"`
	Version    string                               `json:"version"`
}

type RepositoryPolicyDocumentStatement struct {
	// Actions must all be `registry:` actions. A repository's owner cannot grant
	// `storage:GetObject` by writing it here, and ignoring such a
	// statement would leave a document that reads as a grant and is not
	// one.
	Actions    []string           `json:"actions"`
	Conditions []*PolicyCondition `json:"conditions,omitempty"`

	// One of: "allow", "deny".
	Effect string `json:"effect"`

	// Principals Account CRNs, an org CRN (`crn:iam:::org/<org-id>`), or `*` for the
	// public.
	Principals []string `json:"principals"`

	// Resources must be this repository's own CRN, or the bare `*` — the same
	// thing here, since the document is only ever evaluated against this
	// repository.
	Resources []string `json:"resources"`
	Sid       *string  `json:"sid,omitempty"`
}

type RepositoryPolicyRequest struct {
	// Required.
	Document *RepositoryPolicyDocument `json:"document"`
}

type RepositoryResponse struct {
	// PushCommands the login/build/push sequence for this repository. Rendered
	// server-side because the login username, the endpoint, and the
	// account handle's position in the reference all come from service
	// configuration.
	PushCommands []string    `json:"push_commands,omitempty"`
	Repository   *Repository `json:"repository"`
}

type RepositoryStats struct {
	ImageCount int `json:"image_count,omitempty"`

	// LastPushedAt null for a repository that has never received an image.
	LastPushedAt *time.Time `json:"last_pushed_at,omitempty"`

	// SizeBytes the sum of the DISTINCT blobs this repository's manifests reference.
	// Not the sum of its images' sizes: ten tags of one image share their
	// layers, and adding the per-image figures would report ten times the
	// space occupied.
	SizeBytes int64 `json:"size_bytes,omitempty"`
	TagCount  int   `json:"tag_count,omitempty"`
}

type ResourceRef struct {
	CRN string `json:"crn"`
	ID  string `json:"id"`

	// Name display label at response time — the resource's name where it has
	// one, otherwise its human label (an invoice number, an account
	// display name). Absent when the target is gone. Not a lookup key.
	Name string `json:"name,omitempty"`
}

type SetRepositoryPolicyPolicy struct {
	Document  *RepositoryPolicyDocument `json:"document"`
	UpdatedAt time.Time                 `json:"updated_at,omitempty"`
}

type StartLifecyclePolicyPreviewPreview struct {
	Matches []*StartLifecyclePolicyPreviewPreviewMatche `json:"matches"`

	// SizeBytes what the expiry would reclaim: only layers no surviving image still
	// references, each counted once.
	SizeBytes int64 `json:"size_bytes"`

	// TotalImages how many images the repository holds, so 400 matches can be read
	// against the 402 that exist.
	TotalImages int `json:"total_images"`
}

type StartLifecyclePolicyPreviewPreviewMatche struct {
	Digest       string     `json:"digest"`
	LastPulledAt *time.Time `json:"last_pulled_at,omitempty"`
	PushedAt     time.Time  `json:"pushed_at"`

	// RulePriority which rule claimed this image.
	RulePriority int      `json:"rule_priority"`
	Tags         []string `json:"tags,omitempty"`
}

type Tag struct {
	CreatedAt time.Time `json:"created_at"`
	Digest    string    `json:"digest"`
	Name      string    `json:"name"`

	// UpdatedAt when the tag last moved. Equal to created_at for a tag that never
	// has.
	UpdatedAt time.Time `json:"updated_at"`
}

// UpdatePullThroughCacheRuleRequest only the fields present are changed. Absent and empty differ for the
// credential: absent leaves it alone, empty clears it.
type UpdatePullThroughCacheRuleRequest struct {
	UpstreamPrefix   *string `json:"upstream_prefix,omitempty"`
	UpstreamRegistry *string `json:"upstream_registry,omitempty"`
	UpstreamSecret   *string `json:"upstream_secret,omitempty"`
	UpstreamUsername *string `json:"upstream_username,omitempty"`
}

// UpdateRepositoryRequest omitted fields are left alone. `tag_mutability` and
// `deletion_protection` are therefore distinguishable from `false` only
// by absence, which is why a request that changes nothing but `tags` is
// authorized as `registry:TagResource`.
type UpdateRepositoryRequest struct {
	DeletionProtection *bool `json:"deletion_protection,omitempty"`

	// One of: "mutable", "immutable".
	TagMutability *string           `json:"tag_mutability,omitempty"`
	Tags          map[string]string `json:"tags,omitempty"`
}
