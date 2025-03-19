package repository

import (
	"ofm_backend/cmd/ofm_backend/api/my_profile/dto"
	"ofm_backend/cmd/ofm_backend/api/my_profile/model"

	"github.com/jmoiron/sqlx"
)

type MyProfileRepository interface {
	GetMyProfileOrders(params *dto.OrdersPaginationParams) (*model.OrdersData, error)
	ParseOrdersDataFromRows(rows *sqlx.Rows) (*model.OrdersData, error)
}
