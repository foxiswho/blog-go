package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	/**
	今天要和女神 去郊游了，要把带的东西准备好，吃的喝的，玩的，等等

	这里是数组
	 */
	println("数组 开始")
	println("键---值")
	list := []string{"a", "b", "c", "d", "e"};

	// range 是个迭代器,当被调用的时候,从它循环的内容中返回一个键值对
	for k, v := range list {
		fmt.Printf("%d---%s \n", k, v)
	}
	println("方式2 ")
	// range函数用来遍历字符串时，返回Unicode代码点。
	// 第一个返回值是每个字符的起始字节的索引，第二个是字符代码点，
	// 因为Go的字符串是由字节组成的，多个字节组成一个rune类型字符。
	for i, c := range "golang" {
		fmt.Println(i, c)
	}
	// 数值型数组
	var ints [10]int
	ints[0] = 221
	ints[1] = 333

	for k, v := range ints {

		fmt.Printf("%d ----- %d\n", k, v)
	}

	array := []int{11, 22, 33}

	array1 := [3]int{11, 22, 33}

	array2 := [...]int{11, 22, 33}

	println("array 数组个数",len(array))
	println("array 数组个数",len(array1))
	println("array 数组个数",len(array2))
	//
	array3 := [3][2]int{{1, 2}, {2, 3}, {3, 4}}

	for k, v := range array3 {
		for kk, vv := range v {
			fmt.Printf("%d ---- %d ---- %d\n", k, kk, vv)
		}
	}
	//////////////////////////////////////////
	/**
	 * slice append copy map
	 */

	println("------------------------slice-----------------------------")
	array10 := []string{"aa", "bb", "cc", "dd", "ee"}

	array11 := array10[1 : 3]

	array12 := array10[:]

	array13 := array10[:3]

	array14 := array10[:len(array10)]

	fmt.Println(array11)
	fmt.Println(array12)
	fmt.Println(array13)
	fmt.Println(array14)

	var array15 [2]int
	array15[0] = 22
	array15[1] = 33
	fmt.Println(array15)

	println("------------------------append-----------------------------")
	ads := []int{0, 0}
	ads1 := append(ads, 2)
	ads2 := append(ads1, 1, 2, 3, 4)
	ads3 := append(ads2, ads2...)
	fmt.Println(ads1)
	fmt.Println(ads2)
	fmt.Println(ads3)

	println("------------------------copy-----------------------------")
	var a = []int{1, 2, 3, 4, 5, 6, 7}
	var b = make([]int, 6)
	s := copy(b, a[:3])
	fmt.Println(s, b)
	s1 := copy(b, b[2:])
	fmt.Println(s1, b)

	println("------------------------map-----------------------------")
	mymap := map[string]int {
		"Jan": 31, "Feb": 28, "Mar": 31,
		"Apr": 30, "May": 31, "Jun": 30,
		"Jul": 31, "Aug": 31, "Sep": 30,
		"Oct": 31, "Nov": 30, "Dec": 31,
	}
	year := 0
	for _, days := range mymap {
		year += days
		fmt.Println(days)
	}
	fmt.Println("一年共有:", year, "天")

	// 向map中添加元素
	mymap["kkkk"] = 33
	println(mymap["kkkk"])

	// 删除map中某个元素
	delete(mymap, "kkkk")

	// 判断kkkk是否存在
	_, ok := mymap["kkkk"]

	if ok {
		println("有值")
	} else {
		println("无值")
	}
	//////////////////////////////////////////
	///////
	//给女神 买了一台电脑，女神想知道他是什么系统的
	//
	//case 体会自动终止，除非用 fallthrough 语句作为结尾
	//fallthrough表示继续执行下面的Case而不是退出Switch
	//当前系统
	fmt.Print("Go 运行的系统是： ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("苹果的 OS X 系统")
	case "linux":
		fmt.Println("开源的 Linux 系统")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
	// fallthrough表示继续执行下面的Case而不是退出Switch
	os := "无";
	switch os {
	case "无":
		fmt.Println("这就结束了")
		fmt.Println("他会继续执行下一个case")
		fallthrough  //必须是最后一个语句
	case "darwin":
		fmt.Println("苹果的 OS X 系统")
	case "linux":
		fmt.Println("开源的 Linux 系统")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
	//
	//周六约女神去游乐场，还有几天到周六
	//
	//switch 的条件从上到下的执行，当匹配成功的时候停止
	fmt.Println("还有几天到周六？")
	today := time.Now().Weekday()
	fmt.Println("今天是周几：", today)
	fmt.Println("周六表示：", time.Saturday)
	switch time.Saturday {
	case today + 0:
		fmt.Println("今天")
	case today + 1:
		fmt.Println("明天")
	case today + 2:
		fmt.Println("还有2天")
	default:
		fmt.Println("太多了，无法计算")
	}
	//
	//每天早上和女神说个早上好，中午好，晚上好
	//
	//没有条件的 switch 同 switch true 一样。
	//这一构造使得可以用更清晰的形式来编写长的 if-then-else 链
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("早上好!")
	case t.Hour() < 17:
		fmt.Println("中午好")
	default:
		fmt.Println("晚上好.")
	}
}
