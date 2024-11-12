package handler

import (
	"github.com/gofiber/fiber/v2"
	baseDTO "github.com/joshuaautawi/go-api/internal/common/dto"
	"github.com/joshuaautawi/go-api/internal/user/dto"
	"github.com/joshuaautawi/go-api/internal/user/models"
	"github.com/joshuaautawi/go-api/internal/user/service"
	"github.com/joshuaautawi/go-api/pkg/db/postgres"
)

// Create a user
func CreateUser(c *fiber.Ctx) error {
	req := new(dto.CreateOne)
	// Store the body in the user and return error if encountered
	errParser := c.BodyParser(req)
	if errParser != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": errParser})
	}
	user, err := service.CreateOne(req)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}
	// Return the created user
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "User has created", "data": user})

}

// Get All Users from db
func GetAllUsers(c *fiber.Ctx) error {
	input := new(baseDTO.GetAllRequest)
	errParser := c.QueryParser(input)
	if errParser != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": errParser})
	}
	users, err := service.GetAll(input)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not get user", "data": err})
	}
	// Return the created user
	return c.Status(201).JSON(fiber.Map{"status": "success", "data": users})

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
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User Found", "data": user})
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
