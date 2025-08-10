package app

import (
	"github.com/gin-gonic/gin"
	"sheepo.com/emb_assessment/internal/config"
	"sheepo.com/emb_assessment/internal/handlers"
	"sheepo.com/emb_assessment/internal/routes"
)

type App struct {
	gin *gin.Engine
	cfg *config.Config
}

func NewApp(cfg *config.Config) *App {
	gin := gin.Default()

	handler := handlers.NewHandler(cfg)

	routes.SetupRoute(gin, handler)

	return &App{
		gin: gin,
		cfg: cfg,
	}
}

func (a *App) Run() {
	a.gin.Run(a.cfg.Addr)
}
