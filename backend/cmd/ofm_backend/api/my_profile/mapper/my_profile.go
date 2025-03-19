package mapper

import (
	"ofm_backend/cmd/ofm_backend/api/my_profile/dto"
	"ofm_backend/cmd/ofm_backend/api/my_profile/model"
	"ofm_backend/cmd/ofm_backend/utils"
)

func mapOrderTableDataModelsToDtos(orderTableDataModels []model.OrderTableData) []dto.OrderTableData {
	var orderTableDataDtos []dto.OrderTableData = make([]dto.OrderTableData, 0)

	for _, orderTableDataModel := range orderTableDataModels {
		orderTableDataDtos = append(orderTableDataDtos, dto.OrderTableData{
			Id:     orderTableDataModel.Id,
			Title:  orderTableDataModel.Title,
			Status: orderTableDataModel.Status,
			Price:  orderTableDataModel.Price,
			Date:   orderTableDataModel.Date,
			Image:  *utils.AddServerURLToFiles[*string](&orderTableDataModel.Image),
		})
	}

	return orderTableDataDtos
}

func MapOrdersDataModelToOrdersResponseDto(ordersDataModel *model.OrdersData) *dto.OrdersResponse {
	orderTableDataDtos := mapOrderTableDataModelsToDtos(ordersDataModel.OrderTableData)

	return &dto.OrdersResponse{
		Orders:     orderTableDataDtos,
		TotalPages: ordersDataModel.TotalPages,
	}
}
