package database

import (
	"database/sql"
	"fmt"

	"github.com/AlexTrav/GoLangCourseHomeWorks/Projects/BooksAPI/pkg/logger"

	"github.com/golang-migrate/migrate/v4"
	mysqlDriver "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(db *sql.DB, migrationsPath string) {

	logger.Log.Println("running database migrations")

	driver, err := mysqlDriver.WithInstance(db, &mysqlDriver.Config{})
	if err != nil {
		logger.Log.Fatalf("migration driver error: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", migrationsPath),
		"mysql",
		driver,
	)

	if err != nil {
		logger.Log.Fatalf("migration init error: %v", err)
	}

	if err := m.Up(); err != nil {

		if err == migrate.ErrNoChange {
			logger.Log.Println("no new migrations")
			return
		}

		logger.Log.Fatalf("migration failed: %v", err)
	}

	logger.Log.Println("migrations applied successfully")
}
