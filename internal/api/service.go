package api

import (
	"github.com/pancudaniel7/go-restful-api-example/internal/model/dto"
)

type BookService interface {
	AddBook(bookDTO *dto.BookDTO) (*dto.BookDTO, error)
	UpdateBook(bookDTO *dto.BookDTO) (*dto.BookDTO, error)
	DeleteBook(id uint) error
	GetBooks() ([]*dto.BookDTO, error)
	GetBook(id uint) (*dto.BookDTO, error)
}

type StoreService interface {
	AddStore(storeDTO *dto.StoreDTO) (*dto.StoreDTO, error)
	UpdateStore(storeDTO *dto.StoreDTO) (*dto.StoreDTO, error)
	DeleteStore(id uint) error
	GetStores() ([]*dto.StoreDTO, error)
	GetStore(id uint) (*dto.StoreDTO, error)
}
