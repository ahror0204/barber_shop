package api

import (
	v1 "github.com/barber_shop/api-gateway/api/handlers/v1"
	config "github.com/barber_shop/api-gateway/config"
	"github.com/barber_shop/api-gateway/pkg/logger"
	services "github.com/barber_shop/api-gateway/services"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	_ "github.com/barber_shop/api-gateway/api/docs" // for swagger
)

// Option ...
type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
}

// @title           Swagger for barber shop api
// @version         1.0
// @host      localhost:8000
func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
	})

	api := router.Group("/v1")
	api.POST("/users", handlerV1.CreateUser)
	

	
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	return router
}
