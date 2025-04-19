package controllers

import (
	"api_go_medium/database"
	"api_go_medium/models"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func ListarPasajerosSQL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}
	query := r.URL.Query()
	limit := 10
	offset := 0

	if l := query.Get("limit"); l != "" {
		if v, err := strconv.Atoi(l); err == nil {
			limit = v
		}
	}
	if o := query.Get("offset"); o != "" {
		if v, err := strconv.Atoi(o); err == nil {
			offset = v
		}
	}

	// Total de pasajeros activos
	var total int
	err := database.SQLDB.QueryRow(`SELECT count(*) FROM pasajeros`).Scan(&total)
	if err != nil {
		http.Error(w, "Error al contar registros", http.StatusInternalServerError)
		return
	}

	rows, err := database.SQLDB.Query(
		`SELECT id, asiento, nombre, activo
		FROM pasajeros
		ORDER BY asiento
		LIMIT $1 OFFSET $2`, limit, offset)

	if err != nil {
		http.Error(w, "Error en consulta SQL", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	pasajeros := make([]models.Pasajero, 0)
	for rows.Next() {
		var p models.Pasajero
		if err := rows.Scan(&p.ID, &p.Asiento, &p.Nombre, &p.Activo); err != nil {
			http.Error(w, "Error al leer fila", http.StatusInternalServerError)
			return
		}
		pasajeros = append(pasajeros, p)
	}

	resp := map[string]interface{}{
		"success": true,
		"data":    pasajeros,
		"meta": map[string]interface{}{
			"total":  total,
			"limit":  limit,
			"offset": offset,
			"pages":  (total + limit - 1) / limit,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
	}

}

func BuscarPasajeroSQL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/pasajeros-sql/")
	var p models.Pasajero
	err := database.SQLDB.QueryRow(`SELECT id, asiento, nombre, activo FROM pasajeros WHERE id=$1`, id).
		Scan(&p.ID, &p.Asiento, &p.Nombre, &p.Activo)
	if err == sql.ErrNoRows {
		http.Error(w, "Pasajero no encontrado", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Error al consultar pasajero", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(p); err != nil {
		http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
	}
}

func DesactivarPasajero(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}
	id := strings.TrimPrefix(r.URL.Path, "/pasajeros-sql/desactivar/")
	_, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "ID no válido", http.StatusBadRequest)
		return
	}

	res, err := database.SQLDB.Exec(`UPDATE pasajeros SET activo=false WHERE id=$1`, id)
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
	if _, err := w.Write([]byte("Pasajero desactivado correctamente")); err != nil {
		http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
		return
	}
}

func ActivarPasajero(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/pasajeros-sql/activar/")
	_, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "ID no válido", http.StatusBadRequest)
		return
	}

	res, err := database.SQLDB.Exec(`UPDATE pasajeros SET activo=true WHERE id=$1`, id)
	if err != nil {
		http.Error(w, "Error al activar pasajero", http.StatusInternalServerError)
		return
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		http.Error(w, "Pasajero no encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("Pasajero activado correctamente")); err != nil {
		log.Printf("❌ Error al escribir respuesta: %v", err)
	}
}
