package utils

import (
	"math"

	"gorm.io/gorm"
)

type Meta struct {
	Page       int  `json:"page"`
	Limit      int  `json:"limit"`
	TotalItems int  `json:"total_items"`
	TotalPages int  `json:"total_pages"`
	HasPrev    bool `json:"has_prev"`
	HasNext    bool `json:"has_next"`
}

type Pagination struct {
	Data interface{} `json:"data"`
	Meta Meta        `json:"meta"`
}

func Paginate(page, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}

		if limit <= 0 {
			limit = 10
		}

		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}

func NewPagination(page, limit, totalItems int, data interface{}) *Pagination {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10 // default limit
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))

	return &Pagination{
		Data: data,
		Meta: Meta{
			Page:       page,
			Limit:      limit,
			TotalItems: totalItems,
			TotalPages: totalPages,
			HasPrev:    page > 1,
			HasNext:    page < totalPages,
		},
	}
}
