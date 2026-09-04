// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package certificate

import (
	"context"
	"iter"
	"net/url"
	"strconv"

	basaltic "github.com/basaltic-sh/sdk-go"
)

// ListCertificatesParams are the optional filters and pagination controls for
// [Client.ListCertificates]. A nil *ListCertificatesParams sends none of them.
type ListCertificatesParams struct {
	Limit int

	// Marker resume token — the last certificate id from the previous page.
	Marker string
}

// query renders the parameters that are set. A zero value means "no
// filter", which is what leaving one out asks for.
func (p *ListCertificatesParams) query() url.Values {
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
func (p *ListCertificatesParams) withMarker(marker string) *ListCertificatesParams {
	var out ListCertificatesParams
	if p != nil {
		out = *p
	}
	out.Marker = marker
	return &out
}

// CreateCertificate creates certificate.
//
// Create a new certificate. Defaults to ACME issuance; set
// `source=uploaded` to store customer-supplied PEM material directly.
//
// For ACME source: returns 202 with status pending and an empty
// `challenges` list — the challenges are planned by the issuance
// workflow just after the row is created, not by this call. Poll GET
// /v1/certificates/{certificate_id} to pick them up: the row moves to
// pending_dns and grows one `challenges` entry per domain. The CNAME is
// created automatically when our_dns=true, otherwise the customer must
// add `_acme-challenge.<domain> CNAME <expected_cname>` at their
// registrar. The cert flips to active once every challenge verifies and
// issuance completes.
//
// For uploaded source: returns 201 with status active.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) CreateCertificate(ctx context.Context, body *CertificateIssueRequest, opts ...basaltic.RequestOption) (*Certificate, error) {
	op := &basaltic.Operation{
		ID:     "createCertificate",
		Method: "POST",
		Path:   "/v1/certificates",
		Body:   body,
	}
	var out struct {
		Certificate *Certificate `json:"certificate"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Certificate, nil
}

// DeleteCertificate deletes certificate.
//
// Delete a certificate.
func (c *Client) DeleteCertificate(ctx context.Context, certificateID string, opts ...basaltic.RequestOption) error {
	op := &basaltic.Operation{
		ID:       "deleteCertificate",
		Method:   "DELETE",
		Path:     "/v1/certificates/{certificate_id}",
		PathArgs: []string{certificateID},
	}
	if err := c.rt.Do(ctx, op, nil, opts...); err != nil {
		return err
	}
	return nil
}

// GetCertificate gets certificate.
//
// Fetch a certificate by ID. Use this to poll issuance progress.
func (c *Client) GetCertificate(ctx context.Context, certificateID string, opts ...basaltic.RequestOption) (*Certificate, error) {
	op := &basaltic.Operation{
		ID:       "getCertificate",
		Method:   "GET",
		Path:     "/v1/certificates/{certificate_id}",
		PathArgs: []string{certificateID},
	}
	var out struct {
		Certificate *Certificate `json:"certificate"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Certificate, nil
}

// GetCertificateMaterial fetches certificate material (leaf, chain, private key).
//
// Return the certificate's full PEM bundle including the DECRYPTED
// private key. This is the one path that surfaces key material; every
// other read omits it.
//
// Requires `certificate:GetCertificateMaterial` on the certificate's
// CRN. That is a different action from `certificate:GetCertificate`, so
// reading and listing certificates does not reach the key — and it is
// the only thing gating this endpoint, so grant it deliberately.
//
// The `fingerprint` accompanies the PEM blocks. A caller replacing what
// it currently serves can compare it to confirm it has the generation it
// meant to install: a renewed certificate has a new one.
func (c *Client) GetCertificateMaterial(ctx context.Context, certificateID string, opts ...basaltic.RequestOption) (*Material, error) {
	op := &basaltic.Operation{
		ID:       "getCertificateMaterial",
		Method:   "GET",
		Path:     "/v1/certificates/{certificate_id}/material",
		PathArgs: []string{certificateID},
	}
	var out struct {
		Material *Material `json:"material"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Material, nil
}

// ListCertificates lists certificates.
//
// List certificates owned by the requesting account (the one named by
// X-Account-Id), newest-first. Keyset-paginated by certificate id
// (UUIDv7 sorts by creation time) — pass the last id from the previous
// page as `marker` to fetch the next.
//
// Returns one page. Use ListCertificatesAll to walk every page.
func (c *Client) ListCertificates(ctx context.Context, params *ListCertificatesParams, opts ...basaltic.RequestOption) (*basaltic.Page[Certificate], error) {
	op := &basaltic.Operation{
		ID:     "listCertificates",
		Method: "GET",
		Path:   "/v1/certificates",
	}
	op.Query = params.query()
	var out struct {
		Items []Certificate `json:"certificates"`
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
	page := &basaltic.Page[Certificate]{Items: out.Items}
	if out.Meta != nil {
		page.Total = out.Meta.Total
		page.Limit = out.Meta.Limit
		page.Marker = out.Meta.Marker
		page.HasMore = out.Meta.HasMore
	}
	return page, nil
}

// ListCertificatesAll walks every page of ListCertificates, yielding one
// item at a time.
//
// The iterator stops at the first error, yielding it alongside a zero
// value, so check err on every step:
//
//	for item, err := range c.ListCertificatesAll(ctx, nil) {
//		if err != nil {
//			return err
//		}
//		...
//	}
//
// Breaking out of the loop stops the walk; no further requests are made.
// Any Marker on params is overwritten as the walk advances.
func (c *Client) ListCertificatesAll(ctx context.Context, params *ListCertificatesParams, opts ...basaltic.RequestOption) iter.Seq2[Certificate, error] {
	return basaltic.Paginate(ctx, func(ctx context.Context, marker string) (*basaltic.Page[Certificate], error) {
		return c.ListCertificates(ctx, params.withMarker(marker), opts...)
	})
}

// RevokeCertificate revokes certificate.
//
// Request revocation through the platform CA and flip the certificate to
// revoked. Async — returns 202 with the cert in its previous state;
// poll GET until status=revoked.
//
// Uploaded certs are revoked locally without contacting the CA.
// Re-revoking an already-revoked cert is a no-op.
//
// Accepts basaltic.WithIdempotencyKey, which makes the call
// replay-safe and therefore retryable.
func (c *Client) RevokeCertificate(ctx context.Context, certificateID string, opts ...basaltic.RequestOption) (*Certificate, error) {
	op := &basaltic.Operation{
		ID:       "revokeCertificate",
		Method:   "POST",
		Path:     "/v1/certificates/{certificate_id}/revoke",
		PathArgs: []string{certificateID},
	}
	var out struct {
		Certificate *Certificate `json:"certificate"`
	}
	if err := c.rt.Do(ctx, op, &out, opts...); err != nil {
		return nil, err
	}
	return out.Certificate, nil
}
