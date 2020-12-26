package main

import (
	"fmt"
)
//接口
type MM interface {
	Name() string
	Test()
}
//结构体
type Test struct {
	name string
}
//结构体 Test 方法
func (t Test) Name() string {
	return t.name
}
//结构体 Test 方法
func (t Test) Test() {
	fmt.Println("Test")
}

func main() {
	t := Test{"name"}
	t.Test()
	dis(t)
}
//函数
func dis(m MM) {
	if t, ok := m.(Test); ok {
		fmt.Println("okk", t.name)
	}
}
