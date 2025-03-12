package service

import (
	"mime/multipart"
)

type FileService interface {
	UploadFiles(files []*multipart.FileHeader) error
	SaveFilesMetaData(files []*multipart.FileHeader) ([]int, error)
	DeleteFile(fileId int) error
}
