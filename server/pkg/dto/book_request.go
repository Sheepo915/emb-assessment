package dto

import "sheepo.com/emb_assessment/pkg/model"

type GetBooksRequest struct {
	Seed    int  `form:"seed,default=12345"`
	Total   int  `form:"total,default=10"`
	Cache   bool `form:"cache,default=false"`
	Page    int  `form:"page,default=1"`
	PerPage int  `form:"per_page,default=10"`
}

type GetBooksResponse struct {
	Books      []model.Book `json:"books"`
	Page       int          `json:"page"`
	PerPage    int          `json:"per_page"`
	Total      int          `json:"total"`
	TotalPages int          `json:"total_pages"`
}
