package main

import (
	"log"

	"github.com/CristFlores/go-db/pkg/product"
	"github.com/CristFlores/go-db/storage"
)

func main() {
	storage.NewPostgresDB()

	// * Migracion de la tabla products
	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	if err := serviceProduct.Migrate(); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}
}