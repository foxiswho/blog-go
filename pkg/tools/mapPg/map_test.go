package mapPg

import (
	"fmt"
	"sync"
	"testing"
)

func TestMap(t *testing.T) {
	// 创建一个键为 string，值为 int 的并发安全 Map
	numbers := SafeMap[string, int]{}

	// 并发写入
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			numbers.Store(key, i*10)
			fmt.Printf("写入 %s: %d\n", key, i*10)
		}(i)
	}
	wg.Wait()

	// 并发读取
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(readerID int) {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				key := fmt.Sprintf("key%d", j)
				if val, ok := numbers.Load(key); ok {
					fmt.Printf("读取器 %d 读取 %s: %d\n", readerID, key, val)
				}
			}
		}(i)
	}
	wg.Wait()

	// 遍历 Map
	fmt.Println("\nMap 内容:")
	numbers.Range(func(key string, value int) bool {
		fmt.Printf("键: %s, 值: %d\n", key, value)
		return true
	})
}
