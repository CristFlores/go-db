package main

import (
	"fmt"
	"log"

	"github.com/CristFlores/go-db/pkg/product"
	"github.com/CristFlores/go-db/storage"
)

func main() {
	// * Conexion con el motor de base de datos postgres
	driver := storage.Postgres
	storage.New(driver)
	myStorage, err := storage.DAOProduct(driver)
	if err != nil {
		log.Fatalf("DAOProduct: %v", err)
	}

	// * Migracion de la tablas products, invoiceheader e invoiceitem
	// Ver code.md

	// * Insertando un nuevo producto
	// storageProduct := storage.NewPsqlProduct(storage.Pool())
	
	// serviceProduct := product.NewService(storageProduct)
	// ! Misma instrucción que la de arriba, pero usando el factory method DAOProduct
	serviceProduct := product.NewService(myStorage)

	// m := &product.Model{
	// 	Name: "Curso de Go",
	// 	Price: 50,
	// }
	// m := &product.Model{
	// 	Name: "Curso de POO con Go",
	// 	Price: 80,
	// }
	// m := &product.Model{
	// 	Name: "Curso de db con Go",
	// 	Price: 70,
	// 	Observations: "on fire",
	// }

	// if err := serviceProduct.Create(m); err != nil {
	// 	log.Fatalf("product.Create: %v", err)
	// }

	ms, err := serviceProduct.GetAll(); 
	if err != nil {
		log.Fatalf("product.GetAll: %v", err)
	}
	fmt.Println(ms)

	// m, err := serviceProduct.GetByID(1)
	// switch {
	// case errors.Is(err, sql.ErrNoRows):
	// 	fmt.Println("No hay un producto con este id")
	// case err != nil:
	// 	log.Fatalf("product.GetByID: %v", err)
	// default:
	// 	fmt.Println(m)
	// }
	
	// m := &product.Model{
	// 	ID: 56,
	// 	Name: "Curso de Python",
	// 	Observations: "este es un curso de python",
	// 	Price: 100,
	// }
	// err := serviceProduct.Update(m); 
	// if err != nil {
	// 	log.Fatalf("product.Update: %v", err)
	// }

	// err := serviceProduct.Delete(56)
	// if err != nil {
	// 	log.Fatalf("product.Delete: %v", err)
	// }

	// * Insertando una nueva factura
	
	// storageHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	// storageItems := storage.NewPsqlInvoiceItem(storage.Pool())
	// storageInvoice := storage.NewPsqlInvoice(
	// 	storage.Pool(), 
	// 	storageHeader, 
	// 	storageItems,
	// )

	// m := &invoice.Model{
	// 	Header: &invoiceheader.Model{
	// 		Client: "Otro",
	// 	},
	// 	Items: invoiceitem.Models{
	// 		&invoiceitem.Model{ProductID: 8},
	// 		&invoiceitem.Model{ProductID: 3},
	// 	},
	// }

	// serviceInvoice := invoice.NewService(storageInvoice)
	// if err := serviceInvoice.Create(m); err != nil {
	// 	log.Fatalf("invoice.Create: %v", err)
	// }
}