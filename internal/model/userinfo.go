package model

// Userinfo ...
type Userinfo struct {
	Age    int `json:"age" binding:"required"`
	Salary int `json:"salary" binding:"required"`
}
