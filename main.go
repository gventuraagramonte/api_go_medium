package main

import (
	"api_go_medium/models"
	"api_go_medium/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error al cargar .env")
	}
	models.ConectarDB() //Conectamos a la base!

	routes.CargarRutas()
	fmt.Println("🚀 API modular corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
