package models

type Reservacion struct {
	Id string `json:"id" db:"id" valid:"_"`
	ClienteId string `json:"cliente_id" db:"cliente_id" valid:"required"`
	NegocioId string `json:"negocio_id" db:"negocio_id" valid:"required"`
	Email string `json:"email" db:"email" valid:"required"`
	Datos string `json:"datos" db:"datos" valid:"required"`
	Telefono string `json:"telefono" db:"telefono" valid:"required"`
	Celular string `json:"celular" db:"celular" valid:"required"`
	Direccion string `json:"direccion" db:"direccion" valid:"required"`
	Url string `json:"url" db:"url" valid:"required"`
	CreateAt string `json:"created_at" db:"created_at" valid:"_"`
	TextHash string `json:"text_hash" db:"text_hash" valid:"_"`
	NameNegocio string `json:"name_negocio" db:"name_negocio" valid:"_"`
}


