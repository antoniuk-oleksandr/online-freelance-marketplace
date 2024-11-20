package filereader

import (
	"os"
	"path/filepath"
)

func GetHTMLTempalate(fileName string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	htmlFilePath := filepath.Join(wd, "internal", "mailer", "templates", fileName)
	
	htmlContent, err := os.ReadFile(htmlFilePath)
	if err != nil {
		return "", err
	}

	return string(htmlContent), nil
}
