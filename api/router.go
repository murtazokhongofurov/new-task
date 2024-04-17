package api

import (
	h "net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/new-task/api/docs" // docs
	v1 "github.com/new-task/api/handler/v1"
	"github.com/new-task/config"
	"github.com/new-task/pkg/logger"
	"github.com/new-task/storage"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Options struct {
	Cfg     config.Config
	Storage storage.StorageI
	Log     logger.Logger
}

// @title           Task API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @BasePath  /v1

// @securityDefinitions.basic  BasicAuth
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func New(opt *Options) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	corConfig := cors.DefaultConfig()
	corConfig.AllowAllOrigins = true
	corConfig.AllowCredentials = true
	corConfig.AllowHeaders = []string{"*"}
	corConfig.AllowBrowserExtensions = true
	corConfig.AllowMethods = []string{"*"}
	router.Use(cors.New(corConfig))

	handlerV1 := v1.New(&v1.HandlerV1Option{
		Cfg:     &opt.Cfg,
		Storage: opt.Storage,
		Log:     opt.Log,
	})
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(h.StatusOK, gin.H{
			"message": "Server is running!!!",
		})
	})
	router.MaxMultipartMemory = 8 << 20 // 8 Mib

	api := router.Group("/v1")

	// user
	user := api.Group("/customer")
	user.POST("", handlerV1.CreateC)
	user.GET("", handlerV1.GetC)
	user.PUT("", handlerV1.UpdateC)
	user.DELETE("", handlerV1.DeleteC)

	// item
	item := api.Group("/item")
	item.POST("", handlerV1.CreateP)
	item.GET("", handlerV1.GetP)
	item.PUT("", handlerV1.UpdateP)
	item.DELETE("", handlerV1.DeleteP)

	// transaction
	tran := api.Group("/transaction")
	tran.POST("", handlerV1.CreateT)
	tran.GET("", handlerV1.GetT)
	tran.PUT("", handlerV1.UpdateT)
	tran.DELETE("", handlerV1.DeleteT)

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
