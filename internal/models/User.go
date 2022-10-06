package models

import (
	"time"
)

type UserValidate struct {
	ID       int     `json:"id" db:"id" valid:"-"`
	Nickname string  `json:"nickname" db:"nickname" valid:"required,stringlength(3|15),matches(^[a-zA-Z0-9]+$)`
	Name     string  `json:"name" db:"name" valid:"required"`
	IdNumber *string `json:"id_number" db:"id_number" valid:"-"`
	Email    string  `json:"email" db:"email" valid:"email,required"`
}

type DoctypeUser struct {
	Id             int       `json:"id" db:"id" valid:"required"`
	Name           string    `json:"name" db:"name" valid:"required"`
	Description    string    `json:"description" db:"description" valid:"required"`
	Url            string    `json:"url" db:"url" valid:"required"`
	UserId         string    `json:"user_id" db:"user_id" valid:"required"`
	Type           int       `json:"type" db:"type" valid:"required"`
	FirstValue     string    `json:"first_value" db:"first_value" valid:"required"`
	SecondValue    string    `json:"second_value" db:"second_value" valid:"required"`
	Execution      string    `json:"execution" db:"execution" valid:"required"`
	ExecutionGarps string    `json:"execution_garps" db:"execution_garps" valid:"required"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

type User struct {
	Id            int       `json:"id" db:"id" valid:"required"`
	Dni           string    `json:"dni" db:"dni" valid:"required"`
	Matricula     int64     `json:"matricula" db:"matricula" valid:"required"`
	Username      string    `json:"username" db:"username" valid:"required"`
	Names         string    `json:"names" db:"names" valid:"required"`
	Lastnames     string    `json:"lastnames" db:"lastnames" valid:"required"`
	Sexo          string    `json:"sexo" db:"sexo" valid:"required"`
	Status        int       `json:"status" db:"status" valid:"required"`
	DateAdmission time.Time `json:"date_admission" db:"date_admission" valid:"required"`
	DateBirth     time.Time `json:"date_birth" db:"date_birth" valid:"required"`
	Email         string    `json:"email" db:"email" valid:"required"`
	IsDelete      int       `json:"is_delete" db:"is_delete" valid:"required"`
	Role          int       `json:"role" db:"role" valid:"required"`
	Password      string    `json:"password" db:"password" valid:"required"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

type Classroom struct {
	Id          int       `json:"id" db:"id" valid:"required"`
	Name        string    `json:"name" db:"name" valid:"required"`
	Description string    `json:"description" db:"description" valid:"required"`
	Nivel       string    `json:"nivel" db:"nivel" valid:"required"`
	Range       string    `json:"range" db:"range" valid:"required"`
	Status      int       `json:"status" db:"status" valid:"required"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type Subject struct {
	Id          int       `json:"id" db:"id" valid:"required"`
	Name        string    `json:"name" db:"name" valid:"required"`
	Description string    `json:"description" db:"description" valid:"required"`
	Status      int       `json:"status" db:"status" valid:"required"`
	IsDelete    int       `json:"is_delete" db:"is_delete" valid:"required"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type Payment struct {
	Id            int       `json:"id" db:"id" valid:"required"`
	Matricula     int64     `json:"matricula" db:"matricula" valid:"required"`
	Name          string    `json:"name" db:"name" valid:"required"`
	Description   string    `json:"description" db:"description" valid:"required"`
	Motivo        string    `json:"motivo" db:"motivo" valid:"required"`
	Status        int       `json:"status" db:"status" valid:"required"`
	DatePayment   time.Time `json:"date_payment" db:"date_payment" valid:"required"`
	UserMatricula int64     `json:"user_matricula" db:"user_matricula" valid:"required"`
	Amount        float64   `json:"amount" db:"amount" valid:"required"`
	Role          int       `json:"role" db:"role" valid:"required"`
	IsDelete      int       `json:"is_delete" db:"is_delete" valid:"required"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

type Sections struct {
	Id        int       `json:"id" db:"id" valid:"required"`
	Name      string    `json:"name" db:"name" valid:"required"`
	GradoId   int       `json:"grado_id" db:"grado_id" valid:"required"`
	Status    int       `json:"status" db:"status" valid:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
