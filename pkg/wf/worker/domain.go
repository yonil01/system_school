package worker

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Worker  Model struct Worker
type Worker struct {
	ID            int       `json:"id" db:"id" valid:"-"`
	MatriculaUser int64     `json:"matricula_user" db:"matricula_user" valid:"required"`
	Status        int       `json:"status" db:"status" valid:"required"`
	IsDelete      int       `json:"is_delete" db:"is_delete" valid:"-"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

func NewWorker(id int, matriculaUser int64, status int, isDelete int) *Worker {
	return &Worker{
		ID:            id,
		MatriculaUser: matriculaUser,
		Status:        status,
		IsDelete:      isDelete,
	}
}

func NewCreateWorker(matriculaUser int64, status int, isDelete int) *Worker {
	return &Worker{
		MatriculaUser: matriculaUser,
		Status:        status,
		IsDelete:      isDelete,
	}
}

func (m *Worker) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
