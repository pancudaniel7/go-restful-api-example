package dto

import "time"

type BookDTO struct {
	ID            uint      `json:"id"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	PublishedDate time.Time `json:"published_date"`
	StoreID       uint      `json:"store_id"`
}

type PageDTO struct {
	BookID     uint   `json:"book_id"`
	PageNumber int    `json:"page_number"`
	Content    string `json:"content"`
}
