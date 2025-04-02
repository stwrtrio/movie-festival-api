package models

type PaginationRequest struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"page_size" form:"page_size"`
}

type PaginationResponse struct {
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalItems int         `json:"total_items"`
	TotalPages int         `json:"total_pages"`
	Data       interface{} `json:"data"`
}
