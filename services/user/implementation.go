package user

import (
	"context"
	"github.com/hedonicadapter/gopher/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct {
	Db *pgxpool.Pool
}

func (s Service) Get(id string) (models.User, error) {
	res, err := s.Db.Query(context.Background(), "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return models.User{Id: id}, err
	}

	user, err := pgx.CollectExactlyOneRow(res, pgx.RowToStructByName[models.User])
	return user, err
}

func InitService(db *pgxpool.Pool) Service {
	return Service{Db: db}
}
