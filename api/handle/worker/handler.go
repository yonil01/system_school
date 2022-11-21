package worker

import (
	"foro-hotel/internal/logger"
	"foro-hotel/internal/msg"
	"foro-hotel/pkg/wf"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"net/http"
	"strconv"
)

type handlerWorker struct {
	DB   *sqlx.DB
	TxID string
}

func (h *handlerWorker) getWorkerAll(c *fiber.Ctx) error {
	res := responseWorkers{Error: true}

	srvAuth := wf.NewServerWf(h.DB, nil, h.TxID)
	us, err := srvAuth.SrvWorker.GetAllWorker()
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

func (h *handlerWorker) updateWorker(c *fiber.Ctx) error {
	res := responseWorker{Error: true}
	m := requestWorker{}
	err := c.BodyParser(&m)
	if err != nil {
		logger.Error.Printf("couldn't bind model login: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "")
		return c.Status(http.StatusAccepted).JSON(res)
	}

	srvAuth := wf.NewServerWf(h.DB, nil, h.TxID)

	us, cod, err := srvAuth.SrvWorker.UpdateWorker(m.Id, m.Matricula, m.Status, 0)
	if err != nil {
		// TODO implements code
		logger.Warning.Printf("The token was not sent: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(cod, "", "")
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = us
	res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "Se actualizo Los datos")
	res.Error = false
	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerWorker) createWorker(c *fiber.Ctx) error {
	res := responseWorker{Error: true}
	m := requestWorker{}
	err := c.BodyParser(&m)
	if err != nil {
		logger.Error.Printf("couldn't bind model login: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "")
		return c.Status(http.StatusAccepted).JSON(res)
	}

	srvAuth := wf.NewServerWf(h.DB, nil, h.TxID)

	us, cod, err := srvAuth.SrvWorker.CreateWorker(m.Matricula, m.Status, 0)
	if err != nil {
		// TODO implements code
		logger.Warning.Printf("The token was not sent: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(cod, "", "")
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = us
	res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "Creado Correctamente")
	res.Error = false
	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerWorker) deleteWorker(c *fiber.Ctx) error {
	res := responseWorker{Error: true}
	m := requestWorker{}
	err := c.BodyParser(&m)
	if err != nil {
		logger.Error.Printf("couldn't bind model login: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "")
		return c.Status(http.StatusAccepted).JSON(res)
	}

	srvAuth := wf.NewServerWf(h.DB, nil, h.TxID)

	cod, err := srvAuth.SrvWorker.DeleteWorker(m.Id)
	if err != nil {
		// TODO implements code
		logger.Warning.Printf("The token was not sent: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(cod, "", "")
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "Se elimino con exito")
	res.Error = false
	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerWorker) getWorkerById(c *fiber.Ctx) error {
	res := responseWorker{Error: true}

	srvAuth := wf.NewServerWf(h.DB, nil, h.TxID)

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		// TODO implements code
		logger.Warning.Printf("The token was not sent: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "")
		return c.Status(http.StatusAccepted).JSON(res)
	}
	us, cod, err := srvAuth.SrvWorker.GetWorkerByID(id)
	if err != nil {
		// TODO implements code
		logger.Warning.Printf("The token was not sent: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "")
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = us
	res.Code, res.Type, res.Msg = msg.GetByCode(cod, "", "")
	res.Error = false
	return c.Status(http.StatusOK).JSON(res)
}
