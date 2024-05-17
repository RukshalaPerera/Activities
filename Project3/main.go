package main

import (
	"Project3/app/db"
	"Project3/app/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db.Initialize()

	app := fiber.New()

	routes.SetupBookRoutes(app)
	routes.SetupReservationRoutes(app)

	routes.SetUpReportsRoutes(app)
	routes.SetupReservationRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		panic("Failed to start server")
	}
}
