package blog
//同步
type BlogSync struct {

}
//快速初始化
func NewBlogSyncService() *BlogSync {
	return new(BlogSync)
}