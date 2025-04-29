package queue

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Service struct {
	Rdb *redis.Client
}

func (s Service) Enqueue() {
	fmt.Println("enqueued a chungus")
}

func InitService(rdb *redis.Client) Service {
	return Service{Rdb: rdb}
}
