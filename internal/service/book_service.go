package services

import (
	internal "github.com/pancudaniel7/go-restful-api-example/internal/model"
	"gorm.io/gorm"
	"log"
	"time"
)

type BookService struct {
	db *gorm.DB
}

func NewBookService(db *gorm.DB) *BookService {
	return &BookService{db: db}
}

// AddBook adds a new book to a store
func (s *BookService) AddBook(title, author string, publishedDate time.Time, storeID uint) (*internal.Book, error) {
	book := internal.Book{Title: title, Author: author, PublishedDate: publishedDate, StoreID: storeID}
	result := s.db.Create(&book)
	if result.Error != nil {
		log.Println("Error creating book:", result.Error)
		return nil, result.Error
	}
	return &book, nil
}

// AddPage adds a new page to a book
func (s *BookService) AddPage(bookID uint, pageNumber int, content string) (*internal.Page, error) {
	page := internal.Page{BookID: bookID, PageNumber: pageNumber, Content: content}
	result := s.db.Create(&page)
	if result.Error != nil {
		log.Println("Error creating page:", result.Error)
		return nil, result.Error
	}
	return &page, nil
}

// FindBookByTitle finds a book by its title
func (s *BookService) FindBookByTitle(title string) (*internal.Book, error) {
	var book internal.Book
	result := s.db.First(&book, "title = ?", title)
	if result.Error != nil {
		log.Println("Error finding book:", result.Error)
		return nil, result.Error
	}
	return &book, nil
}

// UpdateBookTitle updates the title of a book
func (s *BookService) UpdateBookTitle(bookID uint, newTitle string) error {
	result := s.db.Model(&internal.Book{}).Where("id = ?", bookID).Update("title", newTitle)
	return result.Error
}

// DeleteBook deletes a book
func (s *BookService) DeleteBook(bookID uint) error {
	result := s.db.Delete(&internal.Book{}, bookID)
	return result.Error
}
