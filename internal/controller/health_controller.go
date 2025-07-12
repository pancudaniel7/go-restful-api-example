package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthControllerImpl struct {
}

func NewHealthController() *HealthControllerImpl {
	return &HealthControllerImpl{}
}

func (c *HealthControllerImpl) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "UP",
		"_links": gin.H{
			"self": "/health",
		},
	})
}
