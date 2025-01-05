package repository

import (
	"encoding/json"
	"ofm_backend/cmd/ofm_backend/api/home_data/model"
	"ofm_backend/cmd/ofm_backend/api/home_data/utils"

	"github.com/jmoiron/sqlx"
)

type homeRepository struct {
	db *sqlx.DB
}

func NewHomeRepository(db *sqlx.DB) HomeRepository {
	return &homeRepository{
		db: db,
	}
}

func (homeRepository *homeRepository) GetHomeData() (*model.HomeData, error) {
	var homeData model.HomeData

	var homeDataJSON []byte

	rows, err := homeRepository.db.Queryx(utils.GetHomeDataQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&homeDataJSON); err != nil {
			return nil, err
		}

		if err := json.Unmarshal(homeDataJSON, &homeData); err != nil {
			return nil, err
		}
	}

	return &homeData, nil
}
