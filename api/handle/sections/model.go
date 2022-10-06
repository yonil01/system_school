package sections

import (
	"foro-hotel/internal/models"
)

type requestSection struct {
	Id      int    `json:"matricula" db:"matricula" valid:"required"`
	Name    string `json:"dni" db:"dni" valid:"required"`
	GradoId int    `json:"grado_id" db:"grado_id" valid:"required"`
}

type responseSections struct {
	Error bool               `json:"error"`
	Data  []*models.Sections `json:"data"`
	B64   string             `json:"image"`
	Code  int                `json:"code"`
	Type  string             `json:"type"`
	Msg   string             `json:"msg"`
}

type responseUpdate struct {
	Error bool             `json:"error"`
	Data  *models.Sections `json:"data"`
	B64   string           `json:"image"`
	Code  int              `json:"code"`
	Type  string           `json:"type"`
	Msg   string           `json:"msg"`
}
