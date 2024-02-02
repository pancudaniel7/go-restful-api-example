package services

import (
	internal "github.com/pancudaniel7/go-restful-api-example/internal/model"
	"gorm.io/gorm"
	"log"
)

type StoreService struct {
	db *gorm.DB
}

func NewStoreService(db *gorm.DB) *StoreService {
	return &StoreService{db: db}
}

// AddStore adds a new store to the database
func (s *StoreService) AddStore(name, location string) (*internal.Store, error) {
	store := internal.Store{Name: name, Location: location}
	result := s.db.Create(&store)
	if result.Error != nil {
		log.Println("Error creating store:", result.Error)
		return nil, result.Error
	}
	return &store, nil
}
