package background

import (
	"encoding/json"
	"sort"
	"time"

	"github.com/deepzz0/go-com/log"
	"github.com/deepzz0/goblog/RS"
	"github.com/deepzz0/goblog/helper"
	"github.com/deepzz0/goblog/models"
)

type SocialController struct {
	Common
}

func (this *SocialController) Get() {
	this.TplName = "manage/social/socialTemplate.html"
	this.Data["Title"] = "社交工具 | " + models.Blogger.BlogName
	this.LeftBar("social")
	this.Content()
}
func (this *SocialController) Content() {
	this.Data["Socials"] = models.Blogger.Socials
}

func (this *SocialController) Post() {
	resp := helper.NewResponse()
	defer resp.WriteJson(this.Ctx.ResponseWriter)
	flag := this.GetString("flag")
	log.Debugf("flag = %s", flag)
	switch flag {
	case "save":
		this.saveSocial(resp)
	case "modify":
		this.getSocial(resp)
	case "delete":
		this.doDelete(resp)
	default:
		resp.Status = RS.RS_failed
		resp.Err = helper.Error{Level: helper.WARNING, Msg: "参数错误|未知的flag标志。"}
	}
}
func (this *SocialController) saveSocial(resp *helper.Response) {
	content := this.GetString("json")
	var sc models.Social
	err := json.Unmarshal([]byte(content), &sc)
	if err != nil {
		log.Error(err)
		resp.Status = RS.RS_failed
		resp.Err = helper.Error{Level: helper.WARNING, Msg: "内容错误|要仔细检查哦。"}
		return
	}
	if sc.ID == "TEST" {
		resp.Status = RS.RS_failed
		resp.Err = helper.Error{Level: helper.WARNING, Msg: "内容错误|请修改你需要添加的社交。"}
		return
	}
	if social := models.Blogger.GetSocialByID(sc.ID); social != nil {
		*social = sc
		sort.Sort(models.Blogger.Socials)
	} else {
		sc.CreateTime = time.Now()
		models.Blogger.AddSocial(&sc)
	}
}

func (this *SocialController) getSocial(resp *helper.Response) {
	id := this.GetString("id")
	if id != "" {
		if social := models.Blogger.GetSocialByID(id); social != nil {
			b, _ := json.Marshal(social)
			resp.Data = string(b)
		}
	} else {
		resp.Status = RS.RS_failed
		resp.Err = helper.Error{Level: helper.WARNING, Msg: "错误|参数错误。"}
	}
}

func (this *SocialController) doDelete(resp *helper.Response) {
	id := this.GetString("id")
	if id == "" {
		resp.Status = RS.RS_failed
		resp.Err = helper.Error{Level: helper.WARNING, Msg: "哦噢。。。|参数错误。"}
		return
	}
	if code := models.Blogger.DelSocialByID(id); code != RS.RS_success {
		resp.Status = code
		resp.Err = helper.Error{Level: helper.WARNING, Msg: "抱歉|系统没有找到该工具。"}
		return
	}
}
