package csdn

import (
	mod "blog/service/blog"
	"blog/fox"
	"fmt"
	"blog/app/csdn/blog"
	"strconv"
	"blog/model"
	"blog/fox/db"
	"time"
	"blog/app/csdn/entity"
)

type Blog struct {

}

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
	web := NewAuthorizeWeb()
	acc, err := web.GetAccessTokenCache()
	if err != nil {
		fmt.Println("令牌缓存获取失败", err)
		return err
	}
	fmt.Println("令牌缓存获取成功")
	art := blog.NewSaveArticle()
	art.Article = &entity.Article{}
	art.AccessToken = acc.AccessToken
	if len(b.Blog.Tag)>0{
		art.Tags = b.Blog.Tag
	}
	art.Id = id
	art.Content = b.Content
	if len(b.Blog.Description)>0{
		art.Description = b.Blog.Description
	}
	art.Title = b.Title
	art.Type = "original"
	str, err := art.Post()
	if err != nil {
		return err
	}
	fmt.Println("返回：", str)
	maps := make(map[string]interface{})
	maps["blog_id"] = b.Blog.BlogId
	maps["id"] = id
	maps["type_id"] = type_id
	m := model.NewBlogSyncMapping()
	ok, err := db.Filter(maps).Get(m)
	if err != nil {
		return fox.NewError("更新错误:"+err.Error())
	}
	fmt.Println("查询状态", ok)
	//更新
	if m.MapId > 0 {
		m.IsSync = 1
		m.TimeUpdate = time.Now()
		_, err := db.NewDb().Id(m.MapId).Update(m)
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
func (t *Blog) Create(b *mod.Blog, type_id int) error {
	maps := make(map[string]interface{})
	maps["blog_id"] = b.Blog.BlogId
	maps["type_id"] = type_id
	m := model.NewBlogSyncMapping()
	ok, err := db.Filter(maps).Get(m)
	if err != nil {
		return fox.NewError("更新错误:"+err.Error())
	}
	fmt.Println("查询状态", ok)
	if m.MapId > 0 {
		return fox.NewError("已存在此数据，不能重复插入")
	}
	web := NewAuthorizeWeb()
	acc, err := web.GetAccessTokenCache()
	if err != nil {
		fmt.Println("令牌缓存获取失败", err)
		return err
	}
	fmt.Println("令牌缓存获取成功")
	art := blog.NewSaveArticle()
	art.Article = &entity.Article{}
	art.AccessToken = acc.AccessToken
	art.Tags = b.Blog.Tag
	art.Id = 0
	art.Content = b.Content
	art.Description = b.Description
	art.Title = b.Title
	art.Type = "original"
	str, err := art.Post()
	if err != nil {
		return err
	}
	fmt.Println("返回：", str)
	//更新
	if m.MapId > 0 {
		m.IsSync = 1
		m.TimeUpdate = time.Now()
		_, err := db.NewDb().Id(m.MapId).Update(m)
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
	return nil
}
//获取
func (t *Blog) Read(id string) (*mod.Blog, error) {
	if len(id) < 1 {
		return nil, fox.NewError("id 不能为空")
	}
	web := NewAuthorizeWeb()
	acc, err := web.GetAccessTokenCache()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("令牌缓存获取成功")
	art := blog.NewGetArticle()
	art.AccessToken = acc.AccessToken
	art.Id, _ = strconv.Atoi(id)
	str, err := art.Post()
	if err != nil {
		return nil, err
	}
	fmt.Println("内容获取成功")
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
func Get(type_id, id, blog_id int) (*mod.Blog, error) {
	csdn := NewCsdnBlogApp()
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
			return nil,fox.NewError("更新错误:"+err.Error())
		}
		//更新
		if m.MapId > 0 {
			m.IsSync = 1
			m.TimeUpdate = time.Now()
			_, err := db.NewDb().Id(m.MapId).Update(m)
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