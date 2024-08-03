package dbcli

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

var redisCli *redis.Client

func InitRedis(user, pass, host string, port, db int) error {
	redisCli = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", host, port),
		Password: pass,
		DB: db,
	})

	return nil
}

func Rdb() *redis.Client {
	return redisCli
}
