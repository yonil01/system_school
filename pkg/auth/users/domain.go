package users

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type User struct {
	NumberDocument string    `json:"number_document" db:"number_document" valid:"required"`
	TypeDocument   string    `json:"type_document" db:"type_document" valid:"required"`
	FirstName      string    `json:"first_name" db:"first_name" valid:"required"`
	SecondName     string    `json:"second_name" db:"second_name" valid:"-"`
	FirstLastname  string    `json:"first_lastname" db:"first_lastname" valid:"required"`
	SecondLastname string    `json:"second_lastname" db:"second_lastname" valid:"required"`
	Email          string    `json:"email" db:"email" valid:"required"`
	Password       string    `json:"password" db:"password" valid:"required"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	RealIP         string    `json:"real_ip" db:"real_ip"`
	Url            string    `json:"url" db:"url" valid:"_"`
	TypeUrl        string    `json:"type_url" db:"type_url" valid:"_"`
}

type DoctypeUser struct {
	Id          int       `json:"id" db:"id" valid:"required"`
	Name        string    `json:"name" db:"name" valid:"required"`
	Description string    `json:"description" db:"description" valid:"required"`
	Url         string    `json:"url" db:"url" valid:"required"`
	UserId      string    `json:"user_id" db:"user_id" valid:"required"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

func NewCreateUser(numberDocuemnt string, typeDocument string, firstName string, secondName string, fisrtLastname string, secondLastname string, email string, password string) *User {
	return &User{
		NumberDocument: numberDocuemnt,
		TypeDocument:   typeDocument,
		FirstName:      firstName,
		SecondName:     secondName,
		FirstLastname:  fisrtLastname,
		SecondLastname: secondLastname,
		Email:          email,
		Password:       password,
	}
}

func NewLoginUser(numberDocuemnt string, typeDocument string, firstName string, secondName string, fisrtLastname string, secondLastname string, email string, real_ip string) *User {
	return &User{
		NumberDocument: numberDocuemnt,
		TypeDocument:   typeDocument,
		FirstName:      firstName,
		SecondName:     secondName,
		FirstLastname:  fisrtLastname,
		SecondLastname: secondLastname,
		Email:          email,
		RealIP:         real_ip,
	}
}

func (m *User) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
