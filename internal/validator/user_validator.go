package validator

// UserGetValidator ...
type UserGetValidator struct {
	ID int64 `uri:"id" binding:"required,gte=1"`
}

// UserDeleteValidator ...
type UserDeleteValidator struct {
	ID int64 `uri:"id" binding:"required,gte=1"`
}

// UserAddValidator ...
type UserAddValidator struct {
	Name   string `json:"name" binding:"required"`
	Email  string `json:"email" binding:"required"`
	Age    int    `json:"age" binding:"required"`
	Salary int    `json:"salary" binding:"required"`
}

// UserUpdateValidator ...
type UserUpdateValidator struct {
	ID     int64  `json:"id" binding:"required,gte=1"`
	Name   string `json:"name" binding:"required"`
	Email  string `json:"email" binding:"required"`
	Age    int    `json:"age" binding:"required"`
	Salary int    `json:"salary" binding:"required"`
}

// UserImageValidator ...
type UserImageValidator struct {
	ID int64 `uri:"id" binding:"required,gte=1"`
}
