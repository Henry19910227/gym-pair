package validator

// UserAddValidator ...
type UserAddValidator struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Age    int    `json:"age" binding:"required"`
	Salary int    `json:"salary" binding:"required"`
}

// UserUpdateValidator ...
type UserUpdateValidator struct {
	ID     int64  `json:"id" binding:"required"`
	Name   string `json:"name" binding:"required"`
	Email  string `json:"email" binding:"required"`
	Age    int    `json:"age" binding:"required"`
	Salary int    `json:"salary" binding:"required"`
}
