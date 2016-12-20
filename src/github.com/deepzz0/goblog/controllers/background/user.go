package background

import (
	"bytes"
	"html/template"

	// "github.com/astaxie/beego"
	"github.com/deepzz0/goblog/RS"
	"github.com/deepzz0/goblog/helper"
	"github.com/deepzz0/goblog/models"
)

type UserController struct {
	Common
}

func (this *UserController) Post() {
	resp := helper.NewResponse()
	defer resp.WriteJson(this.Ctx.ResponseWriter)
	flag := this.GetString("flag")
	switch flag {
	case "info":
		this.userInfo(resp)
	case "modifyinfo":
		this.doModifyInfo(resp)
	case "modpasswd":
		this.modifyPasswd(resp)
	case "domodpasswd":
		this.doModifyPasswd(resp)
	default:
		resp.Status = RS.RS_failed
		resp.Err = helper.Error{Level: helper.WARNING, Msg: "参数错误|未知的flag标志。"}
	}
}
func (this *UserController) userInfo(resp *helper.Response) {
	infoT, err := template.ParseFiles("views/manage/user/user.html")
	if err != nil {
		panic(err)
	}
	Map := make(map[string]string)
	Map["BlogName"] = models.Blogger.BlogName
	Map["Icon"] = models.Blogger.HeadIcon
	Map["Introduce"] = models.Blogger.Introduce
	Map["Sex"] = models.Blogger.Sex
	Map["Username"] = models.Blogger.UserName
	Map["Email"] = models.Blogger.Email
	Map["Address"] = models.Blogger.Address
	Map["Education"] = models.Blogger.Education
	Map["LoginTime"] = models.Blogger.LoginTime.Format(helper.Layout_y_m_d_time)
	Map["IP"] = models.Blogger.LoginIp

	var buffer bytes.Buffer
	infoT.Execute(&buffer, Map)
	resp.Data = buffer.String()
}
func (this *UserController) modifyPasswd(resp *helper.Response) {
	modifypasswdT, err := template.ParseFiles("views/manage/user/modifypasswd.html")
	if err != nil {
		panic(err)
	}
	var buffer bytes.Buffer
	modifypasswdT.Execute(&buffer, nil)
	resp.Data = buffer.String()
}
func (this *UserController) doModifyInfo(resp *helper.Response) {
	if blogname := this.GetString("blogname"); blogname != "" {
		models.Blogger.BlogName = blogname
	}
	if icon := this.GetString("icon"); icon != "" {
		models.Blogger.HeadIcon = icon
	}
	if introduce := this.GetString("introduce"); introduce != "" {
		models.Blogger.Introduce = introduce
	}
	if sex := this.GetString("sex"); sex != "" {
		models.Blogger.Sex = sex
	}
	if email := this.GetString("email"); email != "" {
		models.Blogger.Email = email
	}
	if address := this.GetString("address"); address != "" {
		models.Blogger.Address = address
	}
	if education := this.GetString("education"); education != "" {
		models.Blogger.Education = education
	}
	resp.Success()
}
func (this *UserController) doModifyPasswd(resp *helper.Response) {
	oldPasswd := this.GetString("old")
	newPasswd := this.GetString("new")
	if oldPasswd == "" || newPasswd == "" {
		resp.Status = RS.RS_failed
		resp.Err = helper.Error{Level: helper.WARNING, Msg: "错误|参数错误。"}
		return
	}
	if !helper.VerifyPasswd(models.Blogger.PassWord, models.Blogger.UserName, oldPasswd, models.Blogger.Salt) {
		resp.Status = RS.RS_failed
		resp.Err = helper.Error{Level: helper.WARNING, Msg: "错误|原密码错误。"}
		return
	}
	models.Blogger.ChangePassword(newPasswd)
	resp.Success()
}
