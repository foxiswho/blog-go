package data

import (
	"context"

	"github.com/foxiswho/blog-go/infrastructure/entityBasic"
	"github.com/foxiswho/blog-go/infrastructure/repositoryBasic"
	"github.com/foxiswho/blog-go/pkg/consts/constNodePg"
	"github.com/foxiswho/blog-go/pkg/consts/constTags"
	"github.com/foxiswho/blog-go/pkg/consts/constsPg"
	"github.com/foxiswho/blog-go/pkg/enum/enumCommonPg/typeSysPg"
	"github.com/foxiswho/blog-go/pkg/enum/state/enumStatePg"
	"github.com/foxiswho/blog-go/pkg/log2"
	syslog "github.com/go-spring/log"
	"github.com/pangu-2/go-tools/tools/numberPg"
	"gorm.io/datatypes"
)

// IBasicData
// @Description: 初始化基础数据
type IBasicData struct {
	log     *log2.Logger                                 `autowire:"?"`
	country *repositoryBasic.BasicCountryRepository      `autowire:"?"`
	tagsCat *repositoryBasic.BasicTagsCategoryRepository `autowire:"?"`
}

func (b *IBasicData) Run() error {
	syslog.Infof(context.Background(), syslog.TagAppDef, "[init].[基础数据：国家:中国]===================")
	{
		save := entityBasic.BasicCountryEntity{
			ID:           100000,
			No:           "100000",
			Code:         "86",
			Name:         "中国",
			NameFl:       "CHN",
			NameFull:     "中国",
			State:        enumStatePg.ENABLE.Index(),
			Iso3:         "CHN",
			CountryCode:  "86",
			PhoneUse:     enumStatePg.ENABLE.Index(),
			DomainSuffix: "cn",
		}
		//不存在时创建
		_, result := b.country.FindByNo(save.No)
		if !result {
			save.IdLink = constNodePg.NoLinkDefault(numberPg.Int64ToString(save.ID))
			save.NoLink = constNodePg.NoLinkDefault(save.No)
			b.country.Create(&save)
		}
	}
	//
	syslog.Infof(context.Background(), syslog.TagAppDef, "[init].[基础数据：标签:分类]===================")
	{
		save := entityBasic.BasicTagsCategoryEntity{
			ID:       100000,
			No:       constTags.ArticleInfo.Index(),
			Code:     constTags.ArticleInfo.Index(),
			Name:     "文章",
			NameFl:   "文章",
			NameFull: "文章",
			State:    enumStatePg.ENABLE.Index(),
			TenantNo: constsPg.ACCOUNT_MANAGE_No,
			TypeSys:  typeSysPg.System.Index(),
			Tags:     datatypes.NewJSONType(make([]string, 0)),
		}
		//不存在时创建
		_, result := b.tagsCat.FindByNo(save.No)
		if !result {
			save.IdLink = constNodePg.NoLinkDefault(numberPg.Int64ToString(save.ID))
			save.NoLink = constNodePg.NoLinkDefault(save.No)
			b.tagsCat.Create(&save)
		}
	}
	return nil
}
