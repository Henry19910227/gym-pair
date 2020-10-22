package upload

import (
	"io"
	"mime/multipart"
)

// Upload ...
type Upload interface {
	UploadImage(fileHeader *multipart.FileHeader) (string, error)
	CheckUploadImageAllowExt(ext string) bool
	CheckUploadImageMaxSize(file io.Reader) bool
}

// UploadSetting ...
type UploadSetting interface {
	GetUploadSavePath() string
	GetUploadImageAllowExts() []string
	GetUploadImageMaxSize() int
}
