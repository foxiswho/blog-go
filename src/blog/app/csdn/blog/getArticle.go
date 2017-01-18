package blog

import (
	"blog/app/csdn/entity"
	"blog/fox"
	"github.com/astaxie/beego/httplib"
	"time"
	"strconv"
	"blog/app/csdn/conf"
	"encoding/json"
	"fmt"
)
//获取文章
type GetArticle struct {
	AccessToken string `json:"access_token" 是 OAuth授权后获得`
	Id          int `json:"id" 是 文章id`
}
//初始化
func NewGetArticle() *GetArticle {
	return new(GetArticle)
}
//检测
func (t *GetArticle)Check() (error) {
	if len(t.AccessToken) < 1 {
		return fox.NewError("access_token 不能为空")
	}
	if t.Id < 1 {
		return fox.NewError("文章id 不能为空")
	}
	return nil
}
//发送
func (t *GetArticle)Post() (*entity.Article, error) {
	err := t.Check()
	if err != nil {
		return nil, err
	}
	fmt.Println("验证通过")
	req := httplib.Post(conf.BLOG_ID_URL)
	//超时
	req.SetTimeout(100 * time.Second, 30 * time.Second)
	//参数
	req.Param("access_token", t.AccessToken)
	req.Param("id", strconv.Itoa(t.Id))
	//返回
	s, err := req.String()
	if err != nil {
		return nil, err
	}
	fmt.Println("通信返回：",s)
	var art *entity.Article
	err = json.Unmarshal([]byte(s), &art)
	if err != nil {
		fmt.Println("内容反序列化 错误：",err)
		return nil, fox.NewError("内容反序列化 错误："+err.Error())
	}
	if len(art.MarkdownContent)>0{
		art.Content=art.MarkdownContent
	}
	fmt.Println("反序列化：",art)
	return art, nil
}