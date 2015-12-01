package async_cache

type Container struct {
	Serializer *Seria
	Rediscli   *Rediscli
}

var InstanceContainer Container = Container{
	Serializer:  &Seria{},             //serialization.go
	RedisClient: &Rediscli{newPool()}, // 生成Redis连接池 //redis_cli.go
}
