package main

import (
	"database/sql"
	"log"

	"github.com/Henry19910227/gym-pair/pkg/jwt"

	"github.com/Henry19910227/gym-pair/global"
	"github.com/Henry19910227/gym-pair/internal/controller"
	"github.com/Henry19910227/gym-pair/internal/middleware"
	"github.com/Henry19910227/gym-pair/internal/repository"
	"github.com/Henry19910227/gym-pair/internal/service"
	"github.com/Henry19910227/gym-pair/pkg/db"
	"github.com/Henry19910227/gym-pair/pkg/logger"
	"github.com/Henry19910227/gym-pair/pkg/upload"
	"github.com/gin-gonic/gin"
)

var (
	mysqlDB      *sql.DB
	userService  service.UserService
	loginService service.LoginService
	jwtTool      jwt.Tool
)

func init() {
	setupLogger()
	setupDB()
	setupLoginService()
	setupUserService()
	setupTokenTool()
}

func main() {
	router := gin.New()
	router.Use(gin.CustomRecovery(middleware.Recover())) //加入攔截panic中間層
	router.Use(gin.Logger())                             //加入路由Logger
	router.Use(middleware.Cors())                        //加入解決跨域中間層
	controller.NewUserController(router, userService, jwtTool)
	controller.NewLoginController(router, loginService, jwtTool)

	router.Run("127.0.0.1:9090")
}

func setupDB() {
	setting, err := db.NewMysqlSetting()
	if err != nil {
		log.Fatalf(err.Error())
	}
	mysqlDB = db.NewDB(setting)
}

func setupLogger() {
	setting, err := logger.NewGPLogSetting("./config/config.yaml")
	if err != nil {
		log.Fatalf(err.Error())
	}
	logger, err := logger.NewGPLogger(setting)
	if err != nil {
		log.Fatalf(err.Error())
	}
	global.Log = logger
}

func setupLoginService() {
	setting, err := upload.NewUploadSetting("./config/config.yaml")
	if err != nil {
		log.Fatalf(err.Error())
	}
	userService = service.NewUserService(repository.NewUserRepository(mysqlDB), upload.NewUploadTool(setting))
}

func setupUserService() {
	loginService = service.NewLoginService(repository.NewUserRepository(mysqlDB))
}

func setupTokenTool() {
	setting, err := jwt.NewJWTSetting("./config/config.yaml")
	if err != nil {
		log.Fatalf(err.Error())
	}
	jwtTool = jwt.NewJWTTool(setting)
}
