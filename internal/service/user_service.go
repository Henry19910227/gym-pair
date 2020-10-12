package service

import (
	"github.com/Henry19910227/gym-pair/internal/model"
	"github.com/Henry19910227/gym-pair/internal/repository"
	"github.com/Henry19910227/gym-pair/internal/validator"
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

func (us *userService) Add(user *validator.UserAddValidator) (int64, error) {
	return us.userRepo.Add(user.Name, user.Email, user.Age, user.Salary)
}

func (us *userService) DeleteByID(id int64) error {
	return us.userRepo.DeleteByID(id)
}

func (us *userService) Update(user *validator.UserUpdateValidator) (*model.User, error) {
	return us.userRepo.Update(user.ID, user.Name, user.Email, user.Age, user.Salary)
}
