package router

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joaovicdsantos/dayoffreminder/handler"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	dayoffRoutes := api.Group("/dayoff")
	dayoffRoutes.Get("/", handler.GetDayOffs)
	dayoffRoutes.Post("/", handler.CreateDayOff)
	log.Println("Configured Routes")
}
