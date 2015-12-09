package AsyncCache

type container struct {
	serializer        *serializer
	redisClient       *rediscli
	AsyncCacheHandler *cacheHandler
	Exception         *exception
	Loghelper         *loghelper
}

var InstanceContainer container = container{
	serializer:        &serializer{},           //serialization.go
	redisClient:       &rediscli{newPool()},    // 生成Redis连接池 //redis_cli.go
	AsyncCacheHandler: &cacheHandler{},         //async_cache.go
	Exception:         &exception{},            //exception.go
	Loghelper:         &loghelper{newLogger()}, //loghelper.go
}
