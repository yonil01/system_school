package Payment

import (
	"foro-hotel/internal/models"
	"time"
)

type requestPayment struct {
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
}

type responsePayments struct {
	Error bool              `json:"error"`
	Data  []*models.Payment `json:"data"`
	B64   string            `json:"image"`
	Code  int               `json:"code"`
	Type  string            `json:"type"`
	Msg   string            `json:"msg"`
}

type responseUpdate struct {
	Error bool            `json:"error"`
	Data  *models.Payment `json:"data"`
	B64   string          `json:"image"`
	Code  int             `json:"code"`
	Type  string          `json:"type"`
	Msg   string          `json:"msg"`
}
