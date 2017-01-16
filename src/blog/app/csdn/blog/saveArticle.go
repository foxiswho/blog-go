package blog

import (
	"github.com/astaxie/beego/httplib"
	"time"
	"blog/fox"
	"blog/app/csdn/entity"
	"blog/app/csdn/conf"
	"strconv"
	"fmt"
	"strings"
	"encoding/json"
)
//发表/修改文章
type SaveArticle struct {
	AccessToken string `json:"access_token" 是 OAuth授权后获得`
	Ip          string `json:"ip" 否 用户ip`
	*entity.Article
}

func NewSaveArticle() *SaveArticle {
	return new(SaveArticle)
}
func (t *SaveArticle)SetType(str string) (error) {
	if str == "original" || str == "report" || str == "translated" {
		return nil
	}
	return &fox.Error{Msg:"type 不能为空，original|report|translated"}
}
func (t *SaveArticle)Check() (error) {
	if len(t.AccessToken) < 1 {
		return &fox.Error{Msg:"access_token 不能为空"}
	}
	if len(t.Title) < 1 {
		return &fox.Error{Msg:"title 不能为空"}
	}
	if len(t.Type) < 1 {
		return &fox.Error{Msg:"type 不能为空，original|report|translated"}
	}
	if len(t.Description) < 1 {
		return &fox.Error{Msg:"description 不能为空"}
	}
	if len(t.Content) < 1 {
		return &fox.Error{Msg:"content 不能为空"}
	}
	if len(t.Tags) < 1 {
		return &fox.Error{Msg:"tags 不能为空"}
	}
	return nil
}
//发送
func (t *SaveArticle)Post() (*entity.Article, error) {
	err := t.Check()
	if err != nil {
		return nil, err
	}
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
		return nil, &fox.Error{Msg:"返回错误信息：" + err.Error()}
	}
	//是否错误代码
	if strings.Contains(s, "error_code") {
		fmt.Println("返回错误信息：", s)
		return nil, &fox.Error{Msg:"返回错误信息：" + s}
	}
	fmt.Println("返回内容：", s)
	var saveArticle *entity.Article
	if err := json.Unmarshal([]byte(s), &saveArticle); err != nil {
		return nil, &fox.Error{Msg:"反序列化失败：" + err.Error()}
	}
	fmt.Println("反序列化：", saveArticle)
	return saveArticle, nil
}