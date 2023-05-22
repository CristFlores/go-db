package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/CristFlores/go-db/pkg/product"
	_ "github.com/lib/pq"
)

var (
	db *sql.DB
	once sync.Once
)

// * Implementacion patron DAO (Data Access Object) - https://chuidiang.org/index.php?title=Patr%C3%B3n_DAO

// Driver de storage
type Driver string

// Drivers
const (
	Postgres Driver = "POSTGRES"
	MySql Driver = "MYSQL"
)

// New crea una nueva instancia de la base de datos
func New(d Driver) {
	switch d {
	case MySql:
		// newMySqlDB() --> si tuvieramos una base de datos mysql
	case Postgres:
		newPostgresDB()
	}
}

// * Modelo singleton (con esto aseguramos que solo se cree una instancia de la base de datos cuando mandamos a llamar al metodo open, independientemente de las llamadas que se hagan)

func newPostgresDB() {
	once.Do(func ()  {
		var err error
		db, err = sql.Open("postgres", "postgres://cristats:bSb9:5D9@localhost:5432/db_go_course?sslmode=disable")
    if err != nil {
        log.Fatalf("can't open db: %v", err)
			}
			if err = db.Ping(); err != nil {
				log.Fatalf("can't do ping: %v", err)
    }

		fmt.Println("conectado a postgres")
	})
}

//* Pool returns a unique instance of the database connection
func Pool() *sql.DB {
	return db
}

// Helpers para manejo de nulos (funcion no exportada)

// * Recibe un string y retorna una estructura especial del pkg sql, llamado "NullString"
func stringToNull(s string) sql.NullString {
	null := sql.NullString{String: s}
	if null.String != "" {
		null.Valid = true
	}
	return null
}

// helpers para manejo de nulos de tipo time.Time (funcion no exportada)
func timeToNull(t time.Time) sql.NullTime {
	null := sql.NullTime{Time: t}
	if !null.Time.IsZero() {
		null.Valid = true
	}
	return null
}

// DAOProduct factory of product storage
// ! Al implementar este factory, los metodos NewPsqlProduct y NewMySQLProduct los volvemos no exportados
func DAOProduct(driver Driver) (product.Storage, error) {
	switch driver {
	case Postgres:
		return newPsqlProduct(db), nil
	case MySql:
		return nil, fmt.Errorf("driver no implementado")
		// return newMySQLProduct(db), nil --> si tuvieramos una base de datos mysql
	default:
		return nil, fmt.Errorf("driver desconocido")
	}
}



// * Conexion con el motor de base de datos postgres
// "postgres://<user_name>:<pw>@localhost:<port_number>/<db_name>?sslmode=verify-full"