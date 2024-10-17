package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joshuaautawi/go-api/cmd/app/routers"
	"github.com/joshuaautawi/go-api/configs"
	"github.com/joshuaautawi/go-api/databases"
	_ "github.com/lib/pq"
)

func main() {
	databases.ConnectPostgres()
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())
	routers.SetupRoutes(app)
	// handle unavailable route
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
	port := configs.Config("PORT")
	app.Listen(":" + port)
}
