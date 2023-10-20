package models

import "time"

type Article struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	Subtitle  string    `json:"subtitle"`
	Theme     string    `json:"theme"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	QrCode    string    `json:"qr_code"`
}
