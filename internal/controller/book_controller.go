package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pancudaniel7/go-restful-api-example/internal/dto"
	services "github.com/pancudaniel7/go-restful-api-example/internal/service"
	"net/http"
	"strconv"
)

type BookController struct {
	service *services.BookService
}

func NewBookController(service *services.BookService) *BookController {
	return &BookController{service: service}
}

func (c *BookController) AddBook(ctx *gin.Context) {
	var bookDTO dto.BookDTO
	if err := ctx.ShouldBindJSON(&bookDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := c.service.AddBook(bookDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (c *BookController) UpdateBook(ctx *gin.Context) {
	var bookDTO dto.BookDTO
	if err := ctx.ShouldBindJSON(&bookDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := c.service.UpdateBook(bookDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (c *BookController) DeleteBook(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	err = c.service.DeleteBook(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (c *BookController) RegisterRoutes(router *gin.Engine) {
	router.POST("/book", c.AddBook)
	router.PUT("/book", c.UpdateBook)
	router.DELETE("/book/:id", c.DeleteBook)
}
