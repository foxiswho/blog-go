package file

import (
	"os"
	"github.com/astaxie/beego"
	"fox/util/datetime"
	"time"
	"net/http"
	"fox/util/crypt"
	"path"
	"fmt"
	"fox/util"
	"io"
	"strconv"
	"strings"
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
}
// 获取文件信息的接口
type Stat interface {
	Stat() (os.FileInfo, error)
}
//上传
func Upload(field string, r *http.Request) (*UploadFile, error) {
	file, header, err := r.FormFile(field)
	if err != nil {
		return nil, err
	}
	UploadFile := &UploadFile{}
	var spe string
	if os.IsPathSeparator('\\') {
		//前边的判断是否是系统的分隔符
		spe = "\\"
	} else {
		spe = "/"
	}
	root_path, err := beego.GetConfig("upload_default", "root_path", "/uploads/image/")
	if err != nil {
		return nil, err
	}

	//年月
	ym := datetime.Format(time.Now(), "2006_01")

	str := header.Filename + strconv.FormatInt(time.Now().UnixNano(), 10)
	fmt.Println("Filename", header.Filename)
	fmt.Println("md5", str)
	if statInterface, ok := file.(Stat); ok {
		fileInfo, _ := statInterface.Stat()
		//文件大小
		UploadFile.Size = int(fileInfo.Size())
	}
	//文件后缀
	UploadFile.Ext = path.Ext(header.Filename)
	//原文件名
	UploadFile.NameOriginal = header.Filename
	//新文件名
	UploadFile.Name = crypt.Md5(str) + UploadFile.Ext
	//保存目录
	UploadFile.Path = root_path.(string) + ym + spe
	//文件地址
	UploadFile.Url = UploadFile.Path + UploadFile.Name
	fmt.Println("文件数据：", UploadFile)
	//当前的目录
	dir, _ := os.Getwd()
	fmt.Println("当前的目录", dir)
	//判断目录
	ok, _ := PathExists(dir + UploadFile.Path)
	if !ok {
		err = os.Mkdir(dir + UploadFile.Path, os.ModePerm)  //在当前目录下生成目录
		if err != nil {
			fmt.Println("创建目录失败", err)
			return nil, &util.Error{Msg:"目录创建不成功！" + UploadFile.Path}
		}
		fmt.Println("创建目录" + dir + UploadFile.Path + "成功")
	}
	defer file.Close()
	f, err := os.OpenFile(dir + UploadFile.Url, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("写入文件失败", err)
		return nil, &util.Error{Msg:"文件写入不成功！" + UploadFile.Url}
	}
	defer f.Close()
	w, err := io.Copy(f, file)
	fmt.Println("io.Copy", w, err)
	fmt.Println("写入文件" + dir + UploadFile.Url + "成功")
	//最后处理
	//删除 文件后缀 中的点号
	UploadFile.Ext = strings.Replace(UploadFile.Ext, ".", "", -1)
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