package models

type Image struct {
	B64 string `json:"b64" db:"b64" valid:"_"`
	Name string `json:"name" db:"name" valid:"required"`
	Type string `json:"type" db:"type" valid:"_"`
	Url string `json:"url" db:"url" valid:"_"`
}

type ImageSend struct {
	DocumentId string `json:"document_id" db:"document_id" valid:"required"`
	Name string `json:"name" db:"name" valid:"required"`
	Url string `json:"url" db:"url" valid:"_"`
	B64 string `json:"b64" db:"b64" valid:"_"`
}


