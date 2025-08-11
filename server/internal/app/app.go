package app

import (
	"github.com/gin-gonic/gin"
	"sheepo.com/emb_assessment/internal/config"
	"sheepo.com/emb_assessment/internal/handlers"
	"sheepo.com/emb_assessment/internal/routes"
)

type App struct {
	gin   *gin.Engine
	cfg   *config.Config
	cache map[string]any // Simple cache implementation
}

func NewApp(cfg *config.Config) *App {
	gin := gin.Default()

	cache := make(map[string]any)
	handler := handlers.NewHandler(cfg, cache)

	routes.SetupRoute(gin, handler)

	return &App{
		gin:   gin,
		cfg:   cfg,
		cache: cache,
	}
}

func (a *App) Run() {
	a.gin.Run(a.cfg.Addr)
}
