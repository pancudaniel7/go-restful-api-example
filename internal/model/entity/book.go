package entity

import (
	"database/sql"
	"errors"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title         string `gorm:"not null" validate:"required"`
	Author        string `gorm:"not null" validate:"required"`
	PublishedDate sql.NullTime
	StoreID       uint `gorm:"not null" validate:"required"`
}

func (Book) TableName() string {
	return "books"
}

func (b *Book) BeforeSave(tx *gorm.DB) (err error) {
	if b.Title == "" {
		return errors.New("title is required")
	}
	if b.Author == "" {
		return errors.New("author is required")
	}
	if b.StoreID == 0 {
		return errors.New("store_id is required")
	}
	return nil
}
