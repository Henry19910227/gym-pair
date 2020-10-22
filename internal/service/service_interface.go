package service

import (
	"mime/multipart"

	"github.com/Henry19910227/gym-pair/internal/model"
	"github.com/Henry19910227/gym-pair/internal/validator"
)

// UserService ...
type UserService interface {
	GetAll() ([]*model.User, error)
	Get(id int64) (*model.User, error)
	Add(validator *validator.UserAddValidator) (int64, error)
	Delete(validator *validator.UserDeleteValidator) error
	Update(id int64, name string, email string, image string, age int, salary int) (*model.User, error)
	UploadImage(id int64, file multipart.File, fileHeader *multipart.FileHeader) error
}
