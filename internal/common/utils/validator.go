package utils

import (
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/joshuaautawi/go-api/internal/common/dto"
)

var (
	once              sync.Once
	validatorInstance *validator.Validate
)

// GetValidator returns a singleton instance of the validator
func GetValidator() *validator.Validate {
	once.Do(func() {
		validatorInstance = validator.New()
	})
	return validatorInstance
}

func HandleValidation[T any](req T) *dto.Error {
	if validationErr := GetValidator().Struct(req); validationErr != nil {
		err := ValidationError(validationErr.Error())
		return &err
	}
	return nil
}
