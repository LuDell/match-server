package utils

import (
	"github.com/gomodule/redigo/redis"
	"time"
	"match-server/model"
)

var RedisPool *redis.Pool

func init()  {
	RedisPool = globalPool()
}

func globalPool() *redis.Pool {
	config := model.SeeLogConfig.Redis
	pool := &redis.Pool{
		MaxIdle: config.Max_idle,
		MaxActive: config.Max_active,
		IdleTimeout:6000 *time.Second,
		Dial: func() (redis.Conn, error) {
			con,err1 :=redis.Dial("tcp",config.Tcp+":"+config.Port)
			if err1 != nil {
				panic(err1)
			}
			_,err2 := con.Do("AUTH",config.Password)
			if err2 != nil {
				panic(err2)
			}
			return con,err1
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_,err := c.Do("PING")
			return err
		},
	}
	return pool
}