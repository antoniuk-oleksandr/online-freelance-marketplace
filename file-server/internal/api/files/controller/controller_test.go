package controller

import (
	"bytes"
	"file-server/internal/utils"
	"fmt"
	"mime/multipart"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) ConvertToBytes(file *multipart.FileHeader) ([]byte, error) {
	args := m.Called(file)
	if args.Get(0) != nil {
		return args.Get(0).([]byte), args.Error(1)
	}

	return nil, args.Error(1)
}

func (m *MockService) SaveFile(file *multipart.FileHeader, filePath string) error {
	args := m.Called(file, filePath)
	return args.Error(0)
}

func (m *MockService) DeleteFile(fileName string) (int, error) {
	args := m.Called(fileName)
	return args.Int(0), args.Error(1)
}

func (m *MockService) UploadFiles(files []*multipart.FileHeader) (int, error) {
	args := m.Called(files)
	return args.Int(0), args.Error(1)
}

func TestUploadFile(t *testing.T) {
	tests := []struct {
		name           string
		expectedStatus int
		expectedErr    error
		addFiles       bool
	}{
		{
			name:           "Successful file upload",
			expectedStatus: fiber.StatusCreated,
			expectedErr:    nil,
			addFiles:       true,
		},
		{
			name:           "No file upload",
			expectedStatus: fiber.StatusBadRequest,
			expectedErr:    utils.ErrNoFileUploaded,
			addFiles:       false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockService := new(MockService)
			mockService.On("UploadFiles", mock.AnythingOfType("[]*multipart.FileHeader")).
				Return(tc.expectedStatus, tc.expectedErr)
			mockService.On("ConvertToBytes", mock.Anything).Return([]byte{}, nil)
			mockService.On("SaveFile", mock.Anything, mock.Anything).Return(nil)

			fc := NewFileController(mockService)

			app := fiber.New()
			app.Post("/upload", fc.UploadFiles)

			buf, contentType, err := createBody(tc.addFiles)
			assert.NoError(t, err)

			req, _ := http.NewRequest("POST", "/upload", buf)
			req.Header.Set("Content-Type", contentType)

			resp, err := app.Test(req, -1)
			assert.NoError(t, err, "Error should be nil")
			assert.Equal(t, tc.expectedStatus, resp.StatusCode, "Status code should be equal")
		})
	}
}

func TestDeleteFile(t *testing.T) {
	tests := []struct {
		name            string
		extectedStatus  int
		expectedErr     error
		actualFileName  string
		requestFileName string
	}{
		{
			name:            "Successful file deletion",
			extectedStatus:  fiber.StatusOK,
			expectedErr:     nil,
			actualFileName:  "test.txt",
			requestFileName: "test.txt",
		},
		{
			name:            "Not found file",
			extectedStatus:  fiber.StatusNotFound,
			expectedErr:     utils.ErrFileNotFound,
			actualFileName:  "test.txt",
			requestFileName: "req.txt",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockService := new(MockService)
			mockService.On("DeleteFile", mock.AnythingOfType("string")).Return(tc.extectedStatus, tc.expectedErr)

			fc := NewFileController(mockService)

			app := fiber.New()
			app.Delete("/delete/:filename", fc.DeleteFile)

			req, _ := http.NewRequest("DELETE", fmt.Sprintf("/delete/%s", tc.requestFileName), nil)

			resp, err := app.Test(req, -1)
			assert.NoError(t, err, "Error should be nil")
			assert.Equal(t, tc.extectedStatus, resp.StatusCode, "Status code should be equal")
		})
	}
}

func createBody(addFiles bool) (*bytes.Buffer, string, error) {
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)

	if addFiles {
		_, err := w.CreateFormFile("files", "test.txt")
		if err != nil {
			return nil, "", err
		}
	}

	contentType := w.FormDataContentType()
	w.Close()

	return &buf, contentType, nil
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
