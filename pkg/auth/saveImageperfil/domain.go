package saveimageperfil

type Image struct {
	DocumentId string    `json:"document_id" db:"document_id" valid:"required"`
	Base64   []base64    `json:"images" db:"images" valid:"required"`
}

type base64 struct {
	B64 string `json:"b64"  db:"b64" valid:"required"`
	Name string `json:"name"  db:"name" valid:"required"`
	Type string `json:"type"  db:"type" valid:"required"`
	Url string `json:"url"  db:"url" valid:"_"`
}


func NewImage(DocumentId string, Base64 []base64) *Image {
	return &Image{
		DocumentId: DocumentId,
		Base64: Base64,
	}
}

func NewBase64(B64 string, Name string, Type string, Url string) *base64 {
	return &base64{
		B64: B64,
		Name: Name,
		Type: Type,
		Url: Url,
	}
}