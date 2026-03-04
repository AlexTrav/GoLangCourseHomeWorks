// @title Books API
// @version 1.0
// @description REST API for managing books
// @host localhost:3000
// @BasePath /api/v1

package main

import (
	"github.com/AlexTrav/GoLangCourseHomeWorks/Projects/BooksAPI/internal/config"
	"github.com/AlexTrav/GoLangCourseHomeWorks/Projects/BooksAPI/internal/database"
	"github.com/AlexTrav/GoLangCourseHomeWorks/Projects/BooksAPI/internal/http"
	"github.com/AlexTrav/GoLangCourseHomeWorks/Projects/BooksAPI/internal/repository"
	"github.com/AlexTrav/GoLangCourseHomeWorks/Projects/BooksAPI/internal/service"
	"github.com/AlexTrav/GoLangCourseHomeWorks/Projects/BooksAPI/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	logger.Init()
	logger.Log.Println("starting Books API")

	cfg := config.Load()
	logger.Log.Println("config loaded")

	db := database.NewMySQL(cfg.DBDsn)
	logger.Log.Println("database connected")

	database.RunMigrations(db, "migrations")

	repo := repository.NewBookMySQL(db)
	svc := service.NewBookService(repo)

	app := fiber.New()

	app.Use(fiberLogger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))

	http.SetupRoutes(app, svc)

	logger.Log.Printf("server started on port %s\n", cfg.AppPort)

	app.Listen(":" + cfg.AppPort)
}
