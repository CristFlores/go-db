package invoiceheader

import "time"

// Model is the representation of an invoice header
type Model struct {
	ID uint
	Client string
	CreatedAt time.Time
	UpdatedAt time.Time
}