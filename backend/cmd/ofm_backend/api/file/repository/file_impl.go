package repository

import (
	"log"
	"ofm_backend/cmd/ofm_backend/api/file/model"
	"ofm_backend/cmd/ofm_backend/api/file/queries"

	"github.com/jmoiron/sqlx"
)

type fileRepository struct {
	db *sqlx.DB
}

func NewFileRepository(db *sqlx.DB) FileRepository {
	return &fileRepository{
		db: db,
	}
}

func (fileRepository *fileRepository) AddFiles(
	fileData []model.FileData,
) ([]int, error) {
	rows, err := fileRepository.db.NamedQuery(queries.AddFilesQuery, fileData)
	if err != nil {
		log.Println("error in AddFiles", err)
		return nil, err
	}
	defer rows.Close()

	fileIds := make([]int, 0)
	for rows.Next() {
		var fileId int
		if err := rows.Scan(&fileId); err != nil {
			return nil, err
		}

		fileIds = append(fileIds, fileId)
	}

	return fileIds, nil
}

func (fileRepository *fileRepository) DeleteFile(fileId int) error {
	return nil
}

func (fileRepository *fileRepository) StartTransaction() (*sqlx.Tx, error) {
	return fileRepository.db.Beginx()
}

func (fileRepository *fileRepository) CommitTransaction(tx *sqlx.Tx) error {
	return tx.Commit()
}

func (fileRepository *fileRepository) RollbackTransaction(tx *sqlx.Tx) error {
	return tx.Rollback()
}
