package dbo

import (
	"foro-hotel/internal/models"
	"foro-hotel/pkg/dbo/materia"
	"github.com/jmoiron/sqlx"
)

type ServerData struct {
	SrvMateria materia.PortsServerMateria
}

func NewServerDbo(db *sqlx.DB, user *models.User, txID string) *ServerData {
	repoMateria := materia.FactoryStorage(db, user, txID)
	srvMateria := materia.NewMateriaService(repoMateria, user, txID)

	return &ServerData{
		SrvMateria: srvMateria,
	}
}
