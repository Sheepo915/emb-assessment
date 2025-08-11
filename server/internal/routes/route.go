package routes

import (
	"text/template"

	"github.com/gin-gonic/gin"
	"sheepo.com/emb_assessment/internal/handlers"
	"sheepo.com/emb_assessment/pkg/utils"
)

func SetupRoute(g *gin.Engine, h *handlers.Handler) {
	{
		g.Static("/static", "internal/static")
	}

	{
		g.SetFuncMap(template.FuncMap{
			"seq": utils.Seq,
			"sub": utils.Sub,
			"add": utils.Add,
		})
		g.LoadHTMLGlob("internal/templates/*")

		g.GET("/", h.Index)
	}

	{
		v1 := g.Group("v1")

		v1.GET("/api/books", h.GetBooks)
	}
}
