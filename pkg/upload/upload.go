package upload

import (
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"

	"github.com/Henry19910227/gym-pair/utils"
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
func (upload *GPUpload) UploadImage(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	dst, err := upload.createUploadSavePath()
	if err != nil {
		return "", err
	}
	filename := getFileName(path.Ext(fileHeader.Filename))
	out, err := os.Create(dst + "/" + filename)
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

func getFileName(ext string) string {
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	return utils.EncodeMD5(timeStr) + ext
}
