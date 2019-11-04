package db

import (
	"errors"
	"github.com/gomodule/redigo/redis"
	"demo/utility/helper"
	"time"
)

var redisPool *redis.Pool

// 初始化redis
func InitRedis(conf map[string]string) (err error) {
	if !helper.MapSKeysExists(conf, []string{"addr", "max_idle", "max_open", "db"}) {
		return errors.New("redis config is missing")
	}

	redisPool = &redis.Pool{
		MaxIdle:     helper.StrToInt(conf["max_idle"], 1),
		MaxActive:   helper.StrToInt(conf["max_open"], 10),
		IdleTimeout: time.Duration(30) * time.Minute,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", conf["addr"], redis.DialDatabase(helper.StrToInt(conf["db"], 0)))
		},
	}

	conn := GetRedis()
	defer conn.Close()

	if r, _ := redis.String(conn.Do("PING")); r != "PONG" {
		err = errors.New("redis connect failed.")
	}

	return
}

// 获取redis连接
func GetRedis() redis.Conn {
	return redisPool.Get()
}

// 关闭redis
func CloseRedis() {
	if redisPool != nil {
		redisPool.Close()
	}
}

