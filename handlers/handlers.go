package handlers

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
	"web_practicum/database"
	. "web_practicum/models"
	. "web_practicum/services"

	"github.com/gofiber/fiber/v2"
)

// @Summary Get article by title and increment its view counter
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
	if err := db.Model(&article).Update("times_visited", article.TimesVisited+1).Error; err != nil {
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

// @Summary Create new article, return a created article link and QR code
// @Success 200 {object} handlers.CreateArticle.Response
// @Failure 400
// @Failure 500
// @Param request body Article true "Subset of article fields"
// @Router /api/create/article [post]
func CreateArticle(ctx *fiber.Ctx) error {

	type Response struct {
		Link   string `json:"link"`
		QrCode string `json:"qr_code"`
	}
	db := database.DB
	var article Article
	if err := ctx.BodyParser(&article); err != nil {
		fmt.Println(err)
		return fiber.ErrBadRequest
	}
	article.NormalisedTitle = fmt.Sprintf(
		"%s-%s",
		Sanitize(article.Title),
		time.Now().Format("02-01"))
	article.CreatedAt = time.Now().Format("02.01.2006")

	if err := db.Create(&article).Error; err != nil {
		return err
	}
	path, err := WriteQrCode(article)
	if err != nil {
		return err
	}
	return ctx.JSON(Response{Link: MakeLink(article),
		QrCode: fmt.Sprintf("%s/static/%s", os.Getenv("HOST_URL"), path)})
}

// @Summary Get article estimated read time
// @Success 200 {object} int
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
