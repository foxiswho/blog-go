package modEventBasicTags

import "github.com/foxiswho/blog-go/pkg/holderPg"

type TagsRelation struct {
	Category string            `json:"category"`
	Tags     []string          `json:"tags"`
	Holder   holderPg.HolderPg `json:"holder"`
}
