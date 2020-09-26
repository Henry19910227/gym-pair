package main

import (
	"net/http"

	"github.com/Henry19910227/gym-pair/controller"
	"github.com/Henry19910227/gym-pair/db"
	"github.com/Henry19910227/gym-pair/repository"
	"github.com/Henry19910227/gym-pair/service"
)

func main() {
	db := db.NewDB()
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	controller.NewUserController(userService)

	http.ListenAndServe("127.0.0.1:9090", nil)

}
