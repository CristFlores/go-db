package storage

import (
	"database/sql"
	"fmt"

	"github.com/CristFlores/go-db/pkg/product"
)

// ? Firma del metodo rows.Scan() --> Como lo supe? como en typescript, haces hover sobre la funcion deseada, ctrl + click y te lleva a la defincion de la funcion
type scanner interface {
	Scan(dest ...interface{}) error
}

const (
	psqlMigrateProduct = `CREATE TABLE IF NOT EXISTS products(
		id SERIAL NOT NULL,
		name VARCHAR(25) NOT NULL,
		observations VARCHAR(100),
		price INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT products_id_pk PRIMARY KEY (id)
	)`
	psqlCreateProduct = `INSERT INTO products(name, observations, price, created_at) 
												VALUES($1, $2, $3, $4) RETURNING id`
	psqlGetAllProducts = `SELECT id, name, observations, price, created_at, updated_at 
												FROM products`
	psqlGetProductByID = psqlGetAllProducts + " WHERE id = $1"

	psqlUpdateProduct = `UPDATE products SET name = $1, observations = $2, price = $3, updated_at = $4 WHERE id = $5`

	psqlDeleteProduct = `DELETE FROM products WHERE id = $1`
)

// PsqlProduct used for work with postgres - product
type psqlProduct struct {
	db *sql.DB
}

// NewPsqlProduct return a new pointer of PsqlProduct (builder function)
func newPsqlProduct(db *sql.DB) *psqlProduct {
	return &psqlProduct{db}
}

// Migrate implements the interface product.Storage
func (p *psqlProduct) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("product migration executed successfully")
	return nil
}

// Implementacion del metodo Created de la interface product.Storage
func (p *psqlProduct) Create(m *product.Model) error {
	// Preparamos la consulta declarada en las constantes de arriba para insertar datos
	stmt, err := p.db.Prepare(psqlCreateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()
	
	// * Nota que en el siguiente query le estamos indicando que le vamos a pasar, la propiedad "Observations", en la primer insert dejaremos ese campo vacio para observar como es el manejo de datos vacios cuando no especificamos algun tratamiento en particular.
	err = stmt.QueryRow(
		m.Name, 
		// m.Observations, // Ejecucion sin el helper creado en storage.go para manejo de nulos
		stringToNull(m.Observations),
		m.Price, 
		m.CreatedAt,
	).Scan(&m.ID)
	if err != nil {
		return err
	}
	fmt.Println("the product was created successfully")
		// Se retorna nil, porque llegado a este punto la ejecucion del codigo se realizo de forma adecuada, (recuerda que este metodo retorna un error, sin embargo, llegado a este punto, indica que no hubo error)
	return nil
}


// GetAll implements the interface product.Storage
func (p *psqlProduct) GetAll() (product.Models, error) {
	// Preparamos la consulta
	stmt, err := p.db.Prepare(psqlGetAllProducts)
	if err != nil {
		return nil, err
	}
	// Cerramos los recursos al finalizar la ejecucion (stmt) del metodo
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := make(product.Models, 0)
	for rows.Next() {
		// //Mapeo de datos fila por fila
		// m := &product.Model{}

		// // Estructura intermedia para manejo de nulos de las variables Observations y UpdatedAt
		// observationNull := sql.NullString{}
		// updatedAtNull := sql.NullTime{}

		// // Escaneamos los valores de la fila en la variable "m"
		// err := rows.Scan(
		// 	&m.ID,
		// 	&m.Name,
		// 	&observationNull,
		// 	&m.Price,
		// 	&m.CreatedAt,
		// 	&updatedAtNull,
		// )
		// if err != nil {
		// 	return nil, err
		// }
		// // * Antes de agregar los elementos al slide guardamos el valor de la estructura intermedia (que puede ser null) en los campos correspondientes
		// m.Observations = observationNull.String
		// m.UpdatedAt = updatedAtNull.Time
		// // Si no hay error, agregamos la fila al slide ms
		// * NOTA: Todo lo de arriba se comenta porque ya fue implementado a traves de la interfaz interna "scanner" lo cual nos permite reutilizar este metodo 
		m, err := scanRowProduct(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	// Si finaliza el for y no hay error, entonces retornamos el slide "ms" y el error en "nil"
	return ms, nil
}

// GetByID implements the interface product.Storage
func (p *psqlProduct) GetByID(id uint) (*product.Model, error) {
	stmt, err := p.db.Prepare(psqlGetProductByID)
	if err != nil {
		return &product.Model{}, err
	}
	defer stmt.Close()

	return scanRowProduct(stmt.QueryRow(id))
}

//Implementacion del metodo Update de la interface product.Storage
func (p *psqlProduct) Update(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlUpdateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// * En este caso ignoramos el result que retorna el metodo Exec (_), porque no lo necesitamos en el caso del query que estamos ejecutando que es un UPDATE, solo nos interesa saber si hubo un error o no
	res, err := stmt.Exec(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		timeToNull(m.UpdatedAt),
		m.ID,
	)
	if err != nil {
		return err
	}
	
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	
	if rowsAffected == 0 {
		return fmt.Errorf("no existe el producto con id: %d, total de columas afectadas: %d", m.ID, rowsAffected)
	}
	
	fmt.Println("the product was updated successfully")
	return nil
}

//Implementacion del metodo Delete de la interface product.Storage
func (p *psqlProduct) Delete(id uint) error {
	stmt, err := p.db.Prepare(psqlDeleteProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no existe el producto con id: %d, total de columas afectadas: %d", id, rowsAffected)
	}

	fmt.Println("the product was deleted successfully")
	return nil
}


func scanRowProduct(s scanner) (*product.Model , error) {
	//Mapeo de datos fila por fila
	m := &product.Model{}
	// Estructura intermedia para manejo de nulos de las variables Observations y UpdatedAt
	observationNull := sql.NullString{}
	updatedAtNull := sql.NullTime{}
	// Escaneamos los valores de la fila en la variable "m"
	err := s.Scan(
		&m.ID,
		&m.Name,
		&observationNull,
		&m.Price,
		&m.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return &product.Model{}, err
	}

	// * Antes de agregar los elementos al slide guardamos el valor de la estructura intermedia (que puede ser null) en los campos correspondientes
	m.Observations = observationNull.String
	m.UpdatedAt = updatedAtNull.Time

	return m, nil
}