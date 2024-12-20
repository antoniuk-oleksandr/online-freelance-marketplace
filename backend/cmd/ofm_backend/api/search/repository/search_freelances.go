package repository

import (
	"ofm_backend/cmd/ofm_backend/api/search/model"

	"github.com/jmoiron/sqlx"
)

func SearchFreelances(db *sqlx.DB, query string) ([]model.SearchService, error) {
	var services []model.SearchService

	if err := db.Select(&services, query); err != nil {
		return nil, err
	}

	return services, nil
}
