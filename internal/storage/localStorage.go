package storage

import (
	"bytes"
	"io"
	"os"
	"path/filepath"

	"go.uber.org/zap"
)

type LocalStorage struct {
	basePath string
	logger   *zap.Logger
}

func NewLocalStorage(basePath string, logger *zap.Logger) *LocalStorage {
	return &LocalStorage{
		basePath: basePath,
		logger:   logger,
	}
}

func (ls *LocalStorage) UploadFile(objectName string, path string, contentType string, data []byte, reader io.Reader, objectSize int64) error {
	fullPath := filepath.Join(ls.basePath, path)

	err := os.MkdirAll(fullPath, os.ModePerm)
	if err != nil {
		ls.logger.Error("Failed to create directory for file", zap.String("path", fullPath), zap.Error(err))
		return err
	}

	file, err := os.Create(filepath.Join(fullPath, objectName))
	if err != nil {
		ls.logger.Error("Failed to create file", zap.String("path", fullPath), zap.Error(err))
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, bytes.NewReader(data))
	if err != nil {
		ls.logger.Error("Failed to write data to file", zap.String("path", fullPath), zap.Error(err))
		return err
	}

	return nil
}

func (ls *LocalStorage) DownloadFile(objectName string, path string) ([]byte, error) {
	fullPath := filepath.Join(ls.basePath, path, objectName)

	data, err := os.ReadFile(fullPath)
	if err != nil {
		ls.logger.Error("Failed to read file", zap.String("path", fullPath), zap.Error(err))
		return nil, err
	}

	return data, nil
}
