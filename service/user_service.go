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

func (us *userService) GetAll() ([]*model.User, error) {
	return us.userRepo.GetAll()
}

func (us *userService) GetByID(id int64) (*model.User, error) {
	return us.userRepo.GetByID(id)
}

func (us *userService) Add(user *model.User) (int64, error) {
	return us.userRepo.Add(user)
}

func (us *userService) DeleteByID(id int64) error {
	return us.userRepo.DeleteByID(id)
}

func (us *userService) Update(user *model.User) (*model.User, error) {
	return us.userRepo.Update(user)
}
