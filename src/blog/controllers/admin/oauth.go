package admin

import (
	"blog/app/csdn"
	"fmt"
	"blog/fox/config"
)

type Oauth struct {
	BaseNoLogin
}

func (c *Oauth)Get() {
	tp := c.GetString("type")
	if len(tp) < 1 {
		c.Error("类别错误")
	}
}
func (c *Oauth)Csdn() {
	tp := c.GetString("type")
	if len(tp) < 1 {
		c.Error("类别错误")
	}
	if tp == "csdn" {
		web := csdn.NewAuthorizeWeb()
		ok, err := web.SetConfig()
		if err != nil {
			fmt.Println(err)
			c.Error("csdn oauth:" + err.Error())
		}
		fmt.Println("status:", ok);
		web.SetRedirectUri(config.String("http") + "/admin/auth_token")
		url := web.GetAuthorizeUrl()
		c.Redirect(url, 302)
	} else {
		c.Error("类别不存在")
	}

}