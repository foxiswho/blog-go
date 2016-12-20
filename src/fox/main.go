package main

import (
	_ "fox/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)


func main() {
	orm.Debug = true
	beego.Run()
}

