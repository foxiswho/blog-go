package file

import (
	"github.com/foxiswho/blog-go/blog/fox"
	"github.com/foxiswho/blog-go/blog/fox/cache"
	"github.com/foxiswho/blog-go/blog/fox/config"
	"context"
	"fmt"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"mime/multipart"
	"os"
	"strings"
	"time"
)

//七牛云存储
type QiNiu struct {
	Config map[string]string `json:"-"`
}

// 构造返回值字段
type PutRet struct {
	Hash string `json:"hash"`
	Key  string `json:"key"`
}

//初始化
func NewQiNiu() *QiNiu {
	return new(QiNiu)
}

//获取配置
func (t *QiNiu) setConfig() (bool, error) {
	maps, err := config.GetSection("qiniu")
	if err != nil {
		return false, err
	}
	t.Config = maps
	return true, nil
}

//七牛配置读取
func (t *QiNiu) SetQiNiuConfig() (bool, error) {
	ok, err := t.setConfig()
	if err != nil {
		fmt.Println("setConfig err:", err)
		return false, err
	}
	fmt.Println("setConfig", ok)
	if len(t.Config) < 1 {
		return false, fox.NewError("配置文件没有读取")
	}
	return true, nil
}

//设置 token 缓存
func (t *QiNiu) setToken() string {
	// 初始化AK，SK
	putPolicy := storage.PutPolicy{
		Scope:   t.Config["bucket"], // 设置上传到的空间
		Expires: 3600,               //设置Token过期时间
	}
	mac := qbox.NewMac(t.Config["access_key"], t.Config["secret_key"])
	// 生成一个上传token
	token := putPolicy.UploadToken(mac)
	//缓存
	err := cache.Put("qiniu_token", token, 3600*time.Second)
	if err != nil {
		fmt.Println("设置缓存错误", err)
	}
	return token
}

//获取
func (t *QiNiu) GetToken() string {
	if !cache.IsExist("qiniu_token") {
		return t.setToken()
	}
	token := cache.Get("qiniu_token")
	tmp := token.(string)
	if len(tmp) == 0 {
		return t.setToken()
	}
	return tmp
}

//上传
func (t *QiNiu) Upload(file multipart.File, UploadFile *UploadFile) (interface{}, error) {
	//七牛配置填充
	_, err := t.SetQiNiuConfig()
	if err != nil {
		fmt.Println("getConfig err:", err)
		return nil, fox.NewError("七牛配置错误:" + err.Error())
	}
	//token
	token := t.GetToken()
	if len(token) < 1 {
		fmt.Println("token", token)
		return nil, fox.NewError("七牛token不能为空")
	}
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{}

	//当前的目录
	dir, _ := os.Getwd()
	fmt.Println("当前的目录", dir)

	// 设置上传文件的路径
	key := UploadFile.Path + UploadFile.Name
	filePath := dir + UploadFile.GetLocalTmpPath() + UploadFile.Name
	key = strings.Replace(key, "/", "", 1)
	fmt.Println("本地文件绝对路径", filePath)
	fmt.Println("去除第一个字符/后，访问路径", key)
	// 调用PutFileWithoutKey方式上传，没有设置saveasKey以文件的hash命名
	err = formUploader.PutFile(context.Background(), &ret, token, key, filePath, &putExtra)
	// 打印出错信息
	if err != nil {
		fmt.Println("io.Put failed:", err)
		return nil, err
	}
	// 打印返回的信息
	fmt.Println("七牛成功后返回信息", ret)
	return ret, nil
}
