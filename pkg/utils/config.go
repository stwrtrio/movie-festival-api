package utils

import (
	"os"
	"strconv"
)

func GetRedisCacheTTL() int {
	ttlMinutes, err := strconv.Atoi(os.Getenv("REDIS_CACHE_TTL_MINUTES"))
	if err != nil {
		ttlMinutes = 5 // default 5 minutes
	}

	return ttlMinutes
}
