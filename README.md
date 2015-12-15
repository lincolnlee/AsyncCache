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
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lincolnlee/AsyncCache"
)

func main() {
	DoSomething()

	u := getUser(1)

	DoOtherThings()
}

func getUser(uid int) UserInfo {
	key := "UserInfo_" + string(uid)

	v := AsyncCache.InstanceContainer.AsyncCacheHandler.AsyncGetAndUpdateData(
		func() interface{} {
			return getUserDataFromDB(uid)
		},
		key)

	return UserInfo(v)
}

func getUserDataFromDB(uid int) UserInfo {
	db, _ := sql.Open("mysql", dbConnString)
	defer db.Close()

	rows, err := db.Query("SELECT id,username,age FROM users where id = ?", uid)
	checkError(err)

	var user UserInfo = nil

	for rows.Next() {
		var id int32
		var username string
		var age int32
		_ = rows.Scan(&id, &username, &age)

		user = UserInfo{id, username, age}
	}

	return user
}

```

##LICENSE

AsyncCache is licensed under the [Apache Licence, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0.html).