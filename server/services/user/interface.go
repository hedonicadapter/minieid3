package user

import "github.com/hedonicadapter/gopher/models"

type UserService interface {
	Get(id string) (models.User, error)
}
