package utils

import (
	"github.com/joho/godotenv"
)

func LoadEnvValues() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

func Contains(slice []string, element string) bool {
	for _, item := range slice {
		if item == element {
			return true
		}
	}
	return false
}