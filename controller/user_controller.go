package controller

import (
	"net/http"
	"strconv"

	"github.com/Henry19910227/gym-pair/model"
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
	v1 := router.Group("/gympair/v1")
	v1.GET("/user", userController.GetAll)
	v1.GET("/user/:id", userController.GetByID)
	v1.POST("/user", userController.Add)
	v1.DELETE("/user/:id", userController.RemoveByID)
	v1.PUT("/user", userController.UpdateByID)
	v1.GET("/panic", userController.PanicTest)
}

// GetAll 列出所有用戶
func (uc *UserController) GetAll(c *gin.Context) {
	users, err := uc.UserService.GetAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": users, "msg": "success!"})
}

// GetByID 以 uid 查找單個用戶
func (uc *UserController) GetByID(c *gin.Context) {
	uid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": "請輸入數字!"})
		return
	}
	user, err := uc.UserService.GetByID(int64(uid))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": "查無此用戶!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": user, "msg": "success!"})
}

// Add 新增用戶
func (uc *UserController) Add(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": "錯誤的json格式!"})
		return
	}
	uid, err := uc.UserService.Add(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": uid, "msg": "新增成功!"})
}

// RemoveByID 以 uid 刪除用戶
func (uc *UserController) RemoveByID(c *gin.Context) {
	uid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": "請輸入數字!"})
		return
	}
	if err := uc.UserService.DeleteByID(int64(uid)); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": nil, "msg": "刪除成功!"})
}

// UpdateByID 以 uid 更新用戶資料
func (uc *UserController) UpdateByID(c *gin.Context) {
	var user model.User
	// ShouldBindJSON 解析json至model, 並且驗證欄位
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": "錯誤的json格式!"})
		return
	}
	userRes, err := uc.UserService.Update(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": "更新失敗!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": userRes, "msg": "update success!"})
}

// PanicTest 測試 Panic
func (uc *UserController) PanicTest(c *gin.Context) {
	// panic("PanicTest!!!!!")
	var dict map[string]string //不能只有聲明就開始使用
	dict["H"] = "Hello"
}
