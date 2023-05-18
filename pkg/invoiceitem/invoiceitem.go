package invoiceitem

import "time"

// Model is the representation of an invoice item
type Model struct {
	ID uint
	InvoiceHeaderID uint
	ProductID uint
	CreatedAt time.Time
	UpdatedAt time.Time
}