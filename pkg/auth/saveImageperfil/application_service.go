package saveimageperfil

import (
	"foro-hotel/internal/models"
	"foro-hotel/pkg/auth/users"
	"github.com/jmoiron/sqlx"
)

type PortSaveImageNegocios interface {
	SaveImage(document_id string, images *models.Image) ( int, error)
	SaveFileImage(nro_document string, images *models.Image) (int, error)
}


type service struct {
	DB *sqlx.DB
	TxID string
}

func NewSaveImageService(db *sqlx.DB, txID string) PortSaveImageNegocios {
	return &service{DB: db, TxID: txID}
}

func (s *service) SaveImage(document_id string, images *models.Image) (int, error) {
	if document_id == "" {
		return 12, nil
	}
	 cod, err := s.SaveFileImage(document_id, images)
	 if err != nil {
		 return cod, err
	 }

	return 29, nil
}

func (s *service) SaveFileImage(num_document string, images *models.Image) (int, error) {
	repoLogin := users.FactoryStorage(s.DB, nil, s.TxID)
	srvImage := users.NewUserService(repoLogin, nil, s.TxID, s.DB)

	cod, err := srvImage.SaveIamgePerfil(num_document, images)

	if err != nil {
		return cod, err
	}

	return cod, nil
}

