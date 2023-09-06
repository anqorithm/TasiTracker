package controller

import (
	"errors"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/qahta0/tasitracker/database"
	"github.com/qahta0/tasitracker/model"
	"gorm.io/gorm"
)

func GetCompany(c *fiber.Ctx) error {
	db := database.Connect()
	id := c.Params("id")
	var company model.Company
	result := db.Preload("TodayPoints").Where("guid = ?", id).First(&company)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Company not found", "data": nil})
		}
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Internal server error", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Company fetched successfully", "data": company})
}

func GetCompanies(c *fiber.Ctx) error {
	db := database.Connect()
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "5"))
	offset := (page - 1) * limit
	var companies []model.Company
	var total int64
	db.Model(&model.Company{}).Count(&total)
	db.Preload("TodayPoints").Offset(offset).Limit(limit).Find(&companies)
	if len(companies) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Companies not found", "data": nil})
	}
	response := fiber.Map{
		"status":  "success",
		"message": "Companies fetched successfully",
		"data":    companies,
		"meta": fiber.Map{
			"total":    total,
			"page":     page,
			"lastPage": math.Ceil(float64(total) / float64(limit)),
			"pageSize": limit,
		},
	}
	return c.Status(200).JSON(response)
}
