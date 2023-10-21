package main

import (
	"log"
	"os"
	"strings"

	"web_practicum/database"
	. "web_practicum/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/mattn/go-sqlite3"

	_ "web_practicum/docs"

	swagger "github.com/gofiber/swagger"
)

func main() {

	if err := database.Setup(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOriginsFunc: nil,
		AllowOrigins:     "*",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
		AllowHeaders:     "*",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}))
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
