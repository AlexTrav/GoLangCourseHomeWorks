package http

import (
	"github.com/AlexTrav/GoLangCourseHomeWorks/Projects/BooksAPI/internal/service"

	_ "github.com/AlexTrav/GoLangCourseHomeWorks/Projects/BooksAPI/docs"
	"github.com/gofiber/fiber/v2"
	swagger "github.com/gofiber/swagger"
)

func SetupRoutes(app *fiber.App, svc service.BookService) {
	h := NewBookHandler(svc)

	app.Get("/swagger/*", swagger.HandlerDefault)

	api := app.Group("/api")
	v1 := api.Group("/v1")

	books := v1.Group("/books")

	books.Post("/", ValidateBook(), h.Create)
	books.Get("/:id", h.Get)
	books.Put("/:id", ValidateBook(), h.Update)
	books.Delete("/:id", h.Delete)
}
