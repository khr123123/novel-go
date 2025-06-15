package config

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

var RedisClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379", // 改成你的地址
	Password: "",               // 无密码填空字符串
	DB:       6,
})
