package middleware

import (
	"github.com/gin-gonic/gin"
)

// SettingMiddleware ...
type SettingMiddleware struct{}

// NewSettingMiddleware ...
func NewSettingMiddleware() *SettingMiddleware {
	return &SettingMiddleware{}
}

// CORS 解決前端跨域問題
func (sm *SettingMiddleware) Cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
}
