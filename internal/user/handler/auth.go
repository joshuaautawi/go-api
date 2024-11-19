package handler

import (
	"github.com/gofiber/fiber/v2"
	baseDTO "github.com/joshuaautawi/go-api/internal/common/dto"
	"github.com/joshuaautawi/go-api/internal/common/utils"
	"github.com/joshuaautawi/go-api/internal/user/dto"
	"github.com/joshuaautawi/go-api/internal/user/service"
)

func Login(c *fiber.Ctx) error {
	req := new(dto.LoginRequest)
	res := baseDTO.Response[string]{}

	if err := c.BodyParser(req); err != nil {
		err := utils.ParseError(err.Error())
		return utils.HandleErrorResponse(c, &err, &res)
	}
	if err := utils.HandleValidation(req); err != nil {
		return utils.HandleErrorResponse(c, err, &res)
	}
	token, err := service.Login(req)
	if err != nil {
		return utils.HandleErrorResponse(c, err, &res)
	}
	res.Data = token
	return c.Status(fiber.StatusOK).JSON(res)
}
