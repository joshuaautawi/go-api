package repository

import (
	"github.com/joshuaautawi/go-api/internal/user/models"
	"github.com/joshuaautawi/go-api/pkg/db/postgres"
)

func GetAll() models.Users {
	var users models.Users
	db := postgres.DB.Db
	db.Find(&users)
	return users
}

func GetOneByID() models.User {
	var user models.User
	db := postgres.DB.Db
	db.Find(&user)
	return user
}
