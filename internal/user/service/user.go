package service

import (
	baseDTO "github.com/joshuaautawi/go-api/internal/common/dto"
	"github.com/joshuaautawi/go-api/internal/user/dto"
	"github.com/joshuaautawi/go-api/internal/user/models"
	"github.com/joshuaautawi/go-api/internal/user/repository"
)

func CreateOne(req *dto.CreateOne) (*models.User, error) {
	user, err := repository.Create(req)
	return user, err
}

func GetAll(input *baseDTO.GetAllRequest) (*baseDTO.GetAllBaseResponse[[]models.User], error) {
	users, err := repository.GetAll(input)
	return users, err
}
