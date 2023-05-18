package product

import "time"

// Model is the representation of a product
type Model struct {
	ID uint
	Name string
	Observations string
	Price float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Models is a slice of Model
type Models []*Model

type Storage interface {
	Migrate() error
	// Created(*Model) error
	// Update(*Model) error
	// GetAll() (Models, error)
	// GetByID(uint) (*Model, error)
	// Delete(uint) error
}
// Service is the representation of a product service
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