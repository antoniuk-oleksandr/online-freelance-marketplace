package repository

import (
	"ofm_backend/cmd/ofm_backend/api/file/model"

	"github.com/jmoiron/sqlx"
)

func GetFileById(id int, db *sqlx.DB) (*model.File, error) {
	var file model.File

	query := "SELECT F.id, F.name FROM files F WHERE id = $1"

	if err := db.Get(&file, query, id); err != nil {
		return nil, err
	}

	return &file, nil
}

func AddFile(name string, db *sqlx.DB) (int, error) {
	var fileID int
	
	query := `INSERT INTO files (name) VALUES ($1) RETURNING id`
	
	err := db.QueryRow(query, name).Scan(&fileID)
	if err != nil {
		return 0, err
	}
	
	return fileID, nil
}
