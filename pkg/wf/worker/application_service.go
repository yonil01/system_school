package worker

import (
	"fmt"
	"foro-hotel/internal/logger"
	"foro-hotel/internal/models"
)

type PortsServerWorker interface {
	CreateWorker(matriculaUser int64, status int, isDelete int) (*Worker, int, error)
	UpdateWorker(id int, matriculaUser int64, status int, isDelete int) (*Worker, int, error)
	DeleteWorker(id int) (int, error)
	GetWorkerByID(id int) (*Worker, int, error)
	GetAllWorker() ([]*Worker, error)
}

type service struct {
	repository ServicesWorkerRepository
	user       *models.User
	txID       string
}

func NewWorkerService(repository ServicesWorkerRepository, user *models.User, TxID string) PortsServerWorker {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateWorker(matriculaUser int64, status int, isDelete int) (*Worker, int, error) {
	m := NewCreateWorker(matriculaUser, status, isDelete)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create Worker :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateWorker(id int, matriculaUser int64, status int, isDelete int) (*Worker, int, error) {
	m := NewWorker(id, matriculaUser, status, isDelete)
	if id == 0 {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return m, 15, fmt.Errorf("id is required")
	}
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update Worker :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteWorker(id int) (int, error) {
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

func (s *service) GetWorkerByID(id int) (*Worker, int, error) {
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

func (s *service) GetAllWorker() ([]*Worker, error) {
	return s.repository.getAll()
}
