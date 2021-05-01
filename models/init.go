package models

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

// DB pointer
var DB *redis.Client

func InitRedisDB() (*redis.Client, error) {

	redisAddress := fmt.Sprintf("%s:%s", os.Getenv("REDIS_SERVER"), os.Getenv("REDIS_DB_PORT"))

	client := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	err := client.Ping().Err()
	if err != nil {
		return nil, err
	}

	DB = client
	return DB, nil
}
