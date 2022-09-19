package Payment

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterPayment(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerPayment{DB: db, TxID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	autofill := v1.Group("/payments")
	autofill.Get("/all", h.getPaymentAll)
	autofill.Post("/update", h.updatePayment)
	autofill.Post("/create", h.createPayment)
	autofill.Post("/delete", h.deletePayment)
}
