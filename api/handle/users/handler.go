package user

import (
	"foro-hotel/internal/logger"
	"foro-hotel/internal/msg"
	"foro-hotel/pkg/data"
	"foro-hotel/pkg/entity"
	"foro-hotel/pkg/entity/file"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"net/http"
	"strconv"
)

type handlerUser struct {
	DB   *sqlx.DB
	TxID string
}

func (h *handlerUser) getUserAll(c *fiber.Ctx) error {
	res := responseUsers{Error: true}

	srvAuth := data.NewServerData(h.DB, nil, h.TxID)

	role, err := strconv.Atoi(c.Params("role"))
	if err != nil {
		// TODO implements code
		logger.Warning.Printf("The token was not sent: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "")
		return c.Status(http.StatusAccepted).JSON(res)
	}
	us, err := srvAuth.SrvData.GetUser(role)
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

func (h *handlerUser) updateUser(c *fiber.Ctx) error {
	res := responseUpdate{Error: true}
	m := requestUser{}
	err := c.BodyParser(&m)
	if err != nil {
		logger.Error.Printf("couldn't bind model login: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "")
		return c.Status(http.StatusAccepted).JSON(res)
	}

	srvAuth := data.NewServerData(h.DB, nil, h.TxID)

	us, err := srvAuth.SrvData.UpdateUser(m.Matricula, m.Dni, m.Username, m.Names, m.Lastnames, m.Sexo, m.Email, m.DateBirth, m.DateAdmission)
	if err != nil {
		// TODO implements code
		logger.Warning.Printf("The token was not sent: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "")
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = us
	res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "Se actualizo Los datos")
	res.Error = false
	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerUser) createUser(c *fiber.Ctx) error {
	res := responseUpdate{Error: true}
	m := requestUser{}
	err := c.BodyParser(&m)
	if err != nil {
		logger.Error.Printf("couldn't bind model login: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "")
		return c.Status(http.StatusAccepted).JSON(res)
	}

	srvAuth := data.NewServerData(h.DB, nil, h.TxID)

	us, err := srvAuth.SrvData.CreateUser(m.Dni, m.Username, m.Names, m.Lastnames, m.Sexo, m.Email, m.DateBirth, m.DateAdmission, m.Role)
	if err != nil {
		// TODO implements code
		logger.Warning.Printf("The token was not sent: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "")
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = us
	res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "Creado Correctamente")
	res.Error = false
	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerUser) deleteUser(c *fiber.Ctx) error {
	res := responseUpdate{Error: true}
	m := requestUser{}
	err := c.BodyParser(&m)
	if err != nil {
		logger.Error.Printf("couldn't bind model login: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "")
		return c.Status(http.StatusAccepted).JSON(res)
	}

	srvAuth := data.NewServerData(h.DB, nil, h.TxID)

	us, err := srvAuth.SrvData.DeleteUser(m.Matricula, m.Dni, m.Username, m.Names, m.Lastnames, m.Sexo, m.Email, m.DateBirth, m.DateAdmission)
	if err != nil {
		// TODO implements code
		logger.Warning.Printf("The token was not sent: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "")
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = us
	res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "Se elimino con exito")
	res.Error = false
	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerUser) getUserById(c *fiber.Ctx) error {
	res := responseUser{Error: true}

	srvAuth := data.NewServerData(h.DB, nil, h.TxID)

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		// TODO implements code
		logger.Warning.Printf("The token was not sent: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "")
		return c.Status(http.StatusAccepted).JSON(res)
	}
	us, cod, err := srvAuth.SrvData.GetUserById(id)
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

// TODO REPRESENTANTE

func (h *handlerUser) createRepresentante(c *fiber.Ctx) error {
	res := responsRepresentante{Error: true}
	m := RequestRepresentante{}
	err := c.BodyParser(&m)
	if err != nil {
		logger.Error.Printf("couldn't bind model login: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "")
		return c.Status(http.StatusAccepted).JSON(res)
	}

	srvEntity := entity.NewServerEntity(h.DB, nil, h.TxID)

	rep, cod, err := srvEntity.SrvRepresentante.CreateRepresentante(m.MatriculaUser, m.TypeRepresentante, m.Notification, m.Dni, m.Direction, m.Names, m.Lastnames, m.CellPhone, m.Email, 1, 0)
	if err != nil {
		// TODO implements code
		logger.Warning.Printf("The token was not sent: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(cod, "", "")
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = rep
	res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "Creado Correctamente")
	res.Error = false
	return c.Status(http.StatusOK).JSON(res)
}

//TODO FILE
func (h *handlerUser) uploadFileAnexos(c *fiber.Ctx) error {
	res := responseFiles{Error: true}
	m := RequestFile{}
	err := c.BodyParser(&m)
	if err != nil {
		logger.Error.Printf("couldn't bind model login: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "")
		return c.Status(http.StatusAccepted).JSON(res)
	}

	srvEntity := entity.NewServerEntity(h.DB, nil, h.TxID)

	var files []*file.File

	for _, file := range m.Files {
		f, cod, err := srvEntity.SrvFile.CreateFile(file.MatriculaUser, file.Name, file.Description, "files/", file.FileName, file.B64, file.TypeFile, 1, 0)
		if err != nil {
			// TODO implements code
			logger.Warning.Printf("The token was not sent: %v", err)
			res.Code, res.Type, res.Msg = msg.GetByCode(cod, "", "")
			return c.Status(http.StatusAccepted).JSON(res)
		}
		files = append(files, f)
	}

	res.Data = files
	res.Code, res.Type, res.Msg = msg.GetByCode(1, "", "Creado Correctamente")
	res.Error = false
	return c.Status(http.StatusOK).JSON(res)
}
