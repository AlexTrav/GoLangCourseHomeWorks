package config

import (
	"os"

	"github.com/AlexTrav/GoLangCourseHomeWorks/Projects/BooksAPI/pkg/logger"
	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string
	DBDsn   string
}

func Load() *Config {

	logger.Log.Println("loading configuration")

	err := godotenv.Load()
	if err != nil {
		logger.Log.Println(".env file not found, using environment variables")
	}

	cfg := &Config{
		AppPort: os.Getenv("APP_PORT"),
		DBDsn: os.Getenv("DB_USER") + ":" +
			os.Getenv("DB_PASSWORD") + "@tcp(" +
			os.Getenv("DB_HOST") + ":" +
			os.Getenv("DB_PORT") + ")/" +
			os.Getenv("DB_NAME") + "?parseTime=true",
	}

	logger.Log.Printf("configuration loaded (port=%s)\n", cfg.AppPort)

	return cfg
}
