package main

import (
	"Project2/app/db"
	"Project2/app/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db.Initialize()

	app := fiber.New()

	routes.SetupUserRoutes(app)
	routes.SetUpRoleRoutes(app)
	routes.SetUpDocumentRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		panic("Failed to start server")
	}
}
