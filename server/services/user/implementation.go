package user

import (
	"github.com/hedonicadapter/gopher/models"
	"gorm.io/gorm"
)

type Service struct {
	Db *gorm.DB
}

func (s Service) GetById(id string) (*models.User, error) {
	var user models.User
	err := s.Db.First(&user, "id = ?", id).Error

	return &user, err
}
func (s Service) Create(user models.User) (models.User, error) {
	err := s.Db.Create(&user).Error

	return user, err
}
func (s Service) Delete(id string) error {
	err := s.Db.Delete(&models.User{}, id).Error

	return err
}
func (s Service) List() ([]models.User, error) {
	var user []models.User
	res := s.Db.Find(&user)

	return user, res.Error
}
func (s Service) Update(id string, user models.User) (uint, error) {
	usr, err := s.GetById(id)
	if err != nil {
		return user.ID, err
	}

	usr.Name = user.Name
	updatedUser := s.Db.Save(usr)

	// TODO: why not
	// return strconv.Itoa(*usr.ID), updatedUser.Error
	return usr.ID, updatedUser.Error
}

func InitService(db *gorm.DB) Service {
	return Service{Db: db}
}
