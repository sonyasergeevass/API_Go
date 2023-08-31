package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sonyasergeevass/API_Go/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.GetSegm)
	app.Post("/segm", handlers.CreateSegment)
}
