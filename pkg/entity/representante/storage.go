package representante

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

type ServicesRepresentanteRepository interface {
	create(m *Representante) error
	update(m *Representante) error
	delete(id int) error
	getByID(id int) (*Representante, error)
	getAll() ([]*Representante, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesRepresentanteRepository {
	var s ServicesRepresentanteRepository
	engine := db.DriverName()
	switch engine {
	case SqlServer:
		return newRepresentanteSqlServerRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
