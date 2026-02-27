package mapPg

import "sync"

// SafeMap
// @Description: 泛型并发安全 Map
type SafeMap[K comparable, V any] struct {
	m sync.Map
}

// Store
//
//	@Description: 设置键值对
//	@receiver sm
//	@param key
//	@param value
func (sm *SafeMap[K, V]) Store(key K, value V) {
	sm.m.Store(key, value)
}

// Load
//
//	@Description: 获取键对应的值
//	@receiver sm
//	@param key
//	@return V
//	@return bool
func (sm *SafeMap[K, V]) Load(key K) (V, bool) {
	val, ok := sm.m.Load(key)
	var zero V
	if !ok {
		return zero, false
	}
	return val.(V), true
}

// Delete
//
//	@Description:删除键值对
//	@receiver sm
//	@param key
func (sm *SafeMap[K, V]) Delete(key K) {
	sm.m.Delete(key)
}

// Range
//
//	@Description: 遍历 Map
//	@receiver sm
//	@param f
func (sm *SafeMap[K, V]) Range(f func(key K, value V) bool) {
	sm.m.Range(func(k, v interface{}) bool {
		return f(k.(K), v.(V))
	})
}
