package service

import (
	"github.com/Henry19910227/gym-pair/internal/model"
	"github.com/Henry19910227/gym-pair/internal/validator"
)

// UserService ...
type UserService interface {
	GetAll() ([]*model.User, error)
	Get(validator *validator.UserGetValidator) (*model.User, error)
	Add(validator *validator.UserAddValidator) (int64, error)
	Delete(validator *validator.UserDeleteValidator) error
	Update(validator *validator.UserUpdateValidator) (*model.User, error)
}

type UploadService interface {
}
