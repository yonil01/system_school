package models

type City struct {
	PaisCodigo string    `json:"pais_codigo" db:"PaisCodigo" valid:"required"`
	CiudadNombre   string    `json:"ciudad_nombre" db:"CiudadNombre" valid:"required"`
}


