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
		fmt.Println("这个是大于判断，晚上有空:", x)
	} else {
		fmt.Println("这个是大于判断，晚上没空")
	}
	if x := 19; x >= 19 {
		fmt.Println("这个是大于等于判断，晚上有空:", x)
	} else {
		fmt.Println("这个是大于等于判断，晚上没空")
	}
	//女神给我回复
	str := "不约"
	if strings.EqualFold(str, "约") {
		fmt.Println("这个是字符串相等判断，女神今晚有空，走起")
	} else {
		fmt.Println("这个是字符串相等判断，女神今晚没空，女神要看书，自己玩吧")
	}
	//这里为什么不用 双等于呢，双等于 判断 两个对象是否同一个引用对象
	///////
	//////
	/////
	/////
	/////
	//这里是否两个真值
	if true && true {
		fmt.Println("true")
	}
	//不为假的时候，那么就是真值
	if !false {
		fmt.Println("true")
	}
	////
	test(2)
}

func test(i int) bool {
	if i == 1 {
		return true
	} else {
		fmt.Println("返回：false")
		return false
	}
	fmt.Println("继续执行")
	//panic("oo")
	fmt.Println("看看还有么")
	panic("oo")
	//if语句块中去做return处理,而else中不处理，而是继续执行if-else后面的代码，这样能减少一个代码缩进，
	// 不需要在了解代码时去记住else语句块的处理。
	// 当然如果想必须这样写,也可以进行特殊处理，在函数的末行添加语句**panic("")**
}
//http://www.cnblogs.com/howDo/archive/2013/06/01/GoLang-Control.html
