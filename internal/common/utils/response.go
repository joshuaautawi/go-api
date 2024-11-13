package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joshuaautawi/go-api/internal/common/dto"
)

func HandleErrorResponse[T any](c *fiber.Ctx, err *dto.Error, res *dto.Response[T]) error {
	res.Error = err
	res.Status = "error"
	return c.Status(err.Code).JSON(res)
}
