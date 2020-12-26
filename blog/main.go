package main

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	"github.com/foxiswho/blog-go/blog/controllers"
	_ "github.com/foxiswho/blog-go/blog/routers"
	"strconv"
)


func main() {
	orm.Debug = true
	web.ErrorController(&controllers.Error{})
	web.SetStaticPath("/uploads","uploads")
	//beego.AddFuncMap("array",array)
	web.Run()
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
