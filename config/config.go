package config

import (
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	Addr    string
	DB_HOST string
	DB_PORT string
	DB_NAME string
	DB_USER string
	DB_PASS string
}

var Get *config

func LoadConfig(path string) error {
	if err := godotenv.Load(path); err != nil {
		return err
	}

	Get = &config{
		Addr:    os.Getenv("APP_PORT"),
		DB_HOST: os.Getenv("DB_HOST"),
		DB_PORT: os.Getenv("DB_PORT"),
		DB_NAME: os.Getenv("DB_NAME"),
		DB_USER: os.Getenv("DB_USER"),
		DB_PASS: os.Getenv("DB_PASSWORD"),
	}
	return nil

}
