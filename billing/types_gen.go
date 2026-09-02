// Code generated from the Basaltic OpenAPI specifications. DO NOT EDIT.
//
// Regenerate with:
//
//	go run ./internal/gen -spec /path/to/openapi

package billing

import (
	"time"
)

type Credit struct {
	Amount      string    `json:"amount"`
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
	ExpiresAt   time.Time `json:"expires_at,omitempty"`
	ID          string    `json:"id"`
	Remaining   string    `json:"remaining"`

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

type Invoice struct {
	CreatedAt      time.Time `json:"created_at"`
	CreditsApplied string    `json:"credits_applied"`
	Currency       string    `json:"currency"`
	DueAt          time.Time `json:"due_at,omitempty"`
	ID             string    `json:"id"`
	InvoiceNumber  string    `json:"invoice_number"`
	IssuedAt       time.Time `json:"issued_at,omitempty"`

	// Items line items; present only on the detail endpoint.
	Items  []*InvoiceItem `json:"items,omitempty"`
	PaidAt time.Time      `json:"paid_at,omitempty"`

	// PDFURL path of the PDF statement, rendered on demand by GET
	// /v1/invoices/{invoice_id}/pdf under the same authorization as this
	// document.
	PDFURL string `json:"pdf_url,omitempty"`

	// PeriodEnd exclusive end (first day of the following month).
	PeriodEnd string `json:"period_end"`

	// PeriodStart first day of the billed UTC month.
	PeriodStart string `json:"period_start"`

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
	Kind      string `json:"kind"`
	Quantity  string `json:"quantity"`
	Sku       string `json:"sku,omitempty"`
	Unit      string `json:"unit,omitempty"`
	UnitPrice string `json:"unit_price"`
}

type Payment struct {
	Amount string `json:"amount"`

	// Attempt 1-based dunning attempt this payment row belongs to.
	Attempt     int32     `json:"attempt"`
	CompletedAt time.Time `json:"completed_at,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	ID          string    `json:"id"`
	InvoiceID   string    `json:"invoice_id,omitempty"`

	// One of: "pending", "processing", "succeeded", "failed", "refunded".
	Status string `json:"status"`
}

// Price one effective row of the public price catalog — the same
// `billing.billing_prices` row rating charges against. Money is a
// decimal string rather than a JSON number so the quoted rate is exactly
// the one that will be billed.
type Price struct {
	Currency    string `json:"currency"`
	Description string `json:"description,omitempty"`

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

	// ValidFrom when this revision took effect.
	ValidFrom time.Time `json:"valid_from"`

	// ValidTo when the next revision supersedes it, or null while this is the
	// current price.
	ValidTo time.Time `json:"valid_to,omitempty"`
}

type PriceListResponse struct {
	// AsOf the instant the catalog was read as of — the `at` that was asked
	// for, or the server's clock when none was.
	AsOf   time.Time `json:"as_of"`
	Prices []*Price  `json:"prices"`
}

type Transaction struct {
	// Amount always positive; the direction lives in the type.
	Amount        string    `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
	Description   string    `json:"description,omitempty"`
	ID            string    `json:"id"`
	ReferenceType string    `json:"reference_type,omitempty"`

	// Type ledger entry type.
	//
	// One of: "payment", "refund", "adjustment", "credit_grant", "credit_applied".
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
