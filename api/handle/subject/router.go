package Subject

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterSubject(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerSubject{DB: db, TxID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	autofill := v1.Group("/subjects")
	autofill.Get("/all", h.getSubjectAll)
	autofill.Get("/:user_id", h.getSubjectByUser)
	autofill.Post("/update", h.updateSubject)
	autofill.Post("/create", h.createSubject)
	autofill.Post("/delete", h.deleteSubject)
}
