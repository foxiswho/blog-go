package admin

import (
	"blog/fox/file"
	"fmt"
	"encoding/json"
	"blog/fox/editor"
	"blog/fox/array"
)

type Upload struct {
	Base
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
	var maps map[string]interface{}
	var err error
	t := c.GetString("t")
	token := c.GetString("token")
	if len(token) > 0 {
		maps, err = file.TokenDeCode(token)
		if err != nil {
			fmt.Println("令牌：" + token)
			fmt.Println("令牌解密失败：" + err.Error())
			token = ""
		} else {
			fmt.Println("令牌解密", maps)
		}

	}
	if token == "" {
		if t == "markdown" {
			md := &editor.EditorMd{}
			md.Message = "令牌错误"
			md.Success = 0
			c.Data["json"] = md
			c.ServeJSON()
		} else {
			c.Error("令牌错误")
		}
		return
	}
	file_name := "file"
	if t == "markdown" {
		file_name = "editormd-image-file"
	}
	//上传
	f, err := file.Upload(file_name, c.Ctx.Request, maps)
	if t == "markdown" {
		md := &editor.EditorMd{}
		if err != nil {
			md.Message = err.Error()
			md.Success = 0
		} else {
			md.Message = "上传成功"
			md.Url = f.Http
			md.Success = 1
		}
		c.Data["json"] = md
		c.ServeJSON()
	} else {
		if err != nil {
			c.Error(err.Error())
		} else {
			m, err := array.ObjToMap(f)
			if err != nil {
				c.Success("操作成功")
			} else {
				c.Success("操作成功", m)
			}
		}

	}
}