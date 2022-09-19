package classroom

import (
	"foro-hotel/internal/logger"
	"foro-hotel/internal/msg"
	"foro-hotel/pkg/data"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type handlerClassroom struct {
	DB   *sqlx.DB
	TxID string
}

func (h *handlerClassroom) getClassroomAll(c *fiber.Ctx) error {
	res := responseClassrooms{Error: true}

	srvAuth := data.NewServerData(h.DB, nil, h.TxID)

	us, err := srvAuth.SrvData.GetClassrooms()
	if err != nil {
		// TODO implements code
		logger.Warning.Printf("The token was not sent: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "")
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = us
	res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "")
	res.Error = false
	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerClassroom) updateClassroom(c *fiber.Ctx) error {
	res := responseUpdate{Error: true}
	m := requestClassroom{}
	err := c.BodyParser(&m)
	if err != nil {
		logger.Error.Printf("couldn't bind model login: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "")
		return c.Status(http.StatusAccepted).JSON(res)
	}

	srvAuth := data.NewServerData(h.DB, nil, h.TxID)

	us, err := srvAuth.SrvData.UpdateClassroom(m.Id, m.Name, m.Description, m.Nivel, m.Grado, m.Section)
	if err != nil {
		// TODO implements code
		logger.Warning.Printf("The token was not sent: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "")
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = us
	res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "")
	res.Error = false
	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerClassroom) createClassroom(c *fiber.Ctx) error {
	res := responseUpdate{Error: true}
	m := requestClassroom{}
	err := c.BodyParser(&m)
	if err != nil {
		logger.Error.Printf("couldn't bind model login: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "")
		return c.Status(http.StatusAccepted).JSON(res)
	}

	srvAuth := data.NewServerData(h.DB, nil, h.TxID)

	us, err := srvAuth.SrvData.CreateClassroom(m.Name, m.Description, m.Nivel, m.Grado, m.Section)
	if err != nil {
		// TODO implements code
		logger.Warning.Printf("The token was not sent: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "")
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = us
	res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "")
	res.Error = false
	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerClassroom) deleteClassroom(c *fiber.Ctx) error {
	res := responseUpdate{Error: true}
	m := requestClassroom{}
	err := c.BodyParser(&m)
	if err != nil {
		logger.Error.Printf("couldn't bind model login: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "")
		return c.Status(http.StatusAccepted).JSON(res)
	}

	srvAuth := data.NewServerData(h.DB, nil, h.TxID)

	us, err := srvAuth.SrvData.DeleteClassroom(m.Id, m.Name, m.Description, m.Nivel, m.Grado, m.Section)
	if err != nil {
		// TODO implements code
		logger.Warning.Printf("The token was not sent: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "")
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = us
	res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "")
	res.Error = false
	return c.Status(http.StatusOK).JSON(res)
}
