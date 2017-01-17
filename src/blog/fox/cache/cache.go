package cache

import (
	"github.com/astaxie/beego/cache"
	"fmt"
	"time"
)

var Cache cache.Cache

func init() {
	var err error
	Cache, err = cache.NewCache("file", `{"CachePath":"./uploads/cache","FileSuffix":".cache","DirectoryLevel":2,"EmbedExpiry":120}`)
	if err != nil {
		fmt.Println("cache err:", err)
	}
}
// get cached value by key.
func Get(key string) interface{} {
	return Cache.Get(key)
}
// GetMulti is a batch version of Get.
func GetMulti(keys []string) []interface{} {
	return Cache.GetMulti(keys)
}
// set cached value with key and expire time.
func Put(key string, val interface{}, timeout time.Duration) error {
	return Cache.Put(key, val, timeout)
}
// delete cached value by key.
func Delete(key string) error {
	return Cache.Delete(key)
}
// increase cached int value by key, as a counter.
func Incr(key string) error {
	return Cache.Incr(key)
}
// decrease cached int value by key, as a counter.
func Decr(key string) error {
	return Cache.Decr(key)
}
// check if cached value exists or not.
func IsExist(key string) bool {
	return Cache.IsExist(key)
}
// clear all cache.
func ClearAll() error {
	return Cache.ClearAll()
}