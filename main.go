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

	_ "github.com/Henry19910227/gym-pair/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

// @title Henry
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	router := gin.New()
	router.Use(gin.CustomRecovery(middleware.Recover()))            //加入攔截panic中間層
	router.Use(gin.Logger())                                        //加入路由Logger
	router.Use(middleware.Cors())                                   //加入解決跨域中間層
	url := ginSwagger.URL("http://localhost:9090/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	controller.NewUserController(router, userService, jwtTool)
	controller.NewLoginController(router, loginService, jwtTool)

	router.Run(":9090")
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
