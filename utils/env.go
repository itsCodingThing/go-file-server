package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	DATABASE_URL      string
	HOST              string
	DATABASE_USER     string
	DATABASE_PASSWORD string
	USERNAME          string
	PASSWORD          string
	STORAGE           string
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}
}
