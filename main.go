package main

import (
	"fmt"
	"net/http"

	"github.com/YungBenn/yungbenn-url-shortener/internal"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": "Welcome to YungBenn URL Shortener API",
		})
	})

	app.Post("/create-short-url", internal.CreateShortUrl)
	app.Get("/:shortUrl", internal.HandleShortUrlRedirect)

	internal.InitStore()

	err := app.Listen(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}