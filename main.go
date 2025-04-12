package main

import (
	"api_go_medium/models"
	"api_go_medium/routes"
	"fmt"
	"net/http"
)

func main() {
	models.ConectarDB() //Conectamos a la base!

	routes.CargarRutas()
	fmt.Println("🚀 API modular corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
