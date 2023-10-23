package routing

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"

	swagger "github.com/gofiber/swagger"

	. "web_practicum/handlers"
)

func Setup(app *fiber.App) {

	app.Get("/swagger/*", swagger.New(swagger.Config{}))

	static := app.Group("/static")
	static.Static("/qr-codes", "./static/qr-codes")

	api := app.Group("/api")
	api.Post("/create/article", CreateArticle)
	api.Get("/article/:title", GetArticle)
	api.Get("/articles", GetAllArticles)
	app.Use("/", filesystem.New(filesystem.Config{
		Root:         http.Dir("dist"),
		NotFoundFile: "index.html"}))
}
