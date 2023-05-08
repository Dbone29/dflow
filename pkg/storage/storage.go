package storage

import "io"

type DflowStorage interface {
	UploadFile(objectName string, path string, contentType string, data []byte, reader io.Reader, objectSize int64) error
	DownloadFile(objectName string, path string) ([]byte, error)
}
