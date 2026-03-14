package repositoryPg

import (
	"gorm.io/gorm"
)

type Condition func(*gorm.DB) *gorm.DB

func WithCondition(c Condition) Condition {
	return c
}
