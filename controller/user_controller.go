package controller

import (
	"net/http"

	"github.com/Henry19910227/gym-pair/service"
)

// UserController ...
type UserController struct {
	UserService service.UserService
}

// NewUserController ...
func NewUserController(s service.UserService) {
	userController := &UserController{
		UserService: s,
	}
	http.HandleFunc("/getUserByID", userController.GetByID)
}

// GetByID ...
func (c *UserController) GetByID(w http.ResponseWriter, r *http.Request) {

}
