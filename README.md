AsyncCache
===========

An async Cache Middleware in golang

##Install

```shell
go get github.com/lincolnlee/AsyncCache

```

##Example

```go
package main

import (
	"github.com/lincolnlee/AsyncCache"
)

func main() {
	DoSomething()

	AsyncCache.InstanceContainer.Exception.Try(
		func() {
			AsyncCache.InstanceContainer.AsyncCacheHandler.AsyncGetAndUpdateData(
				func() interface{} {
					return "Hello, World!"
				},
				"testKey")
			panic("test error")
		})

	AsyncCache.InstanceContainer.Exception.Catch(
		func(ex interface{}) {
			fmt.Println("catch:", ex)
		})

	DoOtherThings()
}

```

##LICENSE

AsyncCache is licensed under the [Apache Licence, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0.html).