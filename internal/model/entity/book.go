package entity

import (
	"database/sql"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title         string
	Author        string
	PublishedDate sql.NullTime
	StoreID       uint
}

func (Book) TableName() string {
	return "books"
}

type Page struct {
	gorm.Model
	BookID     uint
	PageNumber int
	Content    string
}

func (Page) TableName() string {
	return "pages"
}
