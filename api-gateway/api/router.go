package api

import (
	v1 "github.com/barber_shop/api-gateway/api/handlers/v1"
	config "github.com/barber_shop/api-gateway/config"
	"github.com/barber_shop/api-gateway/pkg/logger"
	services "github.com/barber_shop/api-gateway/services"
	repo "github.com/barber_shop/api-gateway/storage/repo"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	_ "github.com/barber_shop/api-gateway/api/docs" // for swagger
)

// Option ...
type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	RedisRepo      repo.RedisRepositoryStorage
	ServiceManager services.IServiceManager
}

// @title           Swagger for barber shop api
// @version         1.0
// @host      localhost:9090
// @BasePath  /v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @Security ApiKeyAuth
func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		Redis:          option.RedisRepo,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
	})

	router.Static("/media", "./media")

	api := router.Group("/v1")

	api.POST("/file-upload", handlerV1.AuthMiddleware, handlerV1.UploadFile)
	
	api.POST("/customer/create", handlerV1.AuthMiddleware, handlerV1.CreateCustomer)
	api.PUT("/customer/update/:id", handlerV1.AuthMiddleware, handlerV1.UpdateCustomer)
	api.GET("/customer/get/:id", handlerV1.GetCustomerByID)
	api.GET("/customers/list", handlerV1.GetListCustomers)
	api.DELETE("/customer/delete/:id", handlerV1.AuthMiddleware, handlerV1.DeleteCustomer)
	api.GET("/customer/me", handlerV1.AuthMiddleware, handlerV1.GetCustomerProfile)

	api.POST("/customer/register", handlerV1.RegisterCustomer)
	api.POST("/customer/verify", handlerV1.Verify)
	api.POST("/customer/login", handlerV1.CustomerLogIn)
	api.POST("/customer/forgot-password", handlerV1.ForgotPassword)
	api.POST("/customer/verify-forgot-password", handlerV1.VerifyForgotPassword)
	api.POST("/customer/update-password", handlerV1.AuthMiddleware, handlerV1.UpdateCustomerPassword)

	api.POST("/salon/create", handlerV1.CreateSalon)
	api.PUT("/salon/update/:id", handlerV1.UpdateSalon)
	api.GET("/salon/get/:id", handlerV1.GetSalonByID)
	api.GET("/salons/list", handlerV1.GetListSalons)
	api.DELETE("/salon/delete/:id", handlerV1.DeleteSalon)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
