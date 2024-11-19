package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joshuaautawi/go-api/internal/common/dto"
)

func ParseError(message string) dto.Error {
	err := dto.Error{Code: fiber.StatusBadRequest, InternalCode: "102", Message: message}
	return err
}

func FetchDBError(message string) dto.Error {
	err := dto.Error{Code: fiber.StatusInternalServerError, InternalCode: "10002", Message: message}
	return err
}

func HashError(message string) dto.Error {
	err := dto.Error{Code: fiber.StatusInternalServerError, InternalCode: "10002", Message: message}
	return err
}

func JWTError(message string) dto.Error {
	err := dto.Error{Code: fiber.StatusBadRequest, InternalCode: "10002", Message: message}
	return err
}

func ValidationError(message string) dto.Error {
	err := dto.Error{Code: fiber.StatusBadRequest, InternalCode: "10002", Message: message}
	return err
}

func WrongPasswordError() dto.Error {
	err := dto.Error{Code: fiber.StatusForbidden, InternalCode: "10002", Message: "wrong password"}
	return err
}

func JWTMiddlewareError(message string) dto.Error {
	err := dto.Error{Code: fiber.StatusForbidden, InternalCode: "10002", Message: message}
	return err
}
