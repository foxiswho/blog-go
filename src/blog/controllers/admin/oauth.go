package admin

import (
	"blog/app/csdn"
	"fmt"
	"blog/fox/config"
	"blog/service/oauth"
	"blog/service/admin"
	"blog/service/conf"
)
//第三方账号登录
type Oauth struct {
	BaseNoLogin
}

func (c *Oauth) URLMapping() {
	c.Mapping("Get", c.Get)
	c.Mapping("Csdn", c.Csdn)
}
// @router /oauth [get]
func (c *Oauth)Get() {
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
		web.SetRedirectUri(config.String("http") + "/admin/oauth_csdn")
		url := web.GetAuthorizeUrl()
		c.Redirect(url, 302)
	} else {
		c.Error("类别不存在")
	}
}
// @router /oauth_csdn [get]
func (c *Oauth)Csdn() {
	token := c.GetString("code")
	if len(token) < 1 {
		c.Error("token 不存在")
	}
	web := csdn.NewAuthorizeWeb()
	ok, err := web.SetConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("status:", ok);
	web.SetRedirectUri(config.String("http") + "/admin/oauth_csdn")
	acc, err1 := web.GetAccessToken(token)
	if err1 != nil {
		c.Error(err.Error())
	} else {
		//查询用户是否存在
		oau := oauth.NewConnect()
		con, err := oau.Admin(conf.APP_CSDN, acc.Username, false)
		if err == nil {
			fmt.Println("con 值", con)
			adminUser := admin.NewAdminUserService()
			adm, err := adminUser.GetAdminById(con.Uid)
			if err == nil {
				//转换为session
				AdminSession := admin.NewAdminSessionService()
				Session := AdminSession.Convert(adm)
				c.SessionSet(Session)
				fmt.Println("验证通过 跳转到后台")
				url := config.String("http") + "/admin/index"
				c.Redirect(url, 302)
			} else {
				fmt.Println(err)
				c.Error(err.Error())
			}

		} else if err.Error()=="未绑定" {
			c.TplName="app/csdn/get.html"
		}else {
			fmt.Println(err)
			c.Error(err.Error())
		}

	}
}