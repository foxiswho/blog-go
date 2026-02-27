package service

import (
	"context"
	"reflect"

	"github.com/foxiswho/blog-go/app/manage/domainBlog/model/modBlogTopicRelation"
	"github.com/foxiswho/blog-go/infrastructure/entityBlog"
	"github.com/foxiswho/blog-go/infrastructure/repositoryBlog"
	"github.com/foxiswho/blog-go/pkg/enum/state/enumStatePg"
	"github.com/foxiswho/blog-go/pkg/holderPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/gin-gonic/gin"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
	"github.com/jinzhu/copier"
	"github.com/pangu-2/go-tools/tools/dbPg/pagePg"
	"github.com/pangu-2/go-tools/tools/numberPg"
	"github.com/pangu-2/go-tools/tools/strPg"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
)

func init() {
	gs.Provide(new(BlogTopicRelationService)).Init(func(s *BlogTopicRelationService) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type BlogTopicRelationService struct {
	topic    *repositoryBlog.BlogTopicRepository         `autowire:"?"`
	article  *repositoryBlog.BlogArticleRepository       `autowire:"?"`
	relation *repositoryBlog.BlogTopicRelationRepository `autowire:"?"`
	log      *log2.Logger                                `autowire:"?"`
}

// AddByTopic
//
//	@Description: 加入话题
//	@receiver c
//	@param ctx
//	@param ct
//	@return rt
func (c *BlogTopicRelationService) AddByTopic(ctx *gin.Context, ct modBlogTopicRelation.AddByTopicCt) (rt rg.Rs[string]) {
	c.log.Infof("ct=%+v", ct)
	if strPg.IsBlank(ct.TopicNo) {
		return rt.ErrorMessage("话题编号不能为空")
	}
	if nil == ct.Nos || 0 == len(ct.Nos) {
		return rt.ErrorMessage("文章编号不能为空")
	}
	topic, result := c.topic.FindByNo(ct.TopicNo, repositoryPg.GetOption(ctx))
	if !result {
		return rt.ErrorMessage("话题 不存在")
	}
	if !enumStatePg.ENABLE.IsEqualInt8(topic.State) {
		return rt.ErrorMessage("话题 状态异常")
	}
	ids := make([]string, 0)
	for _, no := range ct.Nos {
		if strPg.IsNotBlank(no) {
			ids = append(ids, no)
		}
	}
	if len(ids) < 1 {
		return rt.ErrorMessage("文章编号不能为空")
	}
	articles, result := c.article.FindAllByNoIn(ids)
	if !result {
		return rt.ErrorMessage("文章 不存在")
	}
	holder := holderPg.GetContextAccount(ctx)
	tenantNo := holder.GetTenantNo()
	ano := holder.GetAccountNo()
	//mapArticles := make(map[string]*entityBlog.BlogArticleEntity)
	for _, article := range articles {
		//判断是否存在，如果存在则跳过
		_, r := c.relation.FindByTopicNoAndArticleNo(topic.No, article.No)
		if r {
			continue
		}
		find, b := c.article.FindByNo(article.No, repositoryPg.GetOption(ctx))
		if !b {
			continue
		}
		obj := entityBlog.BlogTopicRelationEntity{}
		obj.TenantNo = tenantNo
		obj.TopicNo = topic.No
		obj.ArticleNo = article.No
		obj.Ano = ano
		obj.Name = find.Name
		obj.Description = find.Description
		err, _ := c.relation.Create(&obj)
		if nil != err {
			c.log.Errorf("save err=%+v", err)
		}
	}
	return rt.Ok()
}

// PhysicalDeletion 物理删除
//
//	@Description:
//	@receiver c
//	@param ct
func (c *BlogTopicRelationService) PhysicalDeletion(ctx *gin.Context, ids []string) (rt rg.Rs[string]) {
	c.log.Infof("ct=%+v", ids)
	if len(ids) < 1 {
		return rt.ErrorMessage("id错误")
	}
	cn := c.relation
	finds, b := cn.FindAllByIdStringIn(ids, repositoryPg.GetOption(ctx))
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	holder := holderPg.GetContextAccount(ctx)
	tenantNo := holder.GetTenantNo()
	idsNew := make([]string, 0)
	for _, info := range finds {
		c.log.Infof("id=%v,TenantNo=%v", info.ID, info.TenantNo)
		idsNew = append(idsNew, numberPg.Int64ToString(info.ID))
	}
	if len(idsNew) > 0 {
		err := cn.DeleteAllByTenantNoAndIdsString(tenantNo, idsNew)
		if err != nil {
			c.log.Errorf("操作 err=%+v", err)
		}
	}
	return rt.Ok()
}

// Query 查询
//
//	@Description:
//	@receiver c
//	@param ct
func (c *BlogTopicRelationService) Query(ctx *gin.Context, ct modBlogTopicRelation.QueryCt) (rt rg.Rs[pagePg.PaginatorPg[modBlogTopicRelation.Vo]]) {
	c.log.Infof("ct=%+v", ct)
	var query entityBlog.BlogTopicRelationEntity
	copier.Copy(&query, &ct)
	slice := make([]modBlogTopicRelation.Vo, 0)
	rt.Data.Data = slice
	r := c.relation
	page, err := r.FindAllPageQuery(query, func(p *pagePg.PageCondition[*entityBlog.BlogTopicRelationEntity]) {
		p.PageOption = func(c *pagePg.PaginatorPg[*entityBlog.BlogTopicRelationEntity]) {
			c.PageNum = ct.PageNum
			c.PageSize = ct.PageSize
			if c.PageSize < 1 {
				c.PageSize = 20
			}
		}
		//自定义查询
		p.Condition = r.DbModel().Order("sort asc").Order("create_at desc")
		//自定义查询
		if "" != ct.Wd {
			p.Condition.Where("name like ?", "%"+ct.Wd+"%")
		}
	}, repositoryPg.GetOption(ctx))
	if nil != err {
		return rt.Ok()
	}

	if page.Total > 0 && page.Data != nil && len(page.Data) > 0 {
		pg := pagePg.NewPaginatorPg(func(c *pagePg.PaginatorPg[modBlogTopicRelation.Vo]) {
			c.TotalPage = page.TotalPage
			c.Total = page.Total
			c.PageSize = page.PageSize
			c.PageNum = page.PageNum
		})
		//字段赋值
		for _, item := range page.Data {
			var vo modBlogTopicRelation.Vo
			copier.Copy(&vo, &item)
			slice = append(slice, vo)
		}
		pg.Data = slice
		pg.Pageable = page.Pageable
		rt.Data = pg
		return rt.Ok()
	}
	return rt.Ok()
}
