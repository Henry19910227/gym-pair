package service

import (
	"github.com/Henry19910227/gym-pair/internal/model"
)

// UserService ...
type UserService interface {
	GetAll() ([]*model.User, error)
	GetByID(id int64) (*model.User, error)
	Add(user *model.User) (int64, error)
	DeleteByID(id int64) error
	Update(user *model.User) (*model.User, error)
}
