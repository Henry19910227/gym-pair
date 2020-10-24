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
	return us.userRepo.Add(validator.Email, validator.Password, validator.Name, validator.Birthday)
}

// Delete Implement UserService interface
func (us *userService) Delete(validator *validator.UserDeleteValidator) error {
	return us.userRepo.DeleteByID(validator.ID)
}

func (us *userService) UpdateUserinfo(uid int64, name string, birthday string) (*model.User, error) {
	return us.userRepo.UpdateUserinfo(uid, name, birthday)
}

func (us *userService) UpdateEmail(uid int64, email string) (*model.User, error) {
	return us.userRepo.UpdateEmail(uid, email)
}

// UpdatePassword ...
func (us *userService) UpdatePassword(uid int64, oldpwd string, newpwd string) error {
	user, err := us.userRepo.GetByID(uid)
	if err != nil {
		return err
	}
	if user.Password != oldpwd {
		return errors.New("與舊密碼不一致")
	}
	return us.userRepo.UpdatePassword(uid, newpwd)
}

// UploadImage Implement UserService interface
func (us *userService) UploadImage(id int64, file multipart.File, fileHeader *multipart.FileHeader) error {

	if !us.uploader.CheckUploadImageAllowExt(path.Ext(fileHeader.Filename)) {
		return errors.New("image ext is not allow")
	}

	if !us.uploader.CheckUploadImageMaxSize(file) {
		return errors.New("exceeded maximum file limit")
	}
	newFilename, err := us.uploader.UploadImage(fileHeader)
	if err != nil {
		return err
	}
	if _, err = us.userRepo.UpdateUserImage(id, newFilename); err != nil {
		return err
	}
	return nil
}
