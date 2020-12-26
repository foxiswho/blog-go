package csdn

import (
	"github.com/foxiswho/blog-go/blog/app/csdn/blog"
	"github.com/foxiswho/blog-go/blog/app/csdn/entity"
	"github.com/foxiswho/blog-go/blog/fox"
	"github.com/foxiswho/blog-go/blog/fox/db"
	"github.com/foxiswho/blog-go/blog/fox/editor"
	"github.com/foxiswho/blog-go/blog/model"
	mod "github.com/foxiswho/blog-go/blog/service/blog"
	"fmt"
	"strconv"
	"time"
)
//博客
type Blog struct {

}
//初始化
func NewCsdnBlogApp() *Blog {
	return new(Blog)
}
//更新
func (t *Blog) Update(b *mod.Blog, type_id int, id int) error {
	fmt.Println("CSDN 更新文章")
	//fmt.Println("更新内容",b.Blog)
	if id < 1 {
		return fox.NewError("id 不能为空")
	}
	//初始化
	web := NewAuthorizeWeb()
	//获取缓存  及判断
	acc, err := web.GetAccessTokenCache()
	if err != nil {
		fmt.Println("令牌缓存获取失败", err)
		return err
	}
	fmt.Println("令牌缓存获取成功")
	//初始化 及赋值
	art := blog.NewSaveArticle()
	art.Article = &entity.Article{}
	art.AccessToken = acc.AccessToken
	if len(b.Blog.Tag) > 0 {
		art.Tags = b.Blog.Tag
	}
	art.Id = id
	//csdn的人 真懒，不支持Markdown格式文本
	//转换为 富文本格式内容
	art.Content = string(editor.Markdown([]byte(b.Content)))
	if len(b.Blog.Description) > 0 {
		art.Description = b.Blog.Description
	}
	art.Title = b.Title
	art.Type = "original"
	//接口传输
	str, err := art.Post()
	if err != nil {
		return err
	}
	fmt.Println("返回：", str)
	maps := make(map[string]interface{})
	maps["blog_id"] = b.Blog.BlogId
	maps["id"] = id
	maps["type_id"] = type_id
	//初始化及数据库查询
	m := model.NewBlogSyncMapping()
	ok, err := db.Filter(maps).Get(m)
	if err != nil {
		return fox.NewError("更新错误:" + err.Error())
	}
	fmt.Println("查询状态", ok)
	//更新
	if m.MapId > 0 {
		m.IsSync = 1
		m.TimeUpdate = time.Now()
		_, err := db.NewDb().ID(m.MapId).Update(m)
		fmt.Println("保存状态", err)
	} else {
		m.Id = strconv.Itoa(id)
		//插入
		m.BlogId = b.Blog.BlogId
		m.TypeId = type_id
		m.IsSync = 1
		_, err := db.NewDb().Insert(m)
		fmt.Println("保存状态", err)
	}
	return nil
}
//创建
func (t *Blog) Create(b *mod.Blog, type_id int)(string,error){
	maps := make(map[string]interface{})
	maps["blog_id"] = b.Blog.BlogId
	maps["type_id"] = type_id
	//初始化及查询
	m := model.NewBlogSyncMapping()
	ok, err := db.Filter(maps).Get(m)
	if err != nil {
		return "",fox.NewError("更新错误:" + err.Error())
	}
	fmt.Println("查询状态", ok)
	if m.MapId > 0 {
		return "",fox.NewError("已存在此数据，不能重复插入")
	}
	//初始化
	web := NewAuthorizeWeb()
	//获取token缓存
	acc, err := web.GetAccessTokenCache()
	if err != nil {
		fmt.Println("令牌缓存获取失败", err)
		return "",err
	}
	fmt.Println("令牌缓存获取成功")
	//初始化
	art := blog.NewSaveArticle()
	art.Article = &entity.Article{}
	art.AccessToken = acc.AccessToken
	art.Tags = b.Blog.Tag
	art.Id = 0
	//csdn的人 真懒，不支持Markdown格式文本
	//转换为 富文本格式内容
	art.Content = string(editor.Markdown([]byte(b.Content)))
	if len(b.Blog.Description) > 0 {
		art.Description = b.Blog.Description
	}
	art.Title = b.Title
	art.Type = "original"
	//接口传输
	str, err := art.Post()
	if err != nil {
		return "",err
	}
	fmt.Println("返回：", str)
	//更新
	if m.MapId > 0 {
		m.IsSync = 1
		m.TimeUpdate = time.Now()
		_, err := db.NewDb().ID(m.MapId).Update(m)
		fmt.Println("保存状态", err)
	} else {
		m.Id = strconv.Itoa(str.Id)
		//插入
		m.BlogId = b.Blog.BlogId
		m.TypeId = type_id
		m.IsSync = 1
		_, err := db.NewDb().Insert(m)
		fmt.Println("保存状态", err)
	}
	return m.Id,nil
}
//获取
func (t *Blog) Read(id string) (*mod.Blog, error) {
	if len(id) < 1 {
		return nil, fox.NewError("id 不能为空")
	}
	//初始化
	web := NewAuthorizeWeb()
	//获取缓存
	acc, err := web.GetAccessTokenCache()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("令牌缓存获取成功")
	//初始化及赋值
	art := blog.NewGetArticle()
	art.AccessToken = acc.AccessToken
	art.Id, _ = strconv.Atoi(id)
	//接口数据传输
	str, err := art.Post()
	if err != nil {
		return nil, err
	}
	fmt.Println("内容获取成功")
	//初始化 赋值
	modBlog := model.NewBlog()
	modBlog.Content = str.Content
	if len(str.Description) > 1 {
		modBlog.Description = str.Description
	}
	if len(str.Tags) > 0 {
		modBlog.Tag = str.Tags
	}
	if len(str.Title) > 0 {
		modBlog.Title = str.Title
	}
	b := mod.NewBlogService()
	b.Blog = modBlog
	b.BlogStatistics = model.NewBlogStatistics()
	//fmt.Println("内容", modBlog)
	fmt.Println("内容", b.Title)
	return b, nil
}
//根据id更新或插入记录
func Get(type_id, id, blog_id int) (*mod.Blog, error) {
	//初始化
	csdn := NewCsdnBlogApp()
	//接口传输获取内容
	b, err := csdn.Read(strconv.Itoa(id))
	if err != nil {
		return nil, err
	}
	//保存
	if blog_id > 0 {
		maps := make(map[string]interface{})
		maps["blog_id"] = blog_id
		maps["id"] = id
		maps["type_id"] = type_id
		m := model.NewBlogSyncMapping()
		_, err := db.Filter(maps).Get(m)
		if err != nil {
			return nil, fox.NewError("更新错误:" + err.Error())
		}
		//更新
		if m.MapId > 0 {
			m.IsSync = 1
			m.TimeUpdate = time.Now()
			_, err := db.NewDb().ID(m.MapId).Update(m)
			fmt.Println("保存状态", err)
		} else {
			//插入
			m.BlogId = blog_id
			m.TypeId = type_id
			m.Id = strconv.Itoa(id)
			m.IsSync = 1
			_, err := db.NewDb().Insert(m)
			fmt.Println("保存状态", err)
		}
	} else {
		m := model.NewBlogSyncMapping()
		//插入
		m.BlogId = blog_id
		m.TypeId = type_id
		m.Id = strconv.Itoa(id)
		m.IsSync = 1
		_, err := db.NewDb().Insert(m)
		fmt.Println("保存状态", err)
	}
	return b, nil
}
