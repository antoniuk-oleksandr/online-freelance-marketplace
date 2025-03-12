package helpers

import (
	"mime/multipart"
	"ofm_backend/cmd/ofm_backend/api/order/body"
	"ofm_backend/cmd/ofm_backend/api/order/model"
	"path/filepath"

	"github.com/google/uuid"
)

func MakeOrderAnswerData(
	orderId int,
	answers []body.OrderQuestionsAnswer,
) []model.OrderAnswer {
	var orderAnswerData []model.OrderAnswer = make([]model.OrderAnswer, len(answers))

	for index, bodyAnswer := range answers {
		orderAnswerData[index] = model.OrderAnswer{
			OrderId:           orderId,
			Content:           bodyAnswer.Content,
			ServiceQuestionId: bodyAnswer.QuestionId,
		}
	}

	return orderAnswerData
}

func MakeOrderFilesData(orderId int, fileIds []int) []model.OrderFile {
	var orderFilesData []model.OrderFile = make([]model.OrderFile, len(fileIds))

	for index, fileId := range fileIds {
		orderFilesData[index] = model.OrderFile{
			OrderId: orderId,
			FileId:  fileId,
		}
	}

	return orderFilesData
}

func RenameFilesWithUUID(files []*multipart.FileHeader) {
	for _, file := range files {
		file.Filename = generateUUIDFilename(file)
	}
}

func generateUUIDFilename(file *multipart.FileHeader) string {
	fileUuid := uuid.New()
	ext := filepath.Ext(file.Filename)
	return fileUuid.String() + ext
}