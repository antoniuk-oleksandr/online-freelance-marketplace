package mapper

import (
	"ofm_backend/cmd/ofm_backend/api/my_profile/dto"
	"ofm_backend/cmd/ofm_backend/api/my_profile/model"
	"ofm_backend/cmd/ofm_backend/utils"
)

func MapOrdersDataModelToOrdersResponseDto(ordersDataModel *model.OrdersData) *dto.OrdersResponse {
	orderTableDataDtos := mapOrderTableDataModelsToDtos(ordersDataModel.OrderTableData)

	return &dto.OrdersResponse{
		Orders:     orderTableDataDtos,
		TotalPages: ordersDataModel.TotalPages,
	}
}

func MapServicesDataModelToServicesResponseDto(servicesDataModel *model.ServicesData) *dto.ServicesResponse {
	serviceTableDataDtos := mapServiceTableDataModelsToDtos(servicesDataModel.ServiceTableData)

	return &dto.ServicesResponse{
		Services:   serviceTableDataDtos,
		TotalPages: servicesDataModel.TotalPages,
	}
}

func MapRequestsDataModelToRequestsResponseDto(requestsDataModel *model.RequestsData) *dto.RequestsResponse {
	requestsTableDataDtos := mapRequestTableDataModelsToDtos(requestsDataModel.RequestTableData)

	return &dto.RequestsResponse{
		Requests:   requestsTableDataDtos,
		TotalPages: requestsDataModel.TotalPages,
	}
}

func mapServiceTableDataModelsToDtos(serviceTableDataModels []model.ServiceTableData) []dto.ServiceTableData {
	var serviceTableDataDtos []dto.ServiceTableData = make([]dto.ServiceTableData, 0)

	for _, serviceTableDataModel := range serviceTableDataModels {
		serviceTableDataDtos = append(serviceTableDataDtos, dto.ServiceTableData{
			Id:          serviceTableDataModel.Id,
			Image:       *utils.AddServerURLToFiles[*string](&serviceTableDataModel.Image),
			Title:       serviceTableDataModel.Title,
			Price:       serviceTableDataModel.Price,
			Category:    serviceTableDataModel.Category,
			Rating:      serviceTableDataModel.Rating,
			OrdersCount: serviceTableDataModel.OrdersCount,
			Date:        serviceTableDataModel.Date,
		})
	}

	return serviceTableDataDtos
}

func mapRequestTableDataModelsToDtos(requestTableDataModels []model.RequestTableData) []dto.RequestTableData {
	var requestTableDataDtos []dto.RequestTableData = make([]dto.RequestTableData, 0)

	for _, requestTableDataModel := range requestTableDataModels {
		requestTableDataDtos = append(requestTableDataDtos, dto.RequestTableData{
			Id:                requestTableDataModel.Id,
			Title:             requestTableDataModel.Title,
			Status:            requestTableDataModel.Status,
			Price:             requestTableDataModel.Price,
			CustomerFirstName: requestTableDataModel.CustomerFirstName,
			CustomerSurname:   requestTableDataModel.CustomerSurname,
			CustomerAvatar:    *utils.AddServerURLToFiles[*string](&requestTableDataModel.CustomerAvatar),
			Date:              requestTableDataModel.Date,
		})
	}

	return requestTableDataDtos
}

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
