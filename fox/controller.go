package fox

import (
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"github.com/foxiswho/blog-go/fox/config"
)

//此处 为以后 更换框架做准备
//控制器基类
type Controller struct {
	web.Controller //继承 beego 控制器
}

//  设置模版
func (c *Controller) SetTpl(str string) {
	fmt.Println("TplName=", str)
	//是否有主题
	theme := config.String("theme")
	if len(theme) > 0 {
		str = theme + "/" + str
	}
	fmt.Println("模版地址", str)
	c.TplName = str
}
