package config

import (
	"errors"
	"os"

	"github.com/hedonicadapter/gopher/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&models.User{})
	return db
}

func CreateDummyData(db *gorm.DB) *gorm.DB {
	errors.New("not implemented")
	os.Exit(1)

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
