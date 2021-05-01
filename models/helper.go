package models

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func SaveUrlRecord(uniqueKey, longURL string) *redis.StatusCmd {
	h, _ := time.ParseDuration("720hours")
	status := DB.Set(uniqueKey, longURL, time.Duration(h.Hours()))
	return status
}

func IsUniqueKeyAlreadyUsed(uniqueKey uint64) bool {
	result := DB.Exists(fmt.Sprint(uniqueKey))
	var isAlreadyUsed bool = int(result.Val()) != 0
	return isAlreadyUsed
}
