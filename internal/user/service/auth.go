package service

import (
	baseDTO "github.com/joshuaautawi/go-api/internal/common/dto"
	"github.com/joshuaautawi/go-api/internal/common/utils"
	"github.com/joshuaautawi/go-api/internal/user/dto"
	"github.com/joshuaautawi/go-api/internal/user/repository"
)

func Login(req *dto.LoginRequest) (string, *baseDTO.Error) {
	user, err := repository.Login(req)
	if err != nil {
		return "", err
	}
	token, err := utils.GenerateJWTToken(int(user.ID))
	return token, err
}
