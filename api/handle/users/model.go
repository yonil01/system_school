package user

import (
	"foro-hotel/internal/models"
	"time"
)

type requestUser struct {
	Matricula     int64     `json:"matricula" db:"matricula" valid:"required"`
	Dni           string    `json:"dni" db:"dni" valid:"required"`
	Username      string    `json:"username" db:"username" valid:"required"`
	Names         string    `json:"names" db:"names" valid:"required"`
	Lastnames     string    `json:"lastnames" db:"lastnames" valid:"required"`
	Sexo          string    `json:"sexo" db:"sexo" valid:"required"`
	DateAdmission time.Time `json:"date_admission" db:"date_admission" valid:"required"`
	DateBirth     time.Time `json:"date_birth" db:"date_birth" valid:"required"`
	Email         string    `json:"email" db:"email" valid:"required"`
}

type responseUser struct {
	Error bool           `json:"error"`
	Data  []*models.User `json:"data"`
	B64   string         `json:"image"`
	Code  int            `json:"code"`
	Type  string         `json:"type"`
	Msg   string         `json:"msg"`
}

type responseUpdate struct {
	Error bool         `json:"error"`
	Data  *models.User `json:"data"`
	B64   string       `json:"image"`
	Code  int          `json:"code"`
	Type  string       `json:"type"`
	Msg   string       `json:"msg"`
}
