package services

import (
	"github.com/pancudaniel7/go-restful-api-example/internal/model/entity"
	"github.com/pancudaniel7/go-restful-api-example/internal/utils"

	"gorm.io/gorm"
)

type BookRepositoryImpl struct {
	db *gorm.DB
}

func NewBookRepositoryImpl(db *gorm.DB) *BookRepositoryImpl {
	return &BookRepositoryImpl{db: db}
}

func (s *BookRepositoryImpl) AddBook(book *entity.Book) (*entity.Book, error) {
	result := s.db.Create(book)
	if result.Error != nil {
		utils.Log().Error("Error creating book:", result.Error)
		return nil, result.Error
	}

	utils.Log().Debug("Book created:", book)
	return book, nil
}

func (s *BookRepositoryImpl) UpdateBook(book *entity.Book) (*entity.Book, error) {
	if err := s.db.Model(&entity.Book{}).Where("id = ?", book.ID).Updates(book).Error; err != nil {
		utils.Log().Error("Error updating book: %s", err.Error())
		return nil, err
	}

	utils.Log().Debug("Book updated for id: %d", book.ID)
	return book, nil
}

func (s *BookRepositoryImpl) DeleteBook(id uint) error {
	result := s.db.Delete(&entity.Book{}, id)
	if result.Error != nil {
		utils.Log().Error("Error deleting book: %s", result.Error.Error())
		return result.Error
	}

	utils.Log().Debug("Book deleted: %d", id)
	return nil
}

func (s *BookRepositoryImpl) GetBooks() ([]entity.Book, error) {
	var books []entity.Book
	result := s.db.Find(&books)
	if result.Error != nil {
		utils.Log().Error("Error retrieving books: %s", result.Error.Error())
		return nil, result.Error
	}

	utils.Log().Debug("All books found")
	return books, nil
}

func (s *BookRepositoryImpl) GetBook(id uint) (*entity.Book, error) {
	book := &entity.Book{}
	result := s.db.First(book, id)
	if result.Error != nil {
		utils.Log().Warn("Error finding book: %v", result.Error)
		return nil, result.Error
	}

	utils.Log().Debug("Book found: %d", book.ID)
	return book, nil
}
