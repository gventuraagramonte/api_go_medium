package routes

import (
	"api_go_medium/controllers"
	"net/http"
)

func CargarRutas() {
	http.Handle("/docs/", http.StripPrefix("/docs/", http.FileServer(http.Dir("docs"))))
	http.Handle("/openapi.yaml", http.FileServer(http.Dir("./")))
	http.HandleFunc("/pasajeros", controllers.CrearPasajero)
	http.HandleFunc("/pasajeros-sql", controllers.ListarPasajerosSQL)
	http.HandleFunc("/pasajeros-sql/", controllers.BuscarPasajeroSQL)
	http.HandleFunc("/pasajeros-sql/desactivar/", controllers.DesactivarPasajero)
	http.HandleFunc("/pasajeros-sql/activar/", controllers.ActivarPasajero)

	// http.HandleFunc("/pasajeros/", func(w http.ResponseWriter, r *http.Request) {
	// 	switch {
	// 	case strings.HasSuffix(r.URL.Path, "/desactivar") && r.Method == http.MethodPut:
	// 		controllers.DesactivarPasajero(w, r)

	// 	case strings.HasSuffix(r.URL.Path, "/activar") && r.Method == http.MethodPut:
	// 		controllers.ActivarPasajero(w, r)

	// 	default:
	// 		controllers.BuscarPasajero(w, r)
	// 	}

	// })
}
