package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	controller "github.com/nchikkam/context-findr-be/controllers"
	_ "github.com/nchikkam/context-findr-be/docs"
	"github.com/nchikkam/context-findr-be/middleware"
	"github.com/nchikkam/context-findr-be/utils"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setUpConfigurations(engine *gin.Engine) {
	// Enable X-Forwarded-For header processing
	engine.ForwardedByClientIP = utils.ForwardedByClientIP
	engine.SetTrustedProxies(utils.Proxies[:])
	engine.MaxMultipartMemory = utils.FileUploadSizeLimit
}

func setUpRoutes(engine *gin.Engine) {
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	engine.GET("/", controller.Home)
	engine.POST("/register", controller.Register)
	engine.POST("/signin", controller.Signin)

	authorized := engine.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.POST("/api/v1/upload", controller.FileUpload)
		authorized.GET("/api/v1/uploads", controller.FileUploads)
		authorized.GET("/api/v1/search", controller.Search)
	}
}

func configCors(engine *gin.Engine) {
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

}

func SetUpServer() *gin.Engine {
	engine := gin.Default()
	configCors(engine)

	utils.ConnectDB()
	setUpConfigurations(engine)
	setUpRoutes(engine)

	return engine
}
