package main

import (
	"api_go_medium/models"
	"api_go_medium/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	fmt.Println("🌍 Modo de ejecución:", env)

	if env != "production" {
		fmt.Println("📂 Intentando cargar archivo .env...")
		if err := godotenv.Load(); err != nil {
			if os.IsNotExist(err) {
				log.Println("❌ No se encontró el archivo .env:", err)
			} else {
				log.Printf("⚠️ Error al cargar el archivo .env: %v\n", err)
			}

			if _, statErr := os.Stat(".env"); statErr != nil {
				log.Printf("⚠️  Verificación: el archivo .env realmente NO existe (%v)\n", statErr)
			} else {
				log.Println("✅ Verificación: el archivo .env sí existe, pero hay otro problema de formato o permisos.")
			}
		} else {
			log.Println("✅ Archivo .env cargado correctamente")
		}
	}

	log.Printf("📌 DB_HOST=%s", os.Getenv("DB_HOST"))
	log.Printf("📌 DB_USER=%s", os.Getenv("DB_USER"))
	log.Printf("📌 DB_NAME=%s", os.Getenv("DB_NAME"))

	log.Println("🔌 Conectando a la base de datos...")
	models.ConectarDB()

	routes.CargarRutas()
	fmt.Println("🚀 API modular corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
