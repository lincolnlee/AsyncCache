package async_cache

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

//SET String to redis
func (this *Rediscli) SetString(key string, v string) error {
	// 从连接池里面获得一个连接
	c := this.pool.Get()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()

	if ok, err := redis.Bool(c.Do("SET", key, v)); ok {
		return nil
	} else {
		log.Print(err)
		return err
	}
}

//SET String to redis with expire time.Set the specified expire time, in milliseconds.
func (this *Rediscli) SetStringWithExpriePX(key string, v string, exprie int32) error {
	// 从连接池里面获得一个连接
	c := this.pool.Get()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()

	if ok, err := redis.Bool(c.Do("SET", key, v, "PX", exprie)); ok {
		return nil
	} else {
		log.Print(err)
		return err
	}
}

//SETNX Int32 to redis
func (this *Rediscli) SetNXInt(key string, v int32) error {
	// 从连接池里面获得一个连接
	c := this.pool.Get()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()

	if ok, err := redis.Bool(c.Do("SETNX", key, v)); ok {
		return nil
	} else {
		log.Print(err)
		return err
	}
}

//SET []byte to redis
func (this *Rediscli) SetBytesSlice(key string, v []byte) error {
	// 从连接池里面获得一个连接
	c := this.pool.Get()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()

	if ok, err := redis.Bool(c.Do("SET", key, v)); ok {
		return nil
	} else {
		log.Print(err)
		return err
	}
}

//SET interface{} to redis
func (this *Rediscli) SetInterface(key string, v interface{}) error {
	// 从连接池里面获得一个连接
	c := this.pool.Get()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()

	if ok, err := redis.Bool(c.Do("SET", key, v)); ok {
		return nil
	} else {
		log.Print(err)
		return err
	}
}

//GET String from redis
func (this *Rediscli) GetString(key string) (v string, err error) {
	// 从连接池里面获得一个连接
	c := this.pool.Get()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()

	if v, err = redis.String(c.Do("GET", key)); err != nil {
		log.Print(err)
		return nil, err
	} else {
		return v, err
	}
}

//GET []byte from redis
func (this *Rediscli) GetBytesSlice(key string) (v []byte, err error) {
	// 从连接池里面获得一个连接
	c := this.pool.Get()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()

	if v, err = redis.Bytes(c.Do("GET", key)); err != nil {
		log.Print(err)
		return nil, err
	} else {
		return v, err
	}
}
