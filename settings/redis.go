package settings

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

var (
	rdb *redis.Client
)

func InitRedis(host, port string) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", host, port),
	})

}

func GetRedisClient() *redis.Client {
	return rdb
}
