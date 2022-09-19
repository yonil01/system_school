package users

import (
	"foro-hotel/internal/logger"
	"foro-hotel/internal/models"
	"foro-hotel/internal/pwd"
	"github.com/jmoiron/sqlx"
)

type PortsServerUser interface {
	CreateUser(numberDocument string, typeDocument string, firstName string, secondName string, fisrtLastname string, secondLastname string, email string, password string) (*User, int, error)
	GetUserByEmail(email string) (*User, int, error)
	InsertLogin(numberDocument string, typeDocument string, firstName string, secondName string, fisrtLastname string, secondLastname string, email string, real_ip string) (int, error)
	SaveIamgePerfil(document_id string, images *models.Image) (int, error)
	GetDoctypeUser(userId int) ([]*models.DoctypeUser, error)
}

type service struct {
	repository ServicesUserRepository
	user       *models.User
	txID       string
	db         *sqlx.DB
}

func NewUserService(repository ServicesUserRepository, user *models.User, TxID string, Db *sqlx.DB) PortsServerUser {
	return &service{repository: repository, user: user, txID: TxID, db: Db}
}

func (s *service) CreateUser(numberDocument string, typeDocument string, firstName string, secondName string, fisrtLastname string, secondLastname string, email string, password string) (*User, int, error) {
	m := NewCreateUser(numberDocument, typeDocument, firstName, secondName, fisrtLastname, secondLastname, email, password)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	m.Password = pwd.Encrypt(password)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	/*cod, err := s.repository.getUser(numberDocument, email)
	if err != nil {
		return nil, cod, err
	}

	m1, err := s.repository.create(m)
	if err != nil {
		if err.Error() == "ecatch:108" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create User :", err)
		return m1, 3, err
	}*/

	return nil, 29, nil
}

func (s *service) GetUserByEmail(email string) (*User, int, error) {
	m, cod, err := s.repository.getUserByEmail(email)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't getByNickName row:", err)
		return nil, cod, err
	}
	return m, cod, nil
}

func (s *service) InsertLogin(numberDocument string, typeDocument string, firstName string, secondName string, fisrtLastname string, secondLastname string, email string, real_ip string) (int, error) {
	m := NewLoginUser(numberDocument, typeDocument, firstName, secondName, fisrtLastname, secondLastname, email, real_ip)

	cod, err := s.repository.createLogin(m)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't getByNickName row:", err)
		return cod, err
	}
	return cod, nil
}

func (s *service) SaveIamgePerfil(document_id string, images *models.Image) (int, error) {
	cod, err := s.repository.saveImagePerfil(document_id, images)
	if err != nil {
		logger.Error.Println(s.txID, " - error in saved iamge:", err)
		return cod, err
	}
	return cod, nil
}

func (s *service) GetDoctypeUser(userId int) ([]*models.DoctypeUser, error) {
	m, err := s.repository.getDoctypeByUser(userId)
	if err != nil {
		logger.Error.Println(s.txID, " - error in saved iamge:", err)
		return m, err
	}
	return m, nil
}
