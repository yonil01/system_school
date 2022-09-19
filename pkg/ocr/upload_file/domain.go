package upload_file

import (

	"github.com/asaskevich/govalidator"
)

// Model estructura de Accounting
type File struct {
	ID        string    `json:"id" db:"id" valid:"-"`
}

func NewAccounting(id string) *File {
	return &File{
		ID:       id,
	}
}


func (m *File) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
