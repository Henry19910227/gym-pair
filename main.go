package main

import (
	"github.com/Henry19910227/gym-pair/controller"
	"github.com/Henry19910227/gym-pair/db"
	"github.com/Henry19910227/gym-pair/repository"
	"github.com/Henry19910227/gym-pair/service"
	"github.com/gin-gonic/gin"
)

func main() {
	db := db.NewDB()
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	router := gin.Default()
	controller.NewUserController(router, userService)

	router.Run("127.0.0.1:9090")
}
