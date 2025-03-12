package utils

import "os"

func Plural(n int) string {
	if n == 1 {
		return ""
	}
	return "s"
}

func CreateDirIfDoesNotExist(uploadDir string) {
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, os.ModePerm)
	}
}
