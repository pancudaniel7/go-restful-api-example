package dto

import (
	"github.com/go-playground/validator/v10"
)

type StoreDTO struct {
	ID       uint      `json:"id"`
	Name     string    `json:"name" validate:"required"`
	Location string    `json:"location" validate:"required"`
	Books    []BookDTO `json:"books"`
}

func (s *StoreDTO) Validate() error {
	validate := validator.New()
	return validate.Struct(s)
}
