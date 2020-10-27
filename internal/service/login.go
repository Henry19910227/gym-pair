package service

import (
	"github.com/Henry19910227/gym-pair/internal/model"
	"github.com/Henry19910227/gym-pair/internal/repository"
)

type loginService struct {
	userRepo repository.UserRepository
}

// NewLoginService ...
func NewLoginService(repo repository.UserRepository) LoginService {
	return &loginService{repo}
}

func (ls *loginService) Login(email string, password string) (*model.User, error) {
	return ls.userRepo.GetUser(email, password)
}
