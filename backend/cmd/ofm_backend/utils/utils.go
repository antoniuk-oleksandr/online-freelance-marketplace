package utils

import "github.com/joho/godotenv"

func LoadEnvValues(){
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}