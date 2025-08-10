package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"sheepo.com/emb_assessment/pkg/model"
)

type bookHandler interface {
	Index(c *gin.Context)
}

func (h *Handler) Index(c *gin.Context) {
	resp, err := http.Get(h.cfg.DummyAPI)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error":       "failed to query dummy api",
			"description": err.Error(),
		})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error":       "failed to query dummy api",
			"description": err.Error(),
		})
		return
	}

	var data model.Response
	if err = json.Unmarshal(body, &data); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error":       "failed to query dummy api",
			"description": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "index.html", data)
}
