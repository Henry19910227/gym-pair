package upload

import (
	"io"
	"mime/multipart"
)

// Upload ...
type Upload interface {
	UploadImage(file multipart.File, filename string) (string, error)
	CheckUploadImageAllowExt(ext string) bool
	CheckUploadImageMaxSize(file io.Reader) bool
}

// UploadSetting ...
type UploadSetting interface {
	GetUploadSavePath() string
	GetUploadImageAllowExts() []string
	GetUploadImageMaxSize() int
}
