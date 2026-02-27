package service

import (
	"context"
	"encoding/json"
	"reflect"
	"strings"

	"github.com/foxiswho/blog-go/app/manage/domainBasic/model/modBasicTagsRelation"
	"github.com/foxiswho/blog-go/app/manage/domainBlog/service/blogArticle"
	"github.com/foxiswho/blog-go/app/manage/domainBlog/utilsBlog"
	"github.com/foxiswho/blog-go/app/web/blog/model/modBlogArticle"
	"github.com/foxiswho/blog-go/app/web/utils/webPg"
	"github.com/foxiswho/blog-go/infrastructure/entityBlog"
	"github.com/foxiswho/blog-go/infrastructure/repositoryBasic"
	"github.com/foxiswho/blog-go/infrastructure/repositoryBlog"
	"github.com/foxiswho/blog-go/middleware/components/markdownPg"
	"github.com/foxiswho/blog-go/pkg/cachePg/rdsPg"
	"github.com/foxiswho/blog-go/pkg/consts/constTags"
	"github.com/foxiswho/blog-go/pkg/enum/blog/attachmentTypePg"
	"github.com/foxiswho/blog-go/pkg/enum/state/enumApprovedPg"
	"github.com/foxiswho/blog-go/pkg/enum/state/enumStatePg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/gin-gonic/gin"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
	"github.com/jinzhu/copier"
	"github.com/pangu-2/go-tools/tools/datetimePg"
	"github.com/pangu-2/go-tools/tools/dbPg/pagePg"
	"github.com/pangu-2/go-tools/tools/strPg"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
)

func init() {
	gs.Provide(new(ArticleService)).Init(func(s *ArticleService) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type ArticleService struct {
	sv        *repositoryBlog.BlogArticleRepository           `autowire:"?"`
	statDb    *repositoryBlog.BlogArticleStatisticsRepository `autowire:"?"`
	catDb     *repositoryBlog.BlogArticleCategoryRepository   `autowire:"?"`
	sp        *blogArticle.Sp                                 `autowire:"?"`
	log       *log2.Logger                                    `autowire:"?"`
	tagsRelat *repositoryBasic.BasicTagsRelationRepository    `autowire:"?"`
	rdu       *rdsPg.BatchString                              `autowire:"?"`
}

// Detail 详情
//
//	@Description:
//	@receiver c
//	@param id
func (c *ArticleService) Detail(ctx *gin.Context, id string) (rt rg.Rs[modBlogArticle.DetailVo]) {
	if strPg.IsBlank(id) {
		return rt.ErrorMessage("id错误")
	}
	id = strings.TrimSpace(id)
	find, b := c.sv.FindByIdString(id)
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	no := webPg.GetTenantNo(ctx)
	if strPg.IsNotBlank(no) && find.TenantNo != no {
		return rt.ErrorMessage("数据不存在")
	}
	var stat modBlogArticle.StatisticsVo
	byId, result := c.statDb.FindById(find.ID)
	if result {
		copier.Copy(&stat, &byId)
	}
	c.log.Infof("find=%+v", find)
	var info modBlogArticle.DetailVo
	copier.Copy(&info, &find)
	//
	tags := make([]string, 0)
	info.Tags = make([]string, 0)
	info.Where = make([]string, 0)
	info.AttachmentsMap = make(map[string]string)
	mapTagsOnly := make(map[string]bool)
	//附件图
	tagsData := make(map[string]modBasicTagsRelation.AllVo)
	if strPg.IsNotBlank(find.Attachments) {
		var imagesMap map[string]string
		err := json.Unmarshal([]byte(find.Attachments), &imagesMap)
		if err == nil {
			info.AttachmentsMap = imagesMap
		}
	}
	//标签
	if nil != find.Tags.Data() {
		tmp := find.Tags.Data()
		if len(tmp) > 0 {
			for _, tag := range tmp {
				//重复过滤掉
				if _, ok := mapTagsOnly[tag]; ok {
					continue
				}
				mapTagsOnly[tag] = false
				tags = append(tags, tag)
				//标签
				//info.Tags = append(info.Tags, tag)
			}
		}
		if len(tags) > 0 {
			{
				infos, result := c.tagsRelat.FindAllByTagNoInAndCategoryRoot(tags, constTags.ArticleInfo.String())
				if result {
					for _, item := range infos {
						var vo modBasicTagsRelation.AllVo
						copier.Copy(&vo, &item)
						vo.AttributeMap = make(map[string]interface{})
						//
						if strPg.IsNotBlank(item.Attribute) {
							err := json.Unmarshal([]byte(item.Attribute), &vo.AttributeMap)
							if err != nil {
								c.log.Errorf("json解析失败 %+v", err)
							}
							if obj, ok := vo.AttributeMap["color"]; ok {
								color := make(map[string]interface{})
								if strPg.IsNotBlank(obj.(string)) {
									err := json.Unmarshal([]byte(obj.(string)), &color)
									if err != nil {
										c.log.Errorf("json解析失败 %+v", err)
									}
								}
								vo.AttributeMap["color"] = color
							}
						} else {
							vo.AttributeMap["bordered"] = true
							vo.AttributeMap["type"] = "default"
							vo.AttributeMap["color"] = struct {
							}{}
							vo.AttributeMap["strong"] = false
							vo.AttributeMap["round"] = false
						}
						//
						tagsData[item.TagNo] = vo
						mapTagsOnly[item.TagNo] = true
					}
				}
			}
			{
				for tag, b2 := range mapTagsOnly {
					if !b2 {
						vo := modBasicTagsRelation.AllVo{
							Name:   tag,
							Code:   tag,
							NameFl: tag,
							TagNo:  tag,
						}
						vo.AttributeMap = make(map[string]interface{})
						vo.AttributeMap["bordered"] = true
						vo.AttributeMap["type"] = "default"
						vo.AttributeMap["color"] = struct {
						}{}
						vo.AttributeMap["strong"] = false
						vo.AttributeMap["round"] = false
						tagsData[tag] = vo
					}
				}
			}
			//
			for _, tag := range tags {
				obj := tagsData[tag]
				info.TagsStyle = append(info.TagsStyle, obj)
				info.Tags = append(info.Tags, obj.Name)
			}
		}
	}

	//统计
	if strPg.IsNotBlank(find.No) {
		stat, result2 := c.statDb.FindByArticleNo(find.No)
		if result2 {
			info.Statistics.Comment = stat.Comment
			info.Statistics.Read = stat.Read
			info.Statistics.SeoKeywords = stat.SeoKeywords
			info.Statistics.SeoDescription = stat.SeoDescription
			info.Statistics.PageTitle = stat.PageTitle
		}
	}
	if strPg.IsNotBlank(find.Content) {
		raw := markdownPg.Markdown([]byte(find.Content))
		info.ContentConv = raw.String()
	}
	syslog.Infof(context.Background(), syslog.TagAppDef, "info:%+v", info)
	syslog.Infof(context.Background(), syslog.TagAppDef, "info.create:%+v", datetimePg.Format(info.CreateAt, "2006"))
	return rt.OkData(info)
}

// Query 查询
//
//	@Description:
//	@receiver c
//	@param ct
func (c *ArticleService) Query(ctx *gin.Context, ct modBlogArticle.QueryCt) (rt rg.Rs[pagePg.PaginatorPg[modBlogArticle.Vo]]) {
	var query entityBlog.BlogArticleEntity
	copier.Copy(&query, &ct)
	no := webPg.GetTenantNo(ctx)
	if strPg.IsNotBlank(no) {
		query.TenantNo = no
	}
	//启用
	query.State = enumStatePg.ENABLE.Index()
	//审批通过
	query.PlatformApproved = enumApprovedPg.ApprovedStateApproved.Index()
	slice := make([]modBlogArticle.Vo, 0)
	rt.Data.Data = slice
	r := c.sv
	page, err := r.FindAllPageQuery(query, func(p *pagePg.PageCondition[*entityBlog.BlogArticleEntity]) {
		p.PageOption = func(c *pagePg.PaginatorPg[*entityBlog.BlogArticleEntity]) {
			c.PageNum = ct.PageNum
			c.PageSize = ct.PageSize
			if c.PageSize < 1 {
				c.PageSize = 20
			}
		}
		//自定义查询
		p.Condition = r.DbModel().Order("create_at desc")
		//自定义查询
		if "" != ct.Wd {
			p.Condition.Where("name like ?", "%"+ct.Wd+"%")
		}
		//标签
		if nil != ct.TagsQuery && len(ct.TagsQuery) > 0 {
			for _, tag := range ct.TagsQuery {
				//获取缓存，得到 标签编号
				get, b := c.rdu.Get(ctx, utilsBlog.TagCacheKey(tag))
				if b {
					p.Condition.Where("tags @> ?", get)
				} else {
					p.Condition.Where("tags @> ?", "[\""+tag+"\"]")
				}
			}
		}
	})
	if nil != err {
		return rt.Ok()
	}

	if page.Total > 0 && page.Data != nil && len(page.Data) > 0 {
		ImgDefault := "/assets/imgs/shop/product-1-1.jpg"
		pg := pagePg.NewPaginatorPg(func(c *pagePg.PaginatorPg[modBlogArticle.Vo]) {
			c.TotalPage = page.TotalPage
			c.Total = page.Total
			c.PageSize = page.PageSize
			c.PageNum = page.PageNum
		})
		mapCategory := make(map[string]*entityBlog.BlogArticleCategoryEntity)
		mapStat := make(map[string]*entityBlog.BlogArticleStatisticsEntity)
		idsCategory := make([]string, 0)
		idsNo := make([]string, 0)
		tags := make([]string, 0)
		mapTagsOnly := make(map[string]bool)
		//标签
		tagsData := make(map[string]modBasicTagsRelation.AllVo)
		for _, item := range page.Data {
			if strPg.IsNotBlank(item.CategoryNo) {
				idsCategory = append(idsCategory, item.CategoryNo)
			}
			if strPg.IsNotBlank(item.No) {
				idsNo = append(idsNo, item.No)
			}
			//标签
			if nil != item.Tags.Data() {
				if len(item.Tags.Data()) > 0 {
					for _, tag := range item.Tags.Data() {
						//重复过滤掉
						if _, ok := mapTagsOnly[tag]; ok {
							continue
						}
						mapTagsOnly[tag] = false
						tags = append(tags, tag)
						//标签
					}
				}
			}
		}
		//分类
		{
			if len(idsCategory) > 0 {
				tmp, result := c.catDb.FindAllByNoIn(idsCategory)
				if result {
					for _, item := range tmp {
						mapCategory[item.No] = item
					}
				}
			}
		}
		//统计
		{
			if len(idsNo) > 0 {
				tmp, result := c.statDb.FindAllByArticleNoIn(idsNo)
				if result {
					for _, item := range tmp {
						mapStat[item.ArticleNo] = item
					}
				}
			}
		}
		//标签
		{
			if len(tags) > 0 {
				infos, result := c.tagsRelat.FindAllByTagNoInAndCategoryRoot(tags, constTags.ArticleInfo.String())
				if result {
					for _, item := range infos {
						var vo modBasicTagsRelation.AllVo
						copier.Copy(&vo, &item)
						//
						vo.AttributeMap = make(map[string]interface{})
						//
						if strPg.IsNotBlank(item.Attribute) {
							err := json.Unmarshal([]byte(item.Attribute), &vo.AttributeMap)
							if err != nil {
								c.log.Errorf("json解析失败 %+v", err)
							}
							if obj, ok := vo.AttributeMap["color"]; ok {
								color := make(map[string]interface{})
								if strPg.IsNotBlank(obj.(string)) {
									err := json.Unmarshal([]byte(obj.(string)), &color)
									if err != nil {
										c.log.Errorf("json解析失败 %+v", err)
									}
								}
								vo.AttributeMap["color"] = color
							}
						} else {
							vo.AttributeMap["bordered"] = true
							vo.AttributeMap["type"] = "default"
							vo.AttributeMap["color"] = struct {
							}{}
							vo.AttributeMap["strong"] = false
							vo.AttributeMap["round"] = false
						}
						//
						tagsData[item.TagNo] = vo
						mapTagsOnly[item.TagNo] = true
					}
				}
				for tag, b := range mapTagsOnly {
					if !b {
						vo := modBasicTagsRelation.AllVo{
							Name:   tag,
							Code:   tag,
							NameFl: tag,
							TagNo:  tag,
						}
						vo.AttributeMap = make(map[string]interface{})
						vo.AttributeMap["bordered"] = true
						vo.AttributeMap["type"] = "default"
						vo.AttributeMap["color"] = struct {
						}{}
						vo.AttributeMap["strong"] = false
						vo.AttributeMap["round"] = false
						tagsData[tag] = vo
					}
				}
			}
		}
		//字段赋值
		for _, item := range page.Data {
			var vo modBlogArticle.Vo
			copier.Copy(&vo, &item)
			vo.Tags = make([]string, 0)
			vo.TagsStyle = make([]modBasicTagsRelation.AllVo, 0)
			//
			if strPg.IsNotBlank(item.Attachments) {
				json.Unmarshal([]byte(item.Attachments), &vo.AttachmentsMap)
			}
			if nil == vo.AttachmentsMap {
				vo.AttachmentsMap = make(map[string]string, 0)
			}
			if _, ok := vo.AttachmentsMap[attachmentTypePg.AttachmentList.String()]; !ok {
				vo.AttachmentsMap[attachmentTypePg.AttachmentList.String()] = ImgDefault
			}
			if _, ok := vo.AttachmentsMap[attachmentTypePg.AttachmentMain.String()]; !ok {
				vo.AttachmentsMap[attachmentTypePg.AttachmentMain.String()] = ImgDefault
			}
			//统计
			if obj, ok := mapStat[item.No]; ok {
				vo.Statistics.Comment = obj.Comment
				vo.Statistics.Read = obj.Read
				vo.Statistics.SeoKeywords = obj.SeoKeywords
				vo.Statistics.SeoDescription = obj.SeoDescription
				vo.Statistics.PageTitle = obj.PageTitle
			}
			//分类
			if obj, ok := mapCategory[item.CategoryNo]; ok {
				vo.CategoryName = obj.Name
				vo.CategoryNo = obj.No
			}
			//标签
			if nil != item.Tags.Data() {
				if len(item.Tags.Data()) > 0 {
					for _, tag := range item.Tags.Data() {
						obj := tagsData[tag]
						vo.Tags = append(vo.Tags, obj.Name)
						vo.TagsStyle = append(vo.TagsStyle, obj)
					}
				}
			}
			//
			slice = append(slice, vo)
		}
		pg.Data = slice
		pg.Pageable = page.Pageable
		rt.Data = pg
		return rt.Ok()
	}
	return rt.Ok()
}
