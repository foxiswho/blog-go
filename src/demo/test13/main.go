package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个channel用以同步goroutine
	done := make(chan bool)

	// 在goroutine中执行输出操作
	go func() {
		println("goroutine message")

		// 告诉main函数执行完毕.
		// 这个channel在goroutine中是可见的
		// 因为它是在相同的地址空间执行的.
		done <- true
	}()

	println("main function message")
	<-done // 等待goroutine结束
	/////////
	/////
	////
	////
	message := make(chan string) // 无缓冲
	count := 3

	go func() {
		for i := 1; i <= count; i++ {
			fmt.Println("send message")
			message <- fmt.Sprintf("message %d", i)
		}
	}()

	time.Sleep(time.Second * 3)

	for i := 1; i <= count; i++ {
		fmt.Println(<-message)
	}
	///////
	//////
	//////
	message2 := make(chan string)
	count2 := 3

	go func() {
		for i := 1; i <= count2; i++ {
			message2 <- fmt.Sprintf("message %d", i)
		}
		close(message2)
	}()

	for msg := range message2 {
		fmt.Println(msg)
	}

	//http://www.oschina.net/translate/golang-channels-tutorial
}
