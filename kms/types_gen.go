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
	// dash, underscore.
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
	// Ciphertext base64-encoded ciphertext. Opaque — store verbatim. For AES-GCM
	// the layout is nonce(12) || ct || tag; for RSA-OAEP the standard
	// PKCS#1 RSAES output.
	Ciphertext []byte `json:"ciphertext,omitempty"`

	// KeyCRN CRN of the key the ciphertext was sealed under.
	KeyCRN string `json:"key_crn,omitempty"`
}

type GenerateDataKeyRequest struct {
	// NumberOfBytes size of the generated data key in bytes: 16 (AES-128), 32 (AES-256,
	// the default) or 64 (HMAC-SHA512). No other size is supported — the
	// HSM mints data keys at those three widths only, and any other value
	// fails the operation.
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

	// DeletionScheduledAt set only while state=pending_deletion. The key (and its
	// cryptographic material) is hard-deleted once now() reaches this
	// timestamp; CancelKeyDeletion before then returns the key to
	// state=disabled.
	DeletionScheduledAt time.Time `json:"deletion_scheduled_at,omitempty"`
	Description         string    `json:"description,omitempty"`
	ID                  string    `json:"id"`
	KeySpec             KeySpec   `json:"key_spec"`
	KeyUsage            KeyUsage  `json:"key_usage"`
	Name                string    `json:"name"`
	State               KeyState  `json:"state"`

	// System present and true on platform-owned envelope keys (credential master,
	// JWT signer, …), which are visible but not yours to operate on.
	// Omitted on customer keys. Console and CLI read it to hide the
	// destructive actions.
	System    bool      `json:"system,omitempty"`
	Tags      Tags      `json:"tags"`
	UpdatedAt time.Time `json:"updated_at"`
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
//     deletion_scheduled_at; cancellable until then
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

type ScheduleKeyDeletionRequest struct {
	// PendingWindowInDays how long the key sits in pending_deletion before it is hard-deleted.
	// Matches AWS KMS bounds; the deletion can be cancelled at any point
	// inside the window.
	PendingWindowInDays *int `json:"pending_window_in_days,omitempty"`
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

// UpdateKeyRequest partial update — only the supplied fields are mutated. The key must
// not be in state=pending_deletion.
type UpdateKeyRequest struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
	Tags        Tags    `json:"tags,omitempty"`
}

type VerifyRequest struct {
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
