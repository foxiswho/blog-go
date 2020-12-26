package main

import (
	"fmt"
)

type i int

type sss struct {
	Name string
}

func main() {
	var test i
	test.kk(100)
	fmt.Println(test)

	test1 := sss{Name: "xxx"}
	test1.s1()
}

func (k *i) kk(num int) {
	*k += i(num)
}

func (s sss) s1() {
	fmt.Println("s1")
	fmt.Println(s.Name)
}
