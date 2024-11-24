package services

import (
	"github.com/pancudaniel7/go-restful-api-example/internal/model/dto"
	internal "github.com/pancudaniel7/go-restful-api-example/internal/model/entity"
)

type BookService interface {
	AddBook(bookDTO dto.BookDTO) (*internal.Book, error)
	UpdateBook(bookDTO dto.BookDTO) (*internal.Book, error)
	DeleteBook(id uint) error
	GetBooks() ([]internal.Book, error)
	GetBook(id uint) (*internal.Book, error)
}
