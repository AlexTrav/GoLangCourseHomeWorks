package repository

import "github.com/AlexTrav/GoLangCourseHomeWorks/Projects/BooksAPI/internal/domain"

type BookRepository interface {
	Create(book *domain.Book) error
	GetByID(id int) (*domain.Book, error)
	Update(book *domain.Book) error
	Delete(id int) error
}
