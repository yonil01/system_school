package Subject

import (
	"foro-hotel/internal/models"
)

type requestSubject struct {
	Id          int    `json:"id" db:"id" valid:"required"`
	Name        string `json:"name" db:"name" valid:"required"`
	Description string `json:"description" db:"description" valid:"required"`
	Status      int    `json:"status" db:"status" valid:"required"`
}

type responseSubjects struct {
	Error bool              `json:"error"`
	Data  []*models.Subject `json:"data"`
	B64   string            `json:"image"`
	Code  int               `json:"code"`
	Type  string            `json:"type"`
	Msg   string            `json:"msg"`
}

type responseUpdate struct {
	Error bool            `json:"error"`
	Data  *models.Subject `json:"data"`
	B64   string          `json:"image"`
	Code  int             `json:"code"`
	Type  string          `json:"type"`
	Msg   string          `json:"msg"`
}
