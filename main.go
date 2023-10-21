package main

import (
	"log"
	"os"

	"web_practicum/database"
	. "web_practicum/handlers"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"

	_ "web_practicum/docs"

	swagger "github.com/gofiber/swagger"
)

func main() {

	if err := database.Setup(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()
	app.Get("/swagger/*", swagger.New(swagger.Config{}))
	api := app.Group("/api")
	// app.Post("api/images", postImage)
	static := app.Group("/static")
	static.Static("/images", "./images")
	static.Static("/qr-codes", "./qr-codes")
	api.Post("/create/article", CreateArticle)
	api.Get("/article/:title", GetArticle)
	api.Get("/qr-code/:id<int>", GetArticleQrcode)
	api.Get("/articles", GetAllArticles)
	// app.Get("/api/articles/:user")
	// app.Get("/api/user/articles")
	log.Fatal(app.Listen(os.Getenv("LISTEN_ADDR")))
}

// func postImage(ctx *fiber.Ctx) error {
// 	return nil
// }
