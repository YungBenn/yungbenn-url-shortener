package internal

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *fiber.Ctx) error {
	var creationRequest UrlCreationRequest
	if err := c.BodyParser(&creationRequest); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	shortUrl := GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)
	SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)

	host := "http://localhost:9808/"
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "short url created successfully",
		"short_url": host + shortUrl,
	})
}

func HandleShortUrlRedirect(c *fiber.Ctx) error {
	shortUrl := c.Params("shortUrl")
	initialUrl := RetrieveInitialUrl(shortUrl)
	return c.Redirect(initialUrl, http.StatusFound)
}