package storage

import (
	"context"
	"io"

	"github.com/Dbone29/dflow/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
)

type S3Storage struct {
	bucketName  string
	minioClient *minio.Client
	logger      *zap.Logger
}

func NewS3Storage(logger *zap.Logger, config *config.StorageConfig) *S3Storage {
	ctx := context.Background()

	minioClient, err := minio.New(config.Host, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKeyID, config.SecretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		logger.Panic("Failed to connect to s3 bucket", zap.Error(err))
	}

	err = minioClient.MakeBucket(ctx, config.BucketName, minio.MakeBucketOptions{Region: "us-east-1"})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, config.BucketName)
		if errBucketExists == nil && exists {
			logger.Info("We already own %s\n", zap.String("Bucket", config.BucketName))
		} else {
			logger.Panic(errBucketExists.Error(), zap.Error(err))
		}
	} else {
		logger.Info("Successfully created %s\n", zap.String("Bucket", config.BucketName))
	}
	return &S3Storage{
		bucketName:  config.BucketName,
		minioClient: minioClient,
		logger:      logger,
	}
}

func (s *S3Storage) UploadFile(objectName string, path string, contentType string, data []byte, reader io.Reader, objectSize int64) error {
	ctx := context.Background()

	_, err := s.minioClient.PutObject(ctx, s.bucketName, objectName, reader, objectSize, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		s.logger.Error("Failed to upload file to S3", zap.String("path", path), zap.Error(err))
		return err
	}

	return nil
}

func (s *S3Storage) DownloadFile(objectName string, path string) ([]byte, error) {
	return nil, nil
}
