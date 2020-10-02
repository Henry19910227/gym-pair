package repository

import (
	"database/sql"
	"errors"

	"github.com/Henry19910227/gym-pair/model"
)

type userRepository struct {
	db *sql.DB
}

// NewUserRepository 創建一個 UserRepository
func NewUserRepository(conn *sql.DB) UserRepository {
	return &userRepository{conn}
}

// GetAll ...
func (ur *userRepository) GetAll() ([]*model.User, error) {
	query := "SELECT id,name,email FROM users"
	rows, err := ur.db.Query(query)
	if err != nil {
		return nil, err
	}
	users := []*model.User{}
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err == nil {
			users = append(users, &user)
		}
	}
	return users, nil
}

// GetById ...
func (ur *userRepository) GetByID(id int64) (*model.User, error) {
	query := "SELECT id,name,email FROM users WHERE id = ?"
	row := ur.db.QueryRow(query, id)
	user := model.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, err
}

// Add ...
func (ur *userRepository) Add(user *model.User) (int64, error) {
	query := "INSERT INTO users (name,email) VALUES (?,?)"
	res, err := ur.db.Exec(query, user.Name, user.Email)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// DeleteByID ...
func (ur *userRepository) DeleteByID(id int64) error {
	query := "DELETE FROM users WHERE id = ?"
	res, err := ur.db.Exec(query, id)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("查無此用戶")
	}
	return nil
}

// Update ...
func (ur *userRepository) Update(user *model.User) (*model.User, error) {
	query := "UPDATE users SET name = ?,email = ? WHERE id = ?"
	_, err := ur.db.Exec(query, user.Name, user.Email, user.ID)
	if err != nil {
		return nil, err
	}
	return ur.GetByID(user.ID)
}
