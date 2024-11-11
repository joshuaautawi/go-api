package repository

import (
	"log"

	"github.com/joshuaautawi/go-api/internal/user/dto"
	"github.com/joshuaautawi/go-api/internal/user/models"
	"github.com/joshuaautawi/go-api/internal/utils"
	"github.com/joshuaautawi/go-api/pkg/db/postgres"
)

func GetAll() models.Users {
	var users models.Users
	db := postgres.DB.Db
	db.Find(&users)
	return users
}

func GetOneByID(req dto.GetOneByIDRequest) (*models.User, error) {
	var user models.User
	if err := postgres.DB.Db.First(&user, "id = ?", req.ID).Error; err != nil {
		log.Println("Error fetching user by ID:", err)
		return nil, err
	}
	return &user, nil
}

func Create(user *dto.CreateOne) (*models.User, error) {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Println("Error hashing password:", err)
		return nil, err
	}

	// Set the password to the hashed version
	user.Password = hashedPassword
	newUser := models.User{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
	}
	if err := postgres.DB.Db.Create(&newUser).Error; err != nil {
		log.Println("Error creating user:", err)
		return nil, err
	}
	return &newUser, nil
}
