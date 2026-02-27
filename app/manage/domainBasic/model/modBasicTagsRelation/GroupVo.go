package modBasicTagsRelation

import "github.com/foxiswho/blog-go/app/manage/domainBasic/model/modBasicTagsCategory"

type GroupVo struct {
	General []modBasicTagsCategory.Vo `json:"general" label:"普通"`
	Sys     []modBasicTagsCategory.Vo `json:"sys" label:"系统"`
}
