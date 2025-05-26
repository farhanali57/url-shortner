package routes

import (
	"example.com/ShortUrl/models"
	"example.com/ShortUrl/service"
	"github.com/gofiber/fiber/v2"
)

func GetShortUrl(c *fiber.Ctx) error {
	var url models.Link
	err := c.BodyParser(&url)
	if err != nil {

		//stringBody := string(c.Body())
		//if stringBody != "" {
		//	url.LongUrl = stringBody
		//} else {
		return c.Status(400).JSON(fiber.Map{"invalid input": err.Error()})
		//}
	}

	shortenUrl := service.GenerateShortUrl(url.Url)
	url.ShortUrl = shortenUrl

	err = url.Save()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"could not save url": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"shorturl": url.ShortUrl})
}

func RedirectUrl(c *fiber.Ctx) error {
	shortUrl := c.Params("shorturl")
	if shortUrl == "" {
		return c.Status(400).JSON(fiber.Map{"error": "shortUrl is empty"})
	}
	url, err := models.GetURL(shortUrl)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "url not found",
			"details": err.Error()})
	}
	return c.Redirect(url, fiber.StatusFound)
}
