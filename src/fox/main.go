package main

import (
	_ "fox/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fox/controllers"
)


func main() {
	orm.Debug = true
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}

