package controller

import (
	"net/http"
	"strconv"
	"webDemo/utils"

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
	http.HandleFunc("/gympair/user", userController.GetByID)
}

// GetByID ...
func (c *UserController) GetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	uidParam, ok := r.URL.Query()["id"]
	if !ok || len(uidParam[0]) == 0 {
		w.Write(utils.JSONData(400, nil, "缺少 name 參數"))
		return
	}
	uid, err := strconv.Atoi(uidParam[0])
	if err != nil {
		w.Write(utils.JSONData(400, nil, "請輸入數字"))
		return
	}
	user, err := c.UserService.GetByID(int64(uid))
	if err != nil {
		w.Write(utils.JSONData(200, nil, "查無此id"))
		return
	}
	w.Write(utils.JSONData(200, user, "成功!"))
}
