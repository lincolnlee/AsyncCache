package AsyncCache

import (
	"fmt"
	//"log"
)

type exception struct{}

func (*exception) CatchException() {
	if ex := recover(); ex != nil {
		//log.Panic(ex)
		fmt.Println("Exception:", ex)
	}
}
