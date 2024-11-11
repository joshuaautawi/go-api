package models

import (
	"gorm.io/gorm"
)

// User struct
type Admin struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Users struct
type Admins struct {
	Users []Admin `json:"users"`
}
