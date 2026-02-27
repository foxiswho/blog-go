package service

import (
	"context"
	"reflect"
	"strings"
	"time"

	"github.com/foxiswho/blog-go/app/manage/domainBlog/service/blogCollect"
	"github.com/foxiswho/blog-go/app/web/api/model/modBlogCollect"
	"github.com/foxiswho/blog-go/infrastructure/entityBlog"
	"github.com/foxiswho/blog-go/infrastructure/repositoryBlog"
	"github.com/foxiswho/blog-go/pkg/enum/blog/typeContentPg"
	"github.com/foxiswho/blog-go/pkg/enum/blog/typeDataSourcePg"
	"github.com/foxiswho/blog-go/pkg/enum/blog/typeReadingPg"
	"github.com/foxiswho/blog-go/pkg/enum/blog/typeSourcePg"
	"github.com/foxiswho/blog-go/pkg/holderPg/holderApiPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/foxiswho/blog-go/pkg/tools/noPg"
	"github.com/gin-gonic/gin"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
	"github.com/pangu-2/go-tools/tools/cryptPg"
	"github.com/pangu-2/go-tools/tools/strPg"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
	"gorm.io/datatypes"
)

func init() {
	gs.Provide(new(CollectService)).Init(func(s *CollectService) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type CollectService struct {
	sv    *repositoryBlog.BlogCollectRepository         `autowire:"?"`
	catDb *repositoryBlog.BlogCollectCategoryRepository `autowire:"?"`
	sp    *blogCollect.Sp                               `autowire:"?"`
	log   *log2.Logger                                  `autowire:"?"`
}

// Push
//
//	@Description: 推送文章连接
//	@receiver c
func (c *CollectService) Push(ctx *gin.Context, ct modBlogCollect.PushCt) (rt rg.Rs[string]) {
	c.log.Infof("ct=%+v", ct)

	if strPg.IsBlank(ct.Title) {
		return rt.ErrorMessage("标题不能为空")
	}
	if strPg.IsBlank(ct.Url) {
		return rt.ErrorMessage("url地址不能为空")
	}
	urlSource := strings.TrimSpace(ct.Url)
	urlSourceMd5 := cryptPg.Md5(urlSource)
	//规则
	//重复检查
	duplicate := true
	if ct.Rule != nil && len(ct.Rule) > 0 {
		for _, v := range ct.Rule {
			if strPg.IsBlank(v) {
				continue
			}
			//重复:允许重复数据
			if "duplicatePass" == strings.TrimSpace(v) {
				//不检查重复数据
				duplicate = false
			}
		}
	}
	//重复
	if duplicate {
		info, result := c.sv.FindAllByUrlSourceMd5(urlSourceMd5)
		if result && nil != info && len(info) > 0 {
			return rt.ErrorMessage("该链接已存在")
		}
	}

	holder := holderApiPg.GetContextAccount(ctx)
	save := entityBlog.BlogCollectEntity{
		Name:        strings.TrimSpace(ct.Title),
		UrlSource:   urlSource,
		Description: strings.TrimSpace(ct.Description),
	}

	if strPg.IsBlank(ct.CategoryNo) {
		save.CategoryNo = "default"
	} else {
		info, result := c.catDb.FindByNo(ct.CategoryNo)
		if !result {
			return rt.ErrorMessage("分类不存在")
		}
		save.CategoryNo = info.No
	}
	save.Content = strings.TrimSpace(ct.Content)
	save.Editor = strings.TrimSpace(ct.Editor)
	save.Author = strings.TrimSpace(ct.Author)
	save.Source = strings.TrimSpace(ct.Source)
	tags := make([]string, 0)
	if nil != ct.Tags && len(ct.Tags) > 0 {
		for _, v := range ct.Tags {
			if strPg.IsNotBlank(v) {
				tags = append(tags, strings.TrimSpace(v))
			}
		}
	}
	save.Tags = datatypes.NewJSONType(tags)

	save.No = noPg.No()
	save.Code = save.No
	save.TenantNo = holder.GetTenantNo()
	save.UrlSourceMd5 = urlSourceMd5
	save.TypeReading = typeReadingPg.UNREAD.Index()
	save.TypeDataSource = typeDataSourcePg.EXTERNAL.Index()
	save.TypeSource = typeSourcePg.COLLECTION.Index()
	save.TypeContent = typeContentPg.REPOST.Index()
	now := time.Now()
	save.OperationTime = &now
	if nil == ct.PublishTime {
		save.PublishTime = &now
	} else {
		toTime := ct.PublishTime.ToTime()
		save.PublishTime = &toTime
	}
	//
	//
	err, _ := c.sv.Create(&save)
	if err != nil {
		c.log.Debugf("save err=%+v", err)
		return rt.ErrorMessage("保存失败：" + err.Error())
	}
	return rt.Ok()
}
