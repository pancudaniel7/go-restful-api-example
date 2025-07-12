package entity

import (
	"database/sql"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title         string `gorm:"not null"`
	Author        string `gorm:"not null"`
	PublishedDate sql.NullTime
	StoreID       uint `gorm:"not null"`
}

func (Book) TableName() string {
	return "books"
}
