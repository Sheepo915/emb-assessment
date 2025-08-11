package handlers

import "sheepo.com/emb_assessment/internal/config"

type Handler struct {
	cfg   *config.Config
	cache map[string]any
	bookHandler
}

func NewHandler(cfg *config.Config, cache map[string]any) *Handler {
	return &Handler{
		cfg:   cfg,
		cache: cache,
	}
}
