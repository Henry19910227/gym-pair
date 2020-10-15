package main

import (
	"database/sql"
	"log"

	"github.com/Henry19910227/gym-pair/global"
	"github.com/Henry19910227/gym-pair/internal/controller"
	"github.com/Henry19910227/gym-pair/internal/middleware"
	"github.com/Henry19910227/gym-pair/internal/repository"
	"github.com/Henry19910227/gym-pair/internal/service"
	"github.com/Henry19910227/gym-pair/pkg/db"
	"github.com/Henry19910227/gym-pair/pkg/logger"
	"github.com/gin-gonic/gin"
)

var (
	mysqlDB       *sql.DB
	userService   service.UserService
	uploadService service.UploadService
	recoverMidd   = middleware.NewRecoverMiddleware()
	settingMidd   = middleware.NewSettingMiddleware()
)

func init() {
	setupLogger()
	setupDB()
	setupUserService()
	setupUploadService()
}

func main() {
	router := gin.New()
	router.Use(gin.CustomRecovery(recoverMidd.Recover)) //加入攔截panic中間層
	router.Use(settingMidd.Cors)                        //加入解決跨域中間層
	router.Use(gin.Logger())                            //加入路由Logger
	controller.NewUserController(router, userService)
	controller.NewUploadController(router, &uploadService)

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

func setupUserService() {
	userService = service.NewUserService(repository.NewUserRepository(mysqlDB))
}

func setupUploadService() {
	uploadService = service.NewUploadService()
}
