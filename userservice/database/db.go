package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"userservice/models"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) {
	Instance, dbError = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")
}
func Migrate() {
	Instance.AutoMigrate(&models.User{})
	log.Println("Database Migration Completed!")
}

func CreateRoleAllowedEnum() error {
	result := Instance.Exec("SELECT 1 FROM pg_type WHERE typname = 'role_allowed';")

	switch {
	case result.RowsAffected == 0:
		if err := Instance.Exec("CREATE TYPE role_allowed AS ENUM ('regular', 'admin', 'rider', 'seller');").Error; err != nil {
			log.Fatal("Error creating role_allowed ENUM")
			return err
		}

		return nil
	case result.Error != nil:
		return result.Error

	default:
		return nil
	}
}

func CreateStatusAllowedEnum() error {
	result := Instance.Exec("SELECT 1 FROM pg_type WHERE typname = 'status_allowed';")

	switch {
	case result.RowsAffected == 0:
		if err := Instance.Exec("CREATE TYPE status_allowed AS ENUM ('suspended', 'active', 'inactive');").Error; err != nil {
			log.Fatal("Error creating status_allowed ENUM")
			return err
		}

		return nil
	case result.Error != nil:
		return result.Error

	default:
		return nil
	}
}
