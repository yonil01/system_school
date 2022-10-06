package profile

import (
	"foro-hotel/internal/logger"
	"foro-hotel/internal/models"
	"github.com/jmoiron/sqlx"
	"time"
)

type PortServiceProfile interface {
	GetUser(role int) ([]*models.User, error)
	UpdateUser(matricula int64, dni string, username string, names string, lastname string, sexo string, email string, dateBirth time.Time, dateAdmission time.Time) (*models.User, error)
	CreateUser(dni string, username string, names string, lastname string, sexo string, email string, dateBirth time.Time, dateAdmission time.Time, role int) (*models.User, error)
	DeleteUser(matricula int64, dni string, username string, names string, lastname string, sexo string, email string, dateBirth time.Time, dateAdmission time.Time) (*models.User, error)
	GetClassrooms() ([]*models.Classroom, error)
	UpdateClassroom(id int, name string, description string, nivel string, _range string) (*models.Classroom, error)
	CreateClassroom(name string, description string, nivel string, _range string) (*models.Classroom, error)
	DeleteClassroom(id int, name string, description string, nivel string, _range string) (*models.Classroom, error)
	GetSubjects() ([]*models.Subject, error)
	UpdateSubject(id int, name string, description string, status int) (*models.Subject, error)
	CreateSubject(name string, description string) (*models.Subject, error)
	DeleteSubject(id int, name string, description string) (*models.Subject, error)
	GetPayments() ([]*models.Payment, error)
	UpdatePayment(id int, name string, description string, motivo string, status int, datePayment time.Time, userMatricula int64, amount float64, role int) (*models.Payment, error)
	CreatePayment(name string, description string, motivo string, status int, datePayment time.Time, userMatricula int64, amount float64, role int) (*models.Payment, error)
	DeletePayment(id int, name string, description string, motivo string, status int, datePayment time.Time, userMatricula int64, amount float64, role int) (*models.Payment, error)
	ExecuteSP(procedure string, parameters map[string]string) ([]map[string]interface{}, error)
	GetSections() ([]*models.Sections, error)
	UpdateSection(id int, name string, gradoId int) (*models.Sections, error)
	CreateSection(name string, gradoId int) (*models.Sections, error)
	DeleteSection(id int, name string, gradoId int) (*models.Sections, error)
}

type service struct {
	repository ServicesDataRepository
	user       *models.User
	txID       string
	db         *sqlx.DB
}

func NewProfileService(repository ServicesDataRepository, user *models.User, TxID string, Db *sqlx.DB) PortServiceProfile {
	return &service{repository: repository, user: user, txID: TxID, db: Db}
}

func (s *service) GetUser(role int) ([]*models.User, error) {
	m, err := s.repository.getUser(role)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't getByNickName row:", err)
		return nil, err
	}
	return m, nil
}

func (s *service) UpdateUser(matricula int64, dni string, username string, names string, lastname string, sexo string, email string, dateBirth time.Time, dateAdmission time.Time) (*models.User, error) {
	student := models.User{
		Matricula:     matricula,
		Dni:           dni,
		Username:      username,
		Names:         names,
		Lastnames:     lastname,
		Sexo:          sexo,
		Email:         email,
		DateAdmission: dateAdmission,
		DateBirth:     dateBirth,
	}
	m, err := s.repository.updateUser(student)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't getByNickName row:", err)
		return nil, err
	}
	return m, nil
}

func (s *service) CreateUser(dni string, username string, names string, lastname string, sexo string, email string, dateBirth time.Time, dateAdmission time.Time, role int) (*models.User, error) {
	student := models.User{
		Dni:           dni,
		Username:      username,
		Names:         names,
		Lastnames:     lastname,
		Sexo:          sexo,
		Email:         email,
		Role:          role,
		DateAdmission: dateAdmission,
		DateBirth:     dateBirth,
		Status:        1,
		IsDelete:      0,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	m, err := s.repository.insertUser(student)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't getByNickName row:", err)
		return nil, err
	}
	return m, nil
}

func (s *service) DeleteUser(matricula int64, dni string, username string, names string, lastname string, sexo string, email string, dateBirth time.Time, dateAdmission time.Time) (*models.User, error) {
	student := models.User{
		Matricula:     matricula,
		Dni:           dni,
		Username:      username,
		Names:         names,
		Lastnames:     lastname,
		Sexo:          sexo,
		Email:         email,
		DateAdmission: dateAdmission,
		DateBirth:     dateBirth,
	}
	m, err := s.repository.deleteUser(student)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't getByNickName row:", err)
		return nil, err
	}
	return m, nil
}

func (s *service) GetClassrooms() ([]*models.Classroom, error) {
	m, err := s.repository.getClassrooms()
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't getByNickName row:", err)
		return nil, err
	}
	return m, nil
}

func (s *service) UpdateClassroom(id int, name string, description string, nivel string, _range string) (*models.Classroom, error) {
	classroom := models.Classroom{
		Id:          id,
		Name:        name,
		Description: description,
		Nivel:       nivel,
		Range:       _range,
		UpdatedAt:   time.Now(),
	}

	m, err := s.repository.updateClassroom(classroom)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't getByNickName row:", err)
		return nil, err
	}
	return m, nil
}

func (s *service) CreateClassroom(name string, description string, nivel string, _range string) (*models.Classroom, error) {
	student := models.Classroom{
		Name:        name,
		Description: description,
		Nivel:       nivel,
		Range:       _range,
		Status:      1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	m, err := s.repository.insertClassroom(student)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't getByNickName row:", err)
		return nil, err
	}
	return m, nil
}

func (s *service) DeleteClassroom(id int, name string, description string, nivel string, _range string) (*models.Classroom, error) {
	classroom := models.Classroom{
		Id:          id,
		Name:        name,
		Description: description,
		Nivel:       nivel,
		Range:       _range,
	}
	m, err := s.repository.deleteClassroom(classroom)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't getByNickName row:", err)
		return nil, err
	}
	return m, nil
}

func (s *service) GetSubjects() ([]*models.Subject, error) {
	m, err := s.repository.getSubjects()
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't getByNickName row:", err)
		return nil, err
	}
	return m, nil
}

func (s *service) UpdateSubject(id int, name string, description string, status int) (*models.Subject, error) {
	classroom := models.Subject{
		Id:          id,
		Name:        name,
		Description: description,
		Status:      status,
		UpdatedAt:   time.Now(),
	}

	m, err := s.repository.updateSubject(classroom)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't getByNickName row:", err)
		return nil, err
	}
	return m, nil
}

func (s *service) CreateSubject(name string, description string) (*models.Subject, error) {
	student := models.Subject{
		Name:        name,
		Description: description,
		Status:      1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	m, err := s.repository.insertSubject(student)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't getByNickName row:", err)
		return nil, err
	}
	return m, nil
}

func (s *service) DeleteSubject(id int, name string, description string) (*models.Subject, error) {
	classroom := models.Subject{
		Id:          id,
		Name:        name,
		Description: description,
	}
	m, err := s.repository.deleteSubject(classroom)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't getByNickName row:", err)
		return nil, err
	}
	return m, nil
}

func (s *service) GetPayments() ([]*models.Payment, error) {
	m, err := s.repository.getPayments()
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't getByNickName row:", err)
		return nil, err
	}
	return m, nil
}

func (s *service) UpdatePayment(id int, name string, description string, motivo string, status int, datePayment time.Time, userMatricula int64, amount float64, role int) (*models.Payment, error) {
	student := models.Payment{
		Id:            id,
		Name:          name,
		Description:   description,
		Motivo:        motivo,
		Status:        status,
		DatePayment:   datePayment,
		UserMatricula: userMatricula,
		Amount:        amount,
		Role:          role,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	m, err := s.repository.updatePayment(student)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't getByNickName row:", err)
		return nil, err
	}
	return m, nil
}

func (s *service) CreatePayment(name string, description string, motivo string, status int, datePayment time.Time, userMatricula int64, amount float64, role int) (*models.Payment, error) {
	student := models.Payment{
		Name:          name,
		Description:   description,
		Motivo:        motivo,
		Status:        status,
		DatePayment:   datePayment,
		UserMatricula: userMatricula,
		Amount:        amount,
		Role:          role,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	m, err := s.repository.insertPayment(student)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't getByNickName row:", err)
		return nil, err
	}
	return m, nil
}

func (s *service) DeletePayment(id int, name string, description string, motivo string, status int, datePayment time.Time, userMatricula int64, amount float64, role int) (*models.Payment, error) {
	student := models.Payment{
		Id:            id,
		Name:          name,
		Description:   description,
		Motivo:        motivo,
		Status:        status,
		DatePayment:   datePayment,
		UserMatricula: userMatricula,
		Amount:        amount,
		Role:          role,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	m, err := s.repository.deletePayment(student)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't getByNickName row:", err)
		return nil, err
	}
	return m, nil
}

func (s *service) ExecuteSP(procedure string, parameters map[string]string) ([]map[string]interface{}, error) {
	report := &models.Report{
		Procedure:  procedure,
		Parameters: parameters,
	}
	m, err := s.repository.ExecuteSP(report)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't getByNickName row:", err)
		return nil, err
	}
	return m, nil
}

func (s *service) GetSections() ([]*models.Sections, error) {
	m, err := s.repository.getSections()
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't getByNickName row:", err)
		return nil, err
	}
	return m, nil
}

func (s *service) UpdateSection(id int, name string, gradoId int) (*models.Sections, error) {
	classroom := models.Sections{
		Id:        id,
		Name:      name,
		GradoId:   gradoId,
		Status:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	m, err := s.repository.updateSection(classroom)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't getByNickName row:", err)
		return nil, err
	}
	return m, nil
}

func (s *service) CreateSection(name string, gradoId int) (*models.Sections, error) {
	classroom := models.Sections{
		Name:      name,
		GradoId:   gradoId,
		Status:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	m, err := s.repository.insertSection(classroom)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't getByNickName row:", err)
		return nil, err
	}
	return m, nil
}

func (s *service) DeleteSection(id int, name string, gradoId int) (*models.Sections, error) {
	classroom := models.Sections{
		Name:      name,
		GradoId:   gradoId,
		Status:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	m, err := s.repository.deleteSection(classroom)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't getByNickName row:", err)
		return nil, err
	}
	return m, nil
}
