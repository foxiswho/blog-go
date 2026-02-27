package tagsBasicEvent

import (
	"context"
	"github.com/foxiswho/blog-go/app/event/basic/model/modEventBasicTags"
	"github.com/foxiswho/blog-go/infrastructure/entityBasic"
	"github.com/foxiswho/blog-go/infrastructure/repositoryBasic"
	"github.com/foxiswho/blog-go/pkg/cachePg/rdsPg"
	"github.com/foxiswho/blog-go/pkg/enum/enumCommonPg/typeSysPg"
	"github.com/foxiswho/blog-go/pkg/enum/state/enumStatePg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/foxiswho/blog-go/pkg/tools/noPg"
	"github.com/pangu-2/go-tools/tools/strPg"
	"strings"
)

type Sp struct {
	log     *log2.Logger                                 `autowire:"?"`
	dao     *repositoryBasic.BasicAttachmentRepository   `autowire:"?"`
	TagRela *repositoryBasic.BasicTagsRelationRepository `autowire:"?"`
	TagCate *repositoryBasic.BasicTagsCategoryRepository `autowire:"?"`
	TagsDb  *repositoryBasic.BasicTagsRepository         `autowire:"?"`
	rdt     *rdsPg.BatchString                           `autowire:"?"`
}

// SaveByCategory
// @Description: 标签处理
type SaveByCategory struct {
	sp  *Sp
	log *log2.Logger `autowire:"?"`
	dto modEventBasicTags.TagsRelation
}

func NewSaveByCategory(sp *Sp, dto modEventBasicTags.TagsRelation) *SaveByCategory {
	return &SaveByCategory{
		sp:  sp,
		log: sp.log,
		dto: dto,
	}
}

// Processor 处理
//
//	@Description:
//	@receiver c
//	@param ctx
//	@return error
func (c *SaveByCategory) Processor() error {
	if strPg.IsBlank(c.dto.Category) {
		return nil
	}
	if nil == c.dto.Tags {
		return nil
	}
	if len(c.dto.Tags) < 1 {
		return nil
	}
	//类别是否存在
	info, result := c.sp.TagCate.FindByNo(c.dto.Category)
	if !result {
		return nil
	}
	categoryRoot := make([]string, 0)
	categoryRoot = append(categoryRoot, c.dto.Category)
	mapTagsOnly := make(map[string]string)
	tagsEntity := make(map[string]*entityBasic.BasicTagsEntity)
	tagsRelation := make(map[string]*entityBasic.BasicTagsRelationEntity)
	tagsNo := make([]string, 0)
	in, r := c.sp.TagsDb.FindAllByNameIn(c.dto.Tags)
	if r {
		for _, item := range in {
			mapTagsOnly[item.Name] = item.No
			tagsEntity[item.No] = item
			tagsNo = append(tagsNo, item.No)
		}
	}
	//如果不存在设置 false
	for _, item := range c.dto.Tags {
		if _, ok := mapTagsOnly[item]; !ok {
			mapTagsOnly[item] = ""
		}
	}
	//获取标签关系
	infos, b := c.sp.TagRela.FindAllByTagNoInAndCategoryRoot(tagsNo, c.dto.Category)
	if b {
		for _, item := range infos {
			tagsRelation[item.TagNo] = item
		}
	}
	for _, v := range c.dto.Tags {
		//标签是否存在
		if get, ok := mapTagsOnly[v]; ok {
			if strPg.IsNotBlank(get) {
				//关系存在，则跳过
				if _, ok2 := tagsRelation[get]; ok2 {
					continue
				}
			}
			save := &entityBasic.BasicTagsEntity{}
			//标签 不存在，创建标签
			if strPg.IsBlank(get) {
				save.Name = strings.TrimSpace(v)
				save.NameFull = save.Name
				save.State = enumStatePg.ENABLE.Index()
				save.TypeSys = typeSysPg.General.Index()
				save.No = noPg.No()
				save.Code = save.No
				save.CategoryNo = info.No
				save.TenantNo = c.dto.Holder.GetTenantNo()
				save.CreateBy = c.dto.Holder.GetAccountNo()
				c.sp.TagsDb.Create(save)
				//
				mapTagsOnly[v] = save.No
				tagsEntity[save.No] = save
			} else {
				save = tagsEntity[get]
			}
			//关系不存在
			if _, ok2 := tagsRelation[save.No]; !ok2 {
				save = tagsEntity[save.No]
				rel := entityBasic.BasicTagsRelationEntity{}
				rel.TagNo = save.No
				rel.CategoryNo = save.No
				rel.State = enumStatePg.ENABLE.Index()
				rel.Name = save.Name
				rel.NameFl = save.NameFl
				rel.NameFull = save.NameFull
				rel.Code = save.Code
				rel.TenantNo = save.TenantNo
				rel.TypeSys = save.TypeSys
				rel.Module = save.Module
				//rel.Attribute = save.Attribute
				rel.CategoryRoot = c.dto.Category
				c.sp.TagRela.Create(&rel)
				//
				tagsRelation[rel.TagNo] = &rel
			}
		}
	}
	c.cache(categoryRoot)
	return nil
}

// 缓存更新
func (t *SaveByCategory) cache(categoryRoot []string) {
	dto := modEventBasicTags.TagsCacheDto{CategoryRoot: categoryRoot}
	err := NewCachePush(t.sp, dto).Processor(context.Background())
	if err != nil {
		t.log.Error("tags.push.error:=", err)
	}
}
