package database

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/gomodule/redigo/redis"
)

var RedisPool *redis.Pool

func init() {
	SetupRedisPool()
}

func SetupRedisPool() {
	pool := &redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial:      getRedisUrl,
	}

	RedisPool = pool
}

func getRedisUrl() (redis.Conn, error) {
	redisUrl, err := web.AppConfig.String("redisUrl")
	if err != nil {
		logs.Critical("Redis url not found: ", err)
	}

	return redis.Dial("tcp", redisUrl)
}

func GetRedisPool() *redis.Pool {
	return RedisPool
}
