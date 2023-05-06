package storage

import (
	"github.com/Dbone29/dflow/internal/config"
	"go.uber.org/zap"
)

type Storage interface {
	UploadFile(objectName string, path string, contentType string, data []byte) error
	DownloadFile(path string) ([]byte, error)
}

func InitStorage(logger *zap.Logger, s3StorageConfig *config.S3StorageConfig, localStorageConfig *config.LocalStorageConfig) Storage {
	var store Storage

	if len(s3StorageConfig.Host+s3StorageConfig.BucketName+s3StorageConfig.AccessKeyID+s3StorageConfig.SecretAccessKey) == 0 {
		logger.Info("using local storage")
		store = NewLocalStorage(logger, localStorageConfig)
	} else {
		logger.Info("using s3 storage")
		store = NewS3Storage(logger, s3StorageConfig)
	}

	return store
}
