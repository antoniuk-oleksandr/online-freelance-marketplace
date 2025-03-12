package repository

import (
	"ofm_backend/cmd/ofm_backend/api/filter_params/model"

	"github.com/jmoiron/sqlx"
)

func FilterParamsGetAll(db *sqlx.DB) (*model.FilterParamsJSON, error) {
	var rawFilterParams model.FilterParamsJSON

	query := `SELECT
			(SELECT json_agg(L) FROM languages L)  as languages,
			(SELECT json_agg(C) FROM categories C) as categories,
			(SELECT json_agg(S) FROM skills S)     as skills`

	row := db.QueryRowx(query)
	if err := row.Scan(
		&rawFilterParams.Languages, 
		&rawFilterParams.Categories, 
		&rawFilterParams.Skills,
	); err != nil {
		return nil, err
	}
	
	return &rawFilterParams, nil
}
