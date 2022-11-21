package file

import (
	"encoding/base64"
	"fmt"
	"foro-hotel/internal/logger"
	"foro-hotel/internal/models"
	"os"
)

type PortsServerFile interface {
	CreateFile(matriculaUser int64, name string, description string, path string, fileName string, b64 string, typeFile int, status int, isDelete int) (*File, int, error)
	UpdateFile(id int, matriculaUser int64, name string, description string, path string, fileName string, typeFile int, status int, isDelete int) (*File, int, error)
	DeleteFile(id int) (int, error)
	GetFileByID(id int) (*File, int, error)
	GetAllFile() ([]*File, error)
}

type service struct {
	repository ServicesFileRepository
	user       *models.User
	txID       string
}

func NewFileService(repository ServicesFileRepository, user *models.User, TxID string) PortsServerFile {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateFile(matriculaUser int64, name string, description string, path string, fileName string, b64 string, typeFile int, status int, isDelete int) (*File, int, error) {
	m := NewCreateFile(matriculaUser, name, description, path, fileName, typeFile, status, isDelete)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	pathFile := fmt.Sprintf("%s%s", path, fileName)
	resp, err := UploadFile(pathFile, b64)
	if err != nil || !resp {
		logger.Error.Println(s.txID, " - couldn't create File :", err)
		return m, 3, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create File :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateFile(id int, matriculaUser int64, name string, description string, path string, fileName string, typeFile int, status int, isDelete int) (*File, int, error) {
	m := NewFile(id, matriculaUser, name, description, path, fileName, typeFile, status, isDelete)
	if id == 0 {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return m, 15, fmt.Errorf("id is required")
	}
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update File :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteFile(id int) (int, error) {
	if id == 0 {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return 15, fmt.Errorf("id is required")
	}

	if err := s.repository.delete(id); err != nil {
		if err.Error() == "ecatch:108" {
			return 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't update row:", err)
		return 20, err
	}
	return 28, nil
}

func (s *service) GetFileByID(id int) (*File, int, error) {
	if id == 0 {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return nil, 15, fmt.Errorf("id is required")
	}
	m, err := s.repository.getByID(id)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getByID row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}

func (s *service) GetAllFile() ([]*File, error) {
	return s.repository.getAll()
}

func UploadFile(filename string, b64 string) (bool, error) {
	dec, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return false, err
	}

	f, err := os.Create(filename)
	if err != nil {
		return false, err
	}
	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		return false, err
	}
	if err := f.Sync(); err != nil {
		return false, err
	}
	return true, nil
}
