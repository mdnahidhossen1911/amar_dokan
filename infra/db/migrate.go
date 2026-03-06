package db

import (
	"amar_dokan/models"
	"fmt"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{}, &models.PandingUser{}, &models.Product{}, &models.AddToCart{}, &models.Category{})
	if err != nil {
		return err
	}
	fmt.Println("✅ Migrations applied")
	return nil
}
