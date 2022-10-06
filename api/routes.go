package api

import (
	"foro-hotel/api/handle/classroom"
	doctype "foro-hotel/api/handle/doctypes"
	Payment "foro-hotel/api/handle/payment-debt"
	Report "foro-hotel/api/handle/report"
	"foro-hotel/api/handle/sections"
	Subject "foro-hotel/api/handle/subject"
	user "foro-hotel/api/handle/users"
	"github.com/ansrivas/fiberprometheus/v2"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func routes(db *sqlx.DB, loggerHttp bool, allowedOrigins string) *fiber.App {
	app := fiber.New()

	prometheus := fiberprometheus.New("nexumApiAuth")
	prometheus.RegisterAt(app, "/metrics")

	//app.Get("/swagger/*", fiberSwagger.Handler)
	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "/swagger/doc.json",
		DeepLinking: false,
	}))

	app.Use(recover.New())
	app.Use(prometheus.Middleware)
	app.Use(cors.New(cors.Config{
		AllowOrigins: allowedOrigins,
		AllowHeaders: "Origin, X-Requested-With, Content-Type, Accept, Authorization",
		AllowMethods: "GET,POST",
	}))
	if loggerHttp {
		app.Use(logger.New())
	}
	TxID := uuid.New().String()

	loadRoutes(app, db, TxID)

	return app
}

func loadRoutes(app *fiber.App, db *sqlx.DB, TxID string) {
	user.RouterUser(app, db, TxID)
	doctype.RouterRegisterUser(app, db, TxID)
	classroom.RouterClassroom(app, db, TxID)
	Subject.RouterSubject(app, db, TxID)
	Payment.RouterPayment(app, db, TxID)
	Report.RouterReport(app, db, TxID)
	sections.RouterSection(app, db, TxID)
}
