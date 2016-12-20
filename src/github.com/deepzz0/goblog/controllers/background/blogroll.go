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

type BlogrollController struct {
	Common
}

func (this *BlogrollController) Get() {
	this.TplName = "manage/blogroll/blogrollTemplate.html"
	this.Data["Title"] = "友情链接 | " + models.Blogger.BlogName
	this.LeftBar("blogroll")
	this.Content()
}

func (this *BlogrollController) Content() {
	this.Data["Blogrolls"] = models.Blogger.Blogrolls
}

func (this *BlogrollController) Post() {
	resp := helper.NewResponse()
	defer resp.WriteJson(this.Ctx.ResponseWriter)
	flag := this.GetString("flag")
	log.Debugf("flag=%s", flag)
	switch flag {
	case "save":
		this.saveBlogroll(resp)
	case "modify":
		this.getBlogroll(resp)
	case "delete":
		this.doDelete(resp)
	default:
		resp.Status = RS.RS_failed
		resp.Err = helper.Error{Level: helper.WARNING, Msg: "参数错误|未知的flag标志。"}
	}
}

func (this *BlogrollController) saveBlogroll(resp *helper.Response) {
	content := this.GetString("json")
	var br models.Blogroll
	err := json.Unmarshal([]byte(content), &br)
	if err != nil {
		log.Error(err)
		resp.Status = RS.RS_failed
		resp.Err = helper.Error{Level: helper.WARNING, Msg: "内容错误|要仔细检查哦。"}
		return
	}
	if br.ID == "TEST" {
		resp.Status = RS.RS_failed
		resp.Err = helper.Error{Level: helper.WARNING, Msg: "内容错误|请修改你需要添加的工具。"}
		return
	}
	if blogroll := models.Blogger.GetBlogrollByID(br.ID); blogroll != nil {
		*blogroll = br
		sort.Sort(models.Blogger.Blogrolls)
	} else {
		br.CreateTime = time.Now()
		models.Blogger.AddBlogroll(&br)
	}
}
func (this *BlogrollController) getBlogroll(resp *helper.Response) {
	id := this.GetString("id")
	if id != "" {
		if br := models.Blogger.GetBlogrollByID(id); br != nil {
			b, _ := json.Marshal(br)
			resp.Data = string(b)
		}
	} else {
		resp.Status = RS.RS_failed
		resp.Err = helper.Error{Level: helper.WARNING, Msg: "错误|参数错误。"}
	}
}

func (this *BlogrollController) doDelete(resp *helper.Response) {
	id := this.GetString("id")
	if id == "" {
		resp.Status = RS.RS_failed
		resp.Err = helper.Error{Level: helper.WARNING, Msg: "哦噢。。。|参数错误。"}
		return
	}
	if code := models.Blogger.DelBlogrollByID(id); code != RS.RS_success {
		resp.Status = code
		resp.Err = helper.Error{Level: helper.WARNING, Msg: "抱歉|系统没有找到该友情链接。"}
		return
	}
}
