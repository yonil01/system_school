package wf

import (
	"foro-hotel/internal/models"
	"foro-hotel/pkg/wf/worker"
	"github.com/jmoiron/sqlx"
)

type ServerData struct {
	SrvWorker worker.PortsServerWorker
}

func NewServerWf(db *sqlx.DB, user *models.User, txID string) *ServerData {
	repoWorker := worker.FactoryStorage(db, user, txID)
	srvWorker := worker.NewWorkerService(repoWorker, user, txID)

	return &ServerData{
		SrvWorker: srvWorker,
	}
}
