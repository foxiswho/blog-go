package blog

import (
	"blog/model"
	"blog/fox/db"
	"blog/fox"
)

type BlogSyncMapping struct {

}

func NewBlogSyncMappingService() *BlogSyncMapping {
	return new(BlogSyncMapping)
}
func (t *BlogSyncMapping)GetAppId(type_id, blog_id int) (*model.BlogSyncMapping, error) {
	maps := make(map[string]interface{})
	maps["blog_id"] = blog_id
	maps["type_id"] = type_id
	m := model.NewBlogSyncMapping()
	_, err := db.Filter(maps).Get(m)
	if err != nil {
		return nil,fox.NewError("不存在")
	}
	if m.MapId < 1 {
		return nil,fox.NewError("不存在")
	}
	return m, nil
}