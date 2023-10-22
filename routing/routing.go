package routing

import (
	"github.com/gofiber/fiber/v2"

	swagger "github.com/gofiber/swagger"

	. "web_practicum/handlers"
)

func Setup(app *fiber.App) {

	app.Get("/swagger/*", swagger.New(swagger.Config{}))
	api := app.Group("/api")
	static := app.Group("/static")
	static.Static("/images", "./images")
	static.Static("/qr-codes", "./qr-codes")
	api.Post("/create/article", CreateArticle)
	api.Get("/article/:title", GetArticle)
	api.Get("/reading-time/:id<int>", GetArticleReadingTime)
	api.Get("/articles", GetAllArticles)
}
