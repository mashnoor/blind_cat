package settings

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

var (
	rdb *redis.Client
)

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", RedisHost, RedisPort),
	})

}

func GetRedisClient() *redis.Client {
	return rdb
}
