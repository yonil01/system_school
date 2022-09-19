package upload_file

import (
	"foro-hotel/internal/logger"
	"foro-hotel/internal/models"
	"github.com/jmoiron/sqlx"

)

const (
	Postgresql = "postgres"
)

type ServicesUploadFileRepository interface {
	uploadFile(m *File) error
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesUploadFileRepository {
	var s ServicesUploadFileRepository
	engine := db.DriverName()
	switch engine {
	case Postgresql:
		return newUploadFilePsqlRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}

