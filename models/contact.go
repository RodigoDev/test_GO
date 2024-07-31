package models

type Contact struct {
	Id       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Telefono string `json:"telefono"`
}
