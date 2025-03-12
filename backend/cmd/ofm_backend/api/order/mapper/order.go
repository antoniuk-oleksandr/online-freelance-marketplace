package mapper

import (
	"ofm_backend/cmd/ofm_backend/api/order/dto"
	"ofm_backend/cmd/ofm_backend/api/order/model"
)

func mapServiceQuestionModelsToDtos(
	serviceQuestionModels []model.FreelanceQuestion,
) []dto.FreelanceQuestion {
	var serviceQuestionDtos = make([]dto.FreelanceQuestion, 0)

	for _, item := range serviceQuestionModels {
		serviceQuestionDtos = append(serviceQuestionDtos, dto.FreelanceQuestion{
			Id:      item.Id,
			Content: item.Content,
		})
	}

	return serviceQuestionDtos
}

func MapOrderByIdResponseModelToDto(
	orderByIdResponseModel model.OrderByIdResponse,
) dto.OrderByIdResponse {
	var serviceQuestions = mapServiceQuestionModelsToDtos(orderByIdResponseModel.ServiceQuestions)

	return dto.OrderByIdResponse{
		Order: dto.Order{
			Id:       orderByIdResponseModel.Order.Id,
			CratedAt: orderByIdResponseModel.Order.CratedAt,
			Status:   orderByIdResponseModel.Order.Status,
		},
		Freelance: dto.Freelance{
			Id:    orderByIdResponseModel.Service.Id,
			Title: orderByIdResponseModel.Service.Title,
			Image: orderByIdResponseModel.Service.Image,
			Package: dto.Package{
				Id:    orderByIdResponseModel.Service.Package.Id,
				Price: orderByIdResponseModel.Service.Package.Price,
				Title: orderByIdResponseModel.Service.Package.Title,
			},
		},
		ServiceQuestions: serviceQuestions,
	}
}
