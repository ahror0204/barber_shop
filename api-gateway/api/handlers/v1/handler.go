package v1

import (
	"strconv"

	"github.com/barber_shop/api-gateway/api/models"
	config "github.com/barber_shop/api-gateway/config"
	"github.com/barber_shop/api-gateway/pkg/logger"
	services "github.com/barber_shop/api-gateway/services"
	repo "github.com/barber_shop/api-gateway/storage/repo"
	"github.com/gin-gonic/gin"
)

type handlerV1 struct {
	log            logger.Logger
	redisStorage   repo.RedisRepositoryStorage
	serviceManager services.IServiceManager
	cfg            config.Config
}

// HandlerV1Config ...
type HandlerV1Config struct {
	Logger         logger.Logger
	Redis          repo.RedisRepositoryStorage
	ServiceManager services.IServiceManager
	Cfg            config.Config
}

// New ...
func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		log:            c.Logger,
		redisStorage:   c.Redis,
		serviceManager: c.ServiceManager,
		cfg:            c.Cfg,
	}
}

func validateGetAllParams(c *gin.Context) (*models.GetListParams, error) {
	var (
		limit int64 = 10
		page  int64 = 1
		err   error
	)

	if c.Query("limit") != "" {
		limit, err = strconv.ParseInt(c.Query("limit"), 10, 64)
		if err != nil {
			return nil, err
		}
	}

	if c.Query("page") != "" {
		page, err = strconv.ParseInt(c.Query("page"), 10, 64)
		if err != nil {
			return nil, err
		}
	}

	return &models.GetListParams{
		Limit:  limit,
		Page:   page,
		Search: c.Query("search"),
	}, nil
}

func errorResponse(err error) *models.ErrorResponse {
	return &models.ErrorResponse{
		Error: err.Error(),
	}
}
