// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package certificate

import (
	"time"
)

type Certificate struct {
	// CertificatePEM PEM-encoded leaf certificate. Empty until active.
	CertificatePEM string `json:"certificate_pem,omitempty"`

	// ChainPEM PEM-encoded intermediate chain.
	ChainPEM string `json:"chain_pem,omitempty"`

	// Challenges Per-domain CNAME delegation state — one entry per domain on the
	// cert. While status=pending_dns issuance waits for every challenge's
	// `verified` to flip true; use `expected_cname` and `our_dns` to tell
	// which records need to be added at the registrar.
	Challenges []*CertificateChallenge `json:"challenges,omitempty"`

	// CRN name-based, so an IAM policy can wildcard a naming convention
	// (`crn:certificate::my-account:certificate/prod-*`). The region slot
	// is empty — certs are not region-bound.
	CRN     string   `json:"crn,omitempty"`
	Domains []string `json:"domains,omitempty"`

	// ErrorMessage last failure reason (set when status=error).
	ErrorMessage string    `json:"error_message,omitempty"`
	ExpiresAt    time.Time `json:"expires_at,omitempty"`

	// Fingerprint Hex SHA-256 of the leaf's DER — the certificate's material
	// version. Changes on every rotation; consumers use it to know when to
	// re-fetch material and to verify they fetched the intended
	// generation. Empty until the cert has a leaf.
	Fingerprint  string                  `json:"fingerprint,omitempty"`
	ID           string                  `json:"id,omitempty"`
	IssuedAt     time.Time               `json:"issued_at,omitempty"`
	KeyAlgorithm CertificateKeyAlgorithm `json:"key_algorithm,omitempty"`
	Name         string                  `json:"name,omitempty"`
	Source       CertificateSource       `json:"source,omitempty"`
	Status       CertificateStatus       `json:"status,omitempty"`
	Tags         Tags                    `json:"tags,omitempty"`
}

type CertificateChallenge struct {
	// CnameRecordName Full LHS of the CNAME record the customer needs to add. For wildcard
	// SANs this is the parent name (`_acme-challenge.example.com.`), not
	// the literal SAN — wildcards validate at their parent under RFC
	// 8555 §8.4.
	CnameRecordName string `json:"cname_record_name,omitempty"`

	// Domain the cert SAN this challenge belongs to (as the customer wrote it).
	Domain string `json:"domain,omitempty"`

	// ErrorMessage last verification failure (e.g. "CNAME = foo, want bar"). Cleared
	// when the challenge eventually verifies.
	ErrorMessage string `json:"error_message,omitempty"`

	// ExpectedCname Target FQDN (RHS of the CNAME). Hosted in the platform's validation
	// zone, where the per-order TXT is published during issuance.
	ExpectedCname string `json:"expected_cname,omitempty"`

	// OurDNS true when the domain is hosted on the platform DNS service and the
	// CNAME was created automatically. False means the customer owns the
	// zone and must add the CNAME at their registrar.
	OurDNS bool `json:"our_dns,omitempty"`

	// Verified true once the CNAME has resolved to expected_cname; cert issuance
	// only proceeds when every challenge is verified.
	Verified   bool      `json:"verified,omitempty"`
	VerifiedAt time.Time `json:"verified_at,omitempty"`
}

type CertificateIssueRequest struct {
	// CertificatePEM PEM-encoded leaf certificate. Required when source=uploaded.
	CertificatePEM *string `json:"certificate_pem,omitempty"`

	// ChainPEM PEM-encoded intermediate chain (optional when source=uploaded).
	ChainPEM *string `json:"chain_pem,omitempty"`

	// Domains capped at 100 to stay inside the certificate authority's per-order
	// limits.
	//
	// Required.
	Domains      []string                 `json:"domains"`
	KeyAlgorithm *CertificateKeyAlgorithm `json:"key_algorithm,omitempty"`

	// Name unique per account. Surfaces in the CRN
	// (`crn:certificate::<account>:certificate/<name>`), so it must be
	// URL-safe — letters, digits, dot, dash, underscore.
	//
	// Required.
	Name string `json:"name"`

	// PrivateKeyPEM PEM-encoded private key. Required when source=uploaded.
	PrivateKeyPEM *string `json:"private_key_pem,omitempty"`

	// Source defaults to "acme" — issued by the platform CA. Set to "uploaded"
	// to store customer-supplied PEM material instead; certificate_pem +
	// private_key_pem must then be provided.
	Source *CertificateSource `json:"source,omitempty"`
	Tags   Tags               `json:"tags,omitempty"`
}

type CertificateKeyAlgorithm string

// Values CertificateKeyAlgorithm accepts.
const (
	CertificateKeyAlgorithmEcdsaP256 CertificateKeyAlgorithm = "ecdsa-p256"
	CertificateKeyAlgorithmEcdsaP384 CertificateKeyAlgorithm = "ecdsa-p384"
	CertificateKeyAlgorithmRSA2048   CertificateKeyAlgorithm = "rsa-2048"
	CertificateKeyAlgorithmRSA4096   CertificateKeyAlgorithm = "rsa-4096"
)

type CertificateSource string

// Values CertificateSource accepts.
const (
	CertificateSourceAcme     CertificateSource = "acme"
	CertificateSourceUploaded CertificateSource = "uploaded"
)

type CertificateStatus string

// Values CertificateStatus accepts.
const (
	CertificateStatusPendingDNS CertificateStatus = "pending_dns"
	CertificateStatusPending    CertificateStatus = "pending"
	CertificateStatusActive     CertificateStatus = "active"
	CertificateStatusError      CertificateStatus = "error"
	CertificateStatusExpired    CertificateStatus = "expired"
	CertificateStatusRevoked    CertificateStatus = "revoked"
)

// Material a certificate's full PEM bundle including the decrypted private key.
// Returned ONLY by the material endpoint, which requires the stronger
// certificate:GetCertificateMaterial action; every other read omits the
// key.
type Material struct {
	// CertificatePEM PEM-encoded leaf certificate.
	CertificatePEM string `json:"certificate_pem,omitempty"`

	// ChainPEM PEM-encoded intermediate chain.
	ChainPEM string `json:"chain_pem,omitempty"`

	// Fingerprint Hex SHA-256 of the leaf's DER. A caller confirms this matches the
	// version the feed named before installing the material.
	Fingerprint string `json:"fingerprint,omitempty"`

	// PrivateKeyPEM PEM-encoded private key (decrypted).
	PrivateKeyPEM string `json:"private_key_pem,omitempty"`
}

type Tags = map[string]string
