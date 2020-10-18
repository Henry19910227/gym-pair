package upload

import (
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strings"
)

// GPUpload ...
type GPUpload struct {
	setting UploadSetting
}

// NewGPUpload ...
func NewGPUpload(setting UploadSetting) *GPUpload {
	return &GPUpload{setting}
}

// UploadImage Implement Upload interface
func (upload *GPUpload) UploadImage(file multipart.File, filename string) (string, error) {
	path, err := upload.createUploadSavePath()
	if err != nil {
		return "", err
	}
	out, err := os.Create(path + "/" + filename)
	if err != nil {
		return "", err
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		return "", err
	}
	return filename, nil
}

// CheckUploadImageAllowExt Implement Upload interface
func (upload *GPUpload) CheckUploadImageAllowExt(ext string) bool {
	ext = strings.ToUpper(ext)
	for _, v := range upload.setting.GetUploadImageAllowExts() {
		if ext == strings.ToUpper(v) {
			return true
		}
	}
	return false
}

// CheckUploadImageMaxSize Implement Upload interface
func (upload *GPUpload) CheckUploadImageMaxSize(file io.Reader) bool {
	content, _ := ioutil.ReadAll(file)
	size := len(content)
	return size <= upload.setting.GetUploadImageMaxSize()*1024*1024
}

func (upload *GPUpload) createUploadSavePath() (string, error) {
	err := os.MkdirAll(upload.setting.GetUploadSavePath(), os.ModePerm)
	if err != nil {
		return "", err
	}
	return upload.setting.GetUploadSavePath(), nil
}
