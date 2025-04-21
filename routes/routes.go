package routes

import (
	"api_go_medium/auth"
	"api_go_medium/controllers"
	"net/http"
)

func CargarRutas() {
	http.Handle("/docs/", http.StripPrefix("/docs/", http.FileServer(http.Dir("docs"))))
	http.Handle("/openapi.yaml", http.FileServer(http.Dir("./")))
	http.Handle("/pasajeros", auth.JWTMiddleware(http.HandlerFunc(controllers.CrearPasajero)))
	http.Handle("/pasajeros-sql", auth.JWTMiddleware(http.HandlerFunc((controllers.ListarPasajerosSQL))))
	http.HandleFunc("/pasajeros-sql/", controllers.BuscarPasajeroSQL)
	http.Handle("/pasajeros-sql/desactivar/", auth.JWTMiddleware(http.HandlerFunc(controllers.DesactivarPasajero)))
	http.Handle("/pasajeros-sql/activar/", auth.JWTMiddleware(http.HandlerFunc(controllers.ActivarPasajero)))
}
