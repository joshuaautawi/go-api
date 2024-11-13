package handler

import (
	"github.com/gofiber/fiber/v2"
	baseDTO "github.com/joshuaautawi/go-api/internal/common/dto"
	"github.com/joshuaautawi/go-api/internal/common/utils"
	"github.com/joshuaautawi/go-api/internal/user/dto"
	"github.com/joshuaautawi/go-api/internal/user/models"
	"github.com/joshuaautawi/go-api/internal/user/service"
	"github.com/joshuaautawi/go-api/pkg/db/postgres"
)

// Create a user
func CreateUser(c *fiber.Ctx) error {
	req := new(dto.CreateOne)
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

// GetSingleUser from db
func GetSingleUser(c *fiber.Ctx) error {
	db := postgres.DB.Db
	// get id params
	id := c.Params("id")
	var user models.User
	// find single user in the database by id
	result := db.Find(&user, "id = ?", id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "User Found", "data": user})
}

// update a user in db
func UpdateUser(c *fiber.Ctx) error {
	type updateUser struct {
		Username string `json:"username"`
	}
	db := postgres.DB.Db
	var user models.User
	// get id params
	id := c.Params("id")
	// find single user in the database by id
	result := db.Find(&user, "id = ?", id)

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}
	var updateUserData updateUser
	err := c.BodyParser(&updateUserData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	user.Username = updateUserData.Username
	// Save the Changes
	db.Save(&user)
	// Return the updated user
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "users Found", "data": user})
}

// delete user in db by ID
func DeleteUserByID(c *fiber.Ctx) error {
	db := postgres.DB.Db
	var user models.User
	// get id params
	id := c.Params("id")
	// find single user in the database by id
	result := db.Find(&user, "id = ?", id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}
	err := db.Delete(&user, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete user", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User deleted"})
}
