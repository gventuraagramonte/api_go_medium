package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var GormDB *gorm.DB

func ConectarGORM() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var err error
	GormDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Error al conectar con GORM:", err)
	}

	// Verificamos la conexión real
	sqlDB, err := GormDB.DB()
	if err != nil {
		log.Fatal("❌ Error al obtener conexión subyacente:", err)
	}

	if err = sqlDB.Ping(); err != nil {
		log.Fatal("❌ No se puede hacer ping a la base de datos con GORM:", err)
	}
	fmt.Println("✅ Conexión exitosa a PostgreSQL con GORM")

}
