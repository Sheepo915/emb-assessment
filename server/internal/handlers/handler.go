package handlers

import (
	"fmt"

	"sheepo.com/emb_assessment/internal/config"
)

type Handler struct {
	cfg   *config.Config
	cache map[string]any
	siteHandler
	bookHandler
}

const (
	Quantity = "1000"
	Seed     = "12345"
)

var DefaultBookKey = fmt.Sprintf("book:%s:%s", Seed, Quantity)

func NewHandler(cfg *config.Config, cache map[string]any) *Handler {
	return &Handler{
		cfg:   cfg,
		cache: cache,
	}
}

func getBookKey(seed, quantity int) string {
	return fmt.Sprintf("book:%d:%d", seed, quantity)
}
