package handlers

import (
	"encoding/json"
	"io"
	"math"
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
	"sheepo.com/emb_assessment/pkg/dto"
	"sheepo.com/emb_assessment/pkg/model"
	"sheepo.com/emb_assessment/pkg/utils"
)

type bookHandler interface {
	Index(c *gin.Context)
}

const (
	quantity = "1000"
	seed     = "12345"
)

func (h *Handler) Index(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limitParam := c.DefaultQuery("limit", "10")
	order := c.DefaultQuery("sort", "asc")

	var limit int
	if limitParam == "unlimited" {
		limit = -1
	} else {
		limit, _ = strconv.Atoi(limitParam)
	}

	totalItems := 0
	if limit == -1 {
		totalItems, _ = strconv.Atoi(quantity) // full dataset size
	} else {
		totalItems, _ = strconv.Atoi(quantity) // still full dataset for pagination calc
	}

	var data model.Response
	if _, ok := h.cache["books"]; !ok {
		api, err := utils.APIConstruct(h.cfg.DummyAPI,
			&utils.QueryKVPair{
				Key: "_quantity", Value: strconv.Itoa(totalItems),
			}, &utils.QueryKVPair{
				Key: "_seed", Value: seed,
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

		h.cache["books"] = data
	} else {
		data, _ = h.cache["books"].(model.Response)
	}

	if order == "desc" {
		sort.Slice(data.Data, func(i, j int) bool {
			return data.Data[i].Title > data.Data[j].Title
		})
	} else {
		sort.Slice(data.Data, func(i, j int) bool {
			return data.Data[i].Title < data.Data[j].Title
		})
	}

	totalPages := 1
	if limit > 0 {
		totalPages = int(math.Ceil(float64(len(data.Data)) / float64(limit)))
	}

	prev := 0
	if page > 1 {
		prev = page - 1
	}

	next := 0
	if page < totalPages {
		next = page + 1
	}

	if limit != -1 {
		start := (page - 1) * limit
		end := start + limit
		if start > len(data.Data) {
			start = len(data.Data)
		}
		if end > len(data.Data) {
			end = len(data.Data)
		}
		data.Data = data.Data[start:end]
	}

	// Replace given placeholder with working placeholder
	for i := range data.Data {
		data.Data[i].Image = "https://placehold.co/480x640"
	}

	var selected model.Book
	idParam := c.Query("id")
	if idParam != "" {
		if id, err := strconv.Atoi(idParam); err == nil {
			for _, book := range data.Data {
				if book.ID == id {
					selected = book
					break
				}
			}
		}
	}

	renderData := dto.RenderBookHome{
		Books:    data.Data,
		Selected: selected,
		Pagination: dto.Pagination{
			Page:       page,
			Prev:       prev,
			Next:       next,
			TotalPages: totalPages,
			Limit:      limit,
			Order:      order,
		},
	}

	c.HTML(http.StatusOK, "index.html", renderData)
}
