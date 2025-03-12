package service

import "mime/multipart"

type FileService interface {
	DeleteFile(fileName string) (int, error)
	UploadFiles(files []*multipart.FileHeader) (int, error)
	SaveFile(file *multipart.FileHeader, filePath string) error
	ConvertToBytes(file *multipart.FileHeader) ([]byte, error)
}
