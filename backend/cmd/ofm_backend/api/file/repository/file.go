package repository

import (
	"ofm_backend/cmd/ofm_backend/api/file/model"

	"github.com/jmoiron/sqlx"
)

type FileRepository interface {
	AddFiles(tx *sqlx.Tx, fileData []model.FileData) ([]int, error)
	DeleteFile(fileId int) error
	StartTransaction() (*sqlx.Tx, error)
	RollbackTransaction(tx *sqlx.Tx) error
	CommitTransaction(tx *sqlx.Tx) error
}
