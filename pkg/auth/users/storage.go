package users

import (
	"foro-hotel/internal/logger"
	"foro-hotel/internal/models"
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "sqlserver"
)

type ServicesUserRepository interface {
	create(m *User) (*User, error)
	// getUser(string, string) (int, error)
	getUserByEmail(email string) (*User, int, error)
	createLogin(m *User) (int, error)
	saveImagePerfil(document_id string, images *models.Image) (int, error)
	getDoctypeByUser(userId int) ([]*models.DoctypeUser, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesUserRepository {
	var s ServicesUserRepository
	engine := db.DriverName()
	switch engine {
	case Postgresql:
		return newUserPsqlRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
