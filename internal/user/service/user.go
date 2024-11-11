package service

import (
	"github.com/joshuaautawi/go-api/internal/user/dto"
	"github.com/joshuaautawi/go-api/internal/user/models"
	"github.com/joshuaautawi/go-api/internal/user/repository"
)

func CreateOne(req *dto.CreateOne) (*models.User, error) {
	user, err := repository.Create(req)
	return user, err
}
