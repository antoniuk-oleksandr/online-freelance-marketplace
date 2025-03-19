package repository

import (
	"ofm_backend/cmd/ofm_backend/api/my_profile/dto"
	"ofm_backend/cmd/ofm_backend/api/my_profile/helpers"
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
	params *dto.MyProfileParams,
) (*model.OrdersData, error) {
	rows, err := mpr.db.Queryx(
		queries.GetMyProfileOrdersQuery, params.UserId, params.Offset, params.Limit,
	)
	if err != nil {
		return nil, err
	}

	data, totalPages, err := helpers.ParseMyProfileDataFromRows[model.OrderTableData](rows)
	if err != nil {
		return nil, err
	}

	return &model.OrdersData{
		OrderTableData: *data,
		TotalPages:     totalPages,
	}, nil
}

func (mpr *myProfileRepository) GetMyProfileServices(
	params *dto.MyProfileParams,
) (*model.ServicesData, error) {
	rows, err := mpr.db.Queryx(
		queries.GetMyProfileServicesQuery, params.UserId, params.Offset, params.Limit,
	)
	if err != nil {
		return nil, err
	}

	data, totalPages, err := helpers.ParseMyProfileDataFromRows[model.ServiceTableData](rows)
	if err != nil {
		return nil, err
	}

	return &model.ServicesData{
		ServiceTableData: *data,
		TotalPages:       totalPages,
	}, nil
}
