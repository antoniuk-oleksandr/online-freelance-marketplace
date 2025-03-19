package repository

import (
	"encoding/json"
	"ofm_backend/cmd/ofm_backend/api/my_profile/dto"
	"ofm_backend/cmd/ofm_backend/api/my_profile/model"
	"ofm_backend/cmd/ofm_backend/api/my_profile/queries"

	"github.com/jmoiron/sqlx"
)

type myProfileRepository struct {
	db *sqlx.DB
}

func NewMyProfileRepository(db *sqlx.DB) MyProfileRepository {
	return &myProfileRepository{db: db}
}

func (mpr *myProfileRepository) GetMyProfileOrders(
	params *dto.OrdersPaginationParams,
) (*model.OrdersData, error) {
	rows, err := mpr.db.Queryx(
		queries.GetMyProfileOrdersQuery, params.UserId, params.Offset, params.OrdersPerPage,
	)
	if err != nil {
		return nil, err
	}

	return mpr.ParseOrdersDataFromRows(rows)
}

func (mpr *myProfileRepository) ParseOrdersDataFromRows(rows *sqlx.Rows) (*model.OrdersData, error) {
	var ordersData model.OrdersData
	var orderTableDataJSON []byte

	if rows.Next() {
		err := rows.Scan(&orderTableDataJSON, &ordersData.TotalPages)
		if err != nil {
			return nil, err
		}
	}

	if err := json.Unmarshal(orderTableDataJSON, &ordersData.OrderTableData); err != nil {
		return nil, err
	}

	return &ordersData, nil
}
