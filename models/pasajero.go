package models

import "github.com/google/uuid"

type Pasajero struct {
	ID      uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Asiento int       `gorm:"unique" json:"asiento"`
	Nombre  string    `json:"nombre"`
	Activo  bool      `gorm:"default:true" json:"activo"`
}
