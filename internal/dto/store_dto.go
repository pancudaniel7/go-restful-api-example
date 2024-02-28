package dto

type StoreDTO struct {
	ID    uint      `json:"id"`
	Name  string    `json:"name"`
	Books []BookDTO `json:"books"`
}
