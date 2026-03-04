package service

import (
	"strings"

	"github.com/AlexTrav/GoLangCourseHomeWorks/Projects/BooksAPI/internal/domain"
	appErrors "github.com/AlexTrav/GoLangCourseHomeWorks/Projects/BooksAPI/internal/errors"
	"github.com/AlexTrav/GoLangCourseHomeWorks/Projects/BooksAPI/internal/repository"
	"github.com/AlexTrav/GoLangCourseHomeWorks/Projects/BooksAPI/pkg/logger"
)

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(r repository.BookRepository) BookService {
	return &bookService{repo: r}
}

func (s *bookService) Create(book *domain.Book) error {
	logger.Log.Printf("creating book: title=%s author=%s\n", book.Title, book.Author)

	if strings.TrimSpace(book.Title) == "" ||
		strings.TrimSpace(book.Author) == "" ||
		strings.TrimSpace(book.ISBN) == "" {

		logger.Log.Println("validation error while creating book")
		return appErrors.ErrValidation
	}

	return s.repo.Create(book)
}

func (s *bookService) Get(id int) (*domain.Book, error) {
	logger.Log.Printf("getting book id=%d\n", id)
	return s.repo.GetByID(id)
}
func (s *bookService) Update(book *domain.Book) error {
	logger.Log.Printf("updating book id=%d\n", book.ID)

	if strings.TrimSpace(book.Title) == "" ||
		strings.TrimSpace(book.Author) == "" ||
		strings.TrimSpace(book.ISBN) == "" {
		logger.Log.Println("validation error while updating book")
		return appErrors.ErrValidation
	}

	return s.repo.Update(book)
}

func (s *bookService) Delete(id int) error {
	logger.Log.Printf("deleting book id=%d\n", id)
	return s.repo.Delete(id)
}
