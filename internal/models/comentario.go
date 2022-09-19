package models

type Comentario struct {
	Id string `json:"id" db:"id" valid:"_"`
	NameUser string `json:"name_user" db:"name_user" valid:"required"`
	NegocioId string `json:"negocio_id" db:"negocio_id" valid:"required"`
	Puntuacion int64 `json:"puntuacion" db:"puntuacion" valid:"required"`
	Text string `json:"text" db:"text" valid:"required"`
	CreateAt string `json:"created_at" db:"created_at" valid:"_"`
	NameNegocio string `json:"name_negocio" db:"name_negocio" valid:"_"`
}


