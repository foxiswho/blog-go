package cache

import (
	"fmt"
	"github.com/beego/beego/v2/client/cache"
	"time"
)
//此处 为以后 更换框架做准备
var Cache cache.Cache

func init() {
	var err error
	Cache, err = cache.NewCache("file", `{"CachePath":"./uploads/cache","FileSuffix":".cache","DirectoryLevel":2,"EmbedExpiry":120}`)
	if err != nil {
		fmt.Println("cache err:", err)
	}
}
// get cached value by key.
func Get(key string) (interface{}) {
	get, _ := Cache.Get(nil, key)
	return get
}
// GetMulti is a batch version of Get.
func GetMulti(keys []string) []interface{} {
	multi, _ := Cache.GetMulti(nil, keys)
	return multi
}
// set cached value with key and expire time.
func Put(key string, val interface{}, timeout time.Duration) error {
	return Cache.Put(nil,key, val, timeout)
}
// delete cached value by key.
func Delete(key string) error {
	return Cache.Delete(nil,key)
}
// increase cached int value by key, as a counter.
func Incr(key string) error {
	return Cache.Incr(nil,key)
}
// decrease cached int value by key, as a counter.
func Decr(key string) error {
	return Cache.Decr(nil,key)
}
// check if cached value exists or not.
func IsExist(key string) bool {
	exist, _ := Cache.IsExist(nil, key)
	return exist
}
// clear all cache.
func ClearAll() error {
	return Cache.ClearAll(nil)
}
