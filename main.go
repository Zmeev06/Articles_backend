package main

import (
	"log"
	"os"

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
	app.Use(cors.New())
	routing.Setup(app)
	log.Fatal(app.Listen(os.Getenv("LISTEN_ADDR")))
}

// func postImage(ctx *fiber.Ctx) error {
// 	return nil
// }
