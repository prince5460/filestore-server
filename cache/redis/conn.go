package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

var (
	pool      *redis.Pool
	redisHost = "127.0.0.1:6379"
	redisPass = ""
)

//newRedisPool:创建redis连接池
func newRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     50,                //最大空闲连接数
		MaxActive:   30,                //客户端与数据库的最大连接数,0表示没有限制
		IdleTimeout: 300 * time.Second, //最大空闲时间
		Dial: func() (conn redis.Conn, e error) {
			//1.打开连接
			dial, err := redis.Dial("tcp", redisHost)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			//2.访问认证
			if _, err := dial.Do("AUTH", redisPass); err != nil {
				dial.Close()
				return nil, err
			}
			return dial, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func init() {
	pool = newRedisPool()
}

func RedisPool() *redis.Pool {
	return pool
}
