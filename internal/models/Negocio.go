package models

type Negocio struct {
	NroDocuemnt string `json:"nro_document" valid:"required"`
	TypeNegocio string `json:"type_negocio" valid:"required"`
	NameNegocio string `json:"name_negocio" valid:"required"`
	CategoryNegocio string `json:"category_negocio"`
	PhoneNegocio string `json:"phone_negocio" valid:"required"`
	MobilNegocio string `json:"mobil_negocio"`
	DireccionNegocio string `json:"direccion_negocio" valid:"required"`
	DescriptionNegocio string `json:"description_negocio"`
	Lat string `json:"lat" valid:"_"`
	Lng string `json:"lng"`
}
