package materia

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Materia  Model struct Materia
type Materia struct {
	ID          int       `json:"id" db:"id" valid:"-"`
	Name        string    `json:"name" db:"name" valid:"required"`
	Description string    `json:"description" db:"description" valid:"required"`
	Status      int       `json:"status" db:"status" valid:"required"`
	IsDelete    int       `json:"is_delete" db:"is_delete" valid:"required"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

func NewMateria(id int, name string, description string, status int, isDelete int) *Materia {
	return &Materia{
		ID:          id,
		Name:        name,
		Description: description,
		Status:      status,
		IsDelete:    isDelete,
	}
}

func NewCreateMateria(name string, description string, status int, isDelete int) *Materia {
	return &Materia{
		Name:        name,
		Description: description,
		Status:      status,
		IsDelete:    isDelete,
	}
}

func (m *Materia) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
