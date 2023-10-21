package models

// import "time"

type Article struct {
	ID              uint64 `json:"id"`
	Title           string `json:"title"`
	NormalisedTitle string `json:"normalised_title"`
	Subtitle        string `json:"subtitle"`
	Theme           string `json:"theme"`
	Content         string `json:"content"`
	CreatedAt       string `json:"created_at"`
	TimesVisited    uint64 `json:"times_visited"`
}
