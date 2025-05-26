package main

import (
	"example.com/ShortUrl/db"
	"example.com/ShortUrl/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	db.ConnectDB()

	app := fiber.New()

	app.Post("/shorturl", routes.GetShortUrl)
	app.Get("/:redirect", routes.RedirectUrl)

	log.Fatal(app.Listen(":3000"))

}
