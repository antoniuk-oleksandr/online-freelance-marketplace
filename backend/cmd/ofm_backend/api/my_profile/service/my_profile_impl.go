package service

import (
	"database/sql"
	"errors"
	fileService "ofm_backend/cmd/ofm_backend/api/file/service"
	file_service "ofm_backend/cmd/ofm_backend/api/file/service"
	"ofm_backend/cmd/ofm_backend/api/my_profile/dto"
	"ofm_backend/cmd/ofm_backend/api/my_profile/helpers"
	"ofm_backend/cmd/ofm_backend/api/my_profile/mapper"
	"ofm_backend/cmd/ofm_backend/api/my_profile/model"
	"ofm_backend/cmd/ofm_backend/api/my_profile/repository"
	"ofm_backend/cmd/ofm_backend/enum"
	"ofm_backend/cmd/ofm_backend/utils"
	"os"
	"strconv"
)

type myProfileService struct {
	myProfileRepository repository.MyProfileRepository
	fileService         fileService.FileService
	serviceFee          float32
}

func NewMyProfileService(repo repository.MyProfileRepository, fileService file_service.FileService) MyProfileService {
	serviceFeeStr := os.Getenv("SERVICE_FEE")
	serviceFee, err := strconv.ParseFloat(serviceFeeStr, 32)
	if err != nil {
		panic("Invalid SERVICE_FEE environment variable: " + serviceFeeStr)
	}

	return &myProfileService{
		myProfileRepository: repo,
		fileService:         fileService,
		serviceFee:          float32(serviceFee),
	}
}

func (mps *myProfileService) GetMyProfileOrders(
	params *dto.MyProfileParams,
) (*dto.OrdersResponse, error) {
	var err error
	params.Limit, err = helpers.GetOrdersPerPage()
	if err != nil {
		return nil, err
	}

	params.Offset = helpers.CalcMyProfileLimit(params)

	ordersDataModel, err := mps.myProfileRepository.GetMyProfileOrders(params)
	if err != nil {
		return nil, err
	}

	ordersReponse := mapper.MapOrdersDataModelToOrdersResponseDto(ordersDataModel)

	return ordersReponse, nil
}

func (mps *myProfileService) GetMyProfileServices(
	params *dto.MyProfileParams,
) (*dto.ServicesResponse, error) {
	var err error
	params.Limit, err = helpers.GetServicesPerPage()
	if err != nil {
		return nil, err
	}

	params.Offset = helpers.CalcMyProfileLimit(params)

	servicesDataModel, err := mps.myProfileRepository.GetMyProfileServices(params)
	if err != nil {
		return nil, err
	}

	ordersReponse := mapper.MapServicesDataModelToServicesResponseDto(servicesDataModel)

	return ordersReponse, nil
}

func (mps *myProfileService) GetMyProfileRequests(params *dto.MyProfileParams) (*dto.RequestsResponse, error) {
	var err error
	params.Limit, err = helpers.GetRequestsPerPage()
	if err != nil {
		return nil, err
	}

	params.Offset = helpers.CalcMyProfileLimit(params)

	requestsDataModel, err := mps.myProfileRepository.GetMyProfileRequests(params)
	if err != nil {
		return nil, err
	}

	ordersReponse := mapper.MapRequestsDataModelToRequestsResponseDto(requestsDataModel)

	return ordersReponse, nil
}

func (mps *myProfileService) GetMyProfileChatByOrderId(orderId int, userId int) (*dto.OrderChat, error) {
	orderChatModel, err := mps.myProfileRepository.GetMyProfileChatByOrderId(orderId, userId)
	if err != nil {
		return nil, err
	}

	return mapper.MapOrderChatDataModelToDto(orderChatModel), nil
}

func (mps *myProfileService) GetMyProfileOverviewByOrderId(
	orderId int, userId int,
) (*dto.OrderOverview, error) {
	orderOverviewModel, err := mps.myProfileRepository.GetMyProfileOverviewByOrderId(orderId, mps.serviceFee, userId)
	if err != nil {
		return nil, err
	}

	orderOverviewDto := mapper.MapMyProfileOrderOverviewModelToDto(orderOverviewModel)

	return &orderOverviewDto, nil
}

func (mps *myProfileService) GetMyProfileRequirementsByOrderId(
	orderId int, userId int,
) ([]model.OrderQuestionAnswer, error) {
	return mps.myProfileRepository.GetMyProfileRequirementsByOrderId(orderId, userId)
}

func (mps *myProfileService) GetMyProfileDeliveryByOrderId(
	orderId int, userId int,
) (*model.OrderDelivery, error) {
	deliveryData, err := mps.myProfileRepository.GetMyProfileDeliveryByOrderId(orderId, userId, mps.serviceFee)
	if err != nil {
		return nil, err
	}

	if deliveryData.Status == enum.Completed {
		deliveryData.DeliveryData.Files = *utils.AddServerURLToFiles(&deliveryData.DeliveryData.Files)
	}

	return deliveryData, nil
}

func (mps *myProfileService) CompleteOrderWithDelivery(
	orderId int, userId int, orderDeliveryBody model.OrderDeliveryBody,
) (*model.OrderDelivery, error) {
	utils.RenameFilesWithUUID(orderDeliveryBody.Files)
	tx, err := mps.myProfileRepository.CreateTransaction()
	if err != nil {
		return nil, err
	}

	deliveryId, err := mps.myProfileRepository.AddOrderDelivery(tx, orderId, userId, orderDeliveryBody.Message)
	if err != nil {
		mps.myProfileRepository.RollbackTransaction(tx)
		return nil, err
	}

	if err := mps.myProfileRepository.CompleteMyProfileOrderById(tx, orderId, userId); err != nil {
		mps.myProfileRepository.RollbackTransaction(tx)
		return nil, err
	}

	filesIds, err := mps.fileService.SaveFilesMetaData(tx, orderDeliveryBody.Files)
	if err != nil {
		mps.myProfileRepository.RollbackTransaction(tx)
		return nil, err
	}

	var deliveryFiles = helpers.CreateDeliveryFiles(deliveryId, filesIds)
	if err := mps.myProfileRepository.AddDeliveryFiles(tx, deliveryFiles); err != nil {
		mps.myProfileRepository.RollbackTransaction(tx)
		return nil, err
	}

	if err := mps.fileService.UploadFiles(orderDeliveryBody.Files); err != nil {
		mps.myProfileRepository.RollbackTransaction(tx)
		return nil, err
	}

	if err := mps.myProfileRepository.CommitTransaction(tx); err != nil {
		return nil, err
	}

	return mps.GetMyProfileDeliveryByOrderId(orderId, userId)
}

func (mps *myProfileService) AddOrderReview(reviewRequestBody model.ReviewRequestBody) (*model.OrderReview, error) {
	tx, err := mps.myProfileRepository.CreateTransaction()
	if err != nil {
		return nil, err
	}

	reviewId, err := mps.myProfileRepository.AddOrderReview(tx, reviewRequestBody)
	if err != nil {
		mps.myProfileRepository.RollbackTransaction(tx)
		return nil, err
	}
	reviewRequestBody.ReviewId = reviewId

	err = mps.myProfileRepository.UpdateOrderReview(tx, reviewRequestBody)
	if err != nil {
		mps.myProfileRepository.RollbackTransaction(tx)
		return nil, err
	}

	err = mps.myProfileRepository.CommitTransaction(tx)
	if err != nil {
		return nil, err
	}

	responseReview, err := mps.myProfileRepository.GetOrderReviewByOrderId(reviewRequestBody.OrderId)
	if err != nil {
		return nil, err
	}

	responseReview.Customer.Avatar = *utils.AddServerURLToFiles(&responseReview.Customer.Avatar)
	return responseReview, nil
}

func (mps *myProfileService) GetOrderReviewByOrderId(orderId int) (*model.OrderReview, error) {
	orderReview, err := mps.myProfileRepository.GetOrderReviewByOrderId(orderId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			status, err := mps.myProfileRepository.GetOrderStatusByOrderId(orderId)
			if err != nil {
				return nil, err
			}

			return &model.OrderReview{
				OrderStatus: status,
			}, sql.ErrNoRows
		}

		return nil, err
	}

	orderReview.Customer.Avatar = *utils.AddServerURLToFiles(&orderReview.Customer.Avatar)

	return orderReview, nil
}

func (mps *myProfileService) UpdateOrderStatusByOrderId(orderId int, userId int, statusId int) error {
	return mps.myProfileRepository.ChangeOrderStatusByOrderId(orderId, userId, statusId)
}
