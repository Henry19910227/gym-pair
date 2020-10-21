package service

import (
	"errors"
	"mime/multipart"
	"path"

	"github.com/Henry19910227/gym-pair/pkg/upload"
)

// GPUploadService ...
type GPUploadService struct {
	uploader upload.Upload
}

// NewUploadService ...
func NewUploadService(uploader upload.Upload) *GPUploadService {
	return &GPUploadService{uploader}
}

// UploadImage ...
func (service *GPUploadService) UploadImage(file multipart.File, filename string) (string, error) {
	// if !service.uploader.CheckUploadImageMaxSize(file) {
	// 	return "", errors.New("exceeded maximum file limit")
	// }
	if !service.uploader.CheckUploadImageAllowExt(path.Ext(filename)) {
		return "", errors.New("image ext is not allow")
	}
	filename, err := service.uploader.UploadImage(file, filename)
	if err != nil {
		return "", err
	}
	return filename, nil
}
