package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pancudaniel7/go-restful-api-example/internal/model/dto"
	services "github.com/pancudaniel7/go-restful-api-example/internal/service"
	"net/http"
	"strconv"
)

type StoreController struct {
	service services.StoreService
}

func NewStoreController(service services.StoreService) *StoreController {
	return &StoreController{service: service}
}

func (c *StoreController) AddStore(ctx *gin.Context) {
	var storeDTO dto.StoreDTO
	if err := ctx.ShouldBindJSON(&storeDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	store, err := c.service.AddStore(storeDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	storeResponse := gin.H{
		"store": store,
		"_links": gin.H{
			"self":   fmt.Sprintf("/stores/%d", store.ID),
			"update": fmt.Sprintf("/stores/%d", store.ID),
			"delete": fmt.Sprintf("/stores/%d", store.ID),
			"create": "/stores",
			"get":    fmt.Sprintf("/stores/%d", store.ID),
		},
	}

	ctx.JSON(http.StatusOK, storeResponse)
}

func (c *StoreController) DeleteStore(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store ID"})
		return
	}

	err = c.service.DeleteStore(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{
		"message": "Store deleted",
		"_links": gin.H{
			"self":   fmt.Sprintf("/stores/%d", id),
			"create": "/stores",
			"getAll": "/stores",
		},
	})
}

func (c *StoreController) UpdateStore(ctx *gin.Context) {
	var storeDTO dto.StoreDTO
	if err := ctx.ShouldBindJSON(&storeDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	store, err := c.service.UpdateStore(storeDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	storeResponse := gin.H{
		"store": store,
		"_links": gin.H{
			"self":   fmt.Sprintf("/stores/%d", store.ID),
			"update": fmt.Sprintf("/stores/%d", store.ID),
			"delete": fmt.Sprintf("/stores/%d", store.ID),
			"create": "/stores",
			"get":    fmt.Sprintf("/stores/%d", store.ID),
		},
	}

	ctx.JSON(http.StatusOK, storeResponse)
}

func (c *StoreController) GetStores(ctx *gin.Context) {
	stores, err := c.service.GetStores()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var storeResponses []gin.H
	for _, store := range stores {
		storeResponses = append(storeResponses, gin.H{
			"store": store,
			"_links": gin.H{
				"self":   fmt.Sprintf("/stores/%d", store.ID),
				"delete": fmt.Sprintf("/stores/%d", store.ID),
				"update": fmt.Sprintf("/stores/%d", store.ID),
			},
		})
	}

	response := gin.H{
		"stores": storeResponses,
		"_links": gin.H{
			"self":   "/stores",
			"create": "/stores",
		},
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *StoreController) GetStore(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store ID"})
		return
	}

	store, err := c.service.GetStore(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	storeResponse := gin.H{
		"store": store,
		"_links": gin.H{
			"self":   fmt.Sprintf("/stores/%d", id),
			"delete": fmt.Sprintf("/stores/%d", id),
			"update": fmt.Sprintf("/stores/%d", id),
			"create": "/stores",
			"get":    fmt.Sprintf("/stores/%d", id),
		},
	}

	ctx.JSON(http.StatusOK, storeResponse)
}

func (c *StoreController) RegisterRoutes(router *gin.Engine) {
	router.POST("/stores", c.AddStore)
	router.PUT("/stores", c.UpdateStore)
	router.DELETE("/stores/:id", c.DeleteStore)
	router.GET("/stores", c.GetStores)
	router.GET("/stores/:id", c.GetStore)
}
