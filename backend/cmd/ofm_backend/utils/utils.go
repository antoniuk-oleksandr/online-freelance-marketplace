package utils

import (
	"time"

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

func GetCurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05.999999999")
}

func ConvertTimeToSting(timeData time.Time) string {
	return timeData.Format("2006-01-02 15:04:05.999999999")
	
} 