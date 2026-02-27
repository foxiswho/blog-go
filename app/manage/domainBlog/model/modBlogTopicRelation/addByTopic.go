package modBlogTopicRelation

type AddByTopicCt struct {
	TopicNo string   `json:"topicNo" label:"话题编号" `
	Nos     []string `json:"nos" label:"文章编号" `
}
