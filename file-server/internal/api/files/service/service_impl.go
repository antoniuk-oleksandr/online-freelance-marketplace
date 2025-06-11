package service

import (
	"bytes"
	"file-server/internal/utils"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

type fileService struct {
	uploadDir string
}

func NewFileService(uploadDir string) FileService {
	return &fileService{
		uploadDir: uploadDir,
	}
}

func (fs fileService) DeleteFile(fileName string) (int, error) {
	path := filepath.Join(fs.uploadDir, fileName)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fiber.StatusNotFound, utils.ErrFileNotFound
	}

	err := os.Remove(path)
	if err != nil {
		return fiber.StatusInternalServerError, err
	}

	return fiber.StatusOK, nil
}

func (fs fileService) UploadFiles(files []*multipart.FileHeader) (int, error) {
	for _, file := range files {
		filePath := filepath.Join(fs.uploadDir, file.Filename)
		err := fs.SaveFile(file, filePath)
		if err != nil {
			return fiber.StatusInternalServerError, err
		}
	}

	return fiber.StatusCreated, nil
}

func (fs fileService) SaveFile(file *multipart.FileHeader, filePath string) error {
	bytes, err := fs.ConvertToBytes(file)
	if err != nil {
		return err
	}

	if err := os.WriteFile(filePath, bytes, 0666); err != nil {
		return utils.ErrFailedToSaveFile
	}
	
	return nil
}

func (fs fileService) ConvertToBytes(file *multipart.FileHeader) ([]byte, error) {
	src, err := file.Open()
	if err != nil {
		return nil, utils.ErrFailedToOpenFile
	}
	defer src.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, src); err != nil {
		return nil, utils.ErrFailedToCopyFile
	}

	return buf.Bytes(), nil
}
