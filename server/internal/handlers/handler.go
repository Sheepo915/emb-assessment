package handlers

import "sheepo.com/emb_assessment/internal/config"

type Handler struct {
	cfg *config.Config
	bookHandler
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{
		cfg: cfg,
	}
}
