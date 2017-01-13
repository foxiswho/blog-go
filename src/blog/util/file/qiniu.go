package file

import (
	"qiniupkg.com/api.v7/kodo"
	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodocli"
	"fmt"
	"github.com/astaxie/beego"
	"blog/util"
	"time"
	"blog/util/cache"
	"mime/multipart"
	"strings"
	"os"
)

type QiNiu struct {
	Config map[string]string `json:"-"`
}


// 构造返回值字段
type PutRet struct {
	Hash string `json:"hash"`
	Key  string `json:"key"`
}

func NewQiNiu() *QiNiu {
	return new(QiNiu)
}
//获取配置
func (t *QiNiu)setConfig() (bool, error) {
	maps, err := beego.AppConfig.GetSection("qiniu")
	if err != nil {
		return false, err
	}
	t.Config = maps
	return true, nil
}
//七牛配置读取
func (t *QiNiu)SetQiNiuConfig() (bool, error) {
	ok, err := t.setConfig()
	if err != nil {
		fmt.Println("setConfig err:", err)
		return false, err
	}
	fmt.Println("setConfig", ok)
	if len(t.Config) < 1 {
		return false, &util.Error{Msg:"配置文件没有读取"}
	}
	// 初始化AK，SK
	conf.ACCESS_KEY = t.Config["access_key"]
	conf.SECRET_KEY = t.Config["secret_key"]
	return true, nil
}
//设置 token 缓存
func (t *QiNiu)setToken() string {
	// 设置上传到的空间
	bucket := t.Config["bucket"]
	// 创建一个Client
	c := kodo.New(0, nil)
	// 设置上传的策略
	policy := &kodo.PutPolicy{
		Scope:   bucket,
		//设置Token过期时间
		Expires: 3600,
	}
	// 生成一个上传token
	token := c.MakeUptoken(policy)
	//缓存
	err := cache.Cache.Put("qiniu_token", token, 3600 * time.Second)
	if err!=nil{
		fmt.Println("设置缓存错误",err)
	}
	return token
}
//获取
func (t *QiNiu)GetToken() string {
	if !cache.Cache.IsExist("qiniu_token") {
		return t.setToken()
	}
	token := cache.Cache.Get("qiniu_token")
	tmp:=token.(string)
	if len(tmp)==0{
		return t.setToken()
	}
	return tmp
}
//上传
func (t *QiNiu)Upload(file multipart.File, UploadFile *UploadFile) (interface{}, error) {
	//七牛配置填充
	_, err := t.SetQiNiuConfig()
	if err != nil {
		fmt.Println("getConfig err:", err)
		return nil, &util.Error{Msg:"七牛配置错误:" + err.Error()}
	}
	//token
	token := t.GetToken()
	if len(token)<1{
		fmt.Println("token",token)
		return nil,&util.Error{Msg:"七牛token不能为空"}
	}
	// 构建一个uploader
	uploader := kodocli.NewUploader(0, nil)
	//当前的目录
	dir, _ := os.Getwd()
	fmt.Println("当前的目录", dir)
	var ret PutRet
	// 设置上传文件的路径
	key := UploadFile.Path + UploadFile.Name
	filePath := dir+UploadFile.GetLocalTmpPath()+UploadFile.Name
	key=strings.Replace(key,"/","",1)
	fmt.Println("本地文件绝对路径", filePath)
	fmt.Println("去除第一个字符/后，访问路径",key)
	// 调用PutFileWithoutKey方式上传，没有设置saveasKey以文件的hash命名
	err = uploader.PutFile(nil, &ret, token, key, filePath, nil)
	// 打印出错信息
	if err != nil {
		fmt.Println("io.Put failed:", err)
		return nil, err
	}
	// 打印返回的信息
	fmt.Println("七牛成功后返回信息",ret)
	return ret, nil
}
