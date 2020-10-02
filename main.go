package main

import (
	"github.com/Henry19910227/gym-pair/controller"
	"github.com/Henry19910227/gym-pair/db"
	"github.com/Henry19910227/gym-pair/middleware"
	"github.com/Henry19910227/gym-pair/repository"
	"github.com/Henry19910227/gym-pair/service"
	"github.com/gin-gonic/gin"
)

func main() {
	db := db.NewDB()
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	recoverMidd := middleware.NewRecoverMiddleware()
	settingMidd := middleware.NewSettingMiddleware()

	router := gin.New()
	router.Use(gin.CustomRecovery(recoverMidd.Recover)) //加入攔截panic中間層
	router.Use(settingMidd.Cors)                        //加入解決跨域中間層
	controller.NewUserController(router, userService)

	router.Run("127.0.0.1:9090")
}
