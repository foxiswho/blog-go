package background

import (
	"encoding/json"
	"sort"
	"time"

	"github.com/deepzz0/go-com/log"
	"github.com/deepzz0/goblog/RS"
	"github.com/deepzz0/goblog/helper"
	"github.com/deepzz0/goblog/models"
	// db "github.com/deepzz0/go-com/mongo"
)

type CategoryController struct {
	Common
}

func (this *CategoryController) Get() {
	this.TplName = "manage/category/categoryTemplate.html"
	this.Data["Title"] = "分类管理 | " + models.Blogger.BlogName
	this.LeftBar("category")
	this.Content()
}
func (this *CategoryController) Content() {
	this.Data["Categories"] = models.Blogger.GetValidCategory()
	this.Data["Tags"] = models.Blogger.Tags
}

func (this *CategoryController) Post() {
	resp := helper.NewResponse()
	defer resp.WriteJson(this.Ctx.ResponseWriter)
	flag := this.GetString("flag")
	log.Debugf("flag = %s", flag)
	switch flag {
	case "save":
		this.saveCategory(resp)
	case "modify":
		this.getCategory(resp)
	case "deletecat":
		this.doDeleteCat(resp)
	case "deletetag":
		this.doDeleteTag(resp)
	default:
		resp.Status = RS.RS_failed
		resp.Err = helper.Error{Level: helper.WARNING, Msg: "参数错误|未知的flag标志。"}
	}
}
func (this *CategoryController) getCategory(resp *helper.Response) {
	id := this.GetString("id")
	if id != "" {
		if cat := models.Blogger.GetCategoryByID(id); cat != nil {
			b, _ := json.Marshal(cat)
			resp.Data = string(b)
		}
	} else {
		resp.Status = RS.RS_failed
		resp.Err = helper.Error{Level: helper.WARNING, Msg: "错误|参数错误。"}
	}
}

func (this *CategoryController) saveCategory(resp *helper.Response) {
	content := this.GetString("json")
	var cat models.Category
	err := json.Unmarshal([]byte(content), &cat)
	if err != nil {
		log.Error(err)
		resp.Status = RS.RS_failed
		resp.Err = helper.Error{Level: helper.WARNING, Msg: "内容错误|要仔细检查哦。"}
		return
	}
	if cat.ID == "TEST" {
		resp.Status = RS.RS_failed
		resp.Err = helper.Error{Level: helper.WARNING, Msg: "内容错误|请修改你需要添加的分类。"}
		return
	}
	if category := models.Blogger.GetCategoryByID(cat.ID); category != nil {
		*category = cat
		sort.Sort(models.Blogger.Categories)
	} else {
		cat.CreateTime = time.Now()
		models.Blogger.AddCategory(&cat)
	}
}
func (this *CategoryController) doDeleteCat(resp *helper.Response) {
	id := this.GetString("id")
	if id == "" || id == "default" {
		resp.Status = RS.RS_failed
		resp.Err = helper.Error{Level: helper.WARNING, Msg: "哦噢。。。|参数错误,default不能删除。"}
		return
	}
	if code := models.Blogger.DelCatgoryByID(id); code != RS.RS_success {
		resp.Status = code
		resp.Err = helper.Error{Level: helper.WARNING, Msg: "抱歉|系统没有找到该分类。"}
		return
	}
}
func (this *CategoryController) doDeleteTag(resp *helper.Response) {
	id := this.GetString("id")
	if id == "" {
		resp.Status = RS.RS_failed
		resp.Err = helper.Error{Level: helper.WARNING, Msg: "哦噢。。。|参数错误。"}
		return
	}
	if code := models.Blogger.DelTagByID(id); code != RS.RS_success {
		resp.Status = code
		resp.Err = helper.Error{Level: helper.WARNING, Msg: "抱歉|系统没有找到该标签。"}
		return
	}
}
