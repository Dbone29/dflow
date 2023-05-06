package storage

import (
	"github.com/Dbone29/dflow/internal/config"
	"github.com/Dbone29/dflow/pkg/storage"
	"go.uber.org/zap"
)

func InitStorage(logger *zap.Logger, s3StorageConfig *config.S3StorageConfig, localStorageConfig *config.LocalStorageConfig) storage.DflowStorage {
	var store storage.DflowStorage

	if len(s3StorageConfig.Host+s3StorageConfig.BucketName+s3StorageConfig.AccessKeyID+s3StorageConfig.SecretAccessKey) == 0 {
		logger.Info("using local storage")
		store = NewLocalStorage(logger, localStorageConfig)
	} else {
		logger.Info("using s3 storage")
		store = NewS3Storage(logger, s3StorageConfig)
	}

	return store
}
