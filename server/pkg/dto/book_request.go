package dto

import "sheepo.com/emb_assessment/pkg/model"

type GetAllBooksResponse struct {
}

type RenderBookHome struct {
	Books []model.Book
}
