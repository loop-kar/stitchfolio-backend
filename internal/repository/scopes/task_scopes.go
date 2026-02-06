package scopes

import (
	"fmt"

	"github.com/loop-kar/pixie/util"
	"gorm.io/gorm"
)

func GetTasks_Search(search string) func(db *gorm.DB) *gorm.DB {
	defaultReturn := func(db *gorm.DB) *gorm.DB { return db }

	if util.IsNilOrEmptyString(&search) {
		return defaultReturn
	}

	return func(db *gorm.DB) *gorm.DB {
		formattedSearch := util.EncloseWithPercentageOperator(search)
		whereClause := fmt.Sprintf(
			`("stich"."Tasks".title ILIKE %s OR "stich"."Tasks".description ILIKE %s)`,
			formattedSearch, formattedSearch,
		)
		return db.Where(whereClause)
	}
}

func GetTasks_Filter(filters string) func(db *gorm.DB) *gorm.DB {
	defaultReturn := func(db *gorm.DB) *gorm.DB { return db }

	if util.IsNilOrEmptyString(&filters) {
		return defaultReturn
	}

	fieldMap := map[string]string{
		"IsCompleted": "is_completed",
		"Priority":    "priority",
		"DueDate":     "due_date",
		"CompletedAt": "completed_at",
	}

	return func(db *gorm.DB) *gorm.DB {
		queryString := util.BuildQuery(filters, fieldMap)
		if !util.IsNilOrEmptyString(&queryString) {
			return db.Where(queryString)
		}
		return db
	}
}
