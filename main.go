package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joaovicdsantos/dayoffreminder/database"
	"github.com/joaovicdsantos/dayoffreminder/router"
)

func main() {

	app := fiber.New()
	app.Use(logger.New())

	router.SetupRoutes(app)
	database.InitDatabase()

	app.Listen(":3000")
}
