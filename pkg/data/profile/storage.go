package profile

import (
	"foro-hotel/internal/logger"
	"foro-hotel/internal/models"
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "sqlserver"
)

type ServicesDataRepository interface {
	getUserbyId(string) (*models.User, int, error)
	getComentariosProcedure(document_id string) ([]*models.Comentario, int, error)
	getPdf(text_hash string) (*models.Reservacion, int, error)
	getReservacionByDocument(id string) ([]*models.Reservacion, int, error)
	getReservaciones(id string) ([]*models.Reservacion, int, error)
	getCountries(data string) ([]*models.Country, int, error)
	getCities(data string) ([]*models.City, int, error)
	getUser(role int) ([]*models.User, error)
	updateUser(mdl models.User) (*models.User, error)
	insertUser(mdl models.User) (*models.User, error)
	deleteUser(mdl models.User) (*models.User, error)
	getClassrooms() ([]*models.Classroom, error)
	updateClassroom(mdl models.Classroom) (*models.Classroom, error)
	insertClassroom(mdl models.Classroom) (*models.Classroom, error)
	deleteClassroom(mdl models.Classroom) (*models.Classroom, error)
	getSubjects() ([]*models.Subject, error)
	updateSubject(mdl models.Subject) (*models.Subject, error)
	insertSubject(mdl models.Subject) (*models.Subject, error)
	deleteSubject(mdl models.Subject) (*models.Subject, error)
	getPayments() ([]*models.Payment, error)
	updatePayment(mdl models.Payment) (*models.Payment, error)
	insertPayment(mdl models.Payment) (*models.Payment, error)
	deletePayment(mdl models.Payment) (*models.Payment, error)
	ExecuteSP(m *models.Report) ([]map[string]interface{}, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesDataRepository {
	var s ServicesDataRepository
	engine := db.DriverName()
	switch engine {
	case Postgresql:
		return newDataPsqlRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
