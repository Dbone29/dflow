package storage

type DflowStorage interface {
	UploadFile(objectName string, path string, contentType string, data []byte) error
	DownloadFile(path string) ([]byte, error)
}
