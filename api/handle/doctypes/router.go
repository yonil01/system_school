package doctype

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterRegisterUser(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerRegisterUser{DB: db, TxID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Post("/doctype-user", h.doctypeUser)
}
