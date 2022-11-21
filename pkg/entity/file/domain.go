package file

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// File  Model struct File
type File struct {
	ID            int       `json:"id" db:"id" valid:"-"`
	MatriculaUser int64     `json:"matricula_user" db:"matricula_user" valid:"required"`
	Name          string    `json:"name" db:"name" valid:"required"`
	Description   string    `json:"description" db:"description" valid:"required"`
	Path          string    `json:"path" db:"path" valid:"-"`
	FileName      string    `json:"file_name" db:"file_name" valid:"required"`
	TypeFile      int       `json:"type_file" db:"type_file" valid:"required"`
	Status        int       `json:"status" db:"status" valid:"required"`
	IsDelete      int       `json:"is_delete" db:"is_delete" valid:"-"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

func NewFile(id int, matriculaUser int64, name string, description string, path string, fileName string, typeFile int, status int, isDelete int) *File {
	return &File{
		ID:            id,
		MatriculaUser: matriculaUser,
		Name:          name,
		Description:   description,
		Path:          path,
		FileName:      fileName,
		TypeFile:      typeFile,
		Status:        status,
		IsDelete:      isDelete,
	}
}

func NewCreateFile(matriculaUser int64, name string, description string, path string, fileName string, typeFile int, status int, isDelete int) *File {
	return &File{
		MatriculaUser: matriculaUser,
		Name:          name,
		Description:   description,
		Path:          path,
		FileName:      fileName,
		TypeFile:      typeFile,
		Status:        status,
		IsDelete:      isDelete,
	}
}

func (m *File) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
