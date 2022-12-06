package file

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"foro-hotel/internal/logger"
	"foro-hotel/internal/models"
	"github.com/google/uuid"
	"io/ioutil"
	"os"
	"strings"
)

type PortsServerFile interface {
	CreateFile(matriculaUser int64, name string, description string, path string, fileName string, b64 string, typeFile int, status int, isDelete int) (*File, int, error)
	UpdateFile(id int, matriculaUser int64, name string, description string, path string, fileName string, typeFile int, status int, isDelete int) (*File, int, error)
	DeleteFile(id int) (int, error)
	GetFileByID(id int) (*File, int, error)
	GetAllFile() ([]*File, error)
	GetB64ByName(fullPath string) (*string, error)
}

type service struct {
	repository ServicesFileRepository
	user       *models.User
	txID       string
}

func (s *service) GetB64ByName(fullPath string) (*string, error) {
	fo, err := os.Open(fullPath)
	if err != nil {
		logger.Error.Printf("couldn't open file from document %v", err)
		return nil, err
	}
	reader := bufio.NewReader(fo)
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		logger.Error.Printf("couldn't get file. %v", err)
		return nil, err
	}
	if err := fo.Close(); err != nil {
		logger.Error.Printf("couldn't close *file. %v", err)
		return nil, err
	}
	encoded := base64.StdEncoding.EncodeToString(content)

	return &encoded, nil
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
	ext := getExtensionFile(fileName)
	fileName = fmt.Sprintf("%s.%s", uuid.New().String(), ext)
	pathFile := fmt.Sprintf("%s%s", path, fileName)
	resp, err := UploadFile(pathFile, b64)
	if err != nil || !resp {
		logger.Error.Println(s.txID, " - couldn't create File :", err)
		return m, 3, err
	}
	m.FileName = fileName

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

func getExtensionFile(path string) string {
	pathArray := strings.Split(path, ".")
	return pathArray[(len(pathArray) - 1)]
}
