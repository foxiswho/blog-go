package admin

import (
	"fox/util/Response"
	"fox/util/file"
	"fmt"
	"encoding/json"
	"fox/util/editor"
)

type Upload struct {
	BaseController
}

func (c *Upload) URLMapping() {
	c.Mapping("Image", c.Image)
	c.Mapping("File", c.File)
	c.Mapping("Post", c.Post)
}
//图片
// @router /upload/image [get]
func (c *Upload)Image() {
	type Option struct {
		TypeId int
		Id     int
	}
	stb := &Option{}
	opt := c.GetString("opt")
	err := json.Unmarshal([]byte(opt), &stb)
	if err != nil {
		fmt.Println(err)
	}
	c.Data["token"] = opt
	c.Data["title"] = "图片上传"
	c.TplName = "admin/upload/image.html"
}
// @router /upload/file [get]
func (c *Upload)File() {
	c.Data["title"] = "文件上传"
	c.TplName = "admin/upload/file.html"
}
//上传图片
// @router /upload/image [post]
func (c *Upload)Post() {
	rsp := Response.NewResponse()
	defer rsp.WriteJson(c.Ctx.ResponseWriter)
	t := c.GetString("t")
	token := c.GetString("token")
	if token == "" {
		if t == "markdown" {
			md := &editor.EditorMd{}
			md.Message = "令牌错误"
			md.Success = 0
			c.Data["json"] = md
			c.ServeJSON()
		} else {
			rsp.Error("令牌错误")
		}
		return
	}

	file_name := "file"
	if t == "markdown" {
		file_name = "editormd-image-file"
	}

	f, err := file.Upload(file_name, c.Ctx.Request, "")
	if t == "markdown" {
		md := &editor.EditorMd{}
		if err != nil {
			md.Message = err.Error()
			md.Success = 0
		} else {
			md.Message = "上传成功"
			md.Url = f.Url
			md.Success = 1
		}
		c.Data["json"] = md
		c.ServeJSON()
	} else {
		if err != nil {
			rsp.Error(err.Error())
			c.StopRun()
		}
		rsp.SetData(f)
		rsp.Success("")
	}
}