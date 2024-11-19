package service

import (
	baseDTO "github.com/joshuaautawi/go-api/internal/common/dto"
	"github.com/joshuaautawi/go-api/internal/user/dto"
	"github.com/joshuaautawi/go-api/internal/user/models"
	"github.com/joshuaautawi/go-api/internal/user/repository"
)

func GetAll(input *baseDTO.GetAllRequest) (*[]models.User, *baseDTO.Meta, *baseDTO.Error) {
	users, meta, err := repository.GetAll(input)
	return users, meta, err
}

func GetOne(id int) (*models.User, *baseDTO.Error) {
	user, err := repository.GetOneByID(id)
	return user, err
}

func CreateOne(req *dto.CreateOneRequest) (*models.User, *baseDTO.Error) {
	user, err := repository.CreateOne(req)
	return user, err
}

func UpdateOne(input *dto.UpdateOneRequest) (*models.User, *baseDTO.Error) {
	user, err := repository.UpdateOne(input)
	return user, err
}

func DeleteOne(id int) (*models.User, *baseDTO.Error) {
	user, err := repository.DeleteOne(id)
	return user, err
}
