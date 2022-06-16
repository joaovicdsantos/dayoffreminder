package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaovicdsantos/dayoffreminder/database"
	"github.com/joaovicdsantos/dayoffreminder/model"
)

func GetDayOffs(c *fiber.Ctx) error {
	var dayOffs []model.DayOff
	db := database.DBConn
	db.Find(&dayOffs)
	return c.JSON(dayOffs)
}

func CreateDayOff(c *fiber.Ctx) error {
	var dayOff model.DayOff
	var slackRequest model.SlackRequest
	db := database.DBConn

	if err := c.BodyParser(&slackRequest); err != nil {
		return c.JSON(fiber.Map{
			"response_type": "ephemeral",
			"text":          err.Error(),
		})
	}

	if err := dayOff.SlackRequestToDayOff(slackRequest); err != nil {
		return c.JSON(fiber.Map{
			"response_type": "ephemeral",
			"text":          err.Error(),
		})
	}

	db.Create(&dayOff)

	return c.JSON(fiber.Map{
		"response_type": "in_channel",
		"text":          "Beleza! Agendado.",
	})
}
