package user

import "github.com/hedonicadapter/gopher/models"

type UserService interface {
	GetById(id string) (*models.User, error)
	List() ([]models.User, error)
	Create(models.User) (models.User, error)
	Update(id string, user models.User) (uint, error)
	Delete(id string) error
}
