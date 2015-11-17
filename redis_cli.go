package redis_cli

import (
	"github.com/garyburd/redigo/redis"
	"log"
)

type Rediscli struct {
	pool *redis.Pool
}

// 生成连接池方法
func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "192.168.0.1:6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

// 生成连接池
var Rc = &Rediscli{newPool()}

//存储字段串到redis
func (rediscli *Rediscli) SetString(key string, v string) error {
	// 从连接池里面获得一个连接
	c := rediscli.pool.Get()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()

	if ok, err := redis.Bool(c.Do("SET", key, v)); ok {
		return nil
	} else {
		log.Print(err)
		return err
	}
}
