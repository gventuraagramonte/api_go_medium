package controllers

import (
	"api_go_medium/database"
	"api_go_medium/models"
	"encoding/json"
	"net/http"
)

type PaginacionResponse struct {
	Data []models.Pasajero `json:"data"`
	Meta struct {
		Page  int `json:"page"`
		Limit int `json:"limit"`
		Total int `json:"total"`
	} `json:"meta"`
}

func CrearPasajero(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var nuevo models.Pasajero
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if err := database.GormDB.Create(&nuevo).Error; err != nil {
		http.Error(w, "Error al crear pasajero", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(nuevo); err != nil {
		http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
	}
}

// func BuscarPasajero(w http.ResponseWriter, r *http.Request) {
// 	partes := strings.Split(r.URL.Path, "/")
// 	if len(partes) != 3 {
// 		http.Error(w, "Ruta inválida", http.StatusBadRequest)
// 		return
// 	}

// 	asientoID, err := strconv.Atoi(partes[2])
// 	if err != nil {
// 		http.Error(w, "Asiento inválido", http.StatusBadRequest)
// 		return
// 	}

// 	var p models.Pasajero
// 	if err := models.DB.Where("asiento = ? AND activo = true", asientoID).First(&p).Error; err != nil {
// 		http.Error(w, "Pasajero no encontrado", http.StatusNotFound)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(p)

// }

// func DesactivarPasajero(w http.ResponseWriter, r *http.Request) {
// 	partes := strings.Split(r.URL.Path, "/")
// 	if len(partes) != 4 || partes[3] != "desactivar" {
// 		http.Error(w, "Ruta inválida", http.StatusBadRequest)
// 		return
// 	}

// 	asientoID, err := strconv.Atoi(partes[2])
// 	if err != nil {
// 		http.Error(w, "Asiento inválido", http.StatusBadRequest)
// 		return
// 	}

// 	if err := models.DB.Model(&models.Pasajero{}).
// 		Where("asiento = ? AND activo = true", asientoID).
// 		Update("activo", false).Error; err != nil {
// 		http.Error(w, "Error al desactivar pasajero", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, "✅ Pasajero del asiento %d fue desactivado", asientoID)
// }

// func ActivarPasajero(w http.ResponseWriter, r *http.Request) {
// 	partes := strings.Split(r.URL.Path, "/")
// 	if len(partes) != 4 || partes[3] != "activar" {
// 		http.Error(w, "Ruta inválida", http.StatusBadRequest)
// 		return
// 	}

// 	asientoID, err := strconv.Atoi(partes[2])
// 	if err != nil {
// 		http.Error(w, "Asiento inválido", http.StatusBadRequest)
// 		return
// 	}

// 	if err := models.DB.Model(&models.Pasajero{}).
// 		Where("asiento = ? AND activo = false", asientoID).
// 		Update("activo", true).Error; err != nil {
// 		http.Error(w, "No se pudo activar el pasajero", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, "🟢 Pasajero del asiento %d ha sido reactivado", asientoID)
// }
