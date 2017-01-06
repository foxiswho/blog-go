package cache

import (
	"github.com/astaxie/beego/cache"
	"fmt"
)

var Cache cache.Cache

func init() {
	var err error
	Cache, err = cache.NewCache("file", `{"CachePath":"./uploads/cache","FileSuffix":".cache","DirectoryLevel":2,"EmbedExpiry":120}`)
	if err!=nil{
		fmt.Println("cache err:",err)
	}
}