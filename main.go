package main

import (
	"log"
	"os"
	"strings"

	"web_practicum/database"
	"web_practicum/routing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/mattn/go-sqlite3"

	_ "web_practicum/docs"

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
	routing.Setup(app)
	log.Fatal(app.Listen(os.Getenv("LISTEN_ADDR")))
}

// func postImage(ctx *fiber.Ctx) error {
// 	return nil
// }
