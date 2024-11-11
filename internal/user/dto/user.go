package dto

type (
	GetOneByIDRequest struct {
		ID int
	}
	CreateOne struct {
		Username string
		Email    string
		Password string
	}
)
