package configurations

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_NAME     string
	DB_USER     string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	APP_PORT    int64
}

var GetEnv = initConfig()

func initConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Env dosyası bulunamadı!")
	}

	APP_PORT, err := strconv.ParseInt(os.Getenv("APP_PORT"), 10, 64)
	if err != nil {
		APP_PORT = 3000
	}

	return Config{
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_HOST:     os.Getenv("DB_HOST"),
		APP_PORT:    APP_PORT,
	}
}
