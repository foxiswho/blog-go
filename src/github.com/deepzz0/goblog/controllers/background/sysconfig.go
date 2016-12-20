package background

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/deepzz0/go-com/log"
	"github.com/deepzz0/goblog/RS"
	"github.com/deepzz0/goblog/helper"
	"github.com/deepzz0/goblog/models"
)

type SysconfigController struct {
	Common
}

func (this *SysconfigController) Get() {
	this.TplName = "manage/system/systemConfig.html"
	this.Data["Title"] = "系统设置 | " + models.Blogger.BlogName
	this.LeftBar("sysconfig")
	this.Content()
}

func (this *SysconfigController) Content() {
	this.Data["SiteVerify"] = models.ManageConf.SiteVerify
}

func (this *SysconfigController) Post() {
	resp := helper.NewResponse()
	defer resp.WriteJson(this.Ctx.ResponseWriter)
	flag := this.GetString("flag")
	log.Debugf("flag = %s", flag)
	switch flag {
	case "deleteverify":
		this.deleteVerify(resp)
	case "addverify":
		this.addVerify(resp)
	case "updatesitemap":
		this.updateSitemap(resp)
	case "getsitemap":
		this.getSitemap(resp)
	default:
		resp.Status = RS.RS_failed
		resp.Err = helper.Error{Level: helper.WARNING, Msg: "参数错误|未知的flag标志。"}
	}
}

func (this *SysconfigController) deleteVerify(resp *helper.Response) {
	name := this.GetString("name")
	if name == "" {
		resp.Status = RS.RS_params_error
		resp.Tips(helper.WARNING, RS.RS_params_error)
		return
	}
	models.ManageConf.DelVerification(name)
}

func (this *SysconfigController) addVerify(resp *helper.Response) {
	name := this.GetString("name")
	content := this.GetString("content")
	if name == "" || content == "" {
		resp.Status = RS.RS_params_error
		resp.Tips(helper.WARNING, RS.RS_params_error)
		return
	}
	verify := models.ManageConf.GetVerification(name)
	if verify != nil {
		resp.Status = RS.RS_duplicate_add
		resp.Tips(helper.WARNING, RS.RS_duplicate_add)
		return
	}
	verify = models.NewVerify()
	verify.Name = name
	verify.Content = content
	models.ManageConf.AddVerification(verify)
}

func (this *SysconfigController) updateSitemap(resp *helper.Response) {
	content := this.GetString("content")
	if content == "" {
		resp.Status = RS.RS_params_error
		resp.Tips(helper.WARNING, RS.RS_params_error)
		return
	}
	_, err := os.Stat(models.SiteFile)
	if err != nil && !strings.Contains(err.Error(), "no such file") {
		log.Error(err)
		return
	} else {
		os.Remove(models.SiteFile)
	}
	f, err := os.Create(models.SiteFile)
	if err != nil {
		log.Error(err)
		return
	}
	defer f.Close()
	_, err = f.WriteString(content)
	if err != nil {
		resp.Status = RS.RS_failed
		resp.Err = helper.Error{Level: helper.WARNING, Msg: "错误|" + err.Error()}
	}
}

func (this *SysconfigController) getSitemap(resp *helper.Response) {
	f, err := os.Open(models.SiteFile)
	if err != nil {
		resp.Status = RS.RS_failed
		resp.Err = helper.Error{Level: helper.WARNING, Msg: "错误|" + err.Error()}
		return
	}
	defer f.Close()
	data, _ := ioutil.ReadAll(f)
	resp.Data = string(data)
}

type DataBackupRecover struct {
	Common
}
