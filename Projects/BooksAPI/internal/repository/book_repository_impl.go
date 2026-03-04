package repository

import (
	"database/sql"
	stdErrors "errors"

	"github.com/AlexTrav/GoLangCourseHomeWorks/Projects/BooksAPI/internal/domain"
	appErrors "github.com/AlexTrav/GoLangCourseHomeWorks/Projects/BooksAPI/internal/errors"
	"github.com/AlexTrav/GoLangCourseHomeWorks/Projects/BooksAPI/pkg/logger"
)

type bookMySQL struct {
	db *sql.DB
}

func NewBookMySQL(db *sql.DB) BookRepository {
	return &bookMySQL{db: db}
}

func (r *bookMySQL) Create(book *domain.Book) error {

	logger.Log.Printf("repository: creating book title=%s author=%s\n", book.Title, book.Author)

	res, err := r.db.Exec(
		"INSERT INTO books(title, author, isbn) VALUES (?, ?, ?)",
		book.Title,
		book.Author,
		book.ISBN,
	)

	if err != nil {
		logger.Log.Printf("repository: db error creating book: %v\n", err)
		return appErrors.ErrInternal
	}

	id, err := res.LastInsertId()
	if err != nil {
		logger.Log.Printf("repository: error getting last insert id: %v\n", err)
		return appErrors.ErrInternal
	}

	book.ID = int(id)

	logger.Log.Printf("repository: book created id=%d\n", book.ID)

	return nil
}

func (r *bookMySQL) GetByID(id int) (*domain.Book, error) {

	logger.Log.Printf("repository: getting book id=%d\n", id)

	book := &domain.Book{}

	err := r.db.QueryRow(
		"SELECT id, title, author, isbn FROM books WHERE id = ?",
		id,
	).Scan(
		&book.ID,
		&book.Title,
		&book.Author,
		&book.ISBN,
	)

	if err != nil {

		if stdErrors.Is(err, sql.ErrNoRows) {
			logger.Log.Printf("repository: book not found id=%d\n", id)
			return nil, appErrors.ErrNotFound
		}

		logger.Log.Printf("repository: db error getting book id=%d error=%v\n", id, err)

		return nil, appErrors.ErrInternal
	}

	return book, nil
}

func (r *bookMySQL) Update(book *domain.Book) error {

	logger.Log.Printf("repository: updating book id=%d\n", book.ID)

	res, err := r.db.Exec(
		"UPDATE books SET title=?, author=?, isbn=? WHERE id=?",
		book.Title,
		book.Author,
		book.ISBN,
		book.ID,
	)

	if err != nil {
		logger.Log.Printf("repository: db error updating book id=%d error=%v\n", book.ID, err)
		return appErrors.ErrInternal
	}

	rows, err := res.RowsAffected()
	if err != nil {
		logger.Log.Printf("repository: rows affected error id=%d error=%v\n", book.ID, err)
		return appErrors.ErrInternal
	}

	if rows == 0 {
		logger.Log.Printf("repository: update failed book not found id=%d\n", book.ID)
		return appErrors.ErrNotFound
	}

	logger.Log.Printf("repository: book updated id=%d\n", book.ID)

	return nil
}

func (r *bookMySQL) Delete(id int) error {

	logger.Log.Printf("repository: deleting book id=%d\n", id)

	res, err := r.db.Exec(
		"DELETE FROM books WHERE id=?",
		id,
	)

	if err != nil {
		logger.Log.Printf("repository: db error deleting book id=%d error=%v\n", id, err)
		return appErrors.ErrInternal
	}

	rows, err := res.RowsAffected()
	if err != nil {
		logger.Log.Printf("repository: rows affected error deleting id=%d error=%v\n", id, err)
		return appErrors.ErrInternal
	}

	if rows == 0 {
		logger.Log.Printf("repository: delete failed book not found id=%d\n", id)
		return appErrors.ErrNotFound
	}

	logger.Log.Printf("repository: book deleted id=%d\n", id)

	return nil
}
