package services

import (
	"github.com/pancudaniel7/go-restful-api-example/internal/dto"
	internal "github.com/pancudaniel7/go-restful-api-example/internal/entity"
	"gorm.io/gorm"
	"log"
)

type BookService struct {
	db *gorm.DB
}

func NewBookService(db *gorm.DB) *BookService {
	return &BookService{db: db}
}

// AddBook adds a new book to a store
func (s *BookService) AddBook(bookDTO dto.BookDTO) (*internal.Book, error) {
	book := internal.Book{Title: bookDTO.Title, Author: bookDTO.Author, PublishedDate: bookDTO.PublishedDate, StoreID: bookDTO.StoreID}
	result := s.db.Create(&book)
	if result.Error != nil {
		log.Println("Error creating book:", result.Error)
		return nil, result.Error
	}
	return &book, nil
}

// UpdateBook updates a book in a store
func (s *BookService) UpdateBook(bookDTO dto.BookDTO) (*internal.Book, error) {
	book := &internal.Book{}
	result := s.db.First(book, bookDTO.ID)
	if result.Error != nil {
		log.Println("Error finding book:", result.Error)
		return nil, result.Error
	}

	book.Title = bookDTO.Title
	book.Author = bookDTO.Author
	book.PublishedDate = bookDTO.PublishedDate
	book.StoreID = bookDTO.StoreID

	result = s.db.Save(&book)
	if result.Error != nil {
		log.Println("Error updating book:", result.Error)
		return nil, result.Error
	}
	return book, nil
}

// DeleteBook deletes a book from a store
func (s *BookService) DeleteBook(id uint) error {
	result := s.db.Delete(&internal.Book{}, id)
	if result.Error != nil {
		log.Println("Error deleting book:", result.Error)
		return result.Error
	}
	return nil
}
