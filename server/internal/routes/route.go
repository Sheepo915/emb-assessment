package routes

import (
	"github.com/gin-gonic/gin"
	"sheepo.com/emb_assessment/internal/handlers"
)

func SetupRoute(g *gin.Engine, h *handlers.Handler) {
	{
		g.Static("/static", "internal/static")
	}

	{
		g.LoadHTMLGlob("internal/templates/*")

		g.GET("/", h.Index)
	}

	{
		v1 := g.Group("v1")

		v1.GET("/books")
		v1.POST("/book")
		v1.PUT("/book")
		v1.DELETE("/book")
		v1.POST("/book/reset")
	}
}
