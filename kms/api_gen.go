// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package kms

import (
	"context"
	"iter"
	"net/url"
	"strconv"

	basaltic "github.com/basaltic-sh/sdk-go"
)

// ListKeysParams are the optional filters and pagination controls for
// [Client.ListKeys]. A nil *ListKeysParams sends none of them.
type ListKeysParams struct {
	Limit int

	// Marker resume token — the last key id from the previous page.
	Marker string

	// Name optional substring filter on key name.
	Name string

	// State optional state filter (enabled / disabled / pending_deletion).
	State KeyState
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListKeysParams) query() url.Values {
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
	if p.State != "" {
		q.Set("state", string(p.State))
	}
	return q
}

// withMarker copies p with the pagination cursor replaced, leaving the
// caller's value untouched across pages.
func (p *ListKeysParams) withMarker(marker string) *ListKeysParams {
	var out ListKeysParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// CancelKeyDeletion cancels a scheduled deletion.
//
// Exit the pending-deletion window. The key returns to state=disabled
// (NOT enabled — matches AWS KMS; caller must explicitly enable).
// Re-reserves quota; fails with 403 if the account is now over its
// kms_keys quota.
func (c *Client) CancelKeyDeletion(ctx context.Context, keyID string, opts ...basaltic.RequestOption) (*Key, error) {
	op := &basaltic.Operation{
		ID:       "cancelKeyDeletion",
		Method:   "POST",
		Path:     "/v1/keys/{key_id}/cancel-deletion",
		PathArgs: []string{keyID},
	}
	var out struct {
		Key *Key `json:"key"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Key, nil
}

// CreateKey creates a KMS key.
//
// Provision a fresh key under the requested spec and usage. Synchronous:
// the response carries the key in state=enabled, ready for crypto
// operations.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateKey(ctx context.Context, body *CreateKeyRequest, opts ...basaltic.RequestOption) (*Key, error) {
	op := &basaltic.Operation{
		ID:     "createKey",
		Method: "POST",
		Path:   "/v1/keys",
		Body:   body,
	}
	var out struct {
		Key *Key `json:"key"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Key, nil
}

// Decrypt decrypts a ciphertext.
func (c *Client) Decrypt(ctx context.Context, keyID string, body *DecryptRequest, opts ...basaltic.RequestOption) ([]byte, error) {
	op := &basaltic.Operation{
		ID:       "decrypt",
		Method:   "POST",
		Path:     "/v1/keys/{key_id}/decrypt",
		PathArgs: []string{keyID},
		Body:     body,
	}
	var out struct {
		Plaintext []byte `json:"plaintext"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Plaintext, nil
}

// DisableKey disables a key.
//
// Park the key in state=disabled. Crypto operations refuse the key in
// this state; the key material stays so the key can be re-enabled at any
// time. Refused while the key is pending deletion.
func (c *Client) DisableKey(ctx context.Context, keyID string, opts ...basaltic.RequestOption) (*Key, error) {
	op := &basaltic.Operation{
		ID:       "disableKey",
		Method:   "POST",
		Path:     "/v1/keys/{key_id}/disable",
		PathArgs: []string{keyID},
	}
	var out struct {
		Key *Key `json:"key"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Key, nil
}

// EnableKey enables a disabled key.
//
// Flip state=disabled back to state=enabled. Refused while the key is
// pending deletion; cancel deletion first.
func (c *Client) EnableKey(ctx context.Context, keyID string, opts ...basaltic.RequestOption) (*Key, error) {
	op := &basaltic.Operation{
		ID:       "enableKey",
		Method:   "POST",
		Path:     "/v1/keys/{key_id}/enable",
		PathArgs: []string{keyID},
	}
	var out struct {
		Key *Key `json:"key"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Key, nil
}

// Encrypt encrypts a payload.
//
// Encrypt plaintext under the key. Symmetric (AES-256) uses AES-GCM with
// a random 96-bit nonce prefixed to the ciphertext; aad is bound as AEAD
// additional data and must be supplied verbatim at decrypt. RSA keys use
// OAEP-SHA256 (aad is ignored). Refused unless key_usage=encrypt_decrypt
// and state=enabled.
func (c *Client) Encrypt(ctx context.Context, keyID string, body *EncryptRequest, opts ...basaltic.RequestOption) (*EncryptResponse, error) {
	op := &basaltic.Operation{
		ID:       "encrypt",
		Method:   "POST",
		Path:     "/v1/keys/{key_id}/encrypt",
		PathArgs: []string{keyID},
		Body:     body,
	}
	var out EncryptResponse
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// GenerateDataKey generates a fresh data key.
//
// Returns a freshly-random plaintext data key plus its wrapped form
// (encrypted under this KMS key). Use the plaintext immediately and drop
// it from memory; persist the ciphertext alongside the data the key
// protects and re-call Decrypt to recover the plaintext on demand.
func (c *Client) GenerateDataKey(ctx context.Context, keyID string, body *GenerateDataKeyRequest, opts ...basaltic.RequestOption) (*GenerateDataKeyResponse, error) {
	op := &basaltic.Operation{
		ID:       "generateDataKey",
		Method:   "POST",
		Path:     "/v1/keys/{key_id}/generate-data-key",
		PathArgs: []string{keyID},
		Body:     body,
	}
	var out GenerateDataKeyResponse
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetKey gets a KMS key.
func (c *Client) GetKey(ctx context.Context, keyID string, opts ...basaltic.RequestOption) (*Key, error) {
	op := &basaltic.Operation{
		ID:       "getKey",
		Method:   "GET",
		Path:     "/v1/keys/{key_id}",
		PathArgs: []string{keyID},
	}
	var out struct {
		Key *Key `json:"key"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Key, nil
}

// ListKeys lists KMS keys.
//
// List KMS keys owned by the requesting account, newest-first.
// Keyset-paginated by key id (UUIDv7 sorts by creation time) — pass
// the last id from the previous page as `marker` to fetch the next.
//
// Returns one page. Use ListKeysAll to walk every page.
func (c *Client) ListKeys(ctx context.Context, params *ListKeysParams, opts ...basaltic.RequestOption) (*basaltic.Page[Key], error) {
	op := &basaltic.Operation{
		ID:     "listKeys",
		Method: "GET",
		Path:   "/v1/keys",
	}
	op.Query = params.query()
	var out struct {
		Items []Key `json:"keys"`
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
	page := &basaltic.Page[Key]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListKeysAll walks every page of ListKeys, yielding one item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListKeysAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListKeysAll(ctx context.Context, params *ListKeysParams, opts ...basaltic.RequestOption) iter.Seq2[Key, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Key], error) {
		return c.ListKeys(ctx, params.withMarker(marker), opts...)
	})
}

// ScheduleKeyDeletion schedules key for deletion.
//
// Move the key into state=pending_deletion. Quota is released
// immediately so the customer can create a fresh key inside the same
// quota; the key material + record are hard-deleted once now() reaches
// deletion_scheduled_at. The caller can cancel any time inside the
// window via cancel-deletion.
func (c *Client) ScheduleKeyDeletion(ctx context.Context, keyID string, body *ScheduleKeyDeletionRequest, opts ...basaltic.RequestOption) (*Key, error) {
	op := &basaltic.Operation{
		ID:       "scheduleKeyDeletion",
		Method:   "POST",
		Path:     "/v1/keys/{key_id}/schedule-deletion",
		PathArgs: []string{keyID},
		Body:     body,
	}
	var out struct {
		Key *Key `json:"key"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Key, nil
}

// Sign signs a message.
//
// Refused unless key_usage=sign_verify and state=enabled. The service
// hashes the message internally via SHA-256, so pass the raw payload —
// do NOT pre-hash.
func (c *Client) Sign(ctx context.Context, keyID string, body *SignRequest, opts ...basaltic.RequestOption) (*SignResponse, error) {
	op := &basaltic.Operation{
		ID:       "sign",
		Method:   "POST",
		Path:     "/v1/keys/{key_id}/sign",
		PathArgs: []string{keyID},
		Body:     body,
	}
	var out SignResponse
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdateKey updates key metadata.
//
// Partial update of name, description, and tags. The key must not be in
// state=pending_deletion — cancel deletion first.
func (c *Client) UpdateKey(ctx context.Context, keyID string, body *UpdateKeyRequest, opts ...basaltic.RequestOption) (*Key, error) {
	op := &basaltic.Operation{
		ID:       "updateKey",
		Method:   "PATCH",
		Path:     "/v1/keys/{key_id}",
		PathArgs: []string{keyID},
		Body:     body,
	}
	var out struct {
		Key *Key `json:"key"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Key, nil
}

// Verify verifies a signature.
func (c *Client) Verify(ctx context.Context, keyID string, body *VerifyRequest, opts ...basaltic.RequestOption) (bool, error) {
	op := &basaltic.Operation{
		ID:       "verify",
		Method:   "POST",
		Path:     "/v1/keys/{key_id}/verify",
		PathArgs: []string{keyID},
		Body:     body,
	}
	var out struct {
		SignatureValid bool `json:"signature_valid"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return false, err
	}
	return out.SignatureValid, nil
}
