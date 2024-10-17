package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct {
}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (c *HealthController) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "UP",
		"_links": gin.H{
			"self": "/health",
		},
	})
}

func (c *HealthController) RegisterRoutes(router *gin.Engine) {
	router.GET("/health", c.Health)
}
