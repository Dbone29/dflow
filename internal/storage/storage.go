package storage

import (
	"os"

	"github.com/Dbone29/dflow/internal/config"
	"github.com/Dbone29/dflow/pkg/storage"
	"go.uber.org/zap"
)

func InitStorage(logger *zap.Logger, config *config.StorageConfig) storage.DflowStorage {
	var store storage.DflowStorage

	if len(config.Host+config.BucketName+config.AccessKeyID+config.SecretAccessKey) == 0 {
		logger.Info("using local storage")

		const dir = "./files"

		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			logger.Warn("failed to create local storage directory", zap.Error(err))
		}

		store = NewLocalStorage(dir, logger)
	} else {
		logger.Info("using s3 storage")
		store = NewS3Storage(logger, config)
	}

	return store
}
