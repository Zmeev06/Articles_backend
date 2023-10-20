package main

import (
	// "log"

	// "fmt"
	"log"
	"web_practicum/database"
	// . "web_practicum/models"

	// _ "github.com/lib/pq"
	"github.com/gofiber/fiber/v2"
	// "github.com/prometheus/common/log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	if err := database.Setup(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()
	// app.Post("api/images")
	app.Static("api/images", "./images")
	// app.Post("/api/create/article")
	app.Get("/api/article/:id<int>", getArticle)
	app.Get("/api/shit", func(c *fiber.Ctx) error {
		c.Response().SetBodyString("hello")
		return nil
	})
	// app.Get("/api/articles", getAllArticles)
	// app.Get("/api/articles/:user")
	// app.Get("/api/user/articles")
	log.Fatal(app.Listen(":8080"))
}
func getArticle(ctx *fiber.Ctx) error {

	db := database.DB
	obj, succ := db.Get(ctx.Params("id"))
	if !succ {
		return fiber.ErrNotFound
	}
	return ctx.JSON(obj)
}

// func getAllArticles(ctx *fiber.Ctx) error {
//
// 	db := database.DB
// 	var articles []Article
// 	obj := db.Find(articles)
// 	if obj.Error != nil {
// 		return fiber.ErrBadRequest
// 	}
// 	return ctx.JSON(obj)
// }


