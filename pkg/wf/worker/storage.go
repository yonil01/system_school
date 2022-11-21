package worker

import (
	"foro-hotel/internal/logger"
	"foro-hotel/internal/models"
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesWorkerRepository interface {
	create(m *Worker) error
	update(m *Worker) error
	delete(id int) error
	getByID(id int) (*Worker, error)
	getAll() ([]*Worker, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesWorkerRepository {
	var s ServicesWorkerRepository
	engine := db.DriverName()
	switch engine {
	case SqlServer:
		return newWorkerSqlServerRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
