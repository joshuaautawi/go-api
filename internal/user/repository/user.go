package repository

import (
	"log"
	"time"

	baseDTO "github.com/joshuaautawi/go-api/internal/common/dto"
	"github.com/joshuaautawi/go-api/internal/common/utils"
	"github.com/joshuaautawi/go-api/internal/user/dto"
	"github.com/joshuaautawi/go-api/internal/user/models"
	"github.com/joshuaautawi/go-api/pkg/db/postgres"
	"gorm.io/gorm"
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

func GetOneByID(id int) (*models.User, *baseDTO.Error) {
	var user models.User
	if fetchErr := postgres.DB.Db.First(&user, "id = ?", id).Error; fetchErr != nil {
		log.Println("Error fetching user by ID:", fetchErr)
		err := utils.FetchDBError(fetchErr.Error())
		return nil, &err
	}
	return &user, nil
}

func CreateOne(user *dto.CreateOneRequest) (*models.User, *baseDTO.Error) {
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

func UpdateOne(input *dto.UpdateOneRequest) (*models.User, *baseDTO.Error) {
	db := postgres.DB.Db
	user, err := GetOneByID(input.ID)
	if err != nil {
		return nil, err
	}
	user.Username = input.Username
	db.Save(&user)
	return user, nil
}

func DeleteOne(id int) (*models.User, *baseDTO.Error) {
	db := postgres.DB.Db
	user, err := GetOneByID(id)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	user.DeletedAt = gorm.DeletedAt{Time: now, Valid: true}
	db.Save(&user)
	return user, nil
}
