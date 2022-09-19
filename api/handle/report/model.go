package Report

import (
	"foro-hotel/internal/models"
)

type Model struct {
	Procedure  string            `json:procedure`
	Parameters map[string]string `json:"parameters"`
}

type requestReport struct {
	Id          int    `json:"id" db:"id" valid:"required"`
	Name        string `json:"name" db:"name" valid:"required"`
	Description string `json:"description" db:"description" valid:"required"`
	Status      int    `json:"status" db:"status" valid:"required"`
}

type responseReports struct {
	Error bool             `json:"error"`
	Data  []*models.Report `json:"data"`
	B64   string           `json:"image"`
	Code  int              `json:"code"`
	Type  string           `json:"type"`
	Msg   string           `json:"msg"`
}

type responseUpdate struct {
	Error bool                     `json:"error"`
	Data  []map[string]interface{} `json:"data"`
	B64   string                   `json:"image"`
	Code  int                      `json:"code"`
	Type  string                   `json:"type"`
	Msg   string                   `json:"msg"`
}
