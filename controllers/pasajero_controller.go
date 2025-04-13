package controllers

import (
	"api_go_medium/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type PaginacionResponse struct {
	Data []models.Pasajero `json:"data"`
	Meta struct {
		Page  int `json:"page"`
		Limit int `json:"limit"`
		Total int `json:"total"`
	} `json:"meta"`
}

func ListarPasajeros(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	//Valores por defecto
	page := 1
	limit := 10

	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = p
	}

	if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
		limit = l
	}

	offset := (page - 1) * limit

	var pasajeros []models.Pasajero
	if err := models.DB.
		Where("activo=?", true).
		Order("asiento").
		Limit(limit).
		Offset(offset).
		Find(&pasajeros).Error; err != nil {
		http.Error(w, "Error al consultar pasajeros", http.StatusInternalServerError)
		return
	}

	// Contar total de pasajeros activos
	var total int64
	if err := models.DB.Model(&models.Pasajero{}).Where("activo = true").Count(&total).Error; err != nil {
		http.Error(w, "Error al contar pasajeros", http.StatusInternalServerError)
		return
	}

	// Armar respuesta
	resp := PaginacionResponse{
		Data: pasajeros,
	}
	resp.Meta.Page = page
	resp.Meta.Limit = limit
	resp.Meta.Total = int(total)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}

func CrearPasajero(w http.ResponseWriter, r *http.Request) {
	var nuevo models.Pasajero
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if err := models.DB.Create(&nuevo).Error; err != nil {
		http.Error(w, "Error al crear pasajero", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nuevo)
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
	if err := models.DB.Where("asiento = ? AND activo = true", asientoID).First(&p).Error; err != nil {
		http.Error(w, "Pasajero no encontrado", http.StatusNotFound)
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

	if err := models.DB.Model(&models.Pasajero{}).
		Where("asiento = ? AND activo = true", asientoID).
		Update("activo", false).Error; err != nil {
		http.Error(w, "Error al desactivar pasajero", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "✅ Pasajero del asiento %d fue desactivado", asientoID)
}

func ActivarPasajero(w http.ResponseWriter, r *http.Request) {
	partes := strings.Split(r.URL.Path, "/")
	if len(partes) != 4 || partes[3] != "activar" {
		http.Error(w, "Ruta inválida", http.StatusBadRequest)
		return
	}

	asientoID, err := strconv.Atoi(partes[2])
	if err != nil {
		http.Error(w, "Asiento inválido", http.StatusBadRequest)
		return
	}

	if err := models.DB.Model(&models.Pasajero{}).
		Where("asiento = ? AND activo = false", asientoID).
		Update("activo", true).Error; err != nil {
		http.Error(w, "No se pudo activar el pasajero", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "🟢 Pasajero del asiento %d ha sido reactivado", asientoID)
}
