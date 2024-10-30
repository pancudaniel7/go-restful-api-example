package dto

type StoreDTO struct {
	ID       uint      `json:"id"`
	Name     string    `json:"name"`
	Location string    `json:"location"`
	Books    []BookDTO `json:"books"`
}
