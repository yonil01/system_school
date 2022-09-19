package models

type Report struct {
	Procedure  string            `json:procedure`
	Parameters map[string]string `json:"parameters"`
}
