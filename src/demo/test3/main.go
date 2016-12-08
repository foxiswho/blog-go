package main

import (
	"fmt"
	"strings"
)

func main() {

	/**
	今天约女神出来，但是女神没空，女神要看书
	条件
	 */
	/**
	 * 控制语句，判断数字大于
	 */
	if x := 19; x > 19 {
		fmt.Println("晚上有空:", x)
	} else {
		fmt.Println("晚上没空")
	}
	//女神给我回复
	str := "不约"
	if strings.EqualFold(str, "约") {
		fmt.Println("女神今晚有空，走起")
	} else {
		fmt.Println("女神今晚没空，女神要看书，自己玩吧")
	}
}
