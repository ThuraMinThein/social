package env

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func GetString(key , fallback string) string {

	loadEnv()

	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}

func GetInt(key string, fallback int) int {
	
	loadEnv()

	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return intValue
}

func loadEnv() {
	 err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }
}