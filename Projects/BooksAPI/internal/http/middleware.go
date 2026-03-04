package http

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ValidateBook() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body map[string]string
		if err := c.BodyParser(&body); err != nil {
			return fiber.ErrBadRequest
		}

		errors := make(map[string]string)

		if strings.TrimSpace(body["title"]) == "" {
			errors["title"] = "required"
		}
		if strings.TrimSpace(body["author"]) == "" {
			errors["author"] = "required"
		}
		if strings.TrimSpace(body["isbn"]) == "" {
			errors["isbn"] = "required"
		}

		if len(errors) > 0 {
			return c.Status(400).JSON(fiber.Map{
				"error":  "validation",
				"fields": errors,
			})
		}

		return c.Next()
	}
}
