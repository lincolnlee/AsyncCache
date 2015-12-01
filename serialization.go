package async_cache

import (
	"gopkg.in/vmihailenco/msgpack.v2"
)

type Seria struct{}

func (*Seria) Serialize(v ...interface{}) []byte {
	b, err := msgpack.Marshal(v)
	if err != nil {
		panic(err)
	}

	return b
}

func (*Seria) DeserializeToSlice(b []byte) []interface{} {
	var out []interface{}
	err = msgpack.Unmarshal(b, &out)
	if err != nil {
		panic(err)
	}
	return out
}
