package ocr

import (
	"foro-hotel/internal/models"
	"foro-hotel/pkg/ocr/upload_file"
	"github.com/jmoiron/sqlx"
)

type Server struct {

	SrvUploadFile upload_file.PortsServerFile

}

func NewServerOcr(db *sqlx.DB, user *models.User, txID string) *Server {

	repoUploadFile := upload_file.FactoryStorage(db, user, txID)
	srvUploadFile := upload_file.NewUploadFileService(repoUploadFile, user, txID)


	return &Server{
		SrvUploadFile: srvUploadFile,
	}
}
