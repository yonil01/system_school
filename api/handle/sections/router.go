package sections

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterSection(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerSection{DB: db, TxID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	autofill := v1.Group("/sections")
	autofill.Get("/all", h.getSectionAll)
	autofill.Post("/update", h.updateSection)
	autofill.Post("/create", h.createSection)
	autofill.Post("/delete", h.deleteSection)
}
