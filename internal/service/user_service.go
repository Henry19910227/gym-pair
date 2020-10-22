package service

import (
	"errors"
	"mime/multipart"
	"path"

	"github.com/Henry19910227/gym-pair/internal/model"
	"github.com/Henry19910227/gym-pair/internal/repository"
	"github.com/Henry19910227/gym-pair/internal/validator"
	"github.com/Henry19910227/gym-pair/pkg/upload"
)

type userService struct {
	userRepo repository.UserRepository
	uploader upload.Upload
}

// NewUserService ...
func NewUserService(repo repository.UserRepository, uploader upload.Upload) UserService {
	return &userService{repo, uploader}
}

// GetAll Implement UserService interface
func (us *userService) GetAll() ([]*model.User, error) {
	return us.userRepo.GetAll()
}

// Get Implement UserService interface
func (us *userService) Get(id int64) (*model.User, error) {
	return us.userRepo.GetByID(id)
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
func (us *userService) Update(id int64, name string, email string, image string, age int, salary int) (*model.User, error) {
	return us.userRepo.Update(id, name, email, image, age, salary)
}

// UploadImage Implement UserService interface
func (us *userService) UploadImage(id int64, file multipart.File, fileHeader *multipart.FileHeader) error {

	if !us.uploader.CheckUploadImageAllowExt(path.Ext(fileHeader.Filename)) {
		return errors.New("image ext is not allow")
	}

	if !us.uploader.CheckUploadImageMaxSize(file) {
		return errors.New("exceeded maximum file limit")
	}

	user, err := us.userRepo.GetByID(id)
	if err != nil {
		return err
	}
	newFilename, err := us.uploader.UploadImage(fileHeader)
	if err != nil {
		return err
	}
	user.Image = newFilename
	if _, err = us.userRepo.Update(user.ID, user.Name, user.Email, user.Image, user.Userinfo.Age, user.Userinfo.Salary); err != nil {
		return err
	}
	return nil
}
