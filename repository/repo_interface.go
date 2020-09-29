package repository

import "github.com/Henry19910227/gym-pair/model"

// UserRepository ...
type UserRepository interface {
	GetByID(id int64) (model.User, error)
}
