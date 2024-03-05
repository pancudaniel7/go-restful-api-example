package internal

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	Name     string
	Location string
	Books    []Book
}

// TableName overrides the table name used by User to `store`
func (Store) TableName() string {
	return "store"
}
