package libs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_URL     string
	JWT_SECRET string
}

func LoadConfig() *Config {
	config := &Config{}

	config.DB_URL = os.Getenv("DB_URL")
	config.JWT_SECRET = os.Getenv("JWT_SECRET")

	if config.DB_URL == "" || config.JWT_SECRET == "" {
		if err := godotenv.Load(); err != nil {
			panic("Error loading .env file")
		}
		config.DB_URL = os.Getenv("DB_URL")
		config.JWT_SECRET = os.Getenv("JWT_SECRET")
	}

	return config
}
