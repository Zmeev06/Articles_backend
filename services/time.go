package services

import (
	"fmt"
	"web_practicum/database"
	. "web_practicum/models"

	"github.com/gofiber/fiber/v2"
)

func TimeToRead(article Article) float64 {
	return float64(len(article.Content)) / 2000
}

func GetArticleById(id uint64) (Article, error) {
	db := database.DB
	var article Article
	if err := db.First(&article, id).Error; err != nil {
		fmt.Println(err)
		return article, fiber.ErrNotFound
	}
	return article, nil
}
