package utils

import (
	"log"
	"reflect"

	admin "github.com/joshuaautawi/go-api/internal/admin/models"
	user "github.com/joshuaautawi/go-api/internal/user/models"

	"gorm.io/gorm"
)

func MigrateModels(db *gorm.DB) {
	// List of models to migrate
	models := []interface{}{
		&user.User{},
		&admin.Admin{},
		// Add more models here as your application grows
	}

	// Automatically migrate all models
	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			log.Fatalf("Error migrating model %v: %v\n", reflect.TypeOf(model), err)
		}
		log.Printf("Successfully migrated model: %v\n", reflect.TypeOf(model))
	}
}
