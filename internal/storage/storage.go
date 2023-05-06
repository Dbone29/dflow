package storage

import (
	"github.com/Dbone29/dflow/internal/config"
	"go.uber.org/zap"
)

type Storage interface {
	UploadFile(objectName string, path string, contentType string, data []byte) error
	DownloadFile(path string) ([]byte, error)
}

func InitStorage(logger *zap.Logger, config *config.StorageConfig) Storage {
	var store Storage

	if len(config.Host+config.BucketName+config.AccessKeyID+config.SecretAccessKey) == 0 {
		logger.Info("using local storage")
		store = NewLocalStorage(logger)
	} else {
		logger.Info("using s3 storage")
		store = NewS3Storage(logger, config)
	}

	return store
}
