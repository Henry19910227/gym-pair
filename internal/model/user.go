package model

import "database/sql"

// User ...
type User struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Userinfo *Userinfo `json:"userinfo"`
}

// NewUser ...
func NewUser(uid int64, name string, email string, nullAge sql.NullInt64, nullSalary sql.NullInt64) *User {

	if nullAge.Valid && nullSalary.Valid {
		return &User{
			ID:    uid,
			Name:  name,
			Email: email,
			Userinfo: &Userinfo{
				Age:    int(nullAge.Int64),
				Salary: int(nullSalary.Int64),
			},
		}
	}
	return &User{
		ID:    uid,
		Name:  name,
		Email: email,
	}
}
