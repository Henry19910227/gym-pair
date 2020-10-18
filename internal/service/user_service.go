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

// GetAll Implement UserService interface
func (us *userService) GetAll() ([]*model.User, error) {
	return us.userRepo.GetAll()
}

// Get Implement UserService interface
func (us *userService) Get(validator *validator.UserGetValidator) (*model.User, error) {
	return us.userRepo.GetByID(validator.ID)
}

// Add Implement UserService interface
func (us *userService) Add(validator *validator.UserAddValidator) (int64, error) {
	return us.userRepo.Add(validator.Name, validator.Email, validator.Age, validator.Salary)
}

// Delete Implement UserService interface
func (us *userService) Delete(validator *validator.UserDeleteValidator) error {
	return us.userRepo.DeleteByID(validator.ID)
}

// Update Implement UserService interface
func (us *userService) Update(validator *validator.UserUpdateValidator) (*model.User, error) {
	return us.userRepo.Update(validator.ID, validator.Name, validator.Email, validator.Age, validator.Salary)
}
