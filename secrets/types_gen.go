// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package secrets

import (
	"time"
)

type CreateSecretRequest struct {
	Description *string `json:"description,omitempty"`

	// KMSKeyID Customer-managed KMS key to encrypt this secret under. Omit to use
	// the platform-managed default key. The key must be enabled and have
	// encrypt/decrypt usage.
	KMSKeyID *string `json:"kms_key_id,omitempty"`

	// Name unique within the calling account.
	//
	// Required.
	Name                  string `json:"name"`
	RecoveryWindowSeconds *int   `json:"recovery_window_seconds,omitempty"`
	Tags                  Tags   `json:"tags,omitempty"`

	// Value base64 of the initial value bytes (1 byte - 64 KiB).
	//
	// Required.
	Value []byte `json:"value"`
}

type DeleteSecretRequest struct {
	// RecoveryWindowSeconds override the secret's default window. Omit to keep it.
	RecoveryWindowSeconds *int `json:"recovery_window_seconds,omitempty"`
}

type PutSecretValueRequest struct {
	// Value base64 of the new value bytes (1 byte - 64 KiB).
	//
	// Required.
	Value []byte `json:"value"`
}

type Secret struct {
	CreatedAt time.Time `json:"created_at"`
	CRN       string    `json:"crn"`

	// CurrentVersion 0 if no version exists yet.
	CurrentVersion int       `json:"current_version,omitempty"`
	DeletedAt      time.Time `json:"deleted_at,omitempty"`
	Description    string    `json:"description,omitempty"`
	ID             string    `json:"id"`

	// KMSKeyID Customer-managed KMS key the secret is encrypted under. Null
	// (omitted) when encrypted with the platform-managed default key.
	KMSKeyID string `json:"kms_key_id,omitempty"`

	// Managed true when a platform service generated this value and reads it back
	// to act on. You can read and delete a managed secret, but
	// UpdateSecret and PutSecretValue answer 403 SECRET_PLATFORM_MANAGED.
	Managed               bool      `json:"managed"`
	Name                  string    `json:"name"`
	RecoveryWindowSeconds int       `json:"recovery_window_seconds"`
	ScheduledPurgeAt      time.Time `json:"scheduled_purge_at,omitempty"`
	Tags                  Tags      `json:"tags,omitempty"`
	UpdatedAt             time.Time `json:"updated_at"`
}

type SecretValue struct {
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	SecretID  string    `json:"secret_id"`

	// Value base64 of the plaintext bytes.
	Value   []byte `json:"value"`
	Version int    `json:"version"`
}

type SecretVersion struct {
	CreatedAt time.Time `json:"created_at"`

	// CreatedBy CRN of the principal that created this version (e.g.
	// crn:iam:::user/<id>, crn:iam:::service-account/<id>).
	CreatedBy string `json:"created_by,omitempty"`
	ID        string `json:"id"`
	IsCurrent bool   `json:"is_current"`
	Version   int    `json:"version"`
}

type Tags = map[string]string

type UpdateSecretRequest struct {
	Description *string `json:"description,omitempty"`
	Tags        Tags    `json:"tags,omitempty"`
}
