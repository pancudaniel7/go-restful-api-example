package services

import (
	"github.com/pancudaniel7/go-restful-api-example/internal/model/dto"
	internal "github.com/pancudaniel7/go-restful-api-example/internal/model/entity"
	"github.com/pancudaniel7/go-restful-api-example/internal/utils"

	"gorm.io/gorm"
)

type BookServiceImpl struct {
	db *gorm.DB
}

func NewBookService(db *gorm.DB) *BookServiceImpl {
	return &BookServiceImpl{db: db}
}

func (s *BookServiceImpl) AddBook(bookDTO dto.BookDTO) (*internal.Book, error) {
	book := internal.Book{
		Title:         bookDTO.Title,
		Author:        bookDTO.Author,
		PublishedDate: utils.ConvertTimeToNullTime(bookDTO.PublishedDate),
		StoreID:       bookDTO.StoreID,
	}
	result := s.db.Create(&book)
	if result.Error != nil {
		utils.Log().Error("Error creating book:", result.Error)
		return nil, result.Error
	}

	utils.Log().Debug("Book created:", book)
	return &book, nil
}

func (s *BookServiceImpl) UpdateBook(bookDTO dto.BookDTO) (*internal.Book, error) {
	book := &internal.Book{}
	result := s.db.First(book, bookDTO.ID)
	if result.Error != nil {
		utils.Log().Error("Error finding book: %s", result.Error.Error())
		return nil, result.Error
	}

	book.Title = bookDTO.Title
	book.Author = bookDTO.Author
	book.PublishedDate = utils.ConvertTimeToNullTime(bookDTO.PublishedDate)
	book.StoreID = bookDTO.StoreID

	result = s.db.Save(&book)
	if result.Error != nil {
		utils.Log().Error("Error updating book: %s", result.Error.Error())
		return nil, result.Error
	}

	utils.Log().Debug("Book updated for id: %d", book.ID)
	return book, nil
}

func (s *BookServiceImpl) DeleteBook(id uint) error {
	result := s.db.Delete(&internal.Book{}, id)
	if result.Error != nil {
		utils.Log().Error("Error deleting book: %s", result.Error.Error())
		return result.Error
	}

	utils.Log().Debug("Book deleted: %d", id)
	return nil
}

func (s *BookServiceImpl) GetBooks() ([]internal.Book, error) {
	var books []internal.Book
	result := s.db.Find(&books)
	if result.Error != nil {
		utils.Log().Error("Error retrieving books: %s", result.Error.Error())
		return nil, result.Error
	}

	utils.Log().Debug("All books found")
	return books, nil
}

func (s *BookServiceImpl) GetBook(id uint) (*internal.Book, error) {
	book := &internal.Book{}
	result := s.db.First(book, id)
	if result.Error != nil {
		utils.Log().Warn("Error finding book: %v", result.Error)
		return nil, result.Error
	}

	utils.Log().Debug("Book found: %d", book.ID)
	return book, nil
}
