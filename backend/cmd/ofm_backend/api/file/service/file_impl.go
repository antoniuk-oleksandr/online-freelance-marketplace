package service

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"ofm_backend/cmd/ofm_backend/api/file/helpers"
	"ofm_backend/cmd/ofm_backend/api/file/model"
	"ofm_backend/cmd/ofm_backend/api/file/repository"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type fileService struct {
	fileRepository repository.FileRepository
	s3Client       *s3.Client
}

func NewFileService(
	fileRepository repository.FileRepository, s3Client *s3.Client,
) FileService {
	return &fileService{
		fileRepository: fileRepository,
		s3Client:       s3Client,
	}
}

func (fileService *fileService) DeleteFile(fileId int) error {
	panic("unimplemented")
}

func (fileService *fileService) UploadFiles(files []*multipart.FileHeader) error {
	ctx := context.Background()

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			return fmt.Errorf("failed to open file %s: %w", fileHeader.Filename, err)
		}
		defer file.Close()

		_, err = fileService.s3Client.PutObject(ctx, &s3.PutObjectInput{
			Bucket:      aws.String("ofm-s3-bucket"),
			Key:         aws.String(fileHeader.Filename),
			Body:        file,
			ContentType: aws.String(fileHeader.Header.Get("Content-Type")),
		})
		if err != nil {
			return fmt.Errorf("failed to upload file %s: %w", fileHeader.Filename, err)
		}
	}

	return nil
}

func (fileService *fileService) SaveFilesMetaData(tx *sqlx.Tx, files []*multipart.FileHeader) ([]int, error) {
	fileData := helpers.MakeFileData(files)

	fileIds, err := fileService.fileRepository.AddFiles(tx, fileData)
	if err != nil {
		return nil, err
	}

	return fileIds, nil
}

func (fileService *fileService) UploadFromURLWithoutTransaction(tx *sqlx.Tx, picURL string) (int, string, error) {
	ctx := context.Background()

	resp, err := http.Get(picURL)
	if err != nil {
		return -1, "", fmt.Errorf("failed to download file from URL: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return -1, "", fmt.Errorf("non-200 response from URL: %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return -1, "", fmt.Errorf("failed to read response body: %w", err)
	}
	contentLength := int64(len(bodyBytes))

	contentType := resp.Header.Get("Content-Type")
	if contentType == "" || !strings.Contains(contentType, "/") {
		return -1, "", fmt.Errorf("invalid or missing Content-Type")
	}

	parts := strings.SplitN(contentType, "/", 2)
	ext := parts[1]
	if ext == "jpeg" {
		ext = "jpg"
	}
	filename := uuid.New().String() + "." + ext

	// Upload to S3
	_, err = fileService.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:        aws.String("ofm-s3-bucket"),
		Key:           aws.String(filename),
		Body:          bytes.NewReader(bodyBytes),
		ContentType:   aws.String(contentType),
		ContentLength: &contentLength,
	})
	if err != nil {
		return -1, "", fmt.Errorf("failed to upload to S3: %w", err)
	}

	fileIDs, err := fileService.fileRepository.AddFiles(tx, []model.FileData{
		{Name: filename},
	})
	if err != nil || len(fileIDs) == 0 {
		return -1, "", fmt.Errorf("failed to store file in database: %w", err)
	}

	return fileIDs[0], filename, nil
}
