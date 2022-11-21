package user

import (
	"foro-hotel/internal/models"
	"foro-hotel/pkg/entity/file"
	"foro-hotel/pkg/entity/representante"
	"time"
)

type requestUser struct {
	Matricula     int64     `json:"matricula" db:"matricula" valid:"required"`
	Dni           string    `json:"dni" db:"dni" valid:"required"`
	Username      string    `json:"username" db:"username" valid:"required"`
	Names         string    `json:"names" db:"names" valid:"required"`
	Lastnames     string    `json:"lastnames" db:"lastnames" valid:"required"`
	Sexo          string    `json:"sexo" db:"sexo" valid:"required"`
	Role          int       `json:"role" db:"role" valid:"required"`
	DateAdmission time.Time `json:"date_admission" db:"date_admission" valid:"required"`
	DateBirth     time.Time `json:"date_birth" db:"date_birth" valid:"required"`
	Email         string    `json:"email" db:"email" valid:"required"`
}

type responseUsers struct {
	Error bool           `json:"error"`
	Data  []*models.User `json:"data"`
	B64   string         `json:"image"`
	Code  int            `json:"code"`
	Type  string         `json:"type"`
	Msg   string         `json:"msg"`
}

type responseUser struct {
	Error bool         `json:"error"`
	Data  *models.User `json:"data"`
	B64   string       `json:"image"`
	Code  int          `json:"code"`
	Type  string       `json:"type"`
	Msg   string       `json:"msg"`
}

type responseUpdate struct {
	Error bool         `json:"error"`
	Data  *models.User `json:"data"`
	B64   string       `json:"image"`
	Code  int          `json:"code"`
	Type  string       `json:"type"`
	Msg   string       `json:"msg"`
}

// TODO REPRESNETNATE
type RequestRepresentante struct {
	MatriculaUser     int64  `json:"matricula_user" db:"matricula_user" valid:"required"`
	TypeRepresentante string `json:"type_representante" db:"type_representante" valid:"required"`
	Notification      string `json:"notification" db:"notification" valid:"required"`
	Dni               string `json:"dni" db:"dni" valid:"required"`
	Direction         string `json:"direction" db:"direction" valid:"required"`
	Names             string `json:"names" db:"names" valid:"required"`
	Lastnames         string `json:"lastnames" db:"lastnames" valid:"required"`
	CellPhone         string `json:"cell_phone" db:"cell_phone" valid:"required"`
	Email             string `json:"email" db:"email" valid:"required"`
}
type responsRepresentante struct {
	Error bool                         `json:"error"`
	Data  *representante.Representante `json:"data"`
	Code  int                          `json:"code"`
	Type  string                       `json:"type"`
	Msg   string                       `json:"msg"`
}

// TODO FILE

type RequestFile struct {
	Files []File `json:"files" db:"files" valid:"required"`
}

type File struct {
	MatriculaUser int64  `json:"matricula_user" db:"matricula_user" valid:"required"`
	Name          string `json:"name" db:"name" valid:"required"`
	Description   string `json:"description" db:"description" valid:"required"`
	FileName      string `json:"file_name" db:"file_name" valid:"required"`
	B64           string `json:"b64" db:"b64" valid:"-"`
	TypeFile      int    `json:"type_file" db:"type_file" valid:"required"`
}

type responseFiles struct {
	Error bool         `json:"error"`
	Data  []*file.File `json:"data"`
	Code  int          `json:"code"`
	Type  string       `json:"type"`
	Msg   string       `json:"msg"`
}
