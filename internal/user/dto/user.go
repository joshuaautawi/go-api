package dto

type (
	CreateOneRequest struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=6"`
	}

	UpdateOneRequest struct {
		ID       int    `json:"id" validate:"required"`
		Username string `json:"username" validate:"required"`
	}
)
