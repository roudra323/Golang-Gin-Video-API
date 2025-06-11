package domain

import "time"

type Video struct {
	ID          uint64    `json:"id"`
	Title       string    `json:"title" bindding:"required"`
	Description string    `json:"description"`
	URL         string    `json:"url" binding:"required,url"`
	Author      string    `json:"author"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
