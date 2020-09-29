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
func (ur *userRepository) GetByID(id int64) (model.User, error) {
	query := "SELECT name, age FROM user WHERE id = ?"
	row := ur.db.QueryRow(query, id)
	user := model.User{}
	err := row.Scan(&user.Name, &user.Age)
	if err != nil {
		return model.User{}, err
	}
	return user, err
}
