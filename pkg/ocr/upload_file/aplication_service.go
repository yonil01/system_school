package upload_file

import (
	"fmt"
	"foro-hotel/internal/logger"
	"foro-hotel/internal/models"

	"github.com/asaskevich/govalidator"
)

type PortsServerFile interface {
	UploadFile(id string) (*File, int, error)
}

type service struct {
	repository ServicesUploadFileRepository
	user       *models.User
	txID       string
}

func NewUploadFileService(repository ServicesUploadFileRepository, user *models.User, TxID string) PortsServerFile {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) UploadFile(id string) (*File, int, error) {
	m := NewAccounting(id)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if !govalidator.IsUUID(id) {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id isn't uuid"))
		return nil, 15, fmt.Errorf("id isn't uuid")
	}
	if err := s.repository.uploadFile(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create Accounting :", err)
		return m, 3, err
	}

	_ = s.saveFile("")
	return m, 29, nil
}


func (s *service) saveFile(path string) error {
	return nil
}
