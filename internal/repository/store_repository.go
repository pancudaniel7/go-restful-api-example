package services

import (
	"github.com/pancudaniel7/go-restful-api-example/internal/model/entity"
	"github.com/pancudaniel7/go-restful-api-example/internal/utils"
	"gorm.io/gorm"
	"log"
)

type StoreRepositoryImpl struct {
	db *gorm.DB
}

func NewStoreRepository(db *gorm.DB) *StoreRepositoryImpl {
	return &StoreRepositoryImpl{db: db}
}

func (s *StoreRepositoryImpl) AddStore(store *entity.Store) (*entity.Store, error) {
	result := s.db.Create(store)
	if result.Error != nil {
		utils.Log().Info("Error creating store:", result.Error)
		return nil, result.Error
	}
	return store, nil
}

func (s *StoreRepositoryImpl) UpdateStore(store *entity.Store) (*entity.Store, error) {
	if err := s.db.Model(&entity.Store{}).Where("id = ?", store.ID).
		Select("*").Updates(store).Error; err != nil {
		utils.Log().Info("Error updating store:", err)
		return nil, err
	}
	return store, nil
}

func (s *StoreRepositoryImpl) DeleteStore(id uint) error {
	result := s.db.Delete(&entity.Store{}, id)
	if result.Error != nil {
		log.Println("Error deleting store:", result.Error)
		return result.Error
	}
	return nil
}

func (s *StoreRepositoryImpl) GetStores() ([]*entity.Store, error) {
	var stores []*entity.Store
	result := s.db.Find(stores)
	if result.Error != nil {
		log.Println("Error retrieving stores:", result.Error)
		return nil, result.Error
	}
	return stores, nil
}

func (s *StoreRepositoryImpl) GetStore(id uint) (*entity.Store, error) {
	var store *entity.Store
	result := s.db.First(store, id)
	if result.Error != nil {
		log.Println("Error finding store:", result.Error)
		return nil, result.Error
	}
	return store, nil
}
