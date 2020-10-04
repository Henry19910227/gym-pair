package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RecoverMiddleware ...
type RecoverMiddleware struct {
}

// NewRecoverMiddleware ...
func NewRecoverMiddleware() *RecoverMiddleware {
	return &RecoverMiddleware{}
}

// Recover 攔截panic事件
func (rm *RecoverMiddleware) Recover(c *gin.Context, recovered interface{}) {
	// 獲取自定義的panic
	if str, ok := recovered.(string); ok {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "data": nil, "msg": str})
		return
	}
	// 獲取系統的panic
	c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "data": nil, "msg": "發生不知名錯誤!"})
}
