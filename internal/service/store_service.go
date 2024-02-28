package services

import (
	"github.com/pancudaniel7/go-restful-api-example/internal/dto"
	internal "github.com/pancudaniel7/go-restful-api-example/internal/entity"
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
func (s *StoreService) AddStore(storeDTO dto.StoreDTO) (*internal.Store, error) {
	store := internal.Store{Name: storeDTO.Name, Location: storeDTO.Location}
	result := s.db.Create(&store)
	if result.Error != nil {
		log.Println("Error creating store:", result.Error)
		return nil, result.Error
	}
	return &store, nil
}

// UpdateStore updates a store in the database
func (s *StoreService) UpdateStore(storeDTO dto.StoreDTO) (*internal.Store, error) {
	store := &internal.Store{}
	result := s.db.First(store, storeDTO.ID)
	if result.Error != nil {
		log.Println("Error finding store:", result.Error)
		return nil, result.Error
	}

	store.Name = storeDTO.Name
	store.Location = storeDTO.Location

	result = s.db.Save(&store)
	if result.Error != nil {
		log.Println("Error updating store:", result.Error)
		return nil, result.Error
	}
	return store, nil
}

// DeleteStore deletes a store from the database
func (s *StoreService) DeleteStore(id uint) error {
	result := s.db.Delete(&internal.Store{}, id)
	if result.Error != nil {
		log.Println("Error deleting store:", result.Error)
		return result.Error
	}
	return nil
}
