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
	query := "SELECT users.id,users.name,users.email,userinfo.age,userinfo.salary\n" +
		"FROM users\n" +
		"LEFT JOIN userinfo\n" +
		"ON users.userinfo_id = userinfo.id "
	rows, err := ur.db.Query(query)
	if err != nil {
		return nil, err
	}
	users := []*model.User{}
	for rows.Next() {
		var uid int64
		var name string
		var email string
		var nullAge sql.NullInt64
		var nullSalary sql.NullInt64
		if err := rows.Scan(&uid, &name, &email, &nullAge, &nullSalary); err == nil {
			user := model.NewUser(uid, name, email, nullAge, nullSalary)
			users = append(users, user)
		}
	}
	return users, nil
}

// GetById ...
func (ur *userRepository) GetByID(id int64) (*model.User, error) {
	query := "SELECT users.id,users.name,users.email,userinfo.age,userinfo.salary\n" +
		"FROM users\n" +
		"LEFT JOIN userinfo\n" +
		"ON users.userinfo_id = userinfo.id\n" +
		"WHERE users.id = ?"
	row := ur.db.QueryRow(query, id)

	var uid int64
	var name string
	var email string
	var nullAge sql.NullInt64
	var nullSalary sql.NullInt64
	if err := row.Scan(&uid, &name, &email, &nullAge, &nullSalary); err != nil {
		return nil, err
	}
	return model.NewUser(uid, name, email, nullAge, nullSalary), nil
}

// Add 新增 user 並且增加關聯的 userinfo
func (ur *userRepository) Add(user *model.User) (int64, error) {
	tx, err := ur.db.Begin()
	defer tx.Rollback()
	if err != nil {
		return 0, err
	}
	query := "INSERT INTO userinfo (age,salary) VALUES (?,?)"
	infoRes, err := tx.Exec(query, user.Userinfo.Age, user.Userinfo.Salary)
	if err != nil {
		return 0, err
	}
	infoLastID, err := infoRes.LastInsertId()
	if err != nil {
		return 0, err
	}
	query = "INSERT INTO users (name,email,userinfo_id) VALUES (?,?,?)"
	userRes, err := tx.Exec(query, user.Name, user.Email, infoLastID)
	if err != nil {
		return 0, err
	}
	tx.Commit()
	return userRes.LastInsertId()
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
	query := "UPDATE users\n" +
		"INNER JOIN userinfo ON users.userinfo_id = userinfo.id\n" +
		"SET users.name = ?,users.email = ?,userinfo.age = ?,userinfo.salary = ?\n" +
		"WHERE users.id = ?"
	_, err := ur.db.Exec(query, user.Name, user.Email, user.Userinfo.Age, user.Userinfo.Salary, user.ID)
	if err != nil {
		return nil, err
	}
	return ur.GetByID(user.ID)
}

func getUser() *model.User {
	return &model.User{}
}
