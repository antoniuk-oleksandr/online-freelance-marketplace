package service

import (
	"ofm_backend/cmd/ofm_backend/api/file/service"
	"ofm_backend/cmd/ofm_backend/api/order/body"
	"ofm_backend/cmd/ofm_backend/api/order/dto"
	"ofm_backend/cmd/ofm_backend/api/order/helpers"
	"ofm_backend/cmd/ofm_backend/api/order/mapper"
	"ofm_backend/cmd/ofm_backend/api/order/repository"
	"ofm_backend/cmd/ofm_backend/enum"
	"ofm_backend/cmd/ofm_backend/utils"
)

type orderService struct {
	orderRepository repository.OrderRepository
	fileService     service.FileService
}

func NewOrderService(
	orderRepository repository.OrderRepository,
	fileService service.FileService,
) OrderService {
	return &orderService{
		orderRepository: orderRepository,
		fileService:     fileService,
	}
}

func (os *orderService) SubmitOrderRequirements(
	orderRequirementsBody *body.OrderRequirementsBody,
) error {
	tx, err := os.orderRepository.CreateTransaction()
	if err != nil {
		return err
	}

	err = os.orderRepository.UpdateOrderCustomerMessageAndStatus(
		orderRequirementsBody.CustomerMessage, orderRequirementsBody.OrderId, enum.AwaitingAcceptance,
	)
	if err != nil {
		os.orderRepository.RollbackTransaction(tx)
		return err
	}

	orderAnswerData := helpers.MakeOrderAnswerData(orderRequirementsBody.OrderId, orderRequirementsBody.Answers)
	err = os.orderRepository.AddOrderQuestionAnswers(orderAnswerData)
	if err != nil {
		os.orderRepository.RollbackTransaction(tx)
		return err
	}

	if len(orderRequirementsBody.Files) > 0 {
		if err := os.AttachOrderFiles(orderRequirementsBody); err != nil {
			os.orderRepository.RollbackTransaction(tx)
		}
	}

	os.orderRepository.CommitTransaction(tx)
	return nil
}

func (os *orderService) AttachOrderFiles(orderRequirementsBody *body.OrderRequirementsBody) error {
	utils.RenameFilesWithUUID(orderRequirementsBody.Files)

	fileIds, err := os.fileService.SaveFilesMetaData(nil, orderRequirementsBody.Files)
	if err != nil {
		return err
	}

	orderFilesData := helpers.MakeOrderFilesData(orderRequirementsBody.OrderId, fileIds)
	if err := os.orderRepository.AddOrderFiles(orderFilesData); err != nil {
		return err
	}

	return os.fileService.UploadFiles(orderRequirementsBody.Files)
}

func (os *orderService) GetOrderById(
	id int,
) (*dto.OrderByIdResponse, error) {
	responseModel, err := os.orderRepository.GetOrderById(id)
	if err != nil {
		return nil, err
	}

	responseDto := mapper.MapOrderByIdResponseModelToDto(*responseModel)

	return &responseDto, nil
}

func (os *orderService) CheckIfOrderRequirementsSubmitted(orderId int) (bool, error) {
	return os.orderRepository.CheckIfOrderRequirementsSubmitted(orderId, enum.AwaitingAcceptance)
}
