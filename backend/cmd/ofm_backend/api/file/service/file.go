package service

import (
	"mime/multipart"

	"github.com/jmoiron/sqlx"
)

type FileService interface {
	UploadFiles(files []*multipart.FileHeader) error
	SaveFilesMetaData(tx *sqlx.Tx, files []*multipart.FileHeader) ([]int, error)
	DeleteFile(fileId int) error
	UploadFromURLWithoutTransaction(tx *sqlx.Tx, picURL string) (int, string, error)
}
