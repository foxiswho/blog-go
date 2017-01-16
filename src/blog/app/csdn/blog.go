package csdn

import (
	mod "blog/service/blog"
	"blog/fox"
	"fmt"
	"blog/app/csdn/blog"
	"strconv"
)

type Blog struct {

}
//更新
func (t *Blog) Update(b *mod.Blog, id string) error {
	if len(id) < 1 {
		return fox.Error{Msg:"id 不能为空"}
	}
	web := NewAuthorizeWeb()
	acc, err := web.GetAccessTokenCache()
	if err != nil {
		fmt.Println(err)
		return err
	}
	art := blog.NewSaveArticle()
	art.AccessToken = acc.AccessToken
	art.Tags = b.Tags
	art.Id, _ = strconv.Atoi(id)
	art.Content = b.Content
	art.Description = b.Description
	art.Title = b.Title
	str, err := art.Post()
	if err != nil {
		return err
	}
	fmt.Println(str)
	return nil
}
//创建
func (t *Blog) Create(b *mod.Blog, id string) error {
	if len(id) < 1 {
		return fox.Error{Msg:"id 不能为空"}
	}
	web := NewAuthorizeWeb()
	acc, err := web.GetAccessTokenCache()
	if err != nil {
		fmt.Println(err)
		return err
	}
	art := blog.NewSaveArticle()
	art.AccessToken = acc.AccessToken
	art.Tags = b.Tags
	art.Id, _ = strconv.Atoi(id)
	art.Content = b.Content
	art.Description = b.Description
	art.Title = b.Title
	str, err := art.Post()
	if err != nil {
		return err
	}
	fmt.Println(str)
	return nil
}