package routes

import (
	"api_go_medium/controllers"
	"net/http"
	"strings"
)

func CargarRutas() {
	http.HandleFunc("/pasajeros", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.ListarPasajeros(w, r)

		case http.MethodPost:
			controllers.CrearPasajero(w, r)

		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)

		}
	})

	http.HandleFunc("/pasajeros/", func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.HasSuffix(r.URL.Path, "/desactivar") && r.Method == http.MethodPut:
			controllers.DesactivarPasajero(w, r)

		case strings.HasSuffix(r.URL.Path, "/activar") && r.Method == http.MethodPut:
			controllers.ActivarPasajero(w, r)

		default:
			controllers.BuscarPasajero(w, r)
		}

	})
}
