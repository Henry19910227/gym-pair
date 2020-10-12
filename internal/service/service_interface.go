package service

import (
	"github.com/Henry19910227/gym-pair/internal/model"
	"github.com/Henry19910227/gym-pair/internal/validator"
)

// UserService ...
type UserService interface {
	GetAll() ([]*model.User, error)
	GetByID(id int64) (*model.User, error)
	Add(user *validator.UserAddValidator) (int64, error)
	DeleteByID(id int64) error
	Update(user *model.User) (*model.User, error)
}
