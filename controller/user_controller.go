package controller

import (
	"strconv"

	"github.com/Henry19910227/gym-pair/service"
	"github.com/gin-gonic/gin"
)

// UserController ...
type UserController struct {
	UserService service.UserService
}

// NewUserController ...
func NewUserController(router *gin.Engine, s service.UserService) {
	userController := &UserController{
		UserService: s,
	}
	router.GET("/gympair/user/:id", userController.GetByID)
	router.POST("/gympair/user", userController.Insert)
}

// GetByID 以 uid 查找用戶
func (uc *UserController) GetByID(c *gin.Context) {
	uid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"code": 200, "data": nil, "msg": "請輸入數字!"})
		return
	}
	user, err := uc.UserService.GetByID(int64(uid))
	if err != nil {
		c.JSON(400, gin.H{"code": 200, "data": nil, "msg": "查無此用戶!"})
		return
	}
	c.JSON(200, gin.H{"code": 200, "data": user, "msg": "success!"})
}

// Insert 新增用戶
func (uc *UserController) Insert(c *gin.Context) {

}
