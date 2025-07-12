package api

import "github.com/gin-gonic/gin"

type HealthController interface {
	Health(c *gin.Context)
}

type BookController interface {
	AddBook(c *gin.Context)
	UpdateBook(ctx *gin.Context)
	DeleteBook(ctx *gin.Context)
	GetBooks(ctx *gin.Context)
	GetBook(ctx *gin.Context)
}

type StoreController interface {
	AddStore(c *gin.Context)
	DeleteStore(c *gin.Context)
	UpdateStore(c *gin.Context)
	GetStores(c *gin.Context)
	GetStore(ctx *gin.Context)
}
