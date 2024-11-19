package middlewares

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/joshuaautawi/go-api/configs"
	"github.com/joshuaautawi/go-api/internal/common/dto"
	"github.com/joshuaautawi/go-api/internal/common/utils"
)

var secret = configs.Config("JWT_SECRET")
var jwtSecret = []byte(secret) // Use a secure key from environment variables

func JWTProtected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtSecret,
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	res := dto.Response[string]{}
	errRes := utils.JWTMiddlewareError(err.Error())
	res.Error = &errRes
	return c.Status(fiber.StatusUnauthorized).JSON(res)
}
