package repository

import (
	"database/sql"

	"github.com/Henry19910227/gym-pair/model"
)

type userRepository struct {
	db *sql.DB
}

// NewUserRepository 創建一個 UserRepository
func NewUserRepository(conn *sql.DB) UserRepository {
	return &userRepository{conn}
}

// GetById ...
func (ur *userRepository) GetByID(id int64) model.User {
	return model.User{}
}
