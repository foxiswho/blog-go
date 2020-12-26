package conf

var ACCESS_KEY string
var SECRET_KEY string

const (
	//Web应用的验证授权
	WEB_URL = "http://api.csdn.net/oauth2/authorize"
	//客户端的验证授权
	APP_URL = "http://api.csdn.net/oauth2/access_token"
	//发表/修改文章
	BLOG_SAVE_URL="http://api.csdn.net/blog/savearticle"
	//获取博主的文章列表
	BLOG_LIST_URL="http://api.csdn.net/blog/getarticlelist"
	//博客频道分类
	BLOG_CHANNEL_URL="http://api.csdn.net/blog/getchannel"
	//博客自定义分类
	BLOG_CATEGORY_URL="http://api.csdn.net/blog/getcategorylist"
	//博客TAG
	BLOG_TAG_URL="http://api.csdn.net/blog/gettaglist"
	//获取文章内容
	BLOG_ID_URL="http://api.csdn.net/blog/getarticle"
)
