package representante

import (
	"fmt"
	"foro-hotel/internal/logger"
	"foro-hotel/internal/models"
)

type PortsServerRepresentante interface {
	CreateRepresentante(matriculaUser int64, typeRepresentante string, notification string, dni string, direction string, names string, lastnames string, cellPhone string, email string, status int, isDelete int) (*Representante, int, error)
	UpdateRepresentante(id int, matriculaUser int64, typeRepresentante string, notification string, dni string, direction string, names string, lastnames string, cellPhone string, email string, status int, isDelete int) (*Representante, int, error)
	DeleteRepresentante(id int) (int, error)
	GetRepresentanteByID(id int) (*Representante, int, error)
	GetAllRepresentante() ([]*Representante, error)
	GetRepresnetanteByMatriculaUser(matriculaUser int64) (*Representante, error)
}

type service struct {
	repository ServicesRepresentanteRepository
	user       *models.User
	txID       string
}

func NewRepresentanteService(repository ServicesRepresentanteRepository, user *models.User, TxID string) PortsServerRepresentante {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateRepresentante(matriculaUser int64, typeRepresentante string, notification string, dni string, direction string, names string, lastnames string, cellPhone string, email string, status int, isDelete int) (*Representante, int, error) {
	m := NewCreateRepresentante(matriculaUser, typeRepresentante, notification, dni, direction, names, lastnames, cellPhone, email, status, isDelete)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.create(m); err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create Representante :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateRepresentante(id int, matriculaUser int64, typeRepresentante string, notification string, dni string, direction string, names string, lastnames string, cellPhone string, email string, status int, isDelete int) (*Representante, int, error) {
	m := NewRepresentante(id, matriculaUser, typeRepresentante, notification, dni, direction, names, lastnames, cellPhone, email, status, isDelete)
	if id == 0 {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return m, 15, fmt.Errorf("id is required")
	}
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update Representante :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteRepresentante(id int) (int, error) {
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

func (s *service) GetRepresentanteByID(id int) (*Representante, int, error) {
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

func (s *service) GetAllRepresentante() ([]*Representante, error) {
	return s.repository.getAll()
}

func (s *service) GetRepresnetanteByMatriculaUser(matriculaUser int64) (*Representante, error) {
	return s.repository.getByMatriculaUser(matriculaUser)
}
