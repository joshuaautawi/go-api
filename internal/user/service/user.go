package service

import (
	baseDTO "github.com/joshuaautawi/go-api/internal/common/dto"
	"github.com/joshuaautawi/go-api/internal/user/dto"
	"github.com/joshuaautawi/go-api/internal/user/models"
	"github.com/joshuaautawi/go-api/internal/user/repository"
)

func CreateOne(req *dto.CreateOne) (*models.User, *baseDTO.Error) {
	user, err := repository.Create(req)
	return user, err
}

func GetAll(input *baseDTO.GetAllRequest) (*[]models.User, *baseDTO.Meta, *baseDTO.Error) {
	users, meta, err := repository.GetAll(input)
	return users, meta, err
}
