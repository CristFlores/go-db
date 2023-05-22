package invoiceitem

import (
	"database/sql"
	"time"
)

// Model is the representation of an invoice item
type Model struct {
	ID uint
	InvoiceHeaderID uint
	ProductID uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Models is the representation of a slice of Model
type Models []*Model

// Storage interface that must implement a db storage
type Storage interface {
	Migrate() error
	//* CreateTx() recibe un puntero a sql.Tx, un uint (que corresponde al id del invoiceheader) y un puntero al slide Model - es decir, a Models, (que son varios invoiceitems)
	CreateTx(*sql.Tx, uint, Models) error
}
// Service is the representation of a invoiceitem service
type Service struct {
	storage Storage
}

// NewService returns a new pointer of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Migrate is used for migrate products
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}