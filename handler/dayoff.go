package handler

import (
	"fmt"

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
	db := database.DBConn

	if err := dayOff.QueryToDayOff(c.Query("text"), c.Query("user_name")); err != nil {
		c.JSON(fiber.Map{
			"message": err.Error(),
		})
		fmt.Print(err.Error())
		return c.SendStatus(400)
	}

	db.Create(&dayOff)

	c.JSON(dayOff)
	return c.SendStatus(201)
}
