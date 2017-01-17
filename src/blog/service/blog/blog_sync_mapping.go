package blog

import (
	"blog/model"
	"blog/fox/db"
	"blog/fox"
)
//同步第三方博客
type BlogSyncMapping struct {

}
//快速初始化
func NewBlogSyncMappingService() *BlogSyncMapping {
	return new(BlogSyncMapping)
}
//根据博客ID和类别ID 获取 第三方ID数据
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