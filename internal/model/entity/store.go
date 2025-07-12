package entity

import (
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
