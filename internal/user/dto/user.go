package dto

type (
	GetOneByIDRequest struct {
		ID int
	}
	CreateOne struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=6"`
	}
)
