package helpers

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"ofm_backend/cmd/ofm_backend/api/file/model"
	"ofm_backend/cmd/ofm_backend/utils"
	"os"
)

func MakeFileData(files []*multipart.FileHeader) []model.FileData {
	fileData := make([]model.FileData, len(files))
	for i, file := range files {
		fileData[i] = model.FileData{
			Name: file.Filename,
		}
	}

	return fileData
}

func WriteMultipartFiles(files []*multipart.FileHeader, writer *multipart.Writer) error {
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			return err
		}
		defer file.Close()

		formFile, err := writer.CreateFormFile("files", fileHeader.Filename)
		if err != nil {
			return err
		}

		_, err = io.Copy(formFile, file)
		if err != nil {
			return err
		}
	}

	return writer.Close()
}

func MakeFileServiceRequest(
	buf bytes.Buffer, writer *multipart.Writer,
) (*http.Response, error) {
	fileServerURL, err := getFileServerURL()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, fileServerURL, &buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return resp, nil
}

func getFileServerURL() (string, error) {
	fileHost := os.Getenv("FILE_SERVER_HOST")
	if fileHost == "" {
		return "", utils.ErrEnvVarNotSet
	}

	return fmt.Sprintf("%s/upload", fileHost), nil
}
