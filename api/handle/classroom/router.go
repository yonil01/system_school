package classroom

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterClassroom(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerClassroom{DB: db, TxID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	autofill := v1.Group("/classrooms")
	autofill.Get("/all", h.getClassroomAll)
	autofill.Post("/update", h.updateClassroom)
	autofill.Post("/create", h.createClassroom)
	autofill.Post("/delete", h.deleteClassroom)
}
