package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
	once sync.Once
)

// * Modelo singleton (con esto aseguramos que solo se cree una instancia de la base de datos cuando mandamos a llamar al metodo open, independientemente de las llamadas que se hagan)

func NewPostgresDB() {
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

// * Conexion con el motor de base de datos postgres
// "postgres://<user_name>:<pw>@localhost:<port_number>/<db_name>?sslmode=verify-full"