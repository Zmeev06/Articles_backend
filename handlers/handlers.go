package handlers

import (
	"fmt"
	"time"
	"web_practicum/database"
	. "web_practicum/models"
	"web_practicum/services"

	"github.com/gofiber/fiber/v2"
)

// @Summary Get article by title
// @Success 200 {object} Article{}
// @Failure 404
// @Param title path string true "normalised article title"
// @Router /api/article/{title} [get]
func GetArticle(ctx *fiber.Ctx) error {

	db := database.DB
	var article Article
	if err := db.First(&article, Article{NormalisedTitle: ctx.Params("title")}).Error; err != nil {
		fmt.Println(err)
		return fiber.ErrNotFound
	}
	return ctx.JSON(article)
}

// @Summary Get all articles
// @Success 200 {object} []Article{}
// @Failure 400
// @Router /api/articles [get]
func GetAllArticles(ctx *fiber.Ctx) error {

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
// @Param request body handlers.CreateArticle.ArticleReqBody true "Subset of article fields"
// @Router /api/create/article [post]
func CreateArticle(ctx *fiber.Ctx) error {

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

