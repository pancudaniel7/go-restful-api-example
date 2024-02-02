package internal

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	gorm.Model
	Title         string
	Author        string
	PublishedDate time.Time
	StoreID       uint
	Pages         []Page
}

// TableName overrides the table name used by User to `store`
func (Book) TableName() string {
	return "books"
}

type Page struct {
	gorm.Model
	BookID     uint
	PageNumber int
	Content    string
}

// TableName overrides the table name used by User to `store`
func (Page) TableName() string {
	return "pages"
}
