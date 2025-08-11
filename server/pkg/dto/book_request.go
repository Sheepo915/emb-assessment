package dto

import "sheepo.com/emb_assessment/pkg/model"

type Pagination struct {
	Page       int
	Prev       int
	Next       int
	TotalPages int
	Limit      int
	Order      string
}

type RenderBookHome struct {
	Pagination Pagination
	Books      []model.Book
	Selected   model.Book
}
