package main

import (
	"api_go_medium/database"
	"api_go_medium/routes"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
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

	log.Println("🔌 Conectando a la base de datos...")
	database.ConectarSQL()
	database.ConectarGORM()

	routes.CargarRutas()

	// Cors middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-Requested-With", "Accept", "Origin"},
		AllowCredentials: true,
	})

	handlerConCORS := c.Handler(http.DefaultServeMux)

	fmt.Println("🚀 API modular corriendo en http://localhost:8080")

	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      handlerConCORS,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("❌ Error al iniciar el servidor:", err)
	}

}
