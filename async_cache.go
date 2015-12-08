package AsyncCache

import (
	"log"
	"time"
)

type cacheHandler struct{}

func (*cacheHandler) AsyncGetAndUpdateData(f func() interface{}, key string) interface{} {
	var cacheValue interface{} = nil
	if v, err := InstanceContainer.redisClient.GetBytesSlice(key); v != nil && err == nil {
		iSlice := InstanceContainer.serializer.DeserializeToSlice(v)
		if len(iSlice) == 2 {
			cacheValue = iSlice[1]
			if time.Now().Sub(Time(iSlice[0])).Minutes() > 5 {
				go asyncDealCacheTask(f, key, cacheValue, true)
			}

			return cacheValue
		}
	} else {
		log.Println(err)

	}
	cacheValue = f()
	go asyncDealCacheTask(f, key, cacheValue, false)
	return cacheValue
}

func (*cacheHandler) asyncDealCacheTask(f func() interface{}, key string, v interface{}, isDoF bool) interface{} {
	if getLock(key) {
		var r interface{} = v
		if isDoF {
			r = f()
		}
		iSlice = []interface{}{time.Now(), r}
		cacheData := InstanceContainer.serializer.Serialize(iSlice)
		InstanceContainer.redisClient.SetBytesSliceWithExpriePX(key, cacheData, 300*1000)
	}

}

func (*cacheHandler) getLock(key string) bool {
	lockKey := key + "_lock"
	if ok, _ := InstanceContainer.redisClient.SetNXInt(lockKey, 1); ok {
		InstanceContainer.redisClient.SetStringWithExpriePX(lockKey, "1", 60)
		return true
	} else {
		return false
	}
}

func (*cacheHandler) releaseLock(key string) {
	lockKey := key + "_lock"
	InstanceContainer.redisClient.RemoveKey(lockKey)
}
