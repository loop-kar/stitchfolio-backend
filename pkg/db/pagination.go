package db

import (
	"context"
	"math"

	"github.com/imkarthi24/sf-backend/pkg/constants"
	"github.com/imkarthi24/sf-backend/pkg/util"
	"gorm.io/gorm"
)

type Pagination struct {
	SkipPagination bool   `json:"skipPagination" form:"skipPagination"`
	Size           int    `json:"size,omitempty" form:"size"`
	CurrentPage    int    `json:"currentPage,omitempty" form:"currentPage"`
	Sort           string `json:"sort,omitempty" form:"sort"`
	TotalRecords   int64  `json:"totalRecords,omitempty"`
	TotalPages     int    `json:"totalPages,omitempty"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetSize()
}

func (p *Pagination) GetSize() int {
	if p.Size == 0 {
		p.Size = 10
	}
	return p.Size
}

func (p *Pagination) GetPage() int {
	if p.CurrentPage == 0 {
		p.CurrentPage = 1
	}
	return p.CurrentPage
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "1 DESC"
	}
	return p.Sort
}

func Paginate(ctx *context.Context) func(db *gorm.DB) *gorm.DB {

	skipPagination := util.ReadValueFromContext(ctx, constants.INTERNAL_SKIP_PAGINATION)

	if skipPagination != nil && skipPagination.(bool) {

		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}
	var pagination *Pagination
	var totalRows int64

	paginationReq := util.ReadValueFromContext(ctx, constants.PAGINATION_KEY)
	if paginationReq != nil {
		pagination = paginationReq.(*Pagination)
	} else {
		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}

	if pagination.SkipPagination {
		return func(db *gorm.DB) *gorm.DB {
			return db.Order(pagination.GetSort())
		}
	}

	return func(db *gorm.DB) *gorm.DB {

		tx := db.WithContext(context.Background())

		tx.Count(&totalRows)

		pagination.TotalRecords = totalRows
		totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.GetSize())))
		pagination.TotalPages = totalPages

		criteria := db.Offset(pagination.GetOffset()).Limit(pagination.GetSize()).Order(pagination.GetSort())

		return criteria
	}
}

func CustomPaginate(ctx *context.Context, pageSize int, rowCountMultiplier float64) func(db *gorm.DB) *gorm.DB {

	skipPagination := util.ReadValueFromContext(ctx, constants.INTERNAL_SKIP_PAGINATION)

	if skipPagination != nil && skipPagination.(bool) {

		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}

	var pagination *Pagination
	var totalRows int64

	paginationReq := util.ReadValueFromContext(ctx, constants.PAGINATION_KEY)
	if paginationReq != nil {
		pagination = paginationReq.(*Pagination)
	} else {

		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}
	if pagination.SkipPagination {
		return func(db *gorm.DB) *gorm.DB {
			return db.Order(pagination.GetSort())
		}
	}

	return func(db *gorm.DB) *gorm.DB {

		tx := db.WithContext(context.Background())

		tx.Count(&totalRows)

		pagination.TotalRecords = int64(float64(totalRows) * rowCountMultiplier)
		totalPages := int(math.Ceil(float64(pagination.TotalRecords) / float64(pagination.GetSize())))
		pagination.TotalPages = totalPages

		offset := float64(pagination.GetOffset()) / rowCountMultiplier
		criteria := db.Offset(int(offset)).Limit(pageSize).Order(pagination.GetSort())

		return criteria
	}
}
