package consultRENIECSUNAT

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterConsultReniecSunat(app *fiber.App, db *sqlx.DB, txID string) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Post("/consult-reniec", consultReniecSunat)
}
