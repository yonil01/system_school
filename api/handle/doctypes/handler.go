package doctype

import (
	"foro-hotel/internal/logger"
	"foro-hotel/internal/msg"
	"foro-hotel/pkg/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type handlerRegisterUser struct {
	DB   *sqlx.DB
	TxID string
}

func (h *handlerRegisterUser) doctypeUser(c *fiber.Ctx) error {

	res := responseRegisterUser{Error: true}
	m := requestDoctype{}
	err := c.BodyParser(&m)
	if err != nil {
		logger.Error.Printf("couldn't bind model login: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "")
		return c.Status(http.StatusAccepted).JSON(res)
	}

	srvNxm := auth.NewServerAuth(h.DB, nil, h.TxID)
	usr, err := srvNxm.SrvUsers.GetDoctypeUser(m.UserId)
	if err != nil {
		res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "")
		logger.Error.Printf("couldn't create user: %v", err)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = usr
	res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "")
	res.Error = false
	return c.Status(http.StatusOK).JSON(res)
}
