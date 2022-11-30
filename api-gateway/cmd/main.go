package main

import (
	"fmt"

	"github.com/barber_shop/api-gateway/api"
	config "github.com/barber_shop/api-gateway/config"
	"github.com/barber_shop/api-gateway/pkg/logger"
	services "github.com/barber_shop/api-gateway/services"
	"github.com/gomodule/redigo/redis"
	rds "github.com/barber_shop/api-gateway/storage/redis"
	
)

func main() {
	cfg := config.Load("./")
	log := logger.New(cfg.LogLevel, "api_gateway")

	pool := redis.Pool{

		MaxIdle: 80,

		MaxActive: 12000,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s%s", cfg.RedisHost, cfg.RedisPort))
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
	redisRepo := rds.NewRedisRepo(&pool)
	serviceManager, err := services.NewServiceManager(&cfg)
	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}

	server := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
		RedisRepo:      redisRepo,
		ServiceManager: serviceManager,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}
}
