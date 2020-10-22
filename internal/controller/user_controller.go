package controller

import (
	"net/http"

	"github.com/Henry19910227/gym-pair/internal/service"
	"github.com/Henry19910227/gym-pair/internal/validator"
	"github.com/gin-gonic/gin"
)

// UserController ...
type UserController struct {
	UserService service.UserService
}

// NewUserController ...
func NewUserController(router *gin.Engine, userService service.UserService) {
	userController := &UserController{
		UserService: userService,
	}
	v1 := router.Group("/gympair/v1")
	v1.GET("/user", userController.GetAll)
	v1.GET("/user/:id", userController.Get)
	v1.POST("/user", userController.Add)
	v1.DELETE("/user/:id", userController.DeleteByID)
	v1.PUT("/user", userController.UpdateByID)
	v1.PUT("/user/:id/image", userController.UploadImage)
	v1.StaticFS("/userimage", http.Dir("./storege"))
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

// Get 以 uid 查找單個用戶
func (uc *UserController) Get(c *gin.Context) {
	var validator validator.UserGetValidator
	if err := c.ShouldBindUri(&validator); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	user, err := uc.UserService.Get(validator.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": "查無此用戶!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": user, "msg": "success!"})
}

// Add 新增用戶
func (uc *UserController) Add(c *gin.Context) {
	var user validator.UserAddValidator
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	uid, err := uc.UserService.Add(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": uid, "msg": "新增成功!"})
}

// DeleteByID 以 uid 刪除用戶
func (uc *UserController) DeleteByID(c *gin.Context) {
	validator := validator.UserDeleteValidator{}
	if err := c.ShouldBindUri(&validator); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	if err := uc.UserService.Delete(&validator); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": nil, "msg": "刪除成功!"})
}

// UpdateByID 以 uid 更新用戶資料
func (uc *UserController) UpdateByID(c *gin.Context) {
	var validator validator.UserUpdateValidator
	// ShouldBindJSON 解析json至model, 並且驗證欄位
	if err := c.ShouldBindJSON(&validator); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	userRes, err := uc.UserService.Update(validator.ID, validator.Name, validator.Email, "", validator.Age, validator.Salary)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": "更新失敗!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": userRes, "msg": "update success!"})
}

// UploadImage 用戶上傳照片
func (uc *UserController) UploadImage(c *gin.Context) {
	var validator validator.UserImageValidator
	if err := c.ShouldBindUri(&validator); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	file, fileHeader, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	if err = uc.UserService.UploadImage(validator.ID, file, fileHeader); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": nil, "msg": "upload success!"})
}

// PanicTest 測試 Panic
func (uc *UserController) PanicTest(c *gin.Context) {
	var dict map[string]string
	dict["H"] = "Hello"
}
