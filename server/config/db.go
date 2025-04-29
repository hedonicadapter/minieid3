package config

import (
	"fmt"
	"os"

	"github.com/hedonicadapter/gopher/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&models.User{})
	return db
}

// INFO: använde ai
func IdempotentDummyData(db *gorm.DB) *gorm.DB {
	var count int64
	if err := db.Model(&models.User{}).Count(&count).Error; err != nil {
		fmt.Println("Failed to check for existing dummy data: ", err.Error())
		os.Exit(1)
	}

	if count == 0 {
		dummyUsers := []models.User{
			{Name: "Alice"},
			{Name: "Bob"},
			{Name: "Charlie"},
		}

		if err := db.Create(&dummyUsers).Error; err != nil {
			fmt.Println("Failed to seed dummy data: ", err.Error())
			os.Exit(1)
		} else {
			fmt.Println("Dummy data seeded successfully.")
		}
	} else {
		fmt.Println("Dummy data already exists. Skipping seeding.")
	}

	return db
}

func InitDb() *gorm.DB {
	dsn := os.Getenv("POSTGRES_DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	migrate(db)
	return db
}
