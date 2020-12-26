package main

import (
	"regexp"
	"fmt"
)

func main() {
	idStr:="1232323-"
	reg := regexp.MustCompile("^\\d+$")
	a:=reg.MatchString(idStr)
	fmt.Println(a)

	ok, err := regexp.Match(`^\\d+$`, []byte(idStr))
	fmt.Println(ok)
	fmt.Println(err)
}
