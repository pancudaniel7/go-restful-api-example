package controller

import (
	"fmt"
	"github.com/pancudaniel7/go-restful-api-example/internal/api"
	"github.com/pancudaniel7/go-restful-api-example/internal/model/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookControllerImpl struct {
	service api.BookService
}

func NewBookController(service api.BookService) *BookControllerImpl {
	return &BookControllerImpl{service: service}
}

func (c *BookControllerImpl) AddBook(ctx *gin.Context) {
	var bookDTO dto.BookDTO
	err := ctx.ShouldBindJSON(&bookDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := c.service.AddBook(&bookDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	bookResponse := gin.H{
		"book": book,
		"_links": gin.H{
			"self":   fmt.Sprintf("/books/%d", book.ID),
			"update": fmt.Sprintf("/books/%d", book.ID),
			"delete": fmt.Sprintf("/books/%d", book.ID),
			"create": "/books",
			"get":    fmt.Sprintf("/books/%d", book.ID),
		},
	}

	ctx.JSON(http.StatusCreated, bookResponse)
}

func (c *BookControllerImpl) UpdateBook(ctx *gin.Context) {
	var bookDTO dto.BookDTO
	if err := ctx.ShouldBindJSON(&bookDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := c.service.UpdateBook(&bookDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	bookResponse := gin.H{
		"book": book,
		"_links": gin.H{
			"self":   fmt.Sprintf("/books/%d", book.ID),
			"update": fmt.Sprintf("/books/%d", book.ID),
			"delete": fmt.Sprintf("/books/%d", book.ID),
			"create": "/books",
			"get":    fmt.Sprintf("/books/%d", book.ID),
		},
	}

	ctx.JSON(http.StatusOK, bookResponse)
}

func (c *BookControllerImpl) DeleteBook(ctx *gin.Context) {
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

	ctx.JSON(http.StatusNoContent, gin.H{
		"message": "Book deleted",
		"_links": gin.H{
			"self":   fmt.Sprintf("/books/%d", id),
			"create": "/books",
			"getAll": "/books",
		},
	})
}

func (c *BookControllerImpl) GetBooks(ctx *gin.Context) {
	books, err := c.service.GetBooks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var bookResponses []gin.H
	for _, book := range books {
		bookResponses = append(bookResponses, gin.H{
			"book": book,
			"_links": gin.H{
				"self":   fmt.Sprintf("/books/%d", book.ID),
				"delete": fmt.Sprintf("/books/%d", book.ID),
				"update": fmt.Sprintf("/books/%d", book.ID),
			},
		})
	}

	response := gin.H{
		"books": bookResponses,
		"_links": gin.H{
			"self":   "/books",
			"create": "/books",
		},
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *BookControllerImpl) GetBook(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	book, err := c.service.GetBook(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	bookResponse := gin.H{
		"book": book,
		"_links": gin.H{
			"self":   fmt.Sprintf("/books/%d", id),
			"delete": fmt.Sprintf("/books/%d", id),
			"update": fmt.Sprintf("/books/%d", id),
			"create": "/books",
			"get":    fmt.Sprintf("/books/%d", id),
		},
	}

	ctx.JSON(http.StatusOK, bookResponse)
}
