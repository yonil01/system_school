package doctype

import "foro-hotel/internal/models"

type requestDoctype struct {
	UserId int `json:"user_id" valid:"required"`
}

type responseRegisterUser struct {
	Error bool                  `json:"error"`
	Data  []*models.DoctypeUser `json:"data"`
	Code  int                   `json:"code"`
	Type  string                `json:"type"`
	Msg   string                `json:"msg"`
}
