package models

type Article struct {
	ID              uint64 `json:"id"`
	Title           string `json:"title" gorm:"type:VARCHAR(300)"`
	NormalisedTitle string `json:"normalised_title" gorm:"type:VARCHAR(300)"`
	Subtitle        string `json:"subtitle" gorm:"type:VARCHAR(300)"`
	Theme           string `json:"theme" gorm:"type:VARCHAR(10)"`
	Content         string `json:"content"`
	CreatedAt       string `json:"created_at" gorm:"type:VARCHAR(10)"`
	TimesVisited    uint64 `json:"times_visited"`
	ReadingTime     uint64 `json:reading_time"`
	Cover           string `json:"cover"`
}
