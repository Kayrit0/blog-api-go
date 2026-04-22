package libs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_URL string
}

func LoadConfig() *Config {
	config := &Config{}

	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	config.DB_URL = os.Getenv("DB_URL")

	return config
}
