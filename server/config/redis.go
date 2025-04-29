package config

import (
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	redisDbUrl := os.Getenv("REDIS_DATABASE_URL")

	opt, err := redis.ParseURL(redisDbUrl)
	if err != nil {
		fmt.Println("rdb error: ", err.Error())
		os.Exit(1)
	}

	return redis.NewClient(opt)
}
