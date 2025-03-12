package service

import (
	"bytes"
	"file-server/internal/utils"
	"mime/multipart"
	"os"
	"path/filepath"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type mockFileService struct {
	uploadDir string
	saveFile  func(file *multipart.FileHeader, path string) error
}

func TestUploadFiles(t *testing.T) {
	tests := []struct {
		name           string
		expectedStatus int
		expectedErr    error
		files          []*multipart.FileHeader
	}{
		{
			name:           "Successful file upload",
			expectedStatus: fiber.StatusCreated,
			expectedErr:    nil,
			files:          []*multipart.FileHeader{createFileHeader("file.txt", "Hello, World!")},
		},
		{
			name:           "Unuccessful file upload",
			expectedStatus: fiber.StatusInternalServerError,
			expectedErr:    utils.ErrFailedToOpenFile,
			files:          []*multipart.FileHeader{&multipart.FileHeader{}},
		},
	}

	tempDir := t.TempDir()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fs := NewFileService(tempDir)

			actualStatus, actualErr := fs.UploadFiles(tc.files)
			assert.Equal(t, tc.expectedErr, actualErr, "Error should be equal")
			assert.Equal(t, tc.expectedStatus, actualStatus, "Error should be equal")
		})
	}

}

func TestDeleteFile(t *testing.T) {
	tempDir := t.TempDir()

	tests := []struct {
		name           string
		success        bool
		expectedErr    error
		expectedStatus int
		fileName       string
	}{
		{
			name:           "Successful file deletion",
			success:        true,
			expectedErr:    nil,
			fileName:       "file.txt",
			expectedStatus: fiber.StatusOK,
		},
		{
			name:           "No file to delete",
			success:        false,
			expectedErr:    utils.ErrFileNotFound,
			fileName:       "file.txt",
			expectedStatus: fiber.StatusNotFound,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fs := NewFileService(tempDir)

			if tc.success {
				saveTempFile(tc.fileName, "Hello, World!", tempDir)
			}
			actualStatus, actualErr := fs.DeleteFile(tc.fileName)
			assert.Equal(t, tc.expectedErr, actualErr, "Error should be equal")
			assert.Equal(t, tc.expectedStatus, actualStatus, "Status should be equal")
		})
	}
}

func TestConvertToBytes(t *testing.T) {
	tempDir := t.TempDir()

	tests := []struct {
		name        string
		expectedErr error
		file        *multipart.FileHeader
	}{
		{
			name:        "Successful conversion",
			expectedErr: nil,
			file:        createFileHeader("test.txt", "Hello, World!"),
		},
		{
			name:        "Unsuccessful conversion",
			expectedErr: utils.ErrFailedToOpenFile,
			file:        &multipart.FileHeader{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fs := NewFileService(tempDir)

			_, actualErr := fs.ConvertToBytes(tc.file)
			assert.Equal(t, tc.expectedErr, actualErr, "Error should be equal")
		})
	}
}

func TestSaveFile(t *testing.T) {
	tempDir := t.TempDir()

	tests := []struct {
		name        string
		expectedErr error
		file        *multipart.FileHeader
	}{
		{
			name:        "Successful file save",
			expectedErr: nil,
			file:        createFileHeader("test.txt", "Hello, World!"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fs := NewFileService(tempDir)

			filePath := filepath.Join(tempDir, "test.txt")
			actualErr := fs.SaveFile(tc.file, filePath)
			assert.Equal(t, tc.expectedErr, actualErr, "Error should be equal")
		})
	}
}

func saveTempFile(fileName string, content string, tempDir string) {
	path := filepath.Join(tempDir, fileName)
	os.WriteFile(path, []byte(content), 0666)
}

func createFileHeader(fileName string, content string) *multipart.FileHeader {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("fileField", fileName)
	if err != nil {
		panic(err)
	}

	_, err = part.Write([]byte(content))
	if err != nil {
		panic(err)
	}
	writer.Close()

	reader := multipart.NewReader(body, writer.Boundary())
	form, err := reader.ReadForm(0)
	if err != nil {
		panic(err)
	}

	return form.File["fileField"][0]
}
