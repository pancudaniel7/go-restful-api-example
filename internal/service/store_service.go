package services

import (
	"github.com/pancudaniel7/go-restful-api-example/internal/model/dto"
	internal "github.com/pancudaniel7/go-restful-api-example/internal/model/entity"
	"log"

	"gorm.io/gorm"
)

type StoreServiceImpl struct {
	db *gorm.DB
}

func NewStoreService(db *gorm.DB) *StoreServiceImpl {
	return &StoreServiceImpl{db: db}
}

// AddStore adds a new store to the database
func (s *StoreServiceImpl) AddStore(storeDTO dto.StoreDTO) (*internal.Store, error) {
	store := internal.Store{Name: storeDTO.Name, Location: storeDTO.Location}
	result := s.db.Create(&store)
	if result.Error != nil {
		log.Println("Error creating store:", result.Error)
		return nil, result.Error
	}
	return &store, nil
}

// UpdateStore updates a store in the database
func (s *StoreServiceImpl) UpdateStore(storeDTO dto.StoreDTO) (*internal.Store, error) {
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
func (s *StoreServiceImpl) DeleteStore(id uint) error {
	result := s.db.Delete(&internal.Store{}, id)
	if result.Error != nil {
		log.Println("Error deleting store:", result.Error)
		return result.Error
	}
	return nil
}

// GetStores retrieves all stores from the database
func (s *StoreServiceImpl) GetStores() ([]internal.Store, error) {
	var stores []internal.Store
	result := s.db.Find(&stores)
	if result.Error != nil {
		log.Println("Error retrieving stores:", result.Error)
		return nil, result.Error
	}
	return stores, nil
}

// GetStore retrieves a store by its ID from the database
func (s *StoreServiceImpl) GetStore(id uint) (*internal.Store, error) {
	store := &internal.Store{}
	result := s.db.First(store, id)
	if result.Error != nil {
		log.Println("Error finding store:", result.Error)
		return nil, result.Error
	}
	return store, nil
}
