package model

type BlogStatistics struct {
	StatisticsId   int    `json:"statistics_id" xorm:"not null pk autoincr INT(11)"`
	BlogId         int    `json:"blog_id" xorm:"not null default 0 index INT(11)"`
	Comment        int    `json:"comment" xorm:"not null default 0 INT(11)"`
	Read           int    `json:"read" xorm:"not null default 0 INT(11)"`
	SeoTitle       string `json:"seo_title" xorm:"not null default '' VARCHAR(255)"`
	SeoDescription string `json:"seo_description" xorm:"not null default '' VARCHAR(255)"`
	SeoKeyword     string `json:"seo_keyword" xorm:"not null default '' VARCHAR(255)"`
}
