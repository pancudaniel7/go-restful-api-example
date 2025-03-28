package services

import (
	"github.com/pancudaniel7/go-restful-api-example/internal/model/dto"
	internal "github.com/pancudaniel7/go-restful-api-example/internal/model/entity"
	"github.com/pancudaniel7/go-restful-api-example/internal/utils"
	"log"

	"gorm.io/gorm"
)

type StoreServiceImpl struct {
	db *gorm.DB
}

func NewStoreService(db *gorm.DB) *StoreServiceImpl {
	return &StoreServiceImpl{db: db}
}

func (s *StoreServiceImpl) AddStore(storeDTO dto.StoreDTO) (*internal.Store, error) {
	store := internal.Store{Name: storeDTO.Name, Location: storeDTO.Location}
	result := s.db.Create(&store)
	if result.Error != nil {
		utils.Log().Info("Error creating store:", result.Error)
		return nil, result.Error
	}
	return &store, nil
}

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

func (s *StoreServiceImpl) DeleteStore(id uint) error {
	result := s.db.Delete(&internal.Store{}, id)
	if result.Error != nil {
		log.Println("Error deleting store:", result.Error)
		return result.Error
	}
	return nil
}

func (s *StoreServiceImpl) GetStores() ([]internal.Store, error) {
	var stores []internal.Store
	result := s.db.Find(&stores)
	if result.Error != nil {
		log.Println("Error retrieving stores:", result.Error)
		return nil, result.Error
	}
	return stores, nil
}

func (s *StoreServiceImpl) GetStore(id uint) (*internal.Store, error) {
	store := &internal.Store{}
	result := s.db.First(store, id)
	if result.Error != nil {
		log.Println("Error finding store:", result.Error)
		return nil, result.Error
	}
	return store, nil
}
