package repository

import "github.com/Henry19910227/gym-pair/internal/model"

// UserRepository ...
type UserRepository interface {
	GetAll() ([]*model.User, error)
	GetByID(id int64) (*model.User, error)
	Add(name string, email string, age int, salary int) (int64, error)
	DeleteByID(id int64) error
	Update(user *model.User) (*model.User, error)
}
