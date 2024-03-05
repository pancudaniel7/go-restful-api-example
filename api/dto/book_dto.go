package dto

import "time"

type BookDTO struct {
	ID            uint      `json:"id"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	PublishedDate time.Time `json:"publishedDate"`
	StoreID       uint      `json:"storeId"`
}

type PageDTO struct {
	BookID     uint   `json:"bookId"`
	PageNumber int    `json:"pageNumber"`
	Content    string `json:"content"`
}
