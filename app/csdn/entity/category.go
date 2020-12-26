package entity
//获取博主的自定义分类
type Category struct {
	Id int `json:"id" 类别id`
	Name string `json:"name" 类别名称`
	Hide bool `json:"hide" 是否隐藏`
	ArticleCount int `json:"article_count" 类别下的文章数量`
}
