package main

import (
	_ "blog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"blog/controllers"
	"strconv"
)


func main() {
	orm.Debug = true
	beego.ErrorController(&controllers.Error{})
	beego.SetStaticPath("/uploads","uploads")
	//beego.AddFuncMap("array",array)
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