package main

import (
	"fmt"
	"log"
	"time"

	"web_practicum/database"
	. "web_practicum/models"
	"web_practicum/services"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	if err := database.Setup(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()
	// app.Post("api/images")
	app.Static("api/images", "./images")
	app.Post("/api/create/article", createArticle)
	app.Get("/api/article/:id<int>", getArticle)
	app.Get("/api/shit", func(c *fiber.Ctx) error {
		c.Response().SetBodyString("hello")
		return nil
	})
	app.Get("/api/articles", getAllArticles)
	// app.Get("/api/articles/:user")
	// app.Get("/api/user/articles")
	log.Fatal(app.Listen(":8080"))
}
func getArticle(ctx *fiber.Ctx) error {

	db := database.DB
	var article Article
	if err := db.First(&article, ctx.Params("id")).Error; err != nil {
		log.Fatalln(err)
	}
	return ctx.JSON(article)
}

func getAllArticles(ctx *fiber.Ctx) error {

	db := database.DB
	var articles []Article
	obj := db.Find(&articles)
	if obj.Error != nil {
		return fiber.ErrBadRequest
	}
	return ctx.JSON(articles)
}

func createArticle(ctx *fiber.Ctx) error {

	type ArticleReqBody struct {
		Title    string `json:"title"`
		Subtitle string `json:"subtitle"`
		Theme    string `json:"theme"`
		Content  string `json:"content"`
	}
	db := database.DB
	var articleReq ArticleReqBody
	if err := ctx.BodyParser(&articleReq); err != nil {
		log.Fatalln(err)
	}
	var article = Article{
		Title: articleReq.Title,
		NormalisedTitle: fmt.Sprintf("%s-%s",
			services.Sanitize(articleReq.Title), time.Now().Format("02-01")),
		Subtitle:  articleReq.Subtitle,
		Theme:     articleReq.Theme,
		Content:   articleReq.Content,
		CreatedAt: time.Now().Format("02.01.2006"),
	}

	if err := db.Create(&article).Error; err != nil {
		log.Fatalln(err)
	}
	return ctx.JSON(article.NormalisedTitle)
}

