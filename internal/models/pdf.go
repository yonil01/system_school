package models

type Pdf struct {
	DocumentId string `json:"document_id" db:"document_id" valid:"required"`
	FirstName string `json:"first_name" db:"first_name" valid:"required"`
	SecondName string `json:"second_name" db:"second_name" valid:"_"`
	FisrtLastname string `json:"fisrt_lastname" db:"fisrt_lastname" valid:"required"`
	SecondLastname string `json:"second_lastname" db:"second_lastname" valid:"required"`
	Email string `json:"email" db:"email" valid:"required"`
	Direction string `json:"direction" db:"direction" valid:"required"`
	Phone string `json:"phone" db:"phone" valid:"required"`
	Mobile string `json:"mobile" db:"mobile" valid:"_"`
	TypeService string `json:"type_service" db:"type_service" valid:"_"`
	NameService string `json:"name_service" db:"name_service" valid:"_"`
	DirectionService string `json:"direction_service" db:"direction_service" valid:"_"`
	PhoneService string `json:"phone_service" db:"phone_service" valid:"_"`
}


