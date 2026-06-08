package withDbPg

import (
	"gorm.io/gorm"
)

type ConditionOption func(*gorm.DB) *gorm.DB

func Condition(c ConditionOption) ConditionOption {
	return c
}
