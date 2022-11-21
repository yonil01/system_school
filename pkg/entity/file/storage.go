package file

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

type ServicesFileRepository interface {
	create(m *File) error
	update(m *File) error
	delete(id int) error
	getByID(id int) (*File, error)
	getAll() ([]*File, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesFileRepository {
	var s ServicesFileRepository
	engine := db.DriverName()
	switch engine {
	case SqlServer:
		return newFileSqlServerRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
