package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joshuaautawi/go-api/internal/common/middlewares"
	"github.com/joshuaautawi/go-api/internal/user/handler"
)

var userPath = "/users"
var loginPath = "/login"

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// grouping
	api := app.Group("/api")
	userV1 := api.Group(userPath)
	loginV1 := api.Group(loginPath)

	userV1.Get("/", handler.GetAllUsers)
	userV1.Get("/:id", middlewares.JWTProtected(), handler.GetOne)
	userV1.Post("/", handler.CreateUser)
	userV1.Put("/", handler.UpdateUser)
	userV1.Delete("/:id", handler.DeleteUserByID)

	loginV1.Post("/", handler.Login)
}
