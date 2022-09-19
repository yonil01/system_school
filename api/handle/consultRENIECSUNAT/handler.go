package consultRENIECSUNAT

import (
	"encoding/json"
	"fmt"
	"foro-hotel/internal/env"
	"foro-hotel/internal/logger"
	"foro-hotel/internal/ws"
	"github.com/gofiber/fiber/v2"
	"net/http"
)


func consultReniecSunat(c *fiber.Ctx) error {
	res := responseReniecSunat{Error: true}
	e := env.NewConfiguration()
	m := RequestGetUserReniec{}
	err := c.BodyParser(&m)

	fmt.Println(m)


	if err != nil {
		res.Code = 1
		res.Type = "Error"
		res.Msg = "data invalid, don´t decoder data"
		logger.Error.Println("No se pudo decodificar la data enviada en el json: %v", err)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	dataBytes, err := json.Marshal(m)
	if err != nil {
		logger.Error.Println(" - couldn't decode request in call ws", err)
		res.Code = 1
		res.Type = "Error"
		res.Msg = "data invalid, don´t decoder data"
		return c.Status(http.StatusAccepted).JSON(res)
	}
	code, response, err := ws.CallApiRest("POST", e.Ws.Url, dataBytes, "Requestverificationtoken", e.Ws.Token)
	if err != nil {
		logger.Error.Println("No se pudo decodificar la data enviada en el json: %v", err)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	fmt.Println(response)
	fmt.Println(code)


	return nil
}