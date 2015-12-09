package AsyncCache

import (
	"fmt"
	//"log"
)

type exception struct {
	exinfo interface{}
}

func (this *exception) catchException() {
	if ex := recover(); ex != nil {
		//log.Panic(ex)
		//fmt.Println("Exception:", ex)
		this.exinfo = ex
	}
}

func (this *exception) Try(f func()) {
	defer this.catchException()
	f()
}

func (this *exception) Catch(f func(interface{})) {
	if this.exinfo != nil {
		f(this.exinfo)
	}
}
