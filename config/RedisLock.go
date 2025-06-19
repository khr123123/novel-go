package config

import (
	"github.com/go-redsync/redsync/v4"

	redsyncredis "github.com/go-redsync/redsync/v4/redis/goredis/v9" // 注意 v9 适配器
)

var Redsync *redsync.Redsync

func InitRedisLock() {
	pool := redsyncredis.NewPool(RedisClient)
	Redsync = redsync.New(pool)
}
