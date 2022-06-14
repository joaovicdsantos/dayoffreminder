package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaovicdsantos/dayoffreminder/handler"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	dayoffRoutes := api.Group("/dayoff")
	dayoffRoutes.Get("/", handler.GetDayOffs)
}
