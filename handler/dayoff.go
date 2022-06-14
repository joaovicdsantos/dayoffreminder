package handler

import "github.com/gofiber/fiber/v2"

func GetDayOffs(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
