package repository

import (
	"log"
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

func (mpr *myProfileRepository) GetMyProfileRequests(
	params *dto.MyProfileParams,
) (*model.RequestsData, error) {
	rows, err := mpr.db.Queryx(
		queries.GetMyProfileRequestsQuery, params.UserId, params.Status, params.Offset, params.Limit,
	)
	if err != nil {
		return nil, err
	}

	data, totalPages, err := helpers.ParseMyProfileDataFromRows[model.RequestTableData](rows)
	if err != nil {
		return nil, err
	}

	return &model.RequestsData{
		RequestTableData: *data,
		TotalPages:       totalPages,
	}, nil
}

func (mpr *myProfileRepository) GetMyProfileChatByOrderId(orderId int) (*model.OrderChat, error) {
	rows, err := mpr.db.Queryx(queries.GetMyProfileChatByOrderIdQuery, orderId)
	if err != nil {
		return nil, err
	}
	
	orderChat, err := helpers.ParseMyProfileChatByOrderIdFromRows(rows)
	if err != nil {
		log.Println("errx", err)
		return nil, err
	}
	
	return orderChat, nil
}