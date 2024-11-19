package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joshuaautawi/go-api/configs"
	"github.com/joshuaautawi/go-api/internal/common/dto"
)

var secret = configs.Config("JWT_SECRET")
var jwtSecret = []byte(secret)

func GenerateJWTToken(id int) (string, *dto.Error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, tokenErr := token.SignedString(jwtSecret)
	if tokenErr != nil {
		err := JWTError(tokenErr.Error())
		return "", &err
	}
	return tokenString, nil
}
