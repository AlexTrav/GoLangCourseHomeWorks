package service

import "github.com/AlexTrav/GoLangCourseHomeWorks/Projects/BooksAPI/internal/domain"

type BookService interface {
	Create(book *domain.Book) error
	Get(id int) (*domain.Book, error)
	Update(book *domain.Book) error
	Delete(id int) error
}
