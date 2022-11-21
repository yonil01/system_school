package worker

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterWorker(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerWorker{DB: db, TxID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	autofill := v1.Group("/worker")
	autofill.Get("/all/:role", h.getWorkerAll)
	autofill.Get("/:id", h.getWorkerById)
	autofill.Post("/update", h.updateWorker)
	autofill.Post("/create", h.createWorker)
	autofill.Post("/delete", h.deleteWorker)

}
