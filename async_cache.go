package AsyncCache

import (
	"log"
)

type cacheHandler struct{}

func (*cacheHandler) AsyncGetAndUpdateData(f func(...interface{}) interface{}, key string) interface{} {
	if v, err := InstanceContainer.redisClient.GetBytesSlice(key); v == nil || err != nil {
		log.Println(err)
		return f()
	} else {
		iSlice := InstanceContainer.serializer.DeserializeToSlice(v)
		if len(iSlice) == 2 {
			return iSlice[1]
		} else {
			return f()
		}

	}
}
