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

	swagger "github.com/gofiber/swagger"
	_ "web_practicum/docs"
)

func main() {

	if err := database.Setup(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()
	app.Get("/swagger/*", swagger.New(swagger.Config{}))
	api := app.Group("/api")
	// app.Post("api/images", postImage)
	app.Static("/images", "./images")
	api.Post("/create/article", createArticle)
	api.Get("/article/:id<int>", getArticle)
	api.Get("/shit", func(c *fiber.Ctx) error {
		c.Response().SetBodyString("hello")
		return nil
	})
	api.Get("/articles", getAllArticles)
	// app.Get("/api/articles/:user")
	// app.Get("/api/user/articles")
	log.Fatal(app.Listen(":8080"))
}

// @Summary Get article by id
// @Success 200 {object} Article{}
// @Failure 404
// @Param id path int true "Book ID"
// @Router /api/article/{id} [get]
func getArticle(ctx *fiber.Ctx) error {

	db := database.DB
	var article Article
	if err := db.First(&article, ctx.Params("id")).Error; err != nil {
		fmt.Println(err)
		return fiber.ErrNotFound
	}
	return ctx.JSON(article)
}

// @Summary Get all articles
// @Success 200 {object} []Article{}
// @Failure 400
// @Router /api/articles [get]
func getAllArticles(ctx *fiber.Ctx) error {

	db := database.DB
	var articles []Article
	obj := db.Find(&articles)
	if obj.Error != nil {
		return fiber.ErrBadRequest
	}
	return ctx.JSON(articles)
}

// @Summary Create new article, get a normalised title string
// @Success 200 {object} string
// @Failure 400
// @Failure 500
// @Param request body main.createArticle.ArticleReqBody true "Subset of article fields"
// @Router /api/create/article [post]
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
		fmt.Println(err)
		return fiber.ErrBadRequest
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
		fmt.Println(err)
		return fiber.ErrInternalServerError
	}
	return ctx.JSON(article.NormalisedTitle)
}

