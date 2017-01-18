package entity
//频道
type Channel struct {
	Id    int `json:"id" 类别id`
	Name  string `json:"name" 类别名称`
	Alias string `json:"alias" 分类别名`
}
