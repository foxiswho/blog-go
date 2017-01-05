package main

import (
	_ "fox/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fox/controllers"
	"strconv"
)


func main() {
	orm.Debug = true
	beego.ErrorController(&controllers.Error{})
	beego.SetStaticPath("/uploads","uploads")
	beego.AddFuncMap("array",array)
	beego.Run()
}

func array(m map[string]interface{},s int)[]string{
	if m !=nil{
		a:=m["tag_"+strconv.Itoa(s)]
		if a!=nil{
			b:=a.([]string)
			if len(b)>0{
				return b
			}
		}
	}

	return []string{}
}