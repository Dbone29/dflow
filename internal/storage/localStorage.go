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

func NewLocalStorage(logger *zap.Logger) *LocalStorage {
	return &LocalStorage{
		basePath: "./files",
		logger:   logger,
	}
}

func (ls *LocalStorage) UploadFile(objectName string, path string, contentType string, data []byte) error {
	fullPath := filepath.Join(ls.basePath, path)

	err := os.MkdirAll(filepath.Dir(fullPath), 0755)
	if err != nil {
		ls.logger.Error("Failed to create directory for file", zap.String("path", fullPath), zap.Error(err))
		return err
	}

	file, err := os.Create(fullPath)
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

func (ls *LocalStorage) DownloadFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}
