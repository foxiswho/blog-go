package admin

import (
	"fox/util/Response"
	"fox/util/file"
	"fmt"
	"encoding/json"
)

type Upload struct {
	BaseController
}

func (c *Upload) URLMapping() {
	c.Mapping("Image", c.Image)
	c.Mapping("File", c.File)
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
func (c *BlogController)Image() {
	rsp := Response.NewResponse()
	defer rsp.WriteJson(c.Ctx.ResponseWriter)
	f, err := file.Upload("file", c.Ctx.Request)
	if err != nil {
		rsp.Error(err.Error())
		c.StopRun()
	}
	rsp.SetData(f)
	rsp.Success("")
}