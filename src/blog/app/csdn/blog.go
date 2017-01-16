package csdn

import (
	mod "blog/service/blog"

	"fmt"
	"blog/app/csdn/blog"
	"blog/fox"
)

type Blog struct {

}
//更新
func (t *Blog) Update(b *mod.Blog, id string) (error) {
	if len(id) < 1 {
		return &fox.Error{Msg:"id 不能为空"}
	}
	web := NewAuthorizeWeb()
	acc, err := web.GetAccessTokenCache()
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(acc)
	art := blog.NewSaveArticle()
	//art.AccessToken = acc.AccessToken
	//art.Tags = b.Tags
	//art.Id, _ = strconv.Atoi(id)
	//art.Content = b.Content
	//art.Description = b.Description
	//art.Title = b.Title
	//str, err := art.Post()
	//if err != nil {
	//	return err
	//}
	fmt.Println(art)
	return nil
}
