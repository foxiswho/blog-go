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
)
func Upload(field string, r *http.Request) (bool, error) {
	file, header, err := r.FormFile(field)
	if err != nil {
		return false, err
	}
	var spe string
	if os.IsPathSeparator('\\') {
		//前边的判断是否是系统的分隔符
		spe = "\\"
	} else {
		spe = "/"
	}
	root_path, err := beego.GetConfig("upload_default", "root_path", "/uploads/image/")
	if err != nil {
		return false, err
	}
	//年月
	ym := datetime.Format(time.Now(), "2006_01")

	str:=header.Filename+strconv.FormatInt(time.Now().UnixNano(),10)
	fmt.Println("Filename",header.Filename)
	fmt.Println("md5",str)
	md5 := crypt.Md5(str)
	ext := path.Ext(header.Filename)
	//当前的目录
	dir, _ := os.Getwd()
	fmt.Println("当前的目录", dir)
	//保存目录
	file_path := root_path.(string) + ym + spe
	//文件地址
	file_new := file_path + md5 + ext
	//判断目录
	ok,err:=PathExists(dir + file_path)
	if !ok {
		err = os.Mkdir(dir + file_path, os.ModePerm)  //在当前目录下生成目录
		if err != nil {
			fmt.Println("创建目录失败",err)
			return false, &util.Error{Msg:"目录创建不成功！"+file_path}
		}
		fmt.Println("创建目录" + dir + file_path + "成功")
	}
	defer file.Close()
	f, err := os.OpenFile(dir + file_new, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("写入文件失败",err)
		return false, &util.Error{Msg:"文件写入不成功！"+file_new}
	}
	fmt.Println("写入文件" + dir + file_new + "成功")
	defer f.Close()
	io.Copy(f, file)
	return true, nil
}
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