package Payment

import (
	"foro-hotel/internal/logger"
	"foro-hotel/internal/msg"
	"foro-hotel/pkg/data"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type handlerPayment struct {
	DB   *sqlx.DB
	TxID string
}

func (h *handlerPayment) getPaymentAll(c *fiber.Ctx) error {
	res := responsePayments{Error: true}

	srvAuth := data.NewServerData(h.DB, nil, h.TxID)

	us, err := srvAuth.SrvData.GetPayments()
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

func (h *handlerPayment) updatePayment(c *fiber.Ctx) error {
	res := responseUpdate{Error: true}
	m := requestPayment{}
	err := c.BodyParser(&m)
	if err != nil {
		logger.Error.Printf("couldn't bind model login: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "")
		return c.Status(http.StatusAccepted).JSON(res)
	}

	srvAuth := data.NewServerData(h.DB, nil, h.TxID)

	us, err := srvAuth.SrvData.UpdatePayment(m.Id, m.Name, m.Description, m.Motivo, m.Status, m.DatePayment, m.UserMatricula, m.Amount, m.Role)
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

func (h *handlerPayment) createPayment(c *fiber.Ctx) error {
	res := responseUpdate{Error: true}
	m := requestPayment{}
	err := c.BodyParser(&m)
	if err != nil {
		logger.Error.Printf("couldn't bind model login: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "")
		return c.Status(http.StatusAccepted).JSON(res)
	}

	srvAuth := data.NewServerData(h.DB, nil, h.TxID)

	us, err := srvAuth.SrvData.CreatePayment(m.Name, m.Description, m.Motivo, m.Status, m.DatePayment, m.UserMatricula, m.Amount, m.Role)
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

func (h *handlerPayment) deletePayment(c *fiber.Ctx) error {
	res := responseUpdate{Error: true}
	m := requestPayment{}
	err := c.BodyParser(&m)
	if err != nil {
		logger.Error.Printf("couldn't bind model login: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "")
		return c.Status(http.StatusAccepted).JSON(res)
	}

	srvAuth := data.NewServerData(h.DB, nil, h.TxID)

	us, err := srvAuth.SrvData.DeletePayment(m.Id, m.Name, m.Description, m.Motivo, m.Status, m.DatePayment, m.UserMatricula, m.Amount, m.Role)
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
