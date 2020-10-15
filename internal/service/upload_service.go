package service

// GPUploadService ...
type GPUploadService struct {
}

// NewUploadService ...
func NewUploadService() UploadService {
	return &GPUploadService{}
}
