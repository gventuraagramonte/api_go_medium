// Importamos los paquetes para conectarnos a una base de datos
package models

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConectarDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ No se pudo conectar a la base de datos:", err)
	}

	err = DB.AutoMigrate(&Pasajero{})
	if err != nil {
		log.Fatal("❌ Error al migrar modelo Pasajero", err)
	}

	fmt.Println("✅ Conexión exitosa a PostgreSQL con GORM.")
}
