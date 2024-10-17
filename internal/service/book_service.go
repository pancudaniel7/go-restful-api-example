package services

import (
	"log"

	"github.com/pancudaniel7/go-restful-api-example/api/dto"
	internal "github.com/pancudaniel7/go-restful-api-example/internal/model"
	"github.com/pancudaniel7/go-restful-api-example/internal/utils"
	"gorm.io/gorm"
)

type BookServiceImpl struct {
	db *gorm.DB
}

func NewBookServiceImpl(db *gorm.DB) *BookServiceImpl {
	return &BookServiceImpl{db: db}
}

// AddBook adds a new book to a store
func (s *BookServiceImpl) AddBook(bookDTO dto.BookDTO) (*internal.Book, error) {
	book := internal.Book{
		Title:         bookDTO.Title,
		Author:        bookDTO.Author,
		PublishedDate: utils.ConvertTimeToNullTime(bookDTO.PublishedDate),
		StoreID:       bookDTO.StoreID,
	}
	result := s.db.Create(&book)
	if result.Error != nil {
		log.Println("Error creating book:", result.Error)
		return nil, result.Error
	}
	return &book, nil
}

// UpdateBook updates a book in a store
func (s *BookServiceImpl) UpdateBook(bookDTO dto.BookDTO) (*internal.Book, error) {
	book := &internal.Book{}
	result := s.db.First(book, bookDTO.ID)
	if result.Error != nil {
		log.Println("Error finding book:", result.Error)
		return nil, result.Error
	}

	book.Title = bookDTO.Title
	book.Author = bookDTO.Author
	book.PublishedDate = utils.ConvertTimeToNullTime(bookDTO.PublishedDate)
	book.StoreID = bookDTO.StoreID

	result = s.db.Save(&book)
	if result.Error != nil {
		log.Println("Error updating book:", result.Error)
		return nil, result.Error
	}
	return book, nil
}

// DeleteBook deletes a book from a store
func (s *BookServiceImpl) DeleteBook(id uint) error {
	result := s.db.Delete(&internal.Book{}, id)
	if result.Error != nil {
		log.Println("Error deleting book:", result.Error)
		return result.Error
	}
	return nil
}

// GetBooks retrieves all books from the database
func (s *BookServiceImpl) GetBooks() ([]internal.Book, error) {
	var books []internal.Book
	result := s.db.Find(&books)
	if result.Error != nil {
		log.Println("Error retrieving books:", result.Error)
		return nil, result.Error
	}
	return books, nil
}

// GetBook retrieves a book by its ID from the database
func (s *BookServiceImpl) GetBook(id uint) (*internal.Book, error) {
	book := &internal.Book{}
	result := s.db.First(book, id)
	if result.Error != nil {
		log.Println("Error finding book:", result.Error)
		return nil, result.Error
	}
	return book, nil
}
