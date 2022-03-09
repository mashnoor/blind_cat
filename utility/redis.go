package utility

import (
	"context"
	"github.com/mashnoor/blind_cat/settings"
	"strconv"
)

func RedisHSet(serviceName string, param string, value int) {
	client := settings.GetRedisClient()

	client.HSet(context.Background(), serviceName, param, value)
}

func RedisHGet(serviceName string, param string) int {
	client := settings.GetRedisClient()
	val, err := client.HGet(context.Background(), serviceName, param).Result()
	if err != nil {
		panic(err)

	}
	i, err := strconv.Atoi(val)
	if err != nil {
		return 0
	}
	return i
}
