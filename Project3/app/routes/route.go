package routes

import (
	"Project3/app/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupBookRoutes(app *fiber.App) {
	app.Get("/books", handler.GetAllBooks)
	app.Post("/book", handler.CreateBook)
	app.Put("/books/:id", handler.UpdateBook)
	app.Delete("/books/:id", handler.DeleteBook)
}

func SetupReservationRoutes(app *fiber.App) {
	app.Get("/reservations", handler.GetAllReservations)
	app.Post("/reservations", handler.CreateReservation)
	app.Put("/reservations/:id", handler.UpdateReservation)
	app.Delete("/reservations/:id", handler.DeleteReservation)
}

func SetUpReportsRoutes(app *fiber.App) {
	app.Get("/generate-book-report", handler.GenerateBookReport)
	app.Get("/generate-reservation-report", handler.GenerateAllReservationReportHandler)
	app.Get("/Show-book-report", handler.ShowReportPage)
}
