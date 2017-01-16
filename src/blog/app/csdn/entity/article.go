package entity


//文章
type Article struct {
	Id                int `json:"id" 否 文章id，修改文章的时候需要`
	Title             string `json:"title" 是 文章标题`
	CreateAt          string `json:"create_at"  发表时间`
	Create            string `json:"create"  发表时间`
	ViewCount         int `json:"view_count"  阅读次数`
	CommentCount      int `json:"comment_count"  评论次数`
	CommentAllowed    string `json:"comment_allowed"  是否允许评论`
	Type              string `json:"type" 是 文章类型（original|report|translated）`
	ChannelId         string `json:"channel_id" 否 系统类别id`
	Channel           int `json:"channel" 否 系统类别id`
	Digg              int `json:"digg" 否 被顶次数`
	Bury              int `json:"bury" 否 被踩次数`
	Description       string `json:"description" 否 文章简介`
	Content           string `json:"content" 是 文章内容`
	MarkdownContent   string `json:"markdowncontent" 是 文章内容`
	markdownDirectory string `json:"-"`
	Categories        string `json:"categories" 否 自定义类别（英文逗号分割）`
	Tags              string `json:"tags" 否 文章标签（英文逗号分割）`
	Status            int `json:"status"`
	ArticleEditType   int `json:"articleedittype"`
	Level             int `json:"level"`
	ArticleMore       string `json:"articlemore"`
}