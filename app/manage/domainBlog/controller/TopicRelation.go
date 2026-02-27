package controller

import (
	"github.com/foxiswho/blog-go/app/manage/domainBlog/model/modBlogTopicRelation"
	"github.com/foxiswho/blog-go/app/manage/domainBlog/service"
	"github.com/foxiswho/blog-go/middleware/authPg"
	"github.com/foxiswho/blog-go/middleware/validatorPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/gin-gonic/gin"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"

	"github.com/foxiswho/blog-go/pkg/model"
)

func init() {

}

// TopicRelationController 文章话题关系
// @Description:
type TopicRelationController struct {
	Sp  *authPg.GroupManageMiddlewareSp   `autowire:""`
	sv  *service.BlogTopicRelationService `autowire:"?"`
	log *log2.Logger                      `autowire:"?"`
}

// AddByTopic 创建
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *TopicRelationController) AddByTopic(ctx *gin.Context) {
	var ct modBlogTopicRelation.AddByTopicCt
	if err := ctx.ShouldBind(&ct); err != nil {
		//对 返回 错误进行转义 成中文
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	ctx.JSON(200, c.sv.AddByTopic(ctx, ct))
}

// PhysicalDeletion 物理删除
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *TopicRelationController) PhysicalDeletion(ctx *gin.Context) {
	var ct model.BaseIdsCt[string]
	if err := ctx.ShouldBind(&ct); err != nil {
		//对 返回 错误进行转义 成中文
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	ctx.JSON(200, c.sv.PhysicalDeletion(ctx, ct.Ids))
}

// Query 查询列表
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *TopicRelationController) Query(ctx *gin.Context) {
	var ct modBlogTopicRelation.QueryCt
	if err := ctx.ShouldBind(&ct); err != nil {
		//对 返回 错误进行转义 成中文
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	ctx.JSON(200, c.sv.Query(ctx, ct))
}
