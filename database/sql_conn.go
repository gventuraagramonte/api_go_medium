// Importamos los paquetes para conectarnos a una base de datos
package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var SQLDB *sql.DB

func ConectarSQL() {
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var err error
	SQLDB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("❌ No se pudo conectar a la base de datos:", err)
	}

	if err = SQLDB.Ping(); err != nil {
		log.Fatal("❌ No se puede hacer ping a la base de datos:", err)
	}

	fmt.Println("✅ Conexión exitosa a PostgreSQL con sql.DB")
}
