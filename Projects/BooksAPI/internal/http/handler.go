package http

import (
	stdErrors "errors"
	"strconv"

	"github.com/AlexTrav/GoLangCourseHomeWorks/Projects/BooksAPI/internal/domain"
	appErrors "github.com/AlexTrav/GoLangCourseHomeWorks/Projects/BooksAPI/internal/errors"
	"github.com/AlexTrav/GoLangCourseHomeWorks/Projects/BooksAPI/internal/http/dto"
	"github.com/AlexTrav/GoLangCourseHomeWorks/Projects/BooksAPI/internal/service"
	"github.com/gofiber/fiber/v2"
)

type BookHandler struct {
	service service.BookService
}

func NewBookHandler(s service.BookService) *BookHandler {
	return &BookHandler{service: s}
}

// Create godoc
// @Summary Create book
// @Description Create a new book
// @Tags Books
// @Accept json
// @Produce json
// @Param book body dto.CreateBookRequest true "Book"
// @Success 201 {object} dto.BookResponse
// @Failure 400 {object} dto.ValidationErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /books [post]
func (h *BookHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateBookRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}
	book := domain.Book{
		Title:  req.Title,
		Author: req.Author,
		ISBN:   req.ISBN,
	}

	err := h.service.Create(&book)
	if err != nil {
		if stdErrors.Is(err, appErrors.ErrValidation) {
			return c.Status(400).JSON(fiber.Map{"error": "validation"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "internal"})
	}

	return c.Status(201).JSON(book)
}

// Get godoc
// @Summary Get book
// @Description Get book by ID
// @Tags Books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} dto.BookResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /books/{id} [get]
func (h *BookHandler) Get(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	book, err := h.service.Get(id)
	if err != nil {
		if stdErrors.Is(err, appErrors.ErrNotFound) {
			return c.Status(404).JSON(fiber.Map{"error": "not_found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "internal"})
	}

	return c.JSON(book)
}

// Update godoc
// @Summary Update book
// @Description Update existing book
// @Tags Books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body dto.UpdateBookRequest true "Book"
// @Success 200 {object} dto.BookResponse
// @Failure 400 {object} dto.ValidationErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /books/{id} [put]
func (h *BookHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	var req dto.UpdateBookRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	book := domain.Book{
		ID:     id,
		Title:  req.Title,
		Author: req.Author,
		ISBN:   req.ISBN,
	}

	err = h.service.Update(&book)
	if err != nil {
		if stdErrors.Is(err, appErrors.ErrValidation) {
			return c.Status(400).JSON(fiber.Map{"error": "validation"})
		}
		if stdErrors.Is(err, appErrors.ErrNotFound) {
			return c.Status(404).JSON(fiber.Map{"error": "not_found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "internal"})
	}

	return c.JSON(book)
}

// Delete godoc
// @Summary Delete book
// @Description Delete book by ID
// @Tags Books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /books/{id} [delete]
func (h *BookHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	err = h.service.Delete(id)
	if err != nil {
		if stdErrors.Is(err, appErrors.ErrNotFound) {
			return c.Status(404).JSON(fiber.Map{"error": "not_found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "internal"})
	}

	return c.JSON(fiber.Map{"status": "ok"})
}
