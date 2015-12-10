package AsyncCache

import (
	"github.com/going/toolkit/to"
	"sync"
	"time"
)

type cacheHandler struct {
	enabledCache bool
	mutex        sync.Mutex
	errAmount    int32
}

func (this *cacheHandler) AsyncGetAndUpdateData(f func() interface{}, key string) interface{} {
	var cacheValue interface{} = nil
	if v, err := InstanceContainer.redisClient.GetBytesSlice(key); v != nil && err == nil {
		iSlice := InstanceContainer.serializer.DeserializeToSlice(v)
		if len(iSlice) == 2 {
			cacheValue = iSlice[1]
			if time.Now().Sub(to.Time(iSlice[0])).Minutes() > 5 {
				go this.asyncDealCacheTask(f, key, cacheValue, true)
			}

			return cacheValue
		}
	} else {
		InstanceContainer.Loghelper.Error(err)

	}
	cacheValue = f()
	go this.asyncDealCacheTask(f, key, cacheValue, false)
	return cacheValue
}

func (this *cacheHandler) asyncDealCacheTask(f func() interface{}, key string, v interface{}, isDoF bool) {
	if this.getLock(key) {
		var r interface{} = v
		if isDoF {
			r = f()
		}
		iSlice := []interface{}{time.Now(), r}
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

func (this *cacheHandler) healthDetect() {
	this.errAmount++
	if this.errAmount > 100 && this.enabledCache {
		this.mutex.Lock()
		if this.errAmount > 100 && this.enabledCache {
			go func() {
				this.enabledCache = true
			}()
		}
	}

}
