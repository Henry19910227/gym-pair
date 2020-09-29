package service

import (
	"github.com/Henry19910227/gym-pair/model"
	"github.com/Henry19910227/gym-pair/repository"
)

type userService struct {
	userRepo repository.UserRepository
}

// NewUserService ...
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (us *userService) GetByID(id int64) (model.User, error) {
	return us.userRepo.GetByID(id)
}
