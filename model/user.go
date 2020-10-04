package model

// User ...
type User struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Userinfo *Userinfo `json:"userinfo" binding:"required"`
}
