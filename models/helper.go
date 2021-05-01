package models

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func SaveUrlRecord(uniqueKey, longURL string) *redis.StatusCmd {
	status := DB.Set(uniqueKey, longURL, 5*time.Minute)
	return status
}

func IsUniqueKeyAlreadyUsed(uniqueKey uint64) bool {
	result := DB.Exists(fmt.Sprint(uniqueKey))
	var isAlreadyUsed bool = int(result.Val()) != 0
	return isAlreadyUsed
}
