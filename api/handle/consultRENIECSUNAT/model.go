package consultRENIECSUNAT

import "foro-hotel/internal/models"

type RequestGetUserReniec struct {
	dni string `json:"dni" valid:"required"`
}


type responseReniecSunat struct {
	Error bool         `json:"error"`
	Data  *models.User `json:"data"`
	Code  int          `json:"code"`
	Type  string       `json:"type"`
	Msg   string       `json:"msg"`
}