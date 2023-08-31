package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sonyasergeevass/API_Go/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.GetSegm)
	app.Post("/segm", handlers.CreateSegment)
	app.Get("/show_user", handlers.GetUser)
	app.Post("/user", handlers.CreateUser)
	app.Delete("/segment", handlers.DeleteSegment)
	app.Post("/usersegment", handlers.CreateUserSegment)
	app.Get("/showus", handlers.GetUserSegments)
}
