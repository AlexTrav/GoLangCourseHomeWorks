package database

import (
	"database/sql"

	"github.com/AlexTrav/GoLangCourseHomeWorks/Projects/BooksAPI/pkg/logger"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQL(dsn string) *sql.DB {
	logger.Log.Println("initializing MySQL connection")

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logger.Log.Fatalf("mysql open error: %v", err)
	}
	if err := db.Ping(); err != nil {
		logger.Log.Fatalf("mysql ping error: %v", err)
	}
	logger.Log.Println("MySQL connection established")

	return db
}
