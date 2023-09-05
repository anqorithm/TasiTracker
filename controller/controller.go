package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/qahta0/tasitracker/database"
	"github.com/qahta0/tasitracker/model"
)

func GetCommpanies(c *fiber.Ctx) error {
	db := database.Connect()
	var comapnies []model.Company
	db.Preload("TodayPoints").Find(&comapnies)
	if len(comapnies) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Comapnies not found", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Comapnies Found", "data": comapnies})
}
