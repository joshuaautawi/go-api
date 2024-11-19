package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	baseDTO "github.com/joshuaautawi/go-api/internal/common/dto"
	"github.com/joshuaautawi/go-api/internal/common/utils"
	"github.com/joshuaautawi/go-api/internal/user/dto"
	"github.com/joshuaautawi/go-api/internal/user/models"
	"github.com/joshuaautawi/go-api/internal/user/service"
)

// Create a user
func CreateUser(c *fiber.Ctx) error {
	req := new(dto.CreateOneRequest)
	res := baseDTO.Response[*models.User]{}

	if err := c.BodyParser(req); err != nil {
		err := utils.ParseError(err.Error())
		return utils.HandleErrorResponse(c, &err, &res)
	}

	if err := utils.HandleValidation(req); err != nil {
		return utils.HandleErrorResponse(c, err, &res)
	}

	user, err := service.CreateOne(req)

	if err != nil {
		return utils.HandleErrorResponse(c, err, &res)
	}
	res.Data = user
	return c.Status(fiber.StatusCreated).JSON(res)

}

// Get All Users from db
func GetAllUsers(c *fiber.Ctx) error {
	req := new(baseDTO.GetAllRequest)
	errParser := c.QueryParser(req)
	res := baseDTO.Response[*[]models.User]{}

	if errParser != nil {
		err := utils.ParseError(errParser.Error())
		return utils.HandleErrorResponse(c, &err, &res)
	}
	if err := utils.HandleValidation(req); err != nil {
		return utils.HandleErrorResponse(c, err, &res)
	}
	users, meta, err := service.GetAll(req)

	if err != nil {
		return utils.HandleErrorResponse(c, err, &res)
	}
	res.Data = users
	res.Meta = meta
	// Return the created user
	return c.Status(fiber.StatusOK).JSON(res)

}

// GetOne from db
func GetOne(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, convErr := strconv.Atoi(idStr)
	res := baseDTO.Response[*models.User]{}
	if convErr != nil {
		err := utils.ParseError(convErr.Error())
		return utils.HandleErrorResponse(c, &err, &res)
	}
	user, err := service.GetOne(id)
	if err != nil {
		return utils.HandleErrorResponse(c, err, &res)
	}

	res.Data = user
	return c.Status(fiber.StatusOK).JSON(res)
}

// update a user in db
func UpdateUser(c *fiber.Ctx) error {
	req := new(dto.UpdateOneRequest)
	res := baseDTO.Response[*models.User]{}
	if err := c.BodyParser(req); err != nil {
		err := utils.ParseError(err.Error())
		return utils.HandleErrorResponse(c, &err, &res)
	}

	if err := utils.HandleValidation(req); err != nil {
		return utils.HandleErrorResponse(c, err, &res)
	}

	user, err := service.UpdateOne(req)

	if err != nil {
		return utils.HandleErrorResponse(c, err, &res)
	}
	res.Data = user
	return c.Status(fiber.StatusCreated).JSON(res)
}

// delete user in db by ID
func DeleteUserByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	res := baseDTO.Response[*models.User]{}

	id, convErr := strconv.Atoi(idStr)
	if convErr != nil {
		err := utils.ParseError(convErr.Error())
		return utils.HandleErrorResponse(c, &err, &res)
	}
	user, err := service.DeleteOne(id)
	if err != nil {
		return utils.HandleErrorResponse(c, err, &res)
	}
	res.Data = user
	return c.Status(fiber.StatusOK).JSON(res)
}
