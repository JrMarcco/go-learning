package gredis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"go_learning/gin_example/pkg/logging"
	"go_learning/gin_example/pkg/setting"
	"time"
)

var RedisConn *redis.Pool

func Setup() {
	RedisConn = &redis.Pool{
		MaxActive:   setting.RedisSetting.MaxActive,
		MaxIdle:     setting.RedisSetting.MaxIdle,
		IdleTimeout: setting.RedisSetting.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", setting.RedisSetting.Host)
			if err != nil {
				return nil, err
			}

			if setting.RedisSetting.Password != "" {
				if _, err := c.Do("AUTH", setting.RedisSetting.Password); err != nil {
					c.Close()
					return nil, err
				}

			}

			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func Set(key string, data any, time int) error {
	conn := RedisConn.Get()
	defer func(conn redis.Conn) {
		if err := conn.Close(); err != nil {
			logging.Error(err)
		}
	}(conn)

	val, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if _, err = conn.Do("SET", key, val); err != nil {
		return err
	}

	if _, err = conn.Do("EXPIRE", key, time); err != nil {
		return err
	}

	return nil
}

func Exists(key string) bool {
	conn := RedisConn.Get()
	defer func(conn redis.Conn) {
		if err := conn.Close(); err != nil {
			logging.Error(err)
		}
	}(conn)

	exist, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}
	return exist
}

func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer func(conn redis.Conn) {
		if err := conn.Close(); err != nil {
			logging.Error(err)
		}
	}(conn)

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer func(conn redis.Conn) {
		if err := conn.Close(); err != nil {
			logging.Error(err)
		}
	}(conn)

	return redis.Bool(conn.Do("DEL", key))
}
