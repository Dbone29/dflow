package storage

import (
	"github.com/Dbone29/dflow/internal/config"
	"go.uber.org/zap"
)

type Storage interface {
	UploadFile(path string, data []byte) error
	DownloadFile(path string) ([]byte, error)
}

func InitStorage(logger *zap.Logger, cfg *config.S3StorageConfig) Storage {
	var store Storage

	if len(cfg.Host+cfg.BucketName+cfg.AccessKeyID+cfg.SecretAccessKey) == 0 {
		logger.Info("using local storage")

		// TODO: use local storage
	} else {
		logger.Info("using s3 storage")

		// TODO: use s3 storage
	}

	return store
}
