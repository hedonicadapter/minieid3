package config

import (
	"os"

	"github.com/hedonicadapter/gopher/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}

func InitDb() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	migrate(db)
	return db
}
