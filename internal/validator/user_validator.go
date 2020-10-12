package validator

// UserAddValidator ...
type UserAddValidator struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Age    int    `json:"age" binding:"required"`
	Salary int    `json:"salary" binding:"required"`
}
