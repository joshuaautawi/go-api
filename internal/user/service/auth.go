package service

import (
	baseDTO "github.com/joshuaautawi/go-api/internal/common/dto"
	"github.com/joshuaautawi/go-api/internal/user/dto"
	"github.com/joshuaautawi/go-api/internal/user/models"
	"github.com/joshuaautawi/go-api/internal/user/repository"
)

func Login(req *dto.LoginRequest) (*models.User, *baseDTO.Error) {
	user, err := repository.Login(req)
	return user, err
}
