package service

import (
	"bytes"
	"context"
	"file-server/internal/utils"
	"io"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/gofiber/fiber/v2"
)

type fileService struct {
	bucketName string
	client     *s3.Client
}

func NewFileService(bucket string) FileService {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("unable to load AWS config, " + err.Error())
	}

	return &fileService{
		bucketName: bucket,
		client:     s3.NewFromConfig(cfg),
	}
}

func (fs fileService) DeleteFile(fileName string) (int, error) {
	_, err := fs.client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: &fs.bucketName,
		Key:    &fileName,
	})
	if err != nil {
		return fiber.StatusInternalServerError, err
	}
	return fiber.StatusOK, nil
}

func (fs fileService) UploadFiles(files []*multipart.FileHeader) (int, error) {
	for _, file := range files {
		err := fs.uploadToS3(file)
		if err != nil {
			return fiber.StatusInternalServerError, err
		}
	}
	return fiber.StatusCreated, nil
}

func (fs fileService) uploadToS3(file *multipart.FileHeader) error {
	content, err := fs.ConvertToBytes(file)
	if err != nil {
		return err
	}

	contentType := file.Header.Get("Content-Type")

	_, err = fs.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      &fs.bucketName,
		Key:         aws.String(file.Filename),
		Body:        bytes.NewReader(content),
		ContentType: aws.String(contentType),
		ACL:         types.ObjectCannedACLPublicRead, 
	})
	if err != nil {
		return utils.ErrFailedToSaveFile
	}

	return nil
}


func (fs fileService) SaveFile(file *multipart.FileHeader, filePath string) error {
	bytes, err := fs.ConvertToBytes(file)
	if err != nil {
		return err
	}

	if err := os.WriteFile(filePath, bytes, 0666); err != nil {
		return utils.ErrFailedToSaveFile
	}

	return nil
}

func (fs fileService) ConvertToBytes(file *multipart.FileHeader) ([]byte, error) {
	src, err := file.Open()
	if err != nil {
		return nil, utils.ErrFailedToOpenFile
	}
	defer src.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, src); err != nil {
		return nil, utils.ErrFailedToCopyFile
	}

	return buf.Bytes(), nil
}
