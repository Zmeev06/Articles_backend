package handlers

import (
	"fmt"
	"math"
	"strconv"
	"time"
	"web_practicum/database"
	. "web_practicum/models"
	. "web_practicum/services"

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
	if err := db.Model(&article).Update("times_visited", article.TimesVisited + 1).Error; err != nil {
		return err
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

// @Summary Create new article, return a created article ID
// @Success 200 {object} int
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
			Sanitize(articleReq.Title), time.Now().Format("02-01")),
		Subtitle:  articleReq.Subtitle,
		Theme:     articleReq.Theme,
		Content:   articleReq.Content,
		CreatedAt: time.Now().Format("02.01.2006"),
	}

	if err := db.Create(&article).Error; err != nil {
		return err
	}
	if err := WriteQrCode(article); err != nil {
		return err
	}
	return ctx.JSON(article.ID)
}

// @Summary Get article qr code image path
// @Success 200 {object} int
// @Failure 400
// @Failure 500
// @Param id path int true "article ID"
// @Router /api/qr-codes/{id} [get]
func GetArticleQrcode(c *fiber.Ctx) error {
	return c.JSON(fmt.Sprintf("static/qr-codes/%s.png", c.Params("id")))
}

// @Summary Get article estimated read time
// @Success 200 {object} int
// @Failure 400
// @Failure 500
// @Param id path int true "article ID"
// @Router /api/reading-time/{id} [get]
func GetArticleReadingTime(c *fiber.Ctx) error {

	id, err := strconv.ParseUint(c.Params("id"), 10, 0)

	if err != nil {
		return err
	}
	article, err := GetArticleById(id)
	if err != nil {
		return err
	}
	return c.JSON(
		math.Ceil(
			TimeToRead(article) / 60))
}
