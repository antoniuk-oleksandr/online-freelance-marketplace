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
