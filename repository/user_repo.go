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
		user := &model.User{}
		userinfo := &model.Userinfo{}
		var nullAge sql.NullInt64
		var nullSalary sql.NullInt64
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &nullAge, &nullSalary); err == nil {
			// 解決 Left Join 查詢時 userinfo 有可能是空值的狀況
			if nullAge.Valid && nullSalary.Valid {
				userinfo.Age = int(nullAge.Int64)
				userinfo.Salary = int(nullSalary.Int64)
				user.Userinfo = userinfo
			} else {
				user.Userinfo = nil
			}
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
	user := &model.User{}
	userinfo := &model.Userinfo{}
	var nullAge sql.NullInt64
	var nullSalary sql.NullInt64
	if err := row.Scan(&user.ID, &user.Name, &user.Email, &nullAge, &nullSalary); err != nil {
		return nil, err
	}
	// 解決 Left Join 查詢時 userinfo 有可能是空值的狀況
	if nullAge.Valid && nullSalary.Valid {
		userinfo.Age = int(nullAge.Int64)
		userinfo.Salary = int(nullSalary.Int64)
		user.Userinfo = userinfo
	} else {
		user.Userinfo = nil
	}
	return user, nil
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
