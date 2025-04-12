// Importamos los paquetes para conectarnos a una base de datos
package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConectarDB() {
	connStr := "postgresql://usuarioneon:passwordneon@hostdelbasededatosneon/nombredelabasededatosneon?sslmode=require"

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("❌ Error al abrir la conexión:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("❌ No se pudo conectar a la base de datos:", err)
	}

	fmt.Println("✅ Conexión a PostgreSQL exitosa.")
}
