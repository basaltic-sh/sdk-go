// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package billing

import (
	"time"
)

// BillingProfile organization billing recipient. Incomplete drafts are saved; ready
// becomes true only after all country-specific fields are valid. New
// onboarding completes after this step.
type BillingProfile struct {
	City *string `json:"city,omitempty"`

	// CompanyName full legal name of the individual or company.
	CompanyName *string `json:"company_name,omitempty"`
	Complement  *string `json:"complement,omitempty"`

	// Country ISO 3166-1 alpha-2 country code.
	Country *string `json:"country,omitempty"`

	// One of: "", "individual", "company".
	CustomerType *string `json:"customer_type,omitempty"`

	// Email billing email for fiscal invoice delivery. The onboarding form
	// prefills this from the signed-in user's email.
	Email *string `json:"email,omitempty"`

	// ForeignTaxID foreign identifier; not validated as a Brazilian document.
	ForeignTaxID  *string  `json:"foreign_tax_id,omitempty"`
	MissingFields []string `json:"missing_fields,omitempty"`

	// MunicipalityCode Seven-digit IBGE municipality code, required for a Brazilian
	// recipient.
	MunicipalityCode *string `json:"municipality_code,omitempty"`
	Neighborhood     *string `json:"neighborhood,omitempty"`

	// NoTaxIDReason required for a foreign recipient without a tax identifier.
	NoTaxIDReason *string `json:"no_tax_id_reason,omitempty"`
	Phone         *string `json:"phone,omitempty"`

	// PostalCode Eight-digit CEP for Brazil; optional international postal code
	// abroad.
	PostalCode *string `json:"postal_code,omitempty"`
	Ready      *bool   `json:"ready,omitempty"`

	// State Two-letter UF for Brazil; free-form state/province abroad.
	State        *string `json:"state,omitempty"`
	StreetName   *string `json:"street_name,omitempty"`
	StreetNumber *string `json:"street_number,omitempty"`

	// TaxID CPF for a Brazilian individual or CNPJ for a Brazilian company.
	// Check digits are validated.
	TaxID *string `json:"tax_id,omitempty"`
}

type Credit struct {
	Amount    string    `json:"amount"`
	CreatedAt time.Time `json:"created_at"`

	// CRN global organization-scoped credit identity.
	CRN         string     `json:"crn"`
	Description string     `json:"description"`
	ExpiresAt   *time.Time `json:"expires_at,omitempty"`
	ID          string     `json:"id"`
	Remaining   string     `json:"remaining"`

	// One of: "promo", "coupon", "adjustment", "migration".
	Source string `json:"source"`
}

type CurrentUsage struct {
	// Amount unbilled usage accrued this UTC month, 2-decimal string.
	Amount string `json:"amount"`

	// Items Per-SKU breakdown, ordered by cost.
	Items       []*UsageLine `json:"items"`
	PeriodStart time.Time    `json:"period_start"`
}

// FiscalInvoice One NFS-e per confirmed provider payment, including partial payments
// or credit top-ups. Consuming prepaid credit never issues the same
// money twice. Recipient details and receipt amount are frozen before
// submission.
type FiscalInvoice struct {
	// Amount actual amount received in BRL, not the billing invoice total.
	Amount    string    `json:"amount"`
	Attempts  int       `json:"attempts"`
	CreatedAt time.Time `json:"created_at"`

	// EmailStatus separate delivery state. Queued means durably accepted by the
	// internal email service; it does not assert recipient delivery.
	//
	// One of: "pending", "queued".
	EmailStatus string    `json:"email_status"`
	ID          string    `json:"id"`
	InvoiceID   *string   `json:"invoice_id,omitempty"`
	IssuedAt    time.Time `json:"issued_at,omitempty"`

	// LastError sanitized operational error or municipal rejection codes.
	LastError      string `json:"last_error,omitempty"`
	Number         string `json:"number,omitempty"`
	OrganizationID string `json:"organization_id"`
	PaymentID      string `json:"payment_id"`

	// RequiresReview refunds preserve the fiscal document and require operator review;
	// cancellation is never inferred automatically.
	RequiresReview bool `json:"requires_review"`

	// One of: "queued", "waiting_details", "retrying", "rejected", "issued", "review_required".
	Status string `json:"status"`

	// URL municipal view/print link available after issuance.
	URL              string `json:"url,omitempty"`
	VerificationCode string `json:"verification_code,omitempty"`
}

type Invoice struct {
	CreatedAt      time.Time `json:"created_at"`
	CreditsApplied string    `json:"credits_applied"`

	// CRN global organization-scoped invoice identity.
	CRN      string `json:"crn"`
	Currency string `json:"currency"`

	// DisputedAmount dispute principal withdrawn less funds reinstated; excludes provider
	// fees.
	DisputedAmount string     `json:"disputed_amount"`
	DueAt          *time.Time `json:"due_at,omitempty"`
	ID             string     `json:"id"`
	InvoiceNumber  string     `json:"invoice_number"`
	IssuedAt       *time.Time `json:"issued_at,omitempty"`

	// Items line items; present only on the detail endpoint.
	Items  []*InvoiceItem `json:"items,omitempty"`
	PaidAt *time.Time     `json:"paid_at,omitempty"`

	// PDFURL path of the PDF statement, rendered on demand by GET
	// /v1/invoices/{invoice_id}/pdf under the same authorization as this
	// document.
	PDFURL string `json:"pdf_url,omitempty"`

	// PeriodEnd exclusive end (first day of the following month).
	PeriodEnd string `json:"period_end"`

	// PeriodStart first day of the billed UTC month.
	PeriodStart string `json:"period_start"`

	// RefundedAmount confirmed refunds less failed-refund reversals, in BRL.
	RefundedAmount string `json:"refunded_amount"`

	// One of: "open", "paid", "past_due", "uncollectible", "void".
	Status   string `json:"status"`
	Subtotal string `json:"subtotal"`
	Total    string `json:"total"`
}

type InvoiceItem struct {
	// Amount rounded line total; negative for credit lines.
	Amount      string `json:"amount"`
	Description string `json:"description"`

	// One of: "usage", "credit".
	Kind      string  `json:"kind"`
	Quantity  string  `json:"quantity"`
	Sku       *string `json:"sku,omitempty"`
	Unit      *string `json:"unit,omitempty"`
	UnitPrice string  `json:"unit_price"`
}

type Payment struct {
	Amount string `json:"amount"`

	// Attempt 1-based dunning attempt this payment row belongs to.
	Attempt     int32      `json:"attempt"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`

	// CRN global organization-scoped payment identity.
	CRN string `json:"crn"`

	// DisputedAmount dispute principal withdrawn less funds reinstated; excludes provider
	// fees.
	DisputedAmount string `json:"disputed_amount"`
	ID             string `json:"id"`

	// Invoice current invoice list shape, without items; null when the invoice has
	// been deleted.
	Invoice *Invoice `json:"invoice"`

	// RefundedAmount confirmed refunds less failed-refund reversals, in BRL.
	RefundedAmount string `json:"refunded_amount"`

	// RetainedAmount settled receipt less refunds and disputed funds. Zero for an
	// unsettled attempt; may be negative if the provider has withdrawn
	// overlapping reversals. Does not change invoice collection status.
	RetainedAmount string `json:"retained_amount"`

	// One of: "pending", "processing", "succeeded", "failed", "refunded".
	Status string `json:"status"`
}

// Price one effective row of the public price catalog — the same
// `billing.billing_prices` row rating charges against. Money is a
// decimal string rather than a JSON number so the quoted rate is exactly
// the one that will be billed.
type Price struct {
	Currency    string  `json:"currency"`
	Description *string `json:"description,omitempty"`

	// Metadata extra facts about the SKU — `class`, `family`, `vcpus`,
	// `memory_gb`, `storage_type`, … `family` separates the managed
	// products (load balancer replicas, database cluster nodes) from the
	// general compute flavors they share a `resource_type` with.
	Metadata map[string]any `json:"metadata"`

	// Name display name. For compute SKUs this is the flavor name.
	Name         string `json:"name"`
	ResourceType string `json:"resource_type"`

	// Service which service bills this SKU.
	Service string `json:"service"`

	// Sku stable catalog key, `{service}.{resource_type}.{variant}`. This is
	// the public identity of a price — the row id is not published.
	Sku string `json:"sku"`

	// Unit what one unit of `unit_price` buys.
	Unit string `json:"unit"`

	// UnitPrice price for one `unit`, as an exact decimal string.
	UnitPrice string `json:"unit_price"`
}

type PriceListResponse struct {
	// AsOf the instant the catalog was read as of — the `at` that was asked
	// for, or the server's clock when none was.
	AsOf   time.Time `json:"as_of"`
	Prices []*Price  `json:"prices"`
}

type Transaction struct {
	// Amount always positive; the direction lives in the type.
	Amount    string    `json:"amount"`
	CreatedAt time.Time `json:"created_at"`

	// CRN global organization-scoped transaction identity.
	CRN         string  `json:"crn"`
	Description *string `json:"description,omitempty"`
	ID          string  `json:"id"`

	// Reference organization-scoped reference to the ledger entry's target:
	// crn:billing:::invoice/<id>, crn:billing:::payment/<id>, or
	// crn:billing:::credit/<id> for a credit grant. Pass this CRN to the
	// corresponding collection's crn filter within the authenticated
	// organization. Null for manual entries, unsupported reference types,
	// or missing references.
	Reference *string `json:"reference,omitempty"`

	// Type ledger entry type.
	//
	// One of: "payment", "refund", "refund_reversal", "dispute", "dispute_reversal", "adjustment", "credit_grant", "credit_applied".
	Type string `json:"type"`
}

type UsageLine struct {
	// Amount accrued cost at 4-decimal precision (sub-centavo lines stay visible
	// mid-month).
	Amount      string `json:"amount"`
	Description string `json:"description"`
	Quantity    string `json:"quantity"`
	Sku         string `json:"sku"`
	Unit        string `json:"unit"`
}
