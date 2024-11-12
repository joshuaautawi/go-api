package repository

import (
	"log"

	baseDTO "github.com/joshuaautawi/go-api/internal/common/dto"
	"github.com/joshuaautawi/go-api/internal/common/utils"
	"github.com/joshuaautawi/go-api/internal/user/dto"
	"github.com/joshuaautawi/go-api/internal/user/models"
	"github.com/joshuaautawi/go-api/pkg/db/postgres"
)

func GetAll(input *baseDTO.GetAllRequest) (*baseDTO.GetAllBaseResponse[[]models.User], error) {
	var users []models.User
	db := postgres.DB.Db
	var totalCount int64
	if err := db.Model(&models.User{}).Count(&totalCount).Error; err != nil {
		return nil, err
	}

	result := db.Limit(input.Limit).Offset((input.Page - input.Limit) * input.Limit).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	meta := baseDTO.Meta{
		TotalCount: 1,
		Page:       input.Page,
		Limit:      input.Limit,
	}
	return &baseDTO.GetAllBaseResponse[[]models.User]{
		Data: users,
		Meta: &meta,
	}, nil
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
