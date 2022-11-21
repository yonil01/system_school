package materia

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

type ServicesMateriaRepository interface {
	create(m *Materia) error
	update(m *Materia) error
	delete(id int) error
	getByID(id int) (*Materia, error)
	getAll() ([]*Materia, error)
	getGradeId(gradeId int) ([]*Materia, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesMateriaRepository {
	var s ServicesMateriaRepository
	engine := db.DriverName()
	switch engine {
	case SqlServer:
		return newMateriaSqlServerRepository(db, user, txID)

		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
