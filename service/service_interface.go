package service

import (
	"github.com/Henry19910227/gym-pair/model"
)

// UserService ...
type UserService interface {
	GetByID(id int64) model.User
}
