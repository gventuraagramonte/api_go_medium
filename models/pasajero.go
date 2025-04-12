package models

type Pasajero struct {
	Asiento int    `json:"asiento"`
	Nombre  string `json:"nombre"`
	Activo  bool   `json:"activo"`
}

// Base de datos en memoria
var Pasajeros = []Pasajero{
	{Asiento: 5, Nombre: "Giorgio"},
	{Asiento: 9, Nombre: "Luis"},
	{Asiento: 18, Nombre: "Ana"},
}
