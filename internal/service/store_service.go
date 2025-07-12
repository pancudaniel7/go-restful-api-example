package services

import (
	"fmt"
	"github.com/pancudaniel7/go-restful-api-example/internal/api"
	"github.com/pancudaniel7/go-restful-api-example/internal/model/dto"
	"github.com/pancudaniel7/go-restful-api-example/internal/model/entity"
)

type StoreServiceImpl struct {
	storeRepository api.StoreRepository
}

func NewStoreServiceImpl(storeRepository api.StoreRepository) *StoreServiceImpl {
	return &StoreServiceImpl{storeRepository: storeRepository}
}

func (ss *StoreServiceImpl) AddStore(storeDTO *dto.StoreDTO) (*dto.StoreDTO, error) {
	store := &entity.Store{Name: storeDTO.Name, Location: storeDTO.Location}
	store, err := ss.storeRepository.AddStore(store)
	if err != nil {
		return nil, err
	}
	return storeDTO, nil
}

func (ss *StoreServiceImpl) UpdateStore(storeDTO *dto.StoreDTO) (*dto.StoreDTO, error) {
	store := &entity.Store{Name: storeDTO.Name, Location: storeDTO.Location}

	if storeDTO.ID == 0 {
		return nil, fmt.Errorf("store ID must be provided for update")
	}
	
	store.ID = storeDTO.ID
	store, err := ss.storeRepository.UpdateStore(store)
	if err != nil {
		return nil, err
	}
	return storeDTO, nil
}

func (ss *StoreServiceImpl) DeleteStore(id uint) error {
	err := ss.storeRepository.DeleteStore(id)
	if err != nil {
		return err
	}
	return nil
}

func (ss *StoreServiceImpl) GetStores() ([]*dto.StoreDTO, error) {
	stores, err := ss.storeRepository.GetStores()
	if err != nil {
		return nil, err
	}

	storeDTOs := make([]*dto.StoreDTO, len(stores))
	for i, store := range stores {
		storeDTOs[i] = &dto.StoreDTO{
			Name:     store.Name,
			Location: store.Location,
		}
	}

	return storeDTOs, nil
}

func (ss *StoreServiceImpl) GetStore(id uint) (*dto.StoreDTO, error) {
	store, err := ss.storeRepository.GetStore(id)
	if err != nil {
		return nil, err
	}

	storeDTO := &dto.StoreDTO{
		Name:     store.Name,
		Location: store.Location,
	}
	return storeDTO, nil
}
