package Report

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterReport(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerReport{DB: db, TxID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	autofill := v1.Group("/report")
	autofill.Post("/procedure", h.getReportAll)
}
