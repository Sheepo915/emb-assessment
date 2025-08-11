package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"sheepo.com/emb_assessment/pkg/dto"
	"sheepo.com/emb_assessment/pkg/model"
	"sheepo.com/emb_assessment/pkg/utils"
)

type bookHandler interface {
	GetBooks(c *gin.Context)
}

func (h *Handler) GetBooks(c *gin.Context) {
	var query dto.GetBooksRequest

	err := c.ShouldBind(&query)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:       "invalid request",
			Description: err.Error(),
		})
		return
	}

	if query.Total < query.PerPage {
		query.Total = query.PerPage
	}

	bookKey := getBookKey(query.Seed, query.Total)

	var data model.Response
	if _, ok := h.cache[bookKey]; !ok {
		api, err := utils.APIConstruct(h.cfg.DummyAPI,
			&utils.QueryKVPair{
				Key: "_quantity", Value: strconv.Itoa(query.Total),
			}, &utils.QueryKVPair{
				Key: "_seed", Value: strconv.Itoa(query.Seed),
			},
		)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error":       "failed to query dummy api",
				"description": err.Error(),
			})
			return
		}

		resp, err := http.Get(api)
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

		if err = json.Unmarshal(body, &data); err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error":       "failed to query dummy api",
				"description": err.Error(),
			})
			return
		}

		if query.Cache {
			h.cache[bookKey] = data
		}
	} else {
		data, _ = h.cache[bookKey].(model.Response)
	}

	totalPages := int(math.Ceil(float64(len(data.Data)) / float64(query.PerPage)))

	start := (query.Page - 1) * query.PerPage
	end := start + query.PerPage
	if start > len(data.Data) {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:       "page overflow",
			Description: fmt.Sprintf("You have exceeded the page limit. Total page at %d", totalPages),
		})
		return
	}
	if end > len(data.Data) {
		end = len(data.Data)
	}

	c.JSON(http.StatusOK, dto.GetBooksResponse{
		Books:      data.Data[start:end],
		Page:       query.Page,
		PerPage:    query.PerPage,
		Total:      len(data.Data),
		TotalPages: totalPages,
	})
}
