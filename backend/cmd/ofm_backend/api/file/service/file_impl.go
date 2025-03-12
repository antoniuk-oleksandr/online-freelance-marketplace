package service

import (
	"bytes"
	"mime/multipart"
	"ofm_backend/cmd/ofm_backend/api/file/helpers"
	"ofm_backend/cmd/ofm_backend/api/file/repository"
	"ofm_backend/cmd/ofm_backend/utils"

	"github.com/gofiber/fiber/v2"
)

type fileService struct {
	fileRepository repository.FileRepository
}

func NewFileService(fileRepository repository.FileRepository) FileService {
	return &fileService{
		fileRepository: fileRepository,
	}
}

func (fileService *fileService) DeleteFile(fileId int) error {
	panic("unimplemented")
}

func (fileService *fileService) UploadFiles(files []*multipart.FileHeader) error {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	err := helpers.WriteMultipartFiles(files, writer)
	if err != nil {
		return err
	}

	resp, err := helpers.MakeFileServiceRequest(buf, writer)
	if err != nil {
		return err
	}

	if resp.StatusCode != fiber.StatusOK {
		return utils.ErrFailedFileUploadRequest
	}

	return nil
}

func (fileService *fileService) SaveFilesMetaData(files []*multipart.FileHeader) ([]int, error) {
	fileData := helpers.MakeFileData(files)

	fileIds, err := fileService.fileRepository.AddFiles(fileData)
	if err != nil {
		return nil, err
	}

	return fileIds, nil
}
