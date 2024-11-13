package models

import (
	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

// Users struct
type Users struct {
	Users []User `json:"users"`
}
