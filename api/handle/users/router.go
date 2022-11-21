package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterUser(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerUser{DB: db, TxID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	autofill := v1.Group("/user")
	autofill.Get("/all/:role", h.getUserAll)
	autofill.Get("/:id", h.getUserById)
	autofill.Post("/update", h.updateUser)
	autofill.Post("/create", h.createUser)
	autofill.Post("/delete", h.deleteUser)

	representante := v1.Group("/representante")
	representante.Post("/create", h.createRepresentante)

	file := v1.Group("/file")
	file.Post("/create", h.uploadFileAnexos)

}
