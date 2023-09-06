package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
	"github.com/qahta0/tasitracker/database"
	"github.com/qahta0/tasitracker/router"
)

func main() {
	dbConnection := database.Connect()
	app := fiber.New()
	app.Use(cors.New())
	// c := cron.New(cron.WithSeconds())
	// c.AddFunc("@every 1m", func() { jobs.FetchAndStoreStocks(dbConnection) })
	// c.Start()
	router.SetupRoutes(dbConnection, app)
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})
	app.Listen(":8080")
}
