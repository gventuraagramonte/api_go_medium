package models

type Pasajero struct {
	ID      string `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Asiento int    `json:"asiento" gorm:"uniqueIndex"`
	Nombre  string `json:"nombre"`
	Activo  string `json:"activo" gorm:"default:true"`
}
