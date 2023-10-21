package services

import (
	"fmt"
	"os"
	. "web_practicum/models"
)

func MakeLink(article Article) string {
	hostname := os.Getenv("HOST_URL")
	return fmt.Sprintf("%s/article/%s", hostname, article.NormalisedTitle)
}
