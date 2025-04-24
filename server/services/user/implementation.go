package user

import (
	"github.com/hedonicadapter/gopher/models"
	"gorm.io/gorm"
)

type Service struct {
	Db *gorm.DB
}

func (s Service) Get(id string) (models.User, error) {
	var user models.User
	err := s.Db.First(&user, "id = ?", id).Error
	return user, err
}

func InitService(db *gorm.DB) Service {
	return Service{Db: db}
}
