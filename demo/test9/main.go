package main

import "fmt"

func main() {
	A()
	B()
	C()
}

func A() {
	fmt.Println("func a")
}

func B() {
	defer func() {
		if error := recover(); error != nil {
			fmt.Println("func b..")
		}
	}()
	panic("func b")
}

func C() {
	fmt.Println("func c")
}
