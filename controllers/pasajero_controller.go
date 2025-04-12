package controllers

import (
	"api_go_medium/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func ListarPasajeros(w http.ResponseWriter, r *http.Request) {
	rows, err := models.DB.Query("SELECT asiento, nombre, activo FROM pasajeros WHERE activo = true ORDER BY asiento")
	if err != nil {
		http.Error(w, "Error al consultar pasajeros", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var pasajeros []models.Pasajero

	for rows.Next() {
		var p models.Pasajero
		err := rows.Scan(&p.Asiento, &p.Nombre, &p.Activo)
		if err != nil {
			http.Error(w, "Error al leer datos", http.StatusInternalServerError)
			return
		}
		pasajeros = append(pasajeros, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pasajeros)

}

func CrearPasajero(w http.ResponseWriter, r *http.Request) {
	var nuevo models.Pasajero
	err := json.NewDecoder(r.Body).Decode(&nuevo)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	_, err = models.DB.Exec("INSERT INTO pasajeros (asiento, nombre) VALUES ($1, $2)", nuevo.Asiento, nuevo.Nombre)
	if err != nil {
		http.Error(w, "Error al insertar pasajero (¿asiento duplicado?)", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Pasajero agregado: %s (Asiento %d)", nuevo.Nombre, nuevo.Asiento)
}

func BuscarPasajero(w http.ResponseWriter, r *http.Request) {
	partes := strings.Split(r.URL.Path, "/")
	if len(partes) != 3 {
		http.Error(w, "Ruta inválida", http.StatusBadRequest)
		return
	}

	asientoID, err := strconv.Atoi(partes[2])
	if err != nil {
		http.Error(w, "Asiento inválido", http.StatusBadRequest)
		return
	}

	var p models.Pasajero
	err = models.DB.QueryRow("SELECT asiento, nombre FROM pasajeros WHERE asiento = $1", asientoID).Scan(&p.Asiento, &p.Nombre)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Pasajero no encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, "Error al consultar la base de datos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)

}

func DesactivarPasajero(w http.ResponseWriter, r *http.Request) {
	partes := strings.Split(r.URL.Path, "/")
	if len(partes) != 4 || partes[3] != "desactivar" {
		http.Error(w, "Ruta inválida", http.StatusBadRequest)
		return
	}

	asientoID, err := strconv.Atoi(partes[2])
	if err != nil {
		http.Error(w, "Asiento inválido", http.StatusBadRequest)
		return
	}

	res, err := models.DB.Exec("UPDATE pasajeros SET activo = false WHERE asiento = $1", asientoID)
	if err != nil {
		http.Error(w, "Error al desactivar pasajero", http.StatusInternalServerError)
		return
	}

	count, _ := res.RowsAffected()
	if count == 0 {
		http.Error(w, "Pasajero no encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "✅ Pasajero del asiento %d fue desactivado", asientoID)
}
