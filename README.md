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
	AsyncCache.InstanceContainer.AsyncCacheHandler.AsyncGetAndUpdateData(f, key)
}

```