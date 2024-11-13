package repository

import (
	"log"

	baseDTO "github.com/joshuaautawi/go-api/internal/common/dto"
	"github.com/joshuaautawi/go-api/internal/common/utils"
	"github.com/joshuaautawi/go-api/internal/user/dto"
	"github.com/joshuaautawi/go-api/internal/user/models"
	"github.com/joshuaautawi/go-api/pkg/db/postgres"
)

func GetAll(input *baseDTO.GetAllRequest) (*[]models.User, *baseDTO.Meta, *baseDTO.Error) {
	var totalCount int64
	var users []models.User
	meta := baseDTO.Meta{
		TotalCount: 0,
		Page:       input.Page,
		Limit:      input.Limit,
	}
	db := postgres.DB.Db

	if err := db.Model(&models.User{}).Count(&totalCount).Error; err != nil {
		err := utils.FetchDBError(err.Error())
		return nil, &meta, &err
	}

	meta.TotalCount = totalCount

	result := db.Limit(input.Limit).Offset((input.Page - input.Limit) * input.Limit).Find(&users)
	if result.Error != nil {
		err := utils.FetchDBError(result.Error.Error())
		return nil, &meta, &err
	}

	return &users, &meta, nil
}

func GetOneByID(req dto.GetOneByIDRequest) (*models.User, error) {
	var user models.User
	if err := postgres.DB.Db.First(&user, "id = ?", req.ID).Error; err != nil {
		log.Println("Error fetching user by ID:", err)
		return nil, err
	}
	return &user, nil
}

func Create(user *dto.CreateOne) (*models.User, *baseDTO.Error) {
	hashedPassword, errHash := utils.HashPassword(user.Password)
	if errHash != nil {
		err := utils.HashError(errHash.Error())
		return nil, &err
	}

	// Set the password to the hashed version
	user.Password = hashedPassword
	newUser := models.User{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
	}

	if errFetch := postgres.DB.Db.Create(&newUser).Error; errFetch != nil {
		err := utils.FetchDBError(errFetch.Error())
		return nil, &err
	}
	return &newUser, nil
}
