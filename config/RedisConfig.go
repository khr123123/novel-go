package config

import (
	"context"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	cacheRedis "github.com/gomodule/redigo/redis" // 用于 gin-contrib/cache
	"github.com/redis/go-redis/v9"                // 这是另一套 Redis 客户端（可用于业务逻辑）
	"time"
)

// 通用上下文对象，可用于 go-redis 的方法调用
var Ctx = context.Background()

// RedisClient 是 go-redis 官方客户端，主要用于普通 Redis 操作（非缓存中间件）
// 用于业务中其他逻辑，如存 Token、记录 PV 等
var RedisClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379", // Redis 地址
	Password: "",               // Redis 无密码时留空
	DB:       6,                // 使用第 6 个 Redis 库（0-15）
})

// NewRedisPool 返回一个 redigo Redis 连接池
// gin-contrib/cache 使用的是 redigo，所以我们在这里初始化 redigo 连接池
func NewRedisPool() *cacheRedis.Pool {
	return &cacheRedis.Pool{
		MaxIdle:     10,                // 最大空闲连接数
		IdleTimeout: 240 * time.Second, // 空闲连接最大存活时间
		Dial: func() (cacheRedis.Conn, error) {
			// 连接 Redis（端口、地址根据实际修改）
			return cacheRedis.Dial("tcp", "localhost:6379", cacheRedis.DialDatabase(0))
		},
		TestOnBorrow: func(c cacheRedis.Conn, t time.Time) error {
			// 每次从连接池取出连接时做一次 PING 测试
			_, err := c.Do("PING")
			return err
		},
	}
}

// RedisCacheHandler 返回一个带 Redis 缓存的 gin handler（中间件）
// 传入 handler 以及缓存有效时间，自动生成缓存包装函数
//
// 参数说明：
// - handler: 你想缓存的路由 handler 方法
// - cacheDuration: 缓存时间（如 time.Minute）
//
// 内部使用 Redis 存储缓存内容，缓存时间为 cacheDuration，
// Redis 中数据的 TTL 是 60 秒（NewRedisCacheWithPool 第二个参数）
func RedisCacheHandler(handler gin.HandlerFunc, cacheDuration time.Duration) gin.HandlerFunc {
	// 创建 Redis 缓存存储器（底层使用 redigo）
	store := persistence.NewRedisCacheWithPool(NewRedisPool(), 60*time.Second)

	// 使用 gin-contrib/cache 中间件封装 handler，启用页面级缓存
	return cache.CachePage(store, cacheDuration, handler)
}

// LocalCacheHandler 是一个基于内存缓存的 Gin 中间件封装函数。
// 它接收一个原始的 gin.HandlerFunc 以及缓存时间，返回一个带缓存功能的 HandlerFunc。
//
// 参数说明：
// - handler: 需要缓存的原始处理函数（请求处理逻辑）
// - cacheDuration: 缓存的有效时长，即同一个请求的响应结果在缓存中保存多久
//
// 内部实现：
//   - 使用 gin-contrib/cache 提供的内存缓存实现 persistence.NewInMemoryStore，
//   - 调用 cache.CachePage 把 handler 包装成支持缓存的 HandlerFunc，
//     在缓存有效期内，同样的请求会直接返回缓存响应，提升性能。
func LocalCacheHandler(handler gin.HandlerFunc, cacheDuration time.Duration) gin.HandlerFunc {
	store := persistence.NewInMemoryStore(60 * time.Second)
	return cache.CachePage(store, cacheDuration, handler)
}
