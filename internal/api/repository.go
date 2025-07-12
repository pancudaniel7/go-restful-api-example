package api

import "github.com/pancudaniel7/go-restful-api-example/internal/model/entity"

type BookRepository interface {
	AddBook(book *entity.Book) (*entity.Book, error)
	UpdateBook(book *entity.Book) (*entity.Book, error)
	DeleteBook(id uint) error
	GetBooks() ([]entity.Book, error)
	GetBook(id uint) (*entity.Book, error)
}

type StoreRepository interface {
	AddStore(store *entity.Store) (*entity.Store, error)
	UpdateStore(store *entity.Store) (*entity.Store, error)
	DeleteStore(id uint) error
	GetStores() ([]*entity.Store, error)
	GetStore(id uint) (*entity.Store, error)
}
