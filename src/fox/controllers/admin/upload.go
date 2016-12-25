package admin

import (

)
import (
	"fox/util/Response"
	"fox/util/file"
	"fmt"
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
	ok, err := file.Upload("file", c.Ctx.Request)
	fmt.Println(ok)
	fmt.Println(err)
	rsp.Success("")
}