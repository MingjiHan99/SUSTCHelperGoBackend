package utils

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

const (
	RedisURL            = "redis://127.0.0.1:6379"
	redisMaxIdle        = 3   //最大空闲连接数
	redisIdleTimeoutSec = 240 //最大空闲连接时间
)

// NewRedisPool 返回redis连接池
func NewRedisPool(redisURL string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     redisMaxIdle,
		IdleTimeout: redisIdleTimeoutSec * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(redisURL)
			if err != nil {
				return nil, fmt.Errorf("redis connection error: %s", err)
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				return fmt.Errorf("ping redis error: %s", err)
			}
			return nil
		},
	}
}

var redisPool = NewRedisPool(RedisURL)

func Set(k, v string) {
	c := redisPool.Get()
	defer c.Close()
	_, err := c.Do("SET", k, v)
	if err != nil {
		fmt.Println("set error", err.Error())
	}
}

func GetStringValue(k string) string {
	c := redisPool.Get()
	defer c.Close()
	username, err := redis.String(c.Do("GET", k))
	if err != nil {
		fmt.Println("Get Error: ", err.Error())
		return ""
	}
	return username
}

func SetKeyExpire(k string, ex int) {
	c := redisPool.Get()
	defer c.Close()
	_, err := c.Do("EXPIRE", k, ex)
	if err != nil {
		fmt.Println("set error", err.Error())
	}
}

func CheckKey(k string) bool {
	c := redisPool.Get()
	defer c.Close()
	exist, err := redis.Bool(c.Do("EXISTS", k))
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return exist
	}
}

func DelKey(k string) error {
	c := redisPool.Get()
	defer c.Close()
	_, err := c.Do("DEL", k)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
