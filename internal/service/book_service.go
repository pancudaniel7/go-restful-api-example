package services

import (
	"fmt"
	"github.com/pancudaniel7/go-restful-api-example/internal/api"
	"github.com/pancudaniel7/go-restful-api-example/internal/model/dto"
	"github.com/pancudaniel7/go-restful-api-example/internal/model/entity"
	"github.com/pancudaniel7/go-restful-api-example/internal/utils"
)

type BookServiceImpl struct {
	bookRepository api.BookRepository
}

func NewBookServiceImpl(bookRepository api.BookRepository) *BookServiceImpl {
	return &BookServiceImpl{bookRepository: bookRepository}
}

func (bs *BookServiceImpl) AddBook(bookDTO *dto.BookDTO) (*dto.BookDTO, error) {
	book := &entity.Book{
		Title:         bookDTO.Title,
		Author:        bookDTO.Author,
		PublishedDate: utils.ConvertTimeToNullTime(bookDTO.PublishedDate),
		StoreID:       bookDTO.StoreID,
	}
	book, err := bs.bookRepository.AddBook(book)
	if err != nil {
		return nil, err
	}

	utils.Log().Debug("Book created:", book)
	return bookDTO, nil
}

func (bs *BookServiceImpl) UpdateBook(bookDTO *dto.BookDTO) (*dto.BookDTO, error) {
	book := &entity.Book{
		Title:         bookDTO.Title,
		Author:        bookDTO.Author,
		PublishedDate: utils.ConvertTimeToNullTime(bookDTO.PublishedDate),
		StoreID:       bookDTO.StoreID,
	}

	if bookDTO.ID == 0 {
		return nil, fmt.Errorf("book ID must be provided for update")
	}

	book.ID = bookDTO.ID
	book, err := bs.bookRepository.UpdateBook(book)
	if err != nil {
		return nil, err
	}

	utils.Log().Debug("Book updated for id: %d", book.ID)
	return bookDTO, nil
}

func (bs *BookServiceImpl) DeleteBook(id uint) error {
	err := bs.bookRepository.DeleteBook(id)
	if err != nil {
		return err
	}

	utils.Log().Debug("Book deleted: %d", id)
	return nil
}

func (bs *BookServiceImpl) GetBooks() ([]*dto.BookDTO, error) {
	books, err := bs.bookRepository.GetBooks()
	if err != nil {
		return nil, err
	}

	bookDTOs := make([]*dto.BookDTO, len(books))
	for i, book := range books {
		bookDTOs[i] = &dto.BookDTO{
			ID:            book.ID,
			Title:         book.Title,
			Author:        book.Author,
			PublishedDate: utils.ConvertNullTimeToTime(book.PublishedDate),
			StoreID:       book.StoreID,
		}
	}

	utils.Log().Debug("All books found, count: %d", len(bookDTOs))
	return bookDTOs, nil
}

func (bs *BookServiceImpl) GetBook(id uint) (*dto.BookDTO, error) {
	book, err := bs.bookRepository.GetBook(id)
	if err != nil {
		return nil, err
	}

	bookDTO := &dto.BookDTO{
		ID:            book.ID,
		Title:         book.Title,
		Author:        book.Author,
		PublishedDate: utils.ConvertNullTimeToTime(book.PublishedDate),
		StoreID:       book.StoreID,
	}

	utils.Log().Debug("Book found: %d", bookDTO.ID)
	return bookDTO, nil
}
