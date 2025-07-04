package entity

import (
	"errors"
	"gorm.io/gorm"
)

type Store struct {
	gorm.Model
	Name     string
	Location string
	Books    []Book
}

func (Store) TableName() string {
	return "store"
}

func (s *Store) BeforeSave(tx *gorm.DB) (err error) {
	if s.Name == "" {
		return errors.New("name is required")
	}
	if s.Location == "" {
		return errors.New("location is required")
	}
	return nil
}
