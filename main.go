package main

import (
	"log"
	"os"

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
	database.Migrate()

	log.Printf("Running the server on the port %s\n", os.Getenv("PORT"))
	app.Listen(":" + os.Getenv("PORT"))
}
