package representante

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Representante  Model struct Representante
type Representante struct {
	ID                int       `json:"id" db:"id" valid:"-"`
	MatriculaUser     int64     `json:"matricula_user" db:"matricula_user" valid:"required"`
	TypeRepresentante string    `json:"type_representante" db:"type_representante" valid:"required"`
	Notification      string    `json:"notification" db:"notification" valid:"required"`
	Dni               string    `json:"dni" db:"dni" valid:"required"`
	Direction         string    `json:"direction" db:"direction" valid:"required"`
	Names             string    `json:"names" db:"names" valid:"required"`
	Lastnames         string    `json:"lastnames" db:"lastnames" valid:"required"`
	CellPhone         string    `json:"cell_phone" db:"cell_phone" valid:"required"`
	Email             string    `json:"email" db:"email" valid:"required"`
	Status            int       `json:"status" db:"status" valid:"required"`
	IsDelete          int       `json:"is_delete" db:"is_delete" valid:"-"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
}

func NewRepresentante(id int, matriculaUser int64, typeRepresentante string, notification string, dni string, direction string, names string, lastnames string, cellPhone string, email string, status int, isDelete int) *Representante {
	return &Representante{
		ID:                id,
		MatriculaUser:     matriculaUser,
		TypeRepresentante: typeRepresentante,
		Notification:      notification,
		Dni:               dni,
		Direction:         direction,
		Names:             names,
		Lastnames:         lastnames,
		CellPhone:         cellPhone,
		Email:             email,
		Status:            status,
		IsDelete:          isDelete,
	}
}

func NewCreateRepresentante(matriculaUser int64, typeRepresentante string, notification string, dni string, direction string, names string, lastnames string, cellPhone string, email string, status int, isDelete int) *Representante {
	return &Representante{
		MatriculaUser:     matriculaUser,
		TypeRepresentante: typeRepresentante,
		Notification:      notification,
		Dni:               dni,
		Direction:         direction,
		Names:             names,
		Lastnames:         lastnames,
		CellPhone:         cellPhone,
		Email:             email,
		Status:            status,
		IsDelete:          isDelete,
	}
}

func (m *Representante) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
