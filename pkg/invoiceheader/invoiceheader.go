package invoiceheader

import (
	"database/sql"
	"time"
)

// Model is the representation of an invoice header
type Model struct {
	ID uint
	Client string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Storage interface that must implement a db storage
type Storage interface {
	Migrate() error
	//* CreateTx() recibe un puntero a sql.Tx y un puntero a Model(que es un invoiceheader)
	CreateTx(*sql.Tx, *Model) error
}
// Service is the representation of a invoiceheader service
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