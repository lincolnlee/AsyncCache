package AsyncCache

import (
	"gopkg.in/vmihailenco/msgpack.v2"
)

type serializer struct{}

func (*serializer) Serialize(v ...interface{}) []byte {
	b, err := msgpack.Marshal(v)
	if err != nil {
		panic(err)
	}

	return b
}

func (*serializer) DeserializeToSlice(b []byte) []interface{} {
	var out []interface{}
	err := msgpack.Unmarshal(b, &out)
	if err != nil {
		panic(err)
	}
	return out
}
