package async_cache

import (
	"log"
)

func AsyncGetAndUpdateData(f func(...interface{}) interface{}, key string) interface{} {
	if v, err := Rc.GetBytesSlice(key); v == nil || err != nil {
		log.Println(err)
		return f()
	} else {
		iSlice := Seri.DeserializeToSlice(v)
		if len(iSlice) == 2 {
			return iSlice[1]
		} else {
			return f()
		}

	}
}
