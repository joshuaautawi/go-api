package repository

import (
	baseDTO "github.com/joshuaautawi/go-api/internal/common/dto"
	"github.com/joshuaautawi/go-api/internal/common/utils"
	"github.com/joshuaautawi/go-api/internal/user/dto"
	"github.com/joshuaautawi/go-api/internal/user/models"
	"github.com/joshuaautawi/go-api/pkg/db/postgres"
)

func Login(req *dto.LoginRequest) (*models.User, *baseDTO.Error) {

	var user models.User
	if fetchErr := postgres.DB.Db.First(&user, "username= ?", req.Username).Error; fetchErr != nil {
		err := utils.FetchDBError(fetchErr.Error())
		return nil, &err
	}
	isSamePassword := utils.ComparePasswords(req.Password, user.Password)
	if !isSamePassword {
		err := utils.WrongPasswordError()
		return nil, &err
	}

	return &user, nil
}
