package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joshuaautawi/go-api/cmd/app/routers"
	"github.com/joshuaautawi/go-api/pkg/db/postgres"
	_ "github.com/lib/pq"
)

func main() {
	postgres.ConnectPostgres()
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())
	routers.SetupRoutes(app)
	// handle unavailable route
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
	port := "6969"
	app.Listen(":" + port)
}
