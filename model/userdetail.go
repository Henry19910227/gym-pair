package model

// UserDetail 詳細用戶資訊
type UserDetail struct {
	User
	Userinfo *Userinfo `json:"userinfo" binding:"required"`
}
