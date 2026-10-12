// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package kms

import (
	"time"
)

type CreateKeyRequest struct {
	Description *string `json:"description,omitempty"`

	// Required.
	KeySpec KeySpec `json:"key_spec"`

	// KeyUsage required for RSA specs (both encrypt_decrypt and sign_verify are
	// valid). Defaults to encrypt_decrypt for AES, sign_verify for ECDSA.
	KeyUsage *KeyUsage `json:"key_usage,omitempty"`

	// Name unique per account. Surfaces in the CRN
	// (crn:kms:<region>:<account>:key/<name>) — letters, digits, dot,
	// dash, underscore. Resource names must not start with the literal
	// crn: prefix or be UUIDs (canonical, compact, braced, or urn:uuid:
	// forms, in either case).
	//
	// Required.
	Name string `json:"name"`
	Tags Tags   `json:"tags,omitempty"`
}

type DecryptRequest struct {
	// Aad optional base64-encoded AAD. Must match what was supplied at Encrypt
	// — different value fails the tag check.
	Aad []byte `json:"aad,omitempty"`

	// Ciphertext base64-encoded ciphertext produced by Encrypt.
	//
	// Required.
	Ciphertext []byte `json:"ciphertext"`
}

type EncryptRequest struct {
	// Aad optional base64-encoded additional authenticated data (AES-GCM
	// AEAD). Must be supplied verbatim to Decrypt; mismatch fails the auth
	// tag check. Ignored for asymmetric keys.
	Aad []byte `json:"aad,omitempty"`

	// Plaintext base64-encoded plaintext.
	//
	// Required.
	Plaintext []byte `json:"plaintext"`
}

type EncryptResponse struct {
	// Ciphertext base64-encoded opaque, versioned ciphertext. Store verbatim; do not
	// parse its internal layout. Decrypt selects the original material
	// version automatically after rotation.
	Ciphertext []byte `json:"ciphertext,omitempty"`

	// KeyCRN CRN of the key the ciphertext was sealed under.
	KeyCRN string `json:"key_crn,omitempty"`
}

type GenerateDataKeyRequest struct {
	// NumberOfBytes size of the generated data key in bytes: 16 (AES-128), 32 (AES-256,
	// the default) or 64 (HMAC-SHA512). No other size is supported; any
	// other value fails the operation.
	//
	// One of: "16", "32", "64".
	NumberOfBytes *int `json:"number_of_bytes,omitempty"`
}

type GenerateDataKeyResponse struct {
	// Ciphertext base64-encoded data key wrapped under the KMS key. Safe to store at
	// rest alongside the data the key protects.
	Ciphertext []byte `json:"ciphertext,omitempty"`

	// Plaintext base64-encoded plaintext data key. Use immediately, then drop —
	// callers must NOT persist this. Re-derive it on demand by calling
	// Decrypt with the stored ciphertext.
	Plaintext []byte `json:"plaintext,omitempty"`
}

type Key struct {
	CreatedAt time.Time `json:"created_at"`
	CRN       string    `json:"crn"`

	// DeletedAt when deletion was requested. Null for active keys and historical
	// pending deletions whose request time is unknown.
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	Description string     `json:"description,omitempty"`
	ID          string     `json:"id"`
	KeySpec     KeySpec    `json:"key_spec"`
	KeyUsage    KeyUsage   `json:"key_usage"`

	// Name resource names must not start with the literal crn: prefix or be
	// UUIDs (canonical, compact, braced, or urn:uuid: forms, in either
	// case).
	Name string `json:"name"`

	// RecoveryWindowDays chosen whole-day window. Null for active keys and historical pending
	// deletions whose window is unknown.
	RecoveryWindowDays *int `json:"recovery_window_days,omitempty"`

	// ScheduledPurgeAt set only while state=pending_deletion. The key (and its
	// cryptographic material) is hard-deleted once now() reaches this
	// timestamp; CancelKeyDeletion before then returns the key to
	// state=disabled.
	ScheduledPurgeAt *time.Time `json:"scheduled_purge_at,omitempty"`
	State            KeyState   `json:"state"`

	// System present and true on platform-owned envelope keys (credential master,
	// JWT signer, …), which are visible but not yours to operate on.
	// Omitted on customer keys. Console and CLI read it to hide the
	// destructive actions.
	System    bool      `json:"system,omitempty"`
	Tags      Tags      `json:"tags"`
	UpdatedAt time.Time `json:"updated_at"`
}

type KeyRotation struct {
	CompletedAt *time.Time `json:"completed_at"`

	// ErrorCode machine-readable cancellation or reconciliation reason. Null on
	// success.
	ErrorCode *string `json:"error_code"`
	ID        string  `json:"id"`
	KeyID     string  `json:"key_id"`

	// KeyVersion new material version, populated only on success.
	KeyVersion      *int       `json:"key_version"`
	PreviousVersion *int       `json:"previous_version"`
	RequestedAt     time.Time  `json:"requested_at"`
	StartedAt       *time.Time `json:"started_at"`

	// Status pending work is durable. Reconciliation_required means the backend
	// outcome is unknown; no second rotation is sent and further requests
	// are blocked until resolved. Failed means the request was cancelled
	// before dispatch.
	//
	// One of: "pending", "in_progress", "reconciliation_required", "succeeded", "failed".
	Status string `json:"status"`

	// One of: "on_demand", "automatic".
	Trigger string `json:"trigger"`
}

type KeyRotationPolicy struct {
	// Enabled automatic rotation is opt-in and defaults to false.
	Enabled bool `json:"enabled"`

	// NextRotationAt next scheduled rotation, or null when automatic rotation is
	// disabled. Disabled keys are skipped until re-enabled.
	NextRotationAt     *time.Time `json:"next_rotation_at"`
	RotationPeriodDays int        `json:"rotation_period_days"`
}

type KeyRotationStatus struct {
	// CurrentVersion current material version used for new cryptographic operations.
	CurrentVersion int `json:"current_version"`

	// Enabled automatic rotation is opt-in and defaults to false.
	Enabled bool `json:"enabled"`

	// NextRotationAt next scheduled rotation, or null when automatic rotation is
	// disabled. Disabled keys are skipped until re-enabled.
	NextRotationAt     *time.Time `json:"next_rotation_at"`
	RotationPeriodDays int        `json:"rotation_period_days"`
}

// KeySpec cryptographic spec. Determines which mechanisms apply:
//   - aes-256:     symmetric AEAD (AES-GCM)
//   - rsa-2048:    asymmetric; OAEP-SHA256 encrypt + PSS-SHA256 sign
//   - rsa-4096:    same as rsa-2048
//   - ecdsa-p256:  asymmetric; SHA-256 sign/verify (no encrypt)
type KeySpec string

// Values KeySpec accepts.
const (
	KeySpecAES256    KeySpec = "aes-256"
	KeySpecRSA2048   KeySpec = "rsa-2048"
	KeySpecRSA4096   KeySpec = "rsa-4096"
	KeySpecEcdsaP256 KeySpec = "ecdsa-p256"
)

// KeyState lifecycle state:
//   - enabled:          usable for spec/usage operations
//   - disabled:         exists; refuses crypto ops; can be re-enabled
//   - pending_deletion: scheduled for hard delete at
//     scheduled_purge_at; cancellable until then
type KeyState string

// Values KeyState accepts.
const (
	KeyStateEnabled         KeyState = "enabled"
	KeyStateDisabled        KeyState = "disabled"
	KeyStatePendingDeletion KeyState = "pending_deletion"
)

// KeyUsage operation set the key is pinned to at create time. Even when the spec
// supports both (RSA), one usage must be chosen — the service refuses
// operations from the other set.
type KeyUsage string

// Values KeyUsage accepts.
const (
	KeyUsageEncryptDecrypt KeyUsage = "encrypt_decrypt"
	KeyUsageSignVerify     KeyUsage = "sign_verify"
)

type PublicKeyResponse struct {
	KeyCRN     string   `json:"key_crn"`
	KeySpec    KeySpec  `json:"key_spec"`
	KeyUsage   KeyUsage `json:"key_usage"`
	KeyVersion int      `json:"key_version"`

	// PublicKey PEM-encoded SubjectPublicKeyInfo for the requested asymmetric
	// material version. No private material is exported.
	PublicKey string `json:"public_key"`
}

type RewrapRequest struct {
	// Ciphertext base64-encoded ciphertext or wrapped data key produced by this KMS
	// key. AAD is neither required nor changed.
	//
	// Required.
	Ciphertext []byte `json:"ciphertext"`
}

type ScheduleKeyDeletionRequest struct {
	// RecoveryWindowDays how long the key sits in pending_deletion before it is hard-deleted.
	// Matches AWS KMS bounds; the deletion can be cancelled at any point
	// inside the window.
	RecoveryWindowDays *int `json:"recovery_window_days,omitempty"`
}

type SetKeyRotationRequest struct {
	// Required.
	Enabled bool `json:"enabled"`

	// RotationPeriodDays whole days between rotations. Omission selects 365 days.
	RotationPeriodDays *int `json:"rotation_period_days,omitempty"`
}

type SignRequest struct {
	// Message base64-encoded message to sign. The service hashes it via SHA-256
	// server-side, so pass the raw payload — do not pre-hash.
	//
	// Required.
	Message          []byte            `json:"message"`
	SigningAlgorithm *SigningAlgorithm `json:"signing_algorithm,omitempty"`
}

type SignResponse struct {
	// KeyVersion material version that produced this signature. Persist with the
	// signature for explicit verification and public-key retrieval.
	KeyVersion int `json:"key_version,omitempty"`

	// Signature base64-encoded signature. RSA-PSS: PKCS#1 octet string with
	// saltLen=hashLen=32. ECDSA: ASN.1 DER (r,s) tuple per ANSI X9.62.
	Signature        []byte           `json:"signature,omitempty"`
	SigningAlgorithm SigningAlgorithm `json:"signing_algorithm,omitempty"`
}

// SigningAlgorithm signing algorithm for asymmetric Sign / Verify. Optional in requests
// — when omitted, the service picks the default for the key spec (RSA
// → RSASSA_PSS_SHA_256, ECDSA → ECDSA_SHA_256). Must match the key
// spec: RSA keys accept the two RSASSA schemes, ECDSA keys accept
// ECDSA_SHA_256.
type SigningAlgorithm string

// Values SigningAlgorithm accepts.
const (
	SigningAlgorithmRsassaPssSHA256      SigningAlgorithm = "RSASSA_PSS_SHA_256"
	SigningAlgorithmRsassaPkcs1V15SHA256 SigningAlgorithm = "RSASSA_PKCS1_V1_5_SHA_256"
	SigningAlgorithmEcdsaSHA256          SigningAlgorithm = "ECDSA_SHA_256"
)

type Tags = map[string]string

// UpdateKeyRequest names are fixed at creation because they form the CRN used by IAM
// policies. Sending name in an update, including an unchanged, empty or
// null value, returns a validation error. Partial update — only the
// supplied fields are mutated. The key must not be in
// state=pending_deletion.
type UpdateKeyRequest struct {
	Description *string `json:"description,omitempty"`
	Tags        Tags    `json:"tags,omitempty"`
}

type VerifyRequest struct {
	// KeyVersion optional material version returned by Sign. When omitted, all
	// retained public-key versions are tried for compatibility with older
	// signatures.
	KeyVersion *int `json:"key_version,omitempty"`

	// Message base64-encoded original message.
	//
	// Required.
	Message []byte `json:"message"`

	// Signature base64-encoded signature produced by Sign.
	//
	// Required.
	Signature        []byte            `json:"signature"`
	SigningAlgorithm *SigningAlgorithm `json:"signing_algorithm,omitempty"`
}
