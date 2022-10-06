package classroom

import (
	"foro-hotel/internal/models"
)

type requestClassroom struct {
	Id          int    `json:"matricula" db:"matricula" valid:"required"`
	Name        string `json:"dni" db:"dni" valid:"required"`
	Description string `json:"username" db:"username" valid:"required"`
	Nivel       string `json:"names" db:"names" valid:"required"`
	Range       string `json:"range" db:"range" valid:"required"`
}

type responseClassrooms struct {
	Error bool                `json:"error"`
	Data  []*models.Classroom `json:"data"`
	B64   string              `json:"image"`
	Code  int                 `json:"code"`
	Type  string              `json:"type"`
	Msg   string              `json:"msg"`
}

type responseUpdate struct {
	Error bool              `json:"error"`
	Data  *models.Classroom `json:"data"`
	B64   string            `json:"image"`
	Code  int               `json:"code"`
	Type  string            `json:"type"`
	Msg   string            `json:"msg"`
}
