package entity

import (
	"foro-hotel/internal/models"
	"foro-hotel/pkg/entity/file"
	"foro-hotel/pkg/entity/representante"
	"github.com/jmoiron/sqlx"
)

type ServerData struct {
	SrvRepresentante representante.PortsServerRepresentante
	SrvFile          file.PortsServerFile
}

func NewServerEntity(db *sqlx.DB, user *models.User, txID string) *ServerData {
	repoRepresentante := representante.FactoryStorage(db, user, txID)
	srvRepresentante := representante.NewRepresentanteService(repoRepresentante, user, txID)

	repoFile := file.FactoryStorage(db, user, txID)
	srvFile := file.NewFileService(repoFile, user, txID)

	return &ServerData{
		SrvRepresentante: srvRepresentante,
		SrvFile:          srvFile,
	}
}
