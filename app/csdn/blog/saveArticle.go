package blog

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/client/httplib"
	"github.com/foxiswho/blog-go/app/csdn/conf"
	"github.com/foxiswho/blog-go/app/csdn/entity"
	"github.com/foxiswho/blog-go/fox"
	"strconv"
	"strings"
	"time"
)
//发表/修改文章
type SaveArticle struct {
	AccessToken string `json:"access_token" 是 OAuth授权后获得`
	Ip          string `json:"ip" 否 用户ip`
	*entity.Article
}
//初始化
func NewSaveArticle() *SaveArticle {
	return new(SaveArticle)
}
//设置类型
func (t *SaveArticle)SetType(str string) (error) {
	if str == "original" || str == "report" || str == "translated" {
		t.Type = str
		return nil
	}
	return fox.NewError("type 不能为空，original|report|translated")
}
//检测
func (t *SaveArticle)Check() (error) {
	if len(t.AccessToken) < 1 {
		return fox.NewError("access_token 不能为空")
	}
	if len(t.Title) < 1 {
		return fox.NewError("title 不能为空")
	}
	if len(t.Type) < 1 {
		return fox.NewError("type 不能为空，original|report|translated")
	}
	if len(t.Description) < 1 {
		return fox.NewError("description 不能为空")
	}
	if len(t.Content) < 1 {
		return fox.NewError("content 不能为空")
	}
	if len(t.Tags) < 1 {
		return fox.NewError("tags 不能为空")
	}
	if err := t.SetType(t.Type); err != nil {
		return err
	}
	return nil
}
//发送
func (t *SaveArticle)Post() (*entity.Article, error) {
	err := t.Check()
	if err != nil {
		return nil, err
	}
	//接口传输数据
	req := httplib.Post(conf.BLOG_SAVE_URL)
	//超时
	req.SetTimeout(100 * time.Second, 30 * time.Second)
	//参数
	req.Param("access_token", t.AccessToken)
	req.Param("id", strconv.Itoa(t.Id))
	req.Param("title", t.Title)
	req.Param("type", t.Type)
	req.Param("description", t.Description)
	req.Param("content", t.Content)
	//req.Param("markdowncontent", t.Content)
	//req.Param("markdowndirectory", "")
	req.Param("articleedittype", "1")
	req.Param("categories", t.Categories)
	req.Param("tags", t.Tags)
	req.Param("ip", t.Ip)
	//返回
	s, err := req.String()
	if err != nil {
		fmt.Println("返回错误信息：", err)
		return nil, fox.NewError("返回错误信息：" + err.Error())
	}
	//是否错误代码
	if strings.Contains(s, "error_code") {
		fmt.Println("返回错误信息：", s)
		return nil, fox.NewError("返回错误信息：" + s)
	}
	fmt.Println("返回内容：", s)
	//json 序列化 成  结构体
	var saveArticle *entity.Article
	if err := json.Unmarshal([]byte(s), &saveArticle); err != nil {
		return nil, fox.NewError("反序列化失败：" + err.Error())
	}
	fmt.Println("反序列化：", saveArticle)
	return saveArticle, nil
}
