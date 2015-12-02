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

##LICENSE

AsyncCache is licensed under the [Apache Licence, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0.html).