package auth

import (
	"foro-hotel/internal/models"
	saveimageperfil "foro-hotel/pkg/auth/saveImageperfil"
	"foro-hotel/pkg/auth/users"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	SrvUsers       users.PortsServerUser
	SrvImagePerfil saveimageperfil.PortSaveImageNegocios
}

func NewServerAuth(db *sqlx.DB, user *models.User, txID string) *Server {

	repoUsers := users.FactoryStorage(db, user, txID)
	srvUsers := users.NewUserService(repoUsers, user, txID, db)

	return &Server{
		SrvUsers: srvUsers,
	}
}
