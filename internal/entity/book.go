package internal

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
