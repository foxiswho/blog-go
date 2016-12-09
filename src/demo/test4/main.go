package main

import "fmt"

func main() {
	/**
	要和女神约会了，比如先吃吃饭啊
	、、
	这里是 函数使用和 goto使用
	 */
	println("for 循环")
	func_for()
	println("流程控制 break 退出循环体")
	func_for2()
	println("流程控制 goto 跳转")
	func_goto()
}
// for 循环
func func_for() {
	sum := 0

	for i := 0; i < 100; i++ {
		sum += i
		if i == 5 {
			println("数字5 跳过")
			continue
		}
	}
	println(sum)
}
// break 退出循环体
func func_for2() {
	J:
	for i := 0; i < 10; i++ {
		for j := 0; j < 5; j++ {
			if j > 5 {
				println("流程控制 break 退出循环体",j)
				break J
			}


		}
	}
	fmt.Println("break 结束")
}
//goto 跳转
func func_goto() {
	i := 0
	Here:
	if i < 100 {
		fmt.Printf("%d,\n", i)
		i++
	} else {
		println("流程控制 goto 结束",i)
		return
	}
	println("流程控制 goto 跳转",i)
	goto Here
}


//http://www.cnblogs.com/howDo/archive/2013/06/01/GoLang-Control.html