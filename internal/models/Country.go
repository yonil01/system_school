package models

type Country struct {
	PaisCodigo string    `json:"pais_codigo" db:"PaisCodigo" valid:"required"`
	PaisNombre   string    `json:"pais_nombre" db:"PaisNombre" valid:"required"`
}


