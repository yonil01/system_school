package data

import (
	"foro-hotel/internal/models"
	"foro-hotel/pkg/data/profile"
	"github.com/jmoiron/sqlx"
)

type ServerData struct {
	SrvData profile.PortServiceProfile
}

func NewServerData(db *sqlx.DB, user *models.User, txID string) *ServerData {
	repoData := profile.FactoryStorage(db, user, txID)
	srvData := profile.NewProfileService(repoData, user, txID, db)

	return &ServerData{
		SrvData: srvData,
	}
}
