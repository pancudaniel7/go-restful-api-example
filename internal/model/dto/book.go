package dto

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type BookDTO struct {
	ID            uint      `json:"id"`
	Title         string    `json:"title" validate:"required"`
	Author        string    `json:"author" validate:"required"`
	PublishedDate time.Time `json:"publishedDate"`
	StoreID       uint      `json:"storeId" validate:"required,gt=0"`
}

func (b *BookDTO) Validate() error {
	validate := validator.New()
	return validate.Struct(b)
}
