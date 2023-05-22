package invoice

import (
	"github.com/CristFlores/go-db/pkg/invoiceheader"
	"github.com/CristFlores/go-db/pkg/invoiceitem"
)

// Model of invoice
type Model struct {
	Header *invoiceheader.Model
	Items invoiceitem.Models
}

// Storage interface that must implement a db storage
type Storage interface {
	Create(*Model) error
}

// Service is the representation of a invoice service
type Service struct {
	// * Este servicio contiene un campo llamado storage que es de tipo Storage
	storage Storage
}

// Funcion constructora de Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Metodo que nos permite crear la factura
func (s *Service) Create(m *Model) error {
	// Este servicio se comunica con el storage (sistema de almacenamiento) para crear la factura por medio del metodo Create de la interface Storage
	return s.storage.Create(m)
}