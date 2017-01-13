package file

import (
	"os"
	"github.com/astaxie/beego"
	"blog/fox/datetime"
	"time"
	"net/http"
	"blog/fox/crypt"
	"path"
	"fmt"
	"blog/fox"
	"io"
	"strconv"
	"strings"
	"mime/multipart"
	"blog/model"
	"blog/fox/db"
	"blog/fox/number"
)
//上传成功后返回结构体
type UploadFile struct {
	NameOriginal string `json:"name_original" name:"保存的文件名"`
	Name         string `json:"name"  name:"原文件名"`
	Path         string `json:"path"  name:"文件路径"`
	Size         int `json:"size"  name:"文件大小"`
	Ext          string `json:"ext"  name:"文件后缀"`
	Md5          string `json:"md5"  name:"md5"`
	Http         string `json:"http"  name:"图片http地址"`
	AttachmentId int `json:"attachment_id"  name:"attachment_id"`
	Id           int `json:"attachment_id"  name:"id"`
	Url          string `json:"url"  name:"完整地址"`
	Config       map[string]string `json:"-"`
	Attachment   *model.Attachment `json:"-"`
}
// 获取文件信息的接口
type Stat interface {
	Stat() (os.FileInfo, error)
}
// 获取文件大小的接口
type Size interface {
	Size() int64
}

func NewUploadFile() *UploadFile {
	return new(UploadFile)
}
//上传
// @field 上传表单名称
// @r     Request
// @map   配置数组map
func Upload(field string, r *http.Request, maps map[string]interface{}) (*UploadFile, error) {
	file, header, err := r.FormFile(field)
	if err != nil {
		return nil, err
	}
	UploadFile := NewUploadFile()
	//fmt.Println("maps",maps)
	upload_type1 := maps["upload_type"]
	upload_type := ""
	if upload_type1 != nil {
		upload_type = upload_type1.(string)
	}
	//数据填充
	_, err = UploadFile.SetUploadFileData(upload_type, file, header)
	if err != nil {
		return nil, err
	}
	//审核 检测 大小，文件后缀
	if _, err := UploadFile.Check(); err != nil {
		return nil, err
	}
	defer file.Close()
	defer UploadFile.LocalTmpFileRemove()
	//临时文件
	if _, err := UploadFile.LocalSaveFile(file, UploadFile.GetLocalTmpPath(), UploadFile.Name); err != nil {
		return nil, err
	}
	//调用第三方存储
	if UploadFile.Config["type"] == "QiNiu" {

		UP := NewQiNiu()
		ret, err := UP.Upload(file, UploadFile)
		if err != nil {
			fmt.Println("七牛回执 ERR：", err)
		}
		fmt.Println("七牛回执", ret)
	}
	//本地是否保存
	if UploadFile.Config["local_save_is"] == "true" {
		//保存到本地
		if _, err := UploadFile.LocalSaveFile(file, UploadFile.Path, UploadFile.Name); err != nil {
			return nil, err
		}
	}
	//保存到数据库
	UploadFile.SaveDataBase(maps)
	//最后处理
	return UploadFile, nil
}
//判断目录或文件是已存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
//设置配置
func (c *UploadFile)SetConfig(mod string) (bool, error) {
	isFind := true
	_, err := beego.GetConfig(mod, "type", "")
	if err != nil {
		isFind = false
	}
	if !isFind {
		maps, err := beego.AppConfig.GetSection("upload_default")
		if err != nil {
			fmt.Println("config error:", err)
			return false, err
		}
		c.Config = maps
	} else {
		maps, err := beego.AppConfig.GetSection(mod)
		if err != nil {
			fmt.Println("config error:", err)
			return false, err
		}
		c.Config = maps
	}
	return true, nil
}
//文件基本数据填充
//@upload_type 上传配置
func (c *UploadFile)SetUploadFileData(upload_type string, file multipart.File, header *multipart.FileHeader) (bool, error) {
	var spe string
	if os.IsPathSeparator('\\') {
		//前边的判断是否是系统的分隔符
		spe = "\\"
	} else {
		spe = "/"
	}
	//配置
	if upload_type == "" {
		upload_type = "upload_default"
	}
	//配置检测
	if _, err := c.SetConfig(upload_type); err != nil {
		return false, err
	}
	root_path := ""
	root_path = c.Config["root_path"]
	if root_path == "" {
		root_path = "/uploads/image/"
	}
	//年月
	ym := datetime.Format(time.Now(), "2006_01")
	str := header.Filename + strconv.FormatInt(time.Now().UnixNano(), 10)
	//fmt.Println("Filename", header.Filename)
	//fmt.Println("md5", str)
	//文件大小
	statInterface, ok := file.(Stat)
	if ok {
		fileInfo, _ := statInterface.Stat()
		//文件大小
		c.Size = int(fileInfo.Size())
	} else {
		num := file.(Size).Size()
		num_int := int(num)
		if num_int > 0 {
			//文件大小
			c.Size = num_int
		} else {
			return false, &fox.Error{Msg:"文件错误."}
		}
	}
	//文件后缀
	c.Ext = path.Ext(header.Filename)
	//原文件名
	c.NameOriginal = header.Filename
	//新文件名
	c.Name = crypt.Md5(str) + c.Ext
	//保存目录
	c.Path = root_path + ym + spe
	//文件地址
	c.Url = c.Path + c.Name
	//删除 文件后缀 中的点号
	c.Ext = strings.Replace(c.Ext, ".", "", -1)
	http := c.Config["http"]
	c.Http = http + strings.Replace(c.Url, "/", "", 1)
	fmt.Println("文件数据：", c)

	return true, nil
}
//审核
func (c *UploadFile)Check() (bool, error) {
	if len(c.Config) == 0 {
		return false, &fox.Error{Msg:"Config 没有配置"}
	}
	//后缀是否找到
	isFind := false
	extArr := strings.Split(c.Config["ext"], ",")
	for _, v := range extArr {
		if v == c.Ext {
			isFind = true
		}
	}
	//检测后缀 不在上传文件中报错
	if !isFind {
		return false, &fox.Error{Msg:"此文件不在允许上传范围内"}
	}
	//文件大小
	size, _ := strconv.Atoi(c.Config["size"])
	if (size > 0) {
		if c.Size > size {
			return false, &fox.Error{Msg:"文件大小 超过上传限制"}
		}
	}
	//检测文件大小
	return true, nil
}
//本地保存
func (c *UploadFile)LocalSaveFile(file multipart.File, path, file_name string) (bool, error) {
	//当前的目录
	dir, _ := os.Getwd()
	fmt.Println("当前的目录", dir)
	url := path + file_name
	//判断目录
	isOk, _ := PathExists(dir + path)
	if !isOk {
		err := os.Mkdir(dir + path, os.ModePerm)  //在当前目录下生成目录
		if err != nil {
			fmt.Println("创建目录失败", err)
			return false, &fox.Error{Msg:"目录创建不成功！" + path}
		}
		fmt.Println("创建目录" + dir + path + "成功")
	}

	f, err := os.OpenFile(dir + url, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("写入文件失败", err)
		return false, &fox.Error{Msg:"文件写入不成功！" + url}
	}
	defer f.Close()
	w, err := io.Copy(f, file)
	fmt.Println("io.Copy", w, err)
	fmt.Println("写入文件" + dir + url + "成功")
	return true, nil
}
//本地临时文件
func (c *UploadFile)LocalTmpFileRemove() error {
	root_path_tmp := c.GetLocalTmpPath()
	//当前的目录
	dir, _ := os.Getwd()
	fmt.Println("当前的目录", dir)
	file := dir + root_path_tmp + c.Name
	fmt.Println("删除的文件", file)
	err := os.Remove(file)               //删除文件test.txt
	if err != nil {
		return err
	} else {
		return nil
	}
	return nil
}
//本地临时文件路径
func (c *UploadFile)GetLocalTmpPath() string {
	root_path_tmp := c.Config["root_path_tmp"]
	if root_path_tmp == "" {
		root_path_tmp = "/uploads/tmp"
	}
	return root_path_tmp
}
//保存到数据库
func (c *UploadFile)SaveDataBase(maps map[string]interface{}) {
	//
	att := model.NewAttachment()
	att.Ext = c.Ext
	att.Size = c.Size
	att.Name = c.Name
	att.NameOriginal = c.NameOriginal
	att.Path = c.Path
	att.Md5 = crypt.Md5(c.Path + c.Name)

	fmt.Println("maps=>", maps)
	//其他字段
	if maps["type_id"] != nil {
		type_id, _ := number.ObjToInt(maps["type_id"])
		//fmt.Println("反射类型reflect.Type",reflect.TypeOf(maps["type_id"]))
		att.TypeId = type_id
	}
	if maps["aid"] != nil {
		aid, _ := number.ObjToInt(maps["aid"])
		att.Aid = aid
		//fmt.Println("att.TypeId",att.Aid)
	}
	if maps["id"] != nil {
		id, _ := number.ObjToInt(maps["id"])
		att.Id = id
		//fmt.Println("att.Id",att.Id)
	}
	c.Attachment = att
	//fmt.Println("att",att)
	o := db.NewDb()
	_, err := o.Insert(c.Attachment)
	if err != nil {
		fmt.Println("保存到数据库失败", err)
	}
}
