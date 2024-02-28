package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pancudaniel7/go-restful-api-example/internal/dto"
	services "github.com/pancudaniel7/go-restful-api-example/internal/service"
	"net/http"
	"strconv"
)

type StoreController struct {
	service *services.StoreService
}

func NewStoreController(service *services.StoreService) *StoreController {
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

	ctx.JSON(http.StatusOK, store)
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

	ctx.Status(http.StatusNoContent)
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

	ctx.JSON(http.StatusOK, store)
}

func (c *StoreController) RegisterRoutes(router *gin.Engine) {
	router.POST("/store", c.AddStore)
	router.PUT("/store", c.UpdateStore)
	router.DELETE("/store/:id", c.DeleteStore)
}
