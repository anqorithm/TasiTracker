package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/qahta0/tasitracker/controller"
	"gorm.io/gorm"
)

func SetupRoutes(dbConnection *gorm.DB, app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	companies := v1.Group("/companies")
	companies.Get("/", controller.GetCommpanies)
}
