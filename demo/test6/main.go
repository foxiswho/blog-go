package main

import "fmt"

func main() {
	/**
	女神生气了，要哄哄，要让女神开心

	 */
	if angry("因为你生气") {
		fmt.Println("女神生气了，快哄哄")
	}else {
		fmt.Println("女神没有生气，但是不能掉以轻心哦")
	}
}

func angry(c string) bool {
	switch c {
	case "因为他人他事而生气", "因为你生气", "你不在乎她",
		"她吃醋了", "你嫌弃她", "你的观点让她难以接受",
		"你的做法让她难以接受", "你在她难受/不高兴/遇到困难的时候不理她":
		return true
	}

	return false
}