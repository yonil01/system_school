package materia

import (
	"fmt"
	"foro-hotel/internal/logger"
	"foro-hotel/internal/models"
)

type PortsServerMateria interface {
	CreateMateria(name string, description string, status int, isDelete int) (*Materia, int, error)
	UpdateMateria(id int, name string, description string, status int, isDelete int) (*Materia, int, error)
	DeleteMateria(id int) (int, error)
	GetMateriaByID(id int) (*Materia, int, error)
	GetAllMateria() ([]*Materia, error)
	GetMateriaByUserId(id int) ([]*Materia, int, error)
}

type service struct {
	repository ServicesMateriaRepository
	user       *models.User
	txID       string
}

func NewMateriaService(repository ServicesMateriaRepository, user *models.User, TxID string) PortsServerMateria {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateMateria(name string, description string, status int, isDelete int) (*Materia, int, error) {
	m := NewCreateMateria(name, description, status, isDelete)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create Materia :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateMateria(id int, name string, description string, status int, isDelete int) (*Materia, int, error) {
	m := NewMateria(id, name, description, status, isDelete)
	if id == 0 {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return m, 15, fmt.Errorf("id is required")
	}
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update Materia :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteMateria(id int) (int, error) {
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

func (s *service) GetMateriaByID(id int) (*Materia, int, error) {
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

func (s *service) GetAllMateria() ([]*Materia, error) {
	return s.repository.getAll()
}

func (s *service) GetMateriaByUserId(id int) ([]*Materia, int, error) {
	//TODO implement me
	panic("implement me")
}
