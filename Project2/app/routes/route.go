package routes

import (
	"Project2/app/handler"
	"Project2/app/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	app.Get("/users", handler.GetAllUsers)
	app.Post("/users", handler.CreateUser)
	app.Put("/users/:id", handler.UpdateUser)
	app.Delete("/users/:id", handler.DeleteUser)
}

func SetUpRoleRoutes(app *fiber.App) {
	app.Get("/roles", handler.GetAllRoles)
	app.Post("/roles", handler.CreateRole)
}

func SetUpDocumentRoutes(app *fiber.App) {
	app.Get("/documents", middleware.DocumentAuthorization("user", "admin", "moderator"), handler.ListDocuments)        //list documents
	app.Get("/documents/:id", middleware.DocumentAuthorization("user", "admin", "moderator"), handler.DownloadDocument) //downloading by ID
	app.Post("/documents", middleware.DocumentAuthorization("moderator"), handler.UploadDocument)                       //filename and title
}
