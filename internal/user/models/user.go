package models

import (
	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Users struct
type Users struct {
	Users []User `json:"users"`
}
